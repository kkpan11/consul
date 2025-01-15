// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: annotations/ratelimit/ratelimit.proto

package ratelimit

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

// OperationType determines the kind of rate limit that will be applied to this
// RPC (i.e. read or write).
type OperationType int32

const (
	OperationType_OPERATION_TYPE_UNSPECIFIED OperationType = 0
	OperationType_OPERATION_TYPE_EXEMPT      OperationType = 1
	OperationType_OPERATION_TYPE_READ        OperationType = 2
	OperationType_OPERATION_TYPE_WRITE       OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_UNSPECIFIED",
		1: "OPERATION_TYPE_EXEMPT",
		2: "OPERATION_TYPE_READ",
		3: "OPERATION_TYPE_WRITE",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_UNSPECIFIED": 0,
		"OPERATION_TYPE_EXEMPT":      1,
		"OPERATION_TYPE_READ":        2,
		"OPERATION_TYPE_WRITE":       3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_ratelimit_ratelimit_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_annotations_ratelimit_ratelimit_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_annotations_ratelimit_ratelimit_proto_rawDescGZIP(), []int{0}
}

type OperationCategory int32

const (
	OperationCategory_OPERATION_CATEGORY_UNSPECIFIED      OperationCategory = 0
	OperationCategory_OPERATION_CATEGORY_ACL              OperationCategory = 1
	OperationCategory_OPERATION_CATEGORY_PEER_STREAM      OperationCategory = 2
	OperationCategory_OPERATION_CATEGORY_CONNECT_CA       OperationCategory = 3
	OperationCategory_OPERATION_CATEGORY_PARTITION        OperationCategory = 4
	OperationCategory_OPERATION_CATEGORY_PEERING          OperationCategory = 5
	OperationCategory_OPERATION_CATEGORY_SERVER_DISCOVERY OperationCategory = 6
	OperationCategory_OPERATION_CATEGORY_DATAPLANE        OperationCategory = 7
	OperationCategory_OPERATION_CATEGORY_DNS              OperationCategory = 8
	OperationCategory_OPERATION_CATEGORY_SUBSCRIBE        OperationCategory = 9
	OperationCategory_OPERATION_CATEGORY_OPERATOR         OperationCategory = 10
	OperationCategory_OPERATION_CATEGORY_RESOURCE         OperationCategory = 11
	OperationCategory_OPERATION_CATEGORY_CONFIGENTRY      OperationCategory = 12
)

// Enum value maps for OperationCategory.
var (
	OperationCategory_name = map[int32]string{
		0:  "OPERATION_CATEGORY_UNSPECIFIED",
		1:  "OPERATION_CATEGORY_ACL",
		2:  "OPERATION_CATEGORY_PEER_STREAM",
		3:  "OPERATION_CATEGORY_CONNECT_CA",
		4:  "OPERATION_CATEGORY_PARTITION",
		5:  "OPERATION_CATEGORY_PEERING",
		6:  "OPERATION_CATEGORY_SERVER_DISCOVERY",
		7:  "OPERATION_CATEGORY_DATAPLANE",
		8:  "OPERATION_CATEGORY_DNS",
		9:  "OPERATION_CATEGORY_SUBSCRIBE",
		10: "OPERATION_CATEGORY_OPERATOR",
		11: "OPERATION_CATEGORY_RESOURCE",
		12: "OPERATION_CATEGORY_CONFIGENTRY",
	}
	OperationCategory_value = map[string]int32{
		"OPERATION_CATEGORY_UNSPECIFIED":      0,
		"OPERATION_CATEGORY_ACL":              1,
		"OPERATION_CATEGORY_PEER_STREAM":      2,
		"OPERATION_CATEGORY_CONNECT_CA":       3,
		"OPERATION_CATEGORY_PARTITION":        4,
		"OPERATION_CATEGORY_PEERING":          5,
		"OPERATION_CATEGORY_SERVER_DISCOVERY": 6,
		"OPERATION_CATEGORY_DATAPLANE":        7,
		"OPERATION_CATEGORY_DNS":              8,
		"OPERATION_CATEGORY_SUBSCRIBE":        9,
		"OPERATION_CATEGORY_OPERATOR":         10,
		"OPERATION_CATEGORY_RESOURCE":         11,
		"OPERATION_CATEGORY_CONFIGENTRY":      12,
	}
)

func (x OperationCategory) Enum() *OperationCategory {
	p := new(OperationCategory)
	*p = x
	return p
}

func (x OperationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_ratelimit_ratelimit_proto_enumTypes[1].Descriptor()
}

func (OperationCategory) Type() protoreflect.EnumType {
	return &file_annotations_ratelimit_ratelimit_proto_enumTypes[1]
}

