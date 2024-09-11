# IRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to [**RefundReason**](RefundReason.md) |  | [optional] 
**ChannelExportStatus** | Pointer to [**ChannelExportStatus**](ChannelExportStatus.md) |  | [optional] 
**SubTotalInclTax** | Pointer to **float32** |  | [optional] 
**OriginalSubTotalInclTax** | Pointer to **float32** |  | [optional] 
**AdditionalAmountInclTax** | Pointer to **float32** |  | [optional] 
**OriginalAdditionalAmountInclTax** | Pointer to **float32** |  | [optional] 
**ShippingCostInclTax** | Pointer to **float32** |  | [optional] 
**OriginalShippingCostInclTax** | Pointer to **float32** |  | [optional] 
**TotalInclTax** | Pointer to **float32** |  | [optional] 
**OriginalTotalInclTax** | Pointer to **float32** |  | [optional] 
**MerchantComment** | Pointer to **NullableString** |  | [optional] 
**MerchantRefundNo** | Pointer to **NullableString** |  | [optional] 
**ChannelRefundNo** | Pointer to **NullableString** |  | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** |  | [optional] 
**CreatedByType** | Pointer to [**CreatedByType**](CreatedByType.md) |  | [optional] 
**RefundDate** | Pointer to **time.Time** |  | [optional] 
**ExternalBatchNo** | Pointer to **NullableString** |  | [optional] 
**ChannelAcknowledgedDate** | Pointer to **NullableTime** |  | [optional] 
**MerchantAcknowledgedDate** | Pointer to **NullableTime** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**ChannelId** | Pointer to **int32** |  | [optional] 
**ReturnId** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to [**IRefundCurrency**](IRefundCurrency.md) |  | [optional] 
**Lines** | Pointer to [**[]IRefundLine**](IRefundLine.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewIRefund

`func NewIRefund() *IRefund`

NewIRefund instantiates a new IRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIRefundWithDefaults

`func NewIRefundWithDefaults() *IRefund`

NewIRefundWithDefaults instantiates a new IRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IRefund) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IRefund) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IRefund) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IRefund) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *IRefund) GetReason() RefundReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IRefund) GetReasonOk() (*RefundReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IRefund) SetReason(v RefundReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IRefund) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetChannelExportStatus

`func (o *IRefund) GetChannelExportStatus() ChannelExportStatus`

GetChannelExportStatus returns the ChannelExportStatus field if non-nil, zero value otherwise.

### GetChannelExportStatusOk

`func (o *IRefund) GetChannelExportStatusOk() (*ChannelExportStatus, bool)`

GetChannelExportStatusOk returns a tuple with the ChannelExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelExportStatus

`func (o *IRefund) SetChannelExportStatus(v ChannelExportStatus)`

SetChannelExportStatus sets ChannelExportStatus field to given value.

### HasChannelExportStatus

`func (o *IRefund) HasChannelExportStatus() bool`

HasChannelExportStatus returns a boolean if a field has been set.

### GetSubTotalInclTax

`func (o *IRefund) GetSubTotalInclTax() float32`

GetSubTotalInclTax returns the SubTotalInclTax field if non-nil, zero value otherwise.

### GetSubTotalInclTaxOk

`func (o *IRefund) GetSubTotalInclTaxOk() (*float32, bool)`

GetSubTotalInclTaxOk returns a tuple with the SubTotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalInclTax

`func (o *IRefund) SetSubTotalInclTax(v float32)`

SetSubTotalInclTax sets SubTotalInclTax field to given value.

### HasSubTotalInclTax

`func (o *IRefund) HasSubTotalInclTax() bool`

HasSubTotalInclTax returns a boolean if a field has been set.

### GetOriginalSubTotalInclTax

`func (o *IRefund) GetOriginalSubTotalInclTax() float32`

GetOriginalSubTotalInclTax returns the OriginalSubTotalInclTax field if non-nil, zero value otherwise.

### GetOriginalSubTotalInclTaxOk

`func (o *IRefund) GetOriginalSubTotalInclTaxOk() (*float32, bool)`

GetOriginalSubTotalInclTaxOk returns a tuple with the OriginalSubTotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubTotalInclTax

`func (o *IRefund) SetOriginalSubTotalInclTax(v float32)`

SetOriginalSubTotalInclTax sets OriginalSubTotalInclTax field to given value.

### HasOriginalSubTotalInclTax

`func (o *IRefund) HasOriginalSubTotalInclTax() bool`

HasOriginalSubTotalInclTax returns a boolean if a field has been set.

### GetAdditionalAmountInclTax

