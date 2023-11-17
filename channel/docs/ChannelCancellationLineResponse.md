# ChannelCancellationLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | **string** | The unique product reference used by the Channel. | 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant. | [optional] 
**OrderLine** | Pointer to [**ChannelOrderLineResponse**](ChannelOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**Quantity** | **int32** | Quantity of the product to cancel. | 

## Methods

### NewChannelCancellationLineResponse

`func NewChannelCancellationLineResponse(channelProductNo string, quantity int32, ) *ChannelCancellationLineResponse`

NewChannelCancellationLineResponse instantiates a new ChannelCancellationLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelCancellationLineResponseWithDefaults

`func NewChannelCancellationLineResponseWithDefaults() *ChannelCancellationLineResponse`

NewChannelCancellationLineResponseWithDefaults instantiates a new ChannelCancellationLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelCancellationLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelCancellationLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelCancellationLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *ChannelCancellationLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelCancellationLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelCancellationLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelCancellationLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelCancellationLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelCancellationLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetOrderLine

`func (o *ChannelCancellationLineResponse) GetOrderLine() ChannelOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *ChannelCancellationLineResponse) GetOrderLineOk() (*ChannelOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *ChannelCancellationLineResponse) SetOrderLine(v ChannelOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *ChannelCancellationLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *ChannelCancellationLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *ChannelCancellationLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *ChannelCancellationLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *ChannelCancellationLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *ChannelCancellationLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelCancellationLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelCancellationLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


