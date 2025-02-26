# 电商系统API设计文档

## 1. 用户服务 (UserService)

### 现有接口
- 用户注册 (Register)
- 用户登录 (Login)

### 建议添加的接口
- 获取用户信息 (GetUserInfo)
- 更新用户信息 (UpdateUserInfo)
- 重置密码 (ResetPassword)
- 更新密码 (UpdatePassword)
- 获取用户地址列表 (ListUserAddresses)
- 添加收货地址 (AddUserAddress)
- 更新收货地址 (UpdateUserAddress)
- 删除收货地址 (DeleteUserAddress)
- 设置默认地址 (SetDefaultAddress)

## 2. 商家服务 (MerchantService)

### 建议接口
- 商家注册 (RegisterMerchant)
- 商家登录 (MerchantLogin)
- 获取商家信息 (GetMerchantInfo)
- 更新商家信息 (UpdateMerchantInfo)
- 更新商家密码 (UpdateMerchantPassword)
- 商家认证申请 (ApplyAuthentication)
- 获取认证状态 (GetAuthenticationStatus)
- 更新商家结算账户 (UpdateSettlementAccount)
- 获取商家统计数据 (GetMerchantStatistics)
- 获取商家订单列表 (ListMerchantOrders)
- 获取商家收入报表 (GetMerchantIncomeReport)
- 获取店铺评分 (GetShopRating)
- 更新店铺信息 (UpdateShopInfo)
- 设置营业状态 (SetOperationStatus)
- 管理店铺分类 (ManageShopCategories)

## 3. 商品服务 (ProductService)

### 建议接口
- 创建商品 (CreateProduct)
- 更新商品 (UpdateProduct)
- 删除商品 (DeleteProduct)
- 获取商品详情 (GetProduct)
- 商品列表查询 (ListProducts)
- 商品分类查询 (ListCategories)
- 商品库存查询 (GetProductStock)
- 批量获取商品信息 (BatchGetProducts)
- 搜索商品 (SearchProducts)
- 获取商品评价 (GetProductReviews)
- 发布商品评价 (CreateProductReview)
- 商品上下架 (UpdateProductStatus)
- 批量更新商品价格 (BatchUpdatePrices)
- 管理商品规格 (ManageProductSpecs)
- 管理商品标签 (ManageProductTags)

## 4. 购物车服务 (CartService)

### 建议接口
- 添加商品到购物车 (AddToCart)
- 从购物车移除商品 (RemoveFromCart)
- 更新购物车商品数量 (UpdateCartItemQuantity)
- 获取购物车列表 (GetCart)
- 清空购物车 (ClearCart)
- 选择/取消选择购物车商品 (SelectCartItems)
- 获取已选商品价格 (GetSelectedItemsPrice)

## 5. 订单服务 (OrderService)

### 建议接口
- 创建订单 (CreateOrder)
- 取消订单 (CancelOrder)
- 获取订单详情 (GetOrder)
- 获取订单列表 (ListOrders)
- 确认收货 (ConfirmReceived)
- 申请退款 (RequestRefund)
- 获取订单支付信息 (GetOrderPayment)
- 获取物流信息 (GetOrderShipment)
- 获取订单统计信息 (GetOrderStatistics)
- 商家接单 (AcceptOrder)
- 商家拒单 (RejectOrder)
- 商家发货 (ShipOrder)
- 处理退款申请 (ProcessRefundRequest)
- 修改订单价格 (UpdateOrderPrice)
- 添加订单备注 (AddOrderNote)

## 6. 支付服务 (PaymentService)

### 建议接口
- 创建支付单 (CreatePayment)
- 支付回调 (PaymentCallback)
- 查询支付状态 (GetPaymentStatus)
- 退款 (Refund)
- 查询退款状态 (GetRefundStatus)
- 获取支付方式列表 (ListPaymentMethods)
- 商家提现 (MerchantWithdraw)
- 获取商家账户余额 (GetMerchantBalance)
- 获取商家交易记录 (ListMerchantTransactions)
- 设置结算账户 (SetSettlementAccount)

## 7. 优惠券服务 (CouponService)

