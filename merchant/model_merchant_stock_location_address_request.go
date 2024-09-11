/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MerchantStockLocationAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantStockLocationAddressRequest{}

// MerchantStockLocationAddressRequest struct for MerchantStockLocationAddressRequest
type MerchantStockLocationAddressRequest struct {
	CountryIso string `json:"CountryIso"`
	StreetName NullableString `json:"StreetName,omitempty"`
	ZipCode NullableString `json:"ZipCode,omitempty"`
	HouseNr NullableString `json:"HouseNr,omitempty"`
	HouseNrAddition NullableString `json:"HouseNrAddition,omitempty"`
	City NullableString `json:"City,omitempty"`
	Region NullableString `json:"Region,omitempty"`
}

type _MerchantStockLocationAddressRequest MerchantStockLocationAddressRequest

// NewMerchantStockLocationAddressRequest instantiates a new MerchantStockLocationAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantStockLocationAddressRequest(countryIso string) *MerchantStockLocationAddressRequest {
	this := MerchantStockLocationAddressRequest{}
	this.CountryIso = countryIso
	return &this
}

// NewMerchantStockLocationAddressRequestWithDefaults instantiates a new MerchantStockLocationAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantStockLocationAddressRequestWithDefaults() *MerchantStockLocationAddressRequest {
	this := MerchantStockLocationAddressRequest{}
	return &this
}

// GetCountryIso returns the CountryIso field value
func (o *MerchantStockLocationAddressRequest) GetCountryIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryIso
}

// GetCountryIsoOk returns a tuple with the CountryIso field value
// and a boolean to check if the value has been set.
func (o *MerchantStockLocationAddressRequest) GetCountryIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryIso, true
}

// SetCountryIso sets field value
func (o *MerchantStockLocationAddressRequest) SetCountryIso(v string) {
	o.CountryIso = v
}

// GetStreetName returns the StreetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetStreetName() string {
	if o == nil || IsNil(o.StreetName.Get()) {
		var ret string
		return ret
	}
	return *o.StreetName.Get()
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetStreetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetName.Get(), o.StreetName.IsSet()
}

// HasStreetName returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasStreetName() bool {
	if o != nil && o.StreetName.IsSet() {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given NullableString and assigns it to the StreetName field.
func (o *MerchantStockLocationAddressRequest) SetStreetName(v string) {
	o.StreetName.Set(&v)
}
// SetStreetNameNil sets the value for StreetName to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetStreetNameNil() {
	o.StreetName.Set(nil)
}

// UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetStreetName() {
	o.StreetName.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *MerchantStockLocationAddressRequest) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetHouseNr returns the HouseNr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetHouseNr() string {
	if o == nil || IsNil(o.HouseNr.Get()) {
		var ret string
		return ret
	}
	return *o.HouseNr.Get()
}

// GetHouseNrOk returns a tuple with the HouseNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetHouseNrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HouseNr.Get(), o.HouseNr.IsSet()
}

// HasHouseNr returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasHouseNr() bool {
	if o != nil && o.HouseNr.IsSet() {
		return true
	}

	return false
}

// SetHouseNr gets a reference to the given NullableString and assigns it to the HouseNr field.
func (o *MerchantStockLocationAddressRequest) SetHouseNr(v string) {
	o.HouseNr.Set(&v)
}
// SetHouseNrNil sets the value for HouseNr to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetHouseNrNil() {
	o.HouseNr.Set(nil)
}

// UnsetHouseNr ensures that no value is present for HouseNr, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetHouseNr() {
	o.HouseNr.Unset()
}

// GetHouseNrAddition returns the HouseNrAddition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetHouseNrAddition() string {
	if o == nil || IsNil(o.HouseNrAddition.Get()) {
		var ret string
		return ret
	}
	return *o.HouseNrAddition.Get()
}

// GetHouseNrAdditionOk returns a tuple with the HouseNrAddition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetHouseNrAdditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HouseNrAddition.Get(), o.HouseNrAddition.IsSet()
}

// HasHouseNrAddition returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasHouseNrAddition() bool {
	if o != nil && o.HouseNrAddition.IsSet() {
		return true
	}

	return false
}

// SetHouseNrAddition gets a reference to the given NullableString and assigns it to the HouseNrAddition field.
func (o *MerchantStockLocationAddressRequest) SetHouseNrAddition(v string) {
	o.HouseNrAddition.Set(&v)
}
// SetHouseNrAdditionNil sets the value for HouseNrAddition to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetHouseNrAdditionNil() {
	o.HouseNrAddition.Set(nil)
}

// UnsetHouseNrAddition ensures that no value is present for HouseNrAddition, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetHouseNrAddition() {
	o.HouseNrAddition.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *MerchantStockLocationAddressRequest) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetCity() {
	o.City.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantStockLocationAddressRequest) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantStockLocationAddressRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *MerchantStockLocationAddressRequest) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *MerchantStockLocationAddressRequest) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *MerchantStockLocationAddressRequest) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *MerchantStockLocationAddressRequest) UnsetRegion() {
	o.Region.Unset()
}

func (o MerchantStockLocationAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantStockLocationAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CountryIso"] = o.CountryIso
	if o.StreetName.IsSet() {
		toSerialize["StreetName"] = o.StreetName.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["ZipCode"] = o.ZipCode.Get()
	}
	if o.HouseNr.IsSet() {
		toSerialize["HouseNr"] = o.HouseNr.Get()
	}
	if o.HouseNrAddition.IsSet() {
		toSerialize["HouseNrAddition"] = o.HouseNrAddition.Get()
	}
	if o.City.IsSet() {
		toSerialize["City"] = o.City.Get()
	}
	if o.Region.IsSet() {
		toSerialize["Region"] = o.Region.Get()
	}
	return toSerialize, nil
}

func (o *MerchantStockLocationAddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CountryIso",
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

	varMerchantStockLocationAddressRequest := _MerchantStockLocationAddressRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantStockLocationAddressRequest)

	if err != nil {
		return err
	}

	*o = MerchantStockLocationAddressRequest(varMerchantStockLocationAddressRequest)

	return err
}

type NullableMerchantStockLocationAddressRequest struct {
	value *MerchantStockLocationAddressRequest
	isSet bool
}

func (v NullableMerchantStockLocationAddressRequest) Get() *MerchantStockLocationAddressRequest {
	return v.value
}

func (v *NullableMerchantStockLocationAddressRequest) Set(val *MerchantStockLocationAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantStockLocationAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantStockLocationAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantStockLocationAddressRequest(val *MerchantStockLocationAddressRequest) *NullableMerchantStockLocationAddressRequest {
	return &NullableMerchantStockLocationAddressRequest{value: val, isSet: true}
}

func (v NullableMerchantStockLocationAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantStockLocationAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


