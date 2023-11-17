# MerchantProductWithBuyBoxPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **NullableString** | Product SKU number | [optional] 
**Gtin** | Pointer to **NullableString** | Product GTIN | [optional] 
**BuyBoxPrice** | Pointer to **NullableFloat32** | Price of Buy box winner (excl. shipping cost)  Note: not all channels have a separate shipping cost field (e.g. bol.com), so can be the same as BuyBoxPriceInclShipping | [optional] 
**BuyBoxPriceInclShipping** | Pointer to **NullableFloat32** | Price of Buy box winner (incl. shipping cost).  If null, then there is no data or no Buy box winner | [optional] 
**IsMerchantBuyBoxWinner** | Pointer to **bool** | Are you the Buy box winner or not? | [optional] 

## Methods

### NewMerchantProductWithBuyBoxPrice

`func NewMerchantProductWithBuyBoxPrice() *MerchantProductWithBuyBoxPrice`

NewMerchantProductWithBuyBoxPrice instantiates a new MerchantProductWithBuyBoxPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductWithBuyBoxPriceWithDefaults

`func NewMerchantProductWithBuyBoxPriceWithDefaults() *MerchantProductWithBuyBoxPrice`

NewMerchantProductWithBuyBoxPriceWithDefaults instantiates a new MerchantProductWithBuyBoxPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *MerchantProductWithBuyBoxPrice) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *MerchantProductWithBuyBoxPrice) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *MerchantProductWithBuyBoxPrice) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *MerchantProductWithBuyBoxPrice) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *MerchantProductWithBuyBoxPrice) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *MerchantProductWithBuyBoxPrice) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetGtin

`func (o *MerchantProductWithBuyBoxPrice) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *MerchantProductWithBuyBoxPrice) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *MerchantProductWithBuyBoxPrice) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *MerchantProductWithBuyBoxPrice) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### SetGtinNil

`func (o *MerchantProductWithBuyBoxPrice) SetGtinNil(b bool)`

 SetGtinNil sets the value for Gtin to be an explicit nil

### UnsetGtin
`func (o *MerchantProductWithBuyBoxPrice) UnsetGtin()`

UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
### GetBuyBoxPrice

`func (o *MerchantProductWithBuyBoxPrice) GetBuyBoxPrice() float32`

GetBuyBoxPrice returns the BuyBoxPrice field if non-nil, zero value otherwise.

### GetBuyBoxPriceOk

`func (o *MerchantProductWithBuyBoxPrice) GetBuyBoxPriceOk() (*float32, bool)`

GetBuyBoxPriceOk returns a tuple with the BuyBoxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyBoxPrice

`func (o *MerchantProductWithBuyBoxPrice) SetBuyBoxPrice(v float32)`

SetBuyBoxPrice sets BuyBoxPrice field to given value.

### HasBuyBoxPrice

`func (o *MerchantProductWithBuyBoxPrice) HasBuyBoxPrice() bool`

HasBuyBoxPrice returns a boolean if a field has been set.

### SetBuyBoxPriceNil

`func (o *MerchantProductWithBuyBoxPrice) SetBuyBoxPriceNil(b bool)`

 SetBuyBoxPriceNil sets the value for BuyBoxPrice to be an explicit nil

### UnsetBuyBoxPrice
`func (o *MerchantProductWithBuyBoxPrice) UnsetBuyBoxPrice()`

UnsetBuyBoxPrice ensures that no value is present for BuyBoxPrice, not even an explicit nil
### GetBuyBoxPriceInclShipping

`func (o *MerchantProductWithBuyBoxPrice) GetBuyBoxPriceInclShipping() float32`

GetBuyBoxPriceInclShipping returns the BuyBoxPriceInclShipping field if non-nil, zero value otherwise.

### GetBuyBoxPriceInclShippingOk

`func (o *MerchantProductWithBuyBoxPrice) GetBuyBoxPriceInclShippingOk() (*float32, bool)`

GetBuyBoxPriceInclShippingOk returns a tuple with the BuyBoxPriceInclShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyBoxPriceInclShipping

`func (o *MerchantProductWithBuyBoxPrice) SetBuyBoxPriceInclShipping(v float32)`

SetBuyBoxPriceInclShipping sets BuyBoxPriceInclShipping field to given value.

### HasBuyBoxPriceInclShipping

`func (o *MerchantProductWithBuyBoxPrice) HasBuyBoxPriceInclShipping() bool`

HasBuyBoxPriceInclShipping returns a boolean if a field has been set.

### SetBuyBoxPriceInclShippingNil

`func (o *MerchantProductWithBuyBoxPrice) SetBuyBoxPriceInclShippingNil(b bool)`

 SetBuyBoxPriceInclShippingNil sets the value for BuyBoxPriceInclShipping to be an explicit nil

### UnsetBuyBoxPriceInclShipping
`func (o *MerchantProductWithBuyBoxPrice) UnsetBuyBoxPriceInclShipping()`

UnsetBuyBoxPriceInclShipping ensures that no value is present for BuyBoxPriceInclShipping, not even an explicit nil
### GetIsMerchantBuyBoxWinner

`func (o *MerchantProductWithBuyBoxPrice) GetIsMerchantBuyBoxWinner() bool`

GetIsMerchantBuyBoxWinner returns the IsMerchantBuyBoxWinner field if non-nil, zero value otherwise.

### GetIsMerchantBuyBoxWinnerOk

`func (o *MerchantProductWithBuyBoxPrice) GetIsMerchantBuyBoxWinnerOk() (*bool, bool)`

GetIsMerchantBuyBoxWinnerOk returns a tuple with the IsMerchantBuyBoxWinner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchantBuyBoxWinner

`func (o *MerchantProductWithBuyBoxPrice) SetIsMerchantBuyBoxWinner(v bool)`

SetIsMerchantBuyBoxWinner sets IsMerchantBuyBoxWinner field to given value.

### HasIsMerchantBuyBoxWinner

`func (o *MerchantProductWithBuyBoxPrice) HasIsMerchantBuyBoxWinner() bool`

HasIsMerchantBuyBoxWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


