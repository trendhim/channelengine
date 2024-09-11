# IReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**[]IReturnLine**](IReturnLine.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**ModuleReturnStatus**](ModuleReturnStatus.md) |  | [optional] 
**ChannelExportStatus** | Pointer to [**ChannelExportStatus**](ChannelExportStatus.md) |  | [optional] 
**Reason** | Pointer to [**ModuleReturnReason**](ModuleReturnReason.md) |  | [optional] 
**MerchantComment** | Pointer to **NullableString** |  | [optional] 
**MerchantReturnNo** | Pointer to **NullableString** |  | [optional] 
**ChannelReturnNo** | Pointer to **NullableString** |  | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** |  | [optional] 
**CreatedByType** | Pointer to [**CreatedByType**](CreatedByType.md) |  | [optional] 
**ReturnDate** | Pointer to **time.Time** |  | [optional] 
**OrderDate** | Pointer to **NullableTime** |  | [optional] 
**ExternalBatchNo** | Pointer to **NullableString** |  | [optional] 
**ChannelAcknowledgedDate** | Pointer to **NullableTime** |  | [optional] 
**MerchantAcknowledgedDate** | Pointer to **NullableTime** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**ChannelId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewIReturn

`func NewIReturn() *IReturn`

NewIReturn instantiates a new IReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIReturnWithDefaults

`func NewIReturnWithDefaults() *IReturn`

NewIReturnWithDefaults instantiates a new IReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *IReturn) GetLines() []IReturnLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *IReturn) GetLinesOk() (*[]IReturnLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *IReturn) SetLines(v []IReturnLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *IReturn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *IReturn) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *IReturn) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetId

`func (o *IReturn) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IReturn) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IReturn) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IReturn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *IReturn) GetStatus() ModuleReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IReturn) GetStatusOk() (*ModuleReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IReturn) SetStatus(v ModuleReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IReturn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetChannelExportStatus

`func (o *IReturn) GetChannelExportStatus() ChannelExportStatus`

GetChannelExportStatus returns the ChannelExportStatus field if non-nil, zero value otherwise.

### GetChannelExportStatusOk

`func (o *IReturn) GetChannelExportStatusOk() (*ChannelExportStatus, bool)`

GetChannelExportStatusOk returns a tuple with the ChannelExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelExportStatus

`func (o *IReturn) SetChannelExportStatus(v ChannelExportStatus)`

SetChannelExportStatus sets ChannelExportStatus field to given value.

### HasChannelExportStatus

`func (o *IReturn) HasChannelExportStatus() bool`

HasChannelExportStatus returns a boolean if a field has been set.

### GetReason

`func (o *IReturn) GetReason() ModuleReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IReturn) GetReasonOk() (*ModuleReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IReturn) SetReason(v ModuleReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IReturn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMerchantComment

`func (o *IReturn) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *IReturn) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *IReturn) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *IReturn) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *IReturn) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *IReturn) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetMerchantReturnNo

`func (o *IReturn) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *IReturn) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *IReturn) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.

### HasMerchantReturnNo

`func (o *IReturn) HasMerchantReturnNo() bool`

HasMerchantReturnNo returns a boolean if a field has been set.

### SetMerchantReturnNoNil

`func (o *IReturn) SetMerchantReturnNoNil(b bool)`

 SetMerchantReturnNoNil sets the value for MerchantReturnNo to be an explicit nil

### UnsetMerchantReturnNo
`func (o *IReturn) UnsetMerchantReturnNo()`

UnsetMerchantReturnNo ensures that no value is present for MerchantReturnNo, not even an explicit nil
### GetChannelReturnNo

`func (o *IReturn) GetChannelReturnNo() string`

GetChannelReturnNo returns the ChannelReturnNo field if non-nil, zero value otherwise.

### GetChannelReturnNoOk

`func (o *IReturn) GetChannelReturnNoOk() (*string, bool)`

GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnNo

`func (o *IReturn) SetChannelReturnNo(v string)`

SetChannelReturnNo sets ChannelReturnNo field to given value.

### HasChannelReturnNo

`func (o *IReturn) HasChannelReturnNo() bool`

HasChannelReturnNo returns a boolean if a field has been set.

### SetChannelReturnNoNil

`func (o *IReturn) SetChannelReturnNoNil(b bool)`

 SetChannelReturnNoNil sets the value for ChannelReturnNo to be an explicit nil

### UnsetChannelReturnNo
`func (o *IReturn) UnsetChannelReturnNo()`

UnsetChannelReturnNo ensures that no value is present for ChannelReturnNo, not even an explicit nil
### GetChannelOrderNo

