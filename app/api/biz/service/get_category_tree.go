package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryTreeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryTreeService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryTreeService {
	return &GetCategoryTreeService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryTreeService) Run(req *product.GetCategoryTreeReq) (resp *product.GetCategoryTreeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