`func (o *IRefund) GetAdditionalAmountInclTax() float32`

GetAdditionalAmountInclTax returns the AdditionalAmountInclTax field if non-nil, zero value otherwise.

### GetAdditionalAmountInclTaxOk

`func (o *IRefund) GetAdditionalAmountInclTaxOk() (*float32, bool)`

GetAdditionalAmountInclTaxOk returns a tuple with the AdditionalAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmountInclTax

`func (o *IRefund) SetAdditionalAmountInclTax(v float32)`

SetAdditionalAmountInclTax sets AdditionalAmountInclTax field to given value.

### HasAdditionalAmountInclTax

`func (o *IRefund) HasAdditionalAmountInclTax() bool`

HasAdditionalAmountInclTax returns a boolean if a field has been set.

### GetOriginalAdditionalAmountInclTax

`func (o *IRefund) GetOriginalAdditionalAmountInclTax() float32`

GetOriginalAdditionalAmountInclTax returns the OriginalAdditionalAmountInclTax field if non-nil, zero value otherwise.

### GetOriginalAdditionalAmountInclTaxOk

`func (o *IRefund) GetOriginalAdditionalAmountInclTaxOk() (*float32, bool)`

GetOriginalAdditionalAmountInclTaxOk returns a tuple with the OriginalAdditionalAmountInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAdditionalAmountInclTax

`func (o *IRefund) SetOriginalAdditionalAmountInclTax(v float32)`

SetOriginalAdditionalAmountInclTax sets OriginalAdditionalAmountInclTax field to given value.

### HasOriginalAdditionalAmountInclTax

`func (o *IRefund) HasOriginalAdditionalAmountInclTax() bool`

HasOriginalAdditionalAmountInclTax returns a boolean if a field has been set.

### GetShippingCostInclTax

`func (o *IRefund) GetShippingCostInclTax() float32`

GetShippingCostInclTax returns the ShippingCostInclTax field if non-nil, zero value otherwise.

### GetShippingCostInclTaxOk

`func (o *IRefund) GetShippingCostInclTaxOk() (*float32, bool)`

GetShippingCostInclTaxOk returns a tuple with the ShippingCostInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostInclTax

`func (o *IRefund) SetShippingCostInclTax(v float32)`

SetShippingCostInclTax sets ShippingCostInclTax field to given value.

### HasShippingCostInclTax

`func (o *IRefund) HasShippingCostInclTax() bool`

HasShippingCostInclTax returns a boolean if a field has been set.

### GetOriginalShippingCostInclTax

`func (o *IRefund) GetOriginalShippingCostInclTax() float32`

GetOriginalShippingCostInclTax returns the OriginalShippingCostInclTax field if non-nil, zero value otherwise.

### GetOriginalShippingCostInclTaxOk

`func (o *IRefund) GetOriginalShippingCostInclTaxOk() (*float32, bool)`

GetOriginalShippingCostInclTaxOk returns a tuple with the OriginalShippingCostInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShippingCostInclTax

`func (o *IRefund) SetOriginalShippingCostInclTax(v float32)`

SetOriginalShippingCostInclTax sets OriginalShippingCostInclTax field to given value.

### HasOriginalShippingCostInclTax

`func (o *IRefund) HasOriginalShippingCostInclTax() bool`

HasOriginalShippingCostInclTax returns a boolean if a field has been set.

### GetTotalInclTax

`func (o *IRefund) GetTotalInclTax() float32`

GetTotalInclTax returns the TotalInclTax field if non-nil, zero value otherwise.

### GetTotalInclTaxOk

`func (o *IRefund) GetTotalInclTaxOk() (*float32, bool)`

GetTotalInclTaxOk returns a tuple with the TotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInclTax

`func (o *IRefund) SetTotalInclTax(v float32)`

SetTotalInclTax sets TotalInclTax field to given value.

### HasTotalInclTax

`func (o *IRefund) HasTotalInclTax() bool`

HasTotalInclTax returns a boolean if a field has been set.

### GetOriginalTotalInclTax

`func (o *IRefund) GetOriginalTotalInclTax() float32`

GetOriginalTotalInclTax returns the OriginalTotalInclTax field if non-nil, zero value otherwise.

### GetOriginalTotalInclTaxOk

`func (o *IRefund) GetOriginalTotalInclTaxOk() (*float32, bool)`

GetOriginalTotalInclTaxOk returns a tuple with the OriginalTotalInclTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotalInclTax

`func (o *IRefund) SetOriginalTotalInclTax(v float32)`

SetOriginalTotalInclTax sets OriginalTotalInclTax field to given value.

