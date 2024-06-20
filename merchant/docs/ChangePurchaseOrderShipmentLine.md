# ChangePurchaseOrderShipmentLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOrderNo** | **string** | Channel&#39;s identifier of the purchase order | 
**MerchantProductNo** | **string** | Merchant&#39;s identifier of the product.  The combination of ChannelOrderNo + MerchantProductNo identifies the order line this shipment line  ships. | 
**QuantityUnitOfMeasure** | Pointer to [**PurchaseOrderLineUnitOfMeasure**](PurchaseOrderLineUnitOfMeasure.md) |  | [optional] 
**Quantity** | Pointer to **int32** | The quantity | [optional] 
**QuantityUnitSize** | Pointer to **NullableInt32** | The case size, when QuantityUnitOfMeasure is &#39;CASES&#39;. Otherwise, it is 1. | [optional] 
**ExpiryDate** | Pointer to **NullableTime** | The date that determines the limit of consumption or use of a product.  For perishable products. | [optional] 

## Methods

### NewChangePurchaseOrderShipmentLine

`func NewChangePurchaseOrderShipmentLine(channelOrderNo string, merchantProductNo string, ) *ChangePurchaseOrderShipmentLine`

NewChangePurchaseOrderShipmentLine instantiates a new ChangePurchaseOrderShipmentLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePurchaseOrderShipmentLineWithDefaults

`func NewChangePurchaseOrderShipmentLineWithDefaults() *ChangePurchaseOrderShipmentLine`

NewChangePurchaseOrderShipmentLineWithDefaults instantiates a new ChangePurchaseOrderShipmentLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOrderNo

`func (o *ChangePurchaseOrderShipmentLine) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChangePurchaseOrderShipmentLine) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChangePurchaseOrderShipmentLine) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.


### GetMerchantProductNo

`func (o *ChangePurchaseOrderShipmentLine) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChangePurchaseOrderShipmentLine) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChangePurchaseOrderShipmentLine) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetQuantityUnitOfMeasure

`func (o *ChangePurchaseOrderShipmentLine) GetQuantityUnitOfMeasure() PurchaseOrderLineUnitOfMeasure`

GetQuantityUnitOfMeasure returns the QuantityUnitOfMeasure field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureOk

`func (o *ChangePurchaseOrderShipmentLine) GetQuantityUnitOfMeasureOk() (*PurchaseOrderLineUnitOfMeasure, bool)`

GetQuantityUnitOfMeasureOk returns a tuple with the QuantityUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasure

`func (o *ChangePurchaseOrderShipmentLine) SetQuantityUnitOfMeasure(v PurchaseOrderLineUnitOfMeasure)`

SetQuantityUnitOfMeasure sets QuantityUnitOfMeasure field to given value.

### HasQuantityUnitOfMeasure

`func (o *ChangePurchaseOrderShipmentLine) HasQuantityUnitOfMeasure() bool`

HasQuantityUnitOfMeasure returns a boolean if a field has been set.

### GetQuantity

`func (o *ChangePurchaseOrderShipmentLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChangePurchaseOrderShipmentLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChangePurchaseOrderShipmentLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ChangePurchaseOrderShipmentLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityUnitSize

`func (o *ChangePurchaseOrderShipmentLine) GetQuantityUnitSize() int32`

GetQuantityUnitSize returns the QuantityUnitSize field if non-nil, zero value otherwise.

### GetQuantityUnitSizeOk

`func (o *ChangePurchaseOrderShipmentLine) GetQuantityUnitSizeOk() (*int32, bool)`

GetQuantityUnitSizeOk returns a tuple with the QuantityUnitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitSize

`func (o *ChangePurchaseOrderShipmentLine) SetQuantityUnitSize(v int32)`

SetQuantityUnitSize sets QuantityUnitSize field to given value.

### HasQuantityUnitSize

`func (o *ChangePurchaseOrderShipmentLine) HasQuantityUnitSize() bool`

HasQuantityUnitSize returns a boolean if a field has been set.

### SetQuantityUnitSizeNil

`func (o *ChangePurchaseOrderShipmentLine) SetQuantityUnitSizeNil(b bool)`

 SetQuantityUnitSizeNil sets the value for QuantityUnitSize to be an explicit nil

### UnsetQuantityUnitSize
`func (o *ChangePurchaseOrderShipmentLine) UnsetQuantityUnitSize()`

UnsetQuantityUnitSize ensures that no value is present for QuantityUnitSize, not even an explicit nil
### GetExpiryDate

`func (o *ChangePurchaseOrderShipmentLine) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ChangePurchaseOrderShipmentLine) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ChangePurchaseOrderShipmentLine) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ChangePurchaseOrderShipmentLine) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *ChangePurchaseOrderShipmentLine) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *ChangePurchaseOrderShipmentLine) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


