# CollectionOfMerchantReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]MerchantReportsResponse**](MerchantReportsResponse.md) |  | [optional] 
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

### NewCollectionOfMerchantReportsResponse

`func NewCollectionOfMerchantReportsResponse() *CollectionOfMerchantReportsResponse`

NewCollectionOfMerchantReportsResponse instantiates a new CollectionOfMerchantReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfMerchantReportsResponseWithDefaults

`func NewCollectionOfMerchantReportsResponseWithDefaults() *CollectionOfMerchantReportsResponse`

NewCollectionOfMerchantReportsResponseWithDefaults instantiates a new CollectionOfMerchantReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionOfMerchantReportsResponse) GetContent() []MerchantReportsResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionOfMerchantReportsResponse) GetContentOk() (*[]MerchantReportsResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionOfMerchantReportsResponse) SetContent(v []MerchantReportsResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionOfMerchantReportsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CollectionOfMerchantReportsResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CollectionOfMerchantReportsResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCount

`func (o *CollectionOfMerchantReportsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionOfMerchantReportsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionOfMerchantReportsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionOfMerchantReportsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CollectionOfMerchantReportsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CollectionOfMerchantReportsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CollectionOfMerchantReportsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CollectionOfMerchantReportsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *CollectionOfMerchantReportsResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *CollectionOfMerchantReportsResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *CollectionOfMerchantReportsResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *CollectionOfMerchantReportsResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionOfMerchantReportsResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionOfMerchantReportsResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionOfMerchantReportsResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionOfMerchantReportsResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionOfMerchantReportsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionOfMerchantReportsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionOfMerchantReportsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionOfMerchantReportsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionOfMerchantReportsResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionOfMerchantReportsResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionOfMerchantReportsResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionOfMerchantReportsResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionOfMerchantReportsResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionOfMerchantReportsResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionOfMerchantReportsResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionOfMerchantReportsResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionOfMerchantReportsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionOfMerchantReportsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionOfMerchantReportsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionOfMerchantReportsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionOfMerchantReportsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionOfMerchantReportsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionOfMerchantReportsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionOfMerchantReportsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionOfMerchantReportsResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionOfMerchantReportsResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValidationErrors

`func (o *CollectionOfMerchantReportsResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionOfMerchantReportsResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionOfMerchantReportsResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionOfMerchantReportsResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionOfMerchantReportsResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionOfMerchantReportsResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


