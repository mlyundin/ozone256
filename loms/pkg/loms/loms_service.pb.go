// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: loms_service.proto

package loms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	OrderStatus_None            OrderStatus = 0
	OrderStatus_New             OrderStatus = 1
	OrderStatus_AwaitingPayment OrderStatus = 2
	OrderStatus_Failed          OrderStatus = 3
	OrderStatus_Payed           OrderStatus = 4
	OrderStatus_Cancelled       OrderStatus = 5
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "None",
		1: "New",
		2: "AwaitingPayment",
		3: "Failed",
		4: "Payed",
		5: "Cancelled",
	}
	OrderStatus_value = map[string]int32{
		"None":            0,
		"New":             1,
		"AwaitingPayment": 2,
		"Failed":          3,
		"Payed":           4,
		"Cancelled":       5,
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
	return file_loms_service_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_loms_service_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{0}
}

type StocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *StocksRequest) Reset() {
	*x = StocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksRequest) ProtoMessage() {}

func (x *StocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksRequest.ProtoReflect.Descriptor instead.
func (*StocksRequest) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{0}
}

func (x *StocksRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type StockItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseID int64  `protobuf:"varint,1,opt,name=warehouseID,proto3" json:"warehouseID,omitempty"`
	Count       uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StockItem) Reset() {
	*x = StockItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockItem) ProtoMessage() {}

func (x *StockItem) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockItem.ProtoReflect.Descriptor instead.
func (*StockItem) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{1}
}

func (x *StockItem) GetWarehouseID() int64 {
	if x != nil {
		return x.WarehouseID
	}
	return 0
}

func (x *StockItem) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*StockItem `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *StocksResponse) Reset() {
	*x = StocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksResponse) ProtoMessage() {}

func (x *StocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksResponse.ProtoReflect.Descriptor instead.
func (*StocksResponse) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{2}
}

func (x *StocksResponse) GetStocks() []*StockItem {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrderResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=route256.loms.OrderStatus" json:"status,omitempty"`
	User   int64       `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*Item     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrderResponse) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_None
}

func (x *ListOrderResponse) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ListOrderResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{8}
}

func (x *CancelOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type OrderPayedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OrderPayedRequest) Reset() {
	*x = OrderPayedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedRequest) ProtoMessage() {}

func (x *OrderPayedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedRequest.ProtoReflect.Descriptor instead.
func (*OrderPayedRequest) Descriptor() ([]byte, []int) {
	return file_loms_service_proto_rawDescGZIP(), []int{9}
}

func (x *OrderPayedRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_loms_service_proto protoreflect.FileDescriptor

var file_loms_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x21, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x73, 0x6b, 0x75, 0x22, 0x43, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36,
	0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x86, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35,
	0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x5b, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x77,
	0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x61, 0x79, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x64, 0x10, 0x05, 0x32, 0x85, 0x03, 0x0a, 0x04, 0x4c, 0x6f, 0x6d, 0x73, 0x12, 0x45,
	0x0a, 0x06, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35,
	0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32,
	0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1d, 0x5a,
	0x1b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x3b, 0x6c, 0x6f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loms_service_proto_rawDescOnce sync.Once
	file_loms_service_proto_rawDescData = file_loms_service_proto_rawDesc
)

func file_loms_service_proto_rawDescGZIP() []byte {
	file_loms_service_proto_rawDescOnce.Do(func() {
		file_loms_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_loms_service_proto_rawDescData)
	})
	return file_loms_service_proto_rawDescData
}

var file_loms_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loms_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_loms_service_proto_goTypes = []interface{}{
	(OrderStatus)(0),            // 0: route256.loms.OrderStatus
	(*StocksRequest)(nil),       // 1: route256.loms.StocksRequest
	(*StockItem)(nil),           // 2: route256.loms.StockItem
	(*StocksResponse)(nil),      // 3: route256.loms.StocksResponse
	(*Item)(nil),                // 4: route256.loms.Item
	(*CreateOrderRequest)(nil),  // 5: route256.loms.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 6: route256.loms.CreateOrderResponse
	(*ListOrderRequest)(nil),    // 7: route256.loms.ListOrderRequest
	(*ListOrderResponse)(nil),   // 8: route256.loms.ListOrderResponse
	(*CancelOrderRequest)(nil),  // 9: route256.loms.CancelOrderRequest
	(*OrderPayedRequest)(nil),   // 10: route256.loms.OrderPayedRequest
	(*emptypb.Empty)(nil),       // 11: google.protobuf.Empty
}
var file_loms_service_proto_depIdxs = []int32{
	2,  // 0: route256.loms.StocksResponse.stocks:type_name -> route256.loms.StockItem
	4,  // 1: route256.loms.CreateOrderRequest.items:type_name -> route256.loms.Item
	0,  // 2: route256.loms.ListOrderResponse.status:type_name -> route256.loms.OrderStatus
	4,  // 3: route256.loms.ListOrderResponse.items:type_name -> route256.loms.Item
	1,  // 4: route256.loms.Loms.Stocks:input_type -> route256.loms.StocksRequest
	5,  // 5: route256.loms.Loms.CreateOrder:input_type -> route256.loms.CreateOrderRequest
	7,  // 6: route256.loms.Loms.ListOrder:input_type -> route256.loms.ListOrderRequest
	9,  // 7: route256.loms.Loms.CancelOrder:input_type -> route256.loms.CancelOrderRequest
	10, // 8: route256.loms.Loms.OrderPayed:input_type -> route256.loms.OrderPayedRequest
	3,  // 9: route256.loms.Loms.Stocks:output_type -> route256.loms.StocksResponse
	6,  // 10: route256.loms.Loms.CreateOrder:output_type -> route256.loms.CreateOrderResponse
	8,  // 11: route256.loms.Loms.ListOrder:output_type -> route256.loms.ListOrderResponse
	11, // 12: route256.loms.Loms.CancelOrder:output_type -> google.protobuf.Empty
	11, // 13: route256.loms.Loms.OrderPayed:output_type -> google.protobuf.Empty
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_loms_service_proto_init() }
func file_loms_service_proto_init() {
	if File_loms_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loms_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksRequest); i {
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
		file_loms_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockItem); i {
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
		file_loms_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksResponse); i {
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
		file_loms_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_loms_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_loms_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_loms_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
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
		file_loms_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResponse); i {
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
		file_loms_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_loms_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedRequest); i {
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
			RawDescriptor: file_loms_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loms_service_proto_goTypes,
		DependencyIndexes: file_loms_service_proto_depIdxs,
		EnumInfos:         file_loms_service_proto_enumTypes,
		MessageInfos:      file_loms_service_proto_msgTypes,
	}.Build()
	File_loms_service_proto = out.File
	file_loms_service_proto_rawDesc = nil
	file_loms_service_proto_goTypes = nil
	file_loms_service_proto_depIdxs = nil
}
