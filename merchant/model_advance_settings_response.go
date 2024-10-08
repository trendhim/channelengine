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

// checks if the AdvanceSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvanceSettingsResponse{}

// AdvanceSettingsResponse struct for AdvanceSettingsResponse
type AdvanceSettingsResponse struct {
	ManageStock *bool `json:"ManageStock,omitempty"`
	DisableAddressValidation *bool `json:"DisableAddressValidation,omitempty"`
	SkipHouseNumberValidation *bool `json:"SkipHouseNumberValidation,omitempty"`
	SkipHouseNumberValidationForCountryCodes []string `json:"SkipHouseNumberValidationForCountryCodes,omitempty"`
	SetOrdersToClosedOnImport *bool `json:"SetOrdersToClosedOnImport,omitempty"`
	DisableStockReservation *bool `json:"DisableStockReservation,omitempty"`
	DisableAutoOrderCancellation *bool `json:"DisableAutoOrderCancellation,omitempty"`
	AutomaticallySetPhoneNumberIfEmpty NullableString `json:"AutomaticallySetPhoneNumberIfEmpty,omitempty"`
	DefaultVatRate *float32 `json:"DefaultVatRate,omitempty"`
	OrderTooLongOnNewHoursOffset *int32 `json:"OrderTooLongOnNewHoursOffset,omitempty"`
}

// NewAdvanceSettingsResponse instantiates a new AdvanceSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvanceSettingsResponse() *AdvanceSettingsResponse {
	this := AdvanceSettingsResponse{}
	return &this
}

// NewAdvanceSettingsResponseWithDefaults instantiates a new AdvanceSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvanceSettingsResponseWithDefaults() *AdvanceSettingsResponse {
	this := AdvanceSettingsResponse{}
	return &this
}

// GetManageStock returns the ManageStock field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetManageStock() bool {
	if o == nil || IsNil(o.ManageStock) {
		var ret bool
		return ret
	}
	return *o.ManageStock
}

// GetManageStockOk returns a tuple with the ManageStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetManageStockOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageStock) {
		return nil, false
	}
	return o.ManageStock, true
}

// HasManageStock returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasManageStock() bool {
	if o != nil && !IsNil(o.ManageStock) {
		return true
	}

	return false
}

// SetManageStock gets a reference to the given bool and assigns it to the ManageStock field.
func (o *AdvanceSettingsResponse) SetManageStock(v bool) {
	o.ManageStock = &v
}

// GetDisableAddressValidation returns the DisableAddressValidation field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetDisableAddressValidation() bool {
	if o == nil || IsNil(o.DisableAddressValidation) {
		var ret bool
		return ret
	}
	return *o.DisableAddressValidation
}

// GetDisableAddressValidationOk returns a tuple with the DisableAddressValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetDisableAddressValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAddressValidation) {
		return nil, false
	}
	return o.DisableAddressValidation, true
}

// HasDisableAddressValidation returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasDisableAddressValidation() bool {
	if o != nil && !IsNil(o.DisableAddressValidation) {
		return true
	}

	return false
}

// SetDisableAddressValidation gets a reference to the given bool and assigns it to the DisableAddressValidation field.
func (o *AdvanceSettingsResponse) SetDisableAddressValidation(v bool) {
	o.DisableAddressValidation = &v
}

// GetSkipHouseNumberValidation returns the SkipHouseNumberValidation field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidation() bool {
	if o == nil || IsNil(o.SkipHouseNumberValidation) {
		var ret bool
		return ret
	}
	return *o.SkipHouseNumberValidation
}

// GetSkipHouseNumberValidationOk returns a tuple with the SkipHouseNumberValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipHouseNumberValidation) {
		return nil, false
	}
	return o.SkipHouseNumberValidation, true
}

// HasSkipHouseNumberValidation returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasSkipHouseNumberValidation() bool {
	if o != nil && !IsNil(o.SkipHouseNumberValidation) {
		return true
	}

	return false
}

// SetSkipHouseNumberValidation gets a reference to the given bool and assigns it to the SkipHouseNumberValidation field.
func (o *AdvanceSettingsResponse) SetSkipHouseNumberValidation(v bool) {
	o.SkipHouseNumberValidation = &v
}

// GetSkipHouseNumberValidationForCountryCodes returns the SkipHouseNumberValidationForCountryCodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationForCountryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SkipHouseNumberValidationForCountryCodes
}

// GetSkipHouseNumberValidationForCountryCodesOk returns a tuple with the SkipHouseNumberValidationForCountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationForCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.SkipHouseNumberValidationForCountryCodes) {
		return nil, false
	}
	return o.SkipHouseNumberValidationForCountryCodes, true
}

// HasSkipHouseNumberValidationForCountryCodes returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasSkipHouseNumberValidationForCountryCodes() bool {
	if o != nil && !IsNil(o.SkipHouseNumberValidationForCountryCodes) {
		return true
	}

	return false
}

