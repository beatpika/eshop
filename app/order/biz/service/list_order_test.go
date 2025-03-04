package service

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/app/order/biz/dal/mysql"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
)

func TestListOrder_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
