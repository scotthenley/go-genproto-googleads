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
// source: google/ads/googleads/v8/resources/customer.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v8/enums"
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

// A customer.
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the customer.
	Id *int64 `protobuf:"varint,19,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *string `protobuf:"bytes,20,opt,name=descriptive_name,json=descriptiveName,proto3,oneof" json:"descriptive_name,omitempty"`
	// Immutable. The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *string `protobuf:"bytes,21,opt,name=currency_code,json=currencyCode,proto3,oneof" json:"currency_code,omitempty"`
	// Immutable. The local timezone ID of the customer.
	TimeZone *string `protobuf:"bytes,22,opt,name=time_zone,json=timeZone,proto3,oneof" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *string `protobuf:"bytes,23,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3,oneof" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *string `protobuf:"bytes,24,opt,name=final_url_suffix,json=finalUrlSuffix,proto3,oneof" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *bool `protobuf:"varint,25,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3,oneof" json:"auto_tagging_enabled,omitempty"`
	// Output only. Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *bool `protobuf:"varint,26,opt,name=has_partners_badge,json=hasPartnersBadge,proto3,oneof" json:"has_partners_badge,omitempty"`
	// Output only. Whether the customer is a manager.
	Manager *bool `protobuf:"varint,27,opt,name=manager,proto3,oneof" json:"manager,omitempty"`
	// Output only. Whether the customer is a test account.
	TestAccount *bool `protobuf:"varint,28,opt,name=test_account,json=testAccount,proto3,oneof" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Output only. Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	// Output only. Remarketing setting for a customer.
	RemarketingSetting *RemarketingSetting `protobuf:"bytes,15,opt,name=remarketing_setting,json=remarketingSetting,proto3" json:"remarketing_setting,omitempty"`
	// Output only. Reasons why the customer is not eligible to use PaymentMode.CONVERSIONS. If
	// the list is empty, the customer is eligible. This field is read-only.
	PayPerConversionEligibilityFailureReasons []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason `protobuf:"varint,16,rep,packed,name=pay_per_conversion_eligibility_failure_reasons,json=payPerConversionEligibilityFailureReasons,proto3,enum=google.ads.googleads.v8.enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason" json:"pay_per_conversion_eligibility_failure_reasons,omitempty"`
	// Output only. Optimization score of the customer.
	//
	// Optimization score is an estimate of how well a customer's campaigns are
	// set to perform. It ranges from 0% (0.0) to 100% (1.0). This field is null
	// for all manager customers, and for unscored non-manager customers.
	//
	// See "About optimization score" at
	// https://support.google.com/google-ads/answer/9061546.
	//
	// This field is read-only.
	OptimizationScore *float64 `protobuf:"fixed64,29,opt,name=optimization_score,json=optimizationScore,proto3,oneof" json:"optimization_score,omitempty"`
	// Output only. Optimization score weight of the customer.
	//
	// Optimization score weight can be used to compare/aggregate optimization
	// scores across multiple non-manager customers. The aggregate optimization
	// score of a manager is computed as the sum over all of their customers of
	// `Customer.optimization_score * Customer.optimization_score_weight`. This
	// field is 0 for all manager customers, and for unscored non-manager
	// customers.
	//
	// This field is read-only.
	OptimizationScoreWeight float64 `protobuf:"fixed64,30,opt,name=optimization_score_weight,json=optimizationScoreWeight,proto3" json:"optimization_score_weight,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Customer) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Customer) GetDescriptiveName() string {
	if x != nil && x.DescriptiveName != nil {
		return *x.DescriptiveName
	}
	return ""
}

func (x *Customer) GetCurrencyCode() string {
	if x != nil && x.CurrencyCode != nil {
		return *x.CurrencyCode
	}
	return ""
}

func (x *Customer) GetTimeZone() string {
	if x != nil && x.TimeZone != nil {
		return *x.TimeZone
	}
	return ""
}

func (x *Customer) GetTrackingUrlTemplate() string {
	if x != nil && x.TrackingUrlTemplate != nil {
		return *x.TrackingUrlTemplate
	}
	return ""
}

func (x *Customer) GetFinalUrlSuffix() string {
	if x != nil && x.FinalUrlSuffix != nil {
		return *x.FinalUrlSuffix
	}
	return ""
}

