package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCategoryService(Context context.Context, RequestContext *app.RequestContext) *UpdateCategoryService {
	return &UpdateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCategoryService) Run(req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
