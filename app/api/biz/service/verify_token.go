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

type VerifyTokenService struct {
	ctx    context.Context
	c      *app.RequestContext
	client authservice.Client
}

func NewVerifyTokenService(ctx context.Context, c *app.RequestContext) *VerifyTokenService {
	return &VerifyTokenService{
		ctx:    ctx,
		c:      c,
		client: rpc.TokenClient,
	}
}

// ForTest 返回一个用于测试的服务实例，允许注入mock客户端
func NewVerifyTokenServiceForTest(ctx context.Context, c *app.RequestContext, client authservice.Client) *VerifyTokenService {
	return &VerifyTokenService{
		ctx:    ctx,
		c:      c,
		client: client,
	}
}

func (s *VerifyTokenService) Run(req *token.VerifyTokenReq) (*token.VerifyTokenResp, error) {
	// 构造RPC请求
	rpcReq := &auth.VerifyTokenRequest{
		Token: req.Token,
	}

	// 调用RPC服务
	rpcResp, err := s.client.VerifyToken(s.ctx, rpcReq)
	if err != nil {
		return &token.VerifyTokenResp{
			Base: &common.BaseResp{
				StatusCode:    500,
				StatusMessage: "Internal server error: " + err.Error(),
			},
		}, err
	}

	// 检查RPC响应中的错误
	if rpcResp.ErrorCode != auth.ErrorCode_ERROR_CODE_UNSPECIFIED {
		return &token.VerifyTokenResp{
			Base: &common.BaseResp{
				StatusCode:    400,
				StatusMessage: rpcResp.ErrorMessage,
			},
			IsValid: false,
		}, nil
	}

	// 构造成功响应
	return &token.VerifyTokenResp{
		Base: &common.BaseResp{
			StatusCode:    200,
			StatusMessage: "Success",
		},
		IsValid: rpcResp.IsValid,
		UserId:  rpcResp.UserId,
	}, nil
}
