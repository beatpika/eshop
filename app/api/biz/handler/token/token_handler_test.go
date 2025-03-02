package token

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/app/api/biz/service"
	"github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/stretchr/testify/assert"
)

// MockTokenClient 实现 authservice.Client 接口用于测试
type MockTokenClient struct {
	generateTokenFunc func(context.Context, *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error)
	refreshTokenFunc  func(context.Context, *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	verifyTokenFunc   func(context.Context, *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error)
	revokeTokenFunc   func(context.Context, *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error)
}

func (m *MockTokenClient) GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest, opts ...callopt.Option) (*auth.GenerateTokenResponse, error) {
	return m.generateTokenFunc(ctx, req)
}

func (m *MockTokenClient) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest, opts ...callopt.Option) (*auth.RefreshTokenResponse, error) {
	return m.refreshTokenFunc(ctx, req)
}

func (m *MockTokenClient) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, opts ...callopt.Option) (*auth.VerifyTokenResponse, error) {
	return m.verifyTokenFunc(ctx, req)
}

func (m *MockTokenClient) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest, opts ...callopt.Option) (*auth.RevokeTokenResponse, error) {
	return m.revokeTokenFunc(ctx, req)
}

func TestGenerateToken(t *testing.T) {
	tests := []struct {
		name           string
		req            *token.GenerateTokenReq
		rpcResp        *auth.GenerateTokenResponse
		rpcErr         error
		expectedStatus int
		expectedResp   *token.GenerateTokenResp
	}{
		{
			name: "success",
			req: &token.GenerateTokenReq{
				UserId: 123,
			},
			rpcResp: &auth.GenerateTokenResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				ExpiresAt:    3600,
				ErrorCode:    auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.GenerateTokenResp{
				Base: &common.BaseResp{
					StatusCode:    200,
					StatusMessage: "Success",
				},
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				ExpiresIn:    3600,
			},
		},
		{
			name: "error",
			req: &token.GenerateTokenReq{
				UserId: 123,
			},
			rpcResp: &auth.GenerateTokenResponse{
				ErrorCode:    auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
				ErrorMessage: "Invalid request",
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.GenerateTokenResp{
				Base: &common.BaseResp{
					StatusCode:    400,
					StatusMessage: "Invalid request",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock客户端
			mockClient := &MockTokenClient{
				generateTokenFunc: func(ctx context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
					assert.Equal(t, tt.req.UserId, req.UserId)
					return tt.rpcResp, tt.rpcErr
				},
			}

			// 创建测试服务
			ctx := context.Background()
			c := app.NewContext(0)
			svc := service.NewGenerateTokenServiceForTest(ctx, c, mockClient)

			// 执行请求
			resp, err := svc.Run(tt.req)

			// 验证结果
			if tt.rpcErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {
	tests := []struct {
		name           string
		req            *token.RefreshTokenReq
		rpcResp        *auth.RefreshTokenResponse
		rpcErr         error
		expectedStatus int
		expectedResp   *token.RefreshTokenResp
	}{
		{
			name: "success",
			req: &token.RefreshTokenReq{
				RefreshToken: "old_refresh_token",
			},
			rpcResp: &auth.RefreshTokenResponse{
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				ExpiresAt:    3600,
				ErrorCode:    auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.RefreshTokenResp{
				Base: &common.BaseResp{
					StatusCode:    200,
					StatusMessage: "Success",
				},
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				ExpiresIn:    3600,
			},
		},
		{
			name: "error",
			req: &token.RefreshTokenReq{
				RefreshToken: "invalid_token",
			},
			rpcResp: &auth.RefreshTokenResponse{
				ErrorCode:    auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
				ErrorMessage: "Invalid refresh token",
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.RefreshTokenResp{
				Base: &common.BaseResp{
					StatusCode:    400,
					StatusMessage: "Invalid refresh token",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock客户端
			mockClient := &MockTokenClient{
				refreshTokenFunc: func(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
					assert.Equal(t, tt.req.RefreshToken, req.RefreshToken)
					return tt.rpcResp, tt.rpcErr
				},
			}

			// 创建测试服务
			ctx := context.Background()
			c := app.NewContext(0)
			svc := service.NewRefreshTokenServiceForTest(ctx, c, mockClient)

			// 执行请求
			resp, err := svc.Run(tt.req)

			// 验证结果
			if tt.rpcErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	tests := []struct {
		name           string
		req            *token.VerifyTokenReq
		rpcResp        *auth.VerifyTokenResponse
		rpcErr         error
		expectedStatus int
		expectedResp   *token.VerifyTokenResp
	}{
		{
			name: "success",
			req: &token.VerifyTokenReq{
				Token: "valid_token",
			},
			rpcResp: &auth.VerifyTokenResponse{
				IsValid:   true,
				UserId:    123,
				ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.VerifyTokenResp{
				Base: &common.BaseResp{
					StatusCode:    200,
					StatusMessage: "Success",
				},
				IsValid: true,
				UserId:  123,
			},
		},
		{
			name: "error",
			req: &token.VerifyTokenReq{
				Token: "invalid_token",
			},
			rpcResp: &auth.VerifyTokenResponse{
				IsValid:      false,
				ErrorCode:    auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
				ErrorMessage: "Invalid token",
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.VerifyTokenResp{
				Base: &common.BaseResp{
					StatusCode:    400,
					StatusMessage: "Invalid token",
				},
				IsValid: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock客户端
			mockClient := &MockTokenClient{
				verifyTokenFunc: func(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
					assert.Equal(t, tt.req.Token, req.Token)
					return tt.rpcResp, tt.rpcErr
				},
			}

			// 创建测试服务
			ctx := context.Background()
			c := app.NewContext(0)
			svc := service.NewVerifyTokenServiceForTest(ctx, c, mockClient)

			// 执行请求
			resp, err := svc.Run(tt.req)

			// 验证结果
			if tt.rpcErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestRevokeToken(t *testing.T) {
	tests := []struct {
		name           string
		req            *token.RevokeTokenReq
		rpcResp        *auth.RevokeTokenResponse
		rpcErr         error
		expectedStatus int
		expectedResp   *token.RevokeTokenResp
	}{
		{
			name: "success",
			req: &token.RevokeTokenReq{
				Token: "valid_token",
			},
			rpcResp: &auth.RevokeTokenResponse{
				Success:   true,
				ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.RevokeTokenResp{
				Base: &common.BaseResp{
					StatusCode:    200,
					StatusMessage: "Success",
				},
			},
		},
		{
			name: "error",
			req: &token.RevokeTokenReq{
				Token: "invalid_token",
			},
			rpcResp: &auth.RevokeTokenResponse{
				Success:      false,
				ErrorCode:    auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
				ErrorMessage: "Invalid token",
			},
			expectedStatus: consts.StatusOK,
			expectedResp: &token.RevokeTokenResp{
				Base: &common.BaseResp{
					StatusCode:    400,
					StatusMessage: "Invalid token",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock客户端
			mockClient := &MockTokenClient{
				revokeTokenFunc: func(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
					assert.Equal(t, tt.req.Token, req.Token)
					return tt.rpcResp, tt.rpcErr
				},
			}

			// 创建测试服务
			ctx := context.Background()
			c := app.NewContext(0)
			svc := service.NewRevokeTokenServiceForTest(ctx, c, mockClient)

			// 执行请求
			resp, err := svc.Run(tt.req)

			// 验证结果
			if tt.rpcErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
