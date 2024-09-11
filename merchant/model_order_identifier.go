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

// OrderIdentifier the model 'OrderIdentifier'
type OrderIdentifier string

// List of OrderIdentifier
const (
	ORDERIDENTIFIER_ORDER_ID OrderIdentifier = "ORDER_ID"
	ORDERIDENTIFIER_CHANNEL_ORDER_NO OrderIdentifier = "CHANNEL_ORDER_NO"
	ORDERIDENTIFIER_MERCHANT_ORDER_NO OrderIdentifier = "MERCHANT_ORDER_NO"
)

// All allowed values of OrderIdentifier enum
var AllowedOrderIdentifierEnumValues = []OrderIdentifier{
	"ORDER_ID",
	"CHANNEL_ORDER_NO",
	"MERCHANT_ORDER_NO",
}

func (v *OrderIdentifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderIdentifier(value)
	for _, existing := range AllowedOrderIdentifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderIdentifier", value)
}

// NewOrderIdentifierFromValue returns a pointer to a valid OrderIdentifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderIdentifierFromValue(v string) (*OrderIdentifier, error) {
	ev := OrderIdentifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderIdentifier: valid values are %v", v, AllowedOrderIdentifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderIdentifier) IsValid() bool {
	for _, existing := range AllowedOrderIdentifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderIdentifier value
func (v OrderIdentifier) Ptr() *OrderIdentifier {
	return &v
}

type NullableOrderIdentifier struct {
	value *OrderIdentifier
	isSet bool
}

func (v NullableOrderIdentifier) Get() *OrderIdentifier {
	return v.value
}

func (v *NullableOrderIdentifier) Set(val *OrderIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderIdentifier(val *OrderIdentifier) *NullableOrderIdentifier {
	return &NullableOrderIdentifier{value: val, isSet: true}
}

func (v NullableOrderIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
