package product

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (resp *product.CreateProductResp, err error) {
	resp, err = defaultClient.CreateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.UpdateProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProductStatus(ctx context.Context, req *product.UpdateProductStatusReq, callOptions ...callopt.Option) (resp *product.UpdateProductStatusResp, err error) {
	resp, err = defaultClient.UpdateProductStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProductStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateSKU(ctx context.Context, req *product.CreateSKUReq, callOptions ...callopt.Option) (resp *product.CreateSKUResp, err error) {
	resp, err = defaultClient.CreateSKU(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateSKU call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateSKU(ctx context.Context, req *product.UpdateSKUReq, callOptions ...callopt.Option) (resp *product.UpdateSKUResp, err error) {
	resp, err = defaultClient.UpdateSKU(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateSKU call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteSKU(ctx context.Context, req *product.DeleteSKUReq, callOptions ...callopt.Option) (resp *product.DeleteSKUResp, err error) {
	resp, err = defaultClient.DeleteSKU(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteSKU call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateSKUStock(ctx context.Context, req *product.UpdateSKUStockReq, callOptions ...callopt.Option) (resp *product.UpdateSKUStockResp, err error) {
	resp, err = defaultClient.UpdateSKUStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateSKUStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateCategory(ctx context.Context, req *product.CreateCategoryReq, callOptions ...callopt.Option) (resp *product.CreateCategoryResp, err error) {
	resp, err = defaultClient.CreateCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq, callOptions ...callopt.Option) (resp *product.UpdateCategoryResp, err error) {
	resp, err = defaultClient.UpdateCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq, callOptions ...callopt.Option) (resp *product.DeleteCategoryResp, err error) {
	resp, err = defaultClient.DeleteCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListCategories(ctx context.Context, req *product.ListCategoriesReq, callOptions ...callopt.Option) (resp *product.ListCategoriesResp, err error) {
	resp, err = defaultClient.ListCategories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListCategories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategoryTree(ctx context.Context, req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (resp *product.GetCategoryTreeResp, err error) {
	resp, err = defaultClient.GetCategoryTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductsByCategory(ctx context.Context, req *product.GetProductsByCategoryReq, callOptions ...callopt.Option) (resp *product.GetProductsByCategoryResp, err error) {
	resp, err = defaultClient.GetProductsByCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductsByCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
