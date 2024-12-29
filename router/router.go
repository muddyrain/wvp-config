package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
	"time"
	"wvp-config/router/internal"
)

func CustomizeRegister(h *server.Hertz) {
	// 1 获取wvp和zlm架构详情
	h.POST("/details", details)
	// 2 根据预制配置文件生成对应的配置文件
	h.PUT("/config", create)
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
