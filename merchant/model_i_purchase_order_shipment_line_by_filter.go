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

// checks if the IPurchaseOrderShipmentLineByFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPurchaseOrderShipmentLineByFilter{}

// IPurchaseOrderShipmentLineByFilter struct for IPurchaseOrderShipmentLineByFilter
type IPurchaseOrderShipmentLineByFilter struct {
	// ChannelEngine identifier of the shipment line
	Id NullableInt32 `json:"Id,omitempty"`
	// The number the channel uses to identify the purchase order,  which this line (partially) ships.
	ChannelOrderNo NullableString `json:"ChannelOrderNo,omitempty"`
	// Item sequence number for the item. The first item will be 001, the second 002, and so on.  This number is used as a reference to refer to this item from the carton or pallet level.
	ItemSequenceNumber NullableString `json:"ItemSequenceNumber,omitempty"`
	// The number the channel uses to identify the product
	ChannelProductNo NullableString `json:"ChannelProductNo,omitempty"`
	// The number the merchant uses to identify the product
	MerchantProductNo NullableString `json:"MerchantProductNo,omitempty"`
	QuantityUnitOfMeasure *PurchaseOrderLineUnitOfMeasure `json:"QuantityUnitOfMeasure,omitempty"`
	// The quantity
	Quantity *int32 `json:"Quantity,omitempty"`
	// The case size, in the event that we ordered using cases. Otherwise, it is 1.
	QuantityUnitSize NullableInt32 `json:"QuantityUnitSize,omitempty"`
	// The date that determines the limit of consumption or use of a product.  For perishable products.
	ExpiryDate NullableTime `json:"ExpiryDate,omitempty"`
}

// NewIPurchaseOrderShipmentLineByFilter instantiates a new IPurchaseOrderShipmentLineByFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPurchaseOrderShipmentLineByFilter() *IPurchaseOrderShipmentLineByFilter {
	this := IPurchaseOrderShipmentLineByFilter{}
	return &this
}

// NewIPurchaseOrderShipmentLineByFilterWithDefaults instantiates a new IPurchaseOrderShipmentLineByFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPurchaseOrderShipmentLineByFilterWithDefaults() *IPurchaseOrderShipmentLineByFilter {
	this := IPurchaseOrderShipmentLineByFilter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IPurchaseOrderShipmentLineByFilter) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetId() {
	o.Id.Unset()
}

// GetChannelOrderNo returns the ChannelOrderNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetChannelOrderNo() string {
	if o == nil || IsNil(o.ChannelOrderNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelOrderNo.Get()
}

// GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetChannelOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelOrderNo.Get(), o.ChannelOrderNo.IsSet()
}

// HasChannelOrderNo returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasChannelOrderNo() bool {
	if o != nil && o.ChannelOrderNo.IsSet() {
		return true
	}

	return false
}

// SetChannelOrderNo gets a reference to the given NullableString and assigns it to the ChannelOrderNo field.
func (o *IPurchaseOrderShipmentLineByFilter) SetChannelOrderNo(v string) {
	o.ChannelOrderNo.Set(&v)
}
// SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetChannelOrderNoNil() {
	o.ChannelOrderNo.Set(nil)
}

// UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetChannelOrderNo() {
	o.ChannelOrderNo.Unset()
}

// GetItemSequenceNumber returns the ItemSequenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetItemSequenceNumber() string {
	if o == nil || IsNil(o.ItemSequenceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.ItemSequenceNumber.Get()
}

// GetItemSequenceNumberOk returns a tuple with the ItemSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetItemSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemSequenceNumber.Get(), o.ItemSequenceNumber.IsSet()
}

// HasItemSequenceNumber returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasItemSequenceNumber() bool {
	if o != nil && o.ItemSequenceNumber.IsSet() {
		return true
	}

	return false
}

// SetItemSequenceNumber gets a reference to the given NullableString and assigns it to the ItemSequenceNumber field.
func (o *IPurchaseOrderShipmentLineByFilter) SetItemSequenceNumber(v string) {
	o.ItemSequenceNumber.Set(&v)
}
// SetItemSequenceNumberNil sets the value for ItemSequenceNumber to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetItemSequenceNumberNil() {
	o.ItemSequenceNumber.Set(nil)
}

// UnsetItemSequenceNumber ensures that no value is present for ItemSequenceNumber, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetItemSequenceNumber() {
	o.ItemSequenceNumber.Unset()
}

// GetChannelProductNo returns the ChannelProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetChannelProductNo() string {
	if o == nil || IsNil(o.ChannelProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelProductNo.Get()
}

// GetChannelProductNoOk returns a tuple with the ChannelProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetChannelProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProductNo.Get(), o.ChannelProductNo.IsSet()
}

// HasChannelProductNo returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasChannelProductNo() bool {
	if o != nil && o.ChannelProductNo.IsSet() {
		return true
	}

	return false
}

// SetChannelProductNo gets a reference to the given NullableString and assigns it to the ChannelProductNo field.
func (o *IPurchaseOrderShipmentLineByFilter) SetChannelProductNo(v string) {
	o.ChannelProductNo.Set(&v)
}
// SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetChannelProductNoNil() {
	o.ChannelProductNo.Set(nil)
}

// UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetChannelProductNo() {
	o.ChannelProductNo.Unset()
}

// GetMerchantProductNo returns the MerchantProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetMerchantProductNo() string {
	if o == nil || IsNil(o.MerchantProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantProductNo.Get()
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantProductNo.Get(), o.MerchantProductNo.IsSet()
}

