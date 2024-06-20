# PurchaseOrderInvoiceTaxDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxType** | Pointer to [**ModulesTaxType**](ModulesTaxType.md) |  | [optional] 
**TaxRate** | Pointer to **NullableFloat32** |  | [optional] 
**TaxAmount** | Pointer to **NullableFloat32** |  | [optional] 
**TaxAmountCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**TaxableAmount** | Pointer to **NullableFloat32** |  | [optional] 
**TaxableAmountCurrencyCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPurchaseOrderInvoiceTaxDetails

`func NewPurchaseOrderInvoiceTaxDetails() *PurchaseOrderInvoiceTaxDetails`

NewPurchaseOrderInvoiceTaxDetails instantiates a new PurchaseOrderInvoiceTaxDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderInvoiceTaxDetailsWithDefaults

`func NewPurchaseOrderInvoiceTaxDetailsWithDefaults() *PurchaseOrderInvoiceTaxDetails`

NewPurchaseOrderInvoiceTaxDetailsWithDefaults instantiates a new PurchaseOrderInvoiceTaxDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxType

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxType() ModulesTaxType`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxTypeOk() (*ModulesTaxType, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxType(v ModulesTaxType)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetTaxRate

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *PurchaseOrderInvoiceTaxDetails) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxAmount

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *PurchaseOrderInvoiceTaxDetails) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetTaxAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxAmountCurrencyCode() string`

GetTaxAmountCurrencyCode returns the TaxAmountCurrencyCode field if non-nil, zero value otherwise.

### GetTaxAmountCurrencyCodeOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxAmountCurrencyCodeOk() (*string, bool)`

GetTaxAmountCurrencyCodeOk returns a tuple with the TaxAmountCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxAmountCurrencyCode(v string)`

SetTaxAmountCurrencyCode sets TaxAmountCurrencyCode field to given value.

### HasTaxAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxAmountCurrencyCode() bool`

HasTaxAmountCurrencyCode returns a boolean if a field has been set.

### SetTaxAmountCurrencyCodeNil

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxAmountCurrencyCodeNil(b bool)`

 SetTaxAmountCurrencyCodeNil sets the value for TaxAmountCurrencyCode to be an explicit nil

### UnsetTaxAmountCurrencyCode
`func (o *PurchaseOrderInvoiceTaxDetails) UnsetTaxAmountCurrencyCode()`

UnsetTaxAmountCurrencyCode ensures that no value is present for TaxAmountCurrencyCode, not even an explicit nil
### GetTaxableAmount

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxableAmount() float32`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxableAmountOk() (*float32, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxableAmount(v float32)`

SetTaxableAmount sets TaxableAmount field to given value.

### HasTaxableAmount

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxableAmount() bool`

HasTaxableAmount returns a boolean if a field has been set.

### SetTaxableAmountNil

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxableAmountNil(b bool)`

 SetTaxableAmountNil sets the value for TaxableAmount to be an explicit nil

### UnsetTaxableAmount
`func (o *PurchaseOrderInvoiceTaxDetails) UnsetTaxableAmount()`

UnsetTaxableAmount ensures that no value is present for TaxableAmount, not even an explicit nil
### GetTaxableAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxableAmountCurrencyCode() string`

GetTaxableAmountCurrencyCode returns the TaxableAmountCurrencyCode field if non-nil, zero value otherwise.

### GetTaxableAmountCurrencyCodeOk

`func (o *PurchaseOrderInvoiceTaxDetails) GetTaxableAmountCurrencyCodeOk() (*string, bool)`

GetTaxableAmountCurrencyCodeOk returns a tuple with the TaxableAmountCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxableAmountCurrencyCode(v string)`

SetTaxableAmountCurrencyCode sets TaxableAmountCurrencyCode field to given value.

### HasTaxableAmountCurrencyCode

`func (o *PurchaseOrderInvoiceTaxDetails) HasTaxableAmountCurrencyCode() bool`

HasTaxableAmountCurrencyCode returns a boolean if a field has been set.

### SetTaxableAmountCurrencyCodeNil

`func (o *PurchaseOrderInvoiceTaxDetails) SetTaxableAmountCurrencyCodeNil(b bool)`

 SetTaxableAmountCurrencyCodeNil sets the value for TaxableAmountCurrencyCode to be an explicit nil

### UnsetTaxableAmountCurrencyCode
`func (o *PurchaseOrderInvoiceTaxDetails) UnsetTaxableAmountCurrencyCode()`

UnsetTaxableAmountCurrencyCode ensures that no value is present for TaxableAmountCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


