# SingleMerchantHandleReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**LineIdentifierType** | Pointer to [**ReturnLineIdentifier**](ReturnLineIdentifier.md) |  | [optional] 
**IdentifierType** | Pointer to [**ReturnIdentifier**](ReturnIdentifier.md) |  | [optional] 
**Model** | Pointer to [**MerchantHandleReturn**](MerchantHandleReturn.md) |  | [optional] 

## Methods

### NewSingleMerchantHandleReturnRequest

`func NewSingleMerchantHandleReturnRequest() *SingleMerchantHandleReturnRequest`

NewSingleMerchantHandleReturnRequest instantiates a new SingleMerchantHandleReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantHandleReturnRequestWithDefaults

`func NewSingleMerchantHandleReturnRequestWithDefaults() *SingleMerchantHandleReturnRequest`

NewSingleMerchantHandleReturnRequestWithDefaults instantiates a new SingleMerchantHandleReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantHandleReturnRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantHandleReturnRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantHandleReturnRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantHandleReturnRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantHandleReturnRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantHandleReturnRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetLineIdentifierType

`func (o *SingleMerchantHandleReturnRequest) GetLineIdentifierType() ReturnLineIdentifier`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleMerchantHandleReturnRequest) GetLineIdentifierTypeOk() (*ReturnLineIdentifier, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleMerchantHandleReturnRequest) SetLineIdentifierType(v ReturnLineIdentifier)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleMerchantHandleReturnRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleMerchantHandleReturnRequest) GetIdentifierType() ReturnIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantHandleReturnRequest) GetIdentifierTypeOk() (*ReturnIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantHandleReturnRequest) SetIdentifierType(v ReturnIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantHandleReturnRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantHandleReturnRequest) GetModel() MerchantHandleReturn`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantHandleReturnRequest) GetModelOk() (*MerchantHandleReturn, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantHandleReturnRequest) SetModel(v MerchantHandleReturn)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantHandleReturnRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


