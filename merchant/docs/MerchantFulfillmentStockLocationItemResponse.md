# MerchantFulfillmentStockLocationItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ChannelEngine id of the stock location. | [optional] 
**Name** | Pointer to **NullableString** | The ChannelEngine name of the stock location. | [optional] 
**ReservedQuantity** | Pointer to **int32** | Total quantity reserved for orders. | [optional] 
**AvailableQuantity** | Pointer to **int32** | The quantity that is available for fulfillment orders. | [optional] 
**AllocatedQuantity** | Pointer to **int32** | Quantity pending customer order | [optional] 
**InTransitQuantity** | Pointer to **int32** | Quantity in transit / &#39;transshipment&#39; (Amazon) | [optional] 
**FulfillmentCenterProcessingQuantity** | Pointer to **int32** | Quantity that is in processing at the fulfillment warehouse (center) | [optional] 
**DefectiveQuantity** | Pointer to **int32** | The number of units in defective disposition. | [optional] 
**ExpiredQuantity** | Pointer to **int32** | The number of units in expired disposition. | [optional] 
**WarehouseDamagedQuantity** | Pointer to **int32** | The number of units in warehouse damaged disposition. | [optional] 
**CustomerDamagedQuantity** | Pointer to **int32** | The number of units in customer damaged disposition. | [optional] 
**CarrierDamagedQuantity** | Pointer to **int32** | The number of units in carrier damaged disposition. | [optional] 
**PendingPickupQuantity** | Pointer to **int32** | The number of units in pending pickup disposition. | [optional] 
**InboundQuantity** | Pointer to **int32** | Total quantity that is inbound (in inbound [aka fulfillment] shipment from the seller to the fulfillment warehouse) | [optional] 
**ReturnQuantity** | Pointer to **int32** | Total quantity in on going returns | [optional] 
**ResearchingQuantity** | Pointer to **int32** | Quantity that is being researched | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp of the last stock update for the stock location. | [optional] 

## Methods

### NewMerchantFulfillmentStockLocationItemResponse

`func NewMerchantFulfillmentStockLocationItemResponse() *MerchantFulfillmentStockLocationItemResponse`

NewMerchantFulfillmentStockLocationItemResponse instantiates a new MerchantFulfillmentStockLocationItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantFulfillmentStockLocationItemResponseWithDefaults

`func NewMerchantFulfillmentStockLocationItemResponseWithDefaults() *MerchantFulfillmentStockLocationItemResponse`

NewMerchantFulfillmentStockLocationItemResponseWithDefaults instantiates a new MerchantFulfillmentStockLocationItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantFulfillmentStockLocationItemResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantFulfillmentStockLocationItemResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantFulfillmentStockLocationItemResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MerchantFulfillmentStockLocationItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantFulfillmentStockLocationItemResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantFulfillmentStockLocationItemResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantFulfillmentStockLocationItemResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantFulfillmentStockLocationItemResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReservedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetReservedQuantity() int32`

GetReservedQuantity returns the ReservedQuantity field if non-nil, zero value otherwise.

### GetReservedQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetReservedQuantityOk() (*int32, bool)`

GetReservedQuantityOk returns a tuple with the ReservedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetReservedQuantity(v int32)`

SetReservedQuantity sets ReservedQuantity field to given value.

### HasReservedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasReservedQuantity() bool`

HasReservedQuantity returns a boolean if a field has been set.

### GetAvailableQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetAvailableQuantity() int32`

GetAvailableQuantity returns the AvailableQuantity field if non-nil, zero value otherwise.

### GetAvailableQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetAvailableQuantityOk() (*int32, bool)`

GetAvailableQuantityOk returns a tuple with the AvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetAvailableQuantity(v int32)`

SetAvailableQuantity sets AvailableQuantity field to given value.

### HasAvailableQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasAvailableQuantity() bool`

HasAvailableQuantity returns a boolean if a field has been set.

### GetAllocatedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetAllocatedQuantity() int32`

GetAllocatedQuantity returns the AllocatedQuantity field if non-nil, zero value otherwise.

### GetAllocatedQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetAllocatedQuantityOk() (*int32, bool)`

GetAllocatedQuantityOk returns a tuple with the AllocatedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetAllocatedQuantity(v int32)`

SetAllocatedQuantity sets AllocatedQuantity field to given value.

### HasAllocatedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasAllocatedQuantity() bool`

HasAllocatedQuantity returns a boolean if a field has been set.

### GetInTransitQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetInTransitQuantity() int32`

GetInTransitQuantity returns the InTransitQuantity field if non-nil, zero value otherwise.

### GetInTransitQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetInTransitQuantityOk() (*int32, bool)`

GetInTransitQuantityOk returns a tuple with the InTransitQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTransitQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetInTransitQuantity(v int32)`

SetInTransitQuantity sets InTransitQuantity field to given value.

### HasInTransitQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasInTransitQuantity() bool`

HasInTransitQuantity returns a boolean if a field has been set.

### GetFulfillmentCenterProcessingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetFulfillmentCenterProcessingQuantity() int32`

GetFulfillmentCenterProcessingQuantity returns the FulfillmentCenterProcessingQuantity field if non-nil, zero value otherwise.

### GetFulfillmentCenterProcessingQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetFulfillmentCenterProcessingQuantityOk() (*int32, bool)`

GetFulfillmentCenterProcessingQuantityOk returns a tuple with the FulfillmentCenterProcessingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterProcessingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetFulfillmentCenterProcessingQuantity(v int32)`

SetFulfillmentCenterProcessingQuantity sets FulfillmentCenterProcessingQuantity field to given value.

### HasFulfillmentCenterProcessingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasFulfillmentCenterProcessingQuantity() bool`

