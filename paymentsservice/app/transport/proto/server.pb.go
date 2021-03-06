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

type ByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIDRequest) Reset() {
	*x = ByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIDRequest) ProtoMessage() {}

func (x *ByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ByIDRequest.ProtoReflect.Descriptor instead.
func (*ByIDRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *ByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
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
		mi := &file_transport_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{1}
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
		mi := &file_transport_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_transport_proto_server_proto_rawDescGZIP(), []int{2}
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
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x0b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x22, 0x41, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0xe6, 0x01, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x00, 0x42, 0x3e,
	0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x61, 0x42,
	0x0f, 0x4d, 0x6f, 0x6e, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x1a, 0x6d, 0x6f, 0x6e, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_transport_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transport_proto_server_proto_goTypes = []interface{}{
	(*ByIDRequest)(nil), // 0: router.ByIDRequest
	(*Payment)(nil),     // 1: router.Payment
	(*Invoice)(nil),     // 2: router.Invoice
}
var file_transport_proto_server_proto_depIdxs = []int32{
	2, // 0: router.Payment.invoice:type_name -> router.Invoice
	0, // 1: router.Router.GetPaymentByID:input_type -> router.ByIDRequest
	1, // 2: router.Router.CreatePayment:input_type -> router.Payment
	0, // 3: router.Router.GetInvoiceByID:input_type -> router.ByIDRequest
	2, // 4: router.Router.CreateInvoice:input_type -> router.Invoice
	1, // 5: router.Router.GetPaymentByID:output_type -> router.Payment
	1, // 6: router.Router.CreatePayment:output_type -> router.Payment
	2, // 7: router.Router.GetInvoiceByID:output_type -> router.Invoice
	2, // 8: router.Router.CreateInvoice:output_type -> router.Invoice
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_proto_server_proto_init() }
func file_transport_proto_server_proto_init() {
	if File_transport_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   3,
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
