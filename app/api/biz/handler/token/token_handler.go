package token

import (
	"context"

	"github.com/beatpika/eshop/app/api/biz/service"
	"github.com/beatpika/eshop/app/api/biz/utils"
	token "github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GenerateToken .
// @router /token/generate [POST]
func GenerateToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req token.GenerateTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &token.GenerateTokenResp{}
	resp, err = service.NewGenerateTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RefreshToken .
// @router /token/refresh [POST]
func RefreshToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req token.RefreshTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &token.RefreshTokenResp{}
	resp, err = service.NewRefreshTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// VerifyToken .
// @router /token/verify [POST]
func VerifyToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req token.VerifyTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &token.VerifyTokenResp{}
	resp, err = service.NewVerifyTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RevokeToken .
// @router /token/revoke [POST]
func RevokeToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req token.RevokeTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &token.RevokeTokenResp{}
	resp, err = service.NewRevokeTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
