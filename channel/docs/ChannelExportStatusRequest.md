# ChannelExportStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]ChannelExportStatus**](ChannelExportStatus.md) |  | [optional] 
**MaxNumberOfExportAttempts** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewChannelExportStatusRequest

`func NewChannelExportStatusRequest() *ChannelExportStatusRequest`

NewChannelExportStatusRequest instantiates a new ChannelExportStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelExportStatusRequestWithDefaults

`func NewChannelExportStatusRequestWithDefaults() *ChannelExportStatusRequest`

NewChannelExportStatusRequestWithDefaults instantiates a new ChannelExportStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *ChannelExportStatusRequest) GetStatuses() []ChannelExportStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ChannelExportStatusRequest) GetStatusesOk() (*[]ChannelExportStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ChannelExportStatusRequest) SetStatuses(v []ChannelExportStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ChannelExportStatusRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *ChannelExportStatusRequest) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *ChannelExportStatusRequest) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetMaxNumberOfExportAttempts

`func (o *ChannelExportStatusRequest) GetMaxNumberOfExportAttempts() int32`

GetMaxNumberOfExportAttempts returns the MaxNumberOfExportAttempts field if non-nil, zero value otherwise.

### GetMaxNumberOfExportAttemptsOk

`func (o *ChannelExportStatusRequest) GetMaxNumberOfExportAttemptsOk() (*int32, bool)`

GetMaxNumberOfExportAttemptsOk returns a tuple with the MaxNumberOfExportAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfExportAttempts

`func (o *ChannelExportStatusRequest) SetMaxNumberOfExportAttempts(v int32)`

SetMaxNumberOfExportAttempts sets MaxNumberOfExportAttempts field to given value.

### HasMaxNumberOfExportAttempts

`func (o *ChannelExportStatusRequest) HasMaxNumberOfExportAttempts() bool`

HasMaxNumberOfExportAttempts returns a boolean if a field has been set.

### SetMaxNumberOfExportAttemptsNil

`func (o *ChannelExportStatusRequest) SetMaxNumberOfExportAttemptsNil(b bool)`

 SetMaxNumberOfExportAttemptsNil sets the value for MaxNumberOfExportAttempts to be an explicit nil

### UnsetMaxNumberOfExportAttempts
`func (o *ChannelExportStatusRequest) UnsetMaxNumberOfExportAttempts()`

UnsetMaxNumberOfExportAttempts ensures that no value is present for MaxNumberOfExportAttempts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


