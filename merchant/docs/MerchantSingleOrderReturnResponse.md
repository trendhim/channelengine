# MerchantSingleOrderReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant. | [optional] 
**Lines** | Pointer to [**[]MerchantSingleOrderReturnLineResponse**](MerchantSingleOrderReturnLineResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date at which the return was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date at which the return was last modified in ChannelEngine. | [optional] 
**MerchantReturnNo** | Pointer to **NullableString** | The unique return reference used by the Merchant, will be empty in case of a channel or unacknowledged return. | [optional] 
**ChannelReturnNo** | Pointer to **NullableString** | The unique return reference used by the Channel, will be empty in case of a merchant return. | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The id of the channel. | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** | The id of the Global Channel. | [optional] 
**GlobalChannelName** | Pointer to **NullableString** | The name of the Global Channel. | [optional] 
**Status** | Pointer to [**ReturnStatus**](ReturnStatus.md) |  | [optional] 
**Id** | Pointer to **int32** | The unique return reference used by ChannelEngine. | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**CustomerComment** | Pointer to **NullableString** | Optional. Comment of customer on the (reason of) the return. | [optional] 
**MerchantComment** | Pointer to **NullableString** | Optional. Comment of merchant on the return. | [optional] 
**RefundInclVat** | Pointer to **float32** | Refund amount incl. VAT. | [optional] 
**RefundExclVat** | Pointer to **float32** | Refund amount excl. VAT. | [optional] 
**ReturnDate** | Pointer to **NullableTime** | The date at which the return was originally created in the source system (if available). | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the return. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantSingleOrderReturnResponse

`func NewMerchantSingleOrderReturnResponse() *MerchantSingleOrderReturnResponse`

NewMerchantSingleOrderReturnResponse instantiates a new MerchantSingleOrderReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSingleOrderReturnResponseWithDefaults

`func NewMerchantSingleOrderReturnResponseWithDefaults() *MerchantSingleOrderReturnResponse`

NewMerchantSingleOrderReturnResponseWithDefaults instantiates a new MerchantSingleOrderReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderNo

`func (o *MerchantSingleOrderReturnResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantSingleOrderReturnResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantSingleOrderReturnResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantSingleOrderReturnResponse) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantSingleOrderReturnResponse) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantSingleOrderReturnResponse) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetLines

`func (o *MerchantSingleOrderReturnResponse) GetLines() []MerchantSingleOrderReturnLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantSingleOrderReturnResponse) GetLinesOk() (*[]MerchantSingleOrderReturnLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantSingleOrderReturnResponse) SetLines(v []MerchantSingleOrderReturnLineResponse)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantSingleOrderReturnResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantSingleOrderReturnResponse) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantSingleOrderReturnResponse) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetCreatedAt

`func (o *MerchantSingleOrderReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantSingleOrderReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantSingleOrderReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantSingleOrderReturnResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantSingleOrderReturnResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantSingleOrderReturnResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantSingleOrderReturnResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantSingleOrderReturnResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMerchantReturnNo

`func (o *MerchantSingleOrderReturnResponse) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *MerchantSingleOrderReturnResponse) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *MerchantSingleOrderReturnResponse) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.

### HasMerchantReturnNo

`func (o *MerchantSingleOrderReturnResponse) HasMerchantReturnNo() bool`

HasMerchantReturnNo returns a boolean if a field has been set.

### SetMerchantReturnNoNil

`func (o *MerchantSingleOrderReturnResponse) SetMerchantReturnNoNil(b bool)`

 SetMerchantReturnNoNil sets the value for MerchantReturnNo to be an explicit nil

### UnsetMerchantReturnNo
`func (o *MerchantSingleOrderReturnResponse) UnsetMerchantReturnNo()`

UnsetMerchantReturnNo ensures that no value is present for MerchantReturnNo, not even an explicit nil
### GetChannelReturnNo

`func (o *MerchantSingleOrderReturnResponse) GetChannelReturnNo() string`

GetChannelReturnNo returns the ChannelReturnNo field if non-nil, zero value otherwise.

### GetChannelReturnNoOk

`func (o *MerchantSingleOrderReturnResponse) GetChannelReturnNoOk() (*string, bool)`

GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnNo

`func (o *MerchantSingleOrderReturnResponse) SetChannelReturnNo(v string)`

SetChannelReturnNo sets ChannelReturnNo field to given value.

### HasChannelReturnNo

`func (o *MerchantSingleOrderReturnResponse) HasChannelReturnNo() bool`

HasChannelReturnNo returns a boolean if a field has been set.

### SetChannelReturnNoNil

`func (o *MerchantSingleOrderReturnResponse) SetChannelReturnNoNil(b bool)`

 SetChannelReturnNoNil sets the value for ChannelReturnNo to be an explicit nil

### UnsetChannelReturnNo
`func (o *MerchantSingleOrderReturnResponse) UnsetChannelReturnNo()`

UnsetChannelReturnNo ensures that no value is present for ChannelReturnNo, not even an explicit nil
### GetChannelId

`func (o *MerchantSingleOrderReturnResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantSingleOrderReturnResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantSingleOrderReturnResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantSingleOrderReturnResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantSingleOrderReturnResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantSingleOrderReturnResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetGlobalChannelId

`func (o *MerchantSingleOrderReturnResponse) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *MerchantSingleOrderReturnResponse) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *MerchantSingleOrderReturnResponse) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *MerchantSingleOrderReturnResponse) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *MerchantSingleOrderReturnResponse) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *MerchantSingleOrderReturnResponse) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetGlobalChannelName

`func (o *MerchantSingleOrderReturnResponse) GetGlobalChannelName() string`

GetGlobalChannelName returns the GlobalChannelName field if non-nil, zero value otherwise.

### GetGlobalChannelNameOk

`func (o *MerchantSingleOrderReturnResponse) GetGlobalChannelNameOk() (*string, bool)`

GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelName

`func (o *MerchantSingleOrderReturnResponse) SetGlobalChannelName(v string)`

SetGlobalChannelName sets GlobalChannelName field to given value.

### HasGlobalChannelName

`func (o *MerchantSingleOrderReturnResponse) HasGlobalChannelName() bool`

HasGlobalChannelName returns a boolean if a field has been set.

### SetGlobalChannelNameNil

`func (o *MerchantSingleOrderReturnResponse) SetGlobalChannelNameNil(b bool)`

 SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil

### UnsetGlobalChannelName
`func (o *MerchantSingleOrderReturnResponse) UnsetGlobalChannelName()`

UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil
### GetStatus

`func (o *MerchantSingleOrderReturnResponse) GetStatus() ReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantSingleOrderReturnResponse) GetStatusOk() (*ReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantSingleOrderReturnResponse) SetStatus(v ReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MerchantSingleOrderReturnResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *MerchantSingleOrderReturnResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantSingleOrderReturnResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantSingleOrderReturnResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantSingleOrderReturnResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *MerchantSingleOrderReturnResponse) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantSingleOrderReturnResponse) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantSingleOrderReturnResponse) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantSingleOrderReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCustomerComment

`func (o *MerchantSingleOrderReturnResponse) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *MerchantSingleOrderReturnResponse) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *MerchantSingleOrderReturnResponse) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *MerchantSingleOrderReturnResponse) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### SetCustomerCommentNil

`func (o *MerchantSingleOrderReturnResponse) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *MerchantSingleOrderReturnResponse) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetMerchantComment

`func (o *MerchantSingleOrderReturnResponse) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantSingleOrderReturnResponse) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantSingleOrderReturnResponse) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantSingleOrderReturnResponse) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantSingleOrderReturnResponse) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantSingleOrderReturnResponse) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetRefundInclVat

`func (o *MerchantSingleOrderReturnResponse) GetRefundInclVat() float32`

GetRefundInclVat returns the RefundInclVat field if non-nil, zero value otherwise.

### GetRefundInclVatOk

`func (o *MerchantSingleOrderReturnResponse) GetRefundInclVatOk() (*float32, bool)`

GetRefundInclVatOk returns a tuple with the RefundInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInclVat

`func (o *MerchantSingleOrderReturnResponse) SetRefundInclVat(v float32)`

SetRefundInclVat sets RefundInclVat field to given value.

### HasRefundInclVat

`func (o *MerchantSingleOrderReturnResponse) HasRefundInclVat() bool`

HasRefundInclVat returns a boolean if a field has been set.

### GetRefundExclVat

`func (o *MerchantSingleOrderReturnResponse) GetRefundExclVat() float32`

GetRefundExclVat returns the RefundExclVat field if non-nil, zero value otherwise.

### GetRefundExclVatOk

`func (o *MerchantSingleOrderReturnResponse) GetRefundExclVatOk() (*float32, bool)`

GetRefundExclVatOk returns a tuple with the RefundExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExclVat

`func (o *MerchantSingleOrderReturnResponse) SetRefundExclVat(v float32)`

SetRefundExclVat sets RefundExclVat field to given value.

### HasRefundExclVat

`func (o *MerchantSingleOrderReturnResponse) HasRefundExclVat() bool`

HasRefundExclVat returns a boolean if a field has been set.

### GetReturnDate

`func (o *MerchantSingleOrderReturnResponse) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *MerchantSingleOrderReturnResponse) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *MerchantSingleOrderReturnResponse) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *MerchantSingleOrderReturnResponse) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### SetReturnDateNil

`func (o *MerchantSingleOrderReturnResponse) SetReturnDateNil(b bool)`

 SetReturnDateNil sets the value for ReturnDate to be an explicit nil

### UnsetReturnDate
`func (o *MerchantSingleOrderReturnResponse) UnsetReturnDate()`

UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
### GetExtraData

`func (o *MerchantSingleOrderReturnResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantSingleOrderReturnResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantSingleOrderReturnResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantSingleOrderReturnResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantSingleOrderReturnResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantSingleOrderReturnResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


