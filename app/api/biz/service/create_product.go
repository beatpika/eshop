package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/biz/utils"
	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	rpc_product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
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
	resp = new(product.CreateProductResponse)
	resp.Base = new(common.BaseResp)

	// 构造RPC请求
	rpcReq := &rpc_product.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.CreateProduct(h.Context, rpcReq)
	if err != nil {
		resp.Base = utils.BuildBaseResp(err)
		return resp, nil
	}

	// 转换响应
	if rpcResp.Product != nil {
		resp.Product = &product.ProductInfo{
			Id:          rpcResp.Product.Id,
			Name:        rpcResp.Product.Name,
			Description: rpcResp.Product.Description,
			Picture:     rpcResp.Product.Picture,
			Price:       rpcResp.Product.Price,
			Categories:  rpcResp.Product.Categories,
		}
	}

	resp.Base = utils.BuildBaseResp(nil)
	return resp, nil
}
