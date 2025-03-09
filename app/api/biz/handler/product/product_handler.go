package product

import (
	"context"

	"github.com/beatpika/eshop/app/api/biz/service"
	"github.com/beatpika/eshop/app/api/biz/utils"
	product "github.com/beatpika/eshop/app/api/hertz_gen/basic/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateProduct .
// @router /product/manage [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.CreateProductResponse{}
	resp, err = service.NewCreateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProduct .
// @router /product/manage/update/:id [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateProductResponse{}
	resp, err = service.NewUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProduct .
// @router /product/manage/delete/:id [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteProductRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.DeleteProductResponse{}
	resp, err = service.NewDeleteProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProduct .
// @router /product/detail/:id [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.GetProductResponse{}
	resp, err = service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListProducts .
// @router /products [GET]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListProductsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.ListProductsResponse{}
	resp, err = service.NewListProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// SearchProducts .
// @router /products/search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.SearchProductsResponse{}
	resp, err = service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
