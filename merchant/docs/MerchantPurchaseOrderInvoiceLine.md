# MerchantPurchaseOrderInvoiceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelPurchaseOrderNo** | Pointer to **NullableString** |  | [optional] 
**MerchantPurchaseOrderNo** | Pointer to **NullableString** |  | [optional] 
**ChannelProductNo** | Pointer to **NullableString** |  | [optional] 
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**NetCostCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**NetCostAmount** | Pointer to **float32** |  | [optional] 
**HsnCode** | Pointer to **NullableString** |  | [optional] 
**CnReferenceInvoiceNumber** | Pointer to **NullableString** |  | [optional] 
**CnDebitNoteNumber** | Pointer to **NullableString** |  | [optional] 
**CnReturnsReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**CnRmaId** | Pointer to **NullableString** |  | [optional] 
**CnGoodsReturnDate** | Pointer to **time.Time** |  | [optional] 
**CnCoopReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**CnConsignorsReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**AllowanceDetails** | Pointer to [**[]PurchaseOrderInvoiceAllowanceDetails**](PurchaseOrderInvoiceAllowanceDetails.md) |  | [optional] 
**TaxDetails** | Pointer to [**[]PurchaseOrderInvoiceTaxDetails**](PurchaseOrderInvoiceTaxDetails.md) |  | [optional] 
**ChargeDetails** | Pointer to [**[]PurchaseOrderInvoiceChargeDetails**](PurchaseOrderInvoiceChargeDetails.md) |  | [optional] 

## Methods

### NewMerchantPurchaseOrderInvoiceLine

`func NewMerchantPurchaseOrderInvoiceLine() *MerchantPurchaseOrderInvoiceLine`

NewMerchantPurchaseOrderInvoiceLine instantiates a new MerchantPurchaseOrderInvoiceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPurchaseOrderInvoiceLineWithDefaults

`func NewMerchantPurchaseOrderInvoiceLineWithDefaults() *MerchantPurchaseOrderInvoiceLine`

NewMerchantPurchaseOrderInvoiceLineWithDefaults instantiates a new MerchantPurchaseOrderInvoiceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) GetChannelPurchaseOrderNo() string`

GetChannelPurchaseOrderNo returns the ChannelPurchaseOrderNo field if non-nil, zero value otherwise.

### GetChannelPurchaseOrderNoOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetChannelPurchaseOrderNoOk() (*string, bool)`

GetChannelPurchaseOrderNoOk returns a tuple with the ChannelPurchaseOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) SetChannelPurchaseOrderNo(v string)`

SetChannelPurchaseOrderNo sets ChannelPurchaseOrderNo field to given value.

### HasChannelPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) HasChannelPurchaseOrderNo() bool`

HasChannelPurchaseOrderNo returns a boolean if a field has been set.

### SetChannelPurchaseOrderNoNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetChannelPurchaseOrderNoNil(b bool)`

 SetChannelPurchaseOrderNoNil sets the value for ChannelPurchaseOrderNo to be an explicit nil

### UnsetChannelPurchaseOrderNo
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetChannelPurchaseOrderNo()`

UnsetChannelPurchaseOrderNo ensures that no value is present for ChannelPurchaseOrderNo, not even an explicit nil
### GetMerchantPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) GetMerchantPurchaseOrderNo() string`

GetMerchantPurchaseOrderNo returns the MerchantPurchaseOrderNo field if non-nil, zero value otherwise.

### GetMerchantPurchaseOrderNoOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetMerchantPurchaseOrderNoOk() (*string, bool)`

GetMerchantPurchaseOrderNoOk returns a tuple with the MerchantPurchaseOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) SetMerchantPurchaseOrderNo(v string)`

SetMerchantPurchaseOrderNo sets MerchantPurchaseOrderNo field to given value.

### HasMerchantPurchaseOrderNo

`func (o *MerchantPurchaseOrderInvoiceLine) HasMerchantPurchaseOrderNo() bool`

HasMerchantPurchaseOrderNo returns a boolean if a field has been set.

### SetMerchantPurchaseOrderNoNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetMerchantPurchaseOrderNoNil(b bool)`

 SetMerchantPurchaseOrderNoNil sets the value for MerchantPurchaseOrderNo to be an explicit nil

