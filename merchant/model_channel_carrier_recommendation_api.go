/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// ChannelCarrierRecommendationApi the model 'ChannelCarrierRecommendationApi'
type ChannelCarrierRecommendationApi string

// List of ChannelCarrierRecommendationApi
const (
	CHANNELCARRIERRECOMMENDATIONAPI_NEUTRAL ChannelCarrierRecommendationApi = "NEUTRAL"
	CHANNELCARRIERRECOMMENDATIONAPI_RECOMMENDED ChannelCarrierRecommendationApi = "RECOMMENDED"
	CHANNELCARRIERRECOMMENDATIONAPI_DISCOMMENDED ChannelCarrierRecommendationApi = "DISCOMMENDED"
)

// All allowed values of ChannelCarrierRecommendationApi enum
var AllowedChannelCarrierRecommendationApiEnumValues = []ChannelCarrierRecommendationApi{
	"NEUTRAL",
	"RECOMMENDED",
	"DISCOMMENDED",
}

func (v *ChannelCarrierRecommendationApi) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelCarrierRecommendationApi(value)
	for _, existing := range AllowedChannelCarrierRecommendationApiEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelCarrierRecommendationApi", value)
}

// NewChannelCarrierRecommendationApiFromValue returns a pointer to a valid ChannelCarrierRecommendationApi
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelCarrierRecommendationApiFromValue(v string) (*ChannelCarrierRecommendationApi, error) {
	ev := ChannelCarrierRecommendationApi(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelCarrierRecommendationApi: valid values are %v", v, AllowedChannelCarrierRecommendationApiEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelCarrierRecommendationApi) IsValid() bool {
	for _, existing := range AllowedChannelCarrierRecommendationApiEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelCarrierRecommendationApi value
func (v ChannelCarrierRecommendationApi) Ptr() *ChannelCarrierRecommendationApi {
	return &v
}

type NullableChannelCarrierRecommendationApi struct {
	value *ChannelCarrierRecommendationApi
	isSet bool
}

func (v NullableChannelCarrierRecommendationApi) Get() *ChannelCarrierRecommendationApi {
	return v.value
}

func (v *NullableChannelCarrierRecommendationApi) Set(val *ChannelCarrierRecommendationApi) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCarrierRecommendationApi) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCarrierRecommendationApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCarrierRecommendationApi(val *ChannelCarrierRecommendationApi) *NullableChannelCarrierRecommendationApi {
	return &NullableChannelCarrierRecommendationApi{value: val, isSet: true}
}

func (v NullableChannelCarrierRecommendationApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCarrierRecommendationApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

