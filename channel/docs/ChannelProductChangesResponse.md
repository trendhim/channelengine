# ChannelProductChangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToBeCreated** | Pointer to [**[]ChannelProductResponse**](ChannelProductResponse.md) |  | [optional] 
**ToBeUpdated** | Pointer to [**[]ChannelProductResponse**](ChannelProductResponse.md) |  | [optional] 
**ToBeRemoved** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChannelProductChangesResponse

`func NewChannelProductChangesResponse() *ChannelProductChangesResponse`

NewChannelProductChangesResponse instantiates a new ChannelProductChangesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProductChangesResponseWithDefaults

`func NewChannelProductChangesResponseWithDefaults() *ChannelProductChangesResponse`

NewChannelProductChangesResponseWithDefaults instantiates a new ChannelProductChangesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToBeCreated

`func (o *ChannelProductChangesResponse) GetToBeCreated() []ChannelProductResponse`

GetToBeCreated returns the ToBeCreated field if non-nil, zero value otherwise.

### GetToBeCreatedOk

`func (o *ChannelProductChangesResponse) GetToBeCreatedOk() (*[]ChannelProductResponse, bool)`

GetToBeCreatedOk returns a tuple with the ToBeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeCreated

`func (o *ChannelProductChangesResponse) SetToBeCreated(v []ChannelProductResponse)`

SetToBeCreated sets ToBeCreated field to given value.

### HasToBeCreated

`func (o *ChannelProductChangesResponse) HasToBeCreated() bool`

HasToBeCreated returns a boolean if a field has been set.

### SetToBeCreatedNil

`func (o *ChannelProductChangesResponse) SetToBeCreatedNil(b bool)`

 SetToBeCreatedNil sets the value for ToBeCreated to be an explicit nil

### UnsetToBeCreated
`func (o *ChannelProductChangesResponse) UnsetToBeCreated()`

UnsetToBeCreated ensures that no value is present for ToBeCreated, not even an explicit nil
### GetToBeUpdated

`func (o *ChannelProductChangesResponse) GetToBeUpdated() []ChannelProductResponse`

GetToBeUpdated returns the ToBeUpdated field if non-nil, zero value otherwise.

### GetToBeUpdatedOk

`func (o *ChannelProductChangesResponse) GetToBeUpdatedOk() (*[]ChannelProductResponse, bool)`

GetToBeUpdatedOk returns a tuple with the ToBeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeUpdated

`func (o *ChannelProductChangesResponse) SetToBeUpdated(v []ChannelProductResponse)`

SetToBeUpdated sets ToBeUpdated field to given value.

### HasToBeUpdated

`func (o *ChannelProductChangesResponse) HasToBeUpdated() bool`

HasToBeUpdated returns a boolean if a field has been set.

### SetToBeUpdatedNil

`func (o *ChannelProductChangesResponse) SetToBeUpdatedNil(b bool)`

 SetToBeUpdatedNil sets the value for ToBeUpdated to be an explicit nil

### UnsetToBeUpdated
`func (o *ChannelProductChangesResponse) UnsetToBeUpdated()`

UnsetToBeUpdated ensures that no value is present for ToBeUpdated, not even an explicit nil
### GetToBeRemoved

`func (o *ChannelProductChangesResponse) GetToBeRemoved() []string`

GetToBeRemoved returns the ToBeRemoved field if non-nil, zero value otherwise.

### GetToBeRemovedOk

`func (o *ChannelProductChangesResponse) GetToBeRemovedOk() (*[]string, bool)`

GetToBeRemovedOk returns a tuple with the ToBeRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeRemoved

`func (o *ChannelProductChangesResponse) SetToBeRemoved(v []string)`

SetToBeRemoved sets ToBeRemoved field to given value.

### HasToBeRemoved

`func (o *ChannelProductChangesResponse) HasToBeRemoved() bool`

HasToBeRemoved returns a boolean if a field has been set.

### SetToBeRemovedNil

`func (o *ChannelProductChangesResponse) SetToBeRemovedNil(b bool)`

 SetToBeRemovedNil sets the value for ToBeRemoved to be an explicit nil

### UnsetToBeRemoved
`func (o *ChannelProductChangesResponse) UnsetToBeRemoved()`

UnsetToBeRemoved ensures that no value is present for ToBeRemoved, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


