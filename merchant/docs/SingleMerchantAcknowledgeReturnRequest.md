# SingleMerchantAcknowledgeReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**IdentifierType** | Pointer to [**ReturnIdentifier**](ReturnIdentifier.md) |  | [optional] 
**Model** | Pointer to [**MerchantAcknowledgeReturn**](MerchantAcknowledgeReturn.md) |  | [optional] 

## Methods

### NewSingleMerchantAcknowledgeReturnRequest

`func NewSingleMerchantAcknowledgeReturnRequest() *SingleMerchantAcknowledgeReturnRequest`

NewSingleMerchantAcknowledgeReturnRequest instantiates a new SingleMerchantAcknowledgeReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantAcknowledgeReturnRequestWithDefaults

`func NewSingleMerchantAcknowledgeReturnRequestWithDefaults() *SingleMerchantAcknowledgeReturnRequest`

NewSingleMerchantAcknowledgeReturnRequestWithDefaults instantiates a new SingleMerchantAcknowledgeReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantAcknowledgeReturnRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantAcknowledgeReturnRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantAcknowledgeReturnRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantAcknowledgeReturnRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantAcknowledgeReturnRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantAcknowledgeReturnRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetIdentifierType

`func (o *SingleMerchantAcknowledgeReturnRequest) GetIdentifierType() ReturnIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantAcknowledgeReturnRequest) GetIdentifierTypeOk() (*ReturnIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantAcknowledgeReturnRequest) SetIdentifierType(v ReturnIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantAcknowledgeReturnRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantAcknowledgeReturnRequest) GetModel() MerchantAcknowledgeReturn`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantAcknowledgeReturnRequest) GetModelOk() (*MerchantAcknowledgeReturn, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantAcknowledgeReturnRequest) SetModel(v MerchantAcknowledgeReturn)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantAcknowledgeReturnRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


