package service

import (
	"context"
	"testing"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func TestGetCart(t *testing.T) {
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

	// 先清空购物车
	_, err := EmptyCart(ctx, &cart.EmptyCartReq{UserId: userId})
	if err != nil {
		t.Fatalf("Failed to empty cart before test: %v", err)
	}

	// 添加测试商品
	for _, item := range testItems {
		_, err := AddItem(ctx, &cart.AddItemReq{
			UserId: userId,
			Item:   item,
		})
		if err != nil {
			t.Fatalf("Failed to add test item: %v", err)
		}
	}

	// 测试获取购物车
	tests := []struct {
		name    string
		req     *cart.GetCartReq
		wantErr bool
	}{
		{
			name:    "get existing cart",
			req:     &cart.GetCartReq{UserId: userId},
			wantErr: false,
		},
		{
			name:    "get non-existing cart",
			req:     &cart.GetCartReq{UserId: 99999},
			wantErr: false, // 不存在的购物车应该返回空而不是错误
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := GetCart(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.req.UserId == userId {
				// 验证返回的商品数量
				if len(resp.Cart.Items) != len(testItems) {
					t.Errorf("Got %d items, want %d", len(resp.Cart.Items), len(testItems))
				}

				// 验证每个商品的信息
				itemMap := make(map[uint32]int32)
				for _, item := range resp.Cart.Items {
					itemMap[item.ProductId] = item.Quantity
				}

				for _, want := range testItems {
					if got, exists := itemMap[want.ProductId]; !exists || got != want.Quantity {
						t.Errorf("Item %d: got quantity %d, want %d", want.ProductId, got, want.Quantity)
					}
				}
			} else {
				// 对于不存在的购物车，应该返回空列表
				if len(resp.Cart.Items) != 0 {
					t.Errorf("Expected empty cart for non-existing user, got %d items", len(resp.Cart.Items))
				}
			}
		})
	}
}
