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

// checks if the SingleMerchantAcknowledgeReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleMerchantAcknowledgeReturnRequest{}

// SingleMerchantAcknowledgeReturnRequest struct for SingleMerchantAcknowledgeReturnRequest
type SingleMerchantAcknowledgeReturnRequest struct {
	ChannelId NullableInt32 `json:"ChannelId,omitempty"`
	IdentifierType *ReturnIdentifier `json:"IdentifierType,omitempty"`
	Model *MerchantAcknowledgeReturn `json:"Model,omitempty"`
}

// NewSingleMerchantAcknowledgeReturnRequest instantiates a new SingleMerchantAcknowledgeReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleMerchantAcknowledgeReturnRequest() *SingleMerchantAcknowledgeReturnRequest {
	this := SingleMerchantAcknowledgeReturnRequest{}
	return &this
}

// NewSingleMerchantAcknowledgeReturnRequestWithDefaults instantiates a new SingleMerchantAcknowledgeReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleMerchantAcknowledgeReturnRequestWithDefaults() *SingleMerchantAcknowledgeReturnRequest {
	this := SingleMerchantAcknowledgeReturnRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleMerchantAcknowledgeReturnRequest) GetChannelId() int32 {
	if o == nil || IsNil(o.ChannelId.Get()) {
		var ret int32
		return ret
	}
	return *o.ChannelId.Get()
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleMerchantAcknowledgeReturnRequest) GetChannelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelId.Get(), o.ChannelId.IsSet()
}

// HasChannelId returns a boolean if a field has been set.
func (o *SingleMerchantAcknowledgeReturnRequest) HasChannelId() bool {
	if o != nil && o.ChannelId.IsSet() {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given NullableInt32 and assigns it to the ChannelId field.
func (o *SingleMerchantAcknowledgeReturnRequest) SetChannelId(v int32) {
	o.ChannelId.Set(&v)
}
// SetChannelIdNil sets the value for ChannelId to be an explicit nil
func (o *SingleMerchantAcknowledgeReturnRequest) SetChannelIdNil() {
	o.ChannelId.Set(nil)
}

// UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
func (o *SingleMerchantAcknowledgeReturnRequest) UnsetChannelId() {
	o.ChannelId.Unset()
}

// GetIdentifierType returns the IdentifierType field value if set, zero value otherwise.
func (o *SingleMerchantAcknowledgeReturnRequest) GetIdentifierType() ReturnIdentifier {
	if o == nil || IsNil(o.IdentifierType) {
		var ret ReturnIdentifier
		return ret
	}
	return *o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantAcknowledgeReturnRequest) GetIdentifierTypeOk() (*ReturnIdentifier, bool) {
	if o == nil || IsNil(o.IdentifierType) {
		return nil, false
	}
	return o.IdentifierType, true
}

// HasIdentifierType returns a boolean if a field has been set.
func (o *SingleMerchantAcknowledgeReturnRequest) HasIdentifierType() bool {
	if o != nil && !IsNil(o.IdentifierType) {
		return true
	}

	return false
}

// SetIdentifierType gets a reference to the given ReturnIdentifier and assigns it to the IdentifierType field.
func (o *SingleMerchantAcknowledgeReturnRequest) SetIdentifierType(v ReturnIdentifier) {
	o.IdentifierType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *SingleMerchantAcknowledgeReturnRequest) GetModel() MerchantAcknowledgeReturn {
	if o == nil || IsNil(o.Model) {
		var ret MerchantAcknowledgeReturn
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleMerchantAcknowledgeReturnRequest) GetModelOk() (*MerchantAcknowledgeReturn, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *SingleMerchantAcknowledgeReturnRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given MerchantAcknowledgeReturn and assigns it to the Model field.
func (o *SingleMerchantAcknowledgeReturnRequest) SetModel(v MerchantAcknowledgeReturn) {
	o.Model = &v
}

func (o SingleMerchantAcknowledgeReturnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleMerchantAcknowledgeReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelId.IsSet() {
		toSerialize["ChannelId"] = o.ChannelId.Get()
	}
	if !IsNil(o.IdentifierType) {
		toSerialize["IdentifierType"] = o.IdentifierType
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	return toSerialize, nil
}

type NullableSingleMerchantAcknowledgeReturnRequest struct {
	value *SingleMerchantAcknowledgeReturnRequest
	isSet bool
}

func (v NullableSingleMerchantAcknowledgeReturnRequest) Get() *SingleMerchantAcknowledgeReturnRequest {
	return v.value
}

func (v *NullableSingleMerchantAcknowledgeReturnRequest) Set(val *SingleMerchantAcknowledgeReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleMerchantAcknowledgeReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleMerchantAcknowledgeReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleMerchantAcknowledgeReturnRequest(val *SingleMerchantAcknowledgeReturnRequest) *NullableSingleMerchantAcknowledgeReturnRequest {
	return &NullableSingleMerchantAcknowledgeReturnRequest{value: val, isSet: true}
}

func (v NullableSingleMerchantAcknowledgeReturnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleMerchantAcknowledgeReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

