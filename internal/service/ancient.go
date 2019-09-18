package service

import (
	"fmt"
	"github.com/hpifu/go-ancient/internal/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AncientReq struct {
	ID int `uri:"id"`
}

func (s *Service) Ancient(c *gin.Context) {
	var res *mysql.Ancient
	var err error
	var buf []byte
	req := &AncientReq{}
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())

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

	if err := c.BindUri(req); err != nil {
		err = fmt.Errorf("bind failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.getAncient(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("getAncient failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) getAncient(req *AncientReq) (*mysql.Ancient, error) {
	return s.db.SelectAncientByID(req.ID)
}
