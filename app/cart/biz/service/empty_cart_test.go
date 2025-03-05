package service

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func TestEmptyCart(t *testing.T) {
	ctx := context.Background()

	// 准备测试数据
	userId := uint32(1)
	testItems := []*cart.CartItem{
		{
			ProductId: 100,
			Quantity:  2,
		},
		{
			ProductId: 101,
			Quantity:  1,
		},
	}

	// 先添加一些商品
	for _, item := range testItems {
		_, err := AddItem(ctx, &cart.AddItemReq{
			UserId: userId,
			Item:   item,
		})
		if err != nil {
			t.Fatalf("Failed to add test item: %v", err)
		}
	}

	tests := []struct {
		name    string
		req     *cart.EmptyCartReq
		wantErr bool
	}{
		{
			name:    "empty existing cart",
			req:     &cart.EmptyCartReq{UserId: userId},
			wantErr: false,
		},
		{
			name:    "empty non-existing cart",
			req:     &cart.EmptyCartReq{UserId: 99999},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行清空操作
			_, err := EmptyCart(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmptyCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 验证购物车是否为空
			resp, err := GetCart(ctx, &cart.GetCartReq{UserId: tt.req.UserId})
			if err != nil {
				t.Errorf("Failed to get cart after emptying: %v", err)
				return
			}

			if len(resp.Cart.Items) != 0 {
				t.Errorf("Cart should be empty after EmptyCart, got %d items", len(resp.Cart.Items))
			}
		})
	}
}
