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
// source: google/ads/googleads/v8/enums/lead_form_field_user_input_type.proto

package enums

import (
	proto "github.com/golang/protobuf/proto"
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

// Enum describing the input type of a lead form field.
type LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType int32

const (
	// Not specified.
	LeadFormFieldUserInputTypeEnum_UNSPECIFIED LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 0
	// Used for return value only. Represents value unknown in this version.
	LeadFormFieldUserInputTypeEnum_UNKNOWN LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1
	// The user will be asked to fill in their given and family name. This field
	// cannot be set at the same time as GIVEN_NAME or FAMILY_NAME.
	LeadFormFieldUserInputTypeEnum_FULL_NAME LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 2
	// The user will be asked to fill in their email address.
	LeadFormFieldUserInputTypeEnum_EMAIL LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 3
	// The user will be asked to fill in their phone number.
	LeadFormFieldUserInputTypeEnum_PHONE_NUMBER LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 4
	// The user will be asked to fill in their zip code.
	LeadFormFieldUserInputTypeEnum_POSTAL_CODE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 5
	// The user will be asked to fill in their city.
	LeadFormFieldUserInputTypeEnum_CITY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 9
	// The user will be asked to fill in their region part of the address (e.g.
	// state for US, province for Canada).
	LeadFormFieldUserInputTypeEnum_REGION LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 10
	// The user will be asked to fill in their country.
	LeadFormFieldUserInputTypeEnum_COUNTRY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 11
	// The user will be asked to fill in their work email address.
	LeadFormFieldUserInputTypeEnum_WORK_EMAIL LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 12
	// The user will be asked to fill in their company name.
	LeadFormFieldUserInputTypeEnum_COMPANY_NAME LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 13
	// The user will be asked to fill in their work phone.
	LeadFormFieldUserInputTypeEnum_WORK_PHONE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 14
	// The user will be asked to fill in their job title.
	LeadFormFieldUserInputTypeEnum_JOB_TITLE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 15
	// The user will be asked to fill in their first name. This
	// field can not be set at the same time as FULL_NAME.
	LeadFormFieldUserInputTypeEnum_FIRST_NAME LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 23
	// The user will be asked to fill in their last name. This
	// field can not be set at the same time as FULL_NAME.
	LeadFormFieldUserInputTypeEnum_LAST_NAME LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 24
	// Question: "Which model are you interested in?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_MODEL LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1001
	// Question: "Which type of vehicle are you interested in?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_TYPE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1002
	// Question: "What is your preferred dealership?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_PREFERRED_DEALERSHIP LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1003
	// Question: "When do you plan on purchasing a vehicle?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_PURCHASE_TIMELINE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1004
	// Question: "Do you own a vehicle?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_OWNERSHIP LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1005
	// Question: "What vehicle ownership option are you interested in?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_PAYMENT_TYPE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1009
	// Question: "What type of vehicle condition are you interested in?"
	// Category: "Auto"
	LeadFormFieldUserInputTypeEnum_VEHICLE_CONDITION LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1010
	// Question: "What size is your company?"
	// Category: "Business"
	LeadFormFieldUserInputTypeEnum_COMPANY_SIZE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1006
	// Question: "What is your annual sales volume?"
	// Category: "Business"
	LeadFormFieldUserInputTypeEnum_ANNUAL_SALES LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1007
	// Question: "How many years have you been in business?"
	// Category: "Business"
	LeadFormFieldUserInputTypeEnum_YEARS_IN_BUSINESS LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1008
	// Question: "What is your job department?"
	// Category: "Business"
	LeadFormFieldUserInputTypeEnum_JOB_DEPARTMENT LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1011
	// Question: "What is your job role?"
	// Category: "Business"
	LeadFormFieldUserInputTypeEnum_JOB_ROLE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1012
	// Question: "Which program are you interested in?"
	// Category: "Education"
	LeadFormFieldUserInputTypeEnum_EDUCATION_PROGRAM LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1013
	// Question: "Which course are you interested in?"
	// Category: "Education"
	LeadFormFieldUserInputTypeEnum_EDUCATION_COURSE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1014
	// Question: "Which product are you interested in?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_PRODUCT LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1016
	// Question: "Which service are you interested in?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_SERVICE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1017
	// Question: "Which offer are you interested in?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_OFFER LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1018
	// Question: "Which category are you interested in?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_CATEGORY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1019
	// Question: "What is your preferred method of contact?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_PREFERRED_CONTACT_METHOD LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1020
	// Question: "What is your preferred location?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_PREFERRED_LOCATION LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1021
	// Question: "What is the best time to contact you?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_PREFERRED_CONTACT_TIME LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1022
	// Question: "When are you looking to make a purchase?"
	// Category: "General"
	LeadFormFieldUserInputTypeEnum_PURCHASE_TIMELINE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1023
	// Question: "How many years of work experience do you have?"
	// Category: "Jobs"
	LeadFormFieldUserInputTypeEnum_YEARS_OF_EXPERIENCE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1048
	// Question: "What industry do you work in?"
	// Category: "Jobs"
	LeadFormFieldUserInputTypeEnum_JOB_INDUSTRY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1049
	// Question: "What is your highest level of education?"
	// Category: "Jobs"
	LeadFormFieldUserInputTypeEnum_LEVEL_OF_EDUCATION LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1050
	// Question: "What type of property are you looking for?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_PROPERTY_TYPE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1024
	// Question: "What do you need a realtor's help with?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_REALTOR_HELP_GOAL LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1025
	// Question: "What neighborhood are you interested in?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_PROPERTY_COMMUNITY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1026
	// Question: "What price range are you looking for?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_PRICE_RANGE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1027
	// Question: "How many bedrooms are you looking for?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_NUMBER_OF_BEDROOMS LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1028
	// Question: "Are you looking for a fully furnished property?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_FURNISHED_PROPERTY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1029
	// Question: "Are you looking for properties that allow pets?"
	// Category: "Real Estate"
	LeadFormFieldUserInputTypeEnum_PETS_ALLOWED_PROPERTY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1030
	// Question: "What is the next product you plan to purchase?"
	// Category: "Retail"
	LeadFormFieldUserInputTypeEnum_NEXT_PLANNED_PURCHASE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1031
	// Question: "Would you like to sign up for an event?"
	// Category: "Retail"
	LeadFormFieldUserInputTypeEnum_EVENT_SIGNUP_INTEREST LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1033
	// Question: "Where are you interested in shopping?"
	// Category: "Retail"
	LeadFormFieldUserInputTypeEnum_PREFERRED_SHOPPING_PLACES LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1034
	// Question: "What is your favorite brand?"
	// Category: "Retail"
	LeadFormFieldUserInputTypeEnum_FAVORITE_BRAND LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1035
	// Question: "Which type of valid commercial license do you have?"
	// Category: "Transportation"
	LeadFormFieldUserInputTypeEnum_TRANSPORTATION_COMMERCIAL_LICENSE_TYPE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1036
	// Question: "Interested in booking an event?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_EVENT_BOOKING_INTEREST LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1038
	// Question: "What is your destination country?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_DESTINATION_COUNTRY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1039
	// Question: "What is your destination city?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_DESTINATION_CITY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1040
	// Question: "What is your departure country?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_DEPARTURE_COUNTRY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1041
	// Question: "What is your departure city?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_DEPARTURE_CITY LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1042
	// Question: "What is your departure date?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_DEPARTURE_DATE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1043
	// Question: "What is your return date?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_RETURN_DATE LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1044
	// Question: "How many people are you traveling with?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_NUMBER_OF_TRAVELERS LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1045
	// Question: "What is your travel budget?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_TRAVEL_BUDGET LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1046
	// Question: "Where do you want to stay during your travel?"
	// Category: "Travel"
	LeadFormFieldUserInputTypeEnum_TRAVEL_ACCOMMODATION LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType = 1047
)

// Enum value maps for LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType.
var (
	LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType_name = map[int32]string{
		0:    "UNSPECIFIED",
		1:    "UNKNOWN",
		2:    "FULL_NAME",
		3:    "EMAIL",
		4:    "PHONE_NUMBER",
		5:    "POSTAL_CODE",
		9:    "CITY",
		10:   "REGION",
		11:   "COUNTRY",
		12:   "WORK_EMAIL",
		13:   "COMPANY_NAME",
		14:   "WORK_PHONE",
		15:   "JOB_TITLE",
		23:   "FIRST_NAME",
		24:   "LAST_NAME",
		1001: "VEHICLE_MODEL",
		1002: "VEHICLE_TYPE",
		1003: "PREFERRED_DEALERSHIP",
		1004: "VEHICLE_PURCHASE_TIMELINE",
		1005: "VEHICLE_OWNERSHIP",
		1009: "VEHICLE_PAYMENT_TYPE",
		1010: "VEHICLE_CONDITION",
		1006: "COMPANY_SIZE",
		1007: "ANNUAL_SALES",
		1008: "YEARS_IN_BUSINESS",
		1011: "JOB_DEPARTMENT",
		1012: "JOB_ROLE",
		1013: "EDUCATION_PROGRAM",
		1014: "EDUCATION_COURSE",
		1016: "PRODUCT",
		1017: "SERVICE",
		1018: "OFFER",
		1019: "CATEGORY",
		1020: "PREFERRED_CONTACT_METHOD",
		1021: "PREFERRED_LOCATION",
		1022: "PREFERRED_CONTACT_TIME",
		1023: "PURCHASE_TIMELINE",
		1048: "YEARS_OF_EXPERIENCE",
		1049: "JOB_INDUSTRY",
		1050: "LEVEL_OF_EDUCATION",
		1024: "PROPERTY_TYPE",
		1025: "REALTOR_HELP_GOAL",
		1026: "PROPERTY_COMMUNITY",
		1027: "PRICE_RANGE",
		1028: "NUMBER_OF_BEDROOMS",
		1029: "FURNISHED_PROPERTY",
		1030: "PETS_ALLOWED_PROPERTY",
		1031: "NEXT_PLANNED_PURCHASE",
		1033: "EVENT_SIGNUP_INTEREST",
		1034: "PREFERRED_SHOPPING_PLACES",
		1035: "FAVORITE_BRAND",
		1036: "TRANSPORTATION_COMMERCIAL_LICENSE_TYPE",
		1038: "EVENT_BOOKING_INTEREST",
		1039: "DESTINATION_COUNTRY",
		1040: "DESTINATION_CITY",
		1041: "DEPARTURE_COUNTRY",
		1042: "DEPARTURE_CITY",
		1043: "DEPARTURE_DATE",
		1044: "RETURN_DATE",
		1045: "NUMBER_OF_TRAVELERS",
		1046: "TRAVEL_BUDGET",
		1047: "TRAVEL_ACCOMMODATION",
	}
	LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType_value = map[string]int32{
		"UNSPECIFIED":                            0,
		"UNKNOWN":                                1,
		"FULL_NAME":                              2,
		"EMAIL":                                  3,
		"PHONE_NUMBER":                           4,
		"POSTAL_CODE":                            5,
		"CITY":                                   9,
		"REGION":                                 10,
		"COUNTRY":                                11,
		"WORK_EMAIL":                             12,
		"COMPANY_NAME":                           13,
		"WORK_PHONE":                             14,
		"JOB_TITLE":                              15,
		"FIRST_NAME":                             23,
		"LAST_NAME":                              24,
		"VEHICLE_MODEL":                          1001,
		"VEHICLE_TYPE":                           1002,
		"PREFERRED_DEALERSHIP":                   1003,
		"VEHICLE_PURCHASE_TIMELINE":              1004,
		"VEHICLE_OWNERSHIP":                      1005,
		"VEHICLE_PAYMENT_TYPE":                   1009,
		"VEHICLE_CONDITION":                      1010,
		"COMPANY_SIZE":                           1006,
		"ANNUAL_SALES":                           1007,
		"YEARS_IN_BUSINESS":                      1008,
		"JOB_DEPARTMENT":                         1011,
		"JOB_ROLE":                               1012,
		"EDUCATION_PROGRAM":                      1013,
		"EDUCATION_COURSE":                       1014,
		"PRODUCT":                                1016,
		"SERVICE":                                1017,
		"OFFER":                                  1018,
		"CATEGORY":                               1019,
		"PREFERRED_CONTACT_METHOD":               1020,
		"PREFERRED_LOCATION":                     1021,
		"PREFERRED_CONTACT_TIME":                 1022,
		"PURCHASE_TIMELINE":                      1023,
		"YEARS_OF_EXPERIENCE":                    1048,
		"JOB_INDUSTRY":                           1049,
		"LEVEL_OF_EDUCATION":                     1050,
		"PROPERTY_TYPE":                          1024,
		"REALTOR_HELP_GOAL":                      1025,
		"PROPERTY_COMMUNITY":                     1026,
		"PRICE_RANGE":                            1027,
		"NUMBER_OF_BEDROOMS":                     1028,
		"FURNISHED_PROPERTY":                     1029,
		"PETS_ALLOWED_PROPERTY":                  1030,
		"NEXT_PLANNED_PURCHASE":                  1031,
		"EVENT_SIGNUP_INTEREST":                  1033,
		"PREFERRED_SHOPPING_PLACES":              1034,
		"FAVORITE_BRAND":                         1035,
		"TRANSPORTATION_COMMERCIAL_LICENSE_TYPE": 1036,
		"EVENT_BOOKING_INTEREST":                 1038,
		"DESTINATION_COUNTRY":                    1039,
		"DESTINATION_CITY":                       1040,
		"DEPARTURE_COUNTRY":                      1041,
		"DEPARTURE_CITY":                         1042,
		"DEPARTURE_DATE":                         1043,
		"RETURN_DATE":                            1044,
		"NUMBER_OF_TRAVELERS":                    1045,
		"TRAVEL_BUDGET":                          1046,
		"TRAVEL_ACCOMMODATION":                   1047,
	}
)

func (x LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) Enum() *LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType {
	p := new(LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType)
	*p = x
	return p
}

func (x LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_enumTypes[0].Descriptor()
}

func (LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_enumTypes[0]
}

func (x LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType.Descriptor instead.
func (LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescGZIP(), []int{0, 0}
}

// Describes the input type of a lead form field.
type LeadFormFieldUserInputTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeadFormFieldUserInputTypeEnum) Reset() {
	*x = LeadFormFieldUserInputTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadFormFieldUserInputTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadFormFieldUserInputTypeEnum) ProtoMessage() {}

func (x *LeadFormFieldUserInputTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadFormFieldUserInputTypeEnum.ProtoReflect.Descriptor instead.
func (*LeadFormFieldUserInputTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xef, 0x0a, 0x0a, 0x1e, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xcc, 0x0a, 0x0a, 0x1a, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f,
	0x72, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x4f, 0x53, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x49, 0x54, 0x59, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59,
	0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x50, 0x48, 0x4f,
	0x4e, 0x45, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x4f, 0x42, 0x5f, 0x54, 0x49, 0x54, 0x4c,
	0x45, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x17, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
	0x10, 0x18, 0x12, 0x12, 0x0a, 0x0d, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x4c, 0x10, 0xe9, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0xea, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x52, 0x45,
	0x46, 0x45, 0x52, 0x52, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x53, 0x48, 0x49,
	0x50, 0x10, 0xeb, 0x07, 0x12, 0x1e, 0x0a, 0x19, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f,
	0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0xec, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f,
	0x4f, 0x57, 0x4e, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x10, 0xed, 0x07, 0x12, 0x19, 0x0a, 0x14,
	0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0xf1, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x56, 0x45, 0x48, 0x49, 0x43,
	0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xf2, 0x07, 0x12,
	0x11, 0x0a, 0x0c, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10,
	0xee, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x41, 0x4e, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x53, 0x41, 0x4c,
	0x45, 0x53, 0x10, 0xef, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x59, 0x45, 0x41, 0x52, 0x53, 0x5f, 0x49,
	0x4e, 0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0xf0, 0x07, 0x12, 0x13, 0x0a,
	0x0e, 0x4a, 0x4f, 0x42, 0x5f, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0xf3, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x4a, 0x4f, 0x42, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0xf4,
	0x07, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xf5, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x44, 0x55,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0xf6, 0x07,
	0x12, 0x0c, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0xf8, 0x07, 0x12, 0x0c,
	0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xf9, 0x07, 0x12, 0x0a, 0x0a, 0x05,
	0x4f, 0x46, 0x46, 0x45, 0x52, 0x10, 0xfa, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x10, 0xfb, 0x07, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x52, 0x45, 0x46, 0x45,
	0x52, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x10, 0xfc, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52,
	0x52, 0x45, 0x44, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xfd, 0x07, 0x12,
	0x1b, 0x0a, 0x16, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0xfe, 0x07, 0x12, 0x16, 0x0a, 0x11,
	0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0xff, 0x07, 0x12, 0x18, 0x0a, 0x13, 0x59, 0x45, 0x41, 0x52, 0x53, 0x5f, 0x4f, 0x46,
	0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x98, 0x08, 0x12, 0x11,
	0x0a, 0x0c, 0x4a, 0x4f, 0x42, 0x5f, 0x49, 0x4e, 0x44, 0x55, 0x53, 0x54, 0x52, 0x59, 0x10, 0x99,
	0x08, 0x12, 0x17, 0x0a, 0x12, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4f, 0x46, 0x5f, 0x45, 0x44,
	0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x9a, 0x08, 0x12, 0x12, 0x0a, 0x0d, 0x50, 0x52,
	0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x80, 0x08, 0x12, 0x16,
	0x0a, 0x11, 0x52, 0x45, 0x41, 0x4c, 0x54, 0x4f, 0x52, 0x5f, 0x48, 0x45, 0x4c, 0x50, 0x5f, 0x47,
	0x4f, 0x41, 0x4c, 0x10, 0x81, 0x08, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52,
	0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x54, 0x59, 0x10, 0x82, 0x08, 0x12,
	0x10, 0x0a, 0x0b, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x83,
	0x08, 0x12, 0x17, 0x0a, 0x12, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x5f, 0x42,
	0x45, 0x44, 0x52, 0x4f, 0x4f, 0x4d, 0x53, 0x10, 0x84, 0x08, 0x12, 0x17, 0x0a, 0x12, 0x46, 0x55,
	0x52, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59,
	0x10, 0x85, 0x08, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x45, 0x54, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x10, 0x86, 0x08, 0x12,
	0x1a, 0x0a, 0x15, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x5f,
	0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x10, 0x87, 0x08, 0x12, 0x1a, 0x0a, 0x15, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x45, 0x53, 0x54, 0x10, 0x89, 0x08, 0x12, 0x1e, 0x0a, 0x19, 0x50, 0x52, 0x45, 0x46, 0x45,
	0x52, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c,
	0x41, 0x43, 0x45, 0x53, 0x10, 0x8a, 0x08, 0x12, 0x13, 0x0a, 0x0e, 0x46, 0x41, 0x56, 0x4f, 0x52,
	0x49, 0x54, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x8b, 0x08, 0x12, 0x2b, 0x0a, 0x26,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x45, 0x52, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x8c, 0x08, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x45, 0x53, 0x54, 0x10, 0x8e, 0x08, 0x12, 0x18, 0x0a, 0x13, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x8f, 0x08,
	0x12, 0x15, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x49, 0x54, 0x59, 0x10, 0x90, 0x08, 0x12, 0x16, 0x0a, 0x11, 0x44, 0x45, 0x50, 0x41, 0x52,
	0x54, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x91, 0x08, 0x12,
	0x13, 0x0a, 0x0e, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x49, 0x54,
	0x59, 0x10, 0x92, 0x08, 0x12, 0x13, 0x0a, 0x0e, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x93, 0x08, 0x12, 0x10, 0x0a, 0x0b, 0x52, 0x45, 0x54,
	0x55, 0x52, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x94, 0x08, 0x12, 0x18, 0x0a, 0x13, 0x4e,
	0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x45,
	0x52, 0x53, 0x10, 0x95, 0x08, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f,
	0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x96, 0x08, 0x12, 0x19, 0x0a, 0x14, 0x54, 0x52, 0x41,
	0x56, 0x45, 0x4c, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x44, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x97, 0x08, 0x42, 0xf4, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1f, 0x4c, 0x65, 0x61, 0x64,
	0x46, 0x6f, 0x72, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescData = file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDesc
)

func file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDescData
}

var file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_goTypes = []interface{}{
	(LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType)(0), // 0: google.ads.googleads.v8.enums.LeadFormFieldUserInputTypeEnum.LeadFormFieldUserInputType
	(*LeadFormFieldUserInputTypeEnum)(nil),                         // 1: google.ads.googleads.v8.enums.LeadFormFieldUserInputTypeEnum
}
var file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_init() }
func file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_init() {
	if File_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadFormFieldUserInputTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto = out.File
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_rawDesc = nil
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_goTypes = nil
	file_google_ads_googleads_v8_enums_lead_form_field_user_input_type_proto_depIdxs = nil
}
