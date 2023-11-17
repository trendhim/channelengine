# MerchantReturnAcknowledgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnId** | Pointer to **int32** |  | [optional] 
**MerchantReturnNo** | **string** |  | 

## Methods

### NewMerchantReturnAcknowledgeRequest

`func NewMerchantReturnAcknowledgeRequest(merchantReturnNo string, ) *MerchantReturnAcknowledgeRequest`

NewMerchantReturnAcknowledgeRequest instantiates a new MerchantReturnAcknowledgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnAcknowledgeRequestWithDefaults

`func NewMerchantReturnAcknowledgeRequestWithDefaults() *MerchantReturnAcknowledgeRequest`

NewMerchantReturnAcknowledgeRequestWithDefaults instantiates a new MerchantReturnAcknowledgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnId

`func (o *MerchantReturnAcknowledgeRequest) GetReturnId() int32`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *MerchantReturnAcknowledgeRequest) GetReturnIdOk() (*int32, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *MerchantReturnAcknowledgeRequest) SetReturnId(v int32)`

SetReturnId sets ReturnId field to given value.

### HasReturnId

`func (o *MerchantReturnAcknowledgeRequest) HasReturnId() bool`

HasReturnId returns a boolean if a field has been set.

### GetMerchantReturnNo

`func (o *MerchantReturnAcknowledgeRequest) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *MerchantReturnAcknowledgeRequest) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *MerchantReturnAcknowledgeRequest) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


