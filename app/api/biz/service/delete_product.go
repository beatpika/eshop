package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *frontendProduct.DeleteProductReq) (resp *frontendProduct.DeleteProductResp, err error) {
	resp = new(frontendProduct.DeleteProductResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.DeleteProductReq{
		Id: req.ProductId,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.DeleteProduct(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	if rpcResp.Success {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
		resp.Base.StatusMessage = "success"
	} else {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = "failed to delete product"
	}

	return resp, nil
}
