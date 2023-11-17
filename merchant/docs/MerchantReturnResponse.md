# MerchantReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant. | [optional] 
**ChannelOrderNo** | Pointer to **NullableString** | The unique order reference used by the Channel. | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The id of the channel. | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** | The id of the Global Channel. | [optional] 
**GlobalChannelName** | Pointer to **NullableString** | The name of the Global Channel. | [optional] 
**Lines** | Pointer to [**[]MerchantReturnLineResponse**](MerchantReturnLineResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date at which the return was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date at which the return was last modified in ChannelEngine. | [optional] 
**MerchantReturnNo** | Pointer to **NullableString** | The unique return reference used by the Merchant, will be empty in case of a channel or unacknowledged return. | [optional] 
**ChannelReturnNo** | Pointer to **NullableString** | The unique return reference used by the Channel, will be empty in case of a merchant return. | [optional] 
**Status** | Pointer to [**ReturnStatus**](ReturnStatus.md) |  | [optional] 
**AcknowledgedDate** | Pointer to **NullableTime** | Date of acknowledgement | [optional] 
**Id** | Pointer to **int32** | The unique return reference used by ChannelEngine. | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**CustomerComment** | Pointer to **NullableString** | Optional. Comment of customer on the (reason of) the return. | [optional] 
**MerchantComment** | Pointer to **NullableString** | Optional. Comment of merchant on the return. | [optional] 
**RefundInclVat** | Pointer to **float32** | Refund amount incl. VAT. | [optional] 
**RefundExclVat** | Pointer to **float32** | Refund amount excl. VAT. | [optional] 
**ReturnDate** | Pointer to **NullableTime** | The date at which the return was originally created in the source system (if available). | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the return. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantReturnResponse

`func NewMerchantReturnResponse() *MerchantReturnResponse`

NewMerchantReturnResponse instantiates a new MerchantReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnResponseWithDefaults

`func NewMerchantReturnResponseWithDefaults() *MerchantReturnResponse`

NewMerchantReturnResponseWithDefaults instantiates a new MerchantReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderNo

`func (o *MerchantReturnResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantReturnResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantReturnResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *MerchantReturnResponse) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *MerchantReturnResponse) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *MerchantReturnResponse) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetChannelOrderNo

`func (o *MerchantReturnResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *MerchantReturnResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *MerchantReturnResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *MerchantReturnResponse) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *MerchantReturnResponse) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *MerchantReturnResponse) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetChannelId

`func (o *MerchantReturnResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantReturnResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantReturnResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantReturnResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantReturnResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantReturnResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetGlobalChannelId

`func (o *MerchantReturnResponse) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *MerchantReturnResponse) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *MerchantReturnResponse) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *MerchantReturnResponse) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *MerchantReturnResponse) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *MerchantReturnResponse) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetGlobalChannelName

`func (o *MerchantReturnResponse) GetGlobalChannelName() string`

GetGlobalChannelName returns the GlobalChannelName field if non-nil, zero value otherwise.

### GetGlobalChannelNameOk

`func (o *MerchantReturnResponse) GetGlobalChannelNameOk() (*string, bool)`

GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelName

`func (o *MerchantReturnResponse) SetGlobalChannelName(v string)`

SetGlobalChannelName sets GlobalChannelName field to given value.

### HasGlobalChannelName

`func (o *MerchantReturnResponse) HasGlobalChannelName() bool`

HasGlobalChannelName returns a boolean if a field has been set.

### SetGlobalChannelNameNil

`func (o *MerchantReturnResponse) SetGlobalChannelNameNil(b bool)`

 SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil

### UnsetGlobalChannelName
`func (o *MerchantReturnResponse) UnsetGlobalChannelName()`

UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil
### GetLines

