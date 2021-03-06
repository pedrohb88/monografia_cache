// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: transport/proto/server.proto

package proto

import (
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{1}
}

type ByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIDRequest) Reset() {
	*x = ByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIDRequest) ProtoMessage() {}

func (x *ByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIDRequest.ProtoReflect.Descriptor instead.
func (*ByIDRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *ByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *Products) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{5}
}

func (x *Orders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemsQuantity int64    `protobuf:"varint,3,opt,name=items_quantity,json=itemsQuantity,proto3" json:"items_quantity,omitempty"`
	Price         float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	PaymentId     int64    `protobuf:"varint,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Items         []*Item  `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	Payment       *Payment `protobuf:"bytes,7,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[6]
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
	return file_transport_proto_server_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetItemsQuantity() int64 {
	if x != nil {
		return x.ItemsQuantity
	}
	return 0
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId   int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId int64    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Product   *Product `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[7]
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
	return file_transport_proto_server_proto_rawDescGZIP(), []int{7}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Item) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount    float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	InvoiceId int64    `protobuf:"varint,3,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Invoice   *Invoice `protobuf:"bytes,4,opt,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{8}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetInvoiceId() int64 {
	if x != nil {
		return x.InvoiceId
	}
	return 0
}

func (x *Payment) GetInvoice() *Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Link string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{9}
}

func (x *Invoice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Invoice) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Invoice) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_transport_proto_server_proto protoreflect.FileDescriptor

var file_transport_proto_server_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25,
	0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x22, 0x7b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x22, 0x41, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x32, 0x92, 0x05, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x34,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x13,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0d,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0c, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x08, 0x50, 0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x14, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x13, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x3e, 0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x61, 0x42, 0x0f, 0x4d, 0x6f, 0x6e, 0x6f, 0x67,
	0x72, 0x61, 0x66, 0x69, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x6d, 0x6f,
	0x6e, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_server_proto_rawDescOnce sync.Once
	file_transport_proto_server_proto_rawDescData = file_transport_proto_server_proto_rawDesc
)

func file_transport_proto_server_proto_rawDescGZIP() []byte {
	file_transport_proto_server_proto_rawDescOnce.Do(func() {
		file_transport_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_server_proto_rawDescData)
	})
	return file_transport_proto_server_proto_rawDescData
}

var file_transport_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_transport_proto_server_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),  // 0: router.EmptyRequest
	(*EmptyResponse)(nil), // 1: router.EmptyResponse
	(*ByIDRequest)(nil),   // 2: router.ByIDRequest
	(*Product)(nil),       // 3: router.Product
	(*Products)(nil),      // 4: router.Products
	(*Orders)(nil),        // 5: router.Orders
	(*Order)(nil),         // 6: router.Order
	(*Item)(nil),          // 7: router.Item
	(*Payment)(nil),       // 8: router.Payment
	(*Invoice)(nil),       // 9: router.Invoice
}
var file_transport_proto_server_proto_depIdxs = []int32{
	3,  // 0: router.Products.products:type_name -> router.Product
	6,  // 1: router.Orders.orders:type_name -> router.Order
	7,  // 2: router.Order.items:type_name -> router.Item
	8,  // 3: router.Order.payment:type_name -> router.Payment
	3,  // 4: router.Item.product:type_name -> router.Product
	9,  // 5: router.Payment.invoice:type_name -> router.Invoice
	2,  // 6: router.Router.GetOrderByID:input_type -> router.ByIDRequest
	2,  // 7: router.Router.GetOrdersByUserID:input_type -> router.ByIDRequest
	6,  // 8: router.Router.CreateOrder:input_type -> router.Order
	7,  // 9: router.Router.AddItem:input_type -> router.Item
	2,  // 10: router.Router.RemoveItem:input_type -> router.ByIDRequest
	2,  // 11: router.Router.PayOrder:input_type -> router.ByIDRequest
	0,  // 12: router.Router.GetAllProducts:input_type -> router.EmptyRequest
	2,  // 13: router.Router.GetProductByID:input_type -> router.ByIDRequest
	3,  // 14: router.Router.CreateProduct:input_type -> router.Product
	2,  // 15: router.Router.DeleteProduct:input_type -> router.ByIDRequest
	2,  // 16: router.Router.GetPaymentByID:input_type -> router.ByIDRequest
	8,  // 17: router.Router.CreatePayment:input_type -> router.Payment
	6,  // 18: router.Router.GetOrderByID:output_type -> router.Order
	5,  // 19: router.Router.GetOrdersByUserID:output_type -> router.Orders
	6,  // 20: router.Router.CreateOrder:output_type -> router.Order
	6,  // 21: router.Router.AddItem:output_type -> router.Order
	6,  // 22: router.Router.RemoveItem:output_type -> router.Order
	6,  // 23: router.Router.PayOrder:output_type -> router.Order
	4,  // 24: router.Router.GetAllProducts:output_type -> router.Products
	3,  // 25: router.Router.GetProductByID:output_type -> router.Product
	3,  // 26: router.Router.CreateProduct:output_type -> router.Product
	1,  // 27: router.Router.DeleteProduct:output_type -> router.EmptyResponse
	8,  // 28: router.Router.GetPaymentByID:output_type -> router.Payment
	8,  // 29: router.Router.CreatePayment:output_type -> router.Payment
	18, // [18:30] is the sub-list for method output_type
	6,  // [6:18] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_transport_proto_server_proto_init() }
func file_transport_proto_server_proto_init() {
	if File_transport_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_transport_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_transport_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIDRequest); i {
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
		file_transport_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_transport_proto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
		file_transport_proto_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
		file_transport_proto_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_proto_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_proto_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_transport_proto_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
			RawDescriptor: file_transport_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_proto_server_proto_goTypes,
		DependencyIndexes: file_transport_proto_server_proto_depIdxs,
		MessageInfos:      file_transport_proto_server_proto_msgTypes,
	}.Build()
	File_transport_proto_server_proto = out.File
	file_transport_proto_server_proto_rawDesc = nil
	file_transport_proto_server_proto_goTypes = nil
	file_transport_proto_server_proto_depIdxs = nil
}
