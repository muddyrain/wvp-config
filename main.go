package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"time"
	"wvp-config/router"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "0.0.0.0:18080", "address")
	flag.Parse()
	h := server.New(
		server.WithALPN(true),
		server.WithHostPorts(address),
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
	h.StaticFS("/", router.AppFS())
	h.Spin()
}
