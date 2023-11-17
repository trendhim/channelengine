/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
)

// checks if the SingleOfDictionaryOfStringAndListOfString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleOfDictionaryOfStringAndListOfString{}

// SingleOfDictionaryOfStringAndListOfString struct for SingleOfDictionaryOfStringAndListOfString
type SingleOfDictionaryOfStringAndListOfString struct {
	Content map[string][]string `json:"Content,omitempty"`
	StatusCode *int32 `json:"StatusCode,omitempty"`
	RequestId NullableString `json:"RequestId,omitempty"`
	LogId NullableString `json:"LogId,omitempty"`
	Success *bool `json:"Success,omitempty"`
	Message NullableString `json:"Message,omitempty"`
	ValidationErrors map[string][]string `json:"ValidationErrors,omitempty"`
}

// NewSingleOfDictionaryOfStringAndListOfString instantiates a new SingleOfDictionaryOfStringAndListOfString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleOfDictionaryOfStringAndListOfString() *SingleOfDictionaryOfStringAndListOfString {
	this := SingleOfDictionaryOfStringAndListOfString{}
	return &this
}

// NewSingleOfDictionaryOfStringAndListOfStringWithDefaults instantiates a new SingleOfDictionaryOfStringAndListOfString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleOfDictionaryOfStringAndListOfStringWithDefaults() *SingleOfDictionaryOfStringAndListOfString {
	this := SingleOfDictionaryOfStringAndListOfString{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDictionaryOfStringAndListOfString) GetContent() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDictionaryOfStringAndListOfString) GetContentOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasContent() bool {
	if o != nil && IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given map[string][]string and assigns it to the Content field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetContent(v map[string][]string) {
	o.Content = v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SingleOfDictionaryOfStringAndListOfString) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDictionaryOfStringAndListOfString) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDictionaryOfStringAndListOfString) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDictionaryOfStringAndListOfString) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDictionaryOfStringAndListOfString) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetLogId(v string) {
	o.LogId.Set(&v)
}
// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) UnsetLogId() {
	o.LogId.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SingleOfDictionaryOfStringAndListOfString) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDictionaryOfStringAndListOfString) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDictionaryOfStringAndListOfString) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *SingleOfDictionaryOfStringAndListOfString) UnsetMessage() {
	o.Message.Unset()
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDictionaryOfStringAndListOfString) GetValidationErrors() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDictionaryOfStringAndListOfString) GetValidationErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return &o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *SingleOfDictionaryOfStringAndListOfString) HasValidationErrors() bool {
	if o != nil && IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given map[string][]string and assigns it to the ValidationErrors field.
func (o *SingleOfDictionaryOfStringAndListOfString) SetValidationErrors(v map[string][]string) {
	o.ValidationErrors = v
}

func (o SingleOfDictionaryOfStringAndListOfString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleOfDictionaryOfStringAndListOfString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if !IsNil(o.StatusCode) {
		toSerialize["StatusCode"] = o.StatusCode
	}
	if o.RequestId.IsSet() {
		toSerialize["RequestId"] = o.RequestId.Get()
	}
	if o.LogId.IsSet() {
		toSerialize["LogId"] = o.LogId.Get()
	}
	if !IsNil(o.Success) {
		toSerialize["Success"] = o.Success
	}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	if o.ValidationErrors != nil {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableSingleOfDictionaryOfStringAndListOfString struct {
	value *SingleOfDictionaryOfStringAndListOfString
	isSet bool
}

func (v NullableSingleOfDictionaryOfStringAndListOfString) Get() *SingleOfDictionaryOfStringAndListOfString {
	return v.value
}

func (v *NullableSingleOfDictionaryOfStringAndListOfString) Set(val *SingleOfDictionaryOfStringAndListOfString) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleOfDictionaryOfStringAndListOfString) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleOfDictionaryOfStringAndListOfString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleOfDictionaryOfStringAndListOfString(val *SingleOfDictionaryOfStringAndListOfString) *NullableSingleOfDictionaryOfStringAndListOfString {
	return &NullableSingleOfDictionaryOfStringAndListOfString{value: val, isSet: true}
}

func (v NullableSingleOfDictionaryOfStringAndListOfString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleOfDictionaryOfStringAndListOfString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

