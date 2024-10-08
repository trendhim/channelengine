/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MerchantProductExtraDataItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantProductExtraDataItemRequest{}

// MerchantProductExtraDataItemRequest struct for MerchantProductExtraDataItemRequest
type MerchantProductExtraDataItemRequest struct {
	// Name of the extra data field.
	Key string `json:"Key"`
	// Value of the extra data field.
	Value NullableString `json:"Value,omitempty"`
	Type *ExtraDataType `json:"Type,omitempty"`
	// Add this field to the export of the product feed to the channel.
	IsPublic *bool `json:"IsPublic,omitempty"`
	// The 2-letter ISO code of the extra data
	LanguageIsoCode NullableString `json:"LanguageIsoCode,omitempty"`
}

type _MerchantProductExtraDataItemRequest MerchantProductExtraDataItemRequest

// NewMerchantProductExtraDataItemRequest instantiates a new MerchantProductExtraDataItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantProductExtraDataItemRequest(key string) *MerchantProductExtraDataItemRequest {
	this := MerchantProductExtraDataItemRequest{}
	this.Key = key
	return &this
}

// NewMerchantProductExtraDataItemRequestWithDefaults instantiates a new MerchantProductExtraDataItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantProductExtraDataItemRequestWithDefaults() *MerchantProductExtraDataItemRequest {
	this := MerchantProductExtraDataItemRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *MerchantProductExtraDataItemRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MerchantProductExtraDataItemRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MerchantProductExtraDataItemRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProductExtraDataItemRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProductExtraDataItemRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MerchantProductExtraDataItemRequest) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MerchantProductExtraDataItemRequest) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MerchantProductExtraDataItemRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MerchantProductExtraDataItemRequest) UnsetValue() {
	o.Value.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MerchantProductExtraDataItemRequest) GetType() ExtraDataType {
	if o == nil || IsNil(o.Type) {
		var ret ExtraDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductExtraDataItemRequest) GetTypeOk() (*ExtraDataType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MerchantProductExtraDataItemRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ExtraDataType and assigns it to the Type field.
func (o *MerchantProductExtraDataItemRequest) SetType(v ExtraDataType) {
	o.Type = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *MerchantProductExtraDataItemRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductExtraDataItemRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *MerchantProductExtraDataItemRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *MerchantProductExtraDataItemRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetLanguageIsoCode returns the LanguageIsoCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProductExtraDataItemRequest) GetLanguageIsoCode() string {
	if o == nil || IsNil(o.LanguageIsoCode.Get()) {
		var ret string
		return ret
	}
	return *o.LanguageIsoCode.Get()
}

// GetLanguageIsoCodeOk returns a tuple with the LanguageIsoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProductExtraDataItemRequest) GetLanguageIsoCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LanguageIsoCode.Get(), o.LanguageIsoCode.IsSet()
}

// HasLanguageIsoCode returns a boolean if a field has been set.
func (o *MerchantProductExtraDataItemRequest) HasLanguageIsoCode() bool {
	if o != nil && o.LanguageIsoCode.IsSet() {
		return true
	}

	return false
}

// SetLanguageIsoCode gets a reference to the given NullableString and assigns it to the LanguageIsoCode field.
func (o *MerchantProductExtraDataItemRequest) SetLanguageIsoCode(v string) {
	o.LanguageIsoCode.Set(&v)
}
// SetLanguageIsoCodeNil sets the value for LanguageIsoCode to be an explicit nil
func (o *MerchantProductExtraDataItemRequest) SetLanguageIsoCodeNil() {
	o.LanguageIsoCode.Set(nil)
}

// UnsetLanguageIsoCode ensures that no value is present for LanguageIsoCode, not even an explicit nil
func (o *MerchantProductExtraDataItemRequest) UnsetLanguageIsoCode() {
	o.LanguageIsoCode.Unset()
}

func (o MerchantProductExtraDataItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantProductExtraDataItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Key"] = o.Key
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.IsPublic) {
		toSerialize["IsPublic"] = o.IsPublic
	}
	if o.LanguageIsoCode.IsSet() {
		toSerialize["LanguageIsoCode"] = o.LanguageIsoCode.Get()
	}
	return toSerialize, nil
}

func (o *MerchantProductExtraDataItemRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantProductExtraDataItemRequest := _MerchantProductExtraDataItemRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantProductExtraDataItemRequest)

	if err != nil {
		return err
	}

	*o = MerchantProductExtraDataItemRequest(varMerchantProductExtraDataItemRequest)

	return err
}

type NullableMerchantProductExtraDataItemRequest struct {
	value *MerchantProductExtraDataItemRequest
	isSet bool
}

func (v NullableMerchantProductExtraDataItemRequest) Get() *MerchantProductExtraDataItemRequest {
	return v.value
}

func (v *NullableMerchantProductExtraDataItemRequest) Set(val *MerchantProductExtraDataItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantProductExtraDataItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantProductExtraDataItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantProductExtraDataItemRequest(val *MerchantProductExtraDataItemRequest) *NullableMerchantProductExtraDataItemRequest {
	return &NullableMerchantProductExtraDataItemRequest{value: val, isSet: true}
}

func (v NullableMerchantProductExtraDataItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantProductExtraDataItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


