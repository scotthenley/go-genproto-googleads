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
// source: google/ads/googleads/v7/services/ad_group_asset_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v7/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Request message for [AdGroupAssetService.GetAdGroupAsset][google.ads.googleads.v7.services.AdGroupAssetService.GetAdGroupAsset].
type GetAdGroupAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the ad group asset to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetAdGroupAssetRequest) Reset() {
	*x = GetAdGroupAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdGroupAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdGroupAssetRequest) ProtoMessage() {}

func (x *GetAdGroupAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdGroupAssetRequest.ProtoReflect.Descriptor instead.
func (*GetAdGroupAssetRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdGroupAssetRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [AdGroupAssetService.MutateAdGroupAssets][google.ads.googleads.v7.services.AdGroupAssetService.MutateAdGroupAssets].
type MutateAdGroupAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose ad group assets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad group assets.
	Operations []*AdGroupAssetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAdGroupAssetsRequest) Reset() {
	*x = MutateAdGroupAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetsRequest) ProtoMessage() {}

func (x *MutateAdGroupAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateAdGroupAssetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupAssetsRequest) GetOperations() []*AdGroupAssetOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupAssetsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupAssetsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group asset.
type AdGroupAssetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AdGroupAssetOperation_Create
	//	*AdGroupAssetOperation_Update
	//	*AdGroupAssetOperation_Remove
	Operation isAdGroupAssetOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AdGroupAssetOperation) Reset() {
	*x = AdGroupAssetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAssetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAssetOperation) ProtoMessage() {}

func (x *AdGroupAssetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAssetOperation.ProtoReflect.Descriptor instead.
func (*AdGroupAssetOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP(), []int{2}
}

func (x *AdGroupAssetOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AdGroupAssetOperation) GetOperation() isAdGroupAssetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AdGroupAssetOperation) GetCreate() *resources.AdGroupAsset {
	if x, ok := x.GetOperation().(*AdGroupAssetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AdGroupAssetOperation) GetUpdate() *resources.AdGroupAsset {
	if x, ok := x.GetOperation().(*AdGroupAssetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *AdGroupAssetOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AdGroupAssetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAdGroupAssetOperation_Operation interface {
	isAdGroupAssetOperation_Operation()
}

type AdGroupAssetOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group
	// asset.
	Create *resources.AdGroupAsset `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAssetOperation_Update struct {
	// Update operation: The ad group asset is expected to have a valid resource
	// name.
	Update *resources.AdGroupAsset `protobuf:"bytes,3,opt,name=update,proto3,oneof"`
}

type AdGroupAssetOperation_Remove struct {
	// Remove operation: A resource name for the removed ad group asset is
	// expected, in this format:
	//
	// `customers/{customer_id}/adGroupAssets/{ad_group_id}~{asset_id}~{field_type}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAssetOperation_Create) isAdGroupAssetOperation_Operation() {}

func (*AdGroupAssetOperation_Update) isAdGroupAssetOperation_Operation() {}

func (*AdGroupAssetOperation_Remove) isAdGroupAssetOperation_Operation() {}

// Response message for an ad group asset mutate.
type MutateAdGroupAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateAdGroupAssetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAdGroupAssetsResponse) Reset() {
	*x = MutateAdGroupAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetsResponse) ProtoMessage() {}

func (x *MutateAdGroupAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupAssetsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupAssetsResponse) GetResults() []*MutateAdGroupAssetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the ad group asset mutate.
type MutateAdGroupAssetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAdGroupAssetResult) Reset() {
	*x = MutateAdGroupAssetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAssetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetResult) ProtoMessage() {}

func (x *MutateAdGroupAssetResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateAdGroupAssetResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v7_services_ad_group_asset_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27,
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x91, 0x02, 0x0a, 0x15, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x49, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x1b, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x54, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x18, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x91, 0x04, 0x0a, 0x13, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x76, 0x37, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xea, 0x01, 0x0a, 0x13, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x32, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0xff, 0x01,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x18, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x37, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescData = file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDesc
)

func file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDescData
}

