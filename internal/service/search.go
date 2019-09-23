package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SearchReq struct {
	Q      string `form:"q"`
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
}

func (s *Service) Search(c *gin.Context) {
	var res []*es.Ancient
	var err error
	var buf []byte
	req := &SearchReq{}
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

	res, err = s.search(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("search failed")
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

func (s *Service) search(req *SearchReq) ([]*es.Ancient, error) {
	return s.es.SearchAncient(req.Q, req.Offset, req.Limit)
}
