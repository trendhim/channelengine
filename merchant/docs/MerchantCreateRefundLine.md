# MerchantCreateRefundLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderLineIdentifier** | Pointer to **NullableString** |  | [optional] 
**LineAmountInclTax** | Pointer to **float32** |  | [optional] 
**MerchantRefundLineNo** | Pointer to **NullableString** |  | [optional] 
**ExtraData** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMerchantCreateRefundLine

`func NewMerchantCreateRefundLine() *MerchantCreateRefundLine`

NewMerchantCreateRefundLine instantiates a new MerchantCreateRefundLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateRefundLineWithDefaults

`func NewMerchantCreateRefundLineWithDefaults() *MerchantCreateRefundLine`

NewMerchantCreateRefundLineWithDefaults instantiates a new MerchantCreateRefundLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderLineIdentifier

`func (o *MerchantCreateRefundLine) GetOrderLineIdentifier() string`

GetOrderLineIdentifier returns the OrderLineIdentifier field if non-nil, zero value otherwise.

### GetOrderLineIdentifierOk

`func (o *MerchantCreateRefundLine) GetOrderLineIdentifierOk() (*string, bool)`

GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineIdentifier

`func (o *MerchantCreateRefundLine) SetOrderLineIdentifier(v string)`

SetOrderLineIdentifier sets OrderLineIdentifier field to given value.

### HasOrderLineIdentifier

`func (o *MerchantCreateRefundLine) HasOrderLineIdentifier() bool`

HasOrderLineIdentifier returns a boolean if a field has been set.

### SetOrderLineIdentifierNil

`func (o *MerchantCreateRefundLine) SetOrderLineIdentifierNil(b bool)`

 SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil

### UnsetOrderLineIdentifier
`func (o *MerchantCreateRefundLine) UnsetOrderLineIdentifier()`

UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
### GetLineAmountInclTax

`func (o *MerchantCreateRefundLine) GetLineAmountInclTax() float32`

GetLineAmountInclTax returns the LineAmountInclTax field if non-nil, zero value otherwise.

### GetLineAmountInclTaxOk

`func (o *MerchantCreateRefundLine) GetLineAmountInclTaxOk() (*float32, bool)`

GetLineAmountInclTaxOk returns a tuple with the LineAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountInclTax

`func (o *MerchantCreateRefundLine) SetLineAmountInclTax(v float32)`

SetLineAmountInclTax sets LineAmountInclTax field to given value.

### HasLineAmountInclTax

`func (o *MerchantCreateRefundLine) HasLineAmountInclTax() bool`

HasLineAmountInclTax returns a boolean if a field has been set.

### GetMerchantRefundLineNo

`func (o *MerchantCreateRefundLine) GetMerchantRefundLineNo() string`

GetMerchantRefundLineNo returns the MerchantRefundLineNo field if non-nil, zero value otherwise.

### GetMerchantRefundLineNoOk

`func (o *MerchantCreateRefundLine) GetMerchantRefundLineNoOk() (*string, bool)`

GetMerchantRefundLineNoOk returns a tuple with the MerchantRefundLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundLineNo

`func (o *MerchantCreateRefundLine) SetMerchantRefundLineNo(v string)`

SetMerchantRefundLineNo sets MerchantRefundLineNo field to given value.

### HasMerchantRefundLineNo

`func (o *MerchantCreateRefundLine) HasMerchantRefundLineNo() bool`

HasMerchantRefundLineNo returns a boolean if a field has been set.

### SetMerchantRefundLineNoNil

`func (o *MerchantCreateRefundLine) SetMerchantRefundLineNoNil(b bool)`

 SetMerchantRefundLineNoNil sets the value for MerchantRefundLineNo to be an explicit nil

### UnsetMerchantRefundLineNo
`func (o *MerchantCreateRefundLine) UnsetMerchantRefundLineNo()`

UnsetMerchantRefundLineNo ensures that no value is present for MerchantRefundLineNo, not even an explicit nil
### GetExtraData

`func (o *MerchantCreateRefundLine) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantCreateRefundLine) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantCreateRefundLine) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantCreateRefundLine) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantCreateRefundLine) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantCreateRefundLine) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


