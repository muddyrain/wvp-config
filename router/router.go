package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
	"time"
	"wvp-config/conf"
	"wvp-config/router/internal"
)

func CustomizeRegister(h *server.Hertz) {
	// 1 获取wvp和zlm架构详情
	g := h.Group("/api/v1/")
	g.GET("/details", details)
	// 2 根据预制配置文件生成对应的配置文件
	g.PUT("/standalone", config)
	g.PUT("/standalone-intercom", config)
}

func AppFS() *app.FS {
	return &app.FS{
		Root:        "./static/",
		PathRewrite: app.NewPathSlashesStripper(0),
		PathNotFound: func(_ context.Context, c *app.RequestContext) {
			c.JSON(http.StatusOK, internal.Response{
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

func config(_ context.Context, c *app.RequestContext) {
	var req internal.Request
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusOK, internal.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	switch req.Architecture {
	case conf.Standalone:
		c.JSON(http.StatusOK, internal.StandaloneConversion(req))
	case conf.StandaloneIntercom:
		c.JSON(http.StatusOK, internal.StandaloneIntercomConversion(req))
	default:
		c.JSON(http.StatusOK, internal.Response{
			Code:    http.StatusBadRequest,
			Message: "暂不支持",
		})
	}
}
