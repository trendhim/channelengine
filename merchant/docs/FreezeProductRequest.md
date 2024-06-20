# FreezeProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** |  | 
**Reason** | **string** |  | 
**Action** | [**FreezingActionRequest**](FreezingActionRequest.md) |  | 

## Methods

### NewFreezeProductRequest

`func NewFreezeProductRequest(merchantProductNo string, reason string, action FreezingActionRequest, ) *FreezeProductRequest`

NewFreezeProductRequest instantiates a new FreezeProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreezeProductRequestWithDefaults

`func NewFreezeProductRequestWithDefaults() *FreezeProductRequest`

NewFreezeProductRequestWithDefaults instantiates a new FreezeProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *FreezeProductRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *FreezeProductRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *FreezeProductRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetReason

`func (o *FreezeProductRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FreezeProductRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FreezeProductRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAction

`func (o *FreezeProductRequest) GetAction() FreezingActionRequest`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FreezeProductRequest) GetActionOk() (*FreezingActionRequest, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FreezeProductRequest) SetAction(v FreezingActionRequest)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


