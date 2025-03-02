package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/hertz/pkg/app"
)

type RefreshTokenService struct {
	ctx    context.Context
	c      *app.RequestContext
	client authservice.Client
}

func NewRefreshTokenService(ctx context.Context, c *app.RequestContext) *RefreshTokenService {
	return &RefreshTokenService{
		ctx:    ctx,
		c:      c,
		client: rpc.TokenClient,
	}
}

// ForTest 返回一个用于测试的服务实例，允许注入mock客户端
func NewRefreshTokenServiceForTest(ctx context.Context, c *app.RequestContext, client authservice.Client) *RefreshTokenService {
	return &RefreshTokenService{
		ctx:    ctx,
		c:      c,
		client: client,
	}
}

func (s *RefreshTokenService) Run(req *token.RefreshTokenReq) (*token.RefreshTokenResp, error) {
	// 构造RPC请求
	rpcReq := &auth.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	// 调用RPC服务
	rpcResp, err := s.client.RefreshToken(s.ctx, rpcReq)
	if err != nil {
		return &token.RefreshTokenResp{
			Base: &common.BaseResp{
				StatusCode:    500,
				StatusMessage: "Internal server error: " + err.Error(),
			},
		}, err
	}

	// 检查RPC响应中的错误
	if rpcResp.ErrorCode != auth.ErrorCode_ERROR_CODE_UNSPECIFIED {
		return &token.RefreshTokenResp{
			Base: &common.BaseResp{
				StatusCode:    400,
				StatusMessage: rpcResp.ErrorMessage,
			},
		}, nil
	}

	// 构造成功响应
	return &token.RefreshTokenResp{
		Base: &common.BaseResp{
			StatusCode:    200,
			StatusMessage: "Success",
		},
		AccessToken:  rpcResp.AccessToken,
		RefreshToken: rpcResp.RefreshToken,
		ExpiresIn:    rpcResp.ExpiresAt,
	}, nil
}
