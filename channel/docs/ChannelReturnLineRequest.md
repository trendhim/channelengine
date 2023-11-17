# ChannelReturnLineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel. | [optional] 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant. | [optional] 
**Quantity** | **int32** | Number of items of the product in this return. | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the returnline. Each item must have an unqiue key | [optional] 

## Methods

### NewChannelReturnLineRequest

`func NewChannelReturnLineRequest(quantity int32, ) *ChannelReturnLineRequest`

NewChannelReturnLineRequest instantiates a new ChannelReturnLineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelReturnLineRequestWithDefaults

`func NewChannelReturnLineRequestWithDefaults() *ChannelReturnLineRequest`

NewChannelReturnLineRequestWithDefaults instantiates a new ChannelReturnLineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelReturnLineRequest) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelReturnLineRequest) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelReturnLineRequest) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *ChannelReturnLineRequest) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *ChannelReturnLineRequest) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *ChannelReturnLineRequest) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetMerchantProductNo

`func (o *ChannelReturnLineRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelReturnLineRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelReturnLineRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelReturnLineRequest) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelReturnLineRequest) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelReturnLineRequest) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *ChannelReturnLineRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelReturnLineRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelReturnLineRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetExtraData

`func (o *ChannelReturnLineRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelReturnLineRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelReturnLineRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelReturnLineRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelReturnLineRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelReturnLineRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


