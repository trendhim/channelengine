/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the MerchantReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantReturnRequest{}

// MerchantReturnRequest struct for MerchantReturnRequest
type MerchantReturnRequest struct {
	// The unique order reference used by the Merchant (sku).
	MerchantOrderNo string `json:"MerchantOrderNo"`
	// The unique return reference used by the Merchant (sku).
	MerchantReturnNo string `json:"MerchantReturnNo"`
	Lines []MerchantReturnLineRequest `json:"Lines"`
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

type _MerchantReturnRequest MerchantReturnRequest

// NewMerchantReturnRequest instantiates a new MerchantReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantReturnRequest(merchantOrderNo string, merchantReturnNo string, lines []MerchantReturnLineRequest) *MerchantReturnRequest {
	this := MerchantReturnRequest{}
	this.MerchantOrderNo = merchantOrderNo
	this.MerchantReturnNo = merchantReturnNo
	this.Lines = lines
	return &this
}

// NewMerchantReturnRequestWithDefaults instantiates a new MerchantReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantReturnRequestWithDefaults() *MerchantReturnRequest {
	this := MerchantReturnRequest{}
	return &this
}

// GetMerchantOrderNo returns the MerchantOrderNo field value
func (o *MerchantReturnRequest) GetMerchantOrderNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantOrderNo
}

// GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetMerchantOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantOrderNo, true
}

// SetMerchantOrderNo sets field value
func (o *MerchantReturnRequest) SetMerchantOrderNo(v string) {
	o.MerchantOrderNo = v
}

// GetMerchantReturnNo returns the MerchantReturnNo field value
func (o *MerchantReturnRequest) GetMerchantReturnNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantReturnNo
}

// GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetMerchantReturnNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantReturnNo, true
}

// SetMerchantReturnNo sets field value
func (o *MerchantReturnRequest) SetMerchantReturnNo(v string) {
	o.MerchantReturnNo = v
}

// GetLines returns the Lines field value
func (o *MerchantReturnRequest) GetLines() []MerchantReturnLineRequest {
	if o == nil {
		var ret []MerchantReturnLineRequest
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetLinesOk() ([]MerchantReturnLineRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *MerchantReturnRequest) SetLines(v []MerchantReturnLineRequest) {
	o.Lines = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MerchantReturnRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MerchantReturnRequest) SetId(v int32) {
	o.Id = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MerchantReturnRequest) GetReason() ReturnReason {
	if o == nil || IsNil(o.Reason) {
		var ret ReturnReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetReasonOk() (*ReturnReason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given ReturnReason and assigns it to the Reason field.
func (o *MerchantReturnRequest) SetReason(v ReturnReason) {
	o.Reason = &v
}

// GetCustomerComment returns the CustomerComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnRequest) GetCustomerComment() string {
	if o == nil || IsNil(o.CustomerComment.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerComment.Get()
}

// GetCustomerCommentOk returns a tuple with the CustomerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnRequest) GetCustomerCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerComment.Get(), o.CustomerComment.IsSet()
}

// HasCustomerComment returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasCustomerComment() bool {
	if o != nil && o.CustomerComment.IsSet() {
		return true
	}

	return false
}

// SetCustomerComment gets a reference to the given NullableString and assigns it to the CustomerComment field.
func (o *MerchantReturnRequest) SetCustomerComment(v string) {
	o.CustomerComment.Set(&v)
}
// SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil
func (o *MerchantReturnRequest) SetCustomerCommentNil() {
	o.CustomerComment.Set(nil)
}

// UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
func (o *MerchantReturnRequest) UnsetCustomerComment() {
	o.CustomerComment.Unset()
}

// GetMerchantComment returns the MerchantComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnRequest) GetMerchantComment() string {
	if o == nil || IsNil(o.MerchantComment.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantComment.Get()
}

// GetMerchantCommentOk returns a tuple with the MerchantComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnRequest) GetMerchantCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantComment.Get(), o.MerchantComment.IsSet()
}

// HasMerchantComment returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasMerchantComment() bool {
	if o != nil && o.MerchantComment.IsSet() {
		return true
	}

	return false
}

// SetMerchantComment gets a reference to the given NullableString and assigns it to the MerchantComment field.
func (o *MerchantReturnRequest) SetMerchantComment(v string) {
	o.MerchantComment.Set(&v)
}
// SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil
func (o *MerchantReturnRequest) SetMerchantCommentNil() {
	o.MerchantComment.Set(nil)
}

// UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
func (o *MerchantReturnRequest) UnsetMerchantComment() {
	o.MerchantComment.Unset()
}

// GetRefundInclVat returns the RefundInclVat field value if set, zero value otherwise.
func (o *MerchantReturnRequest) GetRefundInclVat() float32 {
	if o == nil || IsNil(o.RefundInclVat) {
		var ret float32
		return ret
	}
	return *o.RefundInclVat
}

// GetRefundInclVatOk returns a tuple with the RefundInclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetRefundInclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundInclVat) {
		return nil, false
	}
	return o.RefundInclVat, true
}

