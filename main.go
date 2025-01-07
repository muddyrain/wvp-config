package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"os"
	"time"
	"wvp-config/conf"
	"wvp-config/router"
)

func init() {
	conf.InitConfig("./config.yaml")
	conf.GetData().Show()

	for _, templateConfig := range conf.GetData().TemplateConfigs {
		_ = os.MkdirAll(templateConfig.Dir, os.ModePerm)
	}
}

func main() {
	h := server.New(
		server.WithALPN(true),
		server.WithHostPorts(conf.GetData().Address),
		server.WithHandleMethodNotAllowed(true),
	)
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	h.Use(accesslog.New(accesslog.WithTimeFormat(time.DateTime)))
	router.CustomizeRegister(h)
	h.StaticFS("/files/", router.AppFS())
	h.Spin()
}
