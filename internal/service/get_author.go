package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorReq struct {
	Author string `uri:"author"`
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
}

func (s *Service) GETAuthor(c *gin.Context) (interface{}, interface{}, int, error) {
	req := &AuthorReq{Limit: 20}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind uri failed. err: [%v]", err)
	}

	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	if req.Limit > 50 {
		req.Limit = 50
	}

	ancients, err := s.db.SelectAncientByAuthor(req.Author, req.Offset, req.Limit)

	if err != nil {
		return req, nil, http.StatusInternalServerError, fmt.Errorf("mysql select ancient failed. err: [%v]", err)
	}

	if ancients == nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, ancients, http.StatusOK, nil
}
