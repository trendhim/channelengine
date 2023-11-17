# MerchantCancellationLineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 
**Quantity** | **int32** | Quantity of the product to cancel. | 

## Methods

### NewMerchantCancellationLineRequest

`func NewMerchantCancellationLineRequest(merchantProductNo string, quantity int32, ) *MerchantCancellationLineRequest`

NewMerchantCancellationLineRequest instantiates a new MerchantCancellationLineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCancellationLineRequestWithDefaults

`func NewMerchantCancellationLineRequestWithDefaults() *MerchantCancellationLineRequest`

NewMerchantCancellationLineRequestWithDefaults instantiates a new MerchantCancellationLineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantCancellationLineRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantCancellationLineRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantCancellationLineRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetQuantity

`func (o *MerchantCancellationLineRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantCancellationLineRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantCancellationLineRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


