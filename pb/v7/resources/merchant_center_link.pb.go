// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/ads/googleads/v7/resources/merchant_center_link.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	enums "github.com/dictav/go-genproto-googleads/pb/v7/enums"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A data sharing connection, proposed or in use,
// between a Google Ads Customer and a Merchant Center account.
type MerchantCenterLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the merchant center link.
	// Merchant center link resource names have the form:
	//
	// `customers/{customer_id}/merchantCenterLinks/{merchant_center_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the Merchant Center account.
	// This field is readonly.
	Id *int64 `protobuf:"varint,6,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The name of the Merchant Center account.
	// This field is readonly.
	MerchantCenterAccountName *string `protobuf:"bytes,7,opt,name=merchant_center_account_name,json=merchantCenterAccountName,proto3,oneof" json:"merchant_center_account_name,omitempty"`
	// The status of the link.
	Status enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v7.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus" json:"status,omitempty"`
}

func (x *MerchantCenterLink) Reset() {
	*x = MerchantCenterLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantCenterLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantCenterLink) ProtoMessage() {}

func (x *MerchantCenterLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantCenterLink.ProtoReflect.Descriptor instead.
func (*MerchantCenterLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescGZIP(), []int{0}
}

func (x *MerchantCenterLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MerchantCenterLink) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MerchantCenterLink) GetMerchantCenterAccountName() string {
	if x != nil && x.MerchantCenterAccountName != nil {
		return *x.MerchantCenterAccountName
	}
	return ""
}

func (x *MerchantCenterLink) GetStatus() enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.MerchantCenterLinkStatusEnum_UNSPECIFIED
}

var File_google_ads_googleads_v7_resources_merchant_center_link_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a,
	0x12, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x49, 0x0a, 0x1c, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x01, 0x52, 0x19, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x6c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3a, 0x72, 0xea, 0x41, 0x6f, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x40, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f,
	0x7b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x1f, 0x0a, 0x1d, 0x5f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x84, 0x02, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescData = file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDesc
)

func file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDescData
}

var file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v7_resources_merchant_center_link_proto_goTypes = []interface{}{
	(*MerchantCenterLink)(nil),                                       // 0: google.ads.googleads.v7.resources.MerchantCenterLink
	(enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus)(0), // 1: google.ads.googleads.v7.enums.MerchantCenterLinkStatusEnum.MerchantCenterLinkStatus
}
var file_google_ads_googleads_v7_resources_merchant_center_link_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v7.resources.MerchantCenterLink.status:type_name -> google.ads.googleads.v7.enums.MerchantCenterLinkStatusEnum.MerchantCenterLinkStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_resources_merchant_center_link_proto_init() }
func file_google_ads_googleads_v7_resources_merchant_center_link_proto_init() {
	if File_google_ads_googleads_v7_resources_merchant_center_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantCenterLink); i {
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
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v7_resources_merchant_center_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_resources_merchant_center_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v7_resources_merchant_center_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_resources_merchant_center_link_proto = out.File
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_rawDesc = nil
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_goTypes = nil
	file_google_ads_googleads_v7_resources_merchant_center_link_proto_depIdxs = nil
}
