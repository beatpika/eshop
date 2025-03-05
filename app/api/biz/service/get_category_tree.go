package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryTreeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryTreeService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryTreeService {
	return &GetCategoryTreeService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryTreeService) Run(req *frontendProduct.GetCategoryTreeReq) (resp *frontendProduct.GetCategoryTreeResp, err error) {
	resp = new(frontendProduct.GetCategoryTreeResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.GetCategoryTreeReq{}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.GetCategoryTree(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"

	// 转换分类树
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
