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

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsRequest) (resp *product.SearchProductsResponse, err error) {
	resp = new(product.SearchProductsResponse)
	resp.Base = new(common.BaseResp)

	// 构造RPC请求
	rpcReq := &rpc_product.SearchProductsReq{
		Keywords: req.Keywords,
		Page:     int32(req.Page),
		PageSize: int64(req.PageSize),
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.SearchProducts(h.Context, rpcReq)
	if err != nil {
		resp.Base = utils.BuildBaseResp(err)
		return resp, nil
	}

	// 转换响应
	resp.Products = make([]*product.ProductInfo, 0, len(rpcResp.Results))
	for _, p := range rpcResp.Results {
		resp.Products = append(resp.Products, &product.ProductInfo{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  p.Categories,
		})
	}
	resp.Total = rpcResp.Total

	resp.Base = utils.BuildBaseResp(nil)
	return resp, nil
}
