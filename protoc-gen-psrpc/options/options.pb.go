// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: options.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For RPCs, each client request will receive a response from every server.
	// For subscriptions, every client will receive every update.
	Multi bool `protobuf:"varint,1,opt,name=multi,proto3" json:"multi,omitempty"`
	// This method is a pub/sub.
	Subscription bool `protobuf:"varint,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// This method uses topics.
	Topics      bool               `protobuf:"varint,3,opt,name=topics,proto3" json:"topics,omitempty"`
	TopicParams *TopicParamOptions `protobuf:"bytes,4,opt,name=topic_params,json=topicParams,proto3" json:"topic_params,omitempty"`
	// Your service will supply an affinity function for handler selection.
	AffinityFunc bool `protobuf:"varint,5,opt,name=affinity_func,json=affinityFunc,proto3" json:"affinity_func,omitempty"`
	// The method uses bidirectional streaming.
	Stream bool `protobuf:"varint,6,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetMulti() bool {
	if x != nil {
		return x.Multi
	}
	return false
}

func (x *Options) GetSubscription() bool {
	if x != nil {
		return x.Subscription
	}
	return false
}

func (x *Options) GetTopics() bool {
	if x != nil {
		return x.Topics
	}
	return false
}

func (x *Options) GetTopicParams() *TopicParamOptions {
	if x != nil {
		return x.TopicParams
	}
	return nil
}

func (x *Options) GetAffinityFunc() bool {
	if x != nil {
		return x.AffinityFunc
	}
	return false
}

func (x *Options) GetStream() bool {
	if x != nil {
		return x.Stream
	}
	return false
}

type TopicParamOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rpc can be registered/deregistered atomically with other group members
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The topic is composed of one or more string-like parameters.
	Names []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	// The topic parameters have associated string-like type parameters
	Typed bool `protobuf:"varint,3,opt,name=typed,proto3" json:"typed,omitempty"`
	// At most one server will be registered for each topic
	SingleServer bool `protobuf:"varint,4,opt,name=single_server,json=singleServer,proto3" json:"single_server,omitempty"`
}

func (x *TopicParamOptions) Reset() {
	*x = TopicParamOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicParamOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicParamOptions) ProtoMessage() {}

func (x *TopicParamOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicParamOptions.ProtoReflect.Descriptor instead.
func (*TopicParamOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *TopicParamOptions) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *TopicParamOptions) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *TopicParamOptions) GetTyped() bool {
	if x != nil {
		return x.Typed
	}
	return false
}

func (x *TopicParamOptions) GetSingleServer() bool {
	if x != nil {
		return x.SingleServer
	}
	return false
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         2198,
		Name:          "psrpc.options",
		Tag:           "bytes,2198,opt,name=options",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional psrpc.Options options = 2198;
	E_Options = &file_options_proto_extTypes[0]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x73, 0x72, 0x70, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x73, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x5f,
	0x66, 0x75, 0x6e, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x22, 0x7a, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x4c, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x96, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2f, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_options_proto_goTypes = []interface{}{
	(*Options)(nil),                    // 0: psrpc.Options
	(*TopicParamOptions)(nil),          // 1: psrpc.TopicParamOptions
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_options_proto_depIdxs = []int32{
	1, // 0: psrpc.Options.topic_params:type_name -> psrpc.TopicParamOptions
	2, // 1: psrpc.options:extendee -> google.protobuf.MethodOptions
	0, // 2: psrpc.options:type_name -> psrpc.Options
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicParamOptions); i {
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
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
