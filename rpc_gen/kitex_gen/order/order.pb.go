// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: order.proto

package order

import (
	context "context"
	cart "github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_PENDING OrderStatus = 0 // 待处理
	OrderStatus_ORDER_STATUS_FINISH  OrderStatus = 1 // 已完成
	OrderStatus_ORDER_STATUS_CANCEL  OrderStatus = 2 // 已取消
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_PENDING",
		1: "ORDER_STATUS_FINISH",
		2: "ORDER_STATUS_CANCEL",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_PENDING": 0,
		"ORDER_STATUS_FINISH":  1,
		"ORDER_STATUS_CANCEL":  2,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency string       `protobuf:"bytes,2,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address      *Address     `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email        string       `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	OrderItems   []*OrderItem `protobuf:"bytes,5,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *PlaceOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *PlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceOrderReq) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *cart.CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cost float32        `protobuf:"fixed32,2,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetItem() *cart.CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *OrderItem) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type OrderResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderResult) Reset() {
	*x = OrderResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResult) ProtoMessage() {}

func (x *OrderResult) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResult.ProtoReflect.Descriptor instead.
func (*OrderResult) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderResult) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type PlaceOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *OrderResult `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *PlaceOrderResp) Reset() {
	*x = PlaceOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResp) ProtoMessage() {}

func (x *PlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResp.ProtoReflect.Descriptor instead.
func (*PlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderResp) GetOrder() *OrderResult {
	if x != nil {
		return x.Order
	}
	return nil
}

type ListOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListOrderReq) Reset() {
	*x = ListOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderReq) ProtoMessage() {}

func (x *ListOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderReq.ProtoReflect.Descriptor instead.
func (*ListOrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderItems   []*OrderItem `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	OrderId      string       `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId       uint32       `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency string       `protobuf:"bytes,4,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address      *Address     `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email        string       `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt    int32        `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status       OrderStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=order.OrderStatus" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Order) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_PENDING
}

type ListOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListOrderResp) Reset() {
	*x = ListOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResp) ProtoMessage() {}

func (x *ListOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResp.ProtoReflect.Descriptor instead.
func (*ListOrderResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrderResp) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type UpdateOrderStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId string      `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  OrderStatus `protobuf:"varint,3,opt,name=status,proto3,enum=order.OrderStatus" json:"status,omitempty"`
}

func (x *UpdateOrderStatusReq) Reset() {
	*x = UpdateOrderStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusReq) ProtoMessage() {}

func (x *UpdateOrderStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderStatusReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOrderStatusReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderStatusReq) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_PENDING
}

type UpdateOrderStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderStatusResp) Reset() {
	*x = UpdateOrderStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResp) ProtoMessage() {}

func (x *UpdateOrderStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResp.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateOrderStatusResp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpdateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Address *Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email   string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateOrderReq) Reset() {
	*x = UpdateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderReq) ProtoMessage() {}

func (x *UpdateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderReq.ProtoReflect.Descriptor instead.
func (*UpdateOrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UpdateOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderResp) Reset() {
	*x = UpdateOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResp) ProtoMessage() {}

func (x *UpdateOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResp.ProtoReflect.Descriptor instead.
func (*UpdateOrderResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateOrderResp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9e, 0x02, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x76, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x35, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2a, 0x59, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x02, 0x32,
	0x97, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61, 0x74, 0x70, 0x69, 0x6b, 0x61,
	0x2f, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_order_proto_goTypes = []interface{}{
	(OrderStatus)(0),              // 0: order.OrderStatus
	(*Address)(nil),               // 1: order.Address
	(*PlaceOrderReq)(nil),         // 2: order.PlaceOrderReq
	(*OrderItem)(nil),             // 3: order.OrderItem
	(*OrderResult)(nil),           // 4: order.OrderResult
	(*PlaceOrderResp)(nil),        // 5: order.PlaceOrderResp
	(*ListOrderReq)(nil),          // 6: order.ListOrderReq
	(*Order)(nil),                 // 7: order.Order
	(*ListOrderResp)(nil),         // 8: order.ListOrderResp
	(*UpdateOrderStatusReq)(nil),  // 9: order.UpdateOrderStatusReq
	(*UpdateOrderStatusResp)(nil), // 10: order.UpdateOrderStatusResp
	(*UpdateOrderReq)(nil),        // 11: order.UpdateOrderReq
	(*UpdateOrderResp)(nil),       // 12: order.UpdateOrderResp
	(*cart.CartItem)(nil),         // 13: cart.CartItem
}
var file_order_proto_depIdxs = []int32{
	1,  // 0: order.PlaceOrderReq.address:type_name -> order.Address
	3,  // 1: order.PlaceOrderReq.order_items:type_name -> order.OrderItem
	13, // 2: order.OrderItem.item:type_name -> cart.CartItem
	4,  // 3: order.PlaceOrderResp.order:type_name -> order.OrderResult
	3,  // 4: order.Order.order_items:type_name -> order.OrderItem
	1,  // 5: order.Order.address:type_name -> order.Address
	0,  // 6: order.Order.status:type_name -> order.OrderStatus
	7,  // 7: order.ListOrderResp.orders:type_name -> order.Order
	0,  // 8: order.UpdateOrderStatusReq.status:type_name -> order.OrderStatus
	7,  // 9: order.UpdateOrderStatusResp.order:type_name -> order.Order
	1,  // 10: order.UpdateOrderReq.address:type_name -> order.Address
	7,  // 11: order.UpdateOrderResp.order:type_name -> order.Order
	2,  // 12: order.OrderService.PlaceOrder:input_type -> order.PlaceOrderReq
	6,  // 13: order.OrderService.ListOrder:input_type -> order.ListOrderReq
	11, // 14: order.OrderService.UpdateOrder:input_type -> order.UpdateOrderReq
	9,  // 15: order.OrderService.UpdateOrderStatus:input_type -> order.UpdateOrderStatusReq
	5,  // 16: order.OrderService.PlaceOrder:output_type -> order.PlaceOrderResp
	8,  // 17: order.OrderService.ListOrder:output_type -> order.ListOrderResp
	12, // 18: order.OrderService.UpdateOrder:output_type -> order.UpdateOrderResp
	10, // 19: order.OrderService.UpdateOrderStatus:output_type -> order.UpdateOrderStatusResp
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderStatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderStatusResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		EnumInfos:         file_order_proto_enumTypes,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type OrderService interface {
	PlaceOrder(ctx context.Context, req *PlaceOrderReq) (res *PlaceOrderResp, err error)
	ListOrder(ctx context.Context, req *ListOrderReq) (res *ListOrderResp, err error)
	UpdateOrder(ctx context.Context, req *UpdateOrderReq) (res *UpdateOrderResp, err error)
	UpdateOrderStatus(ctx context.Context, req *UpdateOrderStatusReq) (res *UpdateOrderStatusResp, err error)
}
