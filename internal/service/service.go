package service

import (
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
)

type Service struct {
	db        *mysql.Mysql
	es        *es.ES
	secure    bool
	domain    string
	infoLog   *logrus.Logger
	warnLog   *logrus.Logger
	accessLog *logrus.Logger
}

func (s *Service) SetLogger(infoLog, warnLog, accessLog *logrus.Logger) {
	s.infoLog = infoLog
	s.warnLog = warnLog
	s.accessLog = accessLog
}

func NewService(
	db *mysql.Mysql,
	es *es.ES,
	secure bool,
	domain string,
) *Service {
	return &Service{
		db:        db,
		es:        es,
		secure:    secure,
		domain:    domain,
		infoLog:   logrus.New(),
		warnLog:   logrus.New(),
		accessLog: logrus.New(),
	}
}
