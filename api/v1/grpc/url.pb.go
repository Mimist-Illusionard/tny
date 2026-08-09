// Code generated from api/v1/grpc/url.proto. DO NOT EDIT.

package urlgrpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateShortLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalUrl   string                 `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShortLinkRequest) Reset() {
	*x = CreateShortLinkRequest{}
	mi := &file_api_v1_grpc_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShortLinkRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateShortLinkRequest) ProtoMessage()    {}

func (x *CreateShortLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_grpc_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateShortLinkRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_grpc_url_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortLinkRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateShortLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortCode     string                 `protobuf:"bytes,2,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	OriginalUrl   string                 `protobuf:"bytes,3,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShortLinkResponse) Reset() {
	*x = CreateShortLinkResponse{}
	mi := &file_api_v1_grpc_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShortLinkResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateShortLinkResponse) ProtoMessage()    {}

func (x *CreateShortLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_grpc_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateShortLinkResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_grpc_url_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortLinkResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateShortLinkResponse) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

func (x *CreateShortLinkResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type GetOriginalLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortCode     string                 `protobuf:"bytes,1,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOriginalLinkRequest) Reset() {
	*x = GetOriginalLinkRequest{}
	mi := &file_api_v1_grpc_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginalLinkRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetOriginalLinkRequest) ProtoMessage()    {}

func (x *GetOriginalLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_grpc_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetOriginalLinkRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_grpc_url_proto_rawDescGZIP(), []int{2}
}

func (x *GetOriginalLinkRequest) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

type GetOriginalLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalUrl   string                 `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOriginalLinkResponse) Reset() {
	*x = GetOriginalLinkResponse{}
	mi := &file_api_v1_grpc_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginalLinkResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetOriginalLinkResponse) ProtoMessage()    {}

func (x *GetOriginalLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_grpc_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetOriginalLinkResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_grpc_url_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginalLinkResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_api_v1_grpc_url_proto protoreflect.FileDescriptor

var file_api_v1_grpc_url_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x72,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x12, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x12, 0x14, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x22, 0x2f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x32, 0xce, 0x01, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x2d, 0x49, 0x6c, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x64, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x3b, 0x75, 0x72, 0x6c, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_grpc_url_proto_rawDescOnce sync.Once
	file_api_v1_grpc_url_proto_rawDescData = file_api_v1_grpc_url_proto_rawDesc
)

func file_api_v1_grpc_url_proto_rawDescGZIP() []byte {
	file_api_v1_grpc_url_proto_rawDescOnce.Do(func() {
		file_api_v1_grpc_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_grpc_url_proto_rawDescData)
	})
	return file_api_v1_grpc_url_proto_rawDescData
}

var file_api_v1_grpc_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_grpc_url_proto_goTypes = []any{
	(*CreateShortLinkRequest)(nil),
	(*CreateShortLinkResponse)(nil),
	(*GetOriginalLinkRequest)(nil),
	(*GetOriginalLinkResponse)(nil),
}
var file_api_v1_grpc_url_proto_depIdxs = []int32{
	0, // 0: shortener.v1.URLShortener.CreateShortLink:input_type -> shortener.v1.CreateShortLinkRequest
	2, // 1: shortener.v1.URLShortener.GetOriginalLink:input_type -> shortener.v1.GetOriginalLinkRequest
	1, // 2: shortener.v1.URLShortener.CreateShortLink:output_type -> shortener.v1.CreateShortLinkResponse
	3, // 3: shortener.v1.URLShortener.GetOriginalLink:output_type -> shortener.v1.GetOriginalLinkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_grpc_url_proto_init() }
func file_api_v1_grpc_url_proto_init() {
	if File_api_v1_grpc_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_grpc_url_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateShortLinkRequest); i {
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
		file_api_v1_grpc_url_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateShortLinkResponse); i {
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
		file_api_v1_grpc_url_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetOriginalLinkRequest); i {
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
		file_api_v1_grpc_url_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetOriginalLinkResponse); i {
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
			RawDescriptor: file_api_v1_grpc_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_grpc_url_proto_goTypes,
		DependencyIndexes: file_api_v1_grpc_url_proto_depIdxs,
		MessageInfos:      file_api_v1_grpc_url_proto_msgTypes,
	}.Build()
	File_api_v1_grpc_url_proto = out.File
	file_api_v1_grpc_url_proto_rawDesc = nil
	file_api_v1_grpc_url_proto_goTypes = nil
	file_api_v1_grpc_url_proto_depIdxs = nil
}
