# SingleChannelCreateCancellationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineIdentifierType** | Pointer to [**ChannelCreateCancellationsOrderLineIdentifierType**](ChannelCreateCancellationsOrderLineIdentifierType.md) |  | [optional] 
**IdentifierType** | Pointer to [**ChannelCreateCancellationsOrderIdentifierType**](ChannelCreateCancellationsOrderIdentifierType.md) |  | [optional] 
**Model** | Pointer to [**ChannelCancellation**](ChannelCancellation.md) |  | [optional] 

## Methods

### NewSingleChannelCreateCancellationRequest

`func NewSingleChannelCreateCancellationRequest() *SingleChannelCreateCancellationRequest`

NewSingleChannelCreateCancellationRequest instantiates a new SingleChannelCreateCancellationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleChannelCreateCancellationRequestWithDefaults

`func NewSingleChannelCreateCancellationRequestWithDefaults() *SingleChannelCreateCancellationRequest`

NewSingleChannelCreateCancellationRequestWithDefaults instantiates a new SingleChannelCreateCancellationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineIdentifierType

`func (o *SingleChannelCreateCancellationRequest) GetLineIdentifierType() ChannelCreateCancellationsOrderLineIdentifierType`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleChannelCreateCancellationRequest) GetLineIdentifierTypeOk() (*ChannelCreateCancellationsOrderLineIdentifierType, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleChannelCreateCancellationRequest) SetLineIdentifierType(v ChannelCreateCancellationsOrderLineIdentifierType)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleChannelCreateCancellationRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleChannelCreateCancellationRequest) GetIdentifierType() ChannelCreateCancellationsOrderIdentifierType`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleChannelCreateCancellationRequest) GetIdentifierTypeOk() (*ChannelCreateCancellationsOrderIdentifierType, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleChannelCreateCancellationRequest) SetIdentifierType(v ChannelCreateCancellationsOrderIdentifierType)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleChannelCreateCancellationRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleChannelCreateCancellationRequest) GetModel() ChannelCancellation`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleChannelCreateCancellationRequest) GetModelOk() (*ChannelCancellation, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleChannelCreateCancellationRequest) SetModel(v ChannelCancellation)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleChannelCreateCancellationRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


