# BulkChannelCreateCancellationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineIdentifierType** | Pointer to [**ChannelCreateCancellationsOrderLineIdentifierType**](ChannelCreateCancellationsOrderLineIdentifierType.md) |  | [optional] 
**IdentifierType** | Pointer to [**ChannelCreateCancellationsOrderIdentifierType**](ChannelCreateCancellationsOrderIdentifierType.md) |  | [optional] 
**Models** | Pointer to [**[]ChannelCancellation**](ChannelCancellation.md) |  | [optional] 

## Methods

### NewBulkChannelCreateCancellationsRequest

`func NewBulkChannelCreateCancellationsRequest() *BulkChannelCreateCancellationsRequest`

NewBulkChannelCreateCancellationsRequest instantiates a new BulkChannelCreateCancellationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkChannelCreateCancellationsRequestWithDefaults

`func NewBulkChannelCreateCancellationsRequestWithDefaults() *BulkChannelCreateCancellationsRequest`

NewBulkChannelCreateCancellationsRequestWithDefaults instantiates a new BulkChannelCreateCancellationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) GetLineIdentifierType() ChannelCreateCancellationsOrderLineIdentifierType`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *BulkChannelCreateCancellationsRequest) GetLineIdentifierTypeOk() (*ChannelCreateCancellationsOrderLineIdentifierType, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) SetLineIdentifierType(v ChannelCreateCancellationsOrderLineIdentifierType)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) GetIdentifierType() ChannelCreateCancellationsOrderIdentifierType`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *BulkChannelCreateCancellationsRequest) GetIdentifierTypeOk() (*ChannelCreateCancellationsOrderIdentifierType, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) SetIdentifierType(v ChannelCreateCancellationsOrderIdentifierType)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *BulkChannelCreateCancellationsRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModels

`func (o *BulkChannelCreateCancellationsRequest) GetModels() []ChannelCancellation`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *BulkChannelCreateCancellationsRequest) GetModelsOk() (*[]ChannelCancellation, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *BulkChannelCreateCancellationsRequest) SetModels(v []ChannelCancellation)`

SetModels sets Models field to given value.

### HasModels

`func (o *BulkChannelCreateCancellationsRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *BulkChannelCreateCancellationsRequest) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *BulkChannelCreateCancellationsRequest) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


