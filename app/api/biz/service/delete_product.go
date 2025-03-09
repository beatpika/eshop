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

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductRequest) (resp *product.DeleteProductResponse, err error) {
	resp = new(product.DeleteProductResponse)
	resp.Base = new(common.BaseResp)

	// 构造RPC请求
	rpcReq := &rpc_product.DeleteProductReq{
		Id: req.Id,
	}

	// 调用RPC服务
	_, err = rpc.ProductClient.DeleteProduct(h.Context, rpcReq)
	if err != nil {
		resp.Base = utils.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = utils.BuildBaseResp(nil)
	return resp, nil
}
