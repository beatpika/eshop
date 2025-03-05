package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateSKUStockService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateSKUStockService(Context context.Context, RequestContext *app.RequestContext) *UpdateSKUStockService {
	return &UpdateSKUStockService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateSKUStockService) Run(req *frontendProduct.UpdateSKUStockReq) (resp *frontendProduct.UpdateSKUStockResp, err error) {
	resp = new(frontendProduct.UpdateSKUStockResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.UpdateSKUStockReq{
		Id:    req.SkuId,
		Stock: req.Stock,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.UpdateSKUStock(h.Context, rpcReq)
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
