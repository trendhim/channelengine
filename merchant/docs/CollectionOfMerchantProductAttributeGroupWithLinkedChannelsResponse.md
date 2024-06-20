# CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]MerchantProductAttributeGroupWithLinkedChannelsResponse**](MerchantProductAttributeGroupWithLinkedChannelsResponse.md) |  | [optional] 
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

### NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse

`func NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse() *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse`

NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse instantiates a new CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponseWithDefaults

`func NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponseWithDefaults() *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse`

NewCollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponseWithDefaults instantiates a new CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetContent() []MerchantProductAttributeGroupWithLinkedChannelsResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetContentOk() (*[]MerchantProductAttributeGroupWithLinkedChannelsResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetContent(v []MerchantProductAttributeGroupWithLinkedChannelsResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStatusCode

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetLogId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetSuccess

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExceptionType

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetValidationErrors

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### SetValidationErrorsNil

`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) SetValidationErrorsNil(b bool)`

 SetValidationErrorsNil sets the value for ValidationErrors to be an explicit nil

### UnsetValidationErrors
`func (o *CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse) UnsetValidationErrors()`

UnsetValidationErrors ensures that no value is present for ValidationErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


