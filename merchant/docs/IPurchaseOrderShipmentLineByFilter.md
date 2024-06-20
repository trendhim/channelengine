# IPurchaseOrderShipmentLineByFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | ChannelEngine identifier of the shipment line | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** | The number the channel uses to identify the purchase order,  which this line (partially) ships. | [optional] 
**ItemSequenceNumber** | Pointer to **NullableString** | Item sequence number for the item. The first item will be 001, the second 002, and so on.  This number is used as a reference to refer to this item from the carton or pallet level. | [optional] 
**ChannelProductNo** | Pointer to **NullableString** | The number the channel uses to identify the product | [optional] 
**MerchantProductNo** | Pointer to **NullableString** | The number the merchant uses to identify the product | [optional] 
**QuantityUnitOfMeasure** | Pointer to [**PurchaseOrderLineUnitOfMeasure**](PurchaseOrderLineUnitOfMeasure.md) |  | [optional] 
**Quantity** | Pointer to **int32** | The quantity | [optional] 
**QuantityUnitSize** | Pointer to **NullableInt32** | The case size, in the event that we ordered using cases. Otherwise, it is 1. | [optional] 
**ExpiryDate** | Pointer to **NullableTime** | The date that determines the limit of consumption or use of a product.  For perishable products. | [optional] 

## Methods

### NewIPurchaseOrderShipmentLineByFilter

`func NewIPurchaseOrderShipmentLineByFilter() *IPurchaseOrderShipmentLineByFilter`

NewIPurchaseOrderShipmentLineByFilter instantiates a new IPurchaseOrderShipmentLineByFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPurchaseOrderShipmentLineByFilterWithDefaults

`func NewIPurchaseOrderShipmentLineByFilterWithDefaults() *IPurchaseOrderShipmentLineByFilter`

NewIPurchaseOrderShipmentLineByFilterWithDefaults instantiates a new IPurchaseOrderShipmentLineByFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPurchaseOrderShipmentLineByFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPurchaseOrderShipmentLineByFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPurchaseOrderShipmentLineByFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetChannelOrderNo

`func (o *IPurchaseOrderShipmentLineByFilter) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *IPurchaseOrderShipmentLineByFilter) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *IPurchaseOrderShipmentLineByFilter) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetItemSequenceNumber

`func (o *IPurchaseOrderShipmentLineByFilter) GetItemSequenceNumber() string`

GetItemSequenceNumber returns the ItemSequenceNumber field if non-nil, zero value otherwise.

### GetItemSequenceNumberOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetItemSequenceNumberOk() (*string, bool)`

GetItemSequenceNumberOk returns a tuple with the ItemSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSequenceNumber

`func (o *IPurchaseOrderShipmentLineByFilter) SetItemSequenceNumber(v string)`

SetItemSequenceNumber sets ItemSequenceNumber field to given value.

### HasItemSequenceNumber

`func (o *IPurchaseOrderShipmentLineByFilter) HasItemSequenceNumber() bool`

HasItemSequenceNumber returns a boolean if a field has been set.

### SetItemSequenceNumberNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetItemSequenceNumberNil(b bool)`

 SetItemSequenceNumberNil sets the value for ItemSequenceNumber to be an explicit nil

### UnsetItemSequenceNumber
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetItemSequenceNumber()`

UnsetItemSequenceNumber ensures that no value is present for ItemSequenceNumber, not even an explicit nil
### GetChannelProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetMerchantProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *IPurchaseOrderShipmentLineByFilter) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantityUnitOfMeasure

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitOfMeasure() PurchaseOrderLineUnitOfMeasure`

GetQuantityUnitOfMeasure returns the QuantityUnitOfMeasure field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitOfMeasureOk() (*PurchaseOrderLineUnitOfMeasure, bool)`

GetQuantityUnitOfMeasureOk returns a tuple with the QuantityUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasure

`func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitOfMeasure(v PurchaseOrderLineUnitOfMeasure)`

SetQuantityUnitOfMeasure sets QuantityUnitOfMeasure field to given value.

### HasQuantityUnitOfMeasure

`func (o *IPurchaseOrderShipmentLineByFilter) HasQuantityUnitOfMeasure() bool`

HasQuantityUnitOfMeasure returns a boolean if a field has been set.

### GetQuantity

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IPurchaseOrderShipmentLineByFilter) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IPurchaseOrderShipmentLineByFilter) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityUnitSize

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitSize() int32`

GetQuantityUnitSize returns the QuantityUnitSize field if non-nil, zero value otherwise.

### GetQuantityUnitSizeOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetQuantityUnitSizeOk() (*int32, bool)`

GetQuantityUnitSizeOk returns a tuple with the QuantityUnitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitSize

`func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitSize(v int32)`

SetQuantityUnitSize sets QuantityUnitSize field to given value.

### HasQuantityUnitSize

`func (o *IPurchaseOrderShipmentLineByFilter) HasQuantityUnitSize() bool`

HasQuantityUnitSize returns a boolean if a field has been set.

### SetQuantityUnitSizeNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetQuantityUnitSizeNil(b bool)`

 SetQuantityUnitSizeNil sets the value for QuantityUnitSize to be an explicit nil

### UnsetQuantityUnitSize
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetQuantityUnitSize()`

UnsetQuantityUnitSize ensures that no value is present for QuantityUnitSize, not even an explicit nil
### GetExpiryDate

`func (o *IPurchaseOrderShipmentLineByFilter) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IPurchaseOrderShipmentLineByFilter) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IPurchaseOrderShipmentLineByFilter) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IPurchaseOrderShipmentLineByFilter) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *IPurchaseOrderShipmentLineByFilter) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *IPurchaseOrderShipmentLineByFilter) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


