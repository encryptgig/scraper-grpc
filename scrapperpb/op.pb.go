// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.5.1
// source: op.proto

package scrapperpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TokenizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KID  string `protobuf:"bytes,1,opt,name=KID,proto3" json:"KID,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenizationRequest) Reset() {
	*x = TokenizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationRequest) ProtoMessage() {}

func (x *TokenizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_op_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationRequest.ProtoReflect.Descriptor instead.
func (*TokenizationRequest) Descriptor() ([]byte, []int) {
	return file_op_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizationRequest) GetKID() string {
	if x != nil {
		return x.KID
	}
	return ""
}

func (x *TokenizationRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type TokenizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KID  string `protobuf:"bytes,1,opt,name=KID,proto3" json:"KID,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenizationResponse) Reset() {
	*x = TokenizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationResponse) ProtoMessage() {}

func (x *TokenizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_op_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationResponse.ProtoReflect.Descriptor instead.
func (*TokenizationResponse) Descriptor() ([]byte, []int) {
	return file_op_proto_rawDescGZIP(), []int{1}
}

func (x *TokenizationResponse) GetKID() string {
	if x != nil {
		return x.KID
	}
	return ""
}

func (x *TokenizationResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EncryptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KID  string `protobuf:"bytes,1,opt,name=KID,proto3" json:"KID,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptionRequest) Reset() {
	*x = EncryptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionRequest) ProtoMessage() {}

func (x *EncryptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_op_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionRequest.ProtoReflect.Descriptor instead.
func (*EncryptionRequest) Descriptor() ([]byte, []int) {
	return file_op_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptionRequest) GetKID() string {
	if x != nil {
		return x.KID
	}
	return ""
}

func (x *EncryptionRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EncryptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KID  string `protobuf:"bytes,1,opt,name=KID,proto3" json:"KID,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptionResponse) Reset() {
	*x = EncryptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionResponse) ProtoMessage() {}

func (x *EncryptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_op_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionResponse.ProtoReflect.Descriptor instead.
func (*EncryptionResponse) Descriptor() ([]byte, []int) {
	return file_op_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptionResponse) GetKID() string {
	if x != nil {
		return x.KID
	}
	return ""
}

func (x *EncryptionResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_op_proto protoreflect.FileDescriptor

var file_op_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x22, 0x3b, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x39, 0x0a, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x12,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcf, 0x02, 0x0a, 0x11, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f,
	0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0a, 0x44, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x1d, 0x2e,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x63, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_op_proto_rawDescOnce sync.Once
	file_op_proto_rawDescData = file_op_proto_rawDesc
)

func file_op_proto_rawDescGZIP() []byte {
	file_op_proto_rawDescOnce.Do(func() {
		file_op_proto_rawDescData = protoimpl.X.CompressGZIP(file_op_proto_rawDescData)
	})
	return file_op_proto_rawDescData
}

var file_op_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_op_proto_goTypes = []interface{}{
	(*TokenizationRequest)(nil),  // 0: scrapperpb.TokenizationRequest
	(*TokenizationResponse)(nil), // 1: scrapperpb.TokenizationResponse
	(*EncryptionRequest)(nil),    // 2: scrapperpb.EncryptionRequest
	(*EncryptionResponse)(nil),   // 3: scrapperpb.EncryptionResponse
}
var file_op_proto_depIdxs = []int32{
	0, // 0: scrapperpb.EncryptionService.Tokenize:input_type -> scrapperpb.TokenizationRequest
	0, // 1: scrapperpb.EncryptionService.DeTokenize:input_type -> scrapperpb.TokenizationRequest
	2, // 2: scrapperpb.EncryptionService.Encrypt:input_type -> scrapperpb.EncryptionRequest
	2, // 3: scrapperpb.EncryptionService.Decrypt:input_type -> scrapperpb.EncryptionRequest
	1, // 4: scrapperpb.EncryptionService.Tokenize:output_type -> scrapperpb.TokenizationResponse
	1, // 5: scrapperpb.EncryptionService.DeTokenize:output_type -> scrapperpb.TokenizationResponse
	3, // 6: scrapperpb.EncryptionService.Encrypt:output_type -> scrapperpb.EncryptionResponse
	3, // 7: scrapperpb.EncryptionService.Decrypt:output_type -> scrapperpb.EncryptionResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_op_proto_init() }
func file_op_proto_init() {
	if File_op_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_op_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationRequest); i {
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
		file_op_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationResponse); i {
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
		file_op_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionRequest); i {
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
		file_op_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionResponse); i {
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
			RawDescriptor: file_op_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_op_proto_goTypes,
		DependencyIndexes: file_op_proto_depIdxs,
		MessageInfos:      file_op_proto_msgTypes,
	}.Build()
	File_op_proto = out.File
	file_op_proto_rawDesc = nil
	file_op_proto_goTypes = nil
	file_op_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EncryptionServiceClient is the client API for EncryptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EncryptionServiceClient interface {
	Tokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error)
	DeTokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error)
	Encrypt(ctx context.Context, in *EncryptionRequest, opts ...grpc.CallOption) (*EncryptionResponse, error)
	Decrypt(ctx context.Context, in *EncryptionRequest, opts ...grpc.CallOption) (*EncryptionResponse, error)
}

type encryptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncryptionServiceClient(cc grpc.ClientConnInterface) EncryptionServiceClient {
	return &encryptionServiceClient{cc}
}

func (c *encryptionServiceClient) Tokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error) {
	out := new(TokenizationResponse)
	err := c.cc.Invoke(ctx, "/scrapperpb.EncryptionService/Tokenize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptionServiceClient) DeTokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error) {
	out := new(TokenizationResponse)
	err := c.cc.Invoke(ctx, "/scrapperpb.EncryptionService/DeTokenize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptionServiceClient) Encrypt(ctx context.Context, in *EncryptionRequest, opts ...grpc.CallOption) (*EncryptionResponse, error) {
	out := new(EncryptionResponse)
	err := c.cc.Invoke(ctx, "/scrapperpb.EncryptionService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptionServiceClient) Decrypt(ctx context.Context, in *EncryptionRequest, opts ...grpc.CallOption) (*EncryptionResponse, error) {
	out := new(EncryptionResponse)
	err := c.cc.Invoke(ctx, "/scrapperpb.EncryptionService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptionServiceServer is the server API for EncryptionService service.
type EncryptionServiceServer interface {
	Tokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error)
	DeTokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error)
	Encrypt(context.Context, *EncryptionRequest) (*EncryptionResponse, error)
	Decrypt(context.Context, *EncryptionRequest) (*EncryptionResponse, error)
}

// UnimplementedEncryptionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEncryptionServiceServer struct {
}

func (*UnimplementedEncryptionServiceServer) Tokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tokenize not implemented")
}
func (*UnimplementedEncryptionServiceServer) DeTokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeTokenize not implemented")
}
func (*UnimplementedEncryptionServiceServer) Encrypt(context.Context, *EncryptionRequest) (*EncryptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (*UnimplementedEncryptionServiceServer) Decrypt(context.Context, *EncryptionRequest) (*EncryptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func RegisterEncryptionServiceServer(s *grpc.Server, srv EncryptionServiceServer) {
	s.RegisterService(&_EncryptionService_serviceDesc, srv)
}

func _EncryptionService_Tokenize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).Tokenize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapperpb.EncryptionService/Tokenize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).Tokenize(ctx, req.(*TokenizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncryptionService_DeTokenize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).DeTokenize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapperpb.EncryptionService/DeTokenize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).DeTokenize(ctx, req.(*TokenizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncryptionService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapperpb.EncryptionService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).Encrypt(ctx, req.(*EncryptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncryptionService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapperpb.EncryptionService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).Decrypt(ctx, req.(*EncryptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EncryptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scrapperpb.EncryptionService",
	HandlerType: (*EncryptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tokenize",
			Handler:    _EncryptionService_Tokenize_Handler,
		},
		{
			MethodName: "DeTokenize",
			Handler:    _EncryptionService_DeTokenize_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _EncryptionService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _EncryptionService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "op.proto",
}