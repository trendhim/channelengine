# MerchantStockLocationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stock** | Pointer to **int32** | The stock of the product. Should not be negative. | [optional] 
**StockLocationId** | Pointer to **NullableInt32** | The stock location id of updated stok.  If not provided stock from default stock location will be updated. | [optional] 

## Methods

### NewMerchantStockLocationUpdateRequest

`func NewMerchantStockLocationUpdateRequest() *MerchantStockLocationUpdateRequest`

NewMerchantStockLocationUpdateRequest instantiates a new MerchantStockLocationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantStockLocationUpdateRequestWithDefaults

`func NewMerchantStockLocationUpdateRequestWithDefaults() *MerchantStockLocationUpdateRequest`

NewMerchantStockLocationUpdateRequestWithDefaults instantiates a new MerchantStockLocationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStock

`func (o *MerchantStockLocationUpdateRequest) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *MerchantStockLocationUpdateRequest) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *MerchantStockLocationUpdateRequest) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *MerchantStockLocationUpdateRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetStockLocationId

`func (o *MerchantStockLocationUpdateRequest) GetStockLocationId() int32`

GetStockLocationId returns the StockLocationId field if non-nil, zero value otherwise.

### GetStockLocationIdOk

`func (o *MerchantStockLocationUpdateRequest) GetStockLocationIdOk() (*int32, bool)`

GetStockLocationIdOk returns a tuple with the StockLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocationId

`func (o *MerchantStockLocationUpdateRequest) SetStockLocationId(v int32)`

SetStockLocationId sets StockLocationId field to given value.

### HasStockLocationId

`func (o *MerchantStockLocationUpdateRequest) HasStockLocationId() bool`

HasStockLocationId returns a boolean if a field has been set.

### SetStockLocationIdNil

`func (o *MerchantStockLocationUpdateRequest) SetStockLocationIdNil(b bool)`

 SetStockLocationIdNil sets the value for StockLocationId to be an explicit nil

### UnsetStockLocationId
`func (o *MerchantStockLocationUpdateRequest) UnsetStockLocationId()`

UnsetStockLocationId ensures that no value is present for StockLocationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