`func (o *MerchantReturnResponse) GetLines() []MerchantReturnLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantReturnResponse) GetLinesOk() (*[]MerchantReturnLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantReturnResponse) SetLines(v []MerchantReturnLineResponse)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantReturnResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantReturnResponse) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantReturnResponse) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetCreatedAt

`func (o *MerchantReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantReturnResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantReturnResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantReturnResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantReturnResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantReturnResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMerchantReturnNo

`func (o *MerchantReturnResponse) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *MerchantReturnResponse) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *MerchantReturnResponse) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.

### HasMerchantReturnNo

`func (o *MerchantReturnResponse) HasMerchantReturnNo() bool`

HasMerchantReturnNo returns a boolean if a field has been set.

### SetMerchantReturnNoNil

`func (o *MerchantReturnResponse) SetMerchantReturnNoNil(b bool)`

 SetMerchantReturnNoNil sets the value for MerchantReturnNo to be an explicit nil

### UnsetMerchantReturnNo
`func (o *MerchantReturnResponse) UnsetMerchantReturnNo()`

UnsetMerchantReturnNo ensures that no value is present for MerchantReturnNo, not even an explicit nil
### GetChannelReturnNo

`func (o *MerchantReturnResponse) GetChannelReturnNo() string`

GetChannelReturnNo returns the ChannelReturnNo field if non-nil, zero value otherwise.

### GetChannelReturnNoOk

`func (o *MerchantReturnResponse) GetChannelReturnNoOk() (*string, bool)`

GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnNo

`func (o *MerchantReturnResponse) SetChannelReturnNo(v string)`

SetChannelReturnNo sets ChannelReturnNo field to given value.

### HasChannelReturnNo

`func (o *MerchantReturnResponse) HasChannelReturnNo() bool`

HasChannelReturnNo returns a boolean if a field has been set.

### SetChannelReturnNoNil

`func (o *MerchantReturnResponse) SetChannelReturnNoNil(b bool)`

 SetChannelReturnNoNil sets the value for ChannelReturnNo to be an explicit nil

### UnsetChannelReturnNo
`func (o *MerchantReturnResponse) UnsetChannelReturnNo()`

UnsetChannelReturnNo ensures that no value is present for ChannelReturnNo, not even an explicit nil
### GetStatus

`func (o *MerchantReturnResponse) GetStatus() ReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantReturnResponse) GetStatusOk() (*ReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantReturnResponse) SetStatus(v ReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MerchantReturnResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcknowledgedDate

`func (o *MerchantReturnResponse) GetAcknowledgedDate() time.Time`

GetAcknowledgedDate returns the AcknowledgedDate field if non-nil, zero value otherwise.

### GetAcknowledgedDateOk

`func (o *MerchantReturnResponse) GetAcknowledgedDateOk() (*time.Time, bool)`

GetAcknowledgedDateOk returns a tuple with the AcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedDate

`func (o *MerchantReturnResponse) SetAcknowledgedDate(v time.Time)`

SetAcknowledgedDate sets AcknowledgedDate field to given value.

### HasAcknowledgedDate

`func (o *MerchantReturnResponse) HasAcknowledgedDate() bool`

HasAcknowledgedDate returns a boolean if a field has been set.

### SetAcknowledgedDateNil

`func (o *MerchantReturnResponse) SetAcknowledgedDateNil(b bool)`

 SetAcknowledgedDateNil sets the value for AcknowledgedDate to be an explicit nil

### UnsetAcknowledgedDate
`func (o *MerchantReturnResponse) UnsetAcknowledgedDate()`

UnsetAcknowledgedDate ensures that no value is present for AcknowledgedDate, not even an explicit nil
### GetId

`func (o *MerchantReturnResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantReturnResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantReturnResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantReturnResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *MerchantReturnResponse) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantReturnResponse) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantReturnResponse) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCustomerComment

`func (o *MerchantReturnResponse) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *MerchantReturnResponse) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *MerchantReturnResponse) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *MerchantReturnResponse) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### SetCustomerCommentNil

`func (o *MerchantReturnResponse) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *MerchantReturnResponse) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetMerchantComment

`func (o *MerchantReturnResponse) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantReturnResponse) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantReturnResponse) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantReturnResponse) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantReturnResponse) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantReturnResponse) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetRefundInclVat

`func (o *MerchantReturnResponse) GetRefundInclVat() float32`

GetRefundInclVat returns the RefundInclVat field if non-nil, zero value otherwise.

### GetRefundInclVatOk

`func (o *MerchantReturnResponse) GetRefundInclVatOk() (*float32, bool)`

GetRefundInclVatOk returns a tuple with the RefundInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInclVat

`func (o *MerchantReturnResponse) SetRefundInclVat(v float32)`

SetRefundInclVat sets RefundInclVat field to given value.

### HasRefundInclVat

`func (o *MerchantReturnResponse) HasRefundInclVat() bool`

HasRefundInclVat returns a boolean if a field has been set.

### GetRefundExclVat

`func (o *MerchantReturnResponse) GetRefundExclVat() float32`

GetRefundExclVat returns the RefundExclVat field if non-nil, zero value otherwise.

### GetRefundExclVatOk

`func (o *MerchantReturnResponse) GetRefundExclVatOk() (*float32, bool)`

GetRefundExclVatOk returns a tuple with the RefundExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExclVat

`func (o *MerchantReturnResponse) SetRefundExclVat(v float32)`

SetRefundExclVat sets RefundExclVat field to given value.

### HasRefundExclVat

`func (o *MerchantReturnResponse) HasRefundExclVat() bool`

HasRefundExclVat returns a boolean if a field has been set.

### GetReturnDate

`func (o *MerchantReturnResponse) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *MerchantReturnResponse) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *MerchantReturnResponse) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *MerchantReturnResponse) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### SetReturnDateNil

`func (o *MerchantReturnResponse) SetReturnDateNil(b bool)`

 SetReturnDateNil sets the value for ReturnDate to be an explicit nil

### UnsetReturnDate
`func (o *MerchantReturnResponse) UnsetReturnDate()`

UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
### GetExtraData

`func (o *MerchantReturnResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantReturnResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantReturnResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantReturnResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantReturnResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantReturnResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


