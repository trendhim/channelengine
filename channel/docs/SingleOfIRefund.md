# SingleOfIRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**IRefund**](IRefund.md) |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewSingleOfIRefund

`func NewSingleOfIRefund() *SingleOfIRefund`

NewSingleOfIRefund instantiates a new SingleOfIRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleOfIRefundWithDefaults

`func NewSingleOfIRefundWithDefaults() *SingleOfIRefund`

NewSingleOfIRefundWithDefaults instantiates a new SingleOfIRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SingleOfIRefund) GetContent() IRefund`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SingleOfIRefund) GetContentOk() (*IRefund, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SingleOfIRefund) SetContent(v IRefund)`

SetContent sets Content field to given value.

### HasContent

`func (o *SingleOfIRefund) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatusCode

`func (o *SingleOfIRefund) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SingleOfIRefund) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SingleOfIRefund) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SingleOfIRefund) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *SingleOfIRefund) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SingleOfIRefund) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SingleOfIRefund) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SingleOfIRefund) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *SingleOfIRefund) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SingleOfIRefund) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *SingleOfIRefund) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *SingleOfIRefund) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *SingleOfIRefund) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *SingleOfIRefund) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *SingleOfIRefund) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *SingleOfIRefund) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *SingleOfIRefund) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SingleOfIRefund) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SingleOfIRefund) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SingleOfIRefund) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *SingleOfIRefund) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SingleOfIRefund) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SingleOfIRefund) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SingleOfIRefund) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SingleOfIRefund) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SingleOfIRefund) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *SingleOfIRefund) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *SingleOfIRefund) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *SingleOfIRefund) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *SingleOfIRefund) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *SingleOfIRefund) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *SingleOfIRefund) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *SingleOfIRefund) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *SingleOfIRefund) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *SingleOfIRefund) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *SingleOfIRefund) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *SingleOfIRefund) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *SingleOfIRefund) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