### HasOriginalTotalInclTax

`func (o *IRefund) HasOriginalTotalInclTax() bool`

HasOriginalTotalInclTax returns a boolean if a field has been set.

### GetMerchantComment

`func (o *IRefund) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *IRefund) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *IRefund) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *IRefund) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *IRefund) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *IRefund) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetMerchantRefundNo

`func (o *IRefund) GetMerchantRefundNo() string`

GetMerchantRefundNo returns the MerchantRefundNo field if non-nil, zero value otherwise.

### GetMerchantRefundNoOk

`func (o *IRefund) GetMerchantRefundNoOk() (*string, bool)`

GetMerchantRefundNoOk returns a tuple with the MerchantRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundNo

`func (o *IRefund) SetMerchantRefundNo(v string)`

SetMerchantRefundNo sets MerchantRefundNo field to given value.

### HasMerchantRefundNo

`func (o *IRefund) HasMerchantRefundNo() bool`

HasMerchantRefundNo returns a boolean if a field has been set.

### SetMerchantRefundNoNil

`func (o *IRefund) SetMerchantRefundNoNil(b bool)`

 SetMerchantRefundNoNil sets the value for MerchantRefundNo to be an explicit nil

### UnsetMerchantRefundNo
`func (o *IRefund) UnsetMerchantRefundNo()`

UnsetMerchantRefundNo ensures that no value is present for MerchantRefundNo, not even an explicit nil
### GetChannelRefundNo

`func (o *IRefund) GetChannelRefundNo() string`

GetChannelRefundNo returns the ChannelRefundNo field if non-nil, zero value otherwise.

### GetChannelRefundNoOk

`func (o *IRefund) GetChannelRefundNoOk() (*string, bool)`

GetChannelRefundNoOk returns a tuple with the ChannelRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelRefundNo

`func (o *IRefund) SetChannelRefundNo(v string)`

SetChannelRefundNo sets ChannelRefundNo field to given value.

### HasChannelRefundNo

`func (o *IRefund) HasChannelRefundNo() bool`

HasChannelRefundNo returns a boolean if a field has been set.

### SetChannelRefundNoNil

`func (o *IRefund) SetChannelRefundNoNil(b bool)`

 SetChannelRefundNoNil sets the value for ChannelRefundNo to be an explicit nil

### UnsetChannelRefundNo
`func (o *IRefund) UnsetChannelRefundNo()`

UnsetChannelRefundNo ensures that no value is present for ChannelRefundNo, not even an explicit nil
### GetChannelOrderNo

`func (o *IRefund) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *IRefund) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *IRefund) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *IRefund) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *IRefund) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *IRefund) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetCreatedByType

`func (o *IRefund) GetCreatedByType() CreatedByType`

GetCreatedByType returns the CreatedByType field if non-nil, zero value otherwise.

### GetCreatedByTypeOk

`func (o *IRefund) GetCreatedByTypeOk() (*CreatedByType, bool)`

GetCreatedByTypeOk returns a tuple with the CreatedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByType

`func (o *IRefund) SetCreatedByType(v CreatedByType)`

SetCreatedByType sets CreatedByType field to given value.

### HasCreatedByType

`func (o *IRefund) HasCreatedByType() bool`

HasCreatedByType returns a boolean if a field has been set.

### GetRefundDate

`func (o *IRefund) GetRefundDate() time.Time`

GetRefundDate returns the RefundDate field if non-nil, zero value otherwise.

### GetRefundDateOk

`func (o *IRefund) GetRefundDateOk() (*time.Time, bool)`

GetRefundDateOk returns a tuple with the RefundDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDate

`func (o *IRefund) SetRefundDate(v time.Time)`

SetRefundDate sets RefundDate field to given value.

### HasRefundDate

`func (o *IRefund) HasRefundDate() bool`

HasRefundDate returns a boolean if a field has been set.

### GetExternalBatchNo

`func (o *IRefund) GetExternalBatchNo() string`

GetExternalBatchNo returns the ExternalBatchNo field if non-nil, zero value otherwise.

### GetExternalBatchNoOk

`func (o *IRefund) GetExternalBatchNoOk() (*string, bool)`

GetExternalBatchNoOk returns a tuple with the ExternalBatchNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBatchNo

`func (o *IRefund) SetExternalBatchNo(v string)`

SetExternalBatchNo sets ExternalBatchNo field to given value.

### HasExternalBatchNo

`func (o *IRefund) HasExternalBatchNo() bool`

HasExternalBatchNo returns a boolean if a field has been set.

### SetExternalBatchNoNil