// SetSkipHouseNumberValidationForCountryCodes gets a reference to the given []string and assigns it to the SkipHouseNumberValidationForCountryCodes field.
func (o *AdvanceSettingsResponse) SetSkipHouseNumberValidationForCountryCodes(v []string) {
	o.SkipHouseNumberValidationForCountryCodes = v
}

// GetSetOrdersToClosedOnImport returns the SetOrdersToClosedOnImport field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetSetOrdersToClosedOnImport() bool {
	if o == nil || IsNil(o.SetOrdersToClosedOnImport) {
		var ret bool
		return ret
	}
	return *o.SetOrdersToClosedOnImport
}

// GetSetOrdersToClosedOnImportOk returns a tuple with the SetOrdersToClosedOnImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetSetOrdersToClosedOnImportOk() (*bool, bool) {
	if o == nil || IsNil(o.SetOrdersToClosedOnImport) {
		return nil, false
	}
	return o.SetOrdersToClosedOnImport, true
}

// HasSetOrdersToClosedOnImport returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasSetOrdersToClosedOnImport() bool {
	if o != nil && !IsNil(o.SetOrdersToClosedOnImport) {
		return true
	}

	return false
}

// SetSetOrdersToClosedOnImport gets a reference to the given bool and assigns it to the SetOrdersToClosedOnImport field.
func (o *AdvanceSettingsResponse) SetSetOrdersToClosedOnImport(v bool) {
	o.SetOrdersToClosedOnImport = &v
}

// GetDisableStockReservation returns the DisableStockReservation field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetDisableStockReservation() bool {
	if o == nil || IsNil(o.DisableStockReservation) {
		var ret bool
		return ret
	}
	return *o.DisableStockReservation
}

// GetDisableStockReservationOk returns a tuple with the DisableStockReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetDisableStockReservationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableStockReservation) {
		return nil, false
	}
	return o.DisableStockReservation, true
}

// HasDisableStockReservation returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasDisableStockReservation() bool {
	if o != nil && !IsNil(o.DisableStockReservation) {
		return true
	}

	return false
}

// SetDisableStockReservation gets a reference to the given bool and assigns it to the DisableStockReservation field.
func (o *AdvanceSettingsResponse) SetDisableStockReservation(v bool) {
	o.DisableStockReservation = &v
}

// GetDisableAutoOrderCancellation returns the DisableAutoOrderCancellation field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetDisableAutoOrderCancellation() bool {
	if o == nil || IsNil(o.DisableAutoOrderCancellation) {
		var ret bool
		return ret
	}
	return *o.DisableAutoOrderCancellation
}

// GetDisableAutoOrderCancellationOk returns a tuple with the DisableAutoOrderCancellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetDisableAutoOrderCancellationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAutoOrderCancellation) {
		return nil, false
	}
	return o.DisableAutoOrderCancellation, true
}

// HasDisableAutoOrderCancellation returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasDisableAutoOrderCancellation() bool {
	if o != nil && !IsNil(o.DisableAutoOrderCancellation) {
		return true
	}

	return false
}

// SetDisableAutoOrderCancellation gets a reference to the given bool and assigns it to the DisableAutoOrderCancellation field.
func (o *AdvanceSettingsResponse) SetDisableAutoOrderCancellation(v bool) {
	o.DisableAutoOrderCancellation = &v
}

// GetAutomaticallySetPhoneNumberIfEmpty returns the AutomaticallySetPhoneNumberIfEmpty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdvanceSettingsResponse) GetAutomaticallySetPhoneNumberIfEmpty() string {
	if o == nil || IsNil(o.AutomaticallySetPhoneNumberIfEmpty.Get()) {
		var ret string
		return ret
	}
	return *o.AutomaticallySetPhoneNumberIfEmpty.Get()
}

// GetAutomaticallySetPhoneNumberIfEmptyOk returns a tuple with the AutomaticallySetPhoneNumberIfEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdvanceSettingsResponse) GetAutomaticallySetPhoneNumberIfEmptyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomaticallySetPhoneNumberIfEmpty.Get(), o.AutomaticallySetPhoneNumberIfEmpty.IsSet()
}

// HasAutomaticallySetPhoneNumberIfEmpty returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasAutomaticallySetPhoneNumberIfEmpty() bool {
	if o != nil && o.AutomaticallySetPhoneNumberIfEmpty.IsSet() {
		return true
	}

	return false
}

// SetAutomaticallySetPhoneNumberIfEmpty gets a reference to the given NullableString and assigns it to the AutomaticallySetPhoneNumberIfEmpty field.
func (o *AdvanceSettingsResponse) SetAutomaticallySetPhoneNumberIfEmpty(v string) {
	o.AutomaticallySetPhoneNumberIfEmpty.Set(&v)
}
// SetAutomaticallySetPhoneNumberIfEmptyNil sets the value for AutomaticallySetPhoneNumberIfEmpty to be an explicit nil
func (o *AdvanceSettingsResponse) SetAutomaticallySetPhoneNumberIfEmptyNil() {
	o.AutomaticallySetPhoneNumberIfEmpty.Set(nil)
}

