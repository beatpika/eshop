package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductStatusService {
	return &UpdateProductStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductStatusService) Run(req *product.UpdateProductStatusReq) (resp *product.UpdateProductStatusResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
