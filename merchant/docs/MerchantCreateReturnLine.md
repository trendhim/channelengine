# MerchantCreateReturnLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderLineIdentifier** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**MerchantReturnLineNo** | Pointer to **NullableString** |  | [optional] 
**ExtraData** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMerchantCreateReturnLine

`func NewMerchantCreateReturnLine() *MerchantCreateReturnLine`

NewMerchantCreateReturnLine instantiates a new MerchantCreateReturnLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateReturnLineWithDefaults

`func NewMerchantCreateReturnLineWithDefaults() *MerchantCreateReturnLine`

NewMerchantCreateReturnLineWithDefaults instantiates a new MerchantCreateReturnLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderLineIdentifier

`func (o *MerchantCreateReturnLine) GetOrderLineIdentifier() string`

GetOrderLineIdentifier returns the OrderLineIdentifier field if non-nil, zero value otherwise.

### GetOrderLineIdentifierOk

`func (o *MerchantCreateReturnLine) GetOrderLineIdentifierOk() (*string, bool)`

GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineIdentifier

`func (o *MerchantCreateReturnLine) SetOrderLineIdentifier(v string)`

SetOrderLineIdentifier sets OrderLineIdentifier field to given value.

### HasOrderLineIdentifier

`func (o *MerchantCreateReturnLine) HasOrderLineIdentifier() bool`

HasOrderLineIdentifier returns a boolean if a field has been set.

### SetOrderLineIdentifierNil

`func (o *MerchantCreateReturnLine) SetOrderLineIdentifierNil(b bool)`

 SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil

### UnsetOrderLineIdentifier
`func (o *MerchantCreateReturnLine) UnsetOrderLineIdentifier()`

UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
### GetQuantity

`func (o *MerchantCreateReturnLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantCreateReturnLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantCreateReturnLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MerchantCreateReturnLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMerchantReturnLineNo

`func (o *MerchantCreateReturnLine) GetMerchantReturnLineNo() string`

GetMerchantReturnLineNo returns the MerchantReturnLineNo field if non-nil, zero value otherwise.

### GetMerchantReturnLineNoOk

`func (o *MerchantCreateReturnLine) GetMerchantReturnLineNoOk() (*string, bool)`

GetMerchantReturnLineNoOk returns a tuple with the MerchantReturnLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnLineNo

`func (o *MerchantCreateReturnLine) SetMerchantReturnLineNo(v string)`

SetMerchantReturnLineNo sets MerchantReturnLineNo field to given value.

### HasMerchantReturnLineNo

`func (o *MerchantCreateReturnLine) HasMerchantReturnLineNo() bool`

HasMerchantReturnLineNo returns a boolean if a field has been set.

### SetMerchantReturnLineNoNil

`func (o *MerchantCreateReturnLine) SetMerchantReturnLineNoNil(b bool)`

 SetMerchantReturnLineNoNil sets the value for MerchantReturnLineNo to be an explicit nil

### UnsetMerchantReturnLineNo
`func (o *MerchantCreateReturnLine) UnsetMerchantReturnLineNo()`

UnsetMerchantReturnLineNo ensures that no value is present for MerchantReturnLineNo, not even an explicit nil
### GetExtraData

`func (o *MerchantCreateReturnLine) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantCreateReturnLine) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantCreateReturnLine) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantCreateReturnLine) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantCreateReturnLine) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantCreateReturnLine) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


