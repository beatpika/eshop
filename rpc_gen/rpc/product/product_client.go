package product

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productservice.Client
	Service() string
	CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	UpdateProductStatus(ctx context.Context, Req *product.UpdateProductStatusReq, callOptions ...callopt.Option) (r *product.UpdateProductStatusResp, err error)
	CreateSKU(ctx context.Context, Req *product.CreateSKUReq, callOptions ...callopt.Option) (r *product.CreateSKUResp, err error)
	UpdateSKU(ctx context.Context, Req *product.UpdateSKUReq, callOptions ...callopt.Option) (r *product.UpdateSKUResp, err error)
	DeleteSKU(ctx context.Context, Req *product.DeleteSKUReq, callOptions ...callopt.Option) (r *product.DeleteSKUResp, err error)
	UpdateSKUStock(ctx context.Context, Req *product.UpdateSKUStockReq, callOptions ...callopt.Option) (r *product.UpdateSKUStockResp, err error)
	CreateCategory(ctx context.Context, Req *product.CreateCategoryReq, callOptions ...callopt.Option) (r *product.CreateCategoryResp, err error)
	UpdateCategory(ctx context.Context, Req *product.UpdateCategoryReq, callOptions ...callopt.Option) (r *product.UpdateCategoryResp, err error)
	DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *product.DeleteCategoryResp, err error)
	ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.ListCategoriesResp, err error)
	GetCategoryTree(ctx context.Context, Req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (r *product.GetCategoryTreeResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	GetProductsByCategory(ctx context.Context, Req *product.GetProductsByCategoryReq, callOptions ...callopt.Option) (r *product.GetProductsByCategoryResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error) {
	return c.kitexClient.CreateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProductStatus(ctx context.Context, Req *product.UpdateProductStatusReq, callOptions ...callopt.Option) (r *product.UpdateProductStatusResp, err error) {
	return c.kitexClient.UpdateProductStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateSKU(ctx context.Context, Req *product.CreateSKUReq, callOptions ...callopt.Option) (r *product.CreateSKUResp, err error) {
	return c.kitexClient.CreateSKU(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateSKU(ctx context.Context, Req *product.UpdateSKUReq, callOptions ...callopt.Option) (r *product.UpdateSKUResp, err error) {
	return c.kitexClient.UpdateSKU(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteSKU(ctx context.Context, Req *product.DeleteSKUReq, callOptions ...callopt.Option) (r *product.DeleteSKUResp, err error) {
	return c.kitexClient.DeleteSKU(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateSKUStock(ctx context.Context, Req *product.UpdateSKUStockReq, callOptions ...callopt.Option) (r *product.UpdateSKUStockResp, err error) {
	return c.kitexClient.UpdateSKUStock(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateCategory(ctx context.Context, Req *product.CreateCategoryReq, callOptions ...callopt.Option) (r *product.CreateCategoryResp, err error) {
	return c.kitexClient.CreateCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateCategory(ctx context.Context, Req *product.UpdateCategoryReq, callOptions ...callopt.Option) (r *product.UpdateCategoryResp, err error) {
	return c.kitexClient.UpdateCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *product.DeleteCategoryResp, err error) {
	return c.kitexClient.DeleteCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.ListCategoriesResp, err error) {
	return c.kitexClient.ListCategories(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCategoryTree(ctx context.Context, Req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (r *product.GetCategoryTreeResp, err error) {
	return c.kitexClient.GetCategoryTree(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProductsByCategory(ctx context.Context, Req *product.GetProductsByCategoryReq, callOptions ...callopt.Option) (r *product.GetProductsByCategoryResp, err error) {
	return c.kitexClient.GetProductsByCategory(ctx, Req, callOptions...)
}
