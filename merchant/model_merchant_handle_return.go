/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the MerchantHandleReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantHandleReturn{}

// MerchantHandleReturn struct for MerchantHandleReturn
type MerchantHandleReturn struct {
	ReturnIdentifier NullableString `json:"ReturnIdentifier,omitempty"`
	ReturnLineIdentifier NullableString `json:"ReturnLineIdentifier,omitempty"`
	Quantity *int32 `json:"Quantity,omitempty"`
	Action *ReturnHandlingAction `json:"Action,omitempty"`
}

// NewMerchantHandleReturn instantiates a new MerchantHandleReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantHandleReturn() *MerchantHandleReturn {
	this := MerchantHandleReturn{}
	return &this
}

// NewMerchantHandleReturnWithDefaults instantiates a new MerchantHandleReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantHandleReturnWithDefaults() *MerchantHandleReturn {
	this := MerchantHandleReturn{}
	return &this
}

// GetReturnIdentifier returns the ReturnIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantHandleReturn) GetReturnIdentifier() string {
	if o == nil || IsNil(o.ReturnIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ReturnIdentifier.Get()
}

// GetReturnIdentifierOk returns a tuple with the ReturnIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantHandleReturn) GetReturnIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnIdentifier.Get(), o.ReturnIdentifier.IsSet()
}

// HasReturnIdentifier returns a boolean if a field has been set.
func (o *MerchantHandleReturn) HasReturnIdentifier() bool {
	if o != nil && o.ReturnIdentifier.IsSet() {
		return true
	}

	return false
}

// SetReturnIdentifier gets a reference to the given NullableString and assigns it to the ReturnIdentifier field.
func (o *MerchantHandleReturn) SetReturnIdentifier(v string) {
	o.ReturnIdentifier.Set(&v)
}
// SetReturnIdentifierNil sets the value for ReturnIdentifier to be an explicit nil
func (o *MerchantHandleReturn) SetReturnIdentifierNil() {
	o.ReturnIdentifier.Set(nil)
}

// UnsetReturnIdentifier ensures that no value is present for ReturnIdentifier, not even an explicit nil
func (o *MerchantHandleReturn) UnsetReturnIdentifier() {
	o.ReturnIdentifier.Unset()
}

// GetReturnLineIdentifier returns the ReturnLineIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantHandleReturn) GetReturnLineIdentifier() string {
	if o == nil || IsNil(o.ReturnLineIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ReturnLineIdentifier.Get()
}

// GetReturnLineIdentifierOk returns a tuple with the ReturnLineIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantHandleReturn) GetReturnLineIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnLineIdentifier.Get(), o.ReturnLineIdentifier.IsSet()
}

// HasReturnLineIdentifier returns a boolean if a field has been set.
func (o *MerchantHandleReturn) HasReturnLineIdentifier() bool {
	if o != nil && o.ReturnLineIdentifier.IsSet() {
		return true
	}

	return false
}

// SetReturnLineIdentifier gets a reference to the given NullableString and assigns it to the ReturnLineIdentifier field.
func (o *MerchantHandleReturn) SetReturnLineIdentifier(v string) {
	o.ReturnLineIdentifier.Set(&v)
}
// SetReturnLineIdentifierNil sets the value for ReturnLineIdentifier to be an explicit nil
func (o *MerchantHandleReturn) SetReturnLineIdentifierNil() {
	o.ReturnLineIdentifier.Set(nil)
}

// UnsetReturnLineIdentifier ensures that no value is present for ReturnLineIdentifier, not even an explicit nil
func (o *MerchantHandleReturn) UnsetReturnLineIdentifier() {
	o.ReturnLineIdentifier.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MerchantHandleReturn) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantHandleReturn) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MerchantHandleReturn) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *MerchantHandleReturn) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *MerchantHandleReturn) GetAction() ReturnHandlingAction {
	if o == nil || IsNil(o.Action) {
		var ret ReturnHandlingAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantHandleReturn) GetActionOk() (*ReturnHandlingAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *MerchantHandleReturn) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ReturnHandlingAction and assigns it to the Action field.
func (o *MerchantHandleReturn) SetAction(v ReturnHandlingAction) {
	o.Action = &v
}

func (o MerchantHandleReturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantHandleReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReturnIdentifier.IsSet() {
		toSerialize["ReturnIdentifier"] = o.ReturnIdentifier.Get()
	}
	if o.ReturnLineIdentifier.IsSet() {
		toSerialize["ReturnLineIdentifier"] = o.ReturnLineIdentifier.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	return toSerialize, nil
}

type NullableMerchantHandleReturn struct {
	value *MerchantHandleReturn
	isSet bool
}

func (v NullableMerchantHandleReturn) Get() *MerchantHandleReturn {
	return v.value
}

func (v *NullableMerchantHandleReturn) Set(val *MerchantHandleReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantHandleReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantHandleReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantHandleReturn(val *MerchantHandleReturn) *NullableMerchantHandleReturn {
	return &NullableMerchantHandleReturn{value: val, isSet: true}
}

func (v NullableMerchantHandleReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantHandleReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


