package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DynastyReq struct {
	Dynasty string `uri:"dynasty"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit"`
}

func (s *Service) Dynasty(c *gin.Context) {
	var res []*mysql.Ancient
	var err error
	var buf []byte
	req := &DynastyReq{}
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

	if err := c.Bind(req); err != nil {
		err = fmt.Errorf("bind failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.dynasty(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("dynasty failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	if res == nil {
		status = http.StatusNoContent
		c.Status(status)
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) dynasty(req *DynastyReq) ([]*mysql.Ancient, error) {
	return s.db.SelectAncientByDynasty(req.Dynasty, req.Offset, req.Limit)
}
