package service

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/app/order/biz/dal/mysql"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
)

func TestPlaceOrder_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value
	req := &order.PlaceOrderReq{
		UserId:       1,
		UserCurrency: "usa",
		Email:        "demo@demo.com",
		Address: &order.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "USA",
			ZipCode:       12345,
		},
		OrderItems: []*order.OrderItem{
			{
				Cost: 100,
				Item: &cart.CartItem{
					ProductId: 1,
					Quantity:  2,
				},
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// todo: edit your unit test
}