### 建议接口
- 创建优惠券 (CreateCoupon)
- 领取优惠券 (ClaimCoupon)
- 获取用户优惠券列表 (ListUserCoupons)
- 查询优惠券信息 (GetCoupon)
- 验证优惠券有效性 (ValidateCoupon)
- 使用优惠券 (UseCoupon)
- 获取可用优惠券列表 (ListAvailableCoupons)
- 停用优惠券 (DisableCoupon)
- 获取优惠券使用记录 (GetCouponUsageHistory)
- 批量发放优惠券 (BatchIssueCoupons)

## 8. 库存服务 (InventoryService)

### 建议接口
- 更新库存 (UpdateStock)
- 锁定库存 (LockStock)
- 释放库存 (ReleaseStock)
- 查询库存 (GetStock)
- 批量查询库存 (BatchGetStock)
- 设置库存警戒值 (SetStockAlert)
- 获取库存变更记录 (GetStockHistory)
- 库存盘点 (StockTaking)
- 设置自动补货规则 (SetAutoReplenishmentRules)
- 获取库存预警列表 (ListStockAlerts)

## 9. 搜索服务 (SearchService)

### 建议接口
- 搜索商品 (SearchProducts)
- 获取搜索建议 (GetSearchSuggestions)
- 获取热门搜索 (GetHotSearches)
- 记录搜索历史 (RecordSearchHistory)
- 获取用户搜索历史 (GetUserSearchHistory)
- 清除搜索历史 (ClearSearchHistory)
- 搜索店铺 (SearchShops)
- 更新搜索索引 (UpdateSearchIndex)
- 获取相关商品推荐 (GetRelatedProducts)
- 获取类目热门商品 (GetCategoryHotProducts)

## 10. 评价服务 (ReviewService)

### 建议接口
- 创建商品评价 (CreateReview)
- 回复评价 (ReplyReview)
- 获取评价列表 (ListReviews)
- 获取评价详情 (GetReview)
- 点赞/取消点赞评价 (LikeReview)
- 获取待评价订单商品 (ListPendingReviews)
- 商家回复评价 (MerchantReplyReview)
- 举报评价 (ReportReview)
- 获取店铺评价统计 (GetShopReviewStats)
- 隐藏/显示评价 (UpdateReviewVisibility)

## 11. 消息服务 (MessageService)

### 建议接口
- 发送消息 (SendMessage)
- 获取消息列表 (ListMessages)
- 标记消息已读 (MarkMessageRead)
- 删除消息 (DeleteMessage)
- 获取未读消息数 (GetUnreadCount)
- 批量标记已读 (BatchMarkRead)
- 系统通知发送 (SendSystemNotification)
- 商家消息通知 (SendMerchantNotification)
- 订单状态通知 (SendOrderStatusNotification)
- 优惠活动通知 (SendPromotionNotification)

## 数据结构设计建议

所有的请求响应都应该遵循以下原则：

1. 统一的响应格式
```protobuf
message BaseResponse {
    int32 code = 1;        // 状态码
    string message = 2;    // 错误信息
    string request_id = 3; // 请求ID，用于跟踪
}
```

2. 分页请求格式
```protobuf
message PaginationRequest {
    int32 page_size = 1;   // 每页数量
    int32 page_number = 2; // 页码
}
```

3. 分页响应格式
```protobuf
message PaginationResponse {
    int32 total = 1;       // 总数
    int32 page_size = 2;   // 每页数量
    int32 page_number = 3; // 当前页码
}
```

## 安全性考虑

1. 所有接口都应该进行适当的权限验证
   - 用户权限验证
   - 商家权限验证
   - 管理员权限验证
2. 敏感数据传输应使用HTTPS
3. 实现速率限制以防止滥用
4. 实现数据验证以确保输入安全
5. 使用JWT或类似机制进行身份验证
6. 实现审计日志记录关键操作
7. 商家操作需要额外的安全验证

## 性能优化建议

1. 使用缓存减少数据库负载
2. 实现批量接口减少网络请求
3. 使用异步处理进行耗时操作
4. 合理设计索引提高查询效率
5. 使用消息队列处理异步任务
6. 实现服务降级和熔断机制
7. 针对商家端的高并发请求进行优化
8. 实现数据预热机制

## 商家特定的设计考虑

1. 多店铺管理支持
2. 商家权限分级管理
3. 店铺运营数据分析
4. 商品批量操作支持
5. 订单处理流程优化
6. 结算周期和规则管理
7. 商家端专用的缓存策略
8. 店铺装修和展示管理