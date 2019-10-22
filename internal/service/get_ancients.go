package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AncientsReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Service) GETAncients(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &AncientsReq{Limit: 20}

	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	if req.Limit > 50 {
		req.Limit = 50
	}

	ancients, err := s.db.SelectAncients(req.Offset, req.Limit)
	if err != nil {
		return req, nil, http.StatusInternalServerError, fmt.Errorf("mysql select ancients failed. err: [%v]", err)
	}

	if ancients == nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, ancients, http.StatusOK, nil
}
