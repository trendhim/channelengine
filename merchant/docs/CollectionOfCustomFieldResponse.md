# CollectionOfCustomFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CustomFieldResponse**](CustomFieldResponse.md) |  | [optional] 
**Count** | Pointer to **int32** | The number of items in the current response. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items. | [optional] 
**ItemsPerPage** | Pointer to **int32** | The number of items per page. | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCollectionOfCustomFieldResponse

`func NewCollectionOfCustomFieldResponse() *CollectionOfCustomFieldResponse`

NewCollectionOfCustomFieldResponse instantiates a new CollectionOfCustomFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfCustomFieldResponseWithDefaults

`func NewCollectionOfCustomFieldResponseWithDefaults() *CollectionOfCustomFieldResponse`

NewCollectionOfCustomFieldResponseWithDefaults instantiates a new CollectionOfCustomFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionOfCustomFieldResponse) GetContent() []CustomFieldResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionOfCustomFieldResponse) GetContentOk() (*[]CustomFieldResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionOfCustomFieldResponse) SetContent(v []CustomFieldResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionOfCustomFieldResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CollectionOfCustomFieldResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CollectionOfCustomFieldResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCount

`func (o *CollectionOfCustomFieldResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionOfCustomFieldResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionOfCustomFieldResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionOfCustomFieldResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CollectionOfCustomFieldResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CollectionOfCustomFieldResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CollectionOfCustomFieldResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CollectionOfCustomFieldResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *CollectionOfCustomFieldResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *CollectionOfCustomFieldResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *CollectionOfCustomFieldResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *CollectionOfCustomFieldResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionOfCustomFieldResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionOfCustomFieldResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionOfCustomFieldResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionOfCustomFieldResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionOfCustomFieldResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionOfCustomFieldResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionOfCustomFieldResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionOfCustomFieldResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionOfCustomFieldResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionOfCustomFieldResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionOfCustomFieldResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionOfCustomFieldResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionOfCustomFieldResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionOfCustomFieldResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionOfCustomFieldResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionOfCustomFieldResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionOfCustomFieldResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionOfCustomFieldResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionOfCustomFieldResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionOfCustomFieldResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionOfCustomFieldResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionOfCustomFieldResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionOfCustomFieldResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionOfCustomFieldResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionOfCustomFieldResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionOfCustomFieldResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *CollectionOfCustomFieldResponse) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *CollectionOfCustomFieldResponse) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *CollectionOfCustomFieldResponse) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *CollectionOfCustomFieldResponse) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *CollectionOfCustomFieldResponse) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *CollectionOfCustomFieldResponse) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *CollectionOfCustomFieldResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionOfCustomFieldResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionOfCustomFieldResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionOfCustomFieldResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionOfCustomFieldResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionOfCustomFieldResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


