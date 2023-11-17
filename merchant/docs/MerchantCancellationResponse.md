# MerchantCancellationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique cancellation identifier used by ChannelEngine. | [optional] 
**MerchantCancellationNo** | **string** | The unique cancellation reference used by the Merchant (sku). | 
**MerchantOrderNo** | **string** | The unique order reference used by the Merchant. | 
**ChannelOrderNo** | Pointer to **NullableString** | The unique order reference used by the Channel. | [optional] 
**Lines** | [**[]MerchantCancellationLineResponse**](MerchantCancellationLineResponse.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The date at which the cancellation was created in ChannelEngine. | [optional] 
**Reason** | Pointer to **NullableString** | Reason for cancellation (text). | [optional] 
**ReasonCode** | Pointer to [**MancoReason**](MancoReason.md) |  | [optional] 

## Methods

### NewMerchantCancellationResponse

`func NewMerchantCancellationResponse(merchantCancellationNo string, merchantOrderNo string, lines []MerchantCancellationLineResponse, ) *MerchantCancellationResponse`

NewMerchantCancellationResponse instantiates a new MerchantCancellationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCancellationResponseWithDefaults

`func NewMerchantCancellationResponseWithDefaults() *MerchantCancellationResponse`

NewMerchantCancellationResponseWithDefaults instantiates a new MerchantCancellationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantCancellationResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantCancellationResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantCancellationResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantCancellationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantCancellationNo

`func (o *MerchantCancellationResponse) GetMerchantCancellationNo() string`

GetMerchantCancellationNo returns the MerchantCancellationNo field if non-nil, zero value otherwise.

### GetMerchantCancellationNoOk

`func (o *MerchantCancellationResponse) GetMerchantCancellationNoOk() (*string, bool)`

GetMerchantCancellationNoOk returns a tuple with the MerchantCancellationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCancellationNo

`func (o *MerchantCancellationResponse) SetMerchantCancellationNo(v string)`

SetMerchantCancellationNo sets MerchantCancellationNo field to given value.


### GetMerchantOrderNo

`func (o *MerchantCancellationResponse) GetMerchantOrderNo() string`

GetMerchantOrderNo returns the MerchantOrderNo field if non-nil, zero value otherwise.

### GetMerchantOrderNoOk

`func (o *MerchantCancellationResponse) GetMerchantOrderNoOk() (*string, bool)`

GetMerchantOrderNoOk returns a tuple with the MerchantOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderNo

`func (o *MerchantCancellationResponse) SetMerchantOrderNo(v string)`

SetMerchantOrderNo sets MerchantOrderNo field to given value.


### GetChannelOrderNo

`func (o *MerchantCancellationResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *MerchantCancellationResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *MerchantCancellationResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.

### HasChannelOrderNo

`func (o *MerchantCancellationResponse) HasChannelOrderNo() bool`

HasChannelOrderNo returns a boolean if a field has been set.

### SetChannelOrderNoNil

`func (o *MerchantCancellationResponse) SetChannelOrderNoNil(b bool)`

 SetChannelOrderNoNil sets the value for ChannelOrderNo to be an explicit nil

### UnsetChannelOrderNo
`func (o *MerchantCancellationResponse) UnsetChannelOrderNo()`

UnsetChannelOrderNo ensures that no value is present for ChannelOrderNo, not even an explicit nil
### GetLines

`func (o *MerchantCancellationResponse) GetLines() []MerchantCancellationLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantCancellationResponse) GetLinesOk() (*[]MerchantCancellationLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantCancellationResponse) SetLines(v []MerchantCancellationLineResponse)`

SetLines sets Lines field to given value.


### GetCreatedAt

`func (o *MerchantCancellationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantCancellationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantCancellationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantCancellationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetReason

`func (o *MerchantCancellationResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantCancellationResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantCancellationResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantCancellationResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MerchantCancellationResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MerchantCancellationResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *MerchantCancellationResponse) GetReasonCode() MancoReason`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *MerchantCancellationResponse) GetReasonCodeOk() (*MancoReason, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *MerchantCancellationResponse) SetReasonCode(v MancoReason)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *MerchantCancellationResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


