# CreateEditTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**Targets** | Pointer to [**[]CreateEditTargetView**](CreateEditTargetView.md) |  | [optional] 

## Methods

### NewCreateEditTargetRequest

`func NewCreateEditTargetRequest() *CreateEditTargetRequest`

NewCreateEditTargetRequest instantiates a new CreateEditTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEditTargetRequestWithDefaults

`func NewCreateEditTargetRequestWithDefaults() *CreateEditTargetRequest`

NewCreateEditTargetRequestWithDefaults instantiates a new CreateEditTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *CreateEditTargetRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateEditTargetRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateEditTargetRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *CreateEditTargetRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *CreateEditTargetRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *CreateEditTargetRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetTargets

`func (o *CreateEditTargetRequest) GetTargets() []CreateEditTargetView`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CreateEditTargetRequest) GetTargetsOk() (*[]CreateEditTargetView, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CreateEditTargetRequest) SetTargets(v []CreateEditTargetView)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CreateEditTargetRequest) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *CreateEditTargetRequest) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *CreateEditTargetRequest) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


