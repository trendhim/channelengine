# ChannelShipmentLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | **string** | The unique product reference used by the Channel. | 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant. | [optional] 
**OrderLine** | Pointer to [**ChannelOrderLineResponse**](ChannelOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the order. Each item must have an unqiue key | [optional] 
**Quantity** | **int32** | Number of items of the product in the shipment. | 

## Methods

### NewChannelShipmentLineResponse

`func NewChannelShipmentLineResponse(channelProductNo string, quantity int32, ) *ChannelShipmentLineResponse`

NewChannelShipmentLineResponse instantiates a new ChannelShipmentLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelShipmentLineResponseWithDefaults

`func NewChannelShipmentLineResponseWithDefaults() *ChannelShipmentLineResponse`

NewChannelShipmentLineResponseWithDefaults instantiates a new ChannelShipmentLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelShipmentLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelShipmentLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelShipmentLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *ChannelShipmentLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelShipmentLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelShipmentLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelShipmentLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelShipmentLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelShipmentLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetOrderLine

`func (o *ChannelShipmentLineResponse) GetOrderLine() ChannelOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *ChannelShipmentLineResponse) GetOrderLineOk() (*ChannelOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *ChannelShipmentLineResponse) SetOrderLine(v ChannelOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *ChannelShipmentLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *ChannelShipmentLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *ChannelShipmentLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *ChannelShipmentLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *ChannelShipmentLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetExtraData

`func (o *ChannelShipmentLineResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelShipmentLineResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelShipmentLineResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelShipmentLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelShipmentLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelShipmentLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetQuantity

`func (o *ChannelShipmentLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelShipmentLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelShipmentLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


