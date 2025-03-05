package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCategoryService(Context context.Context, RequestContext *app.RequestContext) *DeleteCategoryService {
	return &DeleteCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
