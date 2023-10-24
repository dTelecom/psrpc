// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: multiple2.proto

// Multiple proto files in one package

package multiple

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

type Msg2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Msg2) Reset() {
	*x = Msg2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiple2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg2) ProtoMessage() {}

func (x *Msg2) ProtoReflect() protoreflect.Message {
	mi := &file_multiple2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg2.ProtoReflect.Descriptor instead.
func (*Msg2) Descriptor() ([]byte, []int) {
	return file_multiple2_proto_rawDescGZIP(), []int{0}
}

var File_multiple2_proto protoreflect.FileDescriptor

var file_multiple2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x1a,
	0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x06, 0x0a, 0x04, 0x4d, 0x73, 0x67, 0x32, 0x32, 0xb8, 0x01, 0x0a, 0x04, 0x53, 0x76, 0x63,
	0x32, 0x12, 0x4e, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x22, 0x2e, 0x70, 0x73, 0x72, 0x70,
	0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x32, 0x1a, 0x22, 0x2e,
	0x70, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x73, 0x67,
	0x32, 0x12, 0x60, 0x0a, 0x16, 0x53, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x73,
	0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x31, 0x1a,
	0x22, 0x2e, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x31, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_multiple2_proto_rawDescOnce sync.Once
	file_multiple2_proto_rawDescData = file_multiple2_proto_rawDesc
)

func file_multiple2_proto_rawDescGZIP() []byte {
	file_multiple2_proto_rawDescOnce.Do(func() {
		file_multiple2_proto_rawDescData = protoimpl.X.CompressGZIP(file_multiple2_proto_rawDescData)
	})
	return file_multiple2_proto_rawDescData
}

var file_multiple2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_multiple2_proto_goTypes = []interface{}{
	(*Msg2)(nil), // 0: psrpc.internal.test.multiple.Msg2
	(*Msg1)(nil), // 1: psrpc.internal.test.multiple.Msg1
}
var file_multiple2_proto_depIdxs = []int32{
	0, // 0: psrpc.internal.test.multiple.Svc2.Send:input_type -> psrpc.internal.test.multiple.Msg2
	1, // 1: psrpc.internal.test.multiple.Svc2.SamePackageProtoImport:input_type -> psrpc.internal.test.multiple.Msg1
	0, // 2: psrpc.internal.test.multiple.Svc2.Send:output_type -> psrpc.internal.test.multiple.Msg2
	1, // 3: psrpc.internal.test.multiple.Svc2.SamePackageProtoImport:output_type -> psrpc.internal.test.multiple.Msg1
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_multiple2_proto_init() }
func file_multiple2_proto_init() {
	if File_multiple2_proto != nil {
		return
	}
	file_multiple1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_multiple2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg2); i {
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
			RawDescriptor: file_multiple2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_multiple2_proto_goTypes,
		DependencyIndexes: file_multiple2_proto_depIdxs,
		MessageInfos:      file_multiple2_proto_msgTypes,
	}.Build()
	File_multiple2_proto = out.File
	file_multiple2_proto_rawDesc = nil
	file_multiple2_proto_goTypes = nil
	file_multiple2_proto_depIdxs = nil
}
