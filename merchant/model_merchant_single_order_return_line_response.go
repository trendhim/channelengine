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

// checks if the MerchantSingleOrderReturnLineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSingleOrderReturnLineResponse{}

// MerchantSingleOrderReturnLineResponse struct for MerchantSingleOrderReturnLineResponse
type MerchantSingleOrderReturnLineResponse struct {
	// The unique product reference used by the Merchant (sku).
	MerchantProductNo NullableString `json:"MerchantProductNo,omitempty"`
	// The accepted quantity of returned products in this orderline.
	AcceptedQuantity NullableInt32 `json:"AcceptedQuantity,omitempty"`
	// The rejected quantity of returned products in this orderline.
	RejectedQuantity NullableInt32 `json:"RejectedQuantity,omitempty"`
	OrderLine *MerchantOrderLineResponse `json:"OrderLine,omitempty"`
	ShipmentStatus *ShipmentLineStatus `json:"ShipmentStatus,omitempty"`
	// Number of items of the product in this return.
	Quantity int32 `json:"Quantity"`
	// Extra data on the returnline. Each item must have an unqiue key
	ExtraData map[string]string `json:"ExtraData,omitempty"`
}

type _MerchantSingleOrderReturnLineResponse MerchantSingleOrderReturnLineResponse

// NewMerchantSingleOrderReturnLineResponse instantiates a new MerchantSingleOrderReturnLineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSingleOrderReturnLineResponse(quantity int32) *MerchantSingleOrderReturnLineResponse {
	this := MerchantSingleOrderReturnLineResponse{}
	this.Quantity = quantity
	return &this
}

// NewMerchantSingleOrderReturnLineResponseWithDefaults instantiates a new MerchantSingleOrderReturnLineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSingleOrderReturnLineResponseWithDefaults() *MerchantSingleOrderReturnLineResponse {
	this := MerchantSingleOrderReturnLineResponse{}
	return &this
}

// GetMerchantProductNo returns the MerchantProductNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSingleOrderReturnLineResponse) GetMerchantProductNo() string {
	if o == nil || IsNil(o.MerchantProductNo.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantProductNo.Get()
}

// GetMerchantProductNoOk returns a tuple with the MerchantProductNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSingleOrderReturnLineResponse) GetMerchantProductNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantProductNo.Get(), o.MerchantProductNo.IsSet()
}

// HasMerchantProductNo returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasMerchantProductNo() bool {
	if o != nil && o.MerchantProductNo.IsSet() {
		return true
	}

	return false
}

// SetMerchantProductNo gets a reference to the given NullableString and assigns it to the MerchantProductNo field.
func (o *MerchantSingleOrderReturnLineResponse) SetMerchantProductNo(v string) {
	o.MerchantProductNo.Set(&v)
}
// SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) SetMerchantProductNoNil() {
	o.MerchantProductNo.Set(nil)
}

// UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) UnsetMerchantProductNo() {
	o.MerchantProductNo.Unset()
}

// GetAcceptedQuantity returns the AcceptedQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSingleOrderReturnLineResponse) GetAcceptedQuantity() int32 {
	if o == nil || IsNil(o.AcceptedQuantity.Get()) {
		var ret int32
		return ret
	}
	return *o.AcceptedQuantity.Get()
}

// GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSingleOrderReturnLineResponse) GetAcceptedQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedQuantity.Get(), o.AcceptedQuantity.IsSet()
}

// HasAcceptedQuantity returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasAcceptedQuantity() bool {
	if o != nil && o.AcceptedQuantity.IsSet() {
		return true
	}

	return false
}

// SetAcceptedQuantity gets a reference to the given NullableInt32 and assigns it to the AcceptedQuantity field.
func (o *MerchantSingleOrderReturnLineResponse) SetAcceptedQuantity(v int32) {
	o.AcceptedQuantity.Set(&v)
}
// SetAcceptedQuantityNil sets the value for AcceptedQuantity to be an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) SetAcceptedQuantityNil() {
	o.AcceptedQuantity.Set(nil)
}

// UnsetAcceptedQuantity ensures that no value is present for AcceptedQuantity, not even an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) UnsetAcceptedQuantity() {
	o.AcceptedQuantity.Unset()
}

// GetRejectedQuantity returns the RejectedQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSingleOrderReturnLineResponse) GetRejectedQuantity() int32 {
	if o == nil || IsNil(o.RejectedQuantity.Get()) {
		var ret int32
		return ret
	}
	return *o.RejectedQuantity.Get()
}

// GetRejectedQuantityOk returns a tuple with the RejectedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSingleOrderReturnLineResponse) GetRejectedQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RejectedQuantity.Get(), o.RejectedQuantity.IsSet()
}

// HasRejectedQuantity returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasRejectedQuantity() bool {
	if o != nil && o.RejectedQuantity.IsSet() {
		return true
	}

	return false
}

// SetRejectedQuantity gets a reference to the given NullableInt32 and assigns it to the RejectedQuantity field.
func (o *MerchantSingleOrderReturnLineResponse) SetRejectedQuantity(v int32) {
	o.RejectedQuantity.Set(&v)
}
// SetRejectedQuantityNil sets the value for RejectedQuantity to be an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) SetRejectedQuantityNil() {
	o.RejectedQuantity.Set(nil)
}

