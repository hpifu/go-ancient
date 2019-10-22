package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorsReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *Service) GETAuthors(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &AuthorsReq{Limit: 20}

	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	if req.Limit > 50 {
		req.Limit = 50
	}

	authors, err := s.db.SelectAuthors(req.Offset, req.Limit)
	if err != nil {
		return req, nil, http.StatusInternalServerError, fmt.Errorf("mysql select authors failed. err: [%v]", err)
	}

	if authors == nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, authors, http.StatusOK, nil
}
