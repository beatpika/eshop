package service

import (
	"context"
	"testing"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

func TestRefreshToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefreshTokenService(ctx)
	// init req and assert value

	req := &auth.RefreshTokenRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