`func (o *IReturn) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *IReturn) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *IReturn) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *IReturn) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *IReturn) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *IReturn) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetCreatedByType

`func (o *IReturn) GetCreatedByType() CreatedByType`

GetCreatedByType returns the CreatedByType field if non-nil, zero value otherwise.

### GetCreatedByTypeOk

`func (o *IReturn) GetCreatedByTypeOk() (*CreatedByType, bool)`

GetCreatedByTypeOk returns a tuple with the CreatedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByType

`func (o *IReturn) SetCreatedByType(v CreatedByType)`

SetCreatedByType sets CreatedByType field to given value.

### HasCreatedByType

`func (o *IReturn) HasCreatedByType() bool`

HasCreatedByType returns a boolean if a field has been set.

### GetReturnDate

`func (o *IReturn) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *IReturn) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *IReturn) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *IReturn) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### GetOrderDate

`func (o *IReturn) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *IReturn) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *IReturn) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *IReturn) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### SetOrderDateNil

`func (o *IReturn) SetOrderDateNil(b bool)`

 SetOrderDateNil sets the value for OrderDate to be an explicit nil

### UnsetOrderDate
`func (o *IReturn) UnsetOrderDate()`

UnsetOrderDate ensures that no value is present for OrderDate, not even an explicit nil
### GetExternalBatchNo

`func (o *IReturn) GetExternalBatchNo() string`

GetExternalBatchNo returns the ExternalBatchNo field if non-nil, zero value otherwise.

### GetExternalBatchNoOk

`func (o *IReturn) GetExternalBatchNoOk() (*string, bool)`

GetExternalBatchNoOk returns a tuple with the ExternalBatchNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBatchNo

`func (o *IReturn) SetExternalBatchNo(v string)`

SetExternalBatchNo sets ExternalBatchNo field to given value.

### HasExternalBatchNo

`func (o *IReturn) HasExternalBatchNo() bool`

HasExternalBatchNo returns a boolean if a field has been set.

### SetExternalBatchNoNil

`func (o *IReturn) SetExternalBatchNoNil(b bool)`

 SetExternalBatchNoNil sets the value for ExternalBatchNo to be an explicit nil

### UnsetExternalBatchNo
`func (o *IReturn) UnsetExternalBatchNo()`

UnsetExternalBatchNo ensures that no value is present for ExternalBatchNo, not even an explicit nil
### GetChannelAcknowledgedDate

`func (o *IReturn) GetChannelAcknowledgedDate() time.Time`

GetChannelAcknowledgedDate returns the ChannelAcknowledgedDate field if non-nil, zero value otherwise.

### GetChannelAcknowledgedDateOk

`func (o *IReturn) GetChannelAcknowledgedDateOk() (*time.Time, bool)`

GetChannelAcknowledgedDateOk returns a tuple with the ChannelAcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAcknowledgedDate

`func (o *IReturn) SetChannelAcknowledgedDate(v time.Time)`

SetChannelAcknowledgedDate sets ChannelAcknowledgedDate field to given value.

### HasChannelAcknowledgedDate

`func (o *IReturn) HasChannelAcknowledgedDate() bool`

HasChannelAcknowledgedDate returns a boolean if a field has been set.

### SetChannelAcknowledgedDateNil

`func (o *IReturn) SetChannelAcknowledgedDateNil(b bool)`

 SetChannelAcknowledgedDateNil sets the value for ChannelAcknowledgedDate to be an explicit nil

### UnsetChannelAcknowledgedDate
`func (o *IReturn) UnsetChannelAcknowledgedDate()`

UnsetChannelAcknowledgedDate ensures that no value is present for ChannelAcknowledgedDate, not even an explicit nil
### GetMerchantAcknowledgedDate

`func (o *IReturn) GetMerchantAcknowledgedDate() time.Time`

GetMerchantAcknowledgedDate returns the MerchantAcknowledgedDate field if non-nil, zero value otherwise.

### GetMerchantAcknowledgedDateOk

`func (o *IReturn) GetMerchantAcknowledgedDateOk() (*time.Time, bool)`

GetMerchantAcknowledgedDateOk returns a tuple with the MerchantAcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcknowledgedDate

`func (o *IReturn) SetMerchantAcknowledgedDate(v time.Time)`

SetMerchantAcknowledgedDate sets MerchantAcknowledgedDate field to given value.

### HasMerchantAcknowledgedDate

`func (o *IReturn) HasMerchantAcknowledgedDate() bool`

HasMerchantAcknowledgedDate returns a boolean if a field has been set.

### SetMerchantAcknowledgedDateNil

`func (o *IReturn) SetMerchantAcknowledgedDateNil(b bool)`

 SetMerchantAcknowledgedDateNil sets the value for MerchantAcknowledgedDate to be an explicit nil

### UnsetMerchantAcknowledgedDate
`func (o *IReturn) UnsetMerchantAcknowledgedDate()`

UnsetMerchantAcknowledgedDate ensures that no value is present for MerchantAcknowledgedDate, not even an explicit nil
### GetOrderId

`func (o *IReturn) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *IReturn) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *IReturn) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *IReturn) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetChannelId

`func (o *IReturn) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *IReturn) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *IReturn) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *IReturn) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IReturn) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IReturn) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IReturn) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IReturn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IReturn) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IReturn) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IReturn) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IReturn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IReturn) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IReturn) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IReturn) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IReturn) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *IReturn) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *IReturn) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


