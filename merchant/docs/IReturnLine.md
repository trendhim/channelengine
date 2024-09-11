# IReturnLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ReturnId** | Pointer to **int32** |  | [optional] 
**OrderLineId** | Pointer to **int32** |  | [optional] 
**MerchantReturnLineNo** | Pointer to **NullableString** |  | [optional] 
**ChannelReturnLineNo** | Pointer to **NullableString** |  | [optional] 
**MerchantOrderLineNo** | Pointer to **NullableString** |  | [optional] 
**ChannelOrderLineNo** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**HandlingResults** | Pointer to [**[]IReturnLineHandlingResult**](IReturnLineHandlingResult.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIReturnLine

`func NewIReturnLine() *IReturnLine`

NewIReturnLine instantiates a new IReturnLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIReturnLineWithDefaults

`func NewIReturnLineWithDefaults() *IReturnLine`

NewIReturnLineWithDefaults instantiates a new IReturnLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IReturnLine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IReturnLine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IReturnLine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IReturnLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReturnId

`func (o *IReturnLine) GetReturnId() int32`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *IReturnLine) GetReturnIdOk() (*int32, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *IReturnLine) SetReturnId(v int32)`

SetReturnId sets ReturnId field to given value.

### HasReturnId

`func (o *IReturnLine) HasReturnId() bool`

HasReturnId returns a boolean if a field has been set.

### GetOrderLineId

`func (o *IReturnLine) GetOrderLineId() int32`

GetOrderLineId returns the OrderLineId field if non-nil, zero value otherwise.

### GetOrderLineIdOk

`func (o *IReturnLine) GetOrderLineIdOk() (*int32, bool)`

GetOrderLineIdOk returns a tuple with the OrderLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineId

`func (o *IReturnLine) SetOrderLineId(v int32)`

SetOrderLineId sets OrderLineId field to given value.

### HasOrderLineId

`func (o *IReturnLine) HasOrderLineId() bool`

HasOrderLineId returns a boolean if a field has been set.

### GetMerchantReturnLineNo

`func (o *IReturnLine) GetMerchantReturnLineNo() string`

GetMerchantReturnLineNo returns the MerchantReturnLineNo field if non-nil, zero value otherwise.

### GetMerchantReturnLineNoOk

`func (o *IReturnLine) GetMerchantReturnLineNoOk() (*string, bool)`

GetMerchantReturnLineNoOk returns a tuple with the MerchantReturnLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnLineNo

`func (o *IReturnLine) SetMerchantReturnLineNo(v string)`

SetMerchantReturnLineNo sets MerchantReturnLineNo field to given value.

### HasMerchantReturnLineNo

`func (o *IReturnLine) HasMerchantReturnLineNo() bool`

HasMerchantReturnLineNo returns a boolean if a field has been set.

### SetMerchantReturnLineNoNil

`func (o *IReturnLine) SetMerchantReturnLineNoNil(b bool)`

 SetMerchantReturnLineNoNil sets the value for MerchantReturnLineNo to be an explicit nil

### UnsetMerchantReturnLineNo
`func (o *IReturnLine) UnsetMerchantReturnLineNo()`

UnsetMerchantReturnLineNo ensures that no value is present for MerchantReturnLineNo, not even an explicit nil
### GetChannelReturnLineNo

`func (o *IReturnLine) GetChannelReturnLineNo() string`

GetChannelReturnLineNo returns the ChannelReturnLineNo field if non-nil, zero value otherwise.

### GetChannelReturnLineNoOk

`func (o *IReturnLine) GetChannelReturnLineNoOk() (*string, bool)`

GetChannelReturnLineNoOk returns a tuple with the ChannelReturnLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnLineNo

`func (o *IReturnLine) SetChannelReturnLineNo(v string)`

SetChannelReturnLineNo sets ChannelReturnLineNo field to given value.

### HasChannelReturnLineNo

`func (o *IReturnLine) HasChannelReturnLineNo() bool`

HasChannelReturnLineNo returns a boolean if a field has been set.

### SetChannelReturnLineNoNil

`func (o *IReturnLine) SetChannelReturnLineNoNil(b bool)`

 SetChannelReturnLineNoNil sets the value for ChannelReturnLineNo to be an explicit nil

### UnsetChannelReturnLineNo
`func (o *IReturnLine) UnsetChannelReturnLineNo()`

UnsetChannelReturnLineNo ensures that no value is present for ChannelReturnLineNo, not even an explicit nil
### GetMerchantOrderLineNo

`func (o *IReturnLine) GetMerchantOrderLineNo() string`

GetMerchantOrderLineNo returns the MerchantOrderLineNo field if non-nil, zero value otherwise.

### GetMerchantOrderLineNoOk

`func (o *IReturnLine) GetMerchantOrderLineNoOk() (*string, bool)`

GetMerchantOrderLineNoOk returns a tuple with the MerchantOrderLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderLineNo

`func (o *IReturnLine) SetMerchantOrderLineNo(v string)`

SetMerchantOrderLineNo sets MerchantOrderLineNo field to given value.

### HasMerchantOrderLineNo

`func (o *IReturnLine) HasMerchantOrderLineNo() bool`

HasMerchantOrderLineNo returns a boolean if a field has been set.

### SetMerchantOrderLineNoNil

`func (o *IReturnLine) SetMerchantOrderLineNoNil(b bool)`

 SetMerchantOrderLineNoNil sets the value for MerchantOrderLineNo to be an explicit nil

### UnsetMerchantOrderLineNo
`func (o *IReturnLine) UnsetMerchantOrderLineNo()`

UnsetMerchantOrderLineNo ensures that no value is present for MerchantOrderLineNo, not even an explicit nil
### GetChannelOrderLineNo

`func (o *IReturnLine) GetChannelOrderLineNo() string`

GetChannelOrderLineNo returns the ChannelOrderLineNo field if non-nil, zero value otherwise.

### GetChannelOrderLineNoOk

`func (o *IReturnLine) GetChannelOrderLineNoOk() (*string, bool)`

GetChannelOrderLineNoOk returns a tuple with the ChannelOrderLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderLineNo

`func (o *IReturnLine) SetChannelOrderLineNo(v string)`

SetChannelOrderLineNo sets ChannelOrderLineNo field to given value.

### HasChannelOrderLineNo

`func (o *IReturnLine) HasChannelOrderLineNo() bool`

HasChannelOrderLineNo returns a boolean if a field has been set.

### SetChannelOrderLineNoNil

`func (o *IReturnLine) SetChannelOrderLineNoNil(b bool)`

 SetChannelOrderLineNoNil sets the value for ChannelOrderLineNo to be an explicit nil

### UnsetChannelOrderLineNo
`func (o *IReturnLine) UnsetChannelOrderLineNo()`

UnsetChannelOrderLineNo ensures that no value is present for ChannelOrderLineNo, not even an explicit nil
### GetQuantity

`func (o *IReturnLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IReturnLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IReturnLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IReturnLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetHandlingResults

`func (o *IReturnLine) GetHandlingResults() []IReturnLineHandlingResult`

GetHandlingResults returns the HandlingResults field if non-nil, zero value otherwise.

### GetHandlingResultsOk

`func (o *IReturnLine) GetHandlingResultsOk() (*[]IReturnLineHandlingResult, bool)`

GetHandlingResultsOk returns a tuple with the HandlingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingResults

`func (o *IReturnLine) SetHandlingResults(v []IReturnLineHandlingResult)`

SetHandlingResults sets HandlingResults field to given value.

### HasHandlingResults

`func (o *IReturnLine) HasHandlingResults() bool`

HasHandlingResults returns a boolean if a field has been set.

### SetHandlingResultsNil

`func (o *IReturnLine) SetHandlingResultsNil(b bool)`

 SetHandlingResultsNil sets the value for HandlingResults to be an explicit nil

### UnsetHandlingResults
`func (o *IReturnLine) UnsetHandlingResults()`

UnsetHandlingResults ensures that no value is present for HandlingResults, not even an explicit nil
### GetCreatedAt

`func (o *IReturnLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IReturnLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IReturnLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IReturnLine) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IReturnLine) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IReturnLine) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IReturnLine) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IReturnLine) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IReturnLine) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IReturnLine) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IReturnLine) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IReturnLine) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *IReturnLine) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *IReturnLine) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetMerchantProductNo

`func (o *IReturnLine) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *IReturnLine) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *IReturnLine) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *IReturnLine) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *IReturnLine) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *IReturnLine) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetProductName

`func (o *IReturnLine) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IReturnLine) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IReturnLine) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IReturnLine) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *IReturnLine) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *IReturnLine) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetProductId

`func (o *IReturnLine) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *IReturnLine) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *IReturnLine) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *IReturnLine) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *IReturnLine) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *IReturnLine) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


