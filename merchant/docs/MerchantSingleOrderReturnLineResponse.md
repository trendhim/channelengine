# MerchantSingleOrderReturnLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the Merchant (sku). | [optional] 
**AcceptedQuantity** | Pointer to **NullableInt32** | The accepted quantity of returned products in this orderline. | [optional] 
**RejectedQuantity** | Pointer to **NullableInt32** | The rejected quantity of returned products in this orderline. | [optional] 
**OrderLine** | Pointer to [**MerchantOrderLineResponse**](MerchantOrderLineResponse.md) |  | [optional] 
**ShipmentStatus** | Pointer to [**ShipmentLineStatus**](ShipmentLineStatus.md) |  | [optional] 
**Quantity** | **int32** | Number of items of the product in this return. | 
**ExtraData** | Pointer to **map[string]string** | Extra data on the returnline. Each item must have an unqiue key | [optional] 

## Methods

### NewMerchantSingleOrderReturnLineResponse

`func NewMerchantSingleOrderReturnLineResponse(quantity int32, ) *MerchantSingleOrderReturnLineResponse`

NewMerchantSingleOrderReturnLineResponse instantiates a new MerchantSingleOrderReturnLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSingleOrderReturnLineResponseWithDefaults

`func NewMerchantSingleOrderReturnLineResponseWithDefaults() *MerchantSingleOrderReturnLineResponse`

NewMerchantSingleOrderReturnLineResponseWithDefaults instantiates a new MerchantSingleOrderReturnLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantSingleOrderReturnLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantSingleOrderReturnLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantSingleOrderReturnLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantSingleOrderReturnLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantSingleOrderReturnLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantSingleOrderReturnLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetAcceptedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) GetAcceptedQuantity() int32`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *MerchantSingleOrderReturnLineResponse) GetAcceptedQuantityOk() (*int32, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) SetAcceptedQuantity(v int32)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.

### HasAcceptedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) HasAcceptedQuantity() bool`

HasAcceptedQuantity returns a boolean if a field has been set.

### SetAcceptedQuantityNil

`func (o *MerchantSingleOrderReturnLineResponse) SetAcceptedQuantityNil(b bool)`

 SetAcceptedQuantityNil sets the value for AcceptedQuantity to be an explicit nil

### UnsetAcceptedQuantity
`func (o *MerchantSingleOrderReturnLineResponse) UnsetAcceptedQuantity()`

UnsetAcceptedQuantity ensures that no value is present for AcceptedQuantity, not even an explicit nil
### GetRejectedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) GetRejectedQuantity() int32`

GetRejectedQuantity returns the RejectedQuantity field if non-nil, zero value otherwise.

### GetRejectedQuantityOk

`func (o *MerchantSingleOrderReturnLineResponse) GetRejectedQuantityOk() (*int32, bool)`

GetRejectedQuantityOk returns a tuple with the RejectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) SetRejectedQuantity(v int32)`

SetRejectedQuantity sets RejectedQuantity field to given value.

### HasRejectedQuantity

`func (o *MerchantSingleOrderReturnLineResponse) HasRejectedQuantity() bool`

HasRejectedQuantity returns a boolean if a field has been set.

### SetRejectedQuantityNil

`func (o *MerchantSingleOrderReturnLineResponse) SetRejectedQuantityNil(b bool)`

 SetRejectedQuantityNil sets the value for RejectedQuantity to be an explicit nil

### UnsetRejectedQuantity
`func (o *MerchantSingleOrderReturnLineResponse) UnsetRejectedQuantity()`

UnsetRejectedQuantity ensures that no value is present for RejectedQuantity, not even an explicit nil
### GetOrderLine

`func (o *MerchantSingleOrderReturnLineResponse) GetOrderLine() MerchantOrderLineResponse`

GetOrderLine returns the OrderLine field if non-nil, zero value otherwise.

### GetOrderLineOk

`func (o *MerchantSingleOrderReturnLineResponse) GetOrderLineOk() (*MerchantOrderLineResponse, bool)`

GetOrderLineOk returns a tuple with the OrderLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLine

`func (o *MerchantSingleOrderReturnLineResponse) SetOrderLine(v MerchantOrderLineResponse)`

SetOrderLine sets OrderLine field to given value.

### HasOrderLine

`func (o *MerchantSingleOrderReturnLineResponse) HasOrderLine() bool`

HasOrderLine returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *MerchantSingleOrderReturnLineResponse) GetShipmentStatus() ShipmentLineStatus`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *MerchantSingleOrderReturnLineResponse) GetShipmentStatusOk() (*ShipmentLineStatus, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *MerchantSingleOrderReturnLineResponse) SetShipmentStatus(v ShipmentLineStatus)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *MerchantSingleOrderReturnLineResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *MerchantSingleOrderReturnLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantSingleOrderReturnLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantSingleOrderReturnLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetExtraData

`func (o *MerchantSingleOrderReturnLineResponse) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantSingleOrderReturnLineResponse) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantSingleOrderReturnLineResponse) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantSingleOrderReturnLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantSingleOrderReturnLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantSingleOrderReturnLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


