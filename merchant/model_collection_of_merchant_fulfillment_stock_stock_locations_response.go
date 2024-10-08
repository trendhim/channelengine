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

// checks if the CollectionOfMerchantFulfillmentStockStockLocationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfMerchantFulfillmentStockStockLocationsResponse{}

// CollectionOfMerchantFulfillmentStockStockLocationsResponse struct for CollectionOfMerchantFulfillmentStockStockLocationsResponse
type CollectionOfMerchantFulfillmentStockStockLocationsResponse struct {
	Content []MerchantFulfillmentStockStockLocationsResponse `json:"Content,omitempty"`
	// The number of items in the current response.
	Count *int32 `json:"Count,omitempty"`
	// The total number of items.
	TotalCount *int32 `json:"TotalCount,omitempty"`
	// The number of items per page.
	ItemsPerPage *int32 `json:"ItemsPerPage,omitempty"`
	StatusCode *int32 `json:"StatusCode,omitempty"`
	RequestId NullableString `json:"RequestId,omitempty"`
	LogId NullableString `json:"LogId,omitempty"`
	Success *bool `json:"Success,omitempty"`
	Message NullableString `json:"Message,omitempty"`
	ExceptionType NullableString `json:"ExceptionType,omitempty"`
	ValidationErrors map[string][]string `json:"ValidationErrors,omitempty"`
}

// NewCollectionOfMerchantFulfillmentStockStockLocationsResponse instantiates a new CollectionOfMerchantFulfillmentStockStockLocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMerchantFulfillmentStockStockLocationsResponse() *CollectionOfMerchantFulfillmentStockStockLocationsResponse {
	this := CollectionOfMerchantFulfillmentStockStockLocationsResponse{}
	return &this
}

// NewCollectionOfMerchantFulfillmentStockStockLocationsResponseWithDefaults instantiates a new CollectionOfMerchantFulfillmentStockStockLocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMerchantFulfillmentStockStockLocationsResponseWithDefaults() *CollectionOfMerchantFulfillmentStockStockLocationsResponse {
	this := CollectionOfMerchantFulfillmentStockStockLocationsResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetContent() []MerchantFulfillmentStockStockLocationsResponse {
	if o == nil {
		var ret []MerchantFulfillmentStockStockLocationsResponse
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetContentOk() ([]MerchantFulfillmentStockStockLocationsResponse, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []MerchantFulfillmentStockStockLocationsResponse and assigns it to the Content field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetContent(v []MerchantFulfillmentStockStockLocationsResponse) {
	o.Content = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetItemsPerPage() int32 {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetLogId(v string) {
	o.LogId.Set(&v)
}
// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) UnsetLogId() {
	o.LogId.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) UnsetMessage() {
	o.Message.Unset()
}

// GetExceptionType returns the ExceptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetExceptionType() string {
	if o == nil || IsNil(o.ExceptionType.Get()) {
		var ret string
		return ret
	}
	return *o.ExceptionType.Get()
}

// GetExceptionTypeOk returns a tuple with the ExceptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetExceptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExceptionType.Get(), o.ExceptionType.IsSet()
}

// HasExceptionType returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasExceptionType() bool {
	if o != nil && o.ExceptionType.IsSet() {
		return true
	}

	return false
}

// SetExceptionType gets a reference to the given NullableString and assigns it to the ExceptionType field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetExceptionType(v string) {
	o.ExceptionType.Set(&v)
}
// SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetExceptionTypeNil() {
	o.ExceptionType.Set(nil)
}

// UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) UnsetExceptionType() {
	o.ExceptionType.Unset()
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetValidationErrors() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) GetValidationErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return &o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given map[string][]string and assigns it to the ValidationErrors field.
func (o *CollectionOfMerchantFulfillmentStockStockLocationsResponse) SetValidationErrors(v map[string][]string) {
	o.ValidationErrors = v
}

func (o CollectionOfMerchantFulfillmentStockStockLocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfMerchantFulfillmentStockStockLocationsResponse) ToMap() (map[string]interface{}, error) {
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
	if o.ExceptionType.IsSet() {
		toSerialize["ExceptionType"] = o.ExceptionType.Get()
	}
	if o.ValidationErrors != nil {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse struct {
	value *CollectionOfMerchantFulfillmentStockStockLocationsResponse
	isSet bool
}

func (v NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) Get() *CollectionOfMerchantFulfillmentStockStockLocationsResponse {
	return v.value
}

func (v *NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) Set(val *CollectionOfMerchantFulfillmentStockStockLocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMerchantFulfillmentStockStockLocationsResponse(val *CollectionOfMerchantFulfillmentStockStockLocationsResponse) *NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse {
	return &NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse{value: val, isSet: true}
}

func (v NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMerchantFulfillmentStockStockLocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


