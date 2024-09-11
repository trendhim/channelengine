# IRefundLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LineAmountInclTax** | Pointer to **float32** |  | [optional] 
**OriginalLineAmountInclTax** | Pointer to **float32** |  | [optional] 
**LineTotalInclTax** | Pointer to **float32** |  | [optional] 
**OriginalLineTotalInclTax** | Pointer to **float32** |  | [optional] 
**OriginalUnitTax** | Pointer to **float32** |  | [optional] 
**RefundId** | Pointer to **int32** |  | [optional] 
**OrderLineId** | Pointer to **int32** |  | [optional] 
**ReturnLineId** | Pointer to **NullableInt32** |  | [optional] 
**ChannelOrderLineNo** | Pointer to **NullableString** |  | [optional] 
**MerchantRefundLineNo** | Pointer to **NullableString** |  | [optional] 
**ChannelProductNo** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIRefundLine

`func NewIRefundLine() *IRefundLine`

NewIRefundLine instantiates a new IRefundLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIRefundLineWithDefaults

`func NewIRefundLineWithDefaults() *IRefundLine`

NewIRefundLineWithDefaults instantiates a new IRefundLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IRefundLine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IRefundLine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IRefundLine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IRefundLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineAmountInclTax

`func (o *IRefundLine) GetLineAmountInclTax() float32`

GetLineAmountInclTax returns the LineAmountInclTax field if non-nil, zero value otherwise.

### GetLineAmountInclTaxOk

`func (o *IRefundLine) GetLineAmountInclTaxOk() (*float32, bool)`

GetLineAmountInclTaxOk returns a tuple with the LineAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountInclTax

`func (o *IRefundLine) SetLineAmountInclTax(v float32)`

SetLineAmountInclTax sets LineAmountInclTax field to given value.

### HasLineAmountInclTax

`func (o *IRefundLine) HasLineAmountInclTax() bool`

HasLineAmountInclTax returns a boolean if a field has been set.

### GetOriginalLineAmountInclTax

`func (o *IRefundLine) GetOriginalLineAmountInclTax() float32`

GetOriginalLineAmountInclTax returns the OriginalLineAmountInclTax field if non-nil, zero value otherwise.

### GetOriginalLineAmountInclTaxOk

`func (o *IRefundLine) GetOriginalLineAmountInclTaxOk() (*float32, bool)`

GetOriginalLineAmountInclTaxOk returns a tuple with the OriginalLineAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineAmountInclTax

`func (o *IRefundLine) SetOriginalLineAmountInclTax(v float32)`

SetOriginalLineAmountInclTax sets OriginalLineAmountInclTax field to given value.

### HasOriginalLineAmountInclTax

`func (o *IRefundLine) HasOriginalLineAmountInclTax() bool`

HasOriginalLineAmountInclTax returns a boolean if a field has been set.

### GetLineTotalInclTax

`func (o *IRefundLine) GetLineTotalInclTax() float32`

GetLineTotalInclTax returns the LineTotalInclTax field if non-nil, zero value otherwise.

### GetLineTotalInclTaxOk

`func (o *IRefundLine) GetLineTotalInclTaxOk() (*float32, bool)`

GetLineTotalInclTaxOk returns a tuple with the LineTotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalInclTax

`func (o *IRefundLine) SetLineTotalInclTax(v float32)`

SetLineTotalInclTax sets LineTotalInclTax field to given value.

### HasLineTotalInclTax

`func (o *IRefundLine) HasLineTotalInclTax() bool`

HasLineTotalInclTax returns a boolean if a field has been set.

### GetOriginalLineTotalInclTax

`func (o *IRefundLine) GetOriginalLineTotalInclTax() float32`

GetOriginalLineTotalInclTax returns the OriginalLineTotalInclTax field if non-nil, zero value otherwise.

### GetOriginalLineTotalInclTaxOk

`func (o *IRefundLine) GetOriginalLineTotalInclTaxOk() (*float32, bool)`

GetOriginalLineTotalInclTaxOk returns a tuple with the OriginalLineTotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineTotalInclTax

`func (o *IRefundLine) SetOriginalLineTotalInclTax(v float32)`

SetOriginalLineTotalInclTax sets OriginalLineTotalInclTax field to given value.

### HasOriginalLineTotalInclTax

`func (o *IRefundLine) HasOriginalLineTotalInclTax() bool`

HasOriginalLineTotalInclTax returns a boolean if a field has been set.

### GetOriginalUnitTax

`func (o *IRefundLine) GetOriginalUnitTax() float32`

GetOriginalUnitTax returns the OriginalUnitTax field if non-nil, zero value otherwise.

### GetOriginalUnitTaxOk

`func (o *IRefundLine) GetOriginalUnitTaxOk() (*float32, bool)`

GetOriginalUnitTaxOk returns a tuple with the OriginalUnitTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitTax

`func (o *IRefundLine) SetOriginalUnitTax(v float32)`

SetOriginalUnitTax sets OriginalUnitTax field to given value.

### HasOriginalUnitTax

`func (o *IRefundLine) HasOriginalUnitTax() bool`

HasOriginalUnitTax returns a boolean if a field has been set.

### GetRefundId

