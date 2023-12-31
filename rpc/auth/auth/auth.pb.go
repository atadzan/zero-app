// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: auth.proto

package auth

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

type Hello int32

const (
	Hello_Priwet Hello = 0
	Hello_Salam  Hello = 1
)

// Enum value maps for Hello.
var (
	Hello_name = map[int32]string{
		0: "Priwet",
		1: "Salam",
	}
	Hello_value = map[string]int32{
		"Priwet": 0,
		"Salam":  1,
	}
)

func (x Hello) Enum() *Hello {
	p := new(Hello)
	*p = x
	return p
}

func (x Hello) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hello) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (Hello) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x Hello) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hello.Descriptor instead.
func (Hello) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Age       int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *SignUpReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SignUpReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type SignUpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignUpRes) Reset() {
	*x = SignUpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRes) ProtoMessage() {}

func (x *SignUpRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRes.ProtoReflect.Descriptor instead.
func (*SignUpRes) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignInReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignInRes) Reset() {
	*x = SignInRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRes) ProtoMessage() {}

func (x *SignInRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRes.ProtoReflect.Descriptor instead.
func (*SignInRes) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignInRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EnumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input Hello `protobuf:"varint,1,opt,name=input,proto3,enum=auth.Hello" json:"input,omitempty"`
}

func (x *EnumReq) Reset() {
	*x = EnumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumReq) ProtoMessage() {}

func (x *EnumReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumReq.ProtoReflect.Descriptor instead.
func (*EnumReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *EnumReq) GetInput() Hello {
	if x != nil {
		return x.Input
	}
	return Hello_Priwet
}

type EnumRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EnumRes) Reset() {
	*x = EnumRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumRes) ProtoMessage() {}

func (x *EnumRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumRes.ProtoReflect.Descriptor instead.
func (*EnumRes) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *EnumRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x8f, 0x01, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x21, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x21,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x23, 0x0a, 0x07, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1e, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x77, 0x65, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x61, 0x6c, 0x61, 0x6d, 0x10, 0x01, 0x32, 0x88, 0x01, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12,
	0x2a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(Hello)(0),        // 0: auth.Hello
	(*SignUpReq)(nil), // 1: auth.signUpReq
	(*SignUpRes)(nil), // 2: auth.signUpRes
	(*SignInReq)(nil), // 3: auth.signInReq
	(*SignInRes)(nil), // 4: auth.signInRes
	(*EnumReq)(nil),   // 5: auth.enumReq
	(*EnumRes)(nil),   // 6: auth.enumRes
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: auth.enumReq.input:type_name -> auth.Hello
	3, // 1: auth.auth.SignIn:input_type -> auth.signInReq
	1, // 2: auth.auth.SignUp:input_type -> auth.signUpReq
	5, // 3: auth.auth.EnumTest:input_type -> auth.enumReq
	4, // 4: auth.auth.SignIn:output_type -> auth.signInRes
	2, // 5: auth.auth.SignUp:output_type -> auth.signUpRes
	6, // 6: auth.auth.EnumTest:output_type -> auth.enumRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRes); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRes); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumReq); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumRes); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
