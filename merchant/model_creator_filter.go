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

// CreatorFilter the model 'CreatorFilter'
type CreatorFilter string

// List of CreatorFilter
const (
	CREATORFILTER_ONLY_FROM_MERCHANT CreatorFilter = "ONLY_FROM_MERCHANT"
	CREATORFILTER_ONLY_FROM_CHANNEL CreatorFilter = "ONLY_FROM_CHANNEL"
	CREATORFILTER_MIXED CreatorFilter = "MIXED"
)

// All allowed values of CreatorFilter enum
var AllowedCreatorFilterEnumValues = []CreatorFilter{
	"ONLY_FROM_MERCHANT",
	"ONLY_FROM_CHANNEL",
	"MIXED",
}

func (v *CreatorFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreatorFilter(value)
	for _, existing := range AllowedCreatorFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreatorFilter", value)
}

// NewCreatorFilterFromValue returns a pointer to a valid CreatorFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreatorFilterFromValue(v string) (*CreatorFilter, error) {
	ev := CreatorFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreatorFilter: valid values are %v", v, AllowedCreatorFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreatorFilter) IsValid() bool {
	for _, existing := range AllowedCreatorFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreatorFilter value
func (v CreatorFilter) Ptr() *CreatorFilter {
	return &v
}

type NullableCreatorFilter struct {
	value *CreatorFilter
	isSet bool
}

func (v NullableCreatorFilter) Get() *CreatorFilter {
	return v.value
}

func (v *NullableCreatorFilter) Set(val *CreatorFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatorFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatorFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatorFilter(val *CreatorFilter) *NullableCreatorFilter {
	return &NullableCreatorFilter{value: val, isSet: true}
}

func (v NullableCreatorFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatorFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

