# IPurchaseOrderLineByFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**ChannelOrderLineNo** | Pointer to **NullableString** |  | [optional] 
**ChannelProductNo** | Pointer to **NullableString** |  | [optional] 
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**IsBackOrderAllowed** | Pointer to **bool** |  | [optional] 
**UnitOfMeasure** | Pointer to [**PurchaseOrderLineUnitOfMeasure**](PurchaseOrderLineUnitOfMeasure.md) |  | [optional] 
**UnitSize** | Pointer to **NullableInt32** |  | [optional] 
**NetCostAmount** | Pointer to **NullableFloat32** |  | [optional] 
**NetCostCurrency** | Pointer to **NullableString** |  | [optional] 
**ListPriceAmount** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**AcknowledgedDate** | Pointer to **NullableTime** |  | [optional] 
**LineTotal** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewIPurchaseOrderLineByFilter

`func NewIPurchaseOrderLineByFilter() *IPurchaseOrderLineByFilter`

NewIPurchaseOrderLineByFilter instantiates a new IPurchaseOrderLineByFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPurchaseOrderLineByFilterWithDefaults

`func NewIPurchaseOrderLineByFilterWithDefaults() *IPurchaseOrderLineByFilter`

NewIPurchaseOrderLineByFilterWithDefaults instantiates a new IPurchaseOrderLineByFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPurchaseOrderLineByFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPurchaseOrderLineByFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPurchaseOrderLineByFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPurchaseOrderLineByFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IPurchaseOrderLineByFilter) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IPurchaseOrderLineByFilter) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetChannelOrderLineNo

`func (o *IPurchaseOrderLineByFilter) GetChannelOrderLineNo() string`

GetChannelOrderLineNo returns the ChannelOrderLineNo field if non-nil, zero value otherwise.

### GetChannelOrderLineNoOk

`func (o *IPurchaseOrderLineByFilter) GetChannelOrderLineNoOk() (*string, bool)`

GetChannelOrderLineNoOk returns a tuple with the ChannelOrderLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderLineNo

`func (o *IPurchaseOrderLineByFilter) SetChannelOrderLineNo(v string)`

SetChannelOrderLineNo sets ChannelOrderLineNo field to given value.

### HasChannelOrderLineNo

`func (o *IPurchaseOrderLineByFilter) HasChannelOrderLineNo() bool`

HasChannelOrderLineNo returns a boolean if a field has been set.

### SetChannelOrderLineNoNil

`func (o *IPurchaseOrderLineByFilter) SetChannelOrderLineNoNil(b bool)`

 SetChannelOrderLineNoNil sets the value for ChannelOrderLineNo to be an explicit nil

### UnsetChannelOrderLineNo
`func (o *IPurchaseOrderLineByFilter) UnsetChannelOrderLineNo()`

UnsetChannelOrderLineNo ensures that no value is present for ChannelOrderLineNo, not even an explicit nil
### GetChannelProductNo

`func (o *IPurchaseOrderLineByFilter) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *IPurchaseOrderLineByFilter) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *IPurchaseOrderLineByFilter) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *IPurchaseOrderLineByFilter) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *IPurchaseOrderLineByFilter) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *IPurchaseOrderLineByFilter) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetMerchantProductNo

`func (o *IPurchaseOrderLineByFilter) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *IPurchaseOrderLineByFilter) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *IPurchaseOrderLineByFilter) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *IPurchaseOrderLineByFilter) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *IPurchaseOrderLineByFilter) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *IPurchaseOrderLineByFilter) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *IPurchaseOrderLineByFilter) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IPurchaseOrderLineByFilter) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IPurchaseOrderLineByFilter) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IPurchaseOrderLineByFilter) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIsBackOrderAllowed

`func (o *IPurchaseOrderLineByFilter) GetIsBackOrderAllowed() bool`

GetIsBackOrderAllowed returns the IsBackOrderAllowed field if non-nil, zero value otherwise.

### GetIsBackOrderAllowedOk

`func (o *IPurchaseOrderLineByFilter) GetIsBackOrderAllowedOk() (*bool, bool)`

