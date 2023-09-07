// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: rpc_create_user_details.proto

package pb

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

type CreateUserDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateUserDetailsRequest) Reset() {
	*x = CreateUserDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_user_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserDetailsRequest) ProtoMessage() {}

func (x *CreateUserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserDetailsRequest.ProtoReflect.Descriptor instead.
func (*CreateUserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_details_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserDetailsRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CreateUserDetailsRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserDetailsRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type CreateUserDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserDetailsResponse) Reset() {
	*x = CreateUserDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_user_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserDetailsResponse) ProtoMessage() {}

func (x *CreateUserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserDetailsResponse.ProtoReflect.Descriptor instead.
func (*CreateUserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_details_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserDetailsResponse) GetUser() *UserDetails {
	if x != nil {
		return x.User
	}
	return nil
}

type FetchUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchUserParams) Reset() {
	*x = FetchUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_user_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserParams) ProtoMessage() {}

func (x *FetchUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserParams.ProtoReflect.Descriptor instead.
func (*FetchUserParams) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_details_proto_rawDescGZIP(), []int{2}
}

type FetchUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	IsSignUpd bool         `protobuf:"varint,2,opt,name=isSignUpd,proto3" json:"isSignUpd,omitempty"`
}

func (x *FetchUserResponse) Reset() {
	*x = FetchUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_user_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserResponse) ProtoMessage() {}

func (x *FetchUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserResponse.ProtoReflect.Descriptor instead.
func (*FetchUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_details_proto_rawDescGZIP(), []int{3}
}

func (x *FetchUserResponse) GetUser() *UserDetails {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FetchUserResponse) GetIsSignUpd() bool {
	if x != nil {
		return x.IsSignUpd
	}
	return false
}

var File_rpc_create_user_details_proto protoreflect.FileDescriptor

var file_rpc_create_user_details_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x56, 0x0a, 0x11, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x2d, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x74, 0x65, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_user_details_proto_rawDescOnce sync.Once
	file_rpc_create_user_details_proto_rawDescData = file_rpc_create_user_details_proto_rawDesc
)

func file_rpc_create_user_details_proto_rawDescGZIP() []byte {
	file_rpc_create_user_details_proto_rawDescOnce.Do(func() {
		file_rpc_create_user_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_user_details_proto_rawDescData)
	})
	return file_rpc_create_user_details_proto_rawDescData
}

var file_rpc_create_user_details_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_create_user_details_proto_goTypes = []interface{}{
	(*CreateUserDetailsRequest)(nil),  // 0: pb.CreateUserDetailsRequest
	(*CreateUserDetailsResponse)(nil), // 1: pb.CreateUserDetailsResponse
	(*FetchUserParams)(nil),           // 2: pb.FetchUserParams
	(*FetchUserResponse)(nil),         // 3: pb.FetchUserResponse
	(*UserDetails)(nil),               // 4: pb.UserDetails
}
var file_rpc_create_user_details_proto_depIdxs = []int32{
	4, // 0: pb.CreateUserDetailsResponse.user:type_name -> pb.UserDetails
	4, // 1: pb.FetchUserResponse.user:type_name -> pb.UserDetails
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_create_user_details_proto_init() }
func file_rpc_create_user_details_proto_init() {
	if File_rpc_create_user_details_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_user_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserDetailsRequest); i {
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
		file_rpc_create_user_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserDetailsResponse); i {
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
		file_rpc_create_user_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserParams); i {
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
		file_rpc_create_user_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserResponse); i {
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
			RawDescriptor: file_rpc_create_user_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_user_details_proto_goTypes,
		DependencyIndexes: file_rpc_create_user_details_proto_depIdxs,
		MessageInfos:      file_rpc_create_user_details_proto_msgTypes,
	}.Build()
	File_rpc_create_user_details_proto = out.File
	file_rpc_create_user_details_proto_rawDesc = nil
	file_rpc_create_user_details_proto_goTypes = nil
	file_rpc_create_user_details_proto_depIdxs = nil
}
