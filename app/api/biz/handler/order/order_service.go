package order

import (
	"context"

	"github.com/beatpika/eshop/app/api/biz/service"
	"github.com/beatpika/eshop/app/api/biz/utils"
	order "github.com/beatpika/eshop/app/api/hertz_gen/basic/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListOrders .
// @router /order [POST]
func ListOrders(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrdersReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.ListOrdersResp{}
	resp, err = service.NewListOrdersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
