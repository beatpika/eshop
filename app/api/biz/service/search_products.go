package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	productRpc "github.com/beatpika/eshop/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *frontendProduct.SearchProductsReq) (resp *frontendProduct.SearchProductsResp, err error) {
	resp = new(frontendProduct.SearchProductsResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.SearchProductsReq{
		Keyword:    req.Keyword,
		CategoryId: req.CategoryId,
		PageSize:   req.PageSize,
		Page:       req.Page,
		SortBy:     req.SortBy,
	}

	// 调用RPC服务
	rpcResp, err := productRpc.DefaultClient().SearchProducts(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"
	resp.Total = rpcResp.Total

	// 转换商品列表
	if len(rpcResp.Products) > 0 {
		resp.Products = make([]*frontendProduct.ProductInfo, 0, len(rpcResp.Products))
		for _, p := range rpcResp.Products {
			product := &frontendProduct.ProductInfo{
				Id:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				CategoryId:  p.CategoryId,
				Images:      p.Images,
				Price:       p.Price,
				Status:      p.Status,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
			}

			// 转换SKU列表
			if len(p.Skus) > 0 {
				product.Skus = make([]*frontendProduct.SKUInfo, 0, len(p.Skus))
				for _, sku := range p.Skus {
					product.Skus = append(product.Skus, &frontendProduct.SKUInfo{
						Id:        sku.Id,
						ProductId: sku.ProductId,
						Specs:     sku.Specs,
						Price:     sku.Price,
						Stock:     sku.Stock,
						Code:      sku.Code,
					})
				}
			}

			resp.Products = append(resp.Products, product)
		}
	}

	return resp, nil
}
