# MerchantShipmentLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant. | 
**MerchantBundleProductNo** | Pointer to **NullableString** | The unique bundle product reference used by the Merchant. | [optional] 
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel. | [optional] 
**OrderLine** | Pointer to [**MerchantOrderLineResponse**](MerchantOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the shipment line. Each item must have an unqiue key | [optional] 
**Quantity** | **int32** | Number of items of the product in the shipment. | 

## Methods

### NewMerchantShipmentLineResponse

`func NewMerchantShipmentLineResponse(merchantProductNo string, quantity int32, ) *MerchantShipmentLineResponse`

NewMerchantShipmentLineResponse instantiates a new MerchantShipmentLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantShipmentLineResponseWithDefaults

`func NewMerchantShipmentLineResponseWithDefaults() *MerchantShipmentLineResponse`

NewMerchantShipmentLineResponseWithDefaults instantiates a new MerchantShipmentLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantShipmentLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantShipmentLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantShipmentLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetMerchantBundleProductNo

`func (o *MerchantShipmentLineResponse) GetMerchantBundleProductNo() string`

GetMerchantBundleProductNo returns the MerchantBundleProductNo field if non-nil, zero value otherwise.

### GetMerchantBundleProductNoOk

`func (o *MerchantShipmentLineResponse) GetMerchantBundleProductNoOk() (*string, bool)`

GetMerchantBundleProductNoOk returns a tuple with the MerchantBundleProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantBundleProductNo

`func (o *MerchantShipmentLineResponse) SetMerchantBundleProductNo(v string)`

SetMerchantBundleProductNo sets MerchantBundleProductNo field to given value.

### HasMerchantBundleProductNo

`func (o *MerchantShipmentLineResponse) HasMerchantBundleProductNo() bool`

HasMerchantBundleProductNo returns a boolean if a field has been set.

### SetMerchantBundleProductNoNil

`func (o *MerchantShipmentLineResponse) SetMerchantBundleProductNoNil(b bool)`

 SetMerchantBundleProductNoNil sets the value for MerchantBundleProductNo to be an explicit nil

### UnsetMerchantBundleProductNo
`func (o *MerchantShipmentLineResponse) UnsetMerchantBundleProductNo()`

UnsetMerchantBundleProductNo ensures that no value is present for MerchantBundleProductNo, not even an explicit nil
### GetChannelProductNo

`func (o *MerchantShipmentLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *MerchantShipmentLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *MerchantShipmentLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *MerchantShipmentLineResponse) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *MerchantShipmentLineResponse) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *MerchantShipmentLineResponse) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetOrderLine

`func (o *MerchantShipmentLineResponse) GetOrderLine() MerchantOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *MerchantShipmentLineResponse) GetOrderLineOk() (*MerchantOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *MerchantShipmentLineResponse) SetOrderLine(v MerchantOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *MerchantShipmentLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *MerchantShipmentLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *MerchantShipmentLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *MerchantShipmentLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *MerchantShipmentLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetExtraData

`func (o *MerchantShipmentLineResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantShipmentLineResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantShipmentLineResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantShipmentLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantShipmentLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantShipmentLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetQuantity

`func (o *MerchantShipmentLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantShipmentLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantShipmentLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


