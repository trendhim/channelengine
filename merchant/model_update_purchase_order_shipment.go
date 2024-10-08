/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"time"
)

// checks if the UpdatePurchaseOrderShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePurchaseOrderShipment{}

// UpdatePurchaseOrderShipment struct for UpdatePurchaseOrderShipment
type UpdatePurchaseOrderShipment struct {
	// The number the merchant uses to identify this PO shipment
	MerchantShipmentNo NullableString `json:"MerchantShipmentNo,omitempty"`
	ShipmentType *ShipmentType `json:"ShipmentType,omitempty"`
	// When the shipment will be/was shipped
	ShippedDate *time.Time `json:"ShippedDate,omitempty"`
	// Estimated delivery time in the channel's warehouse
	EstimatedDeliveryDate *time.Time `json:"EstimatedDeliveryDate,omitempty"`
	// The merchant's identifying 'selling party number' at the channel
	SellingPartyId NullableString `json:"SellingPartyId,omitempty"`
	// The destination's 'ship to party' number at the channel
	ShipToPartyId NullableString `json:"ShipToPartyId,omitempty"`
	// Bill Of Lading (BOL) number is the unique number assigned by the vendor. The BOL present in the Shipment Confirmation message ideally matches the paper BOL provided with the shipment, but that is no must. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field.
	BillOfLadingNumber NullableString `json:"BillOfLadingNumber,omitempty"`
	ShipmentWeightUnitOfMeasure *WeightUnitOfMeasure `json:"ShipmentWeightUnitOfMeasure,omitempty"`
	// The shipment's weight
	ShipmentWeight NullableFloat32 `json:"ShipmentWeight,omitempty"`
	ShipmentVolumeUnitOfMeasure *VolumeUnitOfMeasure `json:"ShipmentVolumeUnitOfMeasure,omitempty"`
	// The shipment's volume
	ShipmentVolume NullableFloat32 `json:"ShipmentVolume,omitempty"`
	// Shipment information for each shipped product
	Lines []ChangePurchaseOrderShipmentLine `json:"Lines,omitempty"`
}

// NewUpdatePurchaseOrderShipment instantiates a new UpdatePurchaseOrderShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePurchaseOrderShipment() *UpdatePurchaseOrderShipment {
	this := UpdatePurchaseOrderShipment{}
	return &this
}

// NewUpdatePurchaseOrderShipmentWithDefaults instantiates a new UpdatePurchaseOrderShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePurchaseOrderShipmentWithDefaults() *UpdatePurchaseOrderShipment {
	this := UpdatePurchaseOrderShipment{}
	return &this
}

// GetMerchantShipmentNo returns the MerchantShipmentNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetMerchantShipmentNo() string {
	if o == nil || IsNil(o.MerchantShipmentNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantShipmentNo.Get()
}

// GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetMerchantShipmentNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantShipmentNo.Get(), o.MerchantShipmentNo.IsSet()
}

// HasMerchantShipmentNo returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasMerchantShipmentNo() bool {
	if o != nil && o.MerchantShipmentNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantShipmentNo gets a reference to the given NullableString and assigns it to the MerchantShipmentNo field.
func (o *UpdatePurchaseOrderShipment) SetMerchantShipmentNo(v string) {
	o.MerchantShipmentNo.Set(&v)
}
// SetMerchantShipmentNoNil sets the value for MerchantShipmentNo to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetMerchantShipmentNoNil() {
	o.MerchantShipmentNo.Set(nil)
}

// UnsetMerchantShipmentNo ensures that no value is present for MerchantShipmentNo, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetMerchantShipmentNo() {
	o.MerchantShipmentNo.Unset()
}

// GetShipmentType returns the ShipmentType field value if set, zero value otherwise.
func (o *UpdatePurchaseOrderShipment) GetShipmentType() ShipmentType {
	if o == nil || IsNil(o.ShipmentType) {
		var ret ShipmentType
		return ret
	}
	return *o.ShipmentType
}

// GetShipmentTypeOk returns a tuple with the ShipmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchaseOrderShipment) GetShipmentTypeOk() (*ShipmentType, bool) {
	if o == nil || IsNil(o.ShipmentType) {
		return nil, false
	}
	return o.ShipmentType, true
}

// HasShipmentType returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipmentType() bool {
	if o != nil && !IsNil(o.ShipmentType) {
		return true
	}

	return false
}

