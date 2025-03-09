package utils

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

// Success status code
const (
	StatusOK = 0
)

// Error status codes
const (
	StatusInternalError = 500
)

// BuildBaseResp build base response
func BuildBaseResp(err error) *common.BaseResp {
	if err != nil {
		return &common.BaseResp{
			StatusCode:    StatusInternalError,
			StatusMessage: err.Error(),
		}
	}
	return &common.BaseResp{
		StatusCode:    StatusOK,
		StatusMessage: "success",
	}
}

// SendErrResponse pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	c.JSON(code, BuildBaseResp(err))
}

// SendSuccessResponse pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	c.JSON(code, data)
}
