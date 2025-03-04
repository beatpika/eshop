package service

import (
	"context"

	"github.com/beatpika/eshop/app/order/biz/dal/mysql"
	"github.com/beatpika/eshop/app/order/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	// 查询订单
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var orders []*order.Order // 声明返回的订单
	// 遍历订单
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems { // 遍历订单商品
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  int32(oi.Quantity),
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			CreatedAt: int32(v.CreatedAt.Unix()),
			OrderId:   v.OrderId,
			UserId:    v.UserId,
			Email:     v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAdress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.Zipcode,
			},
			OrderItems: items,
		})
	}

	return &order.ListOrderResp{
		Orders: orders,
	}, nil
}