// HasMerchantProductNo returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasMerchantProductNo() bool {
	if o != nil && o.MerchantProductNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantProductNo gets a reference to the given NullableString and assigns it to the MerchantProductNo field.
func (o *IPurchaseOrderShipmentLineByFilter) SetMerchantProductNo(v string) {
	o.MerchantProductNo.Set(&v)
}
// SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetMerchantProductNoNil() {
	o.MerchantProductNo.Set(nil)
}

// UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetMerchantProductNo() {
	o.MerchantProductNo.Unset()
}

// GetQuantityUnitOfMeasure returns the QuantityUnitOfMeasure field value if set, zero value otherwise.
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitOfMeasure() PurchaseOrderLineUnitOfMeasure {
	if o == nil || IsNil(o.QuantityUnitOfMeasure) {
		var ret PurchaseOrderLineUnitOfMeasure
		return ret
	}
	return *o.QuantityUnitOfMeasure
}

// GetQuantityUnitOfMeasureOk returns a tuple with the QuantityUnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitOfMeasureOk() (*PurchaseOrderLineUnitOfMeasure, bool) {
	if o == nil || IsNil(o.QuantityUnitOfMeasure) {
		return nil, false
	}
	return o.QuantityUnitOfMeasure, true
}

// HasQuantityUnitOfMeasure returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasQuantityUnitOfMeasure() bool {
	if o != nil && !IsNil(o.QuantityUnitOfMeasure) {
		return true
	}

	return false
}

// SetQuantityUnitOfMeasure gets a reference to the given PurchaseOrderLineUnitOfMeasure and assigns it to the QuantityUnitOfMeasure field.
func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitOfMeasure(v PurchaseOrderLineUnitOfMeasure) {
	o.QuantityUnitOfMeasure = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *IPurchaseOrderShipmentLineByFilter) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetQuantityUnitSize returns the QuantityUnitSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitSize() int32 {
	if o == nil || IsNil(o.QuantityUnitSize.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityUnitSize.Get()
}

// GetQuantityUnitSizeOk returns a tuple with the QuantityUnitSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityUnitSize.Get(), o.QuantityUnitSize.IsSet()
}

// HasQuantityUnitSize returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasQuantityUnitSize() bool {
	if o != nil && o.QuantityUnitSize.IsSet() {
		return true
	}

	return false
}

// SetQuantityUnitSize gets a reference to the given NullableInt32 and assigns it to the QuantityUnitSize field.
func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitSize(v int32) {
	o.QuantityUnitSize.Set(&v)
}
// SetQuantityUnitSizeNil sets the value for QuantityUnitSize to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitSizeNil() {
	o.QuantityUnitSize.Set(nil)
}

// UnsetQuantityUnitSize ensures that no value is present for QuantityUnitSize, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetQuantityUnitSize() {
	o.QuantityUnitSize.Unset()
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPurchaseOrderShipmentLineByFilter) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate.Get()
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPurchaseOrderShipmentLineByFilter) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryDate.Get(), o.ExpiryDate.IsSet()
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *IPurchaseOrderShipmentLineByFilter) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate.IsSet() {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given NullableTime and assigns it to the ExpiryDate field.
func (o *IPurchaseOrderShipmentLineByFilter) SetExpiryDate(v time.Time) {
	o.ExpiryDate.Set(&v)
}
// SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) SetExpiryDateNil() {
	o.ExpiryDate.Set(nil)
}

// UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
func (o *IPurchaseOrderShipmentLineByFilter) UnsetExpiryDate() {
	o.ExpiryDate.Unset()
}

func (o IPurchaseOrderShipmentLineByFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPurchaseOrderShipmentLineByFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.ChannelOrderNo.IsSet() {
		toSerialize["ChannelOrderNo"] = o.ChannelOrderNo.Get()
	}
	if o.ItemSequenceNumber.IsSet() {
		toSerialize["ItemSequenceNumber"] = o.ItemSequenceNumber.Get()
	}
	if o.ChannelProductNo.IsSet() {
		toSerialize["ChannelProductNo"] = o.ChannelProductNo.Get()
	}
	if o.MerchantProductNo.IsSet() {
		toSerialize["MerchantProductNo"] = o.MerchantProductNo.Get()
	}
	if !IsNil(o.QuantityUnitOfMeasure) {
		toSerialize["QuantityUnitOfMeasure"] = o.QuantityUnitOfMeasure
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if o.QuantityUnitSize.IsSet() {
		toSerialize["QuantityUnitSize"] = o.QuantityUnitSize.Get()
	}
	if o.ExpiryDate.IsSet() {
		toSerialize["ExpiryDate"] = o.ExpiryDate.Get()
	}
	return toSerialize, nil
}

type NullableIPurchaseOrderShipmentLineByFilter struct {
	value *IPurchaseOrderShipmentLineByFilter
	isSet bool
}

func (v NullableIPurchaseOrderShipmentLineByFilter) Get() *IPurchaseOrderShipmentLineByFilter {
	return v.value
}

func (v *NullableIPurchaseOrderShipmentLineByFilter) Set(val *IPurchaseOrderShipmentLineByFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableIPurchaseOrderShipmentLineByFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableIPurchaseOrderShipmentLineByFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPurchaseOrderShipmentLineByFilter(val *IPurchaseOrderShipmentLineByFilter) *NullableIPurchaseOrderShipmentLineByFilter {
	return &NullableIPurchaseOrderShipmentLineByFilter{value: val, isSet: true}
}

func (v NullableIPurchaseOrderShipmentLineByFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPurchaseOrderShipmentLineByFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


