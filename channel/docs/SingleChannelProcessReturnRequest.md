# SingleChannelProcessReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | Pointer to [**ReturnIdentifier**](ReturnIdentifier.md) |  | [optional] 
**Model** | Pointer to [**ChannelProcessReturn**](ChannelProcessReturn.md) |  | [optional] 

## Methods

### NewSingleChannelProcessReturnRequest

`func NewSingleChannelProcessReturnRequest() *SingleChannelProcessReturnRequest`

NewSingleChannelProcessReturnRequest instantiates a new SingleChannelProcessReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleChannelProcessReturnRequestWithDefaults

`func NewSingleChannelProcessReturnRequestWithDefaults() *SingleChannelProcessReturnRequest`

NewSingleChannelProcessReturnRequestWithDefaults instantiates a new SingleChannelProcessReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *SingleChannelProcessReturnRequest) GetIdentifierType() ReturnIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleChannelProcessReturnRequest) GetIdentifierTypeOk() (*ReturnIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleChannelProcessReturnRequest) SetIdentifierType(v ReturnIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleChannelProcessReturnRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleChannelProcessReturnRequest) GetModel() ChannelProcessReturn`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleChannelProcessReturnRequest) GetModelOk() (*ChannelProcessReturn, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleChannelProcessReturnRequest) SetModel(v ChannelProcessReturn)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleChannelProcessReturnRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