`func (o *IRefundLine) GetRefundId() int32`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *IRefundLine) GetRefundIdOk() (*int32, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *IRefundLine) SetRefundId(v int32)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *IRefundLine) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetOrderLineId

`func (o *IRefundLine) GetOrderLineId() int32`

GetOrderLineId returns the OrderLineId field if non-nil, zero value otherwise.

### GetOrderLineIdOk

`func (o *IRefundLine) GetOrderLineIdOk() (*int32, bool)`

GetOrderLineIdOk returns a tuple with the OrderLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineId

`func (o *IRefundLine) SetOrderLineId(v int32)`

SetOrderLineId sets OrderLineId field to given value.

### HasOrderLineId

`func (o *IRefundLine) HasOrderLineId() bool`

HasOrderLineId returns a boolean if a field has been set.

### GetReturnLineId

`func (o *IRefundLine) GetReturnLineId() int32`

GetReturnLineId returns the ReturnLineId field if non-nil, zero value otherwise.

### GetReturnLineIdOk

`func (o *IRefundLine) GetReturnLineIdOk() (*int32, bool)`

GetReturnLineIdOk returns a tuple with the ReturnLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineId

`func (o *IRefundLine) SetReturnLineId(v int32)`

SetReturnLineId sets ReturnLineId field to given value.

### HasReturnLineId

`func (o *IRefundLine) HasReturnLineId() bool`

HasReturnLineId returns a boolean if a field has been set.

### SetReturnLineIdNil

`func (o *IRefundLine) SetReturnLineIdNil(b bool)`

 SetReturnLineIdNil sets the value for ReturnLineId to be an explicit nil

### UnsetReturnLineId
`func (o *IRefundLine) UnsetReturnLineId()`

UnsetReturnLineId ensures that no value is present for ReturnLineId, not even an explicit nil
### GetChannelOrderLineNo

`func (o *IRefundLine) GetChannelOrderLineNo() string`

GetChannelOrderLineNo returns the ChannelOrderLineNo field if non-nil, zero value otherwise.

### GetChannelOrderLineNoOk

`func (o *IRefundLine) GetChannelOrderLineNoOk() (*string, bool)`

GetChannelOrderLineNoOk returns a tuple with the ChannelOrderLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderLineNo

`func (o *IRefundLine) SetChannelOrderLineNo(v string)`

SetChannelOrderLineNo sets ChannelOrderLineNo field to given value.

### HasChannelOrderLineNo

`func (o *IRefundLine) HasChannelOrderLineNo() bool`

HasChannelOrderLineNo returns a boolean if a field has been set.

### SetChannelOrderLineNoNil

`func (o *IRefundLine) SetChannelOrderLineNoNil(b bool)`

 SetChannelOrderLineNoNil sets the value for ChannelOrderLineNo to be an explicit nil

### UnsetChannelOrderLineNo
`func (o *IRefundLine) UnsetChannelOrderLineNo()`

UnsetChannelOrderLineNo ensures that no value is present for ChannelOrderLineNo, not even an explicit nil
### GetMerchantRefundLineNo

`func (o *IRefundLine) GetMerchantRefundLineNo() string`

GetMerchantRefundLineNo returns the MerchantRefundLineNo field if non-nil, zero value otherwise.

### GetMerchantRefundLineNoOk

`func (o *IRefundLine) GetMerchantRefundLineNoOk() (*string, bool)`

GetMerchantRefundLineNoOk returns a tuple with the MerchantRefundLineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundLineNo

`func (o *IRefundLine) SetMerchantRefundLineNo(v string)`

SetMerchantRefundLineNo sets MerchantRefundLineNo field to given value.

### HasMerchantRefundLineNo

`func (o *IRefundLine) HasMerchantRefundLineNo() bool`

HasMerchantRefundLineNo returns a boolean if a field has been set.

### SetMerchantRefundLineNoNil

`func (o *IRefundLine) SetMerchantRefundLineNoNil(b bool)`

 SetMerchantRefundLineNoNil sets the value for MerchantRefundLineNo to be an explicit nil

### UnsetMerchantRefundLineNo
`func (o *IRefundLine) UnsetMerchantRefundLineNo()`

UnsetMerchantRefundLineNo ensures that no value is present for MerchantRefundLineNo, not even an explicit nil
### GetChannelProductNo

`func (o *IRefundLine) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *IRefundLine) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *IRefundLine) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *IRefundLine) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *IRefundLine) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *IRefundLine) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetCreatedAt

`func (o *IRefundLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IRefundLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IRefundLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IRefundLine) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IRefundLine) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IRefundLine) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IRefundLine) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IRefundLine) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IRefundLine) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IRefundLine) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IRefundLine) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IRefundLine) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *IRefundLine) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *IRefundLine) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetMerchantProductNo

`func (o *IRefundLine) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *IRefundLine) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *IRefundLine) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *IRefundLine) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *IRefundLine) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *IRefundLine) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetProductName

`func (o *IRefundLine) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IRefundLine) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IRefundLine) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IRefundLine) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *IRefundLine) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *IRefundLine) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetProductId

`func (o *IRefundLine) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *IRefundLine) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *IRefundLine) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *IRefundLine) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *IRefundLine) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *IRefundLine) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


