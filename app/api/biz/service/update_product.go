package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *frontendProduct.UpdateProductReq) (resp *frontendProduct.UpdateProductResp, err error) {
	resp = new(frontendProduct.UpdateProductResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.UpdateProductReq{
		Id:          req.ProductId,
		Name:        req.Name,
		Description: req.Description,
		CategoryId:  req.CategoryId,
		Images:      req.Images,
		Price:       req.Price,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.UpdateProduct(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"

	if rpcResp.Product != nil {
		resp.Product = &frontendProduct.ProductInfo{
			Id:          rpcResp.Product.Id,
			Name:        rpcResp.Product.Name,
			Description: rpcResp.Product.Description,
			CategoryId:  rpcResp.Product.CategoryId,
			Images:      rpcResp.Product.Images,
			Price:       rpcResp.Product.Price,
			Status:      rpcResp.Product.Status,
			CreatedAt:   rpcResp.Product.CreatedAt,
			UpdatedAt:   rpcResp.Product.UpdatedAt,
		}

		// 转换SKU列表
		if len(rpcResp.Product.Skus) > 0 {
			resp.Product.Skus = make([]*frontendProduct.SKUInfo, 0, len(rpcResp.Product.Skus))
			for _, sku := range rpcResp.Product.Skus {
				resp.Product.Skus = append(resp.Product.Skus, &frontendProduct.SKUInfo{
					Id:        sku.Id,
					ProductId: sku.ProductId,
					Specs:     sku.Specs,
					Price:     sku.Price,
					Stock:     sku.Stock,
					Code:      sku.Code,
				})
			}
		}
	}

	return resp, nil
}
