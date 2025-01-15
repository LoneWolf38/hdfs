//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// These .proto interfaces are private and stable.
// Please see http://wiki.apache.org/hadoop/Compatibility
// for what changes are allowed for a *stable* .proto interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: RefreshCallQueueProtocol.proto

package hadoop_common

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

// *
//
//	Refresh callqueue request.
type RefreshCallQueueRequestProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshCallQueueRequestProto) Reset() {
	*x = RefreshCallQueueRequestProto{}
	mi := &file_RefreshCallQueueProtocol_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshCallQueueRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshCallQueueRequestProto) ProtoMessage() {}

func (x *RefreshCallQueueRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshCallQueueProtocol_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshCallQueueRequestProto.ProtoReflect.Descriptor instead.
func (*RefreshCallQueueRequestProto) Descriptor() ([]byte, []int) {
	return file_RefreshCallQueueProtocol_proto_rawDescGZIP(), []int{0}
}

// *
// void response.
type RefreshCallQueueResponseProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshCallQueueResponseProto) Reset() {
	*x = RefreshCallQueueResponseProto{}
	mi := &file_RefreshCallQueueProtocol_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshCallQueueResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshCallQueueResponseProto) ProtoMessage() {}

func (x *RefreshCallQueueResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshCallQueueProtocol_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshCallQueueResponseProto.ProtoReflect.Descriptor instead.
func (*RefreshCallQueueResponseProto) Descriptor() ([]byte, []int) {
	return file_RefreshCallQueueProtocol_proto_rawDescGZIP(), []int{1}
}

var File_RefreshCallQueueProtocol_proto protoreflect.FileDescriptor

var file_RefreshCallQueueProtocol_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22,
	0x1e, 0x0a, 0x1c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1f, 0x0a, 0x1d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x90, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43,
	0x61, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x2b, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x43, 0x61, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c,
	0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x81, 0x01, 0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x1e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x6f, 0x6e, 0x65, 0x57, 0x6f, 0x6c, 0x66, 0x33, 0x38, 0x2f, 0x67, 0x6f, 0x68, 0x64, 0x66,
	0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_RefreshCallQueueProtocol_proto_rawDescOnce sync.Once
	file_RefreshCallQueueProtocol_proto_rawDescData = file_RefreshCallQueueProtocol_proto_rawDesc
)

func file_RefreshCallQueueProtocol_proto_rawDescGZIP() []byte {
	file_RefreshCallQueueProtocol_proto_rawDescOnce.Do(func() {
		file_RefreshCallQueueProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_RefreshCallQueueProtocol_proto_rawDescData)
	})
	return file_RefreshCallQueueProtocol_proto_rawDescData
}

var file_RefreshCallQueueProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RefreshCallQueueProtocol_proto_goTypes = []any{
	(*RefreshCallQueueRequestProto)(nil),  // 0: hadoop.common.RefreshCallQueueRequestProto
	(*RefreshCallQueueResponseProto)(nil), // 1: hadoop.common.RefreshCallQueueResponseProto
}
var file_RefreshCallQueueProtocol_proto_depIdxs = []int32{
	0, // 0: hadoop.common.RefreshCallQueueProtocolService.refreshCallQueue:input_type -> hadoop.common.RefreshCallQueueRequestProto
	1, // 1: hadoop.common.RefreshCallQueueProtocolService.refreshCallQueue:output_type -> hadoop.common.RefreshCallQueueResponseProto
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RefreshCallQueueProtocol_proto_init() }
func file_RefreshCallQueueProtocol_proto_init() {
	if File_RefreshCallQueueProtocol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RefreshCallQueueProtocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RefreshCallQueueProtocol_proto_goTypes,
		DependencyIndexes: file_RefreshCallQueueProtocol_proto_depIdxs,
		MessageInfos:      file_RefreshCallQueueProtocol_proto_msgTypes,
	}.Build()
	File_RefreshCallQueueProtocol_proto = out.File
	file_RefreshCallQueueProtocol_proto_rawDesc = nil
	file_RefreshCallQueueProtocol_proto_goTypes = nil
	file_RefreshCallQueueProtocol_proto_depIdxs = nil
}
