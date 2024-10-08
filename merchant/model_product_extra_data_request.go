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

// checks if the ProductExtraDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductExtraDataRequest{}

// ProductExtraDataRequest struct for ProductExtraDataRequest
type ProductExtraDataRequest struct {
	Op string `json:"Op"`
	Key string `json:"Key"`
	Value NullableString `json:"Value,omitempty"`
}

type _ProductExtraDataRequest ProductExtraDataRequest

// NewProductExtraDataRequest instantiates a new ProductExtraDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductExtraDataRequest(op string, key string) *ProductExtraDataRequest {
	this := ProductExtraDataRequest{}
	this.Op = op
	this.Key = key
	return &this
}

// NewProductExtraDataRequestWithDefaults instantiates a new ProductExtraDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductExtraDataRequestWithDefaults() *ProductExtraDataRequest {
	this := ProductExtraDataRequest{}
	return &this
}

// GetOp returns the Op field value
func (o *ProductExtraDataRequest) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *ProductExtraDataRequest) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *ProductExtraDataRequest) SetOp(v string) {
	o.Op = v
}

// GetKey returns the Key field value
func (o *ProductExtraDataRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProductExtraDataRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProductExtraDataRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductExtraDataRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductExtraDataRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProductExtraDataRequest) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProductExtraDataRequest) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ProductExtraDataRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProductExtraDataRequest) UnsetValue() {
	o.Value.Unset()
}

func (o ProductExtraDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductExtraDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Op"] = o.Op
	toSerialize["Key"] = o.Key
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	return toSerialize, nil
}

func (o *ProductExtraDataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Op",
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

	varProductExtraDataRequest := _ProductExtraDataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductExtraDataRequest)

	if err != nil {
		return err
	}

	*o = ProductExtraDataRequest(varProductExtraDataRequest)

	return err
}

type NullableProductExtraDataRequest struct {
	value *ProductExtraDataRequest
	isSet bool
}

func (v NullableProductExtraDataRequest) Get() *ProductExtraDataRequest {
	return v.value
}

func (v *NullableProductExtraDataRequest) Set(val *ProductExtraDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductExtraDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductExtraDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductExtraDataRequest(val *ProductExtraDataRequest) *NullableProductExtraDataRequest {
	return &NullableProductExtraDataRequest{value: val, isSet: true}
}

func (v NullableProductExtraDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductExtraDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


