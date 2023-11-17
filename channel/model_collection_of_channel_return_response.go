/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
)

// checks if the CollectionOfChannelReturnResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfChannelReturnResponse{}

// CollectionOfChannelReturnResponse struct for CollectionOfChannelReturnResponse
type CollectionOfChannelReturnResponse struct {
	Content []ChannelReturnResponse `json:"Content,omitempty"`
	Count *int32 `json:"Count,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty"`
	ItemsPerPage *int32 `json:"ItemsPerPage,omitempty"`
	StatusCode *int32 `json:"StatusCode,omitempty"`
	RequestId NullableString `json:"RequestId,omitempty"`
	LogId NullableString `json:"LogId,omitempty"`
	Success *bool `json:"Success,omitempty"`
	Message NullableString `json:"Message,omitempty"`
	ValidationErrors map[string][]string `json:"ValidationErrors,omitempty"`
}

// NewCollectionOfChannelReturnResponse instantiates a new CollectionOfChannelReturnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfChannelReturnResponse() *CollectionOfChannelReturnResponse {
	this := CollectionOfChannelReturnResponse{}
	return &this
}

// NewCollectionOfChannelReturnResponseWithDefaults instantiates a new CollectionOfChannelReturnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfChannelReturnResponseWithDefaults() *CollectionOfChannelReturnResponse {
	this := CollectionOfChannelReturnResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfChannelReturnResponse) GetContent() []ChannelReturnResponse {
	if o == nil {
		var ret []ChannelReturnResponse
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfChannelReturnResponse) GetContentOk() ([]ChannelReturnResponse, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasContent() bool {
	if o != nil && IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []ChannelReturnResponse and assigns it to the Content field.
func (o *CollectionOfChannelReturnResponse) SetContent(v []ChannelReturnResponse) {
	o.Content = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CollectionOfChannelReturnResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChannelReturnResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CollectionOfChannelReturnResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CollectionOfChannelReturnResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChannelReturnResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *CollectionOfChannelReturnResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *CollectionOfChannelReturnResponse) GetItemsPerPage() int32 {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChannelReturnResponse) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *CollectionOfChannelReturnResponse) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *CollectionOfChannelReturnResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChannelReturnResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *CollectionOfChannelReturnResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfChannelReturnResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfChannelReturnResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *CollectionOfChannelReturnResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *CollectionOfChannelReturnResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *CollectionOfChannelReturnResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfChannelReturnResponse) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfChannelReturnResponse) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *CollectionOfChannelReturnResponse) SetLogId(v string) {
	o.LogId.Set(&v)
}
// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *CollectionOfChannelReturnResponse) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *CollectionOfChannelReturnResponse) UnsetLogId() {
	o.LogId.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CollectionOfChannelReturnResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChannelReturnResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CollectionOfChannelReturnResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfChannelReturnResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfChannelReturnResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *CollectionOfChannelReturnResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *CollectionOfChannelReturnResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CollectionOfChannelReturnResponse) UnsetMessage() {
	o.Message.Unset()
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfChannelReturnResponse) GetValidationErrors() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfChannelReturnResponse) GetValidationErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return &o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *CollectionOfChannelReturnResponse) HasValidationErrors() bool {
	if o != nil && IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given map[string][]string and assigns it to the ValidationErrors field.
func (o *CollectionOfChannelReturnResponse) SetValidationErrors(v map[string][]string) {
	o.ValidationErrors = v
}

func (o CollectionOfChannelReturnResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfChannelReturnResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if !IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	if !IsNil(o.ItemsPerPage) {
		toSerialize["ItemsPerPage"] = o.ItemsPerPage
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

type NullableCollectionOfChannelReturnResponse struct {
	value *CollectionOfChannelReturnResponse
	isSet bool
}

func (v NullableCollectionOfChannelReturnResponse) Get() *CollectionOfChannelReturnResponse {
	return v.value
}

func (v *NullableCollectionOfChannelReturnResponse) Set(val *CollectionOfChannelReturnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfChannelReturnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfChannelReturnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfChannelReturnResponse(val *CollectionOfChannelReturnResponse) *NullableCollectionOfChannelReturnResponse {
	return &NullableCollectionOfChannelReturnResponse{value: val, isSet: true}
}

func (v NullableCollectionOfChannelReturnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfChannelReturnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

