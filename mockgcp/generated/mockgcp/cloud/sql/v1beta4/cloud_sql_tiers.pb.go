// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/cloud/sql/v1beta4/cloud_sql_tiers.proto

package sqlpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SqlTiersListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project ID of the project for which to list tiers.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *SqlTiersListRequest) Reset() {
	*x = SqlTiersListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlTiersListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlTiersListRequest) ProtoMessage() {}

func (x *SqlTiersListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlTiersListRequest.ProtoReflect.Descriptor instead.
func (*SqlTiersListRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescGZIP(), []int{0}
}

func (x *SqlTiersListRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

// Tiers list response.
type TiersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is always `sql#tiersList`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// List of tiers.
	Items []*Tier `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TiersListResponse) Reset() {
	*x = TiersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiersListResponse) ProtoMessage() {}

func (x *TiersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiersListResponse.ProtoReflect.Descriptor instead.
func (*TiersListResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescGZIP(), []int{1}
}

func (x *TiersListResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *TiersListResponse) GetItems() []*Tier {
	if x != nil {
		return x.Items
	}
	return nil
}

// A Google Cloud SQL service tier resource.
type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier for the machine type, for example, `db-custom-1-3840`. For
	// related information, see [Pricing](/sql/pricing).
	Tier string `protobuf:"bytes,1,opt,name=tier,proto3" json:"tier,omitempty"`
	// The maximum RAM usage of this tier in bytes.
	RAM int64 `protobuf:"varint,2,opt,name=RAM,proto3" json:"RAM,omitempty"`
	// This is always `sql#tier`.
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// The maximum disk size of this tier in bytes.
	Disk_Quota int64 `protobuf:"varint,4,opt,name=Disk_Quota,json=DiskQuota,proto3" json:"Disk_Quota,omitempty"`
	// The applicable regions for this tier.
	Region []string `protobuf:"bytes,5,rep,name=region,proto3" json:"region,omitempty"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescGZIP(), []int{2}
}

func (x *Tier) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *Tier) GetRAM() int64 {
	if x != nil {
		return x.RAM
	}
	return 0
}

func (x *Tier) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Tier) GetDisk_Quota() int64 {
	if x != nil {
		return x.Disk_Quota
	}
	return 0
}

func (x *Tier) GetRegion() []string {
	if x != nil {
		return x.Region
	}
	return nil
}

var File_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x73, 0x71, 0x6c, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x71, 0x6c, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x5e, 0x0a, 0x11, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x35, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x77, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x52, 0x41, 0x4d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x52,
	0x41, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x6b, 0x5f, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x32, 0xa5, 0x02,
	0x0a, 0x0f, 0x53, 0x71, 0x6c, 0x54, 0x69, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x93, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27,
	0x12, 0x25, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x7c, 0xca, 0x41, 0x17, 0x73, 0x71, 0x6c, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x65, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x34, 0x42, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c,
	0x54, 0x69, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x34,
	0x2f, 0x73, 0x71, 0x6c, 0x70, 0x62, 0x3b, 0x73, 0x71, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescData = file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDesc
)

func file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescData)
	})
	return file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDescData
}

var file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_goTypes = []interface{}{
	(*SqlTiersListRequest)(nil), // 0: mockgcp.cloud.sql.v1beta4.SqlTiersListRequest
	(*TiersListResponse)(nil),   // 1: mockgcp.cloud.sql.v1beta4.TiersListResponse
	(*Tier)(nil),                // 2: mockgcp.cloud.sql.v1beta4.Tier
}
var file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_depIdxs = []int32{
	2, // 0: mockgcp.cloud.sql.v1beta4.TiersListResponse.items:type_name -> mockgcp.cloud.sql.v1beta4.Tier
	0, // 1: mockgcp.cloud.sql.v1beta4.SqlTiersService.List:input_type -> mockgcp.cloud.sql.v1beta4.SqlTiersListRequest
	1, // 2: mockgcp.cloud.sql.v1beta4.SqlTiersService.List:output_type -> mockgcp.cloud.sql.v1beta4.TiersListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_init() }
func file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_init() {
	if File_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlTiersListRequest); i {
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
		file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TiersListResponse); i {
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
		file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
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
			RawDescriptor: file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto = out.File
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_rawDesc = nil
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_goTypes = nil
	file_mockgcp_cloud_sql_v1beta4_cloud_sql_tiers_proto_depIdxs = nil
}
