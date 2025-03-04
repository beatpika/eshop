package main

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/beatpika/eshop/app/product/biz/service"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// ListProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// UpdateProductStatus implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProductStatus(ctx context.Context, req *product.UpdateProductStatusReq) (resp *product.UpdateProductStatusResp, err error) {
	resp, err = service.NewUpdateProductStatusService(ctx).Run(req)

	return resp, err
}

// CreateSKU implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateSKU(ctx context.Context, req *product.CreateSKUReq) (resp *product.CreateSKUResp, err error) {
	resp, err = service.NewCreateSKUService(ctx).Run(req)

	return resp, err
}

// UpdateSKU implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateSKU(ctx context.Context, req *product.UpdateSKUReq) (resp *product.UpdateSKUResp, err error) {
	resp, err = service.NewUpdateSKUService(ctx).Run(req)

	return resp, err
}

// DeleteSKU implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteSKU(ctx context.Context, req *product.DeleteSKUReq) (resp *product.DeleteSKUResp, err error) {
	resp, err = service.NewDeleteSKUService(ctx).Run(req)

	return resp, err
}

// UpdateSKUStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateSKUStock(ctx context.Context, req *product.UpdateSKUStockReq) (resp *product.UpdateSKUStockResp, err error) {
	resp, err = service.NewUpdateSKUStockService(ctx).Run(req)

	return resp, err
}

// CreateCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	resp, err = service.NewCreateCategoryService(ctx).Run(req)

	return resp, err
}

// UpdateCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	resp, err = service.NewUpdateCategoryService(ctx).Run(req)

	return resp, err
}

// DeleteCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	resp, err = service.NewDeleteCategoryService(ctx).Run(req)

	return resp, err
}

// ListCategories implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListCategories(ctx context.Context, req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	resp, err = service.NewListCategoriesService(ctx).Run(req)

	return resp, err
}

// GetCategoryTree implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetCategoryTree(ctx context.Context, req *product.GetCategoryTreeReq) (resp *product.GetCategoryTreeResp, err error) {
	resp, err = service.NewGetCategoryTreeService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// GetProductsByCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductsByCategory(ctx context.Context, req *product.GetProductsByCategoryReq) (resp *product.GetProductsByCategoryResp, err error) {
	resp, err = service.NewGetProductsByCategoryService(ctx).Run(req)

	return resp, err
}
