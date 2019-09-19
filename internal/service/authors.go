package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthorsReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Service) Authors(c *gin.Context) {
	var res []string
	var err error
	var buf []byte
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())
	req := &AuthorsReq{Limit: 20}

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

	if req.Limit > 50 {
		req.Limit = 50
	}

	res, err = s.authors(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("authors failed")
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

func (s *Service) authors(req *AuthorsReq) ([]string, error) {
	return s.db.SelectAuthors(req.Offset, req.Limit)
}
