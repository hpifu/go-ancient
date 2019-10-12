package service

import (
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
)

var InfoLog *logrus.Logger = logrus.New()
var WarnLog *logrus.Logger = logrus.New()
var AccessLog *logrus.Logger = logrus.New()

type Service struct {
	db     *mysql.Mysql
	es     *es.ES
	secure bool
	domain string
}

func NewService(
	db *mysql.Mysql,
	es *es.ES,
	secure bool,
	domain string,
) *Service {
	return &Service{
		db:     db,
		es:     es,
		secure: secure,
		domain: domain,
	}
}
