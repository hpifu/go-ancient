package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-ancient/internal/mysql"
	api "github.com/hpifu/go-ancient/pkg/ancient"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Service) GetCategory(c *gin.Context) {
	var res *api.GetCategoryRes
	var err error
	var buf []byte
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())
	req := &api.GetCategoryQuery{}

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

	if err = s.checkGetCategoryReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.getCategory(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("getCategory failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkGetCategoryReqBody(req *api.GetCategoryQuery) error {
	return nil
}

func (s *Service) getCategory(req *api.GetCategoryQuery) (*api.GetCategoryRes, error) {
	ancients, err := s.db.SelectAncientByTitleAndAuthor(&mysql.Ancient{
		Title:   req.Title,
		Dynasty: req.Dynasty,
		Author:  req.Author,
	}, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	var as []*api.Ancient
	for _, ancient := range ancients {
		as = append(as, &api.Ancient{
			ID:      ancient.ID,
			Title:   ancient.Title,
			Author:  ancient.Author,
			Dynasty: ancient.Dynasty,
		})
	}

	return &api.GetCategoryRes{
		OK:       true,
		Ancients: as,
	}, nil
}
