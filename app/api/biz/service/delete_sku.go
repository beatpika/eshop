package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteSKUService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteSKUService(Context context.Context, RequestContext *app.RequestContext) *DeleteSKUService {
	return &DeleteSKUService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteSKUService) Run(req *product.DeleteSKUReq) (resp *product.DeleteSKUResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
