# ChannelCancellationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOrderNo** | **string** | The unique order reference used by the Channel. | 
**MerchantCancellationNo** | **string** | The unique cancellation reference used by the Merchant. | 
**Lines** | [**[]ChannelCancellationLineResponse**](ChannelCancellationLineResponse.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The date at which the cancellation was created in ChannelEngine. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date at which the cancellation was last modified in ChannelEngine. | [optional] 
**Reason** | Pointer to **NullableString** | Reason for cancellation (text). | [optional] 
**ReasonCode** | Pointer to [**MancoReason**](MancoReason.md) |  | [optional] 

## Methods

### NewChannelCancellationResponse

`func NewChannelCancellationResponse(channelOrderNo string, merchantCancellationNo string, lines []ChannelCancellationLineResponse, ) *ChannelCancellationResponse`

NewChannelCancellationResponse instantiates a new ChannelCancellationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelCancellationResponseWithDefaults

`func NewChannelCancellationResponseWithDefaults() *ChannelCancellationResponse`

NewChannelCancellationResponseWithDefaults instantiates a new ChannelCancellationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOrderNo

`func (o *ChannelCancellationResponse) GetChannelOrderNo() string`

GetChannelOrderNo returns the ChannelOrderNo field if non-nil, zero value otherwise.

### GetChannelOrderNoOk

`func (o *ChannelCancellationResponse) GetChannelOrderNoOk() (*string, bool)`

GetChannelOrderNoOk returns a tuple with the ChannelOrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOrderNo

`func (o *ChannelCancellationResponse) SetChannelOrderNo(v string)`

SetChannelOrderNo sets ChannelOrderNo field to given value.


### GetMerchantCancellationNo

`func (o *ChannelCancellationResponse) GetMerchantCancellationNo() string`

GetMerchantCancellationNo returns the MerchantCancellationNo field if non-nil, zero value otherwise.

### GetMerchantCancellationNoOk

`func (o *ChannelCancellationResponse) GetMerchantCancellationNoOk() (*string, bool)`

GetMerchantCancellationNoOk returns a tuple with the MerchantCancellationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCancellationNo

`func (o *ChannelCancellationResponse) SetMerchantCancellationNo(v string)`

SetMerchantCancellationNo sets MerchantCancellationNo field to given value.


### GetLines

`func (o *ChannelCancellationResponse) GetLines() []ChannelCancellationLineResponse`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelCancellationResponse) GetLinesOk() (*[]ChannelCancellationLineResponse, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelCancellationResponse) SetLines(v []ChannelCancellationLineResponse)`

SetLines sets Lines field to given value.


### GetCreatedAt

`func (o *ChannelCancellationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelCancellationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelCancellationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelCancellationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChannelCancellationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelCancellationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelCancellationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChannelCancellationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReason

`func (o *ChannelCancellationResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChannelCancellationResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChannelCancellationResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChannelCancellationResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ChannelCancellationResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ChannelCancellationResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *ChannelCancellationResponse) GetReasonCode() MancoReason`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ChannelCancellationResponse) GetReasonCodeOk() (*MancoReason, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ChannelCancellationResponse) SetReasonCode(v MancoReason)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ChannelCancellationResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


