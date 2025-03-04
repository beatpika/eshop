package service

import (
	"context"

	order "github.com/beatpika/eshop/app/api/hertz_gen/basic/order"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	kitex_order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListOrdersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrdersService(Context context.Context, RequestContext *app.RequestContext) *ListOrdersService {
	return &ListOrdersService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrdersService) Run(req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 创建默认响应
	resp = &order.ListOrdersResp{
		Base: &common.BaseResp{
			StatusCode:    0,
			StatusMessage: "success",
		},
	}

	// 参数验证
	if req.UserId == 0 {
		resp.Base.StatusCode = 400
		resp.Base.StatusMessage = "用户ID不能为空"
		return resp, nil
	}

	// 转换为RPC请求
	rpcReq := &kitex_order.ListOrderReq{
		UserId: req.UserId,
	}

	// 调用RPC服务
	rpcResp, err := rpc.OrderClient.ListOrder(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = 500
		resp.Base.StatusMessage = "获取订单列表失败: " + err.Error()
		return resp, nil
	}

	// 转换RPC响应为API响应
	orders := make([]*order.OrderInfo, 0, len(rpcResp.Orders))
	for _, o := range rpcResp.Orders {
		items := make([]*order.OrderItem, 0, len(o.OrderItems))
		for _, item := range o.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &order.ProductItem{
					ProductId: item.Item.ProductId,
					Quantity:  item.Item.Quantity,
				},
				Cost: item.Cost,
			})
		}

		orders = append(orders, &order.OrderInfo{
			OrderId:   o.OrderId,
			UserId:    o.UserId,
			Email:     o.Email,
			CreatedAt: o.CreatedAt,
			Address: &order.Address{
				StreetAddress: o.Address.StreetAddress,
				City:          o.Address.City,
				State:         o.Address.State,
				Country:       o.Address.Country,
				ZipCode:       o.Address.ZipCode,
			},
			OrderItems: items,
		})
	}

	resp.Orders = orders
	return resp, nil
}
