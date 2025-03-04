package service

import (
	"context"

	order "github.com/beatpika/eshop/app/api/hertz_gen/basic/order"
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
	return
}
