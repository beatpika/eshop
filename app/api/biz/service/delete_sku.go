package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteSKUService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteSKUService(Context context.Context, RequestContext *app.RequestContext) *DeleteSKUService {
	return &DeleteSKUService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteSKUService) Run(req *frontendProduct.DeleteSKUReq) (resp *frontendProduct.DeleteSKUResp, err error) {
	resp = new(frontendProduct.DeleteSKUResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.DeleteSKUReq{
		Id: req.SkuId,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.DeleteSKU(h.Context, rpcReq)
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
		resp.Base.StatusMessage = "failed to delete SKU"
	}

	return resp, nil
}
