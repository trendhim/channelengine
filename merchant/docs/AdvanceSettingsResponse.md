# AdvanceSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManageStock** | Pointer to **bool** |  | [optional] 
**DisableAddressValidation** | Pointer to **bool** |  | [optional] 
**SkipHouseNumberValidation** | Pointer to **bool** |  | [optional] 
**SkipHouseNumberValidationForCountryCodes** | Pointer to **[]string** |  | [optional] 
**SetOrdersToClosedOnImport** | Pointer to **bool** |  | [optional] 
**DisableStockReservation** | Pointer to **bool** |  | [optional] 
**DisableAutoOrderCancellation** | Pointer to **bool** |  | [optional] 
**AutomaticallySetPhoneNumberIfEmpty** | Pointer to **NullableString** |  | [optional] 
**DefaultVatRate** | Pointer to **float32** |  | [optional] 
**OrderTooLongOnNewHoursOffset** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdvanceSettingsResponse

`func NewAdvanceSettingsResponse() *AdvanceSettingsResponse`

NewAdvanceSettingsResponse instantiates a new AdvanceSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvanceSettingsResponseWithDefaults

`func NewAdvanceSettingsResponseWithDefaults() *AdvanceSettingsResponse`

NewAdvanceSettingsResponseWithDefaults instantiates a new AdvanceSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManageStock

