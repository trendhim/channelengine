/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ChannelReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelReturnRequest{}

// ChannelReturnRequest struct for ChannelReturnRequest
type ChannelReturnRequest struct {
	// The unique order reference used by the Channel.
	ChannelOrderNo NullableString `json:"ChannelOrderNo,omitempty"`
	// The unique order reference used by the Merchant (sku).
	MerchantOrderNo NullableString `json:"MerchantOrderNo,omitempty"`
	// The unique return reference used by the Channel.
	ChannelReference string `json:"ChannelReference"`
	// Optional. Is the MON used as key for the order (default value is false).
	KeyIsMerchantOrderNo *bool `json:"KeyIsMerchantOrderNo,omitempty"`
	// Optional. Is the MPN used as key for the product (default value is false).
	KeyIsMerchantProductNo *bool `json:"KeyIsMerchantProductNo,omitempty"`
	Lines []ChannelReturnLineRequest `json:"Lines"`
	Status *ChannelReturnStatus `json:"Status,omitempty"`
	// The unique return reference used by ChannelEngine.
	Id *int32 `json:"Id,omitempty"`
	Reason *ReturnReason `json:"Reason,omitempty"`
	// Optional. Comment of customer on the (reason of) the return.
	CustomerComment NullableString `json:"CustomerComment,omitempty"`
	// Optional. Comment of merchant on the return.
	MerchantComment NullableString `json:"MerchantComment,omitempty"`
	// Refund amount incl. VAT.
	RefundInclVat *float32 `json:"RefundInclVat,omitempty"`
	// Refund amount excl. VAT.
	RefundExclVat *float32 `json:"RefundExclVat,omitempty"`
	// The date at which the return was originally created in the source system (if available).
	ReturnDate NullableTime `json:"ReturnDate,omitempty"`
	// Extra data on the return. Each item must have an unqiue key
	ExtraData map[string]string `json:"ExtraData,omitempty"`
}

type _ChannelReturnRequest ChannelReturnRequest

// NewChannelReturnRequest instantiates a new ChannelReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelReturnRequest(channelReference string, lines []ChannelReturnLineRequest) *ChannelReturnRequest {
	this := ChannelReturnRequest{}
	this.ChannelReference = channelReference
	this.Lines = lines
	return &this
}

// NewChannelReturnRequestWithDefaults instantiates a new ChannelReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelReturnRequestWithDefaults() *ChannelReturnRequest {
	this := ChannelReturnRequest{}
	return &this
}

// GetChannelOrderNo returns the ChannelOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetChannelOrderNo() string {
	if o == nil || IsNil(o.ChannelOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelOrderNo.Get()
}

// GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetChannelOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelOrderNo.Get(), o.ChannelOrderNo.IsSet()
}

// HasChannelOrderNo returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasChannelOrderNo() bool {
	if o != nil && o.ChannelOrderNo.IsSet() {
		return true
	}

	return false
}

// SetChannelOrderNo gets a reference to the given NullableString and assigns it to the ChannelOrderNo field.
func (o *ChannelReturnRequest) SetChannelOrderNo(v string) {
	o.ChannelOrderNo.Set(&v)
}
// SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil
func (o *ChannelReturnRequest) SetChannelOrderNoNil() {
	o.ChannelOrderNo.Set(nil)
}

// UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
func (o *ChannelReturnRequest) UnsetChannelOrderNo() {
	o.ChannelOrderNo.Unset()
}

// GetMerchantOrderNo returns the MerchantOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetMerchantOrderNo() string {
	if o == nil || IsNil(o.MerchantOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantOrderNo.Get()
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantOrderNo.Get(), o.MerchantOrderNo.IsSet()
}

// HasMerchantOrderNo returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasMerchantOrderNo() bool {
	if o != nil && o.MerchantOrderNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantOrderNo gets a reference to the given NullableString and assigns it to the MerchantOrderNo field.
func (o *ChannelReturnRequest) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo.Set(&v)
}
// SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil
func (o *ChannelReturnRequest) SetMerchantOrderNoNil() {
	o.MerchantOrderNo.Set(nil)
}

// UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
func (o *ChannelReturnRequest) UnsetMerchantOrderNo() {
	o.MerchantOrderNo.Unset()
}

// GetChannelReference returns the ChannelReference field value
func (o *ChannelReturnRequest) GetChannelReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetChannelReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelReference, true
}

// SetChannelReference sets field value
func (o *ChannelReturnRequest) SetChannelReference(v string) {
	o.ChannelReference = v
}

// GetKeyIsMerchantOrderNo returns the KeyIsMerchantOrderNo field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetKeyIsMerchantOrderNo() bool {
	if o == nil || IsNil(o.KeyIsMerchantOrderNo) {
		var ret bool
		return ret
	}
	return *o.KeyIsMerchantOrderNo
}