GetIsBackOrderAllowedOk returns a tuple with the IsBackOrderAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackOrderAllowed

`func (o *IPurchaseOrderLineByFilter) SetIsBackOrderAllowed(v bool)`

SetIsBackOrderAllowed sets IsBackOrderAllowed field to given value.

### HasIsBackOrderAllowed

`func (o *IPurchaseOrderLineByFilter) HasIsBackOrderAllowed() bool`

HasIsBackOrderAllowed returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *IPurchaseOrderLineByFilter) GetUnitOfMeasure() PurchaseOrderLineUnitOfMeasure`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *IPurchaseOrderLineByFilter) GetUnitOfMeasureOk() (*PurchaseOrderLineUnitOfMeasure, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *IPurchaseOrderLineByFilter) SetUnitOfMeasure(v PurchaseOrderLineUnitOfMeasure)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *IPurchaseOrderLineByFilter) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetUnitSize

`func (o *IPurchaseOrderLineByFilter) GetUnitSize() int32`

GetUnitSize returns the UnitSize field if non-nil, zero value otherwise.

### GetUnitSizeOk

`func (o *IPurchaseOrderLineByFilter) GetUnitSizeOk() (*int32, bool)`

GetUnitSizeOk returns a tuple with the UnitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSize

`func (o *IPurchaseOrderLineByFilter) SetUnitSize(v int32)`

SetUnitSize sets UnitSize field to given value.

### HasUnitSize

`func (o *IPurchaseOrderLineByFilter) HasUnitSize() bool`

HasUnitSize returns a boolean if a field has been set.

### SetUnitSizeNil

`func (o *IPurchaseOrderLineByFilter) SetUnitSizeNil(b bool)`

 SetUnitSizeNil sets the value for UnitSize to be an explicit nil

### UnsetUnitSize
`func (o *IPurchaseOrderLineByFilter) UnsetUnitSize()`

UnsetUnitSize ensures that no value is present for UnitSize, not even an explicit nil
### GetNetCostAmount

`func (o *IPurchaseOrderLineByFilter) GetNetCostAmount() float32`

GetNetCostAmount returns the NetCostAmount field if non-nil, zero value otherwise.

### GetNetCostAmountOk

`func (o *IPurchaseOrderLineByFilter) GetNetCostAmountOk() (*float32, bool)`

GetNetCostAmountOk returns a tuple with the NetCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCostAmount

`func (o *IPurchaseOrderLineByFilter) SetNetCostAmount(v float32)`

SetNetCostAmount sets NetCostAmount field to given value.

### HasNetCostAmount

`func (o *IPurchaseOrderLineByFilter) HasNetCostAmount() bool`

HasNetCostAmount returns a boolean if a field has been set.

### SetNetCostAmountNil

`func (o *IPurchaseOrderLineByFilter) SetNetCostAmountNil(b bool)`

 SetNetCostAmountNil sets the value for NetCostAmount to be an explicit nil

### UnsetNetCostAmount
`func (o *IPurchaseOrderLineByFilter) UnsetNetCostAmount()`

UnsetNetCostAmount ensures that no value is present for NetCostAmount, not even an explicit nil
### GetNetCostCurrency

`func (o *IPurchaseOrderLineByFilter) GetNetCostCurrency() string`

GetNetCostCurrency returns the NetCostCurrency field if non-nil, zero value otherwise.

### GetNetCostCurrencyOk

`func (o *IPurchaseOrderLineByFilter) GetNetCostCurrencyOk() (*string, bool)`

GetNetCostCurrencyOk returns a tuple with the NetCostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCostCurrency

`func (o *IPurchaseOrderLineByFilter) SetNetCostCurrency(v string)`

SetNetCostCurrency sets NetCostCurrency field to given value.

### HasNetCostCurrency

`func (o *IPurchaseOrderLineByFilter) HasNetCostCurrency() bool`

HasNetCostCurrency returns a boolean if a field has been set.

### SetNetCostCurrencyNil

`func (o *IPurchaseOrderLineByFilter) SetNetCostCurrencyNil(b bool)`

 SetNetCostCurrencyNil sets the value for NetCostCurrency to be an explicit nil

### UnsetNetCostCurrency
`func (o *IPurchaseOrderLineByFilter) UnsetNetCostCurrency()`

UnsetNetCostCurrency ensures that no value is present for NetCostCurrency, not even an explicit nil
### GetListPriceAmount

`func (o *IPurchaseOrderLineByFilter) GetListPriceAmount() float32`

GetListPriceAmount returns the ListPriceAmount field if non-nil, zero value otherwise.

### GetListPriceAmountOk

`func (o *IPurchaseOrderLineByFilter) GetListPriceAmountOk() (*float32, bool)`

GetListPriceAmountOk returns a tuple with the ListPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPriceAmount

`func (o *IPurchaseOrderLineByFilter) SetListPriceAmount(v float32)`

SetListPriceAmount sets ListPriceAmount field to given value.

### HasListPriceAmount

`func (o *IPurchaseOrderLineByFilter) HasListPriceAmount() bool`

HasListPriceAmount returns a boolean if a field has been set.

### SetListPriceAmountNil

`func (o *IPurchaseOrderLineByFilter) SetListPriceAmountNil(b bool)`

 SetListPriceAmountNil sets the value for ListPriceAmount to be an explicit nil

### UnsetListPriceAmount
`func (o *IPurchaseOrderLineByFilter) UnsetListPriceAmount()`

UnsetListPriceAmount ensures that no value is present for ListPriceAmount, not even an explicit nil
### GetCreatedAt

`func (o *IPurchaseOrderLineByFilter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPurchaseOrderLineByFilter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPurchaseOrderLineByFilter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPurchaseOrderLineByFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IPurchaseOrderLineByFilter) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPurchaseOrderLineByFilter) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPurchaseOrderLineByFilter) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IPurchaseOrderLineByFilter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IPurchaseOrderLineByFilter) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IPurchaseOrderLineByFilter) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAcknowledgedDate

