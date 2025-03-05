package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductStatusService {
	return &UpdateProductStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductStatusService) Run(req *frontendProduct.UpdateProductStatusReq) (resp *frontendProduct.UpdateProductStatusResp, err error) {
	resp = new(frontendProduct.UpdateProductStatusResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.UpdateProductStatusReq{
		Id:     req.ProductId,
		Status: req.Status,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.UpdateProductStatus(h.Context, rpcReq)
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
		resp.Base.StatusMessage = "failed to update product status"
	}

	return resp, nil
}
