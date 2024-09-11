# SingleOfDeleteTargetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**DeleteTargetsResponse**](DeleteTargetsResponse.md) |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewSingleOfDeleteTargetsResponse

`func NewSingleOfDeleteTargetsResponse() *SingleOfDeleteTargetsResponse`

NewSingleOfDeleteTargetsResponse instantiates a new SingleOfDeleteTargetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleOfDeleteTargetsResponseWithDefaults

`func NewSingleOfDeleteTargetsResponseWithDefaults() *SingleOfDeleteTargetsResponse`

NewSingleOfDeleteTargetsResponseWithDefaults instantiates a new SingleOfDeleteTargetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SingleOfDeleteTargetsResponse) GetContent() DeleteTargetsResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SingleOfDeleteTargetsResponse) GetContentOk() (*DeleteTargetsResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SingleOfDeleteTargetsResponse) SetContent(v DeleteTargetsResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *SingleOfDeleteTargetsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatusCode

`func (o *SingleOfDeleteTargetsResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SingleOfDeleteTargetsResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SingleOfDeleteTargetsResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SingleOfDeleteTargetsResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *SingleOfDeleteTargetsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SingleOfDeleteTargetsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SingleOfDeleteTargetsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SingleOfDeleteTargetsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *SingleOfDeleteTargetsResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SingleOfDeleteTargetsResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *SingleOfDeleteTargetsResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *SingleOfDeleteTargetsResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *SingleOfDeleteTargetsResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *SingleOfDeleteTargetsResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *SingleOfDeleteTargetsResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *SingleOfDeleteTargetsResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *SingleOfDeleteTargetsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SingleOfDeleteTargetsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SingleOfDeleteTargetsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SingleOfDeleteTargetsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *SingleOfDeleteTargetsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SingleOfDeleteTargetsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SingleOfDeleteTargetsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SingleOfDeleteTargetsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SingleOfDeleteTargetsResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SingleOfDeleteTargetsResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *SingleOfDeleteTargetsResponse) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *SingleOfDeleteTargetsResponse) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *SingleOfDeleteTargetsResponse) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *SingleOfDeleteTargetsResponse) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *SingleOfDeleteTargetsResponse) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *SingleOfDeleteTargetsResponse) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *SingleOfDeleteTargetsResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *SingleOfDeleteTargetsResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *SingleOfDeleteTargetsResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *SingleOfDeleteTargetsResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *SingleOfDeleteTargetsResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *SingleOfDeleteTargetsResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


