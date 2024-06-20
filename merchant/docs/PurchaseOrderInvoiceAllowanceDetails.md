# PurchaseOrderInvoiceAllowanceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModulesAllowanceDetailsType**](ModulesAllowanceDetailsType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ChargeAmount** | Pointer to **NullableFloat32** |  | [optional] 
**ChargeAmountCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**TaxDetails** | Pointer to [**PurchaseOrderInvoiceTaxDetails**](PurchaseOrderInvoiceTaxDetails.md) |  | [optional] 

## Methods

### NewPurchaseOrderInvoiceAllowanceDetails

`func NewPurchaseOrderInvoiceAllowanceDetails() *PurchaseOrderInvoiceAllowanceDetails`

NewPurchaseOrderInvoiceAllowanceDetails instantiates a new PurchaseOrderInvoiceAllowanceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderInvoiceAllowanceDetailsWithDefaults

`func NewPurchaseOrderInvoiceAllowanceDetailsWithDefaults() *PurchaseOrderInvoiceAllowanceDetails`

NewPurchaseOrderInvoiceAllowanceDetailsWithDefaults instantiates a new PurchaseOrderInvoiceAllowanceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetType() ModulesAllowanceDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetTypeOk() (*ModulesAllowanceDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetType(v ModulesAllowanceDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseOrderInvoiceAllowanceDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurchaseOrderInvoiceAllowanceDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PurchaseOrderInvoiceAllowanceDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChargeAmount

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetChargeAmount() float32`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetChargeAmountOk() (*float32, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetChargeAmount(v float32)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *PurchaseOrderInvoiceAllowanceDetails) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### SetChargeAmountNil

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetChargeAmountNil(b bool)`

 SetChargeAmountNil sets the value for ChargeAmount to be an explicit nil

### UnsetChargeAmount
`func (o *PurchaseOrderInvoiceAllowanceDetails) UnsetChargeAmount()`

UnsetChargeAmount ensures that no value is present for ChargeAmount, not even an explicit nil
### GetChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetChargeAmountCurrencyCode() string`

GetChargeAmountCurrencyCode returns the ChargeAmountCurrencyCode field if non-nil, zero value otherwise.

### GetChargeAmountCurrencyCodeOk

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetChargeAmountCurrencyCodeOk() (*string, bool)`

GetChargeAmountCurrencyCodeOk returns a tuple with the ChargeAmountCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetChargeAmountCurrencyCode(v string)`

SetChargeAmountCurrencyCode sets ChargeAmountCurrencyCode field to given value.

### HasChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceAllowanceDetails) HasChargeAmountCurrencyCode() bool`

HasChargeAmountCurrencyCode returns a boolean if a field has been set.

### SetChargeAmountCurrencyCodeNil

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetChargeAmountCurrencyCodeNil(b bool)`

 SetChargeAmountCurrencyCodeNil sets the value for ChargeAmountCurrencyCode to be an explicit nil

### UnsetChargeAmountCurrencyCode
`func (o *PurchaseOrderInvoiceAllowanceDetails) UnsetChargeAmountCurrencyCode()`

UnsetChargeAmountCurrencyCode ensures that no value is present for ChargeAmountCurrencyCode, not even an explicit nil
### GetTaxDetails

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetTaxDetails() PurchaseOrderInvoiceTaxDetails`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *PurchaseOrderInvoiceAllowanceDetails) GetTaxDetailsOk() (*PurchaseOrderInvoiceTaxDetails, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *PurchaseOrderInvoiceAllowanceDetails) SetTaxDetails(v PurchaseOrderInvoiceTaxDetails)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *PurchaseOrderInvoiceAllowanceDetails) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


