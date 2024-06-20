# MerchantProductExtraDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** |  | 
**Operations** | [**[]ProductExtraDataRequest**](ProductExtraDataRequest.md) |  | 

## Methods

### NewMerchantProductExtraDataRequest

`func NewMerchantProductExtraDataRequest(merchantProductNo string, operations []ProductExtraDataRequest, ) *MerchantProductExtraDataRequest`

NewMerchantProductExtraDataRequest instantiates a new MerchantProductExtraDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductExtraDataRequestWithDefaults

`func NewMerchantProductExtraDataRequestWithDefaults() *MerchantProductExtraDataRequest`

NewMerchantProductExtraDataRequestWithDefaults instantiates a new MerchantProductExtraDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantProductExtraDataRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantProductExtraDataRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantProductExtraDataRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetOperations

`func (o *MerchantProductExtraDataRequest) GetOperations() []ProductExtraDataRequest`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MerchantProductExtraDataRequest) GetOperationsOk() (*[]ProductExtraDataRequest, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MerchantProductExtraDataRequest) SetOperations(v []ProductExtraDataRequest)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


