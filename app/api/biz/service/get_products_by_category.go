package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductsByCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductsByCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetProductsByCategoryService {
	return &GetProductsByCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductsByCategoryService) Run(req *product.GetProductsByCategoryReq) (resp *product.GetProductsByCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
