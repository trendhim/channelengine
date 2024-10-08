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

// checks if the ProductMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMessage{}

// ProductMessage struct for ProductMessage
type ProductMessage struct {
	Name NullableString `json:"Name,omitempty"`
	Reference NullableString `json:"Reference,omitempty"`
	KeyReference NullableString `json:"KeyReference,omitempty"`
	Warnings []string `json:"Warnings,omitempty"`
	Errors []string `json:"Errors,omitempty"`
}

// NewProductMessage instantiates a new ProductMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMessage() *ProductMessage {
	this := ProductMessage{}
	return &this
}

// NewProductMessageWithDefaults instantiates a new ProductMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMessageWithDefaults() *ProductMessage {
	this := ProductMessage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductMessage) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductMessage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProductMessage) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProductMessage) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProductMessage) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProductMessage) UnsetName() {
	o.Name.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductMessage) GetReference() string {
	if o == nil || IsNil(o.Reference.Get()) {
		var ret string
		return ret
	}
	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductMessage) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// HasReference returns a boolean if a field has been set.
func (o *ProductMessage) HasReference() bool {
	if o != nil && o.Reference.IsSet() {
		return true
	}

	return false
}

// SetReference gets a reference to the given NullableString and assigns it to the Reference field.
func (o *ProductMessage) SetReference(v string) {
	o.Reference.Set(&v)
}
// SetReferenceNil sets the value for Reference to be an explicit nil
func (o *ProductMessage) SetReferenceNil() {
	o.Reference.Set(nil)
}

// UnsetReference ensures that no value is present for Reference, not even an explicit nil
func (o *ProductMessage) UnsetReference() {
	o.Reference.Unset()
}

// GetKeyReference returns the KeyReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductMessage) GetKeyReference() string {
	if o == nil || IsNil(o.KeyReference.Get()) {
		var ret string
		return ret
	}
	return *o.KeyReference.Get()
}

// GetKeyReferenceOk returns a tuple with the KeyReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductMessage) GetKeyReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyReference.Get(), o.KeyReference.IsSet()
}

// HasKeyReference returns a boolean if a field has been set.
func (o *ProductMessage) HasKeyReference() bool {
	if o != nil && o.KeyReference.IsSet() {
		return true
	}

	return false
}

// SetKeyReference gets a reference to the given NullableString and assigns it to the KeyReference field.
func (o *ProductMessage) SetKeyReference(v string) {
	o.KeyReference.Set(&v)
}
// SetKeyReferenceNil sets the value for KeyReference to be an explicit nil
func (o *ProductMessage) SetKeyReferenceNil() {
	o.KeyReference.Set(nil)
}

// UnsetKeyReference ensures that no value is present for KeyReference, not even an explicit nil
func (o *ProductMessage) UnsetKeyReference() {
	o.KeyReference.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductMessage) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductMessage) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProductMessage) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ProductMessage) SetWarnings(v []string) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductMessage) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductMessage) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductMessage) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ProductMessage) SetErrors(v []string) {
	o.Errors = v
}

func (o ProductMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Reference.IsSet() {
		toSerialize["Reference"] = o.Reference.Get()
	}
	if o.KeyReference.IsSet() {
		toSerialize["KeyReference"] = o.KeyReference.Get()
	}
	if o.Warnings != nil {
		toSerialize["Warnings"] = o.Warnings
	}
	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableProductMessage struct {
	value *ProductMessage
	isSet bool
}

func (v NullableProductMessage) Get() *ProductMessage {
	return v.value
}

func (v *NullableProductMessage) Set(val *ProductMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMessage(val *ProductMessage) *NullableProductMessage {
	return &NullableProductMessage{value: val, isSet: true}
}

func (v NullableProductMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


