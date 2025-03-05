package service

import (
	"context"

	frontendProduct "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCategoryService(Context context.Context, RequestContext *app.RequestContext) *CreateCategoryService {
	return &CreateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCategoryService) Run(req *frontendProduct.CreateCategoryReq) (resp *frontendProduct.CreateCategoryResp, err error) {
	resp = new(frontendProduct.CreateCategoryResp)
	resp.Base = new(common.BaseResp)

	// 构建RPC请求
	rpcReq := &product.CreateCategoryReq{
		Name:      req.Name,
		ParentId:  req.ParentId,
		Level:     req.Level,
		SortOrder: req.SortOrder,
	}

	// 调用RPC服务
	rpcResp, err := rpc.ProductClient.CreateCategory(h.Context, rpcReq)
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, nil
	}

	// 转换RPC响应为HTTP响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "success"

	if rpcResp.Category != nil {
		resp.Category = &frontendProduct.CategoryInfo{
			Id:        rpcResp.Category.Id,
			Name:      rpcResp.Category.Name,
			ParentId:  rpcResp.Category.ParentId,
			Level:     rpcResp.Category.Level,
			SortOrder: rpcResp.Category.SortOrder,
		}
	}

	return resp, nil
}
