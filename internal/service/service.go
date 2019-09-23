package service

import (
	"encoding/hex"
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
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

func NewToken() string {
	buf := make([]byte, 32)
	token := make([]byte, 16)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(token)
	hex.Encode(buf, token)
	return string(buf)
}
