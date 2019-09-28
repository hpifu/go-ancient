package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-ancient/internal/es"
	"github.com/hpifu/go-ancient/internal/mysql"
	"github.com/hpifu/go-ancient/internal/service"
	"github.com/hpifu/go-kit/logger"
	"github.com/spf13/viper"
)

// AppVersion name
var AppVersion = "unknown"

func main() {
	version := flag.Bool("v", false, "print current version")
	configfile := flag.String("c", "configs/ancient.json", "config file path")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	// load config
	config := viper.New()
	config.SetEnvPrefix("account")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()
	config.SetConfigType("json")
	fp, err := os.Open(*configfile)
	if err != nil {
		panic(err)
	}
	err = config.ReadConfig(fp)
	if err != nil {
		panic(err)
	}

	// init logger
	infoLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.infoLog"))
	if err != nil {
		panic(err)
	}
	warnLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.warnLog"))
	if err != nil {
		panic(err)
	}
	accessLog, err := logger.NewJsonLoggerWithViper(config.Sub("logger.accessLog"))
	if err != nil {
		panic(err)
	}
	service.InfoLog = infoLog
	service.WarnLog = warnLog
	service.AccessLog = accessLog

	// init mysqldb
	db, err := mysql.NewMysqlDB(config.GetString("mysqldb.uri"))
	if err != nil {
		panic(err)
	}
	infoLog.Infof("init mysqldb success. uri [%v]", config.GetString("mysqldb.uri"))

	// init elasticsearch
	esclient, err := es.NewES(config.GetString("es.uri"))
	if err != nil {
		panic(err)
	}
	infoLog.Infof("init elasticsearch success. uri [%v]", config.GetString("es.uri"))

	secure := config.GetBool("service.cookieSecure")
	domain := config.GetString("service.cookieDomain")
	origin := config.GetString("service.allowOrigin")
	// init services
	svr := service.NewService(db, esclient, secure, domain)

	// init gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// set handler
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	r.GET("/ancient", service.Decorator(svr.Ancients))
	r.GET("/ancient/:id", service.Decorator(svr.Ancient))
	r.GET("/author", service.Decorator(svr.Authors))
	r.GET("/author/:author", service.Decorator(svr.Author))
	r.GET("/dynasty", service.Decorator(svr.Dynastys))
	r.GET("/dynasty/:dynasty", service.Decorator(svr.Dynasty))
	r.GET("/search", service.Decorator(svr.Search))

	infoLog.Infof("%v init success, port [%v]", os.Args[0], config.GetString("service.port"))

	// run server
	if err := r.Run(config.GetString("service.port")); err != nil {
		panic(err)
	}
}