### UnsetMerchantPurchaseOrderNo
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetMerchantPurchaseOrderNo()`

UnsetMerchantPurchaseOrderNo ensures that no value is present for MerchantPurchaseOrderNo, not even an explicit nil
### GetChannelProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetMerchantProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantPurchaseOrderInvoiceLine) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *MerchantPurchaseOrderInvoiceLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantPurchaseOrderInvoiceLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MerchantPurchaseOrderInvoiceLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNetCostCurrencyCode

`func (o *MerchantPurchaseOrderInvoiceLine) GetNetCostCurrencyCode() string`

GetNetCostCurrencyCode returns the NetCostCurrencyCode field if non-nil, zero value otherwise.

### GetNetCostCurrencyCodeOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetNetCostCurrencyCodeOk() (*string, bool)`

GetNetCostCurrencyCodeOk returns a tuple with the NetCostCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCostCurrencyCode

`func (o *MerchantPurchaseOrderInvoiceLine) SetNetCostCurrencyCode(v string)`

SetNetCostCurrencyCode sets NetCostCurrencyCode field to given value.

### HasNetCostCurrencyCode

`func (o *MerchantPurchaseOrderInvoiceLine) HasNetCostCurrencyCode() bool`

HasNetCostCurrencyCode returns a boolean if a field has been set.

### SetNetCostCurrencyCodeNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetNetCostCurrencyCodeNil(b bool)`

 SetNetCostCurrencyCodeNil sets the value for NetCostCurrencyCode to be an explicit nil

### UnsetNetCostCurrencyCode
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetNetCostCurrencyCode()`

UnsetNetCostCurrencyCode ensures that no value is present for NetCostCurrencyCode, not even an explicit nil
### GetNetCostAmount

`func (o *MerchantPurchaseOrderInvoiceLine) GetNetCostAmount() float32`

GetNetCostAmount returns the NetCostAmount field if non-nil, zero value otherwise.

### GetNetCostAmountOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetNetCostAmountOk() (*float32, bool)`

GetNetCostAmountOk returns a tuple with the NetCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCostAmount

`func (o *MerchantPurchaseOrderInvoiceLine) SetNetCostAmount(v float32)`

SetNetCostAmount sets NetCostAmount field to given value.

### HasNetCostAmount

`func (o *MerchantPurchaseOrderInvoiceLine) HasNetCostAmount() bool`

HasNetCostAmount returns a boolean if a field has been set.

### GetHsnCode

`func (o *MerchantPurchaseOrderInvoiceLine) GetHsnCode() string`

GetHsnCode returns the HsnCode field if non-nil, zero value otherwise.

### GetHsnCodeOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetHsnCodeOk() (*string, bool)`

GetHsnCodeOk returns a tuple with the HsnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsnCode

`func (o *MerchantPurchaseOrderInvoiceLine) SetHsnCode(v string)`

SetHsnCode sets HsnCode field to given value.

### HasHsnCode

`func (o *MerchantPurchaseOrderInvoiceLine) HasHsnCode() bool`

HasHsnCode returns a boolean if a field has been set.

### SetHsnCodeNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetHsnCodeNil(b bool)`

 SetHsnCodeNil sets the value for HsnCode to be an explicit nil

### UnsetHsnCode
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetHsnCode()`

UnsetHsnCode ensures that no value is present for HsnCode, not even an explicit nil
### GetCnReferenceInvoiceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnReferenceInvoiceNumber() string`

GetCnReferenceInvoiceNumber returns the CnReferenceInvoiceNumber field if non-nil, zero value otherwise.

### GetCnReferenceInvoiceNumberOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnReferenceInvoiceNumberOk() (*string, bool)`

GetCnReferenceInvoiceNumberOk returns a tuple with the CnReferenceInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnReferenceInvoiceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnReferenceInvoiceNumber(v string)`

SetCnReferenceInvoiceNumber sets CnReferenceInvoiceNumber field to given value.

### HasCnReferenceInvoiceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnReferenceInvoiceNumber() bool`

HasCnReferenceInvoiceNumber returns a boolean if a field has been set.

### SetCnReferenceInvoiceNumberNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnReferenceInvoiceNumberNil(b bool)`

 SetCnReferenceInvoiceNumberNil sets the value for CnReferenceInvoiceNumber to be an explicit nil

### UnsetCnReferenceInvoiceNumber
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnReferenceInvoiceNumber()`

UnsetCnReferenceInvoiceNumber ensures that no value is present for CnReferenceInvoiceNumber, not even an explicit nil
### GetCnDebitNoteNumber

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnDebitNoteNumber() string`

GetCnDebitNoteNumber returns the CnDebitNoteNumber field if non-nil, zero value otherwise.

### GetCnDebitNoteNumberOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnDebitNoteNumberOk() (*string, bool)`

GetCnDebitNoteNumberOk returns a tuple with the CnDebitNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnDebitNoteNumber

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnDebitNoteNumber(v string)`

SetCnDebitNoteNumber sets CnDebitNoteNumber field to given value.

### HasCnDebitNoteNumber

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnDebitNoteNumber() bool`

HasCnDebitNoteNumber returns a boolean if a field has been set.

### SetCnDebitNoteNumberNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnDebitNoteNumberNil(b bool)`

 SetCnDebitNoteNumberNil sets the value for CnDebitNoteNumber to be an explicit nil

### UnsetCnDebitNoteNumber
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnDebitNoteNumber()`

UnsetCnDebitNoteNumber ensures that no value is present for CnDebitNoteNumber, not even an explicit nil
### GetCnReturnsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnReturnsReferenceNumber() string`

GetCnReturnsReferenceNumber returns the CnReturnsReferenceNumber field if non-nil, zero value otherwise.

### GetCnReturnsReferenceNumberOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnReturnsReferenceNumberOk() (*string, bool)`

GetCnReturnsReferenceNumberOk returns a tuple with the CnReturnsReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnReturnsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnReturnsReferenceNumber(v string)`

SetCnReturnsReferenceNumber sets CnReturnsReferenceNumber field to given value.

### HasCnReturnsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnReturnsReferenceNumber() bool`

HasCnReturnsReferenceNumber returns a boolean if a field has been set.

### SetCnReturnsReferenceNumberNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnReturnsReferenceNumberNil(b bool)`

 SetCnReturnsReferenceNumberNil sets the value for CnReturnsReferenceNumber to be an explicit nil

### UnsetCnReturnsReferenceNumber
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnReturnsReferenceNumber()`

UnsetCnReturnsReferenceNumber ensures that no value is present for CnReturnsReferenceNumber, not even an explicit nil
### GetCnRmaId

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnRmaId() string`

GetCnRmaId returns the CnRmaId field if non-nil, zero value otherwise.

### GetCnRmaIdOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnRmaIdOk() (*string, bool)`

GetCnRmaIdOk returns a tuple with the CnRmaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnRmaId

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnRmaId(v string)`

SetCnRmaId sets CnRmaId field to given value.

### HasCnRmaId

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnRmaId() bool`

HasCnRmaId returns a boolean if a field has been set.

### SetCnRmaIdNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnRmaIdNil(b bool)`

 SetCnRmaIdNil sets the value for CnRmaId to be an explicit nil

### UnsetCnRmaId
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnRmaId()`

UnsetCnRmaId ensures that no value is present for CnRmaId, not even an explicit nil
### GetCnGoodsReturnDate

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnGoodsReturnDate() time.Time`

GetCnGoodsReturnDate returns the CnGoodsReturnDate field if non-nil, zero value otherwise.

### GetCnGoodsReturnDateOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnGoodsReturnDateOk() (*time.Time, bool)`

GetCnGoodsReturnDateOk returns a tuple with the CnGoodsReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnGoodsReturnDate

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnGoodsReturnDate(v time.Time)`

SetCnGoodsReturnDate sets CnGoodsReturnDate field to given value.

### HasCnGoodsReturnDate

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnGoodsReturnDate() bool`

HasCnGoodsReturnDate returns a boolean if a field has been set.

### GetCnCoopReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnCoopReferenceNumber() string`

GetCnCoopReferenceNumber returns the CnCoopReferenceNumber field if non-nil, zero value otherwise.

### GetCnCoopReferenceNumberOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnCoopReferenceNumberOk() (*string, bool)`

GetCnCoopReferenceNumberOk returns a tuple with the CnCoopReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnCoopReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnCoopReferenceNumber(v string)`

SetCnCoopReferenceNumber sets CnCoopReferenceNumber field to given value.

### HasCnCoopReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnCoopReferenceNumber() bool`

HasCnCoopReferenceNumber returns a boolean if a field has been set.

### SetCnCoopReferenceNumberNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnCoopReferenceNumberNil(b bool)`

 SetCnCoopReferenceNumberNil sets the value for CnCoopReferenceNumber to be an explicit nil

### UnsetCnCoopReferenceNumber
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnCoopReferenceNumber()`

UnsetCnCoopReferenceNumber ensures that no value is present for CnCoopReferenceNumber, not even an explicit nil
### GetCnConsignorsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnConsignorsReferenceNumber() string`

