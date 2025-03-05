package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListCategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoriesService(Context context.Context, RequestContext *app.RequestContext) *ListCategoriesService {
	return &ListCategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoriesService) Run(req *frontendProduct.ListCategoriesReq) (resp *frontendProduct.ListCategoriesResp, err error) {
	resp = new(frontendProduct.ListCategoriesResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.ListCategoriesReq{
		ParentId: req.ParentId,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.ListCategories(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"

	// 转换分类列表
	if len(rpcResp.Categories) > 0 {
		resp.Categories = make([]*frontendProduct.CategoryInfo, 0, len(rpcResp.Categories))
		for _, c := range rpcResp.Categories {
			resp.Categories = append(resp.Categories, &frontendProduct.CategoryInfo{
				Id:        c.Id,
				Name:      c.Name,
				ParentId:  c.ParentId,
				Level:     c.Level,
				SortOrder: c.SortOrder,
			})
		}
	}

	return resp, nil
}
