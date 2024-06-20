# MerchantFulfillmentStockStockLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** | The product SKU. | [optional] 
**StockLocations** | Pointer to [**[]MerchantFulfillmentStockLocationItemResponse**](MerchantFulfillmentStockLocationItemResponse.md) | The ChannelEngine id of the stock location. | [optional] 

## Methods

### NewMerchantFulfillmentStockStockLocationsResponse

`func NewMerchantFulfillmentStockStockLocationsResponse() *MerchantFulfillmentStockStockLocationsResponse`

NewMerchantFulfillmentStockStockLocationsResponse instantiates a new MerchantFulfillmentStockStockLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantFulfillmentStockStockLocationsResponseWithDefaults

`func NewMerchantFulfillmentStockStockLocationsResponseWithDefaults() *MerchantFulfillmentStockStockLocationsResponse`

NewMerchantFulfillmentStockStockLocationsResponseWithDefaults instantiates a new MerchantFulfillmentStockStockLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantFulfillmentStockStockLocationsResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantFulfillmentStockStockLocationsResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantFulfillmentStockStockLocationsResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantFulfillmentStockStockLocationsResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantFulfillmentStockStockLocationsResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantFulfillmentStockStockLocationsResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetStockLocations

`func (o *MerchantFulfillmentStockStockLocationsResponse) GetStockLocations() []MerchantFulfillmentStockLocationItemResponse`

GetStockLocations returns the StockLocations field if non-nil, zero value otherwise.

### GetStockLocationsOk

`func (o *MerchantFulfillmentStockStockLocationsResponse) GetStockLocationsOk() (*[]MerchantFulfillmentStockLocationItemResponse, bool)`

GetStockLocationsOk returns a tuple with the StockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocations

`func (o *MerchantFulfillmentStockStockLocationsResponse) SetStockLocations(v []MerchantFulfillmentStockLocationItemResponse)`

SetStockLocations sets StockLocations field to given value.

### HasStockLocations

`func (o *MerchantFulfillmentStockStockLocationsResponse) HasStockLocations() bool`

HasStockLocations returns a boolean if a field has been set.

### SetStockLocationsNil

`func (o *MerchantFulfillmentStockStockLocationsResponse) SetStockLocationsNil(b bool)`

 SetStockLocationsNil sets the value for StockLocations to be an explicit nil

### UnsetStockLocations
`func (o *MerchantFulfillmentStockStockLocationsResponse) UnsetStockLocations()`

UnsetStockLocations ensures that no value is present for StockLocations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


