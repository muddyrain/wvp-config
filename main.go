package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"net/http"
	"time"
)

func main() {
	h := server.New(
		server.WithALPN(true),
		server.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", 8080)),
		server.WithHandleMethodNotAllowed(true),
	)
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(defaultRecoveryHandler))) // 自定义异常处理
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	h.Use(accesslog.New(accesslog.WithTimeFormat(time.DateTime)))
	customizeRegister(h)
	h.StaticFS("/", appFS())
	h.Spin()
}

func customizeRegister(h *server.Hertz) {
	// 1 对比配置文件是否有问题
	h.POST("/config", ConfigComparison)
	// 2 根据预制配置文件生成对应的配置文件
	h.PUT("/config", ConfigCreate)
}

func ConfigCreate(ctx context.Context, c *app.RequestContext) {
	type Options struct {
	}
	type Request struct {
		Mode       int     `json:"mode"`       // 预制的类型 1-单机版wvp 2-单机语音版wvp
		IntranetIP string  `json:"intranetIP"` // 内网IP
		ExternalIP string  `json:"externalIP"` // 外网IP
		Options    Options `json:"options"`    // 可选配置 主要是端口范围等
	}
	var req Request
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, DefaultResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	// 根据预制类型选择模版生成配置文件
	type Response struct {
		Mode   int    `json:"mode"`   // 什么类型的配置 1-wvp 2-zlm 3-nginx
		Config string `json:"config"` // 具体的配置文件
	}
	const (
		wvp = iota + 1
		zlm
		nginx
	)
	response := make([]Response, 0)
	response = append(response, Response{
		Mode:   wvp,
		Config: `{}`,
	})
	response = append(response, Response{
		Mode:   zlm,
		Config: `{}`,
	})
	c.JSON(http.StatusOK, DefaultResponse{
		Message: "请下载对应配置文件",
		Code:    http.StatusOK,
		Data:    response,
	})
	return
}

func ConfigComparison(ctx context.Context, c *app.RequestContext) {
	type Request struct {
		WvpConfig string `json:"wvpConfig,required"`
		ZlmConfig string `json:"zlmConfig,required"`
	}
	var req Request
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, DefaultResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// 序列化配置文件

	// 进行对比
	type Response struct {
		Message      string   `json:"message"`      // 提示具体哪里有问题
		Correct      bool     `json:"correct"`      // 是否正确
		Details      []string `json:"details"`      // 问题详情
		NetworkTitle string   `json:"networkTitle"` // 网络情况提示
	}
	response := Response{
		Message:      "wvp ip异常 推荐为xxx",
		Correct:      false,
		Details:      make([]string, 0),
		NetworkTitle: "需要对外开放xxx端口(TCP+UDP)",
	}
	c.JSON(http.StatusOK, DefaultResponse{
		Message: response.Message,
		Code:    http.StatusOK,
		Data:    response,
	})
	return
}

func appFS() *app.FS {
	return &app.FS{
		Root:        "./static/",
		PathRewrite: app.NewPathSlashesStripper(0),
		PathNotFound: func(_ context.Context, c *app.RequestContext) {
			c.JSON(http.StatusOK, DefaultResponse{
				Message: c.Request.URI().String(),
				Code:    http.StatusNotFound,
			})
		},
		CacheDuration:        time.Second * 5,
		IndexNames:           []string{"*index.html"},
		Compress:             true,
		CompressedFileSuffix: "hertz",
		AcceptByteRange:      true,
	}
}

func defaultRecoveryHandler(c context.Context, ctx *app.RequestContext, err any, stack []byte) {
	hlog.SystemLogger().CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", ctx.Request.Header.UserAgent())
	ctx.AbortWithStatus(hconsts.StatusInternalServerError)
}

type DefaultResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}
