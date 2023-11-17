# MerchantOrderAcknowledgementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderNo** | **string** | Your own order reference, this will be used in consecutive order processing API calls. | 
**OrderId** | **int32** | The ChannelEngine order ID of the order you would like to acknowledge. | 

## Methods

### NewMerchantOrderAcknowledgementRequest

`func NewMerchantOrderAcknowledgementRequest(merchantOrderNo string, orderId int32, ) *MerchantOrderAcknowledgementRequest`

NewMerchantOrderAcknowledgementRequest instantiates a new MerchantOrderAcknowledgementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderAcknowledgementRequestWithDefaults

`func NewMerchantOrderAcknowledgementRequestWithDefaults() *MerchantOrderAcknowledgementRequest`

NewMerchantOrderAcknowledgementRequestWithDefaults instantiates a new MerchantOrderAcknowledgementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderNo

`func (o *MerchantOrderAcknowledgementRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantOrderAcknowledgementRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantOrderAcknowledgementRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetOrderId

`func (o *MerchantOrderAcknowledgementRequest) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MerchantOrderAcknowledgementRequest) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MerchantOrderAcknowledgementRequest) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


