package service

import (
	"context"

	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductRequest) (resp *product.CreateProductResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
