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

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{
		RequestContext: RequestContext,
		Context:        Context,
	}
}

func (h *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	resp = &cart.EmptyCartResp{
		Base: &common.BaseResp{},
	}

	_, err = rpc.CartClient.EmptyCart(h.Context, &cartrpc.EmptyCartReq{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		klog.Errorf("cart rpc EmptyCart error: %v", err)
		resp.Base.StatusMessage = "Failed to empty cart"
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		return resp, errors.Wrap(err, "cart EmptyCart error")
	}

	resp.Base.StatusMessage = "Success"
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	return resp, nil
}
