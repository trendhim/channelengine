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

// checks if the PurchaseOrderInvoiceChargeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseOrderInvoiceChargeDetails{}

// PurchaseOrderInvoiceChargeDetails struct for PurchaseOrderInvoiceChargeDetails
type PurchaseOrderInvoiceChargeDetails struct {
	Type *ModulesChargeDetailsType `json:"Type,omitempty"`
	Description NullableString `json:"Description,omitempty"`
	ChargeAmount NullableFloat32 `json:"ChargeAmount,omitempty"`
	ChargeAmountCurrencyCode NullableString `json:"ChargeAmountCurrencyCode,omitempty"`
	TaxDetails *PurchaseOrderInvoiceTaxDetails `json:"TaxDetails,omitempty"`
}

// NewPurchaseOrderInvoiceChargeDetails instantiates a new PurchaseOrderInvoiceChargeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderInvoiceChargeDetails() *PurchaseOrderInvoiceChargeDetails {
	this := PurchaseOrderInvoiceChargeDetails{}
	return &this
}

// NewPurchaseOrderInvoiceChargeDetailsWithDefaults instantiates a new PurchaseOrderInvoiceChargeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderInvoiceChargeDetailsWithDefaults() *PurchaseOrderInvoiceChargeDetails {
	this := PurchaseOrderInvoiceChargeDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PurchaseOrderInvoiceChargeDetails) GetType() ModulesChargeDetailsType {
	if o == nil || IsNil(o.Type) {
		var ret ModulesChargeDetailsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderInvoiceChargeDetails) GetTypeOk() (*ModulesChargeDetailsType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceChargeDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModulesChargeDetailsType and assigns it to the Type field.
func (o *PurchaseOrderInvoiceChargeDetails) SetType(v ModulesChargeDetailsType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderInvoiceChargeDetails) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderInvoiceChargeDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceChargeDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PurchaseOrderInvoiceChargeDetails) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) UnsetDescription() {
	o.Description.Unset()
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmount() float32 {
	if o == nil || IsNil(o.ChargeAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.ChargeAmount.Get()
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeAmount.Get(), o.ChargeAmount.IsSet()
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceChargeDetails) HasChargeAmount() bool {
	if o != nil && o.ChargeAmount.IsSet() {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given NullableFloat32 and assigns it to the ChargeAmount field.
func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmount(v float32) {
	o.ChargeAmount.Set(&v)
}
// SetChargeAmountNil sets the value for ChargeAmount to be an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountNil() {
	o.ChargeAmount.Set(nil)
}

// UnsetChargeAmount ensures that no value is present for ChargeAmount, not even an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) UnsetChargeAmount() {
	o.ChargeAmount.Unset()
}

// GetChargeAmountCurrencyCode returns the ChargeAmountCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountCurrencyCode() string {
	if o == nil || IsNil(o.ChargeAmountCurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.ChargeAmountCurrencyCode.Get()
}

// GetChargeAmountCurrencyCodeOk returns a tuple with the ChargeAmountCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeAmountCurrencyCode.Get(), o.ChargeAmountCurrencyCode.IsSet()
}

// HasChargeAmountCurrencyCode returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceChargeDetails) HasChargeAmountCurrencyCode() bool {
	if o != nil && o.ChargeAmountCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetChargeAmountCurrencyCode gets a reference to the given NullableString and assigns it to the ChargeAmountCurrencyCode field.
func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountCurrencyCode(v string) {
	o.ChargeAmountCurrencyCode.Set(&v)
}
// SetChargeAmountCurrencyCodeNil sets the value for ChargeAmountCurrencyCode to be an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountCurrencyCodeNil() {
	o.ChargeAmountCurrencyCode.Set(nil)
}

// UnsetChargeAmountCurrencyCode ensures that no value is present for ChargeAmountCurrencyCode, not even an explicit nil
func (o *PurchaseOrderInvoiceChargeDetails) UnsetChargeAmountCurrencyCode() {
	o.ChargeAmountCurrencyCode.Unset()
}

// GetTaxDetails returns the TaxDetails field value if set, zero value otherwise.
func (o *PurchaseOrderInvoiceChargeDetails) GetTaxDetails() PurchaseOrderInvoiceTaxDetails {
	if o == nil || IsNil(o.TaxDetails) {
		var ret PurchaseOrderInvoiceTaxDetails
		return ret
	}
	return *o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderInvoiceChargeDetails) GetTaxDetailsOk() (*PurchaseOrderInvoiceTaxDetails, bool) {
	if o == nil || IsNil(o.TaxDetails) {
		return nil, false
	}
	return o.TaxDetails, true
}

// HasTaxDetails returns a boolean if a field has been set.
func (o *PurchaseOrderInvoiceChargeDetails) HasTaxDetails() bool {
	if o != nil && !IsNil(o.TaxDetails) {
		return true
	}

	return false
}

// SetTaxDetails gets a reference to the given PurchaseOrderInvoiceTaxDetails and assigns it to the TaxDetails field.
func (o *PurchaseOrderInvoiceChargeDetails) SetTaxDetails(v PurchaseOrderInvoiceTaxDetails) {
	o.TaxDetails = &v
}

func (o PurchaseOrderInvoiceChargeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseOrderInvoiceChargeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.ChargeAmount.IsSet() {
		toSerialize["ChargeAmount"] = o.ChargeAmount.Get()
	}
	if o.ChargeAmountCurrencyCode.IsSet() {
		toSerialize["ChargeAmountCurrencyCode"] = o.ChargeAmountCurrencyCode.Get()
	}
	if !IsNil(o.TaxDetails) {
		toSerialize["TaxDetails"] = o.TaxDetails
	}
	return toSerialize, nil
}

type NullablePurchaseOrderInvoiceChargeDetails struct {
	value *PurchaseOrderInvoiceChargeDetails
	isSet bool
}

func (v NullablePurchaseOrderInvoiceChargeDetails) Get() *PurchaseOrderInvoiceChargeDetails {
	return v.value
}

func (v *NullablePurchaseOrderInvoiceChargeDetails) Set(val *PurchaseOrderInvoiceChargeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderInvoiceChargeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderInvoiceChargeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderInvoiceChargeDetails(val *PurchaseOrderInvoiceChargeDetails) *NullablePurchaseOrderInvoiceChargeDetails {
	return &NullablePurchaseOrderInvoiceChargeDetails{value: val, isSet: true}
}

func (v NullablePurchaseOrderInvoiceChargeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderInvoiceChargeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


