# MerchantOfferGetStockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** | The product SKU. | [optional] 
**StockLocationId** | Pointer to **int32** | The ChannelEngine id of the stock location. | [optional] 
**Stock** | Pointer to **int32** | The quantity of products in stock at the stock location. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp of the last stock update for the stock location. | [optional] 

## Methods

### NewMerchantOfferGetStockResponse

`func NewMerchantOfferGetStockResponse() *MerchantOfferGetStockResponse`

NewMerchantOfferGetStockResponse instantiates a new MerchantOfferGetStockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOfferGetStockResponseWithDefaults

`func NewMerchantOfferGetStockResponseWithDefaults() *MerchantOfferGetStockResponse`

NewMerchantOfferGetStockResponseWithDefaults instantiates a new MerchantOfferGetStockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantOfferGetStockResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantOfferGetStockResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantOfferGetStockResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantOfferGetStockResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantOfferGetStockResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantOfferGetStockResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetStockLocationId

`func (o *MerchantOfferGetStockResponse) GetStockLocationId() int32`

GetStockLocationId returns the StockLocationId field if non-nil, zero value otherwise.

### GetStockLocationIdOk

`func (o *MerchantOfferGetStockResponse) GetStockLocationIdOk() (*int32, bool)`

GetStockLocationIdOk returns a tuple with the StockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationId

`func (o *MerchantOfferGetStockResponse) SetStockLocationId(v int32)`

SetStockLocationId sets StockLocationId field to given value.

### HasStockLocationId

`func (o *MerchantOfferGetStockResponse) HasStockLocationId() bool`

HasStockLocationId returns a boolean if a field has been set.

### GetStock

`func (o *MerchantOfferGetStockResponse) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *MerchantOfferGetStockResponse) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *MerchantOfferGetStockResponse) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *MerchantOfferGetStockResponse) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantOfferGetStockResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantOfferGetStockResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantOfferGetStockResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantOfferGetStockResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


