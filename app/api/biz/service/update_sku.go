package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateSKUService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateSKUService(Context context.Context, RequestContext *app.RequestContext) *UpdateSKUService {
	return &UpdateSKUService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateSKUService) Run(req *product.UpdateSKUReq) (resp *product.UpdateSKUResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
