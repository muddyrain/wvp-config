package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"wvp-config/router/internal"
)

func create(_ context.Context, c *app.RequestContext) {
	var req internal.Request
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusOK, internal.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	switch req.Architecture {
	case internal.Standalone:
		c.JSON(http.StatusOK, internal.StandaloneConversion(req))
	case internal.StandaloneIntercom:

	default:
		c.JSON(http.StatusOK, internal.Response{
			Code:    http.StatusBadRequest,
			Message: "暂不支持",
		})
	}
}
