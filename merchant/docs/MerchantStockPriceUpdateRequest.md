# MerchantStockPriceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 
**Stock** | Pointer to **NullableInt32** | The stock of the product. Should not be negative. | [optional] 
**Price** | Pointer to **NullableFloat32** | The price of the product. Should not be negative. | [optional] 
**StockLocationId** | Pointer to **NullableInt32** | The stock location id of the updated stock. If not provided, the stock from the default stock location will be updated. | [optional] 

## Methods

### NewMerchantStockPriceUpdateRequest

`func NewMerchantStockPriceUpdateRequest(merchantProductNo string, ) *MerchantStockPriceUpdateRequest`

NewMerchantStockPriceUpdateRequest instantiates a new MerchantStockPriceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantStockPriceUpdateRequestWithDefaults

`func NewMerchantStockPriceUpdateRequestWithDefaults() *MerchantStockPriceUpdateRequest`

NewMerchantStockPriceUpdateRequestWithDefaults instantiates a new MerchantStockPriceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantStockPriceUpdateRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantStockPriceUpdateRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantStockPriceUpdateRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetStock

`func (o *MerchantStockPriceUpdateRequest) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *MerchantStockPriceUpdateRequest) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *MerchantStockPriceUpdateRequest) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *MerchantStockPriceUpdateRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### SetStockNil

`func (o *MerchantStockPriceUpdateRequest) SetStockNil(b bool)`

 SetStockNil sets the value for Stock to be an explicit nil

### UnsetStock
`func (o *MerchantStockPriceUpdateRequest) UnsetStock()`

UnsetStock ensures that no value is present for Stock, not even an explicit nil
### GetPrice

`func (o *MerchantStockPriceUpdateRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MerchantStockPriceUpdateRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MerchantStockPriceUpdateRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MerchantStockPriceUpdateRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MerchantStockPriceUpdateRequest) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MerchantStockPriceUpdateRequest) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetStockLocationId

`func (o *MerchantStockPriceUpdateRequest) GetStockLocationId() int32`

GetStockLocationId returns the StockLocationId field if non-nil, zero value otherwise.

### GetStockLocationIdOk

`func (o *MerchantStockPriceUpdateRequest) GetStockLocationIdOk() (*int32, bool)`

GetStockLocationIdOk returns a tuple with the StockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationId

`func (o *MerchantStockPriceUpdateRequest) SetStockLocationId(v int32)`

SetStockLocationId sets StockLocationId field to given value.

### HasStockLocationId

`func (o *MerchantStockPriceUpdateRequest) HasStockLocationId() bool`

HasStockLocationId returns a boolean if a field has been set.

### SetStockLocationIdNil

`func (o *MerchantStockPriceUpdateRequest) SetStockLocationIdNil(b bool)`

 SetStockLocationIdNil sets the value for StockLocationId to be an explicit nil

### UnsetStockLocationId
`func (o *MerchantStockPriceUpdateRequest) UnsetStockLocationId()`

UnsetStockLocationId ensures that no value is present for StockLocationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


