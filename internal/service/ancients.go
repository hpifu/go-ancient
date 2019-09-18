package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
)

type AncientReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Service) Ancient(c *gin.Context) {
	var res []*mysql.Ancient
	var err error
	var buf []byte
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())
	req := &AncientReq{}

	defer func() {
		AccessLog.WithFields(logrus.Fields{
			"host":   c.Request.Host,
			"body":   string(buf),
			"url":    c.Request.URL.String(),
			"req":    req,
			"res":    res,
			"rid":    rid,
			"err":    err,
			"status": status,
		}).Info()
	}()

	if err := c.Bind(req); err != nil {
		err = fmt.Errorf("bind failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}
	fmt.Println(req)

	if err = s.checkAncientReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.ancient(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("ancient failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkAncientReqBody(req *AncientReq) error {
	return nil
}

func (s *Service) ancient(req *AncientReq) ([]*mysql.Ancient, error) {
	return s.db.SelectAncients(req.Offset, req.Limit)
}
