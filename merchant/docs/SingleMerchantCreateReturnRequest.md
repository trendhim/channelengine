# SingleMerchantCreateReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**LineIdentifierType** | Pointer to [**OrderLineIdentifier**](OrderLineIdentifier.md) |  | [optional] 
**IdentifierType** | Pointer to [**OrderIdentifier**](OrderIdentifier.md) |  | [optional] 
**Model** | Pointer to [**MerchantCreateReturn**](MerchantCreateReturn.md) |  | [optional] 

## Methods

### NewSingleMerchantCreateReturnRequest

`func NewSingleMerchantCreateReturnRequest() *SingleMerchantCreateReturnRequest`

NewSingleMerchantCreateReturnRequest instantiates a new SingleMerchantCreateReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantCreateReturnRequestWithDefaults

`func NewSingleMerchantCreateReturnRequestWithDefaults() *SingleMerchantCreateReturnRequest`

NewSingleMerchantCreateReturnRequestWithDefaults instantiates a new SingleMerchantCreateReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantCreateReturnRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantCreateReturnRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantCreateReturnRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantCreateReturnRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantCreateReturnRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantCreateReturnRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetLineIdentifierType

`func (o *SingleMerchantCreateReturnRequest) GetLineIdentifierType() OrderLineIdentifier`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleMerchantCreateReturnRequest) GetLineIdentifierTypeOk() (*OrderLineIdentifier, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleMerchantCreateReturnRequest) SetLineIdentifierType(v OrderLineIdentifier)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleMerchantCreateReturnRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleMerchantCreateReturnRequest) GetIdentifierType() OrderIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantCreateReturnRequest) GetIdentifierTypeOk() (*OrderIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantCreateReturnRequest) SetIdentifierType(v OrderIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantCreateReturnRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantCreateReturnRequest) GetModel() MerchantCreateReturn`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantCreateReturnRequest) GetModelOk() (*MerchantCreateReturn, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantCreateReturnRequest) SetModel(v MerchantCreateReturn)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantCreateReturnRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


