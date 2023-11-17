# ChannelReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOrderNo** | Pointer to **NullableString** | The unique order reference used by the Channel. | [optional] 
**MerchantOrderNo** | Pointer to **NullableString** | The unique order reference used by the Merchant (sku). | [optional] 
**ChannelReference** | **string** | The unique return reference used by the Channel. | 
**KeyIsMerchantOrderNo** | Pointer to **bool** | Optional. Is the MON used as key for the order (default value is false). | [optional] 
**KeyIsMerchantProductNo** | Pointer to **bool** | Optional. Is the MPN used as key for the product (default value is false). | [optional] 
**Lines** | [**[]ChannelReturnLineRequest**](ChannelReturnLineRequest.md) |  | 
**Status** | Pointer to [**ChannelReturnStatus**](ChannelReturnStatus.md) |  | [optional] 
**Id** | Pointer to **int32** | The unique return reference used by ChannelEngine. | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**CustomerComment** | Pointer to **NullableString** | Optional. Comment of customer on the (reason of) the return. | [optional] 
**MerchantComment** | Pointer to **NullableString** | Optional. Comment of merchant on the return. | [optional] 
**RefundInclVat** | Pointer to **float32** | Refund amount incl. VAT. | [optional] 
**RefundExclVat** | Pointer to **float32** | Refund amount excl. VAT. | [optional] 
**ReturnDate** | Pointer to **NullableTime** | The date at which the return was originally created in the source system (if available). | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the return. Each item must have an unqiue key | [optional] 

## Methods

### NewChannelReturnRequest

`func NewChannelReturnRequest(channelReference string, lines []ChannelReturnLineRequest, ) *ChannelReturnRequest`

NewChannelReturnRequest instantiates a new ChannelReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelReturnRequestWithDefaults

`func NewChannelReturnRequestWithDefaults() *ChannelReturnRequest`

NewChannelReturnRequestWithDefaults instantiates a new ChannelReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOrderNo

`func (o *ChannelReturnRequest) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChannelReturnRequest) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChannelReturnRequest) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *ChannelReturnRequest) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *ChannelReturnRequest) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *ChannelReturnRequest) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetMerchantOrderNo

`func (o *ChannelReturnRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *ChannelReturnRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *ChannelReturnRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.

### HasMerchantOrderNo

`func (o *ChannelReturnRequest) HasMerchantOrderNo() bool`

HasMerchantOrderNo returns a boolean if a field has been set.

### SetMerchantOrderNoNil

`func (o *ChannelReturnRequest) SetMerchantOrderNoNil(b bool)`

 SetMerchantOrderNoNil sets the value for MerchantOrderNo to be an explicit nil

### UnsetMerchantOrderNo
`func (o *ChannelReturnRequest) UnsetMerchantOrderNo()`

UnsetMerchantOrderNo ensures that no value is present for MerchantOrderNo, not even an explicit nil
### GetChannelReference

`func (o *ChannelReturnRequest) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ChannelReturnRequest) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ChannelReturnRequest) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetKeyIsMerchantOrderNo

`func (o *ChannelReturnRequest) GetKeyIsMerchantOrderNo() bool`

GetKeyIsMerchantOrderNo returns the KeyIsMerchantOrderNo field if non-nil, zero value otherwise.

### GetKeyIsMerchantOrderNoOk

`func (o *ChannelReturnRequest) GetKeyIsMerchantOrderNoOk() (*bool, bool)`

GetKeyIsMerchantOrderNoOk returns a tuple with the KeyIsMerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIsMerchantOrderNo

`func (o *ChannelReturnRequest) SetKeyIsMerchantOrderNo(v bool)`

SetKeyIsMerchantOrderNo sets KeyIsMerchantOrderNo field to given value.

### HasKeyIsMerchantOrderNo

`func (o *ChannelReturnRequest) HasKeyIsMerchantOrderNo() bool`

HasKeyIsMerchantOrderNo returns a boolean if a field has been set.

### GetKeyIsMerchantProductNo

`func (o *ChannelReturnRequest) GetKeyIsMerchantProductNo() bool`

GetKeyIsMerchantProductNo returns the KeyIsMerchantProductNo field if non-nil, zero value otherwise.

### GetKeyIsMerchantProductNoOk

`func (o *ChannelReturnRequest) GetKeyIsMerchantProductNoOk() (*bool, bool)`

GetKeyIsMerchantProductNoOk returns a tuple with the KeyIsMerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIsMerchantProductNo

`func (o *ChannelReturnRequest) SetKeyIsMerchantProductNo(v bool)`

SetKeyIsMerchantProductNo sets KeyIsMerchantProductNo field to given value.

### HasKeyIsMerchantProductNo

`func (o *ChannelReturnRequest) HasKeyIsMerchantProductNo() bool`

HasKeyIsMerchantProductNo returns a boolean if a field has been set.

### GetLines

`func (o *ChannelReturnRequest) GetLines() []ChannelReturnLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelReturnRequest) GetLinesOk() (*[]ChannelReturnLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelReturnRequest) SetLines(v []ChannelReturnLineRequest)`

SetLines sets Lines field to given value.


### GetStatus

`func (o *ChannelReturnRequest) GetStatus() ChannelReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChannelReturnRequest) GetStatusOk() (*ChannelReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChannelReturnRequest) SetStatus(v ChannelReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChannelReturnRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *ChannelReturnRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelReturnRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelReturnRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelReturnRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *ChannelReturnRequest) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChannelReturnRequest) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChannelReturnRequest) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChannelReturnRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCustomerComment

