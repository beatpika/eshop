package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"

	cart "github.com/beatpika/eshop/app/api/hertz_gen/basic/cart"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	cartrpc "github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
	productrpc "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{
		RequestContext: RequestContext,
		Context:        Context,
	}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp = &cart.GetCartResp{
		Base: &common.BaseResp{},
		Cart: &cart.CartInfo{},
	}

	// Get cart items from cart service
	rpcResp, err := rpc.CartClient.GetCart(h.Context, &cartrpc.GetCartReq{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		klog.Errorf("cart rpc GetCart error: %v", err)
		resp.Base.StatusMessage = "Failed to get cart"
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		return resp, errors.Wrap(err, "cart GetCart error")
	}

	// Convert cart info and enrich with product details
	cartInfo := rpcResp.GetCart()
	if cartInfo == nil {
		// Return empty cart with success status
		resp.Base.StatusMessage = "Success"
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
		resp.Cart.UserId = req.UserId
		return resp, nil
	}

	resp.Cart.UserId = req.UserId

	// Get product details for each item
	var total int64
	for _, item := range cartInfo.Items {
		rpcProductResp, err := rpc.ProductClient.GetProduct(h.Context, &productrpc.GetProductReq{
			Id: uint32(item.ProductId),
		})
		if err != nil {
			klog.Errorf("product rpc GetProduct error: %v", err)
			resp.Base.StatusMessage = "Failed to get product details"
			resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
			return resp, errors.Wrap(err, "product GetProduct error")
		}

		productInfo := rpcProductResp.GetProduct()
		if productInfo == nil {
			klog.Warnf("product not found for id: %d", item.ProductId)
			resp.Base.StatusMessage = "Some products not found"
			resp.Base.StatusCode = int32(common.StatusCode_STATUS_NOT_FOUND)
			return resp, errors.Errorf("product not found for id: %d", item.ProductId)
		}

		// Convert price from float to int64 (cents)
		priceInCents := int64(productInfo.Price * 100)
		subtotal := int64(item.Quantity) * priceInCents

		cartItem := &cart.CartItem{
			ProductId:    int64(item.ProductId),
			ProductName:  productInfo.Name,
			ProductImage: productInfo.Picture,
			Price:        priceInCents,
			Quantity:     item.Quantity,
			Subtotal:     subtotal,
		}
		resp.Cart.Items = append(resp.Cart.Items, cartItem)
		total += subtotal
	}

	resp.Cart.Total = total

	resp.Base.StatusMessage = "Success"
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	return resp, nil
}
