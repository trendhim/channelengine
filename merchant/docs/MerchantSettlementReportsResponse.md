# MerchantSettlementReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementId** | Pointer to **int32** |  | [optional] 
**ChannelSettlementNo** | Pointer to **NullableString** |  | [optional] 
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrentReserveAmount** | Pointer to **NullableFloat32** |  | [optional] 
**PreviousReserveAmount** | Pointer to **NullableFloat32** |  | [optional] 
**PaymentAmount** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMerchantSettlementReportsResponse

`func NewMerchantSettlementReportsResponse() *MerchantSettlementReportsResponse`

NewMerchantSettlementReportsResponse instantiates a new MerchantSettlementReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSettlementReportsResponseWithDefaults

`func NewMerchantSettlementReportsResponseWithDefaults() *MerchantSettlementReportsResponse`

NewMerchantSettlementReportsResponseWithDefaults instantiates a new MerchantSettlementReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementId

`func (o *MerchantSettlementReportsResponse) GetSettlementId() int32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *MerchantSettlementReportsResponse) GetSettlementIdOk() (*int32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *MerchantSettlementReportsResponse) SetSettlementId(v int32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *MerchantSettlementReportsResponse) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetChannelSettlementNo

`func (o *MerchantSettlementReportsResponse) GetChannelSettlementNo() string`

GetChannelSettlementNo returns the ChannelSettlementNo field if non-nil, zero value otherwise.

### GetChannelSettlementNoOk

`func (o *MerchantSettlementReportsResponse) GetChannelSettlementNoOk() (*string, bool)`

GetChannelSettlementNoOk returns a tuple with the ChannelSettlementNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSettlementNo

`func (o *MerchantSettlementReportsResponse) SetChannelSettlementNo(v string)`

SetChannelSettlementNo sets ChannelSettlementNo field to given value.

### HasChannelSettlementNo

`func (o *MerchantSettlementReportsResponse) HasChannelSettlementNo() bool`

HasChannelSettlementNo returns a boolean if a field has been set.

### SetChannelSettlementNoNil

`func (o *MerchantSettlementReportsResponse) SetChannelSettlementNoNil(b bool)`

 SetChannelSettlementNoNil sets the value for ChannelSettlementNo to be an explicit nil

### UnsetChannelSettlementNo
`func (o *MerchantSettlementReportsResponse) UnsetChannelSettlementNo()`

UnsetChannelSettlementNo ensures that no value is present for ChannelSettlementNo, not even an explicit nil
### GetChannelId

`func (o *MerchantSettlementReportsResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantSettlementReportsResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantSettlementReportsResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantSettlementReportsResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantSettlementReportsResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantSettlementReportsResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetStartDate

`func (o *MerchantSettlementReportsResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MerchantSettlementReportsResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MerchantSettlementReportsResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MerchantSettlementReportsResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *MerchantSettlementReportsResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MerchantSettlementReportsResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MerchantSettlementReportsResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MerchantSettlementReportsResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MerchantSettlementReportsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantSettlementReportsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantSettlementReportsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantSettlementReportsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantSettlementReportsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantSettlementReportsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantSettlementReportsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantSettlementReportsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCurrentReserveAmount

`func (o *MerchantSettlementReportsResponse) GetCurrentReserveAmount() float32`

GetCurrentReserveAmount returns the CurrentReserveAmount field if non-nil, zero value otherwise.

### GetCurrentReserveAmountOk

`func (o *MerchantSettlementReportsResponse) GetCurrentReserveAmountOk() (*float32, bool)`

GetCurrentReserveAmountOk returns a tuple with the CurrentReserveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReserveAmount

`func (o *MerchantSettlementReportsResponse) SetCurrentReserveAmount(v float32)`

SetCurrentReserveAmount sets CurrentReserveAmount field to given value.

### HasCurrentReserveAmount

`func (o *MerchantSettlementReportsResponse) HasCurrentReserveAmount() bool`

HasCurrentReserveAmount returns a boolean if a field has been set.

### SetCurrentReserveAmountNil

`func (o *MerchantSettlementReportsResponse) SetCurrentReserveAmountNil(b bool)`

 SetCurrentReserveAmountNil sets the value for CurrentReserveAmount to be an explicit nil

### UnsetCurrentReserveAmount
`func (o *MerchantSettlementReportsResponse) UnsetCurrentReserveAmount()`

UnsetCurrentReserveAmount ensures that no value is present for CurrentReserveAmount, not even an explicit nil
### GetPreviousReserveAmount

`func (o *MerchantSettlementReportsResponse) GetPreviousReserveAmount() float32`

GetPreviousReserveAmount returns the PreviousReserveAmount field if non-nil, zero value otherwise.

### GetPreviousReserveAmountOk

`func (o *MerchantSettlementReportsResponse) GetPreviousReserveAmountOk() (*float32, bool)`

GetPreviousReserveAmountOk returns a tuple with the PreviousReserveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousReserveAmount

`func (o *MerchantSettlementReportsResponse) SetPreviousReserveAmount(v float32)`

SetPreviousReserveAmount sets PreviousReserveAmount field to given value.

### HasPreviousReserveAmount

`func (o *MerchantSettlementReportsResponse) HasPreviousReserveAmount() bool`

HasPreviousReserveAmount returns a boolean if a field has been set.

### SetPreviousReserveAmountNil

`func (o *MerchantSettlementReportsResponse) SetPreviousReserveAmountNil(b bool)`

 SetPreviousReserveAmountNil sets the value for PreviousReserveAmount to be an explicit nil

### UnsetPreviousReserveAmount
`func (o *MerchantSettlementReportsResponse) UnsetPreviousReserveAmount()`

UnsetPreviousReserveAmount ensures that no value is present for PreviousReserveAmount, not even an explicit nil
### GetPaymentAmount

`func (o *MerchantSettlementReportsResponse) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *MerchantSettlementReportsResponse) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *MerchantSettlementReportsResponse) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *MerchantSettlementReportsResponse) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### SetPaymentAmountNil

`func (o *MerchantSettlementReportsResponse) SetPaymentAmountNil(b bool)`

 SetPaymentAmountNil sets the value for PaymentAmount to be an explicit nil

### UnsetPaymentAmount
`func (o *MerchantSettlementReportsResponse) UnsetPaymentAmount()`

UnsetPaymentAmount ensures that no value is present for PaymentAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