func (x OperationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationCategory.Descriptor instead.
func (OperationCategory) EnumDescriptor() ([]byte, []int) {
	return file_annotations_ratelimit_ratelimit_proto_rawDescGZIP(), []int{1}
}

// Spec describes the kind of rate limit that will be applied to this RPC.
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType     OperationType     `protobuf:"varint,1,opt,name=operation_type,json=operationType,proto3,enum=hashicorp.consul.internal.ratelimit.OperationType" json:"operation_type,omitempty"`
	OperationCategory OperationCategory `protobuf:"varint,2,opt,name=operation_category,json=operationCategory,proto3,enum=hashicorp.consul.internal.ratelimit.OperationCategory" json:"operation_category,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_annotations_ratelimit_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_annotations_ratelimit_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_annotations_ratelimit_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *Spec) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_OPERATION_TYPE_UNSPECIFIED
}

func (x *Spec) GetOperationCategory() OperationCategory {
	if x != nil {
		return x.OperationCategory
	}
	return OperationCategory_OPERATION_CATEGORY_UNSPECIFIED
}

var file_annotations_ratelimit_ratelimit_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Spec)(nil),
		Field:         8300,
		Name:          "hashicorp.consul.internal.ratelimit.spec",
		Tag:           "bytes,8300,opt,name=spec",
		Filename:      "annotations/ratelimit/ratelimit.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional hashicorp.consul.internal.ratelimit.Spec spec = 8300;
	E_Spec = &file_annotations_ratelimit_ratelimit_proto_extTypes[0]
)

var File_annotations_ratelimit_ratelimit_proto protoreflect.FileDescriptor

var file_annotations_ratelimit_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8,
	0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x59, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x32, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x65, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2a, 0x7d, 0x0a, 0x0d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x45,
	0x4d, 0x50, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x2a, 0xcb, 0x03, 0x0a, 0x11, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x22,
	0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x43, 0x4c, 0x10, 0x01, 0x12, 0x22,
	0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x5f, 0x43, 0x41, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x45,
	0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x27, 0x0a, 0x23, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x06,
	0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x50, 0x4c, 0x41, 0x4e, 0x45,
	0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x4e, 0x53, 0x10, 0x08, 0x12, 0x20,
	0x0a, 0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x09,
	0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10,
	0x0a, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x10, 0x0b, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x45,
	0x4e, 0x54, 0x52, 0x59, 0x10, 0x0c, 0x3a, 0x5e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec,
	0x40, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0xa9, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49, 0x52, 0xaa, 0x02, 0x23, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0xca, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x52, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0xe2, 0x02, 0x2f, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5c, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_annotations_ratelimit_ratelimit_proto_rawDescOnce sync.Once
	file_annotations_ratelimit_ratelimit_proto_rawDescData = file_annotations_ratelimit_ratelimit_proto_rawDesc
)

func file_annotations_ratelimit_ratelimit_proto_rawDescGZIP() []byte {
	file_annotations_ratelimit_ratelimit_proto_rawDescOnce.Do(func() {
		file_annotations_ratelimit_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_annotations_ratelimit_ratelimit_proto_rawDescData)
	})
	return file_annotations_ratelimit_ratelimit_proto_rawDescData
}

var file_annotations_ratelimit_ratelimit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_annotations_ratelimit_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_annotations_ratelimit_ratelimit_proto_goTypes = []interface{}{
	(OperationType)(0),                 // 0: hashicorp.consul.internal.ratelimit.OperationType
	(OperationCategory)(0),             // 1: hashicorp.consul.internal.ratelimit.OperationCategory
	(*Spec)(nil),                       // 2: hashicorp.consul.internal.ratelimit.Spec
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_annotations_ratelimit_ratelimit_proto_depIdxs = []int32{
	0, // 0: hashicorp.consul.internal.ratelimit.Spec.operation_type:type_name -> hashicorp.consul.internal.ratelimit.OperationType
	1, // 1: hashicorp.consul.internal.ratelimit.Spec.operation_category:type_name -> hashicorp.consul.internal.ratelimit.OperationCategory
	3, // 2: hashicorp.consul.internal.ratelimit.spec:extendee -> google.protobuf.MethodOptions
	2, // 3: hashicorp.consul.internal.ratelimit.spec:type_name -> hashicorp.consul.internal.ratelimit.Spec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_annotations_ratelimit_ratelimit_proto_init() }
func file_annotations_ratelimit_ratelimit_proto_init() {
	if File_annotations_ratelimit_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_annotations_ratelimit_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_annotations_ratelimit_ratelimit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_annotations_ratelimit_ratelimit_proto_goTypes,
		DependencyIndexes: file_annotations_ratelimit_ratelimit_proto_depIdxs,
		EnumInfos:         file_annotations_ratelimit_ratelimit_proto_enumTypes,
		MessageInfos:      file_annotations_ratelimit_ratelimit_proto_msgTypes,
		ExtensionInfos:    file_annotations_ratelimit_ratelimit_proto_extTypes,
	}.Build()
	File_annotations_ratelimit_ratelimit_proto = out.File
	file_annotations_ratelimit_ratelimit_proto_rawDesc = nil
	file_annotations_ratelimit_ratelimit_proto_goTypes = nil
	file_annotations_ratelimit_ratelimit_proto_depIdxs = nil
}
