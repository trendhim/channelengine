/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the SingleMerchantUpdatePurchaseOrderShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleMerchantUpdatePurchaseOrderShipmentRequest{}

// SingleMerchantUpdatePurchaseOrderShipmentRequest struct for SingleMerchantUpdatePurchaseOrderShipmentRequest
type SingleMerchantUpdatePurchaseOrderShipmentRequest struct {
	// The identifier of the channel
	ChannelId *int32 `json:"ChannelId,omitempty"`
	Model *UpdatePurchaseOrderShipment `json:"Model,omitempty"`
}

// NewSingleMerchantUpdatePurchaseOrderShipmentRequest instantiates a new SingleMerchantUpdatePurchaseOrderShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleMerchantUpdatePurchaseOrderShipmentRequest() *SingleMerchantUpdatePurchaseOrderShipmentRequest {
	this := SingleMerchantUpdatePurchaseOrderShipmentRequest{}
	return &this
}

// NewSingleMerchantUpdatePurchaseOrderShipmentRequestWithDefaults instantiates a new SingleMerchantUpdatePurchaseOrderShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleMerchantUpdatePurchaseOrderShipmentRequestWithDefaults() *SingleMerchantUpdatePurchaseOrderShipmentRequest {
	this := SingleMerchantUpdatePurchaseOrderShipmentRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetChannelId() int32 {
	if o == nil || IsNil(o.ChannelId) {
		var ret int32
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetChannelIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given int32 and assigns it to the ChannelId field.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) SetChannelId(v int32) {
	o.ChannelId = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetModel() UpdatePurchaseOrderShipment {
	if o == nil || IsNil(o.Model) {
		var ret UpdatePurchaseOrderShipment
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetModelOk() (*UpdatePurchaseOrderShipment, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given UpdatePurchaseOrderShipment and assigns it to the Model field.
func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) SetModel(v UpdatePurchaseOrderShipment) {
	o.Model = &v
}

func (o SingleMerchantUpdatePurchaseOrderShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleMerchantUpdatePurchaseOrderShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelId) {
		toSerialize["ChannelId"] = o.ChannelId
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	return toSerialize, nil
}

type NullableSingleMerchantUpdatePurchaseOrderShipmentRequest struct {
	value *SingleMerchantUpdatePurchaseOrderShipmentRequest
	isSet bool
}

func (v NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) Get() *SingleMerchantUpdatePurchaseOrderShipmentRequest {
	return v.value
}

func (v *NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) Set(val *SingleMerchantUpdatePurchaseOrderShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleMerchantUpdatePurchaseOrderShipmentRequest(val *SingleMerchantUpdatePurchaseOrderShipmentRequest) *NullableSingleMerchantUpdatePurchaseOrderShipmentRequest {
	return &NullableSingleMerchantUpdatePurchaseOrderShipmentRequest{value: val, isSet: true}
}

func (v NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleMerchantUpdatePurchaseOrderShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

