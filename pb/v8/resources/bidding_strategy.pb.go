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
// source: google/ads/googleads/v8/resources/bidding_strategy.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v8/common"
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

// A bidding strategy.
type BiddingStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the bidding strategy.
	Id *int64 `protobuf:"varint,16,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *string `protobuf:"bytes,17,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The status of the bidding strategy.
	//
	// This field is read-only.
	Status enums.BiddingStrategyStatusEnum_BiddingStrategyStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v8.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus" json:"status,omitempty"`
	// Output only. The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v8.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// Immutable. The currency used by the bidding strategy (ISO 4217 three-letter code).
	//
	// For bidding strategies in manager customers, this currency can be set on
	// creation and defaults to the manager customer's currency. For serving
	// customers, this field cannot be set; all strategies in a serving customer
	// implicitly use the serving customer's currency. In all cases the
	// effective_currency_code field returns the currency used by the strategy.
	CurrencyCode string `protobuf:"bytes,23,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Output only. The currency used by the bidding strategy (ISO 4217 three-letter code).
	//
	// For bidding strategies in manager customers, this is the currency set by
	// the advertiser when creating the strategy. For serving customers, this is
	// the customer's currency_code.
	//
	// Bidding strategy metrics are reported in this currency.
	//
	// This field is read-only.
	EffectiveCurrencyCode *string `protobuf:"bytes,20,opt,name=effective_currency_code,json=effectiveCurrencyCode,proto3,oneof" json:"effective_currency_code,omitempty"`
	// Output only. The number of campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	CampaignCount *int64 `protobuf:"varint,18,opt,name=campaign_count,json=campaignCount,proto3,oneof" json:"campaign_count,omitempty"`
	// Output only. The number of non-removed campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	NonRemovedCampaignCount *int64 `protobuf:"varint,19,opt,name=non_removed_campaign_count,json=nonRemovedCampaignCount,proto3,oneof" json:"non_removed_campaign_count,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are assignable to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_MaximizeConversionValue
	//	*BiddingStrategy_MaximizeConversions
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetImpressionShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
}

func (x *BiddingStrategy) Reset() {
	*x = BiddingStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategy) ProtoMessage() {}

func (x *BiddingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategy.ProtoReflect.Descriptor instead.
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *BiddingStrategy) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BiddingStrategy) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BiddingStrategy) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BiddingStrategy) GetStatus() enums.BiddingStrategyStatusEnum_BiddingStrategyStatus {
	if x != nil {
		return x.Status
	}
	return enums.BiddingStrategyStatusEnum_UNSPECIFIED
}

