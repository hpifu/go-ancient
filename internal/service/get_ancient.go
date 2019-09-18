package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	api "github.com/hpifu/go-ancient/pkg/ancient"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Service) GetAncient(c *gin.Context) {
	var res *api.GetAncientRes
	var err error
	var buf []byte
	status := http.StatusOK
	rid := c.DefaultQuery("rid", NewToken())
	req := &api.GetAncientReq{}

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

	if err = s.checkGetAncientReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
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

func (s *Service) checkGetAncientReqBody(req *api.GetAncientReq) error {
	return nil
}

func (s *Service) getAncient(req *api.GetAncientReq) (*api.GetAncientRes, error) {
	ancient, err := s.db.SelectAncientByID(req.ID)
	if err != nil {
		return nil, err
	}
	if ancient == nil {
		return &api.GetAncientRes{
			OK:      true,
			Ancient: nil,
		}, nil
	}

	return &api.GetAncientRes{
		OK: true,
		Ancient: &api.Ancient{
			ID:      ancient.ID,
			Title:   ancient.Title,
			Author:  ancient.Author,
			Dynasty: ancient.Dynasty,
		},
	}, nil
}
