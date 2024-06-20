# CollectionChangesOfChannelProductChangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**ChannelProductChangesResponse**](ChannelProductChangesResponse.md) |  | [optional] 
**ToBeCreatedTotalCount** | Pointer to **int32** | The total number of changes to create. | [optional] 
**ToBeUpdatedTotalCount** | Pointer to **int32** | The total number of changes to update. | [optional] 
**ToBeDeletedTotalCount** | Pointer to **int32** | The total number of changes to delete. | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCollectionChangesOfChannelProductChangesResponse

`func NewCollectionChangesOfChannelProductChangesResponse() *CollectionChangesOfChannelProductChangesResponse`

NewCollectionChangesOfChannelProductChangesResponse instantiates a new CollectionChangesOfChannelProductChangesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionChangesOfChannelProductChangesResponseWithDefaults

`func NewCollectionChangesOfChannelProductChangesResponseWithDefaults() *CollectionChangesOfChannelProductChangesResponse`

NewCollectionChangesOfChannelProductChangesResponseWithDefaults instantiates a new CollectionChangesOfChannelProductChangesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionChangesOfChannelProductChangesResponse) GetContent() ChannelProductChangesResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetContentOk() (*ChannelProductChangesResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionChangesOfChannelProductChangesResponse) SetContent(v ChannelProductChangesResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionChangesOfChannelProductChangesResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetToBeCreatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeCreatedTotalCount() int32`

GetToBeCreatedTotalCount returns the ToBeCreatedTotalCount field if non-nil, zero value otherwise.

### GetToBeCreatedTotalCountOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeCreatedTotalCountOk() (*int32, bool)`

GetToBeCreatedTotalCountOk returns a tuple with the ToBeCreatedTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeCreatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) SetToBeCreatedTotalCount(v int32)`

SetToBeCreatedTotalCount sets ToBeCreatedTotalCount field to given value.

### HasToBeCreatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) HasToBeCreatedTotalCount() bool`

HasToBeCreatedTotalCount returns a boolean if a field has been set.

### GetToBeUpdatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeUpdatedTotalCount() int32`

GetToBeUpdatedTotalCount returns the ToBeUpdatedTotalCount field if non-nil, zero value otherwise.

### GetToBeUpdatedTotalCountOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeUpdatedTotalCountOk() (*int32, bool)`

GetToBeUpdatedTotalCountOk returns a tuple with the ToBeUpdatedTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeUpdatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) SetToBeUpdatedTotalCount(v int32)`

SetToBeUpdatedTotalCount sets ToBeUpdatedTotalCount field to given value.

### HasToBeUpdatedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) HasToBeUpdatedTotalCount() bool`

HasToBeUpdatedTotalCount returns a boolean if a field has been set.

### GetToBeDeletedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeDeletedTotalCount() int32`

GetToBeDeletedTotalCount returns the ToBeDeletedTotalCount field if non-nil, zero value otherwise.

### GetToBeDeletedTotalCountOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetToBeDeletedTotalCountOk() (*int32, bool)`

GetToBeDeletedTotalCountOk returns a tuple with the ToBeDeletedTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeDeletedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) SetToBeDeletedTotalCount(v int32)`

SetToBeDeletedTotalCount sets ToBeDeletedTotalCount field to given value.

### HasToBeDeletedTotalCount

`func (o *CollectionChangesOfChannelProductChangesResponse) HasToBeDeletedTotalCount() bool`

HasToBeDeletedTotalCount returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionChangesOfChannelProductChangesResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionChangesOfChannelProductChangesResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionChangesOfChannelProductChangesResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionChangesOfChannelProductChangesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionChangesOfChannelProductChangesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionChangesOfChannelProductChangesResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionChangesOfChannelProductChangesResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionChangesOfChannelProductChangesResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionChangesOfChannelProductChangesResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionChangesOfChannelProductChangesResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionChangesOfChannelProductChangesResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionChangesOfChannelProductChangesResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionChangesOfChannelProductChangesResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionChangesOfChannelProductChangesResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionChangesOfChannelProductChangesResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionChangesOfChannelProductChangesResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionChangesOfChannelProductChangesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionChangesOfChannelProductChangesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionChangesOfChannelProductChangesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionChangesOfChannelProductChangesResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionChangesOfChannelProductChangesResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *CollectionChangesOfChannelProductChangesResponse) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *CollectionChangesOfChannelProductChangesResponse) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *CollectionChangesOfChannelProductChangesResponse) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *CollectionChangesOfChannelProductChangesResponse) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *CollectionChangesOfChannelProductChangesResponse) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *CollectionChangesOfChannelProductChangesResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionChangesOfChannelProductChangesResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionChangesOfChannelProductChangesResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionChangesOfChannelProductChangesResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionChangesOfChannelProductChangesResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionChangesOfChannelProductChangesResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