// HasRefundInclVat returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasRefundInclVat() bool {
	if o != nil && !IsNil(o.RefundInclVat) {
		return true
	}

	return false
}

// SetRefundInclVat gets a reference to the given float32 and assigns it to the RefundInclVat field.
func (o *MerchantReturnRequest) SetRefundInclVat(v float32) {
	o.RefundInclVat = &v
}

// GetRefundExclVat returns the RefundExclVat field value if set, zero value otherwise.
func (o *MerchantReturnRequest) GetRefundExclVat() float32 {
	if o == nil || IsNil(o.RefundExclVat) {
		var ret float32
		return ret
	}
	return *o.RefundExclVat
}

// GetRefundExclVatOk returns a tuple with the RefundExclVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantReturnRequest) GetRefundExclVatOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundExclVat) {
		return nil, false
	}
	return o.RefundExclVat, true
}

// HasRefundExclVat returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasRefundExclVat() bool {
	if o != nil && !IsNil(o.RefundExclVat) {
		return true
	}

	return false
}

// SetRefundExclVat gets a reference to the given float32 and assigns it to the RefundExclVat field.
func (o *MerchantReturnRequest) SetRefundExclVat(v float32) {
	o.RefundExclVat = &v
}

// GetReturnDate returns the ReturnDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnRequest) GetReturnDate() time.Time {
	if o == nil || IsNil(o.ReturnDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReturnDate.Get()
}

// GetReturnDateOk returns a tuple with the ReturnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnRequest) GetReturnDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnDate.Get(), o.ReturnDate.IsSet()
}

// HasReturnDate returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasReturnDate() bool {
	if o != nil && o.ReturnDate.IsSet() {
		return true
	}

	return false
}

// SetReturnDate gets a reference to the given NullableTime and assigns it to the ReturnDate field.
func (o *MerchantReturnRequest) SetReturnDate(v time.Time) {
	o.ReturnDate.Set(&v)
}
// SetReturnDateNil sets the value for ReturnDate to be an explicit nil
func (o *MerchantReturnRequest) SetReturnDateNil() {
	o.ReturnDate.Set(nil)
}

// UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
func (o *MerchantReturnRequest) UnsetReturnDate() {
	o.ReturnDate.Unset()
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantReturnRequest) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantReturnRequest) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantReturnRequest) HasExtraData() bool {
	if o != nil && IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantReturnRequest) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o MerchantReturnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MerchantOrderNo"] = o.MerchantOrderNo
	toSerialize["MerchantReturnNo"] = o.MerchantReturnNo
	toSerialize["Lines"] = o.Lines
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

func (o *MerchantReturnRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MerchantOrderNo",
		"MerchantReturnNo",
		"Lines",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantReturnRequest := _MerchantReturnRequest{}

	err = json.Unmarshal(bytes, &varMerchantReturnRequest)

	if err != nil {
		return err
	}

	*o = MerchantReturnRequest(varMerchantReturnRequest)

	return err
}

type NullableMerchantReturnRequest struct {
	value *MerchantReturnRequest
	isSet bool
}

func (v NullableMerchantReturnRequest) Get() *MerchantReturnRequest {
	return v.value
}

func (v *NullableMerchantReturnRequest) Set(val *MerchantReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantReturnRequest(val *MerchantReturnRequest) *NullableMerchantReturnRequest {
	return &NullableMerchantReturnRequest{value: val, isSet: true}
}

func (v NullableMerchantReturnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

