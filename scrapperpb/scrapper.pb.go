// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.5.1
// source: scrapper.proto

package scrapperpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	URI       string `protobuf:"bytes,2,opt,name=URI,proto3" json:"URI,omitempty"`
	Account   string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_scrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_scrapper_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Resource) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *Resource) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Resource) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type KeyBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource            *Resource       `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
	KeyFormatType       KeyFormatType   `protobuf:"varint,2,opt,name=KeyFormatType,proto3,enum=scrapperpb.KeyFormatType" json:"KeyFormatType,omitempty"`
	Material            string          `protobuf:"bytes,3,opt,name=Material,proto3" json:"Material,omitempty"`
	CryptoAlgorithm     CryptoAlgorithm `protobuf:"varint,4,opt,name=CryptoAlgorithm,proto3,enum=scrapperpb.CryptoAlgorithm" json:"CryptoAlgorithm,omitempty"`
	CryptographicLength int32           `protobuf:"varint,5,opt,name=CryptographicLength,proto3" json:"CryptographicLength,omitempty"`
	Name                []string        `protobuf:"bytes,6,rep,name=Name,proto3" json:"Name,omitempty"`
	KeyAttributes       *KeyAttributes  `protobuf:"bytes,7,opt,name=KeyAttributes,proto3" json:"KeyAttributes,omitempty"`
	KeyOwners           []*Owners       `protobuf:"bytes,8,rep,name=KeyOwners,proto3" json:"KeyOwners,omitempty"`
	KeyUsers            []*KeyUsers     `protobuf:"bytes,9,rep,name=KeyUsers,proto3" json:"KeyUsers,omitempty"` //wrapping data is pending
}

func (x *KeyBlock) Reset() {
	*x = KeyBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyBlock) ProtoMessage() {}

func (x *KeyBlock) ProtoReflect() protoreflect.Message {
	mi := &file_scrapper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyBlock.ProtoReflect.Descriptor instead.
func (*KeyBlock) Descriptor() ([]byte, []int) {
	return file_scrapper_proto_rawDescGZIP(), []int{1}
}

func (x *KeyBlock) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *KeyBlock) GetKeyFormatType() KeyFormatType {
	if x != nil {
		return x.KeyFormatType
	}
	return KeyFormatType_KeyFormatTypeRaw
}

func (x *KeyBlock) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

func (x *KeyBlock) GetCryptoAlgorithm() CryptoAlgorithm {
	if x != nil {
		return x.CryptoAlgorithm
	}
	return CryptoAlgorithm_NONE
}

func (x *KeyBlock) GetCryptographicLength() int32 {
	if x != nil {
		return x.CryptographicLength
	}
	return 0
}

func (x *KeyBlock) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *KeyBlock) GetKeyAttributes() *KeyAttributes {
	if x != nil {
		return x.KeyAttributes
	}
	return nil
}

func (x *KeyBlock) GetKeyOwners() []*Owners {
	if x != nil {
		return x.KeyOwners
	}
	return nil
}

func (x *KeyBlock) GetKeyUsers() []*KeyUsers {
	if x != nil {
		return x.KeyUsers
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_scrapper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_scrapper_proto_rawDescGZIP(), []int{2}
}

var File_scrapper_proto protoreflect.FileDescriptor

var file_scrapper_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x10, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x49, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xcb, 0x03, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x30,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x45, 0x0a,
	0x0f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x0f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x4b, 0x65,
	0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4b,
	0x65, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0d, 0x4b, 0x65,
	0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x4b,
	0x65, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x09, 0x4b, 0x65, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x11, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_scrapper_proto_rawDescOnce sync.Once
	file_scrapper_proto_rawDescData = file_scrapper_proto_rawDesc
)

func file_scrapper_proto_rawDescGZIP() []byte {
	file_scrapper_proto_rawDescOnce.Do(func() {
		file_scrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_scrapper_proto_rawDescData)
	})
	return file_scrapper_proto_rawDescData
}

var file_scrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scrapper_proto_goTypes = []interface{}{
	(*Resource)(nil),      // 0: scrapperpb.Resource
	(*KeyBlock)(nil),      // 1: scrapperpb.KeyBlock
	(*Empty)(nil),         // 2: scrapperpb.Empty
	(KeyFormatType)(0),    // 3: scrapperpb.KeyFormatType
	(CryptoAlgorithm)(0),  // 4: scrapperpb.CryptoAlgorithm
	(*KeyAttributes)(nil), // 5: scrapperpb.KeyAttributes
	(*Owners)(nil),        // 6: scrapperpb.Owners
	(*KeyUsers)(nil),      // 7: scrapperpb.KeyUsers
}
var file_scrapper_proto_depIdxs = []int32{
	0, // 0: scrapperpb.KeyBlock.Resource:type_name -> scrapperpb.Resource
	3, // 1: scrapperpb.KeyBlock.KeyFormatType:type_name -> scrapperpb.KeyFormatType
	4, // 2: scrapperpb.KeyBlock.CryptoAlgorithm:type_name -> scrapperpb.CryptoAlgorithm
	5, // 3: scrapperpb.KeyBlock.KeyAttributes:type_name -> scrapperpb.KeyAttributes
	6, // 4: scrapperpb.KeyBlock.KeyOwners:type_name -> scrapperpb.Owners
	7, // 5: scrapperpb.KeyBlock.KeyUsers:type_name -> scrapperpb.KeyUsers
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scrapper_proto_init() }
func file_scrapper_proto_init() {
	if File_scrapper_proto != nil {
		return
	}
	file_attributes_proto_init()
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_scrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_scrapper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyBlock); i {
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
		file_scrapper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_scrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scrapper_proto_goTypes,
		DependencyIndexes: file_scrapper_proto_depIdxs,
		MessageInfos:      file_scrapper_proto_msgTypes,
	}.Build()
	File_scrapper_proto = out.File
	file_scrapper_proto_rawDesc = nil
	file_scrapper_proto_goTypes = nil
	file_scrapper_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScrapperServiceClient is the client API for ScrapperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScrapperServiceClient interface {
}

type scrapperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScrapperServiceClient(cc grpc.ClientConnInterface) ScrapperServiceClient {
	return &scrapperServiceClient{cc}
}

// ScrapperServiceServer is the server API for ScrapperService service.
type ScrapperServiceServer interface {
}

// UnimplementedScrapperServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScrapperServiceServer struct {
}

func RegisterScrapperServiceServer(s *grpc.Server, srv ScrapperServiceServer) {
	s.RegisterService(&_ScrapperService_serviceDesc, srv)
}

var _ScrapperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scrapperpb.ScrapperService",
	HandlerType: (*ScrapperServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "scrapper.proto",
}
