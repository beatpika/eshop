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

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.GetProductRequest) (resp *product.GetProductResponse, err error) {
	resp = new(product.GetProductResponse)
	resp.Base = new(common.BaseResp)

	// 构造RPC请求
	rpcReq := &rpc_product.GetProductReq{
		Id: req.Id,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.GetProduct(h.Context, rpcReq)
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