func (x *Customer) GetAutoTaggingEnabled() bool {
	if x != nil && x.AutoTaggingEnabled != nil {
		return *x.AutoTaggingEnabled
	}
	return false
}

func (x *Customer) GetHasPartnersBadge() bool {
	if x != nil && x.HasPartnersBadge != nil {
		return *x.HasPartnersBadge
	}
	return false
}

func (x *Customer) GetManager() bool {
	if x != nil && x.Manager != nil {
		return *x.Manager
	}
	return false
}

func (x *Customer) GetTestAccount() bool {
	if x != nil && x.TestAccount != nil {
		return *x.TestAccount
	}
	return false
}

func (x *Customer) GetCallReportingSetting() *CallReportingSetting {
	if x != nil {
		return x.CallReportingSetting
	}
	return nil
}

func (x *Customer) GetConversionTrackingSetting() *ConversionTrackingSetting {
	if x != nil {
		return x.ConversionTrackingSetting
	}
	return nil
}

func (x *Customer) GetRemarketingSetting() *RemarketingSetting {
	if x != nil {
		return x.RemarketingSetting
	}
	return nil
}

func (x *Customer) GetPayPerConversionEligibilityFailureReasons() []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason {
	if x != nil {
		return x.PayPerConversionEligibilityFailureReasons
	}
	return nil
}

func (x *Customer) GetOptimizationScore() float64 {
	if x != nil && x.OptimizationScore != nil {
		return *x.OptimizationScore
	}
	return 0
}

func (x *Customer) GetOptimizationScoreWeight() float64 {
	if x != nil {
		return x.OptimizationScoreWeight
	}
	return 0
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *bool `protobuf:"varint,10,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3,oneof" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *bool `protobuf:"varint,11,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3,oneof" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *string `protobuf:"bytes,12,opt,name=call_conversion_action,json=callConversionAction,proto3,oneof" json:"call_conversion_action,omitempty"`
}

func (x *CallReportingSetting) Reset() {
	*x = CallReportingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallReportingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallReportingSetting) ProtoMessage() {}

func (x *CallReportingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallReportingSetting.ProtoReflect.Descriptor instead.
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CallReportingSetting) GetCallReportingEnabled() bool {
	if x != nil && x.CallReportingEnabled != nil {
		return *x.CallReportingEnabled
	}
	return false
}

func (x *CallReportingSetting) GetCallConversionReportingEnabled() bool {
	if x != nil && x.CallConversionReportingEnabled != nil {
		return *x.CallConversionReportingEnabled
	}
	return false
}

func (x *CallReportingSetting) GetCallConversionAction() string {
	if x != nil && x.CallConversionAction != nil {
		return *x.CallConversionAction
	}
	return ""
}

// A collection of customer-wide settings related to Google Ads Conversion
// Tracking.
type ConversionTrackingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *int64 `protobuf:"varint,3,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3,oneof" json:"conversion_tracking_id,omitempty"`
	// Output only. The conversion tracking id of the customer's manager. This is set when the
	// customer is opted into cross account conversion tracking, and it overrides
	// conversion_tracking_id. This field can only be managed through the Google
	// Ads UI. This field is read-only.
	CrossAccountConversionTrackingId *int64 `protobuf:"varint,4,opt,name=cross_account_conversion_tracking_id,json=crossAccountConversionTrackingId,proto3,oneof" json:"cross_account_conversion_tracking_id,omitempty"`
}

func (x *ConversionTrackingSetting) Reset() {
	*x = ConversionTrackingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionTrackingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionTrackingSetting) ProtoMessage() {}

func (x *ConversionTrackingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionTrackingSetting.ProtoReflect.Descriptor instead.
func (*ConversionTrackingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_customer_proto_rawDescGZIP(), []int{2}
}

func (x *ConversionTrackingSetting) GetConversionTrackingId() int64 {
	if x != nil && x.ConversionTrackingId != nil {
		return *x.ConversionTrackingId
	}
	return 0
}

func (x *ConversionTrackingSetting) GetCrossAccountConversionTrackingId() int64 {
	if x != nil && x.CrossAccountConversionTrackingId != nil {
		return *x.CrossAccountConversionTrackingId
	}
	return 0
}

// Remarketing setting for a customer.
type RemarketingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Google global site tag.
	GoogleGlobalSiteTag *string `protobuf:"bytes,2,opt,name=google_global_site_tag,json=googleGlobalSiteTag,proto3,oneof" json:"google_global_site_tag,omitempty"`
}

