// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: user.proto

package gen

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x7d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x29, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x81, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x12, 0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x32, 0x8b, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x76, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x82, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x32, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x75, 0x72, 0x79, 0x61, 0x2d, 0x37, 0x38, 0x39, 0x30, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_proto_goTypes = []any{
	(*UserLoginRequest)(nil),   // 0: book_store.user_service.UserLoginRequest
	(*UserSignupRequest)(nil),  // 1: book_store.user_service.UserSignupRequest
	(*GetUserRequest)(nil),     // 2: book_store.user_service.GetUserRequest
	(*UpdateUserRequest)(nil),  // 3: book_store.user_service.UpdateUserRequest
	(*DeleteUserRequest)(nil),  // 4: book_store.user_service.DeleteUserRequest
	(*UserLoginResponse)(nil),  // 5: book_store.user_service.UserLoginResponse
	(*UserSignupResponse)(nil), // 6: book_store.user_service.UserSignupResponse
	(*GetUserResponse)(nil),    // 7: book_store.user_service.GetUserResponse
	(*UpdateUserResponse)(nil), // 8: book_store.user_service.UpdateUserResponse
	(*DeleteUserResponse)(nil), // 9: book_store.user_service.DeleteUserResponse
}
var file_user_proto_depIdxs = []int32{
	0, // 0: book_store.user_service.UserAuth.UserLogin:input_type -> book_store.user_service.UserLoginRequest
	1, // 1: book_store.user_service.UserAuth.UserSignup:input_type -> book_store.user_service.UserSignupRequest
	2, // 2: book_store.user_service.UserProfile.GetUser:input_type -> book_store.user_service.GetUserRequest
	3, // 3: book_store.user_service.UserProfile.UpdateUser:input_type -> book_store.user_service.UpdateUserRequest
	4, // 4: book_store.user_service.UserProfile.DeleteUser:input_type -> book_store.user_service.DeleteUserRequest
	5, // 5: book_store.user_service.UserAuth.UserLogin:output_type -> book_store.user_service.UserLoginResponse
	6, // 6: book_store.user_service.UserAuth.UserSignup:output_type -> book_store.user_service.UserSignupResponse
	7, // 7: book_store.user_service.UserProfile.GetUser:output_type -> book_store.user_service.GetUserResponse
	8, // 8: book_store.user_service.UserProfile.UpdateUser:output_type -> book_store.user_service.UpdateUserResponse
	9, // 9: book_store.user_service.UserProfile.DeleteUser:output_type -> book_store.user_service.DeleteUserResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
