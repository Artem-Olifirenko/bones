// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: citilink/store/stock/v1/stock_api.proto

package stockv1

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

type FindAllBySpaceIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
}

func (x *FindAllBySpaceIdRequest) Reset() {
	*x = FindAllBySpaceIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citilink_store_stock_v1_stock_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllBySpaceIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllBySpaceIdRequest) ProtoMessage() {}

func (x *FindAllBySpaceIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_citilink_store_stock_v1_stock_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllBySpaceIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllBySpaceIdRequest) Descriptor() ([]byte, []int) {
	return file_citilink_store_stock_v1_stock_api_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllBySpaceIdRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

type FindAllBySpaceIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockInfo []*StockInfo `protobuf:"bytes,1,rep,name=stock_info,json=stockInfo,proto3" json:"stock_info,omitempty"`
}

func (x *FindAllBySpaceIdResponse) Reset() {
	*x = FindAllBySpaceIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citilink_store_stock_v1_stock_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllBySpaceIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllBySpaceIdResponse) ProtoMessage() {}

func (x *FindAllBySpaceIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_citilink_store_stock_v1_stock_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllBySpaceIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllBySpaceIdResponse) Descriptor() ([]byte, []int) {
	return file_citilink_store_stock_v1_stock_api_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllBySpaceIdResponse) GetStockInfo() []*StockInfo {
	if x != nil {
		return x.StockInfo
	}
	return nil
}

var File_citilink_store_stock_v1_stock_api_proto protoreflect.FileDescriptor

var file_citilink_store_stock_v1_stock_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x69, 0x74, 0x69, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x1a, 0x23, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x42, 0x79, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x5d, 0x0a,
	0x18, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x83, 0x01, 0x0a,
	0x08, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x12, 0x77, 0x0a, 0x10, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x30, 0x2e,
	0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42,
	0x79, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x42, 0x79, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x8a, 0x02, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x6f, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x6b, 0x65, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x53, 0xaa, 0x02, 0x17, 0x43, 0x69, 0x74, 0x69, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x17, 0x43, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x43,
	0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_citilink_store_stock_v1_stock_api_proto_rawDescOnce sync.Once
	file_citilink_store_stock_v1_stock_api_proto_rawDescData = file_citilink_store_stock_v1_stock_api_proto_rawDesc
)

func file_citilink_store_stock_v1_stock_api_proto_rawDescGZIP() []byte {
	file_citilink_store_stock_v1_stock_api_proto_rawDescOnce.Do(func() {
		file_citilink_store_stock_v1_stock_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_citilink_store_stock_v1_stock_api_proto_rawDescData)
	})
	return file_citilink_store_stock_v1_stock_api_proto_rawDescData
}

var file_citilink_store_stock_v1_stock_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_citilink_store_stock_v1_stock_api_proto_goTypes = []interface{}{
	(*FindAllBySpaceIdRequest)(nil),  // 0: citilink.store.stock.v1.FindAllBySpaceIdRequest
	(*FindAllBySpaceIdResponse)(nil), // 1: citilink.store.stock.v1.FindAllBySpaceIdResponse
	(*StockInfo)(nil),                // 2: citilink.store.stock.v1.StockInfo
}
var file_citilink_store_stock_v1_stock_api_proto_depIdxs = []int32{
	2, // 0: citilink.store.stock.v1.FindAllBySpaceIdResponse.stock_info:type_name -> citilink.store.stock.v1.StockInfo
	0, // 1: citilink.store.stock.v1.StockAPI.FindAllBySpaceId:input_type -> citilink.store.stock.v1.FindAllBySpaceIdRequest
	1, // 2: citilink.store.stock.v1.StockAPI.FindAllBySpaceId:output_type -> citilink.store.stock.v1.FindAllBySpaceIdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_citilink_store_stock_v1_stock_api_proto_init() }
func file_citilink_store_stock_v1_stock_api_proto_init() {
	if File_citilink_store_stock_v1_stock_api_proto != nil {
		return
	}
	file_citilink_store_stock_v1_stock_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_citilink_store_stock_v1_stock_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllBySpaceIdRequest); i {
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
		file_citilink_store_stock_v1_stock_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllBySpaceIdResponse); i {
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
			RawDescriptor: file_citilink_store_stock_v1_stock_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_citilink_store_stock_v1_stock_api_proto_goTypes,
		DependencyIndexes: file_citilink_store_stock_v1_stock_api_proto_depIdxs,
		MessageInfos:      file_citilink_store_stock_v1_stock_api_proto_msgTypes,
	}.Build()
	File_citilink_store_stock_v1_stock_api_proto = out.File
	file_citilink_store_stock_v1_stock_api_proto_rawDesc = nil
	file_citilink_store_stock_v1_stock_api_proto_goTypes = nil
	file_citilink_store_stock_v1_stock_api_proto_depIdxs = nil
}