`func (o *IPurchaseOrderLineByFilter) GetAcknowledgedDate() time.Time`

GetAcknowledgedDate returns the AcknowledgedDate field if non-nil, zero value otherwise.

### GetAcknowledgedDateOk

`func (o *IPurchaseOrderLineByFilter) GetAcknowledgedDateOk() (*time.Time, bool)`

GetAcknowledgedDateOk returns a tuple with the AcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedDate

`func (o *IPurchaseOrderLineByFilter) SetAcknowledgedDate(v time.Time)`

SetAcknowledgedDate sets AcknowledgedDate field to given value.

### HasAcknowledgedDate

`func (o *IPurchaseOrderLineByFilter) HasAcknowledgedDate() bool`

HasAcknowledgedDate returns a boolean if a field has been set.

### SetAcknowledgedDateNil

`func (o *IPurchaseOrderLineByFilter) SetAcknowledgedDateNil(b bool)`

 SetAcknowledgedDateNil sets the value for AcknowledgedDate to be an explicit nil

### UnsetAcknowledgedDate
`func (o *IPurchaseOrderLineByFilter) UnsetAcknowledgedDate()`

UnsetAcknowledgedDate ensures that no value is present for AcknowledgedDate, not even an explicit nil
### GetLineTotal

`func (o *IPurchaseOrderLineByFilter) GetLineTotal() float32`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *IPurchaseOrderLineByFilter) GetLineTotalOk() (*float32, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *IPurchaseOrderLineByFilter) SetLineTotal(v float32)`

SetLineTotal sets LineTotal field to given value.

### HasLineTotal

`func (o *IPurchaseOrderLineByFilter) HasLineTotal() bool`

HasLineTotal returns a boolean if a field has been set.

### SetLineTotalNil

`func (o *IPurchaseOrderLineByFilter) SetLineTotalNil(b bool)`

 SetLineTotalNil sets the value for LineTotal to be an explicit nil

### UnsetLineTotal
`func (o *IPurchaseOrderLineByFilter) UnsetLineTotal()`

UnsetLineTotal ensures that no value is present for LineTotal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


