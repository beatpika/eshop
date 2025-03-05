package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateSKUStockService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateSKUStockService(Context context.Context, RequestContext *app.RequestContext) *UpdateSKUStockService {
	return &UpdateSKUStockService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateSKUStockService) Run(req *product.UpdateSKUStockReq) (resp *product.UpdateSKUStockResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