// UnsetAutomaticallySetPhoneNumberIfEmpty ensures that no value is present for AutomaticallySetPhoneNumberIfEmpty, not even an explicit nil
func (o *AdvanceSettingsResponse) UnsetAutomaticallySetPhoneNumberIfEmpty() {
	o.AutomaticallySetPhoneNumberIfEmpty.Unset()
}

// GetDefaultVatRate returns the DefaultVatRate field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetDefaultVatRate() float32 {
	if o == nil || IsNil(o.DefaultVatRate) {
		var ret float32
		return ret
	}
	return *o.DefaultVatRate
}

// GetDefaultVatRateOk returns a tuple with the DefaultVatRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetDefaultVatRateOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultVatRate) {
		return nil, false
	}
	return o.DefaultVatRate, true
}

// HasDefaultVatRate returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasDefaultVatRate() bool {
	if o != nil && !IsNil(o.DefaultVatRate) {
		return true
	}

	return false
}

// SetDefaultVatRate gets a reference to the given float32 and assigns it to the DefaultVatRate field.
func (o *AdvanceSettingsResponse) SetDefaultVatRate(v float32) {
	o.DefaultVatRate = &v
}

// GetOrderTooLongOnNewHoursOffset returns the OrderTooLongOnNewHoursOffset field value if set, zero value otherwise.
func (o *AdvanceSettingsResponse) GetOrderTooLongOnNewHoursOffset() int32 {
	if o == nil || IsNil(o.OrderTooLongOnNewHoursOffset) {
		var ret int32
		return ret
	}
	return *o.OrderTooLongOnNewHoursOffset
}

// GetOrderTooLongOnNewHoursOffsetOk returns a tuple with the OrderTooLongOnNewHoursOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceSettingsResponse) GetOrderTooLongOnNewHoursOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderTooLongOnNewHoursOffset) {
		return nil, false
	}
	return o.OrderTooLongOnNewHoursOffset, true
}

// HasOrderTooLongOnNewHoursOffset returns a boolean if a field has been set.
func (o *AdvanceSettingsResponse) HasOrderTooLongOnNewHoursOffset() bool {
	if o != nil && !IsNil(o.OrderTooLongOnNewHoursOffset) {
		return true
	}

	return false
}

// SetOrderTooLongOnNewHoursOffset gets a reference to the given int32 and assigns it to the OrderTooLongOnNewHoursOffset field.
func (o *AdvanceSettingsResponse) SetOrderTooLongOnNewHoursOffset(v int32) {
	o.OrderTooLongOnNewHoursOffset = &v
}

func (o AdvanceSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvanceSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManageStock) {
		toSerialize["ManageStock"] = o.ManageStock
	}
	if !IsNil(o.DisableAddressValidation) {
		toSerialize["DisableAddressValidation"] = o.DisableAddressValidation
	}
	if !IsNil(o.SkipHouseNumberValidation) {
		toSerialize["SkipHouseNumberValidation"] = o.SkipHouseNumberValidation
	}
	if o.SkipHouseNumberValidationForCountryCodes != nil {
		toSerialize["SkipHouseNumberValidationForCountryCodes"] = o.SkipHouseNumberValidationForCountryCodes
	}
	if !IsNil(o.SetOrdersToClosedOnImport) {
		toSerialize["SetOrdersToClosedOnImport"] = o.SetOrdersToClosedOnImport
	}
	if !IsNil(o.DisableStockReservation) {
		toSerialize["DisableStockReservation"] = o.DisableStockReservation
	}
	if !IsNil(o.DisableAutoOrderCancellation) {
		toSerialize["DisableAutoOrderCancellation"] = o.DisableAutoOrderCancellation
	}
	if o.AutomaticallySetPhoneNumberIfEmpty.IsSet() {
		toSerialize["AutomaticallySetPhoneNumberIfEmpty"] = o.AutomaticallySetPhoneNumberIfEmpty.Get()
	}
	if !IsNil(o.DefaultVatRate) {
		toSerialize["DefaultVatRate"] = o.DefaultVatRate
	}
	if !IsNil(o.OrderTooLongOnNewHoursOffset) {
		toSerialize["OrderTooLongOnNewHoursOffset"] = o.OrderTooLongOnNewHoursOffset
	}
	return toSerialize, nil
}

type NullableAdvanceSettingsResponse struct {
	value *AdvanceSettingsResponse
	isSet bool
}

func (v NullableAdvanceSettingsResponse) Get() *AdvanceSettingsResponse {
	return v.value
}

func (v *NullableAdvanceSettingsResponse) Set(val *AdvanceSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvanceSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvanceSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvanceSettingsResponse(val *AdvanceSettingsResponse) *NullableAdvanceSettingsResponse {
	return &NullableAdvanceSettingsResponse{value: val, isSet: true}
}

func (v NullableAdvanceSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvanceSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


