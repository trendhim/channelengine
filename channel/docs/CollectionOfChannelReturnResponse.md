# CollectionOfChannelReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]ChannelReturnResponse**](ChannelReturnResponse.md) |  | [optional] 
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

### NewCollectionOfChannelReturnResponse

`func NewCollectionOfChannelReturnResponse() *CollectionOfChannelReturnResponse`

NewCollectionOfChannelReturnResponse instantiates a new CollectionOfChannelReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfChannelReturnResponseWithDefaults

`func NewCollectionOfChannelReturnResponseWithDefaults() *CollectionOfChannelReturnResponse`

NewCollectionOfChannelReturnResponseWithDefaults instantiates a new CollectionOfChannelReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionOfChannelReturnResponse) GetContent() []ChannelReturnResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionOfChannelReturnResponse) GetContentOk() (*[]ChannelReturnResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionOfChannelReturnResponse) SetContent(v []ChannelReturnResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionOfChannelReturnResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CollectionOfChannelReturnResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CollectionOfChannelReturnResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCount

`func (o *CollectionOfChannelReturnResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionOfChannelReturnResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionOfChannelReturnResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionOfChannelReturnResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CollectionOfChannelReturnResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CollectionOfChannelReturnResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CollectionOfChannelReturnResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CollectionOfChannelReturnResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *CollectionOfChannelReturnResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *CollectionOfChannelReturnResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *CollectionOfChannelReturnResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *CollectionOfChannelReturnResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionOfChannelReturnResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionOfChannelReturnResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionOfChannelReturnResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionOfChannelReturnResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionOfChannelReturnResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionOfChannelReturnResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionOfChannelReturnResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionOfChannelReturnResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionOfChannelReturnResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionOfChannelReturnResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionOfChannelReturnResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionOfChannelReturnResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionOfChannelReturnResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionOfChannelReturnResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionOfChannelReturnResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionOfChannelReturnResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionOfChannelReturnResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionOfChannelReturnResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionOfChannelReturnResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionOfChannelReturnResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionOfChannelReturnResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionOfChannelReturnResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionOfChannelReturnResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionOfChannelReturnResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionOfChannelReturnResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionOfChannelReturnResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValidationErrors

`func (o *CollectionOfChannelReturnResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionOfChannelReturnResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionOfChannelReturnResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionOfChannelReturnResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionOfChannelReturnResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionOfChannelReturnResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


