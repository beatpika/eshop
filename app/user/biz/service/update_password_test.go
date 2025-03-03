package service

import (
	"context"
	"testing"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestUpdatePassword_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdatePasswordService(ctx)
	// init req and assert value

	req := &user.UpdatePasswordReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