// UnsetRejectedQuantity ensures that no value is present for RejectedQuantity, not even an explicit nil
func (o *MerchantSingleOrderReturnLineResponse) UnsetRejectedQuantity() {
	o.RejectedQuantity.Unset()
}

// GetOrderLine returns the OrderLine field value if set, zero value otherwise.
func (o *MerchantSingleOrderReturnLineResponse) GetOrderLine() MerchantOrderLineResponse {
	if o == nil || IsNil(o.OrderLine) {
		var ret MerchantOrderLineResponse
		return ret
	}
	return *o.OrderLine
}

// GetOrderLineOk returns a tuple with the OrderLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSingleOrderReturnLineResponse) GetOrderLineOk() (*MerchantOrderLineResponse, bool) {
	if o == nil || IsNil(o.OrderLine) {
		return nil, false
	}
	return o.OrderLine, true
}

// HasOrderLine returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasOrderLine() bool {
	if o != nil && !IsNil(o.OrderLine) {
		return true
	}

	return false
}

// SetOrderLine gets a reference to the given MerchantOrderLineResponse and assigns it to the OrderLine field.
func (o *MerchantSingleOrderReturnLineResponse) SetOrderLine(v MerchantOrderLineResponse) {
	o.OrderLine = &v
}

// GetShipmentStatus returns the ShipmentStatus field value if set, zero value otherwise.
func (o *MerchantSingleOrderReturnLineResponse) GetShipmentStatus() ShipmentLineStatus {
	if o == nil || IsNil(o.ShipmentStatus) {
		var ret ShipmentLineStatus
		return ret
	}
	return *o.ShipmentStatus
}

// GetShipmentStatusOk returns a tuple with the ShipmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSingleOrderReturnLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool) {
	if o == nil || IsNil(o.ShipmentStatus) {
		return nil, false
	}
	return o.ShipmentStatus, true
}

// HasShipmentStatus returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasShipmentStatus() bool {
	if o != nil && !IsNil(o.ShipmentStatus) {
		return true
	}

	return false
}

// SetShipmentStatus gets a reference to the given ShipmentLineStatus and assigns it to the ShipmentStatus field.
func (o *MerchantSingleOrderReturnLineResponse) SetShipmentStatus(v ShipmentLineStatus) {
	o.ShipmentStatus = &v
}

// GetQuantity returns the Quantity field value
func (o *MerchantSingleOrderReturnLineResponse) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MerchantSingleOrderReturnLineResponse) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MerchantSingleOrderReturnLineResponse) SetQuantity(v int32) {
	o.Quantity = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantSingleOrderReturnLineResponse) GetExtraData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantSingleOrderReturnLineResponse) GetExtraDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return &o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *MerchantSingleOrderReturnLineResponse) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]string and assigns it to the ExtraData field.
func (o *MerchantSingleOrderReturnLineResponse) SetExtraData(v map[string]string) {
	o.ExtraData = v
}

func (o MerchantSingleOrderReturnLineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSingleOrderReturnLineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantProductNo.IsSet() {
		toSerialize["MerchantProductNo"] = o.MerchantProductNo.Get()
	}
	if o.AcceptedQuantity.IsSet() {
		toSerialize["AcceptedQuantity"] = o.AcceptedQuantity.Get()
	}
	if o.RejectedQuantity.IsSet() {
		toSerialize["RejectedQuantity"] = o.RejectedQuantity.Get()
	}
	if !IsNil(o.OrderLine) {
		toSerialize["OrderLine"] = o.OrderLine
	}
	if !IsNil(o.ShipmentStatus) {
		toSerialize["ShipmentStatus"] = o.ShipmentStatus
	}
	toSerialize["Quantity"] = o.Quantity
	if o.ExtraData != nil {
		toSerialize["ExtraData"] = o.ExtraData
	}
	return toSerialize, nil
}

func (o *MerchantSingleOrderReturnLineResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Quantity",
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

	varMerchantSingleOrderReturnLineResponse := _MerchantSingleOrderReturnLineResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantSingleOrderReturnLineResponse)

	if err != nil {
		return err
	}

	*o = MerchantSingleOrderReturnLineResponse(varMerchantSingleOrderReturnLineResponse)

	return err
}

type NullableMerchantSingleOrderReturnLineResponse struct {
	value *MerchantSingleOrderReturnLineResponse
	isSet bool
}

func (v NullableMerchantSingleOrderReturnLineResponse) Get() *MerchantSingleOrderReturnLineResponse {
	return v.value
}

func (v *NullableMerchantSingleOrderReturnLineResponse) Set(val *MerchantSingleOrderReturnLineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSingleOrderReturnLineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSingleOrderReturnLineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSingleOrderReturnLineResponse(val *MerchantSingleOrderReturnLineResponse) *NullableMerchantSingleOrderReturnLineResponse {
	return &NullableMerchantSingleOrderReturnLineResponse{value: val, isSet: true}
}

func (v NullableMerchantSingleOrderReturnLineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSingleOrderReturnLineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


