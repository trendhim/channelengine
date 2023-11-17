# ChannelProcessedChangesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**[]ChannelProductReferencesRequest**](ChannelProductReferencesRequest.md) | A collection of pairs of merchant and channel references  of the products which are successfully created. The channel references  are saved such that in subsequent calls these can be used instead of the  merchant references. | [optional] 
**Updated** | Pointer to **[]string** | The channel references of the products which are successfully updated. | [optional] 
**Removed** | Pointer to **[]string** | The channel references of the products which are successfully removed. | [optional] 

## Methods

### NewChannelProcessedChangesRequest

`func NewChannelProcessedChangesRequest() *ChannelProcessedChangesRequest`

NewChannelProcessedChangesRequest instantiates a new ChannelProcessedChangesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProcessedChangesRequestWithDefaults

`func NewChannelProcessedChangesRequestWithDefaults() *ChannelProcessedChangesRequest`

NewChannelProcessedChangesRequestWithDefaults instantiates a new ChannelProcessedChangesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ChannelProcessedChangesRequest) GetCreated() []ChannelProductReferencesRequest`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChannelProcessedChangesRequest) GetCreatedOk() (*[]ChannelProductReferencesRequest, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChannelProcessedChangesRequest) SetCreated(v []ChannelProductReferencesRequest)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChannelProcessedChangesRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ChannelProcessedChangesRequest) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ChannelProcessedChangesRequest) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *ChannelProcessedChangesRequest) GetUpdated() []string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ChannelProcessedChangesRequest) GetUpdatedOk() (*[]string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ChannelProcessedChangesRequest) SetUpdated(v []string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ChannelProcessedChangesRequest) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *ChannelProcessedChangesRequest) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *ChannelProcessedChangesRequest) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetRemoved

`func (o *ChannelProcessedChangesRequest) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *ChannelProcessedChangesRequest) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *ChannelProcessedChangesRequest) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *ChannelProcessedChangesRequest) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### SetRemovedNil

`func (o *ChannelProcessedChangesRequest) SetRemovedNil(b bool)`

 SetRemovedNil sets the value for Removed to be an explicit nil

### UnsetRemoved
`func (o *ChannelProcessedChangesRequest) UnsetRemoved()`

UnsetRemoved ensures that no value is present for Removed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


