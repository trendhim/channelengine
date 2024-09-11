# SingleChannelCreateReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineIdentifierType** | Pointer to [**OrderLineIdentifier**](OrderLineIdentifier.md) |  | [optional] 
**IdentifierType** | Pointer to [**OrderIdentifier**](OrderIdentifier.md) |  | [optional] 
**Model** | Pointer to [**ChannelCreateReturn**](ChannelCreateReturn.md) |  | [optional] 

## Methods

### NewSingleChannelCreateReturnRequest

`func NewSingleChannelCreateReturnRequest() *SingleChannelCreateReturnRequest`

NewSingleChannelCreateReturnRequest instantiates a new SingleChannelCreateReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleChannelCreateReturnRequestWithDefaults

`func NewSingleChannelCreateReturnRequestWithDefaults() *SingleChannelCreateReturnRequest`

NewSingleChannelCreateReturnRequestWithDefaults instantiates a new SingleChannelCreateReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineIdentifierType

`func (o *SingleChannelCreateReturnRequest) GetLineIdentifierType() OrderLineIdentifier`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleChannelCreateReturnRequest) GetLineIdentifierTypeOk() (*OrderLineIdentifier, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleChannelCreateReturnRequest) SetLineIdentifierType(v OrderLineIdentifier)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleChannelCreateReturnRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleChannelCreateReturnRequest) GetIdentifierType() OrderIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleChannelCreateReturnRequest) GetIdentifierTypeOk() (*OrderIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleChannelCreateReturnRequest) SetIdentifierType(v OrderIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleChannelCreateReturnRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleChannelCreateReturnRequest) GetModel() ChannelCreateReturn`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleChannelCreateReturnRequest) GetModelOk() (*ChannelCreateReturn, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleChannelCreateReturnRequest) SetModel(v ChannelCreateReturn)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleChannelCreateReturnRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