func (x *RemarketingSetting) Reset() {
	*x = RemarketingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemarketingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemarketingSetting) ProtoMessage() {}

func (x *RemarketingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemarketingSetting.ProtoReflect.Descriptor instead.
func (*RemarketingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_customer_proto_rawDescGZIP(), []int{3}
}

func (x *RemarketingSetting) GetGoogleGlobalSiteTag() string {
	if x != nil && x.GoogleGlobalSiteTag != nil {
		return *x.GoogleGlobalSiteTag
	}
	return ""
}

var File_google_ads_googleads_v8_resources_customer_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_customer_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x5a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x0b, 0x0a, 0x08,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x02,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x13, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01,
	0x12, 0x35, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06,
	0x52, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x12, 0x68, 0x61, 0x73, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x07, 0x52, 0x10, 0x68, 0x61, 0x73, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x61, 0x64, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x09,
	0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x6d, 0x0a, 0x16, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x14, 0x63, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x81, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0xed, 0x01, 0x0a, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x84, 0x01, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x50, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x29, 0x70, 0x61, 0x79, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x37, 0x0a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x0a, 0x52, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x19, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x3f, 0xea, 0x41, 0x3c, 0x0a,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x5f,
	0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73,
	0x5f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x14, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x16, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x14, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x4e,
	0x0a, 0x21, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x1e, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x69,
	0x0a, 0x16, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e,
	0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02,
	0x52, 0x14, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x42, 0x24, 0x0a, 0x22, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x24, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x20, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x19, 0x0a,
	0x17, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x27, 0x0a, 0x25, 0x5f, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x22, 0x6e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x13,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x88, 0x01, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x67, 0x42, 0xfa, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0d, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_customer_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_customer_proto_rawDescData = file_google_ads_googleads_v8_resources_customer_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_customer_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_customer_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_customer_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_customer_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v8_resources_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),                  // 0: google.ads.googleads.v8.resources.Customer
	(*CallReportingSetting)(nil),      // 1: google.ads.googleads.v8.resources.CallReportingSetting
	(*ConversionTrackingSetting)(nil), // 2: google.ads.googleads.v8.resources.ConversionTrackingSetting
	(*RemarketingSetting)(nil),        // 3: google.ads.googleads.v8.resources.RemarketingSetting
	(enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason)(0), // 4: google.ads.googleads.v8.enums.CustomerPayPerConversionEligibilityFailureReasonEnum.CustomerPayPerConversionEligibilityFailureReason
}
var file_google_ads_googleads_v8_resources_customer_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.Customer.call_reporting_setting:type_name -> google.ads.googleads.v8.resources.CallReportingSetting
	2, // 1: google.ads.googleads.v8.resources.Customer.conversion_tracking_setting:type_name -> google.ads.googleads.v8.resources.ConversionTrackingSetting
	3, // 2: google.ads.googleads.v8.resources.Customer.remarketing_setting:type_name -> google.ads.googleads.v8.resources.RemarketingSetting
	4, // 3: google.ads.googleads.v8.resources.Customer.pay_per_conversion_eligibility_failure_reasons:type_name -> google.ads.googleads.v8.enums.CustomerPayPerConversionEligibilityFailureReasonEnum.CustomerPayPerConversionEligibilityFailureReason
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_customer_proto_init() }
func file_google_ads_googleads_v8_resources_customer_proto_init() {
	if File_google_ads_googleads_v8_resources_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_google_ads_googleads_v8_resources_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallReportingSetting); i {
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
		file_google_ads_googleads_v8_resources_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionTrackingSetting); i {
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
		file_google_ads_googleads_v8_resources_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemarketingSetting); i {
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
	file_google_ads_googleads_v8_resources_customer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_customer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_customer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_customer_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_customer_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_customer_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_customer_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_customer_proto = out.File
	file_google_ads_googleads_v8_resources_customer_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_customer_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_customer_proto_depIdxs = nil
}
