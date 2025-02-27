package token

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	GenerateToken(ctx context.Context, Req *auth.GenerateTokenRequest, callOptions ...callopt.Option) (r *auth.GenerateTokenResponse, err error)
	VerifyToken(ctx context.Context, Req *auth.VerifyTokenRequest, callOptions ...callopt.Option) (r *auth.VerifyTokenResponse, err error)
	RefreshToken(ctx context.Context, Req *auth.RefreshTokenRequest, callOptions ...callopt.Option) (r *auth.RefreshTokenResponse, err error)
	RevokeToken(ctx context.Context, Req *auth.RevokeTokenRequest, callOptions ...callopt.Option) (r *auth.RevokeTokenResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GenerateToken(ctx context.Context, Req *auth.GenerateTokenRequest, callOptions ...callopt.Option) (r *auth.GenerateTokenResponse, err error) {
	return c.kitexClient.GenerateToken(ctx, Req, callOptions...)
}

func (c *clientImpl) VerifyToken(ctx context.Context, Req *auth.VerifyTokenRequest, callOptions ...callopt.Option) (r *auth.VerifyTokenResponse, err error) {
	return c.kitexClient.VerifyToken(ctx, Req, callOptions...)
}

func (c *clientImpl) RefreshToken(ctx context.Context, Req *auth.RefreshTokenRequest, callOptions ...callopt.Option) (r *auth.RefreshTokenResponse, err error) {
	return c.kitexClient.RefreshToken(ctx, Req, callOptions...)
}

func (c *clientImpl) RevokeToken(ctx context.Context, Req *auth.RevokeTokenRequest, callOptions ...callopt.Option) (r *auth.RevokeTokenResponse, err error) {
	return c.kitexClient.RevokeToken(ctx, Req, callOptions...)
}