`func (o *AdvanceSettingsResponse) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *AdvanceSettingsResponse) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *AdvanceSettingsResponse) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *AdvanceSettingsResponse) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetDisableAddressValidation

`func (o *AdvanceSettingsResponse) GetDisableAddressValidation() bool`

GetDisableAddressValidation returns the DisableAddressValidation field if non-nil, zero value otherwise.

### GetDisableAddressValidationOk

`func (o *AdvanceSettingsResponse) GetDisableAddressValidationOk() (*bool, bool)`

GetDisableAddressValidationOk returns a tuple with the DisableAddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAddressValidation

`func (o *AdvanceSettingsResponse) SetDisableAddressValidation(v bool)`

SetDisableAddressValidation sets DisableAddressValidation field to given value.

### HasDisableAddressValidation

`func (o *AdvanceSettingsResponse) HasDisableAddressValidation() bool`

HasDisableAddressValidation returns a boolean if a field has been set.

### GetSkipHouseNumberValidation

`func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidation() bool`

GetSkipHouseNumberValidation returns the SkipHouseNumberValidation field if non-nil, zero value otherwise.

### GetSkipHouseNumberValidationOk

`func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationOk() (*bool, bool)`

GetSkipHouseNumberValidationOk returns a tuple with the SkipHouseNumberValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHouseNumberValidation

`func (o *AdvanceSettingsResponse) SetSkipHouseNumberValidation(v bool)`

SetSkipHouseNumberValidation sets SkipHouseNumberValidation field to given value.

### HasSkipHouseNumberValidation

`func (o *AdvanceSettingsResponse) HasSkipHouseNumberValidation() bool`

HasSkipHouseNumberValidation returns a boolean if a field has been set.

### GetSkipHouseNumberValidationForCountryCodes

`func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationForCountryCodes() []string`

GetSkipHouseNumberValidationForCountryCodes returns the SkipHouseNumberValidationForCountryCodes field if non-nil, zero value otherwise.

### GetSkipHouseNumberValidationForCountryCodesOk

`func (o *AdvanceSettingsResponse) GetSkipHouseNumberValidationForCountryCodesOk() (*[]string, bool)`

GetSkipHouseNumberValidationForCountryCodesOk returns a tuple with the SkipHouseNumberValidationForCountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHouseNumberValidationForCountryCodes

`func (o *AdvanceSettingsResponse) SetSkipHouseNumberValidationForCountryCodes(v []string)`

SetSkipHouseNumberValidationForCountryCodes sets SkipHouseNumberValidationForCountryCodes field to given value.

### HasSkipHouseNumberValidationForCountryCodes

`func (o *AdvanceSettingsResponse) HasSkipHouseNumberValidationForCountryCodes() bool`

HasSkipHouseNumberValidationForCountryCodes returns a boolean if a field has been set.

### SetSkipHouseNumberValidationForCountryCodesNil

`func (o *AdvanceSettingsResponse) SetSkipHouseNumberValidationForCountryCodesNil(b bool)`

 SetSkipHouseNumberValidationForCountryCodesNil sets the value for SkipHouseNumberValidationForCountryCodes to be an explicit nil

### UnsetSkipHouseNumberValidationForCountryCodes
`func (o *AdvanceSettingsResponse) UnsetSkipHouseNumberValidationForCountryCodes()`

UnsetSkipHouseNumberValidationForCountryCodes ensures that no value is present for SkipHouseNumberValidationForCountryCodes, not even an explicit nil
### GetSetOrdersToClosedOnImport

`func (o *AdvanceSettingsResponse) GetSetOrdersToClosedOnImport() bool`

GetSetOrdersToClosedOnImport returns the SetOrdersToClosedOnImport field if non-nil, zero value otherwise.

### GetSetOrdersToClosedOnImportOk

`func (o *AdvanceSettingsResponse) GetSetOrdersToClosedOnImportOk() (*bool, bool)`

GetSetOrdersToClosedOnImportOk returns a tuple with the SetOrdersToClosedOnImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOrdersToClosedOnImport

`func (o *AdvanceSettingsResponse) SetSetOrdersToClosedOnImport(v bool)`

SetSetOrdersToClosedOnImport sets SetOrdersToClosedOnImport field to given value.

### HasSetOrdersToClosedOnImport

`func (o *AdvanceSettingsResponse) HasSetOrdersToClosedOnImport() bool`

HasSetOrdersToClosedOnImport returns a boolean if a field has been set.

### GetDisableStockReservation

`func (o *AdvanceSettingsResponse) GetDisableStockReservation() bool`

GetDisableStockReservation returns the DisableStockReservation field if non-nil, zero value otherwise.

### GetDisableStockReservationOk

`func (o *AdvanceSettingsResponse) GetDisableStockReservationOk() (*bool, bool)`

GetDisableStockReservationOk returns a tuple with the DisableStockReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStockReservation

`func (o *AdvanceSettingsResponse) SetDisableStockReservation(v bool)`

SetDisableStockReservation sets DisableStockReservation field to given value.

### HasDisableStockReservation

`func (o *AdvanceSettingsResponse) HasDisableStockReservation() bool`

HasDisableStockReservation returns a boolean if a field has been set.

### GetDisableAutoOrderCancellation

`func (o *AdvanceSettingsResponse) GetDisableAutoOrderCancellation() bool`

GetDisableAutoOrderCancellation returns the DisableAutoOrderCancellation field if non-nil, zero value otherwise.

### GetDisableAutoOrderCancellationOk

`func (o *AdvanceSettingsResponse) GetDisableAutoOrderCancellationOk() (*bool, bool)`

GetDisableAutoOrderCancellationOk returns a tuple with the DisableAutoOrderCancellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoOrderCancellation

`func (o *AdvanceSettingsResponse) SetDisableAutoOrderCancellation(v bool)`

SetDisableAutoOrderCancellation sets DisableAutoOrderCancellation field to given value.

### HasDisableAutoOrderCancellation

`func (o *AdvanceSettingsResponse) HasDisableAutoOrderCancellation() bool`

HasDisableAutoOrderCancellation returns a boolean if a field has been set.

### GetAutomaticallySetPhoneNumberIfEmpty

`func (o *AdvanceSettingsResponse) GetAutomaticallySetPhoneNumberIfEmpty() string`

GetAutomaticallySetPhoneNumberIfEmpty returns the AutomaticallySetPhoneNumberIfEmpty field if non-nil, zero value otherwise.

### GetAutomaticallySetPhoneNumberIfEmptyOk

`func (o *AdvanceSettingsResponse) GetAutomaticallySetPhoneNumberIfEmptyOk() (*string, bool)`

GetAutomaticallySetPhoneNumberIfEmptyOk returns a tuple with the AutomaticallySetPhoneNumberIfEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallySetPhoneNumberIfEmpty

`func (o *AdvanceSettingsResponse) SetAutomaticallySetPhoneNumberIfEmpty(v string)`

SetAutomaticallySetPhoneNumberIfEmpty sets AutomaticallySetPhoneNumberIfEmpty field to given value.

### HasAutomaticallySetPhoneNumberIfEmpty

`func (o *AdvanceSettingsResponse) HasAutomaticallySetPhoneNumberIfEmpty() bool`

HasAutomaticallySetPhoneNumberIfEmpty returns a boolean if a field has been set.

### SetAutomaticallySetPhoneNumberIfEmptyNil

`func (o *AdvanceSettingsResponse) SetAutomaticallySetPhoneNumberIfEmptyNil(b bool)`

 SetAutomaticallySetPhoneNumberIfEmptyNil sets the value for AutomaticallySetPhoneNumberIfEmpty to be an explicit nil

### UnsetAutomaticallySetPhoneNumberIfEmpty
`func (o *AdvanceSettingsResponse) UnsetAutomaticallySetPhoneNumberIfEmpty()`

UnsetAutomaticallySetPhoneNumberIfEmpty ensures that no value is present for AutomaticallySetPhoneNumberIfEmpty, not even an explicit nil
### GetDefaultVatRate

`func (o *AdvanceSettingsResponse) GetDefaultVatRate() float32`

GetDefaultVatRate returns the DefaultVatRate field if non-nil, zero value otherwise.

### GetDefaultVatRateOk

`func (o *AdvanceSettingsResponse) GetDefaultVatRateOk() (*float32, bool)`

GetDefaultVatRateOk returns a tuple with the DefaultVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVatRate

`func (o *AdvanceSettingsResponse) SetDefaultVatRate(v float32)`

SetDefaultVatRate sets DefaultVatRate field to given value.

### HasDefaultVatRate

`func (o *AdvanceSettingsResponse) HasDefaultVatRate() bool`

HasDefaultVatRate returns a boolean if a field has been set.

### GetOrderTooLongOnNewHoursOffset

`func (o *AdvanceSettingsResponse) GetOrderTooLongOnNewHoursOffset() int32`

GetOrderTooLongOnNewHoursOffset returns the OrderTooLongOnNewHoursOffset field if non-nil, zero value otherwise.

### GetOrderTooLongOnNewHoursOffsetOk

`func (o *AdvanceSettingsResponse) GetOrderTooLongOnNewHoursOffsetOk() (*int32, bool)`

GetOrderTooLongOnNewHoursOffsetOk returns a tuple with the OrderTooLongOnNewHoursOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTooLongOnNewHoursOffset

`func (o *AdvanceSettingsResponse) SetOrderTooLongOnNewHoursOffset(v int32)`

SetOrderTooLongOnNewHoursOffset sets OrderTooLongOnNewHoursOffset field to given value.

### HasOrderTooLongOnNewHoursOffset

`func (o *AdvanceSettingsResponse) HasOrderTooLongOnNewHoursOffset() bool`

HasOrderTooLongOnNewHoursOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


