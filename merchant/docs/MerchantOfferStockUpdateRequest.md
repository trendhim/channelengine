# MerchantOfferStockUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 
**StockLocations** | [**[]MerchantStockLocationUpdateRequest**](MerchantStockLocationUpdateRequest.md) | Stock locations data | 

## Methods

### NewMerchantOfferStockUpdateRequest

`func NewMerchantOfferStockUpdateRequest(merchantProductNo string, stockLocations []MerchantStockLocationUpdateRequest, ) *MerchantOfferStockUpdateRequest`

NewMerchantOfferStockUpdateRequest instantiates a new MerchantOfferStockUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOfferStockUpdateRequestWithDefaults

`func NewMerchantOfferStockUpdateRequestWithDefaults() *MerchantOfferStockUpdateRequest`

NewMerchantOfferStockUpdateRequestWithDefaults instantiates a new MerchantOfferStockUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantOfferStockUpdateRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantOfferStockUpdateRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantOfferStockUpdateRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetStockLocations

`func (o *MerchantOfferStockUpdateRequest) GetStockLocations() []MerchantStockLocationUpdateRequest`

GetStockLocations returns the StockLocations field if non-nil, zero value otherwise.

### GetStockLocationsOk

`func (o *MerchantOfferStockUpdateRequest) GetStockLocationsOk() (*[]MerchantStockLocationUpdateRequest, bool)`

GetStockLocationsOk returns a tuple with the StockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocations

`func (o *MerchantOfferStockUpdateRequest) SetStockLocations(v []MerchantStockLocationUpdateRequest)`

SetStockLocations sets StockLocations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