// SetShipmentType gets a reference to the given ShipmentType and assigns it to the ShipmentType field.
func (o *UpdatePurchaseOrderShipment) SetShipmentType(v ShipmentType) {
	o.ShipmentType = &v
}

// GetShippedDate returns the ShippedDate field value if set, zero value otherwise.
func (o *UpdatePurchaseOrderShipment) GetShippedDate() time.Time {
	if o == nil || IsNil(o.ShippedDate) {
		var ret time.Time
		return ret
	}
	return *o.ShippedDate
}

// GetShippedDateOk returns a tuple with the ShippedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchaseOrderShipment) GetShippedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShippedDate) {
		return nil, false
	}
	return o.ShippedDate, true
}

// HasShippedDate returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShippedDate() bool {
	if o != nil && !IsNil(o.ShippedDate) {
		return true
	}

	return false
}

// SetShippedDate gets a reference to the given time.Time and assigns it to the ShippedDate field.
func (o *UpdatePurchaseOrderShipment) SetShippedDate(v time.Time) {
	o.ShippedDate = &v
}

// GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field value if set, zero value otherwise.
func (o *UpdatePurchaseOrderShipment) GetEstimatedDeliveryDate() time.Time {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedDeliveryDate
}

// GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchaseOrderShipment) GetEstimatedDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		return nil, false
	}
	return o.EstimatedDeliveryDate, true
}

// HasEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDate gets a reference to the given time.Time and assigns it to the EstimatedDeliveryDate field.
func (o *UpdatePurchaseOrderShipment) SetEstimatedDeliveryDate(v time.Time) {
	o.EstimatedDeliveryDate = &v
}

// GetSellingPartyId returns the SellingPartyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetSellingPartyId() string {
	if o == nil || IsNil(o.SellingPartyId.Get()) {
		var ret string
		return ret
	}
	return *o.SellingPartyId.Get()
}

// GetSellingPartyIdOk returns a tuple with the SellingPartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetSellingPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SellingPartyId.Get(), o.SellingPartyId.IsSet()
}

// HasSellingPartyId returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasSellingPartyId() bool {
	if o != nil && o.SellingPartyId.IsSet() {
		return true
	}

	return false
}

// SetSellingPartyId gets a reference to the given NullableString and assigns it to the SellingPartyId field.
func (o *UpdatePurchaseOrderShipment) SetSellingPartyId(v string) {
	o.SellingPartyId.Set(&v)
}
// SetSellingPartyIdNil sets the value for SellingPartyId to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetSellingPartyIdNil() {
	o.SellingPartyId.Set(nil)
}

// UnsetSellingPartyId ensures that no value is present for SellingPartyId, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetSellingPartyId() {
	o.SellingPartyId.Unset()
}

// GetShipToPartyId returns the ShipToPartyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetShipToPartyId() string {
	if o == nil || IsNil(o.ShipToPartyId.Get()) {
		var ret string
		return ret
	}
	return *o.ShipToPartyId.Get()
}

// GetShipToPartyIdOk returns a tuple with the ShipToPartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetShipToPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipToPartyId.Get(), o.ShipToPartyId.IsSet()
}

// HasShipToPartyId returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipToPartyId() bool {
	if o != nil && o.ShipToPartyId.IsSet() {
		return true
	}

	return false
}

// SetShipToPartyId gets a reference to the given NullableString and assigns it to the ShipToPartyId field.
func (o *UpdatePurchaseOrderShipment) SetShipToPartyId(v string) {
	o.ShipToPartyId.Set(&v)
}
// SetShipToPartyIdNil sets the value for ShipToPartyId to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetShipToPartyIdNil() {
	o.ShipToPartyId.Set(nil)
}

// UnsetShipToPartyId ensures that no value is present for ShipToPartyId, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetShipToPartyId() {
	o.ShipToPartyId.Unset()
}

// GetBillOfLadingNumber returns the BillOfLadingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetBillOfLadingNumber() string {
	if o == nil || IsNil(o.BillOfLadingNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BillOfLadingNumber.Get()
}

// GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetBillOfLadingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillOfLadingNumber.Get(), o.BillOfLadingNumber.IsSet()
}

// HasBillOfLadingNumber returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasBillOfLadingNumber() bool {
	if o != nil && o.BillOfLadingNumber.IsSet() {
		return true
	}

	return false
}

