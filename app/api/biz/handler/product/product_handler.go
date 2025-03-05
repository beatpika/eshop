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
// @router /product [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.CreateProductResp{}
	resp, err = service.NewCreateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProduct .
// @router /product/:product_id [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateProductResp{}
	resp, err = service.NewUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProduct .
// @router /product/:product_id [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.GetProductResp{}
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
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.ListProductsResp{}
	resp, err = service.NewListProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProduct .
// @router /product/:product_id [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.DeleteProductResp{}
	resp, err = service.NewDeleteProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProductStatus .
// @router /product/:product_id/status [PUT]
func UpdateProductStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductStatusReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateProductStatusResp{}
	resp, err = service.NewUpdateProductStatusService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateSKU .
// @router /product/:product_id/sku [POST]
func CreateSKU(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateSKUReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.CreateSKUResp{}
	resp, err = service.NewCreateSKUService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateSKU .
// @router /sku/:sku_id [PUT]
func UpdateSKU(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateSKUReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateSKUResp{}
	resp, err = service.NewUpdateSKUService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteSKU .
// @router /sku/:sku_id [DELETE]
func DeleteSKU(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteSKUReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.DeleteSKUResp{}
	resp, err = service.NewDeleteSKUService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateSKUStock .
// @router /sku/:sku_id/stock [PUT]
func UpdateSKUStock(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateSKUStockReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateSKUStockResp{}
	resp, err = service.NewUpdateSKUStockService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateCategory .
// @router /category [POST]
func CreateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.CreateCategoryResp{}
	resp, err = service.NewCreateCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateCategory .
// @router /category/:category_id [PUT]
func UpdateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.UpdateCategoryResp{}
	resp, err = service.NewUpdateCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteCategory .
// @router /category/:category_id [DELETE]
func DeleteCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.DeleteCategoryResp{}
	resp, err = service.NewDeleteCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListCategories .
// @router /categories [GET]
func ListCategories(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListCategoriesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.ListCategoriesResp{}
	resp, err = service.NewListCategoriesService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCategoryTree .
// @router /categories/tree [GET]
func GetCategoryTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetCategoryTreeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.GetCategoryTreeResp{}
	resp, err = service.NewGetCategoryTreeService(ctx, c).Run(&req)
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
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.SearchProductsResp{}
	resp, err = service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProductsByCategory .
// @router /categories/:category_id/products [GET]
func GetProductsByCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductsByCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.GetProductsByCategoryResp{}
	resp, err = service.NewGetProductsByCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
