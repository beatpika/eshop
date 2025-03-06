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
		resp.Base.StatusMessage = "Success"
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
		return resp, nil
	}

	resp.Cart.UserId = req.UserId

	// Get product details for each item
	var total int64
	for _, item := range cartInfo.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &productrpc.GetProductReq{
			Id: int64(item.ProductId),
		})
		if err != nil {
			klog.Errorf("product rpc GetProduct error: %v", err)
			continue
		}

		product := productResp.GetProduct()
		if product == nil {
			continue
		}

		subtotal := int64(item.Quantity) * product.Price
		var image string
		if len(product.Images) > 0 {
			image = product.Images[0]
		}

		cartItem := &cart.CartItem{
			ProductId:    int64(item.ProductId),
			ProductName:  product.Name,
			ProductImage: image,
			Price:        product.Price,
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
