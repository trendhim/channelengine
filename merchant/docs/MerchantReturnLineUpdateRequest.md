# MerchantReturnLineUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 
**AcceptedQuantity** | **int32** | The amount of items that have been accepted. | 
**RejectedQuantity** | **int32** | The amount of items that have been rejected. | 

## Methods

### NewMerchantReturnLineUpdateRequest

`func NewMerchantReturnLineUpdateRequest(merchantProductNo string, acceptedQuantity int32, rejectedQuantity int32, ) *MerchantReturnLineUpdateRequest`

NewMerchantReturnLineUpdateRequest instantiates a new MerchantReturnLineUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnLineUpdateRequestWithDefaults

`func NewMerchantReturnLineUpdateRequestWithDefaults() *MerchantReturnLineUpdateRequest`

NewMerchantReturnLineUpdateRequestWithDefaults instantiates a new MerchantReturnLineUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantReturnLineUpdateRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantReturnLineUpdateRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantReturnLineUpdateRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetAcceptedQuantity

`func (o *MerchantReturnLineUpdateRequest) GetAcceptedQuantity() int32`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *MerchantReturnLineUpdateRequest) GetAcceptedQuantityOk() (*int32, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *MerchantReturnLineUpdateRequest) SetAcceptedQuantity(v int32)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.


### GetRejectedQuantity

`func (o *MerchantReturnLineUpdateRequest) GetRejectedQuantity() int32`

GetRejectedQuantity returns the RejectedQuantity field if non-nil, zero value otherwise.

### GetRejectedQuantityOk

`func (o *MerchantReturnLineUpdateRequest) GetRejectedQuantityOk() (*int32, bool)`

GetRejectedQuantityOk returns a tuple with the RejectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedQuantity

`func (o *MerchantReturnLineUpdateRequest) SetRejectedQuantity(v int32)`

SetRejectedQuantity sets RejectedQuantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


