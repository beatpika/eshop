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

type RevokeTokenService struct {
	ctx    context.Context
	c      *app.RequestContext
	client authservice.Client
}

func NewRevokeTokenService(ctx context.Context, c *app.RequestContext) *RevokeTokenService {
	return &RevokeTokenService{
		ctx:    ctx,
		c:      c,
		client: rpc.TokenClient,
	}
}

// ForTest 返回一个用于测试的服务实例，允许注入mock客户端
func NewRevokeTokenServiceForTest(ctx context.Context, c *app.RequestContext, client authservice.Client) *RevokeTokenService {
	return &RevokeTokenService{
		ctx:    ctx,
		c:      c,
		client: client,
	}
}

func (s *RevokeTokenService) Run(req *token.RevokeTokenReq) (*token.RevokeTokenResp, error) {
	// 构造RPC请求
	rpcReq := &auth.RevokeTokenRequest{
		Token: req.Token,
	}

	// 调用RPC服务
	rpcResp, err := s.client.RevokeToken(s.ctx, rpcReq)
	if err != nil {
		return &token.RevokeTokenResp{
			Base: &common.BaseResp{
				StatusCode:    500,
				StatusMessage: "Internal server error: " + err.Error(),
			},
		}, err
	}

	// 检查RPC响应中的错误
	if rpcResp.ErrorCode != auth.ErrorCode_ERROR_CODE_UNSPECIFIED {
		return &token.RevokeTokenResp{
			Base: &common.BaseResp{
				StatusCode:    400,
				StatusMessage: rpcResp.ErrorMessage,
			},
		}, nil
	}

	// 构造成功响应
	return &token.RevokeTokenResp{
		Base: &common.BaseResp{
			StatusCode:    200,
			StatusMessage: "Success",
		},
	}, nil
}