var file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v7_services_ad_group_asset_service_proto_goTypes = []interface{}{
	(*GetAdGroupAssetRequest)(nil),      // 0: google.ads.googleads.v7.services.GetAdGroupAssetRequest
	(*MutateAdGroupAssetsRequest)(nil),  // 1: google.ads.googleads.v7.services.MutateAdGroupAssetsRequest
	(*AdGroupAssetOperation)(nil),       // 2: google.ads.googleads.v7.services.AdGroupAssetOperation
	(*MutateAdGroupAssetsResponse)(nil), // 3: google.ads.googleads.v7.services.MutateAdGroupAssetsResponse
	(*MutateAdGroupAssetResult)(nil),    // 4: google.ads.googleads.v7.services.MutateAdGroupAssetResult
	(*fieldmaskpb.FieldMask)(nil),       // 5: google.protobuf.FieldMask
	(*resources.AdGroupAsset)(nil),      // 6: google.ads.googleads.v7.resources.AdGroupAsset
	(*status.Status)(nil),               // 7: google.rpc.Status
}
var file_google_ads_googleads_v7_services_ad_group_asset_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v7.services.MutateAdGroupAssetsRequest.operations:type_name -> google.ads.googleads.v7.services.AdGroupAssetOperation
	5, // 1: google.ads.googleads.v7.services.AdGroupAssetOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v7.services.AdGroupAssetOperation.create:type_name -> google.ads.googleads.v7.resources.AdGroupAsset
	6, // 3: google.ads.googleads.v7.services.AdGroupAssetOperation.update:type_name -> google.ads.googleads.v7.resources.AdGroupAsset
	7, // 4: google.ads.googleads.v7.services.MutateAdGroupAssetsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 5: google.ads.googleads.v7.services.MutateAdGroupAssetsResponse.results:type_name -> google.ads.googleads.v7.services.MutateAdGroupAssetResult
	0, // 6: google.ads.googleads.v7.services.AdGroupAssetService.GetAdGroupAsset:input_type -> google.ads.googleads.v7.services.GetAdGroupAssetRequest
	1, // 7: google.ads.googleads.v7.services.AdGroupAssetService.MutateAdGroupAssets:input_type -> google.ads.googleads.v7.services.MutateAdGroupAssetsRequest
	6, // 8: google.ads.googleads.v7.services.AdGroupAssetService.GetAdGroupAsset:output_type -> google.ads.googleads.v7.resources.AdGroupAsset
	3, // 9: google.ads.googleads.v7.services.AdGroupAssetService.MutateAdGroupAssets:output_type -> google.ads.googleads.v7.services.MutateAdGroupAssetsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_services_ad_group_asset_service_proto_init() }
func file_google_ads_googleads_v7_services_ad_group_asset_service_proto_init() {
	if File_google_ads_googleads_v7_services_ad_group_asset_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdGroupAssetRequest); i {
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
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAssetsRequest); i {
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
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupAssetOperation); i {
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
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAssetsResponse); i {
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
		file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAssetResult); i {
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
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AdGroupAssetOperation_Create)(nil),
		(*AdGroupAssetOperation_Update)(nil),
		(*AdGroupAssetOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v7_services_ad_group_asset_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_services_ad_group_asset_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v7_services_ad_group_asset_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_services_ad_group_asset_service_proto = out.File
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_rawDesc = nil
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_goTypes = nil
	file_google_ads_googleads_v7_services_ad_group_asset_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAssetServiceClient is the client API for AdGroupAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAssetServiceClient interface {
	// Returns the requested ad group asset in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAsset(ctx context.Context, in *GetAdGroupAssetRequest, opts ...grpc.CallOption) (*resources.AdGroupAsset, error)
	// Creates, updates, or removes ad group assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAssets(ctx context.Context, in *MutateAdGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetsResponse, error)
}

type adGroupAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAssetServiceClient(cc grpc.ClientConnInterface) AdGroupAssetServiceClient {
	return &adGroupAssetServiceClient{cc}
}

func (c *adGroupAssetServiceClient) GetAdGroupAsset(ctx context.Context, in *GetAdGroupAssetRequest, opts ...grpc.CallOption) (*resources.AdGroupAsset, error) {
	out := new(resources.AdGroupAsset)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.AdGroupAssetService/GetAdGroupAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAssetServiceClient) MutateAdGroupAssets(ctx context.Context, in *MutateAdGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetsResponse, error) {
	out := new(MutateAdGroupAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.AdGroupAssetService/MutateAdGroupAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAssetServiceServer is the server API for AdGroupAssetService service.
type AdGroupAssetServiceServer interface {
	// Returns the requested ad group asset in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAsset(context.Context, *GetAdGroupAssetRequest) (*resources.AdGroupAsset, error)
	// Creates, updates, or removes ad group assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAssets(context.Context, *MutateAdGroupAssetsRequest) (*MutateAdGroupAssetsResponse, error)
}

// UnimplementedAdGroupAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAssetServiceServer struct {
}

func (*UnimplementedAdGroupAssetServiceServer) GetAdGroupAsset(context.Context, *GetAdGroupAssetRequest) (*resources.AdGroupAsset, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAsset not implemented")
}
func (*UnimplementedAdGroupAssetServiceServer) MutateAdGroupAssets(context.Context, *MutateAdGroupAssetsRequest) (*MutateAdGroupAssetsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAssets not implemented")
}

func RegisterAdGroupAssetServiceServer(s *grpc.Server, srv AdGroupAssetServiceServer) {
	s.RegisterService(&_AdGroupAssetService_serviceDesc, srv)
}

func _AdGroupAssetService_GetAdGroupAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAssetServiceServer).GetAdGroupAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.AdGroupAssetService/GetAdGroupAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAssetServiceServer).GetAdGroupAsset(ctx, req.(*GetAdGroupAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAssetService_MutateAdGroupAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAssetServiceServer).MutateAdGroupAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.AdGroupAssetService/MutateAdGroupAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAssetServiceServer).MutateAdGroupAssets(ctx, req.(*MutateAdGroupAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.AdGroupAssetService",
	HandlerType: (*AdGroupAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAsset",
			Handler:    _AdGroupAssetService_GetAdGroupAsset_Handler,
		},
		{
			MethodName: "MutateAdGroupAssets",
			Handler:    _AdGroupAssetService_MutateAdGroupAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/ad_group_asset_service.proto",
}
