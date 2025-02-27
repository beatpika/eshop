package service

import (
	"context"
	"testing"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

func TestRevokeToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRevokeTokenService(ctx)
	// init req and assert value

	req := &auth.RevokeTokenRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
