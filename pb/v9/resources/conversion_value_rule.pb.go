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
// 	protoc        v3.18.1
// source: google/ads/googleads/v9/resources/conversion_value_rule.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	enums "github.com/scotthenley/go-genproto-googleads/pb/v9/enums"
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

// A conversion value rule
type ConversionValueRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the conversion value rule.
	// Conversion value rule resource names have the form:
	//
	// `customers/{customer_id}/conversionValueRules/{conversion_value_rule_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the conversion value rule.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Action applied when the rule is triggered.
	Action *ConversionValueRule_ValueRuleAction `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	// Condition for Geo location that must be satisfied for the value rule to
	// apply.
	GeoLocationCondition *ConversionValueRule_ValueRuleGeoLocationCondition `protobuf:"bytes,4,opt,name=geo_location_condition,json=geoLocationCondition,proto3" json:"geo_location_condition,omitempty"`
	// Condition for device type that must be satisfied for the value rule to
	// apply.
	DeviceCondition *ConversionValueRule_ValueRuleDeviceCondition `protobuf:"bytes,5,opt,name=device_condition,json=deviceCondition,proto3" json:"device_condition,omitempty"`
	// Condition for audience that must be satisfied for the value rule to apply.
	AudienceCondition *ConversionValueRule_ValueRuleAudienceCondition `protobuf:"bytes,6,opt,name=audience_condition,json=audienceCondition,proto3" json:"audience_condition,omitempty"`
	// Output only. The resource name of the conversion value rule's owner customer.
	// When the value rule is inherited from a manager
	// customer, owner_customer will be the resource name of the manager whereas
	// the customer in the resource_name will be of the requesting serving
	// customer.
	// ** Read-only **
	OwnerCustomer string `protobuf:"bytes,7,opt,name=owner_customer,json=ownerCustomer,proto3" json:"owner_customer,omitempty"`
	// The status of the conversion value rule.
	Status enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus `protobuf:"varint,8,opt,name=status,proto3,enum=google.ads.googleads.v9.enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus" json:"status,omitempty"`
}

func (x *ConversionValueRule) Reset() {
	*x = ConversionValueRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule) ProtoMessage() {}

func (x *ConversionValueRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule.ProtoReflect.Descriptor instead.
func (*ConversionValueRule) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0}
}

func (x *ConversionValueRule) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ConversionValueRule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConversionValueRule) GetAction() *ConversionValueRule_ValueRuleAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *ConversionValueRule) GetGeoLocationCondition() *ConversionValueRule_ValueRuleGeoLocationCondition {
	if x != nil {
		return x.GeoLocationCondition
	}
	return nil
}

func (x *ConversionValueRule) GetDeviceCondition() *ConversionValueRule_ValueRuleDeviceCondition {
	if x != nil {
		return x.DeviceCondition
	}
	return nil
}

func (x *ConversionValueRule) GetAudienceCondition() *ConversionValueRule_ValueRuleAudienceCondition {
	if x != nil {
		return x.AudienceCondition
	}
	return nil
}

func (x *ConversionValueRule) GetOwnerCustomer() string {
	if x != nil {
		return x.OwnerCustomer
	}
	return ""
}

func (x *ConversionValueRule) GetStatus() enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus {
	if x != nil {
		return x.Status
	}
	return enums.ConversionValueRuleStatusEnum_UNSPECIFIED
}

// Action applied when rule is applied.
type ConversionValueRule_ValueRuleAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies applied operation.
	Operation enums.ValueRuleOperationEnum_ValueRuleOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=google.ads.googleads.v9.enums.ValueRuleOperationEnum_ValueRuleOperation" json:"operation,omitempty"`
	// Specifies applied value.
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConversionValueRule_ValueRuleAction) Reset() {
	*x = ConversionValueRule_ValueRuleAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRule_ValueRuleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleAction) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleAction) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleAction.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleAction) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConversionValueRule_ValueRuleAction) GetOperation() enums.ValueRuleOperationEnum_ValueRuleOperation {
	if x != nil {
		return x.Operation
	}
	return enums.ValueRuleOperationEnum_UNSPECIFIED
}

