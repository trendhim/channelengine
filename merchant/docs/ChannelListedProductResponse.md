# ChannelListedProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel | [optional] 
**ChannelStatus** | Pointer to [**ListedProductChannelStatus**](ListedProductChannelStatus.md) |  | [optional] 
**Ean** | Pointer to **NullableString** | EAN | [optional] 
**ExportStatus** | Pointer to [**ListedProductExportStatus**](ListedProductExportStatus.md) |  | [optional] 
**MerchantProductNo** | Pointer to **NullableString** | Your product number (SKU) | [optional] 
**LastExportedPrice** | Pointer to **NullableFloat32** | Your product last exported price | [optional] 
**LastExportedStock** | Pointer to **NullableInt32** | Your product last exported stock | [optional] 

## Methods

### NewChannelListedProductResponse

`func NewChannelListedProductResponse() *ChannelListedProductResponse`

NewChannelListedProductResponse instantiates a new ChannelListedProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelListedProductResponseWithDefaults

`func NewChannelListedProductResponseWithDefaults() *ChannelListedProductResponse`

NewChannelListedProductResponseWithDefaults instantiates a new ChannelListedProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelListedProductResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelListedProductResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelListedProductResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *ChannelListedProductResponse) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *ChannelListedProductResponse) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *ChannelListedProductResponse) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetChannelStatus

`func (o *ChannelListedProductResponse) GetChannelStatus() ListedProductChannelStatus`

GetChannelStatus returns the ChannelStatus field if non-nil, zero value otherwise.

### GetChannelStatusOk

`func (o *ChannelListedProductResponse) GetChannelStatusOk() (*ListedProductChannelStatus, bool)`

GetChannelStatusOk returns a tuple with the ChannelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelStatus

`func (o *ChannelListedProductResponse) SetChannelStatus(v ListedProductChannelStatus)`

SetChannelStatus sets ChannelStatus field to given value.

### HasChannelStatus

`func (o *ChannelListedProductResponse) HasChannelStatus() bool`

HasChannelStatus returns a boolean if a field has been set.

### GetEan

`func (o *ChannelListedProductResponse) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ChannelListedProductResponse) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ChannelListedProductResponse) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ChannelListedProductResponse) HasEan() bool`

HasEan returns a boolean if a field has been set.

### SetEanNil

`func (o *ChannelListedProductResponse) SetEanNil(b bool)`

 SetEanNil sets the value for Ean to be an explicit nil

### UnsetEan
`func (o *ChannelListedProductResponse) UnsetEan()`

UnsetEan ensures that no value is present for Ean, not even an explicit nil
### GetExportStatus

`func (o *ChannelListedProductResponse) GetExportStatus() ListedProductExportStatus`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *ChannelListedProductResponse) GetExportStatusOk() (*ListedProductExportStatus, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *ChannelListedProductResponse) SetExportStatus(v ListedProductExportStatus)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *ChannelListedProductResponse) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetMerchantProductNo

`func (o *ChannelListedProductResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelListedProductResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelListedProductResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelListedProductResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelListedProductResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelListedProductResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetLastExportedPrice

`func (o *ChannelListedProductResponse) GetLastExportedPrice() float32`

GetLastExportedPrice returns the LastExportedPrice field if non-nil, zero value otherwise.

### GetLastExportedPriceOk

`func (o *ChannelListedProductResponse) GetLastExportedPriceOk() (*float32, bool)`

GetLastExportedPriceOk returns a tuple with the LastExportedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExportedPrice

`func (o *ChannelListedProductResponse) SetLastExportedPrice(v float32)`

SetLastExportedPrice sets LastExportedPrice field to given value.

### HasLastExportedPrice

`func (o *ChannelListedProductResponse) HasLastExportedPrice() bool`

HasLastExportedPrice returns a boolean if a field has been set.

### SetLastExportedPriceNil

`func (o *ChannelListedProductResponse) SetLastExportedPriceNil(b bool)`

 SetLastExportedPriceNil sets the value for LastExportedPrice to be an explicit nil

### UnsetLastExportedPrice
`func (o *ChannelListedProductResponse) UnsetLastExportedPrice()`

UnsetLastExportedPrice ensures that no value is present for LastExportedPrice, not even an explicit nil
### GetLastExportedStock

`func (o *ChannelListedProductResponse) GetLastExportedStock() int32`

GetLastExportedStock returns the LastExportedStock field if non-nil, zero value otherwise.

### GetLastExportedStockOk

`func (o *ChannelListedProductResponse) GetLastExportedStockOk() (*int32, bool)`

GetLastExportedStockOk returns a tuple with the LastExportedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExportedStock

`func (o *ChannelListedProductResponse) SetLastExportedStock(v int32)`

SetLastExportedStock sets LastExportedStock field to given value.

### HasLastExportedStock

`func (o *ChannelListedProductResponse) HasLastExportedStock() bool`

HasLastExportedStock returns a boolean if a field has been set.

### SetLastExportedStockNil

`func (o *ChannelListedProductResponse) SetLastExportedStockNil(b bool)`

 SetLastExportedStockNil sets the value for LastExportedStock to be an explicit nil

### UnsetLastExportedStock
`func (o *ChannelListedProductResponse) UnsetLastExportedStock()`

UnsetLastExportedStock ensures that no value is present for LastExportedStock, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


