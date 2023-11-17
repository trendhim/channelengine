# ChannelOfferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel. | [optional] 
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Stock** | Pointer to **int32** |  | [optional] 
**MappedFields** | Pointer to **map[string]string** | A channel can require certain fields to be available. The channel  can provide ChannelEngine with this fields after which the merchants  will be able to fill in this field using custom conditions in ChannelEngine. | [optional] 

## Methods

### NewChannelOfferResponse

`func NewChannelOfferResponse() *ChannelOfferResponse`

NewChannelOfferResponse instantiates a new ChannelOfferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOfferResponseWithDefaults

`func NewChannelOfferResponseWithDefaults() *ChannelOfferResponse`

NewChannelOfferResponseWithDefaults instantiates a new ChannelOfferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelOfferResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelOfferResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelOfferResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *ChannelOfferResponse) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *ChannelOfferResponse) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *ChannelOfferResponse) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetMerchantProductNo

`func (o *ChannelOfferResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelOfferResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelOfferResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelOfferResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelOfferResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelOfferResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetPrice

`func (o *ChannelOfferResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ChannelOfferResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ChannelOfferResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ChannelOfferResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStock

`func (o *ChannelOfferResponse) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ChannelOfferResponse) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ChannelOfferResponse) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ChannelOfferResponse) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetMappedFields

`func (o *ChannelOfferResponse) GetMappedFields() map[string]string`

GetMappedFields returns the MappedFields field if non-nil, zero value otherwise.

### GetMappedFieldsOk

`func (o *ChannelOfferResponse) GetMappedFieldsOk() (*map[string]string, bool)`

GetMappedFieldsOk returns a tuple with the MappedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedFields

`func (o *ChannelOfferResponse) SetMappedFields(v map[string]string)`

SetMappedFields sets MappedFields field to given value.

### HasMappedFields

`func (o *ChannelOfferResponse) HasMappedFields() bool`

HasMappedFields returns a boolean if a field has been set.

### SetMappedFieldsNil

`func (o *ChannelOfferResponse) SetMappedFieldsNil(b bool)`

 SetMappedFieldsNil sets the value for MappedFields to be an explicit nil

### UnsetMappedFields
`func (o *ChannelOfferResponse) UnsetMappedFields()`

UnsetMappedFields ensures that no value is present for MappedFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