func (x *ConversionValueRule_ValueRuleAction) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Condition on Geo dimension.
type ConversionValueRule_ValueRuleGeoLocationCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Geo locations that advertisers want to exclude.
	ExcludedGeoTargetConstants []string `protobuf:"bytes,1,rep,name=excluded_geo_target_constants,json=excludedGeoTargetConstants,proto3" json:"excluded_geo_target_constants,omitempty"`
	// Excluded Geo location match type.
	ExcludedGeoMatchType enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType `protobuf:"varint,2,opt,name=excluded_geo_match_type,json=excludedGeoMatchType,proto3,enum=google.ads.googleads.v9.enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType" json:"excluded_geo_match_type,omitempty"`
	// Geo locations that advertisers want to include.
	GeoTargetConstants []string `protobuf:"bytes,3,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Included Geo location match type.
	GeoMatchType enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType `protobuf:"varint,4,opt,name=geo_match_type,json=geoMatchType,proto3,enum=google.ads.googleads.v9.enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType" json:"geo_match_type,omitempty"`
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) Reset() {
	*x = ConversionValueRule_ValueRuleGeoLocationCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleGeoLocationCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleGeoLocationCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleGeoLocationCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetExcludedGeoTargetConstants() []string {
	if x != nil {
		return x.ExcludedGeoTargetConstants
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetExcludedGeoMatchType() enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType {
	if x != nil {
		return x.ExcludedGeoMatchType
	}
	return enums.ValueRuleGeoLocationMatchTypeEnum_UNSPECIFIED
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetGeoTargetConstants() []string {
	if x != nil {
		return x.GeoTargetConstants
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetGeoMatchType() enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType {
	if x != nil {
		return x.GeoMatchType
	}
	return enums.ValueRuleGeoLocationMatchTypeEnum_UNSPECIFIED
}

// Condition on Device dimension.
type ConversionValueRule_ValueRuleDeviceCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value for device type condition.
	DeviceTypes []enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType `protobuf:"varint,1,rep,packed,name=device_types,json=deviceTypes,proto3,enum=google.ads.googleads.v9.enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType" json:"device_types,omitempty"`
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) Reset() {
	*x = ConversionValueRule_ValueRuleDeviceCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleDeviceCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleDeviceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleDeviceCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleDeviceCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) GetDeviceTypes() []enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType {
	if x != nil {
		return x.DeviceTypes
	}
	return nil
}

// Condition on Audience dimension.
type ConversionValueRule_ValueRuleAudienceCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User Lists.
	UserLists []string `protobuf:"bytes,1,rep,name=user_lists,json=userLists,proto3" json:"user_lists,omitempty"`
	// User Interests.
	UserInterests []string `protobuf:"bytes,2,rep,name=user_interests,json=userInterests,proto3" json:"user_interests,omitempty"`
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) Reset() {
	*x = ConversionValueRule_ValueRuleAudienceCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleAudienceCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleAudienceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleAudienceCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleAudienceCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) GetUserLists() []string {
	if x != nil {
		return x.UserLists
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) GetUserInterests() []string {
	if x != nil {
		return x.UserInterests
	}
	return nil
}

var File_google_ads_googleads_v9_resources_conversion_value_rule_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x0f, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x5e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8a, 0x01, 0x0a, 0x16, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x67, 0x65, 0x6f,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x7a, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x80, 0x01,
	0x0a, 0x12, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x50, 0x0a, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23,
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x6e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x56, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x8f, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x95, 0x04, 0x0a, 0x1d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x1d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x64, 0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2f, 0xfa,
	0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x6f,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x1a,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x95, 0x01, 0x0a, 0x17, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x47, 0x65, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x52, 0x12, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x67, 0x65, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x89, 0x01, 0x0a,
	0x18, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x0c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0xb6, 0x01, 0x0a, 0x1a, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23,
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x51,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2a, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x73, 0x3a, 0x7a, 0xea, 0x41, 0x77, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x47, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x85, 0x02,
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x18, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescData = file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDesc
)

func file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDescData
}

var file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v9_resources_conversion_value_rule_proto_goTypes = []interface{}{
	(*ConversionValueRule)(nil),                                                // 0: google.ads.googleads.v9.resources.ConversionValueRule
	(*ConversionValueRule_ValueRuleAction)(nil),                                // 1: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleAction
	(*ConversionValueRule_ValueRuleGeoLocationCondition)(nil),                  // 2: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleGeoLocationCondition
	(*ConversionValueRule_ValueRuleDeviceCondition)(nil),                       // 3: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleDeviceCondition
	(*ConversionValueRule_ValueRuleAudienceCondition)(nil),                     // 4: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleAudienceCondition
	(enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus)(0),         // 5: google.ads.googleads.v9.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatus
	(enums.ValueRuleOperationEnum_ValueRuleOperation)(0),                       // 6: google.ads.googleads.v9.enums.ValueRuleOperationEnum.ValueRuleOperation
	(enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType)(0), // 7: google.ads.googleads.v9.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	(enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType)(0),                     // 8: google.ads.googleads.v9.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceType
}
var file_google_ads_googleads_v9_resources_conversion_value_rule_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v9.resources.ConversionValueRule.action:type_name -> google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleAction
	2, // 1: google.ads.googleads.v9.resources.ConversionValueRule.geo_location_condition:type_name -> google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleGeoLocationCondition
	3, // 2: google.ads.googleads.v9.resources.ConversionValueRule.device_condition:type_name -> google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleDeviceCondition
	4, // 3: google.ads.googleads.v9.resources.ConversionValueRule.audience_condition:type_name -> google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleAudienceCondition
	5, // 4: google.ads.googleads.v9.resources.ConversionValueRule.status:type_name -> google.ads.googleads.v9.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatus
	6, // 5: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleAction.operation:type_name -> google.ads.googleads.v9.enums.ValueRuleOperationEnum.ValueRuleOperation
	7, // 6: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleGeoLocationCondition.excluded_geo_match_type:type_name -> google.ads.googleads.v9.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	7, // 7: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleGeoLocationCondition.geo_match_type:type_name -> google.ads.googleads.v9.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	8, // 8: google.ads.googleads.v9.resources.ConversionValueRule.ValueRuleDeviceCondition.device_types:type_name -> google.ads.googleads.v9.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_resources_conversion_value_rule_proto_init() }
func file_google_ads_googleads_v9_resources_conversion_value_rule_proto_init() {
	if File_google_ads_googleads_v9_resources_conversion_value_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRule); i {
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
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRule_ValueRuleAction); i {
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
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRule_ValueRuleGeoLocationCondition); i {
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
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRule_ValueRuleDeviceCondition); i {
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
		file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRule_ValueRuleAudienceCondition); i {
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
			RawDescriptor: file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_resources_conversion_value_rule_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_resources_conversion_value_rule_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v9_resources_conversion_value_rule_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_resources_conversion_value_rule_proto = out.File
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_rawDesc = nil
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_goTypes = nil
	file_google_ads_googleads_v9_resources_conversion_value_rule_proto_depIdxs = nil
}
