# BulkReturnIdentifierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | Pointer to [**ReturnByFilterIdentifier**](ReturnByFilterIdentifier.md) |  | [optional] 
**Models** | Pointer to **[]string** | The value (of the selected type) to filter on | [optional] 

## Methods

### NewBulkReturnIdentifierRequest

`func NewBulkReturnIdentifierRequest() *BulkReturnIdentifierRequest`

NewBulkReturnIdentifierRequest instantiates a new BulkReturnIdentifierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReturnIdentifierRequestWithDefaults

`func NewBulkReturnIdentifierRequestWithDefaults() *BulkReturnIdentifierRequest`

NewBulkReturnIdentifierRequestWithDefaults instantiates a new BulkReturnIdentifierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *BulkReturnIdentifierRequest) GetIdentifierType() ReturnByFilterIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *BulkReturnIdentifierRequest) GetIdentifierTypeOk() (*ReturnByFilterIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *BulkReturnIdentifierRequest) SetIdentifierType(v ReturnByFilterIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *BulkReturnIdentifierRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModels

`func (o *BulkReturnIdentifierRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *BulkReturnIdentifierRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *BulkReturnIdentifierRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *BulkReturnIdentifierRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *BulkReturnIdentifierRequest) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *BulkReturnIdentifierRequest) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


