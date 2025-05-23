// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: sortify/sortify.proto

package sortify_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthByCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State         string                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthByCodeRequest) Reset() {
	*x = AuthByCodeRequest{}
	mi := &file_sortify_sortify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthByCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByCodeRequest) ProtoMessage() {}

func (x *AuthByCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sortify_sortify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByCodeRequest.ProtoReflect.Descriptor instead.
func (*AuthByCodeRequest) Descriptor() ([]byte, []int) {
	return file_sortify_sortify_proto_rawDescGZIP(), []int{0}
}

func (x *AuthByCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AuthByCodeRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthByCodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthByCodeResponse) Reset() {
	*x = AuthByCodeResponse{}
	mi := &file_sortify_sortify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthByCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByCodeResponse) ProtoMessage() {}

func (x *AuthByCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sortify_sortify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByCodeResponse.ProtoReflect.Descriptor instead.
func (*AuthByCodeResponse) Descriptor() ([]byte, []int) {
	return file_sortify_sortify_proto_rawDescGZIP(), []int{1}
}

func (x *AuthByCodeResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_sortify_sortify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sortify_sortify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_sortify_sortify_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DisplayName   string                 `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	SpotifyId     string                 `protobuf:"bytes,3,opt,name=spotify_id,json=spotifyId,proto3" json:"spotify_id,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_sortify_sortify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sortify_sortify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_sortify_sortify_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetSpotifyId() string {
	if x != nil {
		return x.SpotifyId
	}
	return ""
}

func (x *UserResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_sortify_sortify_proto protoreflect.FileDescriptor

const file_sortify_sortify_proto_rawDesc = "" +
	"\n" +
	"\x15sortify/sortify.proto\x12\asortify\"=\n" +
	"\x11AuthByCodeRequest\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x14\n" +
	"\x05state\x18\x02 \x01(\tR\x05state\"-\n" +
	"\x12AuthByCodeResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"&\n" +
	"\vUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"\x83\x01\n" +
	"\fUserResponse\x12!\n" +
	"\fdisplay_name\x18\x01 \x01(\tR\vdisplayName\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1d\n" +
	"\n" +
	"spotify_id\x18\x03 \x01(\tR\tspotifyId\x12\x1b\n" +
	"\timage_url\x18\x04 \x01(\tR\bimageUrl2\x85\x01\n" +
	"\aSortify\x12E\n" +
	"\n" +
	"AuthByCode\x12\x1a.sortify.AuthByCodeRequest\x1a\x1b.sortify.AuthByCodeResponse\x123\n" +
	"\x04User\x12\x14.sortify.UserRequest\x1a\x15.sortify.UserResponseB\fZ\n" +
	"sortify.v1b\x06proto3"

var (
	file_sortify_sortify_proto_rawDescOnce sync.Once
	file_sortify_sortify_proto_rawDescData []byte
)

func file_sortify_sortify_proto_rawDescGZIP() []byte {
	file_sortify_sortify_proto_rawDescOnce.Do(func() {
		file_sortify_sortify_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sortify_sortify_proto_rawDesc), len(file_sortify_sortify_proto_rawDesc)))
	})
	return file_sortify_sortify_proto_rawDescData
}

var file_sortify_sortify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sortify_sortify_proto_goTypes = []any{
	(*AuthByCodeRequest)(nil),  // 0: sortify.AuthByCodeRequest
	(*AuthByCodeResponse)(nil), // 1: sortify.AuthByCodeResponse
	(*UserRequest)(nil),        // 2: sortify.UserRequest
	(*UserResponse)(nil),       // 3: sortify.UserResponse
}
var file_sortify_sortify_proto_depIdxs = []int32{
	0, // 0: sortify.Sortify.AuthByCode:input_type -> sortify.AuthByCodeRequest
	2, // 1: sortify.Sortify.User:input_type -> sortify.UserRequest
	1, // 2: sortify.Sortify.AuthByCode:output_type -> sortify.AuthByCodeResponse
	3, // 3: sortify.Sortify.User:output_type -> sortify.UserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sortify_sortify_proto_init() }
func file_sortify_sortify_proto_init() {
	if File_sortify_sortify_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sortify_sortify_proto_rawDesc), len(file_sortify_sortify_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sortify_sortify_proto_goTypes,
		DependencyIndexes: file_sortify_sortify_proto_depIdxs,
		MessageInfos:      file_sortify_sortify_proto_msgTypes,
	}.Build()
	File_sortify_sortify_proto = out.File
	file_sortify_sortify_proto_goTypes = nil
	file_sortify_sortify_proto_depIdxs = nil
}
