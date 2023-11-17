# MerchantOrderCommentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderNo** | Pointer to **NullableString** | Your own order reference for the order you would like to update the comment for.  Either this field or OrderId is required | [optional] 
**OrderId** | Pointer to **NullableInt32** | The ChannelEngine order ID of the order you would like to update the comment for.  Either this field or MerchantOrderNo is required | [optional] 
**MerchantComment** | **string** | The merchant comment you would like add / update for the order. | 

## Methods

### NewMerchantOrderCommentUpdateRequest

`func NewMerchantOrderCommentUpdateRequest(merchantComment string, ) *MerchantOrderCommentUpdateRequest`

NewMerchantOrderCommentUpdateRequest instantiates a new MerchantOrderCommentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderCommentUpdateRequestWithDefaults

`func NewMerchantOrderCommentUpdateRequestWithDefaults() *MerchantOrderCommentUpdateRequest`

NewMerchantOrderCommentUpdateRequestWithDefaults instantiates a new MerchantOrderCommentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderNo

`func (o *MerchantOrderCommentUpdateRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantOrderCommentUpdateRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantOrderCommentUpdateRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantOrderCommentUpdateRequest) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantOrderCommentUpdateRequest) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantOrderCommentUpdateRequest) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetOrderId

`func (o *MerchantOrderCommentUpdateRequest) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MerchantOrderCommentUpdateRequest) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MerchantOrderCommentUpdateRequest) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MerchantOrderCommentUpdateRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *MerchantOrderCommentUpdateRequest) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *MerchantOrderCommentUpdateRequest) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetMerchantComment

`func (o *MerchantOrderCommentUpdateRequest) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantOrderCommentUpdateRequest) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantOrderCommentUpdateRequest) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


