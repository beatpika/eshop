package service

import (
	"context"
	"log"

	"github.com/beatpika/eshop/app/order/biz/dal/mysql"
	"github.com/beatpika/eshop/app/order/biz/model"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// 1. 参数校验
	if len(req.OrderItems) == 0 {
		err = kerrors.NewBizStatusError(400, "购物车为空")
		return
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error { // 开启事务,
		orderId, _ := uuid.NewUUID()
		// 创建订单
		o := &model.Order{
			OrderId: orderId.String(),
			UserId:  req.UserId,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		// 提前检查地址，不为空才赋值
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAdress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}
		// 创建订单
		if err := tx.Create(o).Error; err != nil {
			log.Printf("Create order failed: %v", err)
			return err
		}
		// 创建订单商品
		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				ProductId:    v.Item.ProductId,
				OrderIdRefer: orderId.String(),
				Quantity:     uint32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}
		// 写入订单商品子表
		if err := tx.Create(items).Error; err != nil {
			return err
		}
		// 发送resp
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}
