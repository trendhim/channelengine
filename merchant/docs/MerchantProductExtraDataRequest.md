# MerchantProductExtraDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**Operations** | Pointer to [**[]ProductExtraDataRequest**](ProductExtraDataRequest.md) |  | [optional] 

## Methods

### NewMerchantProductExtraDataRequest

`func NewMerchantProductExtraDataRequest() *MerchantProductExtraDataRequest`

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

### HasMerchantProductNo

`func (o *MerchantProductExtraDataRequest) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantProductExtraDataRequest) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantProductExtraDataRequest) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
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

### HasOperations

`func (o *MerchantProductExtraDataRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperationsNil

`func (o *MerchantProductExtraDataRequest) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *MerchantProductExtraDataRequest) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


