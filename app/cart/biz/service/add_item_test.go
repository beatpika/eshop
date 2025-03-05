package service

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func TestAddItem(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		req     *cart.AddItemReq
		wantErr bool
	}{
		{
			name: "normal add",
			req: &cart.AddItemReq{
				UserId: 1,
				Item: &cart.CartItem{
					ProductId: 100,
					Quantity:  2,
				},
			},
			wantErr: false,
		},
		{
			name: "add with zero quantity",
			req: &cart.AddItemReq{
				UserId: 1,
				Item: &cart.CartItem{
					ProductId: 101,
					Quantity:  0,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AddItem(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 验证添加后的商品是否可以获取到
			resp, err := GetCart(ctx, &cart.GetCartReq{UserId: tt.req.UserId})
			if err != nil {
				t.Errorf("Failed to get cart after adding item: %v", err)
				return
			}

			found := false
			for _, item := range resp.Cart.Items {
				if item.ProductId == tt.req.Item.ProductId {
					found = true
					if item.Quantity != tt.req.Item.Quantity {
						t.Errorf("Got quantity = %v, want %v", item.Quantity, tt.req.Item.Quantity)
					}
					break
				}
			}

			if !found && !tt.wantErr {
				t.Errorf("Added item not found in cart")
			}
		})
	}
}
