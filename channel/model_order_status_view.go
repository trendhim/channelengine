/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
	"fmt"
)

// OrderStatusView the model 'OrderStatusView'
type OrderStatusView string

// List of OrderStatusView
const (
	ORDERSTATUSVIEW_IN_PROGRESS OrderStatusView = "IN_PROGRESS"
	ORDERSTATUSVIEW_SHIPPED OrderStatusView = "SHIPPED"
	ORDERSTATUSVIEW_IN_BACKORDER OrderStatusView = "IN_BACKORDER"
	ORDERSTATUSVIEW_MANCO OrderStatusView = "MANCO"
	ORDERSTATUSVIEW_CANCELED OrderStatusView = "CANCELED"
	ORDERSTATUSVIEW_IN_COMBI OrderStatusView = "IN_COMBI"
	ORDERSTATUSVIEW_CLOSED OrderStatusView = "CLOSED"
	ORDERSTATUSVIEW_NEW OrderStatusView = "NEW"
	ORDERSTATUSVIEW_RETURNED OrderStatusView = "RETURNED"
	ORDERSTATUSVIEW_REQUIRES_CORRECTION OrderStatusView = "REQUIRES_CORRECTION"
	ORDERSTATUSVIEW_AWAITING_PAYMENT OrderStatusView = "AWAITING_PAYMENT"
)

// All allowed values of OrderStatusView enum
var AllowedOrderStatusViewEnumValues = []OrderStatusView{
	"IN_PROGRESS",
	"SHIPPED",
	"IN_BACKORDER",
	"MANCO",
	"CANCELED",
	"IN_COMBI",
	"CLOSED",
	"NEW",
	"RETURNED",
	"REQUIRES_CORRECTION",
	"AWAITING_PAYMENT",
}

func (v *OrderStatusView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderStatusView(value)
	for _, existing := range AllowedOrderStatusViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderStatusView", value)
}

// NewOrderStatusViewFromValue returns a pointer to a valid OrderStatusView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderStatusViewFromValue(v string) (*OrderStatusView, error) {
	ev := OrderStatusView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderStatusView: valid values are %v", v, AllowedOrderStatusViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderStatusView) IsValid() bool {
	for _, existing := range AllowedOrderStatusViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderStatusView value
func (v OrderStatusView) Ptr() *OrderStatusView {
	return &v
}

type NullableOrderStatusView struct {
	value *OrderStatusView
	isSet bool
}

func (v NullableOrderStatusView) Get() *OrderStatusView {
	return v.value
}

func (v *NullableOrderStatusView) Set(val *OrderStatusView) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusView) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusView(val *OrderStatusView) *NullableOrderStatusView {
	return &NullableOrderStatusView{value: val, isSet: true}
}

func (v NullableOrderStatusView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

