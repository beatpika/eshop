package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"

	cart "github.com/beatpika/eshop/app/api/hertz_gen/basic/cart"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	cartrpc "github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{
		RequestContext: RequestContext,
		Context:        Context,
	}
}

func (h *AddItemService) Run(req *cart.AddCartItemReq) (resp *cart.AddCartItemResp, err error) {
	resp = &cart.AddCartItemResp{
		Base: &common.BaseResp{},
	}

	// Convert types and call cart RPC
	rpcReq := &cartrpc.AddItemReq{
		UserId: uint32(req.UserId),
		Item: &cartrpc.CartItem{
			ProductId: uint32(req.ProductId),
			Quantity:  req.Quantity,
		},
	}

	_, err = rpc.CartClient.AddItem(h.Context, rpcReq)
	if err != nil {
		klog.Errorf("cart rpc AddItem error: %v", err)
		resp.Base.StatusMessage = "Failed to add item to cart"
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		return resp, errors.Wrap(err, "cart AddItem error")
	}

	resp.Base.StatusMessage = "Success"
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	return resp, nil
}
