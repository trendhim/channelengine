# MerchantReturnLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant (sku). | [optional] 
**OrderLine** | Pointer to [**MerchantOrderLineResponse**](MerchantOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**StockLocation** | Pointer to [**MerchantStockLocationResponse**](MerchantStockLocationResponse.md) |  | [optional] 
**Quantity** | **int32** | Number of items of the product in this return. | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the returnline. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantReturnLineResponse

`func NewMerchantReturnLineResponse(quantity int32, ) *MerchantReturnLineResponse`

NewMerchantReturnLineResponse instantiates a new MerchantReturnLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnLineResponseWithDefaults

`func NewMerchantReturnLineResponseWithDefaults() *MerchantReturnLineResponse`

NewMerchantReturnLineResponseWithDefaults instantiates a new MerchantReturnLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantReturnLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantReturnLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantReturnLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantReturnLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantReturnLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantReturnLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetOrderLine

`func (o *MerchantReturnLineResponse) GetOrderLine() MerchantOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *MerchantReturnLineResponse) GetOrderLineOk() (*MerchantOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *MerchantReturnLineResponse) SetOrderLine(v MerchantOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *MerchantReturnLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *MerchantReturnLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *MerchantReturnLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *MerchantReturnLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *MerchantReturnLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetStockLocation

`func (o *MerchantReturnLineResponse) GetStockLocation() MerchantStockLocationResponse`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *MerchantReturnLineResponse) GetStockLocationOk() (*MerchantStockLocationResponse, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *MerchantReturnLineResponse) SetStockLocation(v MerchantStockLocationResponse)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *MerchantReturnLineResponse) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetQuantity

`func (o *MerchantReturnLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantReturnLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantReturnLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetExtraData

`func (o *MerchantReturnLineResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantReturnLineResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantReturnLineResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantReturnLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantReturnLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantReturnLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


