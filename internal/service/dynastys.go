package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DynastysReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Service) Dynastys(c *gin.Context) {
	var res []string
	var err error
	var buf []byte
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())
	req := &DynastysReq{}

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

	res, err = s.dynastys(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("dynastys failed")
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

func (s *Service) dynastys(req *DynastysReq) ([]string, error) {
	return s.db.SelectDynastys(req.Offset, req.Limit)
}
