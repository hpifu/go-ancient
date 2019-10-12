package service

import (
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
	AccessLog = logrus.New()
}

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
