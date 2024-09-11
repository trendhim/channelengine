# IReturnLineHandlingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ReturnLineId** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**ReturnHandlingAction**](ReturnHandlingAction.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewIReturnLineHandlingResult

`func NewIReturnLineHandlingResult() *IReturnLineHandlingResult`

NewIReturnLineHandlingResult instantiates a new IReturnLineHandlingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIReturnLineHandlingResultWithDefaults

`func NewIReturnLineHandlingResultWithDefaults() *IReturnLineHandlingResult`

NewIReturnLineHandlingResultWithDefaults instantiates a new IReturnLineHandlingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IReturnLineHandlingResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IReturnLineHandlingResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IReturnLineHandlingResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IReturnLineHandlingResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReturnLineId

`func (o *IReturnLineHandlingResult) GetReturnLineId() int32`

GetReturnLineId returns the ReturnLineId field if non-nil, zero value otherwise.

### GetReturnLineIdOk

`func (o *IReturnLineHandlingResult) GetReturnLineIdOk() (*int32, bool)`

GetReturnLineIdOk returns a tuple with the ReturnLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineId

`func (o *IReturnLineHandlingResult) SetReturnLineId(v int32)`

SetReturnLineId sets ReturnLineId field to given value.

### HasReturnLineId

`func (o *IReturnLineHandlingResult) HasReturnLineId() bool`

HasReturnLineId returns a boolean if a field has been set.

### GetQuantity

`func (o *IReturnLineHandlingResult) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IReturnLineHandlingResult) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IReturnLineHandlingResult) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IReturnLineHandlingResult) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAction

`func (o *IReturnLineHandlingResult) GetAction() ReturnHandlingAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IReturnLineHandlingResult) GetActionOk() (*ReturnHandlingAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IReturnLineHandlingResult) SetAction(v ReturnHandlingAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *IReturnLineHandlingResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IReturnLineHandlingResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IReturnLineHandlingResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IReturnLineHandlingResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IReturnLineHandlingResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IReturnLineHandlingResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IReturnLineHandlingResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IReturnLineHandlingResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IReturnLineHandlingResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IReturnLineHandlingResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IReturnLineHandlingResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IReturnLineHandlingResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IReturnLineHandlingResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *IReturnLineHandlingResult) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *IReturnLineHandlingResult) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


