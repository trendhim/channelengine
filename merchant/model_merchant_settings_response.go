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

// checks if the MerchantSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSettingsResponse{}

// MerchantSettingsResponse struct for MerchantSettingsResponse
type MerchantSettingsResponse struct {
	Name NullableString `json:"Name,omitempty"`
	CompanyName NullableString `json:"CompanyName,omitempty"`
	CurrencyCode NullableString `json:"CurrencyCode,omitempty"`
	DefaultCultureCode NullableString `json:"DefaultCultureCode,omitempty"`
	Settings *SettingsResponse `json:"Settings,omitempty"`
	Vat []VatSettingsResponse `json:"Vat,omitempty"`
}

// NewMerchantSettingsResponse instantiates a new MerchantSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSettingsResponse() *MerchantSettingsResponse {
	this := MerchantSettingsResponse{}
	return &this
}

// NewMerchantSettingsResponseWithDefaults instantiates a new MerchantSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSettingsResponseWithDefaults() *MerchantSettingsResponse {
	this := MerchantSettingsResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSettingsResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSettingsResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MerchantSettingsResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MerchantSettingsResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MerchantSettingsResponse) UnsetName() {
	o.Name.Unset()
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSettingsResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSettingsResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasCompanyName() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *MerchantSettingsResponse) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}
// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *MerchantSettingsResponse) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *MerchantSettingsResponse) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSettingsResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSettingsResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *MerchantSettingsResponse) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}
// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *MerchantSettingsResponse) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *MerchantSettingsResponse) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetDefaultCultureCode returns the DefaultCultureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSettingsResponse) GetDefaultCultureCode() string {
	if o == nil || IsNil(o.DefaultCultureCode.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultCultureCode.Get()
}

// GetDefaultCultureCodeOk returns a tuple with the DefaultCultureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSettingsResponse) GetDefaultCultureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultCultureCode.Get(), o.DefaultCultureCode.IsSet()
}

// HasDefaultCultureCode returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasDefaultCultureCode() bool {
	if o != nil && o.DefaultCultureCode.IsSet() {
		return true
	}

	return false
}

// SetDefaultCultureCode gets a reference to the given NullableString and assigns it to the DefaultCultureCode field.
func (o *MerchantSettingsResponse) SetDefaultCultureCode(v string) {
	o.DefaultCultureCode.Set(&v)
}
// SetDefaultCultureCodeNil sets the value for DefaultCultureCode to be an explicit nil
func (o *MerchantSettingsResponse) SetDefaultCultureCodeNil() {
	o.DefaultCultureCode.Set(nil)
}

// UnsetDefaultCultureCode ensures that no value is present for DefaultCultureCode, not even an explicit nil
func (o *MerchantSettingsResponse) UnsetDefaultCultureCode() {
	o.DefaultCultureCode.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *MerchantSettingsResponse) GetSettings() SettingsResponse {
	if o == nil || IsNil(o.Settings) {
		var ret SettingsResponse
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSettingsResponse) GetSettingsOk() (*SettingsResponse, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SettingsResponse and assigns it to the Settings field.
func (o *MerchantSettingsResponse) SetSettings(v SettingsResponse) {
	o.Settings = &v
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSettingsResponse) GetVat() []VatSettingsResponse {
	if o == nil {
		var ret []VatSettingsResponse
		return ret
	}
	return o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSettingsResponse) GetVatOk() ([]VatSettingsResponse, bool) {
	if o == nil || IsNil(o.Vat) {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *MerchantSettingsResponse) HasVat() bool {
	if o != nil && !IsNil(o.Vat) {
		return true
	}

	return false
}

// SetVat gets a reference to the given []VatSettingsResponse and assigns it to the Vat field.
func (o *MerchantSettingsResponse) SetVat(v []VatSettingsResponse) {
	o.Vat = v
}

func (o MerchantSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.CompanyName.IsSet() {
		toSerialize["CompanyName"] = o.CompanyName.Get()
	}
	if o.CurrencyCode.IsSet() {
		toSerialize["CurrencyCode"] = o.CurrencyCode.Get()
	}
	if o.DefaultCultureCode.IsSet() {
		toSerialize["DefaultCultureCode"] = o.DefaultCultureCode.Get()
	}
	if !IsNil(o.Settings) {
		toSerialize["Settings"] = o.Settings
	}
	if o.Vat != nil {
		toSerialize["Vat"] = o.Vat
	}
	return toSerialize, nil
}

type NullableMerchantSettingsResponse struct {
	value *MerchantSettingsResponse
	isSet bool
}

func (v NullableMerchantSettingsResponse) Get() *MerchantSettingsResponse {
	return v.value
}

func (v *NullableMerchantSettingsResponse) Set(val *MerchantSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSettingsResponse(val *MerchantSettingsResponse) *NullableMerchantSettingsResponse {
	return &NullableMerchantSettingsResponse{value: val, isSet: true}
}

func (v NullableMerchantSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


