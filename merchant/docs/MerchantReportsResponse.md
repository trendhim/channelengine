# MerchantReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ChannelSettlementNo** | Pointer to **NullableString** |  | [optional] 
**ChannelSellerNo** | Pointer to **NullableString** |  | [optional] 
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TransactionsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewMerchantReportsResponse

`func NewMerchantReportsResponse() *MerchantReportsResponse`

NewMerchantReportsResponse instantiates a new MerchantReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReportsResponseWithDefaults

`func NewMerchantReportsResponseWithDefaults() *MerchantReportsResponse`

NewMerchantReportsResponseWithDefaults instantiates a new MerchantReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantReportsResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantReportsResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantReportsResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantReportsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelSettlementNo

`func (o *MerchantReportsResponse) GetChannelSettlementNo() string`

GetChannelSettlementNo returns the ChannelSettlementNo field if non-nil, zero value otherwise.

### GetChannelSettlementNoOk

`func (o *MerchantReportsResponse) GetChannelSettlementNoOk() (*string, bool)`

GetChannelSettlementNoOk returns a tuple with the ChannelSettlementNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSettlementNo

`func (o *MerchantReportsResponse) SetChannelSettlementNo(v string)`

SetChannelSettlementNo sets ChannelSettlementNo field to given value.

### HasChannelSettlementNo

`func (o *MerchantReportsResponse) HasChannelSettlementNo() bool`

HasChannelSettlementNo returns a boolean if a field has been set.

### SetChannelSettlementNoNil

`func (o *MerchantReportsResponse) SetChannelSettlementNoNil(b bool)`

 SetChannelSettlementNoNil sets the value for ChannelSettlementNo to be an explicit nil

### UnsetChannelSettlementNo
`func (o *MerchantReportsResponse) UnsetChannelSettlementNo()`

UnsetChannelSettlementNo ensures that no value is present for ChannelSettlementNo, not even an explicit nil
### GetChannelSellerNo

`func (o *MerchantReportsResponse) GetChannelSellerNo() string`

GetChannelSellerNo returns the ChannelSellerNo field if non-nil, zero value otherwise.

### GetChannelSellerNoOk

`func (o *MerchantReportsResponse) GetChannelSellerNoOk() (*string, bool)`

GetChannelSellerNoOk returns a tuple with the ChannelSellerNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSellerNo

`func (o *MerchantReportsResponse) SetChannelSellerNo(v string)`

SetChannelSellerNo sets ChannelSellerNo field to given value.

### HasChannelSellerNo

`func (o *MerchantReportsResponse) HasChannelSellerNo() bool`

HasChannelSellerNo returns a boolean if a field has been set.

### SetChannelSellerNoNil

`func (o *MerchantReportsResponse) SetChannelSellerNoNil(b bool)`

 SetChannelSellerNoNil sets the value for ChannelSellerNo to be an explicit nil

### UnsetChannelSellerNo
`func (o *MerchantReportsResponse) UnsetChannelSellerNo()`

UnsetChannelSellerNo ensures that no value is present for ChannelSellerNo, not even an explicit nil
### GetChannelId

`func (o *MerchantReportsResponse) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MerchantReportsResponse) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MerchantReportsResponse) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MerchantReportsResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MerchantReportsResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MerchantReportsResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetStartDate

`func (o *MerchantReportsResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MerchantReportsResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MerchantReportsResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MerchantReportsResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *MerchantReportsResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MerchantReportsResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MerchantReportsResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MerchantReportsResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MerchantReportsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantReportsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantReportsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantReportsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantReportsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantReportsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantReportsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantReportsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTransactionsCount

`func (o *MerchantReportsResponse) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *MerchantReportsResponse) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *MerchantReportsResponse) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.

### HasTransactionsCount

`func (o *MerchantReportsResponse) HasTransactionsCount() bool`

HasTransactionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


