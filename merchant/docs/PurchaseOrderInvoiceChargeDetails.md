# PurchaseOrderInvoiceChargeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModulesChargeDetailsType**](ModulesChargeDetailsType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ChargeAmount** | Pointer to **NullableFloat32** |  | [optional] 
**ChargeAmountCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**TaxDetails** | Pointer to [**PurchaseOrderInvoiceTaxDetails**](PurchaseOrderInvoiceTaxDetails.md) |  | [optional] 

## Methods

### NewPurchaseOrderInvoiceChargeDetails

`func NewPurchaseOrderInvoiceChargeDetails() *PurchaseOrderInvoiceChargeDetails`

NewPurchaseOrderInvoiceChargeDetails instantiates a new PurchaseOrderInvoiceChargeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderInvoiceChargeDetailsWithDefaults

`func NewPurchaseOrderInvoiceChargeDetailsWithDefaults() *PurchaseOrderInvoiceChargeDetails`

NewPurchaseOrderInvoiceChargeDetailsWithDefaults instantiates a new PurchaseOrderInvoiceChargeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PurchaseOrderInvoiceChargeDetails) GetType() ModulesChargeDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseOrderInvoiceChargeDetails) GetTypeOk() (*ModulesChargeDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseOrderInvoiceChargeDetails) SetType(v ModulesChargeDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseOrderInvoiceChargeDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseOrderInvoiceChargeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseOrderInvoiceChargeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseOrderInvoiceChargeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurchaseOrderInvoiceChargeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PurchaseOrderInvoiceChargeDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PurchaseOrderInvoiceChargeDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChargeAmount

`func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmount() float32`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountOk() (*float32, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmount(v float32)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *PurchaseOrderInvoiceChargeDetails) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### SetChargeAmountNil

`func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountNil(b bool)`

 SetChargeAmountNil sets the value for ChargeAmount to be an explicit nil

### UnsetChargeAmount
`func (o *PurchaseOrderInvoiceChargeDetails) UnsetChargeAmount()`

UnsetChargeAmount ensures that no value is present for ChargeAmount, not even an explicit nil
### GetChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountCurrencyCode() string`

GetChargeAmountCurrencyCode returns the ChargeAmountCurrencyCode field if non-nil, zero value otherwise.

### GetChargeAmountCurrencyCodeOk

`func (o *PurchaseOrderInvoiceChargeDetails) GetChargeAmountCurrencyCodeOk() (*string, bool)`

GetChargeAmountCurrencyCodeOk returns a tuple with the ChargeAmountCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountCurrencyCode(v string)`

SetChargeAmountCurrencyCode sets ChargeAmountCurrencyCode field to given value.

### HasChargeAmountCurrencyCode

`func (o *PurchaseOrderInvoiceChargeDetails) HasChargeAmountCurrencyCode() bool`

HasChargeAmountCurrencyCode returns a boolean if a field has been set.

### SetChargeAmountCurrencyCodeNil

`func (o *PurchaseOrderInvoiceChargeDetails) SetChargeAmountCurrencyCodeNil(b bool)`

 SetChargeAmountCurrencyCodeNil sets the value for ChargeAmountCurrencyCode to be an explicit nil

### UnsetChargeAmountCurrencyCode
`func (o *PurchaseOrderInvoiceChargeDetails) UnsetChargeAmountCurrencyCode()`

UnsetChargeAmountCurrencyCode ensures that no value is present for ChargeAmountCurrencyCode, not even an explicit nil
### GetTaxDetails

`func (o *PurchaseOrderInvoiceChargeDetails) GetTaxDetails() PurchaseOrderInvoiceTaxDetails`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *PurchaseOrderInvoiceChargeDetails) GetTaxDetailsOk() (*PurchaseOrderInvoiceTaxDetails, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *PurchaseOrderInvoiceChargeDetails) SetTaxDetails(v PurchaseOrderInvoiceTaxDetails)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *PurchaseOrderInvoiceChargeDetails) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


