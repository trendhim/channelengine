# MerchantHandleReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnIdentifier** | Pointer to **NullableString** |  | [optional] 
**ReturnLineIdentifier** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to [**ReturnHandlingAction**](ReturnHandlingAction.md) |  | [optional] 

## Methods

### NewMerchantHandleReturn

`func NewMerchantHandleReturn() *MerchantHandleReturn`

NewMerchantHandleReturn instantiates a new MerchantHandleReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantHandleReturnWithDefaults

`func NewMerchantHandleReturnWithDefaults() *MerchantHandleReturn`

NewMerchantHandleReturnWithDefaults instantiates a new MerchantHandleReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnIdentifier

`func (o *MerchantHandleReturn) GetReturnIdentifier() string`

GetReturnIdentifier returns the ReturnIdentifier field if non-nil, zero value otherwise.

### GetReturnIdentifierOk

`func (o *MerchantHandleReturn) GetReturnIdentifierOk() (*string, bool)`

GetReturnIdentifierOk returns a tuple with the ReturnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnIdentifier

`func (o *MerchantHandleReturn) SetReturnIdentifier(v string)`

SetReturnIdentifier sets ReturnIdentifier field to given value.

### HasReturnIdentifier

`func (o *MerchantHandleReturn) HasReturnIdentifier() bool`

HasReturnIdentifier returns a boolean if a field has been set.

### SetReturnIdentifierNil

`func (o *MerchantHandleReturn) SetReturnIdentifierNil(b bool)`

 SetReturnIdentifierNil sets the value for ReturnIdentifier to be an explicit nil

### UnsetReturnIdentifier
`func (o *MerchantHandleReturn) UnsetReturnIdentifier()`

UnsetReturnIdentifier ensures that no value is present for ReturnIdentifier, not even an explicit nil
### GetReturnLineIdentifier

`func (o *MerchantHandleReturn) GetReturnLineIdentifier() string`

GetReturnLineIdentifier returns the ReturnLineIdentifier field if non-nil, zero value otherwise.

### GetReturnLineIdentifierOk

`func (o *MerchantHandleReturn) GetReturnLineIdentifierOk() (*string, bool)`

GetReturnLineIdentifierOk returns a tuple with the ReturnLineIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineIdentifier

`func (o *MerchantHandleReturn) SetReturnLineIdentifier(v string)`

SetReturnLineIdentifier sets ReturnLineIdentifier field to given value.

### HasReturnLineIdentifier

`func (o *MerchantHandleReturn) HasReturnLineIdentifier() bool`

HasReturnLineIdentifier returns a boolean if a field has been set.

### SetReturnLineIdentifierNil

`func (o *MerchantHandleReturn) SetReturnLineIdentifierNil(b bool)`

 SetReturnLineIdentifierNil sets the value for ReturnLineIdentifier to be an explicit nil

### UnsetReturnLineIdentifier
`func (o *MerchantHandleReturn) UnsetReturnLineIdentifier()`

UnsetReturnLineIdentifier ensures that no value is present for ReturnLineIdentifier, not even an explicit nil
### GetQuantity

`func (o *MerchantHandleReturn) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantHandleReturn) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantHandleReturn) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MerchantHandleReturn) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAction

`func (o *MerchantHandleReturn) GetAction() ReturnHandlingAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MerchantHandleReturn) GetActionOk() (*ReturnHandlingAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MerchantHandleReturn) SetAction(v ReturnHandlingAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *MerchantHandleReturn) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


