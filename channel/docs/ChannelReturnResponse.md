# ChannelReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelReturnNo** | **string** | The unique return reference used by the Channel. | 
**ChannelOrderNo** | **string** | The unique order reference used by the Channel. | 
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant (sku). | [optional] 
**Lines** | [**[]ChannelReturnLineResponse**](ChannelReturnLineResponse.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The date at which the return was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date at which the return was last modified in ChannelEngine. | [optional] 
**Id** | Pointer to **int32** | The unique return reference used by ChannelEngine. | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**CustomerComment** | Pointer to **NullableString** | Optional. Comment of customer on the (reason of) the return. | [optional] 
**MerchantComment** | Pointer to **NullableString** | Optional. Comment of merchant on the return. | [optional] 
**RefundInclVat** | Pointer to **float32** | Refund amount incl. VAT. | [optional] 
**RefundExclVat** | Pointer to **float32** | Refund amount excl. VAT. | [optional] 
**ReturnDate** | Pointer to **NullableTime** | The date at which the return was originally created in the source system (if available). | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the return. Each item must have an unqiue key | [optional] 

## Methods

### NewChannelReturnResponse

`func NewChannelReturnResponse(channelReturnNo string, channelOrderNo string, lines []ChannelReturnLineResponse, ) *ChannelReturnResponse`

NewChannelReturnResponse instantiates a new ChannelReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelReturnResponseWithDefaults

`func NewChannelReturnResponseWithDefaults() *ChannelReturnResponse`

NewChannelReturnResponseWithDefaults instantiates a new ChannelReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelReturnNo

`func (o *ChannelReturnResponse) GetChannelReturnNo() string`

GetChannelReturnNo returns the ChannelReturnNo field if non-nil, zero value otherwise.

### GetChannelReturnNoOk

`func (o *ChannelReturnResponse) GetChannelReturnNoOk() (*string, bool)`

GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnNo

`func (o *ChannelReturnResponse) SetChannelReturnNo(v string)`

SetChannelReturnNo sets ChannelReturnNo field to given value.


### GetChannelOrderNo

`func (o *ChannelReturnResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChannelReturnResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChannelReturnResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.


### GetMerchantOrderNo

`func (o *ChannelReturnResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *ChannelReturnResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *ChannelReturnResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *ChannelReturnResponse) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *ChannelReturnResponse) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *ChannelReturnResponse) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetLines

`func (o *ChannelReturnResponse) GetLines() []ChannelReturnLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelReturnResponse) GetLinesOk() (*[]ChannelReturnLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelReturnResponse) SetLines(v []ChannelReturnLineResponse)`

SetLines sets Lines field to given value.


### GetCreatedAt

`func (o *ChannelReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelReturnResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChannelReturnResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelReturnResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelReturnResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChannelReturnResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetId

`func (o *ChannelReturnResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelReturnResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelReturnResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelReturnResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *ChannelReturnResponse) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChannelReturnResponse) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChannelReturnResponse) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChannelReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCustomerComment

`func (o *ChannelReturnResponse) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *ChannelReturnResponse) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *ChannelReturnResponse) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *ChannelReturnResponse) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### SetCustomerCommentNil

`func (o *ChannelReturnResponse) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *ChannelReturnResponse) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetMerchantComment

`func (o *ChannelReturnResponse) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *ChannelReturnResponse) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *ChannelReturnResponse) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *ChannelReturnResponse) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *ChannelReturnResponse) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *ChannelReturnResponse) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetRefundInclVat

`func (o *ChannelReturnResponse) GetRefundInclVat() float32`

GetRefundInclVat returns the RefundInclVat field if non-nil, zero value otherwise.

### GetRefundInclVatOk

`func (o *ChannelReturnResponse) GetRefundInclVatOk() (*float32, bool)`

GetRefundInclVatOk returns a tuple with the RefundInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInclVat

`func (o *ChannelReturnResponse) SetRefundInclVat(v float32)`

SetRefundInclVat sets RefundInclVat field to given value.

### HasRefundInclVat

`func (o *ChannelReturnResponse) HasRefundInclVat() bool`

HasRefundInclVat returns a boolean if a field has been set.

### GetRefundExclVat

`func (o *ChannelReturnResponse) GetRefundExclVat() float32`

GetRefundExclVat returns the RefundExclVat field if non-nil, zero value otherwise.

### GetRefundExclVatOk

`func (o *ChannelReturnResponse) GetRefundExclVatOk() (*float32, bool)`

GetRefundExclVatOk returns a tuple with the RefundExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExclVat

`func (o *ChannelReturnResponse) SetRefundExclVat(v float32)`

SetRefundExclVat sets RefundExclVat field to given value.

### HasRefundExclVat

`func (o *ChannelReturnResponse) HasRefundExclVat() bool`

HasRefundExclVat returns a boolean if a field has been set.

### GetReturnDate

`func (o *ChannelReturnResponse) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *ChannelReturnResponse) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *ChannelReturnResponse) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *ChannelReturnResponse) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### SetReturnDateNil

`func (o *ChannelReturnResponse) SetReturnDateNil(b bool)`

 SetReturnDateNil sets the value for ReturnDate to be an explicit nil

### UnsetReturnDate
`func (o *ChannelReturnResponse) UnsetReturnDate()`

UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
### GetExtraData

`func (o *ChannelReturnResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelReturnResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelReturnResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelReturnResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelReturnResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelReturnResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


