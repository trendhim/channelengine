# MerchantReturnLineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 
**Quantity** | **int32** | Number of items of the product in this return. | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the returnline. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantReturnLineRequest

`func NewMerchantReturnLineRequest(merchantProductNo string, quantity int32, ) *MerchantReturnLineRequest`

NewMerchantReturnLineRequest instantiates a new MerchantReturnLineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnLineRequestWithDefaults

`func NewMerchantReturnLineRequestWithDefaults() *MerchantReturnLineRequest`

NewMerchantReturnLineRequestWithDefaults instantiates a new MerchantReturnLineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantReturnLineRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantReturnLineRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantReturnLineRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetQuantity

`func (o *MerchantReturnLineRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantReturnLineRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantReturnLineRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetExtraData

`func (o *MerchantReturnLineRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantReturnLineRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantReturnLineRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantReturnLineRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantReturnLineRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantReturnLineRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


