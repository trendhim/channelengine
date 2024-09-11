# MerchantCreateRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderIdentifier** | Pointer to **NullableString** |  | [optional] 
**MerchantRefundNo** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to [**RefundReason**](RefundReason.md) |  | [optional] 
**MerchantComment** | Pointer to **NullableString** |  | [optional] 
**AdditionalAmountInclTax** | Pointer to **float32** |  | [optional] 
**ShippingAmountInclTax** | Pointer to **float32** |  | [optional] 
**RefundDate** | Pointer to **time.Time** |  | [optional] 
**ExtraData** | Pointer to **map[string]string** |  | [optional] 
**Lines** | Pointer to [**[]MerchantCreateRefundLine**](MerchantCreateRefundLine.md) |  | [optional] 

## Methods

### NewMerchantCreateRefund

`func NewMerchantCreateRefund() *MerchantCreateRefund`

NewMerchantCreateRefund instantiates a new MerchantCreateRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateRefundWithDefaults

`func NewMerchantCreateRefundWithDefaults() *MerchantCreateRefund`

NewMerchantCreateRefundWithDefaults instantiates a new MerchantCreateRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderIdentifier

`func (o *MerchantCreateRefund) GetOrderIdentifier() string`

GetOrderIdentifier returns the OrderIdentifier field if non-nil, zero value otherwise.

### GetOrderIdentifierOk

`func (o *MerchantCreateRefund) GetOrderIdentifierOk() (*string, bool)`

GetOrderIdentifierOk returns a tuple with the OrderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIdentifier

`func (o *MerchantCreateRefund) SetOrderIdentifier(v string)`

SetOrderIdentifier sets OrderIdentifier field to given value.

### HasOrderIdentifier

`func (o *MerchantCreateRefund) HasOrderIdentifier() bool`

HasOrderIdentifier returns a boolean if a field has been set.

### SetOrderIdentifierNil

`func (o *MerchantCreateRefund) SetOrderIdentifierNil(b bool)`

 SetOrderIdentifierNil sets the value for OrderIdentifier to be an explicit nil

### UnsetOrderIdentifier
`func (o *MerchantCreateRefund) UnsetOrderIdentifier()`

UnsetOrderIdentifier ensures that no value is present for OrderIdentifier, not even an explicit nil
### GetMerchantRefundNo

`func (o *MerchantCreateRefund) GetMerchantRefundNo() string`

GetMerchantRefundNo returns the MerchantRefundNo field if non-nil, zero value otherwise.

### GetMerchantRefundNoOk

`func (o *MerchantCreateRefund) GetMerchantRefundNoOk() (*string, bool)`

GetMerchantRefundNoOk returns a tuple with the MerchantRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundNo

`func (o *MerchantCreateRefund) SetMerchantRefundNo(v string)`

SetMerchantRefundNo sets MerchantRefundNo field to given value.

### HasMerchantRefundNo

`func (o *MerchantCreateRefund) HasMerchantRefundNo() bool`

HasMerchantRefundNo returns a boolean if a field has been set.

### SetMerchantRefundNoNil

`func (o *MerchantCreateRefund) SetMerchantRefundNoNil(b bool)`

 SetMerchantRefundNoNil sets the value for MerchantRefundNo to be an explicit nil

### UnsetMerchantRefundNo
`func (o *MerchantCreateRefund) UnsetMerchantRefundNo()`

UnsetMerchantRefundNo ensures that no value is present for MerchantRefundNo, not even an explicit nil
### GetReason

`func (o *MerchantCreateRefund) GetReason() RefundReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantCreateRefund) GetReasonOk() (*RefundReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantCreateRefund) SetReason(v RefundReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantCreateRefund) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMerchantComment

`func (o *MerchantCreateRefund) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantCreateRefund) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantCreateRefund) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantCreateRefund) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantCreateRefund) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantCreateRefund) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetAdditionalAmountInclTax

`func (o *MerchantCreateRefund) GetAdditionalAmountInclTax() float32`

GetAdditionalAmountInclTax returns the AdditionalAmountInclTax field if non-nil, zero value otherwise.

### GetAdditionalAmountInclTaxOk

`func (o *MerchantCreateRefund) GetAdditionalAmountInclTaxOk() (*float32, bool)`

GetAdditionalAmountInclTaxOk returns a tuple with the AdditionalAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmountInclTax

`func (o *MerchantCreateRefund) SetAdditionalAmountInclTax(v float32)`

SetAdditionalAmountInclTax sets AdditionalAmountInclTax field to given value.

### HasAdditionalAmountInclTax

`func (o *MerchantCreateRefund) HasAdditionalAmountInclTax() bool`

HasAdditionalAmountInclTax returns a boolean if a field has been set.

### GetShippingAmountInclTax

`func (o *MerchantCreateRefund) GetShippingAmountInclTax() float32`

GetShippingAmountInclTax returns the ShippingAmountInclTax field if non-nil, zero value otherwise.

### GetShippingAmountInclTaxOk

`func (o *MerchantCreateRefund) GetShippingAmountInclTaxOk() (*float32, bool)`

GetShippingAmountInclTaxOk returns a tuple with the ShippingAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountInclTax

`func (o *MerchantCreateRefund) SetShippingAmountInclTax(v float32)`

SetShippingAmountInclTax sets ShippingAmountInclTax field to given value.

### HasShippingAmountInclTax

`func (o *MerchantCreateRefund) HasShippingAmountInclTax() bool`

HasShippingAmountInclTax returns a boolean if a field has been set.

### GetRefundDate

`func (o *MerchantCreateRefund) GetRefundDate() time.Time`

GetRefundDate returns the RefundDate field if non-nil, zero value otherwise.

### GetRefundDateOk

`func (o *MerchantCreateRefund) GetRefundDateOk() (*time.Time, bool)`

GetRefundDateOk returns a tuple with the RefundDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDate

`func (o *MerchantCreateRefund) SetRefundDate(v time.Time)`

SetRefundDate sets RefundDate field to given value.

### HasRefundDate

`func (o *MerchantCreateRefund) HasRefundDate() bool`

HasRefundDate returns a boolean if a field has been set.

### GetExtraData

`func (o *MerchantCreateRefund) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantCreateRefund) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantCreateRefund) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantCreateRefund) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantCreateRefund) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantCreateRefund) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetLines

`func (o *MerchantCreateRefund) GetLines() []MerchantCreateRefundLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantCreateRefund) GetLinesOk() (*[]MerchantCreateRefundLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantCreateRefund) SetLines(v []MerchantCreateRefundLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantCreateRefund) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantCreateRefund) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantCreateRefund) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


