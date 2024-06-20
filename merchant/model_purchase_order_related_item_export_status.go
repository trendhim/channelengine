/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"encoding/json"
	"fmt"
)

// PurchaseOrderRelatedItemExportStatus the model 'PurchaseOrderRelatedItemExportStatus'
type PurchaseOrderRelatedItemExportStatus string

// List of PurchaseOrderRelatedItemExportStatus
const (
	PURCHASEORDERRELATEDITEMEXPORTSTATUS_NONE PurchaseOrderRelatedItemExportStatus = "NONE"
	PURCHASEORDERRELATEDITEMEXPORTSTATUS_IN_PROGRESS PurchaseOrderRelatedItemExportStatus = "IN_PROGRESS"
	PURCHASEORDERRELATEDITEMEXPORTSTATUS_SUCCESS PurchaseOrderRelatedItemExportStatus = "SUCCESS"
	PURCHASEORDERRELATEDITEMEXPORTSTATUS_FAILURE PurchaseOrderRelatedItemExportStatus = "FAILURE"
)

// All allowed values of PurchaseOrderRelatedItemExportStatus enum
var AllowedPurchaseOrderRelatedItemExportStatusEnumValues = []PurchaseOrderRelatedItemExportStatus{
	"NONE",
	"IN_PROGRESS",
	"SUCCESS",
	"FAILURE",
}

func (v *PurchaseOrderRelatedItemExportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurchaseOrderRelatedItemExportStatus(value)
	for _, existing := range AllowedPurchaseOrderRelatedItemExportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurchaseOrderRelatedItemExportStatus", value)
}

// NewPurchaseOrderRelatedItemExportStatusFromValue returns a pointer to a valid PurchaseOrderRelatedItemExportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurchaseOrderRelatedItemExportStatusFromValue(v string) (*PurchaseOrderRelatedItemExportStatus, error) {
	ev := PurchaseOrderRelatedItemExportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurchaseOrderRelatedItemExportStatus: valid values are %v", v, AllowedPurchaseOrderRelatedItemExportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurchaseOrderRelatedItemExportStatus) IsValid() bool {
	for _, existing := range AllowedPurchaseOrderRelatedItemExportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurchaseOrderRelatedItemExportStatus value
func (v PurchaseOrderRelatedItemExportStatus) Ptr() *PurchaseOrderRelatedItemExportStatus {
	return &v
}

type NullablePurchaseOrderRelatedItemExportStatus struct {
	value *PurchaseOrderRelatedItemExportStatus
	isSet bool
}

func (v NullablePurchaseOrderRelatedItemExportStatus) Get() *PurchaseOrderRelatedItemExportStatus {
	return v.value
}

func (v *NullablePurchaseOrderRelatedItemExportStatus) Set(val *PurchaseOrderRelatedItemExportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderRelatedItemExportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderRelatedItemExportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderRelatedItemExportStatus(val *PurchaseOrderRelatedItemExportStatus) *NullablePurchaseOrderRelatedItemExportStatus {
	return &NullablePurchaseOrderRelatedItemExportStatus{value: val, isSet: true}
}

func (v NullablePurchaseOrderRelatedItemExportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderRelatedItemExportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
