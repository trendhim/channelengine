/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
)

// checks if the ChannelCreateReturnLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelCreateReturnLine{}

// ChannelCreateReturnLine struct for ChannelCreateReturnLine
type ChannelCreateReturnLine struct {
	OrderLineIdentifier NullableString `json:"OrderLineIdentifier,omitempty"`
	Quantity *int32 `json:"Quantity,omitempty"`
	ChannelReturnLineNo NullableString `json:"ChannelReturnLineNo,omitempty"`
	ExtraData map[string]string `json:"ExtraData,omitempty"`
	StockLocationId NullableInt32 `json:"StockLocationId,omitempty"`
}

// NewChannelCreateReturnLine instantiates a new ChannelCreateReturnLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelCreateReturnLine() *ChannelCreateReturnLine {
	this := ChannelCreateReturnLine{}
	return &this
}

// NewChannelCreateReturnLineWithDefaults instantiates a new ChannelCreateReturnLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelCreateReturnLineWithDefaults() *ChannelCreateReturnLine {
	this := ChannelCreateReturnLine{}
	return &this
}

// GetOrderLineIdentifier returns the OrderLineIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCreateReturnLine) GetOrderLineIdentifier() string {
	if o == nil || IsNil(o.OrderLineIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.OrderLineIdentifier.Get()
}

// GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCreateReturnLine) GetOrderLineIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderLineIdentifier.Get(), o.OrderLineIdentifier.IsSet()
}

// HasOrderLineIdentifier returns a boolean if a field has been set.
func (o *ChannelCreateReturnLine) HasOrderLineIdentifier() bool {
	if o != nil && o.OrderLineIdentifier.IsSet() {
		return true
	}

	return false
}

// SetOrderLineIdentifier gets a reference to the given NullableString and assigns it to the OrderLineIdentifier field.
func (o *ChannelCreateReturnLine) SetOrderLineIdentifier(v string) {
	o.OrderLineIdentifier.Set(&v)
}
// SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil
func (o *ChannelCreateReturnLine) SetOrderLineIdentifierNil() {
	o.OrderLineIdentifier.Set(nil)
}

// UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
func (o *ChannelCreateReturnLine) UnsetOrderLineIdentifier() {
	o.OrderLineIdentifier.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ChannelCreateReturnLine) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelCreateReturnLine) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ChannelCreateReturnLine) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ChannelCreateReturnLine) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetChannelReturnLineNo returns the ChannelReturnLineNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCreateReturnLine) GetChannelReturnLineNo() string {
	if o == nil || IsNil(o.ChannelReturnLineNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelReturnLineNo.Get()
}

// GetChannelReturnLineNoOk returns a tuple with the ChannelReturnLineNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCreateReturnLine) GetChannelReturnLineNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelReturnLineNo.Get(), o.ChannelReturnLineNo.IsSet()
}

// HasChannelReturnLineNo returns a boolean if a field has been set.
func (o *ChannelCreateReturnLine) HasChannelReturnLineNo() bool {
	if o != nil && o.ChannelReturnLineNo.IsSet() {
		return true
	}

	return false
}

// SetChannelReturnLineNo gets a reference to the given NullableString and assigns it to the ChannelReturnLineNo field.
func (o *ChannelCreateReturnLine) SetChannelReturnLineNo(v string) {
	o.ChannelReturnLineNo.Set(&v)
}
// SetChannelReturnLineNoNil sets the value for ChannelReturnLineNo to be an explicit nil
func (o *ChannelCreateReturnLine) SetChannelReturnLineNoNil() {
	o.ChannelReturnLineNo.Set(nil)
}

// UnsetChannelReturnLineNo ensures that no value is present for ChannelReturnLineNo, not even an explicit nil
func (o *ChannelCreateReturnLine) UnsetChannelReturnLineNo() {
	o.ChannelReturnLineNo.Unset()
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCreateReturnLine) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCreateReturnLine) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *ChannelCreateReturnLine) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *ChannelCreateReturnLine) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

// GetStockLocationId returns the StockLocationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCreateReturnLine) GetStockLocationId() int32 {
	if o == nil || IsNil(o.StockLocationId.Get()) {
		var ret int32
		return ret
	}
	return *o.StockLocationId.Get()
}

// GetStockLocationIdOk returns a tuple with the StockLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCreateReturnLine) GetStockLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockLocationId.Get(), o.StockLocationId.IsSet()
}

// HasStockLocationId returns a boolean if a field has been set.
func (o *ChannelCreateReturnLine) HasStockLocationId() bool {
	if o != nil && o.StockLocationId.IsSet() {
		return true
	}

	return false
}

// SetStockLocationId gets a reference to the given NullableInt32 and assigns it to the StockLocationId field.
func (o *ChannelCreateReturnLine) SetStockLocationId(v int32) {
	o.StockLocationId.Set(&v)
}
// SetStockLocationIdNil sets the value for StockLocationId to be an explicit nil
func (o *ChannelCreateReturnLine) SetStockLocationIdNil() {
	o.StockLocationId.Set(nil)
}

// UnsetStockLocationId ensures that no value is present for StockLocationId, not even an explicit nil
func (o *ChannelCreateReturnLine) UnsetStockLocationId() {
	o.StockLocationId.Unset()
}

func (o ChannelCreateReturnLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelCreateReturnLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderLineIdentifier.IsSet() {
		toSerialize["OrderLineIdentifier"] = o.OrderLineIdentifier.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if o.ChannelReturnLineNo.IsSet() {
		toSerialize["ChannelReturnLineNo"] = o.ChannelReturnLineNo.Get()
	}
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	if o.StockLocationId.IsSet() {
		toSerialize["StockLocationId"] = o.StockLocationId.Get()
	}
	return toSerialize, nil
}

type NullableChannelCreateReturnLine struct {
	value *ChannelCreateReturnLine
	isSet bool
}

func (v NullableChannelCreateReturnLine) Get() *ChannelCreateReturnLine {
	return v.value
}

func (v *NullableChannelCreateReturnLine) Set(val *ChannelCreateReturnLine) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCreateReturnLine) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCreateReturnLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCreateReturnLine(val *ChannelCreateReturnLine) *NullableChannelCreateReturnLine {
	return &NullableChannelCreateReturnLine{value: val, isSet: true}
}

func (v NullableChannelCreateReturnLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCreateReturnLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