GetCnConsignorsReferenceNumber returns the CnConsignorsReferenceNumber field if non-nil, zero value otherwise.

### GetCnConsignorsReferenceNumberOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetCnConsignorsReferenceNumberOk() (*string, bool)`

GetCnConsignorsReferenceNumberOk returns a tuple with the CnConsignorsReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnConsignorsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnConsignorsReferenceNumber(v string)`

SetCnConsignorsReferenceNumber sets CnConsignorsReferenceNumber field to given value.

### HasCnConsignorsReferenceNumber

`func (o *MerchantPurchaseOrderInvoiceLine) HasCnConsignorsReferenceNumber() bool`

HasCnConsignorsReferenceNumber returns a boolean if a field has been set.

### SetCnConsignorsReferenceNumberNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetCnConsignorsReferenceNumberNil(b bool)`

 SetCnConsignorsReferenceNumberNil sets the value for CnConsignorsReferenceNumber to be an explicit nil

### UnsetCnConsignorsReferenceNumber
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetCnConsignorsReferenceNumber()`

UnsetCnConsignorsReferenceNumber ensures that no value is present for CnConsignorsReferenceNumber, not even an explicit nil
### GetAllowanceDetails

`func (o *MerchantPurchaseOrderInvoiceLine) GetAllowanceDetails() []PurchaseOrderInvoiceAllowanceDetails`

GetAllowanceDetails returns the AllowanceDetails field if non-nil, zero value otherwise.

### GetAllowanceDetailsOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetAllowanceDetailsOk() (*[]PurchaseOrderInvoiceAllowanceDetails, bool)`

GetAllowanceDetailsOk returns a tuple with the AllowanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceDetails

`func (o *MerchantPurchaseOrderInvoiceLine) SetAllowanceDetails(v []PurchaseOrderInvoiceAllowanceDetails)`

SetAllowanceDetails sets AllowanceDetails field to given value.

### HasAllowanceDetails

`func (o *MerchantPurchaseOrderInvoiceLine) HasAllowanceDetails() bool`

HasAllowanceDetails returns a boolean if a field has been set.

### SetAllowanceDetailsNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetAllowanceDetailsNil(b bool)`

 SetAllowanceDetailsNil sets the value for AllowanceDetails to be an explicit nil

### UnsetAllowanceDetails
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetAllowanceDetails()`

UnsetAllowanceDetails ensures that no value is present for AllowanceDetails, not even an explicit nil
### GetTaxDetails

`func (o *MerchantPurchaseOrderInvoiceLine) GetTaxDetails() []PurchaseOrderInvoiceTaxDetails`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetTaxDetailsOk() (*[]PurchaseOrderInvoiceTaxDetails, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *MerchantPurchaseOrderInvoiceLine) SetTaxDetails(v []PurchaseOrderInvoiceTaxDetails)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *MerchantPurchaseOrderInvoiceLine) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### SetTaxDetailsNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetTaxDetailsNil(b bool)`

 SetTaxDetailsNil sets the value for TaxDetails to be an explicit nil

### UnsetTaxDetails
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetTaxDetails()`

UnsetTaxDetails ensures that no value is present for TaxDetails, not even an explicit nil
### GetChargeDetails

`func (o *MerchantPurchaseOrderInvoiceLine) GetChargeDetails() []PurchaseOrderInvoiceChargeDetails`

GetChargeDetails returns the ChargeDetails field if non-nil, zero value otherwise.

### GetChargeDetailsOk

`func (o *MerchantPurchaseOrderInvoiceLine) GetChargeDetailsOk() (*[]PurchaseOrderInvoiceChargeDetails, bool)`

GetChargeDetailsOk returns a tuple with the ChargeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDetails

`func (o *MerchantPurchaseOrderInvoiceLine) SetChargeDetails(v []PurchaseOrderInvoiceChargeDetails)`

SetChargeDetails sets ChargeDetails field to given value.

### HasChargeDetails

`func (o *MerchantPurchaseOrderInvoiceLine) HasChargeDetails() bool`

HasChargeDetails returns a boolean if a field has been set.

### SetChargeDetailsNil

`func (o *MerchantPurchaseOrderInvoiceLine) SetChargeDetailsNil(b bool)`

 SetChargeDetailsNil sets the value for ChargeDetails to be an explicit nil

### UnsetChargeDetails
`func (o *MerchantPurchaseOrderInvoiceLine) UnsetChargeDetails()`

UnsetChargeDetails ensures that no value is present for ChargeDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


