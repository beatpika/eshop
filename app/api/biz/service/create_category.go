package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCategoryService(Context context.Context, RequestContext *app.RequestContext) *CreateCategoryService {
	return &CreateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
