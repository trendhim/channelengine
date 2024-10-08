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

// checks if the SingleOfDeleteTargetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleOfDeleteTargetsResponse{}

// SingleOfDeleteTargetsResponse struct for SingleOfDeleteTargetsResponse
type SingleOfDeleteTargetsResponse struct {
	Content *DeleteTargetsResponse `json:"Content,omitempty"`
	StatusCode *int32 `json:"StatusCode,omitempty"`
	RequestId NullableString `json:"RequestId,omitempty"`
	LogId NullableString `json:"LogId,omitempty"`
	Success *bool `json:"Success,omitempty"`
	Message NullableString `json:"Message,omitempty"`
	ExceptionType NullableString `json:"ExceptionType,omitempty"`
	ValidationErrors map[string][]string `json:"ValidationErrors,omitempty"`
}

// NewSingleOfDeleteTargetsResponse instantiates a new SingleOfDeleteTargetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleOfDeleteTargetsResponse() *SingleOfDeleteTargetsResponse {
	this := SingleOfDeleteTargetsResponse{}
	return &this
}

// NewSingleOfDeleteTargetsResponseWithDefaults instantiates a new SingleOfDeleteTargetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleOfDeleteTargetsResponseWithDefaults() *SingleOfDeleteTargetsResponse {
	this := SingleOfDeleteTargetsResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SingleOfDeleteTargetsResponse) GetContent() DeleteTargetsResponse {
	if o == nil || IsNil(o.Content) {
		var ret DeleteTargetsResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOfDeleteTargetsResponse) GetContentOk() (*DeleteTargetsResponse, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given DeleteTargetsResponse and assigns it to the Content field.
func (o *SingleOfDeleteTargetsResponse) SetContent(v DeleteTargetsResponse) {
	o.Content = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SingleOfDeleteTargetsResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOfDeleteTargetsResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *SingleOfDeleteTargetsResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDeleteTargetsResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDeleteTargetsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *SingleOfDeleteTargetsResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *SingleOfDeleteTargetsResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *SingleOfDeleteTargetsResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDeleteTargetsResponse) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDeleteTargetsResponse) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *SingleOfDeleteTargetsResponse) SetLogId(v string) {
	o.LogId.Set(&v)
}
// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *SingleOfDeleteTargetsResponse) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *SingleOfDeleteTargetsResponse) UnsetLogId() {
	o.LogId.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SingleOfDeleteTargetsResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOfDeleteTargetsResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SingleOfDeleteTargetsResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDeleteTargetsResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDeleteTargetsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *SingleOfDeleteTargetsResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *SingleOfDeleteTargetsResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *SingleOfDeleteTargetsResponse) UnsetMessage() {
	o.Message.Unset()
}

// GetExceptionType returns the ExceptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDeleteTargetsResponse) GetExceptionType() string {
	if o == nil || IsNil(o.ExceptionType.Get()) {
		var ret string
		return ret
	}
	return *o.ExceptionType.Get()
}

// GetExceptionTypeOk returns a tuple with the ExceptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDeleteTargetsResponse) GetExceptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExceptionType.Get(), o.ExceptionType.IsSet()
}

// HasExceptionType returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasExceptionType() bool {
	if o != nil && o.ExceptionType.IsSet() {
		return true
	}

	return false
}

// SetExceptionType gets a reference to the given NullableString and assigns it to the ExceptionType field.
func (o *SingleOfDeleteTargetsResponse) SetExceptionType(v string) {
	o.ExceptionType.Set(&v)
}
// SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil
func (o *SingleOfDeleteTargetsResponse) SetExceptionTypeNil() {
	o.ExceptionType.Set(nil)
}

// UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
func (o *SingleOfDeleteTargetsResponse) UnsetExceptionType() {
	o.ExceptionType.Unset()
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SingleOfDeleteTargetsResponse) GetValidationErrors() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleOfDeleteTargetsResponse) GetValidationErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return &o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *SingleOfDeleteTargetsResponse) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given map[string][]string and assigns it to the ValidationErrors field.
func (o *SingleOfDeleteTargetsResponse) SetValidationErrors(v map[string][]string) {
	o.ValidationErrors = v
}

func (o SingleOfDeleteTargetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleOfDeleteTargetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
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
	if o.ExceptionType.IsSet() {
		toSerialize["ExceptionType"] = o.ExceptionType.Get()
	}
	if o.ValidationErrors != nil {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableSingleOfDeleteTargetsResponse struct {
	value *SingleOfDeleteTargetsResponse
	isSet bool
}

func (v NullableSingleOfDeleteTargetsResponse) Get() *SingleOfDeleteTargetsResponse {
	return v.value
}

func (v *NullableSingleOfDeleteTargetsResponse) Set(val *SingleOfDeleteTargetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleOfDeleteTargetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleOfDeleteTargetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleOfDeleteTargetsResponse(val *SingleOfDeleteTargetsResponse) *NullableSingleOfDeleteTargetsResponse {
	return &NullableSingleOfDeleteTargetsResponse{value: val, isSet: true}
}

func (v NullableSingleOfDeleteTargetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleOfDeleteTargetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


