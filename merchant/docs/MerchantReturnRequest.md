# MerchantReturnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderNo** | **string** | The unique order reference used by the Merchant (sku). | 
**MerchantReturnNo** | **string** | The unique return reference used by the Merchant (sku). | 
**Lines** | [**[]MerchantReturnLineRequest**](MerchantReturnLineRequest.md) |  | 
**Id** | Pointer to **int32** | The unique return reference used by ChannelEngine. | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**CustomerComment** | Pointer to **NullableString** | Optional. Comment of customer on the (reason of) the return. | [optional] 
**MerchantComment** | Pointer to **NullableString** | Optional. Comment of merchant on the return. | [optional] 
**RefundInclVat** | Pointer to **float32** | Refund amount incl. VAT. | [optional] 
**RefundExclVat** | Pointer to **float32** | Refund amount excl. VAT. | [optional] 
**ReturnDate** | Pointer to **NullableTime** | The date at which the return was originally created in the source system (if available). | [optional] 
**ExtraData** | Pointer to **map[string]string** | Extra data on the return. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantReturnRequest

`func NewMerchantReturnRequest(merchantOrderNo string, merchantReturnNo string, lines []MerchantReturnLineRequest, ) *MerchantReturnRequest`

NewMerchantReturnRequest instantiates a new MerchantReturnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnRequestWithDefaults

`func NewMerchantReturnRequestWithDefaults() *MerchantReturnRequest`

NewMerchantReturnRequestWithDefaults instantiates a new MerchantReturnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderNo

`func (o *MerchantReturnRequest) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantReturnRequest) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantReturnRequest) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetMerchantReturnNo

`func (o *MerchantReturnRequest) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *MerchantReturnRequest) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *MerchantReturnRequest) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.


### GetLines

`func (o *MerchantReturnRequest) GetLines() []MerchantReturnLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantReturnRequest) GetLinesOk() (*[]MerchantReturnLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantReturnRequest) SetLines(v []MerchantReturnLineRequest)`

SetLines sets Lines field to given value.


### GetId

`func (o *MerchantReturnRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantReturnRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantReturnRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantReturnRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *MerchantReturnRequest) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantReturnRequest) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantReturnRequest) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantReturnRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCustomerComment

`func (o *MerchantReturnRequest) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *MerchantReturnRequest) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *MerchantReturnRequest) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *MerchantReturnRequest) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### SetCustomerCommentNil

`func (o *MerchantReturnRequest) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *MerchantReturnRequest) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetMerchantComment

`func (o *MerchantReturnRequest) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantReturnRequest) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantReturnRequest) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantReturnRequest) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantReturnRequest) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantReturnRequest) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetRefundInclVat

`func (o *MerchantReturnRequest) GetRefundInclVat() float32`

GetRefundInclVat returns the RefundInclVat field if non-nil, zero value otherwise.

### GetRefundInclVatOk

`func (o *MerchantReturnRequest) GetRefundInclVatOk() (*float32, bool)`

GetRefundInclVatOk returns a tuple with the RefundInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInclVat

`func (o *MerchantReturnRequest) SetRefundInclVat(v float32)`

SetRefundInclVat sets RefundInclVat field to given value.

### HasRefundInclVat

`func (o *MerchantReturnRequest) HasRefundInclVat() bool`

HasRefundInclVat returns a boolean if a field has been set.

### GetRefundExclVat

`func (o *MerchantReturnRequest) GetRefundExclVat() float32`

GetRefundExclVat returns the RefundExclVat field if non-nil, zero value otherwise.

### GetRefundExclVatOk

`func (o *MerchantReturnRequest) GetRefundExclVatOk() (*float32, bool)`

GetRefundExclVatOk returns a tuple with the RefundExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExclVat

`func (o *MerchantReturnRequest) SetRefundExclVat(v float32)`

SetRefundExclVat sets RefundExclVat field to given value.

### HasRefundExclVat

`func (o *MerchantReturnRequest) HasRefundExclVat() bool`

HasRefundExclVat returns a boolean if a field has been set.

### GetReturnDate

`func (o *MerchantReturnRequest) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *MerchantReturnRequest) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *MerchantReturnRequest) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *MerchantReturnRequest) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### SetReturnDateNil

`func (o *MerchantReturnRequest) SetReturnDateNil(b bool)`

 SetReturnDateNil sets the value for ReturnDate to be an explicit nil

### UnsetReturnDate
`func (o *MerchantReturnRequest) UnsetReturnDate()`

UnsetReturnDate ensures that no value is present for ReturnDate, not even an explicit nil
### GetExtraData

`func (o *MerchantReturnRequest) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantReturnRequest) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantReturnRequest) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantReturnRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantReturnRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantReturnRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