// GetKeyIsMerchantOrderNoOk returns a tuple with the KeyIsMerchantOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetKeyIsMerchantOrderNoOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyIsMerchantOrderNo) {
		return nil, false
	}
	return o.KeyIsMerchantOrderNo, true
}

// HasKeyIsMerchantOrderNo returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasKeyIsMerchantOrderNo() bool {
	if o != nil && !IsNil(o.KeyIsMerchantOrderNo) {
		return true
	}

	return false
}

// SetKeyIsMerchantOrderNo gets a reference to the given bool and assigns it to the KeyIsMerchantOrderNo field.
func (o *ChannelReturnRequest) SetKeyIsMerchantOrderNo(v bool) {
	o.KeyIsMerchantOrderNo = &v
}

// GetKeyIsMerchantProductNo returns the KeyIsMerchantProductNo field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetKeyIsMerchantProductNo() bool {
	if o == nil || IsNil(o.KeyIsMerchantProductNo) {
		var ret bool
		return ret
	}
	return *o.KeyIsMerchantProductNo
}

// GetKeyIsMerchantProductNoOk returns a tuple with the KeyIsMerchantProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetKeyIsMerchantProductNoOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyIsMerchantProductNo) {
		return nil, false
	}
	return o.KeyIsMerchantProductNo, true
}

// HasKeyIsMerchantProductNo returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasKeyIsMerchantProductNo() bool {
	if o != nil && !IsNil(o.KeyIsMerchantProductNo) {
		return true
	}

	return false
}

// SetKeyIsMerchantProductNo gets a reference to the given bool and assigns it to the KeyIsMerchantProductNo field.
func (o *ChannelReturnRequest) SetKeyIsMerchantProductNo(v bool) {
	o.KeyIsMerchantProductNo = &v
}

// GetLines returns the Lines field value
func (o *ChannelReturnRequest) GetLines() []ChannelReturnLineRequest {
	if o == nil {
		var ret []ChannelReturnLineRequest
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetLinesOk() ([]ChannelReturnLineRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *ChannelReturnRequest) SetLines(v []ChannelReturnLineRequest) {
	o.Lines = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetStatus() ChannelReturnStatus {
	if o == nil || IsNil(o.Status) {
		var ret ChannelReturnStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetStatusOk() (*ChannelReturnStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ChannelReturnStatus and assigns it to the Status field.
func (o *ChannelReturnRequest) SetStatus(v ChannelReturnStatus) {
	o.Status = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ChannelReturnRequest) SetId(v int32) {
	o.Id = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetReason() ReturnReason {
	if o == nil || IsNil(o.Reason) {
		var ret ReturnReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetReasonOk() (*ReturnReason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given ReturnReason and assigns it to the Reason field.
func (o *ChannelReturnRequest) SetReason(v ReturnReason) {
	o.Reason = &v
}

// GetCustomerComment returns the CustomerComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetCustomerComment() string {
	if o == nil || IsNil(o.CustomerComment.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerComment.Get()
}

// GetCustomerCommentOk returns a tuple with the CustomerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetCustomerCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerComment.Get(), o.CustomerComment.IsSet()
}

// HasCustomerComment returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasCustomerComment() bool {
	if o != nil && o.CustomerComment.IsSet() {
		return true
	}

	return false
}

// SetCustomerComment gets a reference to the given NullableString and assigns it to the CustomerComment field.
func (o *ChannelReturnRequest) SetCustomerComment(v string) {
	o.CustomerComment.Set(&v)
}
// SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil
func (o *ChannelReturnRequest) SetCustomerCommentNil() {
	o.CustomerComment.Set(nil)
}

// UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
func (o *ChannelReturnRequest) UnsetCustomerComment() {
	o.CustomerComment.Unset()
}

// GetMerchantComment returns the MerchantComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetMerchantComment() string {
	if o == nil || IsNil(o.MerchantComment.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantComment.Get()
}

// GetMerchantCommentOk returns a tuple with the MerchantComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetMerchantCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantComment.Get(), o.MerchantComment.IsSet()
}

// HasMerchantComment returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasMerchantComment() bool {
	if o != nil && o.MerchantComment.IsSet() {
		return true
	}

	return false
}

// SetMerchantComment gets a reference to the given NullableString and assigns it to the MerchantComment field.
func (o *ChannelReturnRequest) SetMerchantComment(v string) {
	o.MerchantComment.Set(&v)
}
// SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil
func (o *ChannelReturnRequest) SetMerchantCommentNil() {
	o.MerchantComment.Set(nil)
}

// UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
func (o *ChannelReturnRequest) UnsetMerchantComment() {
	o.MerchantComment.Unset()
}

// GetRefundInclVat returns the RefundInclVat field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetRefundInclVat() float32 {
	if o == nil || IsNil(o.RefundInclVat) {
		var ret float32
		return ret
	}
	return *o.RefundInclVat
}

// GetRefundInclVatOk returns a tuple with the RefundInclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetRefundInclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundInclVat) {
		return nil, false
	}
	return o.RefundInclVat, true
}

// HasRefundInclVat returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasRefundInclVat() bool {
	if o != nil && !IsNil(o.RefundInclVat) {
		return true
	}

	return false
}

// SetRefundInclVat gets a reference to the given float32 and assigns it to the RefundInclVat field.
func (o *ChannelReturnRequest) SetRefundInclVat(v float32) {
	o.RefundInclVat = &v
}

// GetRefundExclVat returns the RefundExclVat field value if set, zero value otherwise.
func (o *ChannelReturnRequest) GetRefundExclVat() float32 {
	if o == nil || IsNil(o.RefundExclVat) {
		var ret float32
		return ret
	}
	return *o.RefundExclVat
}

// GetRefundExclVatOk returns a tuple with the RefundExclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelReturnRequest) GetRefundExclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundExclVat) {
		return nil, false
	}
	return o.RefundExclVat, true
}