// SetBillOfLadingNumber gets a reference to the given NullableString and assigns it to the BillOfLadingNumber field.
func (o *UpdatePurchaseOrderShipment) SetBillOfLadingNumber(v string) {
	o.BillOfLadingNumber.Set(&v)
}
// SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetBillOfLadingNumberNil() {
	o.BillOfLadingNumber.Set(nil)
}

// UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetBillOfLadingNumber() {
	o.BillOfLadingNumber.Unset()
}

// GetShipmentWeightUnitOfMeasure returns the ShipmentWeightUnitOfMeasure field value if set, zero value otherwise.
func (o *UpdatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasure() WeightUnitOfMeasure {
	if o == nil || IsNil(o.ShipmentWeightUnitOfMeasure) {
		var ret WeightUnitOfMeasure
		return ret
	}
	return *o.ShipmentWeightUnitOfMeasure
}

// GetShipmentWeightUnitOfMeasureOk returns a tuple with the ShipmentWeightUnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasureOk() (*WeightUnitOfMeasure, bool) {
	if o == nil || IsNil(o.ShipmentWeightUnitOfMeasure) {
		return nil, false
	}
	return o.ShipmentWeightUnitOfMeasure, true
}

// HasShipmentWeightUnitOfMeasure returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipmentWeightUnitOfMeasure() bool {
	if o != nil && !IsNil(o.ShipmentWeightUnitOfMeasure) {
		return true
	}

	return false
}

// SetShipmentWeightUnitOfMeasure gets a reference to the given WeightUnitOfMeasure and assigns it to the ShipmentWeightUnitOfMeasure field.
func (o *UpdatePurchaseOrderShipment) SetShipmentWeightUnitOfMeasure(v WeightUnitOfMeasure) {
	o.ShipmentWeightUnitOfMeasure = &v
}

// GetShipmentWeight returns the ShipmentWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetShipmentWeight() float32 {
	if o == nil || IsNil(o.ShipmentWeight.Get()) {
		var ret float32
		return ret
	}
	return *o.ShipmentWeight.Get()
}

// GetShipmentWeightOk returns a tuple with the ShipmentWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetShipmentWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipmentWeight.Get(), o.ShipmentWeight.IsSet()
}

// HasShipmentWeight returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipmentWeight() bool {
	if o != nil && o.ShipmentWeight.IsSet() {
		return true
	}

	return false
}

// SetShipmentWeight gets a reference to the given NullableFloat32 and assigns it to the ShipmentWeight field.
func (o *UpdatePurchaseOrderShipment) SetShipmentWeight(v float32) {
	o.ShipmentWeight.Set(&v)
}
// SetShipmentWeightNil sets the value for ShipmentWeight to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetShipmentWeightNil() {
	o.ShipmentWeight.Set(nil)
}

// UnsetShipmentWeight ensures that no value is present for ShipmentWeight, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetShipmentWeight() {
	o.ShipmentWeight.Unset()
}

// GetShipmentVolumeUnitOfMeasure returns the ShipmentVolumeUnitOfMeasure field value if set, zero value otherwise.
func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasure() VolumeUnitOfMeasure {
	if o == nil || IsNil(o.ShipmentVolumeUnitOfMeasure) {
		var ret VolumeUnitOfMeasure
		return ret
	}
	return *o.ShipmentVolumeUnitOfMeasure
}

// GetShipmentVolumeUnitOfMeasureOk returns a tuple with the ShipmentVolumeUnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasureOk() (*VolumeUnitOfMeasure, bool) {
	if o == nil || IsNil(o.ShipmentVolumeUnitOfMeasure) {
		return nil, false
	}
	return o.ShipmentVolumeUnitOfMeasure, true
}

// HasShipmentVolumeUnitOfMeasure returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipmentVolumeUnitOfMeasure() bool {
	if o != nil && !IsNil(o.ShipmentVolumeUnitOfMeasure) {
		return true
	}

	return false
}

// SetShipmentVolumeUnitOfMeasure gets a reference to the given VolumeUnitOfMeasure and assigns it to the ShipmentVolumeUnitOfMeasure field.
func (o *UpdatePurchaseOrderShipment) SetShipmentVolumeUnitOfMeasure(v VolumeUnitOfMeasure) {
	o.ShipmentVolumeUnitOfMeasure = &v
}

