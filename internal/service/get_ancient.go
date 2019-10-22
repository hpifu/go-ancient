package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AncientReq struct {
	ID int `uri:"id"`
}

func (s *Service) GETAncient(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &AncientReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind uri failed. err: [%v]", err)
	}

	ancient, err := s.db.SelectAncientByID(req.ID)

	if err != nil {
		return req, nil, http.StatusInternalServerError, fmt.Errorf("mysql select ancient failed. err: [%v]", err)
	}

	if ancient == nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, ancient, http.StatusOK, nil
}
