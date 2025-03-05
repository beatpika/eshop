package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateSKUService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateSKUService(Context context.Context, RequestContext *app.RequestContext) *CreateSKUService {
	return &CreateSKUService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateSKUService) Run(req *product.CreateSKUReq) (resp *product.CreateSKUResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
