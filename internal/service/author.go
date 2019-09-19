package service

import (
	"fmt"
	"github.com/hpifu/go-ancient/internal/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthorReq struct {
	Author string `uri:"author"`
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
}

func (s *Service) Author(c *gin.Context) {
	var res []*mysql.Ancient
	var err error
	var buf []byte
	req := &AuthorReq{}
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
		err = fmt.Errorf("bind uri failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	if err := c.Bind(req); err != nil {
		err = fmt.Errorf("bind failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.author(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("author failed")
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

func (s *Service) author(req *AuthorReq) ([]*mysql.Ancient, error) {
	return s.db.SelectAncientByAuthor(req.Author, req.Offset, req.Limit)
}
