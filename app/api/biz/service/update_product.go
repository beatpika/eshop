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

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductRequest) (resp *product.UpdateProductResponse, err error) {
	resp = new(product.UpdateProductResponse)
	resp.Base = new(common.BaseResp)

	// 构造RPC请求
	rpcReq := &rpc_product.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.UpdateProduct(h.Context, rpcReq)
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