HasFulfillmentCenterProcessingQuantity returns a boolean if a field has been set.

### GetDefectiveQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetDefectiveQuantity() int32`

GetDefectiveQuantity returns the DefectiveQuantity field if non-nil, zero value otherwise.

### GetDefectiveQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetDefectiveQuantityOk() (*int32, bool)`

GetDefectiveQuantityOk returns a tuple with the DefectiveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectiveQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetDefectiveQuantity(v int32)`

SetDefectiveQuantity sets DefectiveQuantity field to given value.

### HasDefectiveQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasDefectiveQuantity() bool`

HasDefectiveQuantity returns a boolean if a field has been set.

### GetExpiredQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetExpiredQuantity() int32`

GetExpiredQuantity returns the ExpiredQuantity field if non-nil, zero value otherwise.

### GetExpiredQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetExpiredQuantityOk() (*int32, bool)`

GetExpiredQuantityOk returns a tuple with the ExpiredQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetExpiredQuantity(v int32)`

SetExpiredQuantity sets ExpiredQuantity field to given value.

### HasExpiredQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasExpiredQuantity() bool`

HasExpiredQuantity returns a boolean if a field has been set.

### GetWarehouseDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetWarehouseDamagedQuantity() int32`

GetWarehouseDamagedQuantity returns the WarehouseDamagedQuantity field if non-nil, zero value otherwise.

### GetWarehouseDamagedQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetWarehouseDamagedQuantityOk() (*int32, bool)`

GetWarehouseDamagedQuantityOk returns a tuple with the WarehouseDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetWarehouseDamagedQuantity(v int32)`

SetWarehouseDamagedQuantity sets WarehouseDamagedQuantity field to given value.

### HasWarehouseDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasWarehouseDamagedQuantity() bool`

HasWarehouseDamagedQuantity returns a boolean if a field has been set.

### GetCustomerDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetCustomerDamagedQuantity() int32`

GetCustomerDamagedQuantity returns the CustomerDamagedQuantity field if non-nil, zero value otherwise.

### GetCustomerDamagedQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetCustomerDamagedQuantityOk() (*int32, bool)`

GetCustomerDamagedQuantityOk returns a tuple with the CustomerDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetCustomerDamagedQuantity(v int32)`

SetCustomerDamagedQuantity sets CustomerDamagedQuantity field to given value.

### HasCustomerDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasCustomerDamagedQuantity() bool`

HasCustomerDamagedQuantity returns a boolean if a field has been set.

### GetCarrierDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetCarrierDamagedQuantity() int32`

GetCarrierDamagedQuantity returns the CarrierDamagedQuantity field if non-nil, zero value otherwise.

### GetCarrierDamagedQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetCarrierDamagedQuantityOk() (*int32, bool)`

GetCarrierDamagedQuantityOk returns a tuple with the CarrierDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetCarrierDamagedQuantity(v int32)`

SetCarrierDamagedQuantity sets CarrierDamagedQuantity field to given value.

### HasCarrierDamagedQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasCarrierDamagedQuantity() bool`

HasCarrierDamagedQuantity returns a boolean if a field has been set.

### GetPendingPickupQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetPendingPickupQuantity() int32`

GetPendingPickupQuantity returns the PendingPickupQuantity field if non-nil, zero value otherwise.

### GetPendingPickupQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetPendingPickupQuantityOk() (*int32, bool)`

GetPendingPickupQuantityOk returns a tuple with the PendingPickupQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPickupQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetPendingPickupQuantity(v int32)`

SetPendingPickupQuantity sets PendingPickupQuantity field to given value.

### HasPendingPickupQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasPendingPickupQuantity() bool`

HasPendingPickupQuantity returns a boolean if a field has been set.

### GetInboundQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetInboundQuantity() int32`

GetInboundQuantity returns the InboundQuantity field if non-nil, zero value otherwise.

### GetInboundQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetInboundQuantityOk() (*int32, bool)`

GetInboundQuantityOk returns a tuple with the InboundQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetInboundQuantity(v int32)`

SetInboundQuantity sets InboundQuantity field to given value.

### HasInboundQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasInboundQuantity() bool`

HasInboundQuantity returns a boolean if a field has been set.

### GetReturnQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetReturnQuantity() int32`

GetReturnQuantity returns the ReturnQuantity field if non-nil, zero value otherwise.

### GetReturnQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetReturnQuantityOk() (*int32, bool)`

GetReturnQuantityOk returns a tuple with the ReturnQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetReturnQuantity(v int32)`

SetReturnQuantity sets ReturnQuantity field to given value.

### HasReturnQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasReturnQuantity() bool`

HasReturnQuantity returns a boolean if a field has been set.

### GetResearchingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) GetResearchingQuantity() int32`

GetResearchingQuantity returns the ResearchingQuantity field if non-nil, zero value otherwise.

### GetResearchingQuantityOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetResearchingQuantityOk() (*int32, bool)`

GetResearchingQuantityOk returns a tuple with the ResearchingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) SetResearchingQuantity(v int32)`

SetResearchingQuantity sets ResearchingQuantity field to given value.

### HasResearchingQuantity

`func (o *MerchantFulfillmentStockLocationItemResponse) HasResearchingQuantity() bool`

HasResearchingQuantity returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantFulfillmentStockLocationItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantFulfillmentStockLocationItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantFulfillmentStockLocationItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantFulfillmentStockLocationItemResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


