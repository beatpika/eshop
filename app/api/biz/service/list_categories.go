package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListCategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoriesService(Context context.Context, RequestContext *app.RequestContext) *ListCategoriesService {
	return &ListCategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