// HasRefundExclVat returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasRefundExclVat() bool {
	if o != nil && !IsNil(o.RefundExclVat) {
		return true
	}

	return false
}

// SetRefundExclVat gets a reference to the given float32 and assigns it to the RefundExclVat field.
func (o *ChannelReturnRequest) SetRefundExclVat(v float32) {
	o.RefundExclVat = &v
}

// GetReturnDate returns the ReturnDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetReturnDate() time.Time {
	if o == nil || IsNil(o.ReturnDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReturnDate.Get()
}

// GetReturnDateOk returns a tuple with the ReturnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetReturnDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnDate.Get(), o.ReturnDate.IsSet()
}

// HasReturnDate returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasReturnDate() bool {
	if o != nil && o.ReturnDate.IsSet() {
		return true
	}

	return false
}

// SetReturnDate gets a reference to the given NullableTime and assigns it to the ReturnDate field.
func (o *ChannelReturnRequest) SetReturnDate(v time.Time) {
	o.ReturnDate.Set(&v)
}
// SetReturnDateNil sets the value for ReturnDate to be an explicit nil
func (o *ChannelReturnRequest) SetReturnDateNil() {
	o.ReturnDate.Set(nil)
}

// UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
func (o *ChannelReturnRequest) UnsetReturnDate() {
	o.ReturnDate.Unset()
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelReturnRequest) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelReturnRequest) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *ChannelReturnRequest) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *ChannelReturnRequest) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o ChannelReturnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelOrderNo.IsSet() {
		toSerialize["ChannelOrderNo"] = o.ChannelOrderNo.Get()
	}
	if o.MerchantOrderNo.IsSet() {
		toSerialize["MerchantOrderNo"] = o.MerchantOrderNo.Get()
	}
	toSerialize["ChannelReference"] = o.ChannelReference
	if !IsNil(o.KeyIsMerchantOrderNo) {
		toSerialize["KeyIsMerchantOrderNo"] = o.KeyIsMerchantOrderNo
	}
	if !IsNil(o.KeyIsMerchantProductNo) {
		toSerialize["KeyIsMerchantProductNo"] = o.KeyIsMerchantProductNo
	}
	toSerialize["Lines"] = o.Lines
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if o.CustomerComment.IsSet() {
		toSerialize["CustomerComment"] = o.CustomerComment.Get()
	}
	if o.MerchantComment.IsSet() {
		toSerialize["MerchantComment"] = o.MerchantComment.Get()
	}
	if !IsNil(o.RefundInclVat) {
		toSerialize["RefundInclVat"] = o.RefundInclVat
	}
	if !IsNil(o.RefundExclVat) {
		toSerialize["RefundExclVat"] = o.RefundExclVat
	}
	if o.ReturnDate.IsSet() {
		toSerialize["ReturnDate"] = o.ReturnDate.Get()
	}
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	return toSerialize, nil
}

func (o *ChannelReturnRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ChannelReference",
		"Lines",
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

	varChannelReturnRequest := _ChannelReturnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelReturnRequest)

	if err != nil {
		return err
	}

	*o = ChannelReturnRequest(varChannelReturnRequest)

	return err
}

type NullableChannelReturnRequest struct {
	value *ChannelReturnRequest
	isSet bool
}

func (v NullableChannelReturnRequest) Get() *ChannelReturnRequest {
	return v.value
}

func (v *NullableChannelReturnRequest) Set(val *ChannelReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelReturnRequest(val *ChannelReturnRequest) *NullableChannelReturnRequest {
	return &NullableChannelReturnRequest{value: val, isSet: true}
}

func (v NullableChannelReturnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


