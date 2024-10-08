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

// checks if the PurchaseOrderInvoiceAdditionalDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseOrderInvoiceAdditionalDetails{}

// PurchaseOrderInvoiceAdditionalDetails struct for PurchaseOrderInvoiceAdditionalDetails
type PurchaseOrderInvoiceAdditionalDetails struct {
	Type *ModulesAdditionalDetailsType `json:"Type,omitempty"`
	Detail NullableString `json:"Detail,omitempty"`
	LanguageCode NullableString `json:"LanguageCode,omitempty"`
}

// NewPurchaseOrderInvoiceAdditionalDetails instantiates a new PurchaseOrderInvoiceAdditionalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderInvoiceAdditionalDetails() *PurchaseOrderInvoiceAdditionalDetails {
	this := PurchaseOrderInvoiceAdditionalDetails{}
	return &this
}

// NewPurchaseOrderInvoiceAdditionalDetailsWithDefaults instantiates a new PurchaseOrderInvoiceAdditionalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderInvoiceAdditionalDetailsWithDefaults() *PurchaseOrderInvoiceAdditionalDetails {
	this := PurchaseOrderInvoiceAdditionalDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PurchaseOrderInvoiceAdditionalDetails) GetType() ModulesAdditionalDetailsType {
	if o == nil || IsNil(o.Type) {
		var ret ModulesAdditionalDetailsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderInvoiceAdditionalDetails) GetTypeOk() (*ModulesAdditionalDetailsType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceAdditionalDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModulesAdditionalDetailsType and assigns it to the Type field.
func (o *PurchaseOrderInvoiceAdditionalDetails) SetType(v ModulesAdditionalDetailsType) {
	o.Type = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderInvoiceAdditionalDetails) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderInvoiceAdditionalDetails) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceAdditionalDetails) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *PurchaseOrderInvoiceAdditionalDetails) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *PurchaseOrderInvoiceAdditionalDetails) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *PurchaseOrderInvoiceAdditionalDetails) UnsetDetail() {
	o.Detail.Unset()
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderInvoiceAdditionalDetails) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode.Get()) {
		var ret string
		return ret
	}
	return *o.LanguageCode.Get()
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderInvoiceAdditionalDetails) GetLanguageCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LanguageCode.Get(), o.LanguageCode.IsSet()
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceAdditionalDetails) HasLanguageCode() bool {
	if o != nil && o.LanguageCode.IsSet() {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given NullableString and assigns it to the LanguageCode field.
func (o *PurchaseOrderInvoiceAdditionalDetails) SetLanguageCode(v string) {
	o.LanguageCode.Set(&v)
}
// SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil
func (o *PurchaseOrderInvoiceAdditionalDetails) SetLanguageCodeNil() {
	o.LanguageCode.Set(nil)
}

// UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
func (o *PurchaseOrderInvoiceAdditionalDetails) UnsetLanguageCode() {
	o.LanguageCode.Unset()
}

func (o PurchaseOrderInvoiceAdditionalDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseOrderInvoiceAdditionalDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Detail.IsSet() {
		toSerialize["Detail"] = o.Detail.Get()
	}
	if o.LanguageCode.IsSet() {
		toSerialize["LanguageCode"] = o.LanguageCode.Get()
	}
	return toSerialize, nil
}

type NullablePurchaseOrderInvoiceAdditionalDetails struct {
	value *PurchaseOrderInvoiceAdditionalDetails
	isSet bool
}

func (v NullablePurchaseOrderInvoiceAdditionalDetails) Get() *PurchaseOrderInvoiceAdditionalDetails {
	return v.value
}

func (v *NullablePurchaseOrderInvoiceAdditionalDetails) Set(val *PurchaseOrderInvoiceAdditionalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderInvoiceAdditionalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderInvoiceAdditionalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderInvoiceAdditionalDetails(val *PurchaseOrderInvoiceAdditionalDetails) *NullablePurchaseOrderInvoiceAdditionalDetails {
	return &NullablePurchaseOrderInvoiceAdditionalDetails{value: val, isSet: true}
}

func (v NullablePurchaseOrderInvoiceAdditionalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderInvoiceAdditionalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


