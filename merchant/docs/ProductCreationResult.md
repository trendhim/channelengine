# ProductCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RejectedCount** | Pointer to **int32** |  | [optional] 
**AcceptedCount** | Pointer to **int32** |  | [optional] 
**ProductMessages** | Pointer to [**[]ProductMessage**](ProductMessage.md) | Messages about the rejected products. | [optional] 

## Methods

### NewProductCreationResult

`func NewProductCreationResult() *ProductCreationResult`

NewProductCreationResult instantiates a new ProductCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreationResultWithDefaults

`func NewProductCreationResultWithDefaults() *ProductCreationResult`

NewProductCreationResultWithDefaults instantiates a new ProductCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRejectedCount

`func (o *ProductCreationResult) GetRejectedCount() int32`

GetRejectedCount returns the RejectedCount field if non-nil, zero value otherwise.

### GetRejectedCountOk

`func (o *ProductCreationResult) GetRejectedCountOk() (*int32, bool)`

GetRejectedCountOk returns a tuple with the RejectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCount

`func (o *ProductCreationResult) SetRejectedCount(v int32)`

SetRejectedCount sets RejectedCount field to given value.

### HasRejectedCount

`func (o *ProductCreationResult) HasRejectedCount() bool`

HasRejectedCount returns a boolean if a field has been set.

### GetAcceptedCount

`func (o *ProductCreationResult) GetAcceptedCount() int32`

GetAcceptedCount returns the AcceptedCount field if non-nil, zero value otherwise.

### GetAcceptedCountOk

`func (o *ProductCreationResult) GetAcceptedCountOk() (*int32, bool)`

GetAcceptedCountOk returns a tuple with the AcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCount

`func (o *ProductCreationResult) SetAcceptedCount(v int32)`

SetAcceptedCount sets AcceptedCount field to given value.

### HasAcceptedCount

`func (o *ProductCreationResult) HasAcceptedCount() bool`

HasAcceptedCount returns a boolean if a field has been set.

### GetProductMessages

`func (o *ProductCreationResult) GetProductMessages() []ProductMessage`

GetProductMessages returns the ProductMessages field if non-nil, zero value otherwise.

### GetProductMessagesOk

`func (o *ProductCreationResult) GetProductMessagesOk() (*[]ProductMessage, bool)`

GetProductMessagesOk returns a tuple with the ProductMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMessages

`func (o *ProductCreationResult) SetProductMessages(v []ProductMessage)`

SetProductMessages sets ProductMessages field to given value.

### HasProductMessages

`func (o *ProductCreationResult) HasProductMessages() bool`

HasProductMessages returns a boolean if a field has been set.

### SetProductMessagesNil

`func (o *ProductCreationResult) SetProductMessagesNil(b bool)`

 SetProductMessagesNil sets the value for ProductMessages to be an explicit nil

### UnsetProductMessages
`func (o *ProductCreationResult) UnsetProductMessages()`

UnsetProductMessages ensures that no value is present for ProductMessages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