`func (o *ChannelReturnRequest) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *ChannelReturnRequest) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *ChannelReturnRequest) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *ChannelReturnRequest) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### SetCustomerCommentNil

`func (o *ChannelReturnRequest) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *ChannelReturnRequest) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetMerchantComment

`func (o *ChannelReturnRequest) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *ChannelReturnRequest) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *ChannelReturnRequest) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *ChannelReturnRequest) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *ChannelReturnRequest) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *ChannelReturnRequest) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetRefundInclVat

`func (o *ChannelReturnRequest) GetRefundInclVat() float32`

GetRefundInclVat returns the RefundInclVat field if non-nil, zero value otherwise.

### GetRefundInclVatOk

`func (o *ChannelReturnRequest) GetRefundInclVatOk() (*float32, bool)`

GetRefundInclVatOk returns a tuple with the RefundInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInclVat

`func (o *ChannelReturnRequest) SetRefundInclVat(v float32)`

SetRefundInclVat sets RefundInclVat field to given value.

### HasRefundInclVat

`func (o *ChannelReturnRequest) HasRefundInclVat() bool`

HasRefundInclVat returns a boolean if a field has been set.

### GetRefundExclVat

`func (o *ChannelReturnRequest) GetRefundExclVat() float32`

GetRefundExclVat returns the RefundExclVat field if non-nil, zero value otherwise.

### GetRefundExclVatOk

`func (o *ChannelReturnRequest) GetRefundExclVatOk() (*float32, bool)`

GetRefundExclVatOk returns a tuple with the RefundExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExclVat

`func (o *ChannelReturnRequest) SetRefundExclVat(v float32)`

SetRefundExclVat sets RefundExclVat field to given value.

### HasRefundExclVat

`func (o *ChannelReturnRequest) HasRefundExclVat() bool`

HasRefundExclVat returns a boolean if a field has been set.

### GetReturnDate

`func (o *ChannelReturnRequest) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *ChannelReturnRequest) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *ChannelReturnRequest) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *ChannelReturnRequest) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### SetReturnDateNil

`func (o *ChannelReturnRequest) SetReturnDateNil(b bool)`

 SetReturnDateNil sets the value for ReturnDate to be an explicit nil

### UnsetReturnDate
`func (o *ChannelReturnRequest) UnsetReturnDate()`

UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
### GetExtraData

`func (o *ChannelReturnRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelReturnRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelReturnRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelReturnRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelReturnRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelReturnRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