func (x *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if x != nil {
		return x.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

func (x *BiddingStrategy) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *BiddingStrategy) GetEffectiveCurrencyCode() string {
	if x != nil && x.EffectiveCurrencyCode != nil {
		return *x.EffectiveCurrencyCode
	}
	return ""
}

func (x *BiddingStrategy) GetCampaignCount() int64 {
	if x != nil && x.CampaignCount != nil {
		return *x.CampaignCount
	}
	return 0
}

func (x *BiddingStrategy) GetNonRemovedCampaignCount() int64 {
	if x != nil && x.NonRemovedCampaignCount != nil {
		return *x.NonRemovedCampaignCount
	}
	return 0
}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (x *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := x.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (x *BiddingStrategy) GetMaximizeConversionValue() *common.MaximizeConversionValue {
	if x, ok := x.GetScheme().(*BiddingStrategy_MaximizeConversionValue); ok {
		return x.MaximizeConversionValue
	}
	return nil
}

func (x *BiddingStrategy) GetMaximizeConversions() *common.MaximizeConversions {
	if x, ok := x.GetScheme().(*BiddingStrategy_MaximizeConversions); ok {
		return x.MaximizeConversions
	}
	return nil
}

func (x *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (x *BiddingStrategy) GetTargetImpressionShare() *common.TargetImpressionShare {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetImpressionShare); ok {
		return x.TargetImpressionShare
	}
	return nil
}

func (x *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (x *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	// A bidding strategy that raises bids for clicks that seem more likely to
	// lead to a conversion and lowers them for clicks where they seem less
	// likely.
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_MaximizeConversionValue struct {
	// An automated bidding strategy to help get the most conversion value for
	// your campaigns while spending your budget.
	MaximizeConversionValue *common.MaximizeConversionValue `protobuf:"bytes,21,opt,name=maximize_conversion_value,json=maximizeConversionValue,proto3,oneof"`
}

type BiddingStrategy_MaximizeConversions struct {
	// An automated bidding strategy to help get the most conversions for your
	// campaigns while spending your budget.
	MaximizeConversions *common.MaximizeConversions `protobuf:"bytes,22,opt,name=maximize_conversions,json=maximizeConversions,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	// A bidding strategy that sets bids to help get as many conversions as
	// possible at the target cost-per-acquisition (CPA) you set.
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetImpressionShare struct {
	// A bidding strategy that automatically optimizes towards a desired
	// percentage of impressions.
	TargetImpressionShare *common.TargetImpressionShare `protobuf:"bytes,48,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	// A bidding strategy that helps you maximize revenue while averaging a
	// specific target Return On Ad Spend (ROAS).
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	// A bid strategy that sets your bids to help get as many clicks as
	// possible within your budget.
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_MaximizeConversionValue) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_MaximizeConversions) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetImpressionShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

var File_google_ads_googleads_v8_resources_bidding_strategy_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x0b,
	0x0a, 0x0f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x55, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2a,
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x6b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x63, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x17, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52,
	0x15, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0e, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1a, 0x6e, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x17, 0x6e, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x50, 0x0a, 0x0c, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x63, 0x70,
	0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x43, 0x70, 0x63, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x43, 0x70, 0x63, 0x12, 0x75, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x17, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x68, 0x0a, 0x14, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x69,
	0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52,
	0x13, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x70, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x70, 0x61, 0x48, 0x00, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x61,
	0x12, 0x6f, 0x0a, 0x17, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x48, 0x00, 0x52, 0x15, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x61, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f,
	0x61, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73,
	0x12, 0x50, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x3a, 0x6e, 0xea, 0x41, 0x6b, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x3f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1a, 0x0a, 0x18,
	0x5f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1d, 0x0a, 0x1b, 0x5f,
	0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x81, 0x02, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x14, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f,
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
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescData = file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_resources_bidding_strategy_proto_goTypes = []interface{}{
	(*BiddingStrategy)(nil), // 0: google.ads.googleads.v8.resources.BiddingStrategy
	(enums.BiddingStrategyStatusEnum_BiddingStrategyStatus)(0), // 1: google.ads.googleads.v8.enums.BiddingStrategyStatusEnum.BiddingStrategyStatus
	(enums.BiddingStrategyTypeEnum_BiddingStrategyType)(0),     // 2: google.ads.googleads.v8.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	(*common.EnhancedCpc)(nil),                                 // 3: google.ads.googleads.v8.common.EnhancedCpc
	(*common.MaximizeConversionValue)(nil),                     // 4: google.ads.googleads.v8.common.MaximizeConversionValue
	(*common.MaximizeConversions)(nil),                         // 5: google.ads.googleads.v8.common.MaximizeConversions
	(*common.TargetCpa)(nil),                                   // 6: google.ads.googleads.v8.common.TargetCpa
	(*common.TargetImpressionShare)(nil),                       // 7: google.ads.googleads.v8.common.TargetImpressionShare
	(*common.TargetRoas)(nil),                                  // 8: google.ads.googleads.v8.common.TargetRoas
	(*common.TargetSpend)(nil),                                 // 9: google.ads.googleads.v8.common.TargetSpend
}
var file_google_ads_googleads_v8_resources_bidding_strategy_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.BiddingStrategy.status:type_name -> google.ads.googleads.v8.enums.BiddingStrategyStatusEnum.BiddingStrategyStatus
	2, // 1: google.ads.googleads.v8.resources.BiddingStrategy.type:type_name -> google.ads.googleads.v8.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	3, // 2: google.ads.googleads.v8.resources.BiddingStrategy.enhanced_cpc:type_name -> google.ads.googleads.v8.common.EnhancedCpc
	4, // 3: google.ads.googleads.v8.resources.BiddingStrategy.maximize_conversion_value:type_name -> google.ads.googleads.v8.common.MaximizeConversionValue
	5, // 4: google.ads.googleads.v8.resources.BiddingStrategy.maximize_conversions:type_name -> google.ads.googleads.v8.common.MaximizeConversions
	6, // 5: google.ads.googleads.v8.resources.BiddingStrategy.target_cpa:type_name -> google.ads.googleads.v8.common.TargetCpa
	7, // 6: google.ads.googleads.v8.resources.BiddingStrategy.target_impression_share:type_name -> google.ads.googleads.v8.common.TargetImpressionShare
	8, // 7: google.ads.googleads.v8.resources.BiddingStrategy.target_roas:type_name -> google.ads.googleads.v8.common.TargetRoas
	9, // 8: google.ads.googleads.v8.resources.BiddingStrategy.target_spend:type_name -> google.ads.googleads.v8.common.TargetSpend
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_bidding_strategy_proto_init() }
func file_google_ads_googleads_v8_resources_bidding_strategy_proto_init() {
	if File_google_ads_googleads_v8_resources_bidding_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingStrategy); i {
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
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_MaximizeConversionValue)(nil),
		(*BiddingStrategy_MaximizeConversions)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetImpressionShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_bidding_strategy_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_bidding_strategy_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_bidding_strategy_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_bidding_strategy_proto = out.File
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_bidding_strategy_proto_depIdxs = nil
}
