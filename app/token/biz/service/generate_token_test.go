package service

import (
	"context"
	"testing"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

func TestGenerateToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGenerateTokenService(ctx)
	// init req and assert value

	req := &auth.GenerateTokenRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
