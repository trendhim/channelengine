# SingleOfListOfTargetResponseVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]TargetResponseVm**](TargetResponseVm.md) |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewSingleOfListOfTargetResponseVm

`func NewSingleOfListOfTargetResponseVm() *SingleOfListOfTargetResponseVm`

NewSingleOfListOfTargetResponseVm instantiates a new SingleOfListOfTargetResponseVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleOfListOfTargetResponseVmWithDefaults

`func NewSingleOfListOfTargetResponseVmWithDefaults() *SingleOfListOfTargetResponseVm`

NewSingleOfListOfTargetResponseVmWithDefaults instantiates a new SingleOfListOfTargetResponseVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SingleOfListOfTargetResponseVm) GetContent() []TargetResponseVm`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SingleOfListOfTargetResponseVm) GetContentOk() (*[]TargetResponseVm, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SingleOfListOfTargetResponseVm) SetContent(v []TargetResponseVm)`

SetContent sets Content field to given value.

### HasContent

`func (o *SingleOfListOfTargetResponseVm) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SingleOfListOfTargetResponseVm) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SingleOfListOfTargetResponseVm) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetStatusCode

`func (o *SingleOfListOfTargetResponseVm) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SingleOfListOfTargetResponseVm) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SingleOfListOfTargetResponseVm) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SingleOfListOfTargetResponseVm) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *SingleOfListOfTargetResponseVm) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SingleOfListOfTargetResponseVm) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SingleOfListOfTargetResponseVm) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SingleOfListOfTargetResponseVm) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *SingleOfListOfTargetResponseVm) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SingleOfListOfTargetResponseVm) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *SingleOfListOfTargetResponseVm) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *SingleOfListOfTargetResponseVm) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *SingleOfListOfTargetResponseVm) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *SingleOfListOfTargetResponseVm) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *SingleOfListOfTargetResponseVm) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *SingleOfListOfTargetResponseVm) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *SingleOfListOfTargetResponseVm) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SingleOfListOfTargetResponseVm) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SingleOfListOfTargetResponseVm) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SingleOfListOfTargetResponseVm) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *SingleOfListOfTargetResponseVm) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SingleOfListOfTargetResponseVm) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SingleOfListOfTargetResponseVm) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SingleOfListOfTargetResponseVm) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SingleOfListOfTargetResponseVm) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SingleOfListOfTargetResponseVm) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *SingleOfListOfTargetResponseVm) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *SingleOfListOfTargetResponseVm) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *SingleOfListOfTargetResponseVm) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *SingleOfListOfTargetResponseVm) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *SingleOfListOfTargetResponseVm) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *SingleOfListOfTargetResponseVm) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *SingleOfListOfTargetResponseVm) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *SingleOfListOfTargetResponseVm) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *SingleOfListOfTargetResponseVm) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *SingleOfListOfTargetResponseVm) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *SingleOfListOfTargetResponseVm) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *SingleOfListOfTargetResponseVm) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


