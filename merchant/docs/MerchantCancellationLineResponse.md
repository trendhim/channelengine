# MerchantCancellationLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique cancellation line identifier used by ChannelEngine. | [optional] 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant. | [optional] 
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel. | [optional] 
**Quantity** | **int32** | Quantity of the product to cancel. | 

## Methods

### NewMerchantCancellationLineResponse

`func NewMerchantCancellationLineResponse(quantity int32, ) *MerchantCancellationLineResponse`

NewMerchantCancellationLineResponse instantiates a new MerchantCancellationLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCancellationLineResponseWithDefaults

`func NewMerchantCancellationLineResponseWithDefaults() *MerchantCancellationLineResponse`

NewMerchantCancellationLineResponseWithDefaults instantiates a new MerchantCancellationLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantCancellationLineResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantCancellationLineResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantCancellationLineResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantCancellationLineResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantProductNo

`func (o *MerchantCancellationLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantCancellationLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantCancellationLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantCancellationLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantCancellationLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantCancellationLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetChannelProductNo

`func (o *MerchantCancellationLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *MerchantCancellationLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *MerchantCancellationLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *MerchantCancellationLineResponse) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *MerchantCancellationLineResponse) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *MerchantCancellationLineResponse) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetQuantity

`func (o *MerchantCancellationLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantCancellationLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantCancellationLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


