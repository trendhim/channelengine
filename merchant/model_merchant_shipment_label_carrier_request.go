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

// checks if the MerchantShipmentLabelCarrierRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantShipmentLabelCarrierRequest{}

// MerchantShipmentLabelCarrierRequest struct for MerchantShipmentLabelCarrierRequest
type MerchantShipmentLabelCarrierRequest struct {
	Lines []MerchantShipmentLineRequest `json:"Lines"`
	Dimensions MerchantShipmentPackageDimensionsRequest `json:"Dimensions"`
	Weight MerchantShipmentPackageWeightRequest `json:"Weight"`
}

type _MerchantShipmentLabelCarrierRequest MerchantShipmentLabelCarrierRequest

// NewMerchantShipmentLabelCarrierRequest instantiates a new MerchantShipmentLabelCarrierRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantShipmentLabelCarrierRequest(lines []MerchantShipmentLineRequest, dimensions MerchantShipmentPackageDimensionsRequest, weight MerchantShipmentPackageWeightRequest) *MerchantShipmentLabelCarrierRequest {
	this := MerchantShipmentLabelCarrierRequest{}
	this.Lines = lines
	this.Dimensions = dimensions
	this.Weight = weight
	return &this
}

// NewMerchantShipmentLabelCarrierRequestWithDefaults instantiates a new MerchantShipmentLabelCarrierRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantShipmentLabelCarrierRequestWithDefaults() *MerchantShipmentLabelCarrierRequest {
	this := MerchantShipmentLabelCarrierRequest{}
	return &this
}

// GetLines returns the Lines field value
func (o *MerchantShipmentLabelCarrierRequest) GetLines() []MerchantShipmentLineRequest {
	if o == nil {
		var ret []MerchantShipmentLineRequest
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *MerchantShipmentLabelCarrierRequest) GetLinesOk() ([]MerchantShipmentLineRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *MerchantShipmentLabelCarrierRequest) SetLines(v []MerchantShipmentLineRequest) {
	o.Lines = v
}

// GetDimensions returns the Dimensions field value
func (o *MerchantShipmentLabelCarrierRequest) GetDimensions() MerchantShipmentPackageDimensionsRequest {
	if o == nil {
		var ret MerchantShipmentPackageDimensionsRequest
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *MerchantShipmentLabelCarrierRequest) GetDimensionsOk() (*MerchantShipmentPackageDimensionsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *MerchantShipmentLabelCarrierRequest) SetDimensions(v MerchantShipmentPackageDimensionsRequest) {
	o.Dimensions = v
}

// GetWeight returns the Weight field value
func (o *MerchantShipmentLabelCarrierRequest) GetWeight() MerchantShipmentPackageWeightRequest {
	if o == nil {
		var ret MerchantShipmentPackageWeightRequest
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *MerchantShipmentLabelCarrierRequest) GetWeightOk() (*MerchantShipmentPackageWeightRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *MerchantShipmentLabelCarrierRequest) SetWeight(v MerchantShipmentPackageWeightRequest) {
	o.Weight = v
}

func (o MerchantShipmentLabelCarrierRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantShipmentLabelCarrierRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Lines"] = o.Lines
	toSerialize["Dimensions"] = o.Dimensions
	toSerialize["Weight"] = o.Weight
	return toSerialize, nil
}

func (o *MerchantShipmentLabelCarrierRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Lines",
		"Dimensions",
		"Weight",
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

	varMerchantShipmentLabelCarrierRequest := _MerchantShipmentLabelCarrierRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantShipmentLabelCarrierRequest)

	if err != nil {
		return err
	}

	*o = MerchantShipmentLabelCarrierRequest(varMerchantShipmentLabelCarrierRequest)

	return err
}

type NullableMerchantShipmentLabelCarrierRequest struct {
	value *MerchantShipmentLabelCarrierRequest
	isSet bool
}

func (v NullableMerchantShipmentLabelCarrierRequest) Get() *MerchantShipmentLabelCarrierRequest {
	return v.value
}

func (v *NullableMerchantShipmentLabelCarrierRequest) Set(val *MerchantShipmentLabelCarrierRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantShipmentLabelCarrierRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantShipmentLabelCarrierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantShipmentLabelCarrierRequest(val *MerchantShipmentLabelCarrierRequest) *NullableMerchantShipmentLabelCarrierRequest {
	return &NullableMerchantShipmentLabelCarrierRequest{value: val, isSet: true}
}

func (v NullableMerchantShipmentLabelCarrierRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantShipmentLabelCarrierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