`func (o *IRefund) SetExternalBatchNoNil(b bool)`

 SetExternalBatchNoNil sets the value for ExternalBatchNo to be an explicit nil

### UnsetExternalBatchNo
`func (o *IRefund) UnsetExternalBatchNo()`

UnsetExternalBatchNo ensures that no value is present for ExternalBatchNo, not even an explicit nil
### GetChannelAcknowledgedDate

`func (o *IRefund) GetChannelAcknowledgedDate() time.Time`

GetChannelAcknowledgedDate returns the ChannelAcknowledgedDate field if non-nil, zero value otherwise.

### GetChannelAcknowledgedDateOk

`func (o *IRefund) GetChannelAcknowledgedDateOk() (*time.Time, bool)`

GetChannelAcknowledgedDateOk returns a tuple with the ChannelAcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAcknowledgedDate

`func (o *IRefund) SetChannelAcknowledgedDate(v time.Time)`

SetChannelAcknowledgedDate sets ChannelAcknowledgedDate field to given value.

### HasChannelAcknowledgedDate

`func (o *IRefund) HasChannelAcknowledgedDate() bool`

HasChannelAcknowledgedDate returns a boolean if a field has been set.

### SetChannelAcknowledgedDateNil

`func (o *IRefund) SetChannelAcknowledgedDateNil(b bool)`

 SetChannelAcknowledgedDateNil sets the value for ChannelAcknowledgedDate to be an explicit nil

### UnsetChannelAcknowledgedDate
`func (o *IRefund) UnsetChannelAcknowledgedDate()`

UnsetChannelAcknowledgedDate ensures that no value is present for ChannelAcknowledgedDate, not even an explicit nil
### GetMerchantAcknowledgedDate

`func (o *IRefund) GetMerchantAcknowledgedDate() time.Time`

GetMerchantAcknowledgedDate returns the MerchantAcknowledgedDate field if non-nil, zero value otherwise.

### GetMerchantAcknowledgedDateOk

`func (o *IRefund) GetMerchantAcknowledgedDateOk() (*time.Time, bool)`

GetMerchantAcknowledgedDateOk returns a tuple with the MerchantAcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcknowledgedDate

`func (o *IRefund) SetMerchantAcknowledgedDate(v time.Time)`

SetMerchantAcknowledgedDate sets MerchantAcknowledgedDate field to given value.

### HasMerchantAcknowledgedDate

`func (o *IRefund) HasMerchantAcknowledgedDate() bool`

HasMerchantAcknowledgedDate returns a boolean if a field has been set.

### SetMerchantAcknowledgedDateNil

`func (o *IRefund) SetMerchantAcknowledgedDateNil(b bool)`

 SetMerchantAcknowledgedDateNil sets the value for MerchantAcknowledgedDate to be an explicit nil

### UnsetMerchantAcknowledgedDate
`func (o *IRefund) UnsetMerchantAcknowledgedDate()`

UnsetMerchantAcknowledgedDate ensures that no value is present for MerchantAcknowledgedDate, not even an explicit nil
### GetOrderId

`func (o *IRefund) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *IRefund) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *IRefund) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *IRefund) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetChannelId

`func (o *IRefund) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *IRefund) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *IRefund) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *IRefund) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetReturnId

`func (o *IRefund) GetReturnId() int32`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *IRefund) GetReturnIdOk() (*int32, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *IRefund) SetReturnId(v int32)`

SetReturnId sets ReturnId field to given value.

### HasReturnId

`func (o *IRefund) HasReturnId() bool`

HasReturnId returns a boolean if a field has been set.

### SetReturnIdNil

`func (o *IRefund) SetReturnIdNil(b bool)`

 SetReturnIdNil sets the value for ReturnId to be an explicit nil

### UnsetReturnId
`func (o *IRefund) UnsetReturnId()`

UnsetReturnId ensures that no value is present for ReturnId, not even an explicit nil
### GetCurrency

`func (o *IRefund) GetCurrency() IRefundCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IRefund) GetCurrencyOk() (*IRefundCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IRefund) SetCurrency(v IRefundCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IRefund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLines

`func (o *IRefund) GetLines() []IRefundLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *IRefund) GetLinesOk() (*[]IRefundLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *IRefund) SetLines(v []IRefundLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *IRefund) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *IRefund) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *IRefund) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetCreatedAt

`func (o *IRefund) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IRefund) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IRefund) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IRefund) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IRefund) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IRefund) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IRefund) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IRefund) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IRefund) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IRefund) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IRefund) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IRefund) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *IRefund) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *IRefund) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


