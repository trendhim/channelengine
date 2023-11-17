# CollectionOfChannelCancellationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]ChannelCancellationResponse**](ChannelCancellationResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**LogId** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCollectionOfChannelCancellationResponse

`func NewCollectionOfChannelCancellationResponse() *CollectionOfChannelCancellationResponse`

NewCollectionOfChannelCancellationResponse instantiates a new CollectionOfChannelCancellationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfChannelCancellationResponseWithDefaults

`func NewCollectionOfChannelCancellationResponseWithDefaults() *CollectionOfChannelCancellationResponse`

NewCollectionOfChannelCancellationResponseWithDefaults instantiates a new CollectionOfChannelCancellationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionOfChannelCancellationResponse) GetContent() []ChannelCancellationResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionOfChannelCancellationResponse) GetContentOk() (*[]ChannelCancellationResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionOfChannelCancellationResponse) SetContent(v []ChannelCancellationResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionOfChannelCancellationResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CollectionOfChannelCancellationResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CollectionOfChannelCancellationResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCount

`func (o *CollectionOfChannelCancellationResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionOfChannelCancellationResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionOfChannelCancellationResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionOfChannelCancellationResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CollectionOfChannelCancellationResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CollectionOfChannelCancellationResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CollectionOfChannelCancellationResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CollectionOfChannelCancellationResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *CollectionOfChannelCancellationResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *CollectionOfChannelCancellationResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *CollectionOfChannelCancellationResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *CollectionOfChannelCancellationResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionOfChannelCancellationResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionOfChannelCancellationResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionOfChannelCancellationResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionOfChannelCancellationResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionOfChannelCancellationResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionOfChannelCancellationResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionOfChannelCancellationResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionOfChannelCancellationResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionOfChannelCancellationResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionOfChannelCancellationResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionOfChannelCancellationResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionOfChannelCancellationResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionOfChannelCancellationResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionOfChannelCancellationResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionOfChannelCancellationResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionOfChannelCancellationResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionOfChannelCancellationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionOfChannelCancellationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionOfChannelCancellationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionOfChannelCancellationResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionOfChannelCancellationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionOfChannelCancellationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionOfChannelCancellationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionOfChannelCancellationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionOfChannelCancellationResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionOfChannelCancellationResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValidationErrors

`func (o *CollectionOfChannelCancellationResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionOfChannelCancellationResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionOfChannelCancellationResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionOfChannelCancellationResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionOfChannelCancellationResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionOfChannelCancellationResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


