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

// checks if the CustomFieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldResponse{}

// CustomFieldResponse struct for CustomFieldResponse
type CustomFieldResponse struct {
	Id *int32 `json:"Id,omitempty"`
	Key NullableString `json:"Key,omitempty"`
	IsPublic *bool `json:"IsPublic,omitempty"`
	IsUsed *bool `json:"IsUsed,omitempty"`
}

// NewCustomFieldResponse instantiates a new CustomFieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldResponse() *CustomFieldResponse {
	this := CustomFieldResponse{}
	return &this
}

// NewCustomFieldResponseWithDefaults instantiates a new CustomFieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldResponseWithDefaults() *CustomFieldResponse {
	this := CustomFieldResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFieldResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFieldResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFieldResponse) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldResponse) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CustomFieldResponse) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CustomFieldResponse) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CustomFieldResponse) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CustomFieldResponse) UnsetKey() {
	o.Key.Unset()
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *CustomFieldResponse) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *CustomFieldResponse) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *CustomFieldResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetIsUsed returns the IsUsed field value if set, zero value otherwise.
func (o *CustomFieldResponse) GetIsUsed() bool {
	if o == nil || IsNil(o.IsUsed) {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldResponse) GetIsUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUsed) {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *CustomFieldResponse) HasIsUsed() bool {
	if o != nil && !IsNil(o.IsUsed) {
		return true
	}

	return false
}

// SetIsUsed gets a reference to the given bool and assigns it to the IsUsed field.
func (o *CustomFieldResponse) SetIsUsed(v bool) {
	o.IsUsed = &v
}

func (o CustomFieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	if !IsNil(o.IsPublic) {
		toSerialize["IsPublic"] = o.IsPublic
	}
	if !IsNil(o.IsUsed) {
		toSerialize["IsUsed"] = o.IsUsed
	}
	return toSerialize, nil
}

type NullableCustomFieldResponse struct {
	value *CustomFieldResponse
	isSet bool
}

func (v NullableCustomFieldResponse) Get() *CustomFieldResponse {
	return v.value
}

func (v *NullableCustomFieldResponse) Set(val *CustomFieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldResponse(val *CustomFieldResponse) *NullableCustomFieldResponse {
	return &NullableCustomFieldResponse{value: val, isSet: true}
}

func (v NullableCustomFieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


