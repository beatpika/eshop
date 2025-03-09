package service

import (
	"context"
	"testing"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
)

func TestUpdateOrderStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderStatusService(ctx)
	// init req and assert value

	req := &order.UpdateOrderStatusReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