// GetShipmentVolume returns the ShipmentVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetShipmentVolume() float32 {
	if o == nil || IsNil(o.ShipmentVolume.Get()) {
		var ret float32
		return ret
	}
	return *o.ShipmentVolume.Get()
}

// GetShipmentVolumeOk returns a tuple with the ShipmentVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipmentVolume.Get(), o.ShipmentVolume.IsSet()
}

// HasShipmentVolume returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasShipmentVolume() bool {
	if o != nil && o.ShipmentVolume.IsSet() {
		return true
	}

	return false
}

// SetShipmentVolume gets a reference to the given NullableFloat32 and assigns it to the ShipmentVolume field.
func (o *UpdatePurchaseOrderShipment) SetShipmentVolume(v float32) {
	o.ShipmentVolume.Set(&v)
}
// SetShipmentVolumeNil sets the value for ShipmentVolume to be an explicit nil
func (o *UpdatePurchaseOrderShipment) SetShipmentVolumeNil() {
	o.ShipmentVolume.Set(nil)
}

// UnsetShipmentVolume ensures that no value is present for ShipmentVolume, not even an explicit nil
func (o *UpdatePurchaseOrderShipment) UnsetShipmentVolume() {
	o.ShipmentVolume.Unset()
}

// GetLines returns the Lines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePurchaseOrderShipment) GetLines() []ChangePurchaseOrderShipmentLine {
	if o == nil {
		var ret []ChangePurchaseOrderShipmentLine
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePurchaseOrderShipment) GetLinesOk() ([]ChangePurchaseOrderShipmentLine, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *UpdatePurchaseOrderShipment) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []ChangePurchaseOrderShipmentLine and assigns it to the Lines field.
func (o *UpdatePurchaseOrderShipment) SetLines(v []ChangePurchaseOrderShipmentLine) {
	o.Lines = v
}

func (o UpdatePurchaseOrderShipment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePurchaseOrderShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantShipmentNo.IsSet() {
		toSerialize["MerchantShipmentNo"] = o.MerchantShipmentNo.Get()
	}
	if !IsNil(o.ShipmentType) {
		toSerialize["ShipmentType"] = o.ShipmentType
	}
	if !IsNil(o.ShippedDate) {
		toSerialize["ShippedDate"] = o.ShippedDate
	}
	if !IsNil(o.EstimatedDeliveryDate) {
		toSerialize["EstimatedDeliveryDate"] = o.EstimatedDeliveryDate
	}
	if o.SellingPartyId.IsSet() {
		toSerialize["SellingPartyId"] = o.SellingPartyId.Get()
	}
	if o.ShipToPartyId.IsSet() {
		toSerialize["ShipToPartyId"] = o.ShipToPartyId.Get()
	}
	if o.BillOfLadingNumber.IsSet() {
		toSerialize["BillOfLadingNumber"] = o.BillOfLadingNumber.Get()
	}
	if !IsNil(o.ShipmentWeightUnitOfMeasure) {
		toSerialize["ShipmentWeightUnitOfMeasure"] = o.ShipmentWeightUnitOfMeasure
	}
	if o.ShipmentWeight.IsSet() {
		toSerialize["ShipmentWeight"] = o.ShipmentWeight.Get()
	}
	if !IsNil(o.ShipmentVolumeUnitOfMeasure) {
		toSerialize["ShipmentVolumeUnitOfMeasure"] = o.ShipmentVolumeUnitOfMeasure
	}
	if o.ShipmentVolume.IsSet() {
		toSerialize["ShipmentVolume"] = o.ShipmentVolume.Get()
	}
	if o.Lines != nil {
		toSerialize["Lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableUpdatePurchaseOrderShipment struct {
	value *UpdatePurchaseOrderShipment
	isSet bool
}

func (v NullableUpdatePurchaseOrderShipment) Get() *UpdatePurchaseOrderShipment {
	return v.value
}

func (v *NullableUpdatePurchaseOrderShipment) Set(val *UpdatePurchaseOrderShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePurchaseOrderShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePurchaseOrderShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePurchaseOrderShipment(val *UpdatePurchaseOrderShipment) *NullableUpdatePurchaseOrderShipment {
	return &NullableUpdatePurchaseOrderShipment{value: val, isSet: true}
}

func (v NullableUpdatePurchaseOrderShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePurchaseOrderShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


