/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// checks if the MerchantCancellationLineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCancellationLineResponse{}

// MerchantCancellationLineResponse struct for MerchantCancellationLineResponse
type MerchantCancellationLineResponse struct {
	// The unique cancellation line identifier used by ChannelEngine.
	Id *int32 `json:"Id,omitempty"`
	// The unique product reference used by the Merchant.
	MerchantProductNo NullableString `json:"MerchantProductNo,omitempty"`
	// The unique product reference used by the Channel.
	ChannelProductNo NullableString `json:"ChannelProductNo,omitempty"`
	// Quantity of the product to cancel.
	Quantity int32 `json:"Quantity"`
}

type _MerchantCancellationLineResponse MerchantCancellationLineResponse

// NewMerchantCancellationLineResponse instantiates a new MerchantCancellationLineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCancellationLineResponse(quantity int32) *MerchantCancellationLineResponse {
	this := MerchantCancellationLineResponse{}
	this.Quantity = quantity
	return &this
}

// NewMerchantCancellationLineResponseWithDefaults instantiates a new MerchantCancellationLineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCancellationLineResponseWithDefaults() *MerchantCancellationLineResponse {
	this := MerchantCancellationLineResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MerchantCancellationLineResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCancellationLineResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MerchantCancellationLineResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MerchantCancellationLineResponse) SetId(v int32) {
	o.Id = &v
}

// GetMerchantProductNo returns the MerchantProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCancellationLineResponse) GetMerchantProductNo() string {
	if o == nil || IsNil(o.MerchantProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantProductNo.Get()
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCancellationLineResponse) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantProductNo.Get(), o.MerchantProductNo.IsSet()
}

// HasMerchantProductNo returns a boolean if a field has been set.
func (o *MerchantCancellationLineResponse) HasMerchantProductNo() bool {
	if o != nil && o.MerchantProductNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantProductNo gets a reference to the given NullableString and assigns it to the MerchantProductNo field.
func (o *MerchantCancellationLineResponse) SetMerchantProductNo(v string) {
	o.MerchantProductNo.Set(&v)
}
// SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil
func (o *MerchantCancellationLineResponse) SetMerchantProductNoNil() {
	o.MerchantProductNo.Set(nil)
}

// UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
func (o *MerchantCancellationLineResponse) UnsetMerchantProductNo() {
	o.MerchantProductNo.Unset()
}

// GetChannelProductNo returns the ChannelProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantCancellationLineResponse) GetChannelProductNo() string {
	if o == nil || IsNil(o.ChannelProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelProductNo.Get()
}

// GetChannelProductNoOk returns a tuple with the ChannelProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantCancellationLineResponse) GetChannelProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProductNo.Get(), o.ChannelProductNo.IsSet()
}

// HasChannelProductNo returns a boolean if a field has been set.
func (o *MerchantCancellationLineResponse) HasChannelProductNo() bool {
	if o != nil && o.ChannelProductNo.IsSet() {
		return true
	}

	return false
}

// SetChannelProductNo gets a reference to the given NullableString and assigns it to the ChannelProductNo field.
func (o *MerchantCancellationLineResponse) SetChannelProductNo(v string) {
	o.ChannelProductNo.Set(&v)
}
// SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil
func (o *MerchantCancellationLineResponse) SetChannelProductNoNil() {
	o.ChannelProductNo.Set(nil)
}

// UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
func (o *MerchantCancellationLineResponse) UnsetChannelProductNo() {
	o.ChannelProductNo.Unset()
}

// GetQuantity returns the Quantity field value
func (o *MerchantCancellationLineResponse) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MerchantCancellationLineResponse) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MerchantCancellationLineResponse) SetQuantity(v int32) {
	o.Quantity = v
}

func (o MerchantCancellationLineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCancellationLineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.MerchantProductNo.IsSet() {
		toSerialize["MerchantProductNo"] = o.MerchantProductNo.Get()
	}
	if o.ChannelProductNo.IsSet() {
		toSerialize["ChannelProductNo"] = o.ChannelProductNo.Get()
	}
	toSerialize["Quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *MerchantCancellationLineResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Quantity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantCancellationLineResponse := _MerchantCancellationLineResponse{}

	err = json.Unmarshal(bytes, &varMerchantCancellationLineResponse)

	if err != nil {
		return err
	}

	*o = MerchantCancellationLineResponse(varMerchantCancellationLineResponse)

	return err
}

type NullableMerchantCancellationLineResponse struct {
	value *MerchantCancellationLineResponse
	isSet bool
}

func (v NullableMerchantCancellationLineResponse) Get() *MerchantCancellationLineResponse {
	return v.value
}

func (v *NullableMerchantCancellationLineResponse) Set(val *MerchantCancellationLineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCancellationLineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCancellationLineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCancellationLineResponse(val *MerchantCancellationLineResponse) *NullableMerchantCancellationLineResponse {
	return &NullableMerchantCancellationLineResponse{value: val, isSet: true}
}

func (v NullableMerchantCancellationLineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCancellationLineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

