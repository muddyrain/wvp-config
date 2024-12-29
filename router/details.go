package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"wvp-config/router/internal"
)

func details(_ context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, internal.Response{
		Code:    http.StatusOK,
		Message: "获取当前模版详细",
		Data: []internal.TemplateDetails{
			{
				Architecture: internal.Standalone,
				ImageApi:     "/file/standalone.jpg",
			},
			{
				Architecture: internal.StandaloneIntercom,
				ImageApi:     "/file/standaloneIntercom.jpg",
			},
		},
	})
}
