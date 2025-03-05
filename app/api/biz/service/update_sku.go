package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateSKUService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateSKUService(Context context.Context, RequestContext *app.RequestContext) *UpdateSKUService {
	return &UpdateSKUService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateSKUService) Run(req *frontendProduct.UpdateSKUReq) (resp *frontendProduct.UpdateSKUResp, err error) {
	resp = new(frontendProduct.UpdateSKUResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.UpdateSKUReq{
		Id:    req.SkuId,
		Specs: req.Specs,
		Price: req.Price,
		Code:  req.Code,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.UpdateSKU(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"

	if rpcResp.Sku != nil {
		resp.Sku = &frontendProduct.SKUInfo{
			Id:        rpcResp.Sku.Id,
			ProductId: rpcResp.Sku.ProductId,
			Specs:     rpcResp.Sku.Specs,
			Price:     rpcResp.Sku.Price,
			Stock:     rpcResp.Sku.Stock,
			Code:      rpcResp.Sku.Code,
		}
	}

	return resp, nil
}
