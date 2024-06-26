// TODO это вымышленный пример proto-файла, удалите директорию grpc в реальном приложении
// руководство по описанию протофайлов - https://github.com/uber/prototool
// документация для Proto3 - https://developers.google.com/protocol-buffers/docs/proto3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: citilink/blog/article/v1/article.proto

package articlev1

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

// Category категория.
type Category int32

const (
	Category_CATEGORY_INVALID Category = 0
	Category_CATEGORY_PEOPLE  Category = 1
	Category_CATEGORY_ANIMALS Category = 2
	Category_CATEGORY_TRAVELS Category = 3
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "CATEGORY_INVALID",
		1: "CATEGORY_PEOPLE",
		2: "CATEGORY_ANIMALS",
		3: "CATEGORY_TRAVELS",
	}
	Category_value = map[string]int32{
		"CATEGORY_INVALID": 0,
		"CATEGORY_PEOPLE":  1,
		"CATEGORY_ANIMALS": 2,
		"CATEGORY_TRAVELS": 3,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_citilink_blog_article_v1_article_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_citilink_blog_article_v1_article_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_citilink_blog_article_v1_article_proto_rawDescGZIP(), []int{0}
}

// Article статья.
type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Название
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Контент
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Идентификатор категории
	Category Category `protobuf:"varint,4,opt,name=category,proto3,enum=citilink.blog.article.v1.Category" json:"category,omitempty"`
	// Теги
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// Видимость
	IsVisible bool `protobuf:"varint,6,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citilink_blog_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_citilink_blog_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_citilink_blog_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_CATEGORY_INVALID
}

func (x *Article) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Article) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

var File_citilink_blog_article_v1_article_proto protoreflect.FileDescriptor

var file_citilink_blog_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3e,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x2a, 0x61, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f,
	0x50, 0x45, 0x4f, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x4e, 0x49, 0x4d, 0x41, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45,
	0x4c, 0x53, 0x10, 0x03, 0x42, 0x89, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74,
	0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x6f, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x6b,
	0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x69, 0x74,
	0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x42, 0x41, 0xaa, 0x02, 0x18, 0x43, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x18, 0x43, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x5c, 0x42, 0x6c, 0x6f, 0x67,
	0x5c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x43, 0x69,
	0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x5c, 0x42, 0x6c, 0x6f, 0x67, 0x5c, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x42,
	0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_citilink_blog_article_v1_article_proto_rawDescOnce sync.Once
	file_citilink_blog_article_v1_article_proto_rawDescData = file_citilink_blog_article_v1_article_proto_rawDesc
)

func file_citilink_blog_article_v1_article_proto_rawDescGZIP() []byte {
	file_citilink_blog_article_v1_article_proto_rawDescOnce.Do(func() {
		file_citilink_blog_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_citilink_blog_article_v1_article_proto_rawDescData)
	})
	return file_citilink_blog_article_v1_article_proto_rawDescData
}

var file_citilink_blog_article_v1_article_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_citilink_blog_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_citilink_blog_article_v1_article_proto_goTypes = []interface{}{
	(Category)(0),   // 0: citilink.blog.article.v1.Category
	(*Article)(nil), // 1: citilink.blog.article.v1.Article
}
var file_citilink_blog_article_v1_article_proto_depIdxs = []int32{
	0, // 0: citilink.blog.article.v1.Article.category:type_name -> citilink.blog.article.v1.Category
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_citilink_blog_article_v1_article_proto_init() }
func file_citilink_blog_article_v1_article_proto_init() {
	if File_citilink_blog_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_citilink_blog_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
			RawDescriptor: file_citilink_blog_article_v1_article_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_citilink_blog_article_v1_article_proto_goTypes,
		DependencyIndexes: file_citilink_blog_article_v1_article_proto_depIdxs,
		EnumInfos:         file_citilink_blog_article_v1_article_proto_enumTypes,
		MessageInfos:      file_citilink_blog_article_v1_article_proto_msgTypes,
	}.Build()
	File_citilink_blog_article_v1_article_proto = out.File
	file_citilink_blog_article_v1_article_proto_rawDesc = nil
	file_citilink_blog_article_v1_article_proto_goTypes = nil
	file_citilink_blog_article_v1_article_proto_depIdxs = nil
}
