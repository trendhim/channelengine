# ChannelReturnLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | **string** | The unique product reference used by the Channel. | 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant. | [optional] 
**AcceptedQuantity** | Pointer to **NullableInt32** | The number of accepted items of the product in this return (should normally be refunded). | [optional] 
**RejectedQuantity** | Pointer to **NullableInt32** | The number of rejected items of the product in this return (should normally not be refunded). | [optional] 
**OrderLine** | Pointer to [**ChannelOrderLineResponse**](ChannelOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**Quantity** | **int32** | Number of items of the product in this return. | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the returnline. Each item must have an unqiue key | [optional] 

## Methods

### NewChannelReturnLineResponse

`func NewChannelReturnLineResponse(channelProductNo string, quantity int32, ) *ChannelReturnLineResponse`

NewChannelReturnLineResponse instantiates a new ChannelReturnLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelReturnLineResponseWithDefaults

`func NewChannelReturnLineResponseWithDefaults() *ChannelReturnLineResponse`

NewChannelReturnLineResponseWithDefaults instantiates a new ChannelReturnLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelReturnLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelReturnLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelReturnLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *ChannelReturnLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelReturnLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelReturnLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelReturnLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelReturnLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelReturnLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetAcceptedQuantity

`func (o *ChannelReturnLineResponse) GetAcceptedQuantity() int32`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *ChannelReturnLineResponse) GetAcceptedQuantityOk() (*int32, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *ChannelReturnLineResponse) SetAcceptedQuantity(v int32)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.

### HasAcceptedQuantity

`func (o *ChannelReturnLineResponse) HasAcceptedQuantity() bool`

HasAcceptedQuantity returns a boolean if a field has been set.

### SetAcceptedQuantityNil

`func (o *ChannelReturnLineResponse) SetAcceptedQuantityNil(b bool)`

 SetAcceptedQuantityNil sets the value for AcceptedQuantity to be an explicit nil

### UnsetAcceptedQuantity
`func (o *ChannelReturnLineResponse) UnsetAcceptedQuantity()`

UnsetAcceptedQuantity ensures that no value is present for AcceptedQuantity, not even an explicit nil
### GetRejectedQuantity

`func (o *ChannelReturnLineResponse) GetRejectedQuantity() int32`

GetRejectedQuantity returns the RejectedQuantity field if non-nil, zero value otherwise.

### GetRejectedQuantityOk

`func (o *ChannelReturnLineResponse) GetRejectedQuantityOk() (*int32, bool)`

GetRejectedQuantityOk returns a tuple with the RejectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedQuantity

`func (o *ChannelReturnLineResponse) SetRejectedQuantity(v int32)`

SetRejectedQuantity sets RejectedQuantity field to given value.

### HasRejectedQuantity

`func (o *ChannelReturnLineResponse) HasRejectedQuantity() bool`

HasRejectedQuantity returns a boolean if a field has been set.

### SetRejectedQuantityNil

`func (o *ChannelReturnLineResponse) SetRejectedQuantityNil(b bool)`

 SetRejectedQuantityNil sets the value for RejectedQuantity to be an explicit nil

### UnsetRejectedQuantity
`func (o *ChannelReturnLineResponse) UnsetRejectedQuantity()`

UnsetRejectedQuantity ensures that no value is present for RejectedQuantity, not even an explicit nil
### GetOrderLine

`func (o *ChannelReturnLineResponse) GetOrderLine() ChannelOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *ChannelReturnLineResponse) GetOrderLineOk() (*ChannelOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *ChannelReturnLineResponse) SetOrderLine(v ChannelOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *ChannelReturnLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *ChannelReturnLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *ChannelReturnLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *ChannelReturnLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *ChannelReturnLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *ChannelReturnLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelReturnLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelReturnLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetExtraData

`func (o *ChannelReturnLineResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelReturnLineResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelReturnLineResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelReturnLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelReturnLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelReturnLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


