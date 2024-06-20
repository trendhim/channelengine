# IPurchaseOrderShipmentByFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | ChannelEngine identifier of the shipment | [optional] 
**MerchantShipmentNo** | Pointer to **NullableString** | The number the merchant uses to id this PO shipment | [optional] 
**ChannelShipmentNo** | Pointer to **NullableString** | The number the channel uses to id this PO shipment | [optional] 
**ShipmentType** | Pointer to [**ShipmentType**](ShipmentType.md) |  | [optional] 
**ShippedDate** | Pointer to **time.Time** | When the shipment was shipped | [optional] 
**EstimatedDeliveryDate** | Pointer to **time.Time** | Estimated delivery time in the channel&#39;s warehouse | [optional] 
**CarrierName** | Pointer to **NullableString** | Name of the carrier | [optional] 
**CarrierShipmentNo** | Pointer to **NullableString** | The number the carrier uses to id this PO shipment | [optional] 
**BillOfLadingNumber** | Pointer to **NullableString** | Bill of Lading number (not unique for a shipment) | [optional] 
**ShipmentWeightUnitOfMeasure** | Pointer to [**WeightUnitOfMeasure**](WeightUnitOfMeasure.md) |  | [optional] 
**ShipmentWeight** | Pointer to **NullableFloat32** | The shipment&#39;s weight | [optional] 
**ShipmentVolumeUnitOfMeasure** | Pointer to [**VolumeUnitOfMeasure**](VolumeUnitOfMeasure.md) |  | [optional] 
**ShipmentVolume** | Pointer to **NullableFloat32** | The shipment&#39;s volume | [optional] 
**LastMerchantUpdatedAt** | Pointer to **time.Time** | The last time the shipment was updated by the merchant | [optional] 
**LastExportedAt** | Pointer to **NullableTime** | The last time the shipment was exported to the channel | [optional] 
**LastExportStatus** | Pointer to [**PurchaseOrderRelatedItemExportStatus**](PurchaseOrderRelatedItemExportStatus.md) |  | [optional] 
**Lines** | Pointer to [**[]IPurchaseOrderShipmentLineByFilter**](IPurchaseOrderShipmentLineByFilter.md) | The products in this PO shipment | [optional] 

## Methods

### NewIPurchaseOrderShipmentByFilter

`func NewIPurchaseOrderShipmentByFilter() *IPurchaseOrderShipmentByFilter`

NewIPurchaseOrderShipmentByFilter instantiates a new IPurchaseOrderShipmentByFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPurchaseOrderShipmentByFilterWithDefaults

`func NewIPurchaseOrderShipmentByFilterWithDefaults() *IPurchaseOrderShipmentByFilter`

NewIPurchaseOrderShipmentByFilterWithDefaults instantiates a new IPurchaseOrderShipmentByFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPurchaseOrderShipmentByFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPurchaseOrderShipmentByFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPurchaseOrderShipmentByFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPurchaseOrderShipmentByFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IPurchaseOrderShipmentByFilter) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IPurchaseOrderShipmentByFilter) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMerchantShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *IPurchaseOrderShipmentByFilter) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.

### HasMerchantShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) HasMerchantShipmentNo() bool`

HasMerchantShipmentNo returns a boolean if a field has been set.

### SetMerchantShipmentNoNil

`func (o *IPurchaseOrderShipmentByFilter) SetMerchantShipmentNoNil(b bool)`

 SetMerchantShipmentNoNil sets the value for MerchantShipmentNo to be an explicit nil

### UnsetMerchantShipmentNo
`func (o *IPurchaseOrderShipmentByFilter) UnsetMerchantShipmentNo()`

UnsetMerchantShipmentNo ensures that no value is present for MerchantShipmentNo, not even an explicit nil
### GetChannelShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) GetChannelShipmentNo() string`

GetChannelShipmentNo returns the ChannelShipmentNo field if non-nil, zero value otherwise.

### GetChannelShipmentNoOk

`func (o *IPurchaseOrderShipmentByFilter) GetChannelShipmentNoOk() (*string, bool)`

GetChannelShipmentNoOk returns a tuple with the ChannelShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) SetChannelShipmentNo(v string)`

SetChannelShipmentNo sets ChannelShipmentNo field to given value.

### HasChannelShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) HasChannelShipmentNo() bool`

HasChannelShipmentNo returns a boolean if a field has been set.

### SetChannelShipmentNoNil

`func (o *IPurchaseOrderShipmentByFilter) SetChannelShipmentNoNil(b bool)`

 SetChannelShipmentNoNil sets the value for ChannelShipmentNo to be an explicit nil

### UnsetChannelShipmentNo
`func (o *IPurchaseOrderShipmentByFilter) UnsetChannelShipmentNo()`

UnsetChannelShipmentNo ensures that no value is present for ChannelShipmentNo, not even an explicit nil
### GetShipmentType

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentType() ShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentTypeOk() (*ShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentType(v ShipmentType)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *IPurchaseOrderShipmentByFilter) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetShippedDate

`func (o *IPurchaseOrderShipmentByFilter) GetShippedDate() time.Time`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *IPurchaseOrderShipmentByFilter) GetShippedDateOk() (*time.Time, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *IPurchaseOrderShipmentByFilter) SetShippedDate(v time.Time)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *IPurchaseOrderShipmentByFilter) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *IPurchaseOrderShipmentByFilter) GetEstimatedDeliveryDate() time.Time`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *IPurchaseOrderShipmentByFilter) GetEstimatedDeliveryDateOk() (*time.Time, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *IPurchaseOrderShipmentByFilter) SetEstimatedDeliveryDate(v time.Time)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *IPurchaseOrderShipmentByFilter) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetCarrierName

`func (o *IPurchaseOrderShipmentByFilter) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *IPurchaseOrderShipmentByFilter) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *IPurchaseOrderShipmentByFilter) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *IPurchaseOrderShipmentByFilter) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### SetCarrierNameNil

`func (o *IPurchaseOrderShipmentByFilter) SetCarrierNameNil(b bool)`

 SetCarrierNameNil sets the value for CarrierName to be an explicit nil

### UnsetCarrierName
`func (o *IPurchaseOrderShipmentByFilter) UnsetCarrierName()`

UnsetCarrierName ensures that no value is present for CarrierName, not even an explicit nil
### GetCarrierShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) GetCarrierShipmentNo() string`

GetCarrierShipmentNo returns the CarrierShipmentNo field if non-nil, zero value otherwise.

### GetCarrierShipmentNoOk

`func (o *IPurchaseOrderShipmentByFilter) GetCarrierShipmentNoOk() (*string, bool)`

GetCarrierShipmentNoOk returns a tuple with the CarrierShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) SetCarrierShipmentNo(v string)`

SetCarrierShipmentNo sets CarrierShipmentNo field to given value.

### HasCarrierShipmentNo

`func (o *IPurchaseOrderShipmentByFilter) HasCarrierShipmentNo() bool`

HasCarrierShipmentNo returns a boolean if a field has been set.

### SetCarrierShipmentNoNil

`func (o *IPurchaseOrderShipmentByFilter) SetCarrierShipmentNoNil(b bool)`

 SetCarrierShipmentNoNil sets the value for CarrierShipmentNo to be an explicit nil

### UnsetCarrierShipmentNo
`func (o *IPurchaseOrderShipmentByFilter) UnsetCarrierShipmentNo()`

UnsetCarrierShipmentNo ensures that no value is present for CarrierShipmentNo, not even an explicit nil
### GetBillOfLadingNumber

`func (o *IPurchaseOrderShipmentByFilter) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *IPurchaseOrderShipmentByFilter) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *IPurchaseOrderShipmentByFilter) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *IPurchaseOrderShipmentByFilter) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *IPurchaseOrderShipmentByFilter) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *IPurchaseOrderShipmentByFilter) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetShipmentWeightUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentWeightUnitOfMeasure() WeightUnitOfMeasure`

GetShipmentWeightUnitOfMeasure returns the ShipmentWeightUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentWeightUnitOfMeasureOk

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentWeightUnitOfMeasureOk() (*WeightUnitOfMeasure, bool)`

GetShipmentWeightUnitOfMeasureOk returns a tuple with the ShipmentWeightUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeightUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentWeightUnitOfMeasure(v WeightUnitOfMeasure)`

SetShipmentWeightUnitOfMeasure sets ShipmentWeightUnitOfMeasure field to given value.

### HasShipmentWeightUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) HasShipmentWeightUnitOfMeasure() bool`

HasShipmentWeightUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentWeight

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentWeight() float32`

GetShipmentWeight returns the ShipmentWeight field if non-nil, zero value otherwise.

### GetShipmentWeightOk

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentWeightOk() (*float32, bool)`

GetShipmentWeightOk returns a tuple with the ShipmentWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeight

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentWeight(v float32)`

SetShipmentWeight sets ShipmentWeight field to given value.

### HasShipmentWeight

`func (o *IPurchaseOrderShipmentByFilter) HasShipmentWeight() bool`

HasShipmentWeight returns a boolean if a field has been set.

### SetShipmentWeightNil

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentWeightNil(b bool)`

 SetShipmentWeightNil sets the value for ShipmentWeight to be an explicit nil

### UnsetShipmentWeight
`func (o *IPurchaseOrderShipmentByFilter) UnsetShipmentWeight()`

UnsetShipmentWeight ensures that no value is present for ShipmentWeight, not even an explicit nil
### GetShipmentVolumeUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentVolumeUnitOfMeasure() VolumeUnitOfMeasure`

GetShipmentVolumeUnitOfMeasure returns the ShipmentVolumeUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentVolumeUnitOfMeasureOk

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentVolumeUnitOfMeasureOk() (*VolumeUnitOfMeasure, bool)`

GetShipmentVolumeUnitOfMeasureOk returns a tuple with the ShipmentVolumeUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolumeUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentVolumeUnitOfMeasure(v VolumeUnitOfMeasure)`

SetShipmentVolumeUnitOfMeasure sets ShipmentVolumeUnitOfMeasure field to given value.

### HasShipmentVolumeUnitOfMeasure

`func (o *IPurchaseOrderShipmentByFilter) HasShipmentVolumeUnitOfMeasure() bool`

HasShipmentVolumeUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentVolume

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentVolume() float32`

GetShipmentVolume returns the ShipmentVolume field if non-nil, zero value otherwise.

### GetShipmentVolumeOk

`func (o *IPurchaseOrderShipmentByFilter) GetShipmentVolumeOk() (*float32, bool)`

GetShipmentVolumeOk returns a tuple with the ShipmentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolume

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentVolume(v float32)`

SetShipmentVolume sets ShipmentVolume field to given value.

### HasShipmentVolume

`func (o *IPurchaseOrderShipmentByFilter) HasShipmentVolume() bool`

HasShipmentVolume returns a boolean if a field has been set.

### SetShipmentVolumeNil

`func (o *IPurchaseOrderShipmentByFilter) SetShipmentVolumeNil(b bool)`

 SetShipmentVolumeNil sets the value for ShipmentVolume to be an explicit nil

### UnsetShipmentVolume
`func (o *IPurchaseOrderShipmentByFilter) UnsetShipmentVolume()`

UnsetShipmentVolume ensures that no value is present for ShipmentVolume, not even an explicit nil
### GetLastMerchantUpdatedAt

`func (o *IPurchaseOrderShipmentByFilter) GetLastMerchantUpdatedAt() time.Time`

GetLastMerchantUpdatedAt returns the LastMerchantUpdatedAt field if non-nil, zero value otherwise.

### GetLastMerchantUpdatedAtOk

`func (o *IPurchaseOrderShipmentByFilter) GetLastMerchantUpdatedAtOk() (*time.Time, bool)`

GetLastMerchantUpdatedAtOk returns a tuple with the LastMerchantUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMerchantUpdatedAt

`func (o *IPurchaseOrderShipmentByFilter) SetLastMerchantUpdatedAt(v time.Time)`

SetLastMerchantUpdatedAt sets LastMerchantUpdatedAt field to given value.

### HasLastMerchantUpdatedAt

`func (o *IPurchaseOrderShipmentByFilter) HasLastMerchantUpdatedAt() bool`

HasLastMerchantUpdatedAt returns a boolean if a field has been set.

### GetLastExportedAt

`func (o *IPurchaseOrderShipmentByFilter) GetLastExportedAt() time.Time`

GetLastExportedAt returns the LastExportedAt field if non-nil, zero value otherwise.

### GetLastExportedAtOk

`func (o *IPurchaseOrderShipmentByFilter) GetLastExportedAtOk() (*time.Time, bool)`

GetLastExportedAtOk returns a tuple with the LastExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExportedAt

`func (o *IPurchaseOrderShipmentByFilter) SetLastExportedAt(v time.Time)`

SetLastExportedAt sets LastExportedAt field to given value.

### HasLastExportedAt

`func (o *IPurchaseOrderShipmentByFilter) HasLastExportedAt() bool`

HasLastExportedAt returns a boolean if a field has been set.

### SetLastExportedAtNil

`func (o *IPurchaseOrderShipmentByFilter) SetLastExportedAtNil(b bool)`

 SetLastExportedAtNil sets the value for LastExportedAt to be an explicit nil

### UnsetLastExportedAt
`func (o *IPurchaseOrderShipmentByFilter) UnsetLastExportedAt()`

UnsetLastExportedAt ensures that no value is present for LastExportedAt, not even an explicit nil
### GetLastExportStatus

`func (o *IPurchaseOrderShipmentByFilter) GetLastExportStatus() PurchaseOrderRelatedItemExportStatus`

GetLastExportStatus returns the LastExportStatus field if non-nil, zero value otherwise.

### GetLastExportStatusOk

`func (o *IPurchaseOrderShipmentByFilter) GetLastExportStatusOk() (*PurchaseOrderRelatedItemExportStatus, bool)`

GetLastExportStatusOk returns a tuple with the LastExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExportStatus

`func (o *IPurchaseOrderShipmentByFilter) SetLastExportStatus(v PurchaseOrderRelatedItemExportStatus)`

SetLastExportStatus sets LastExportStatus field to given value.

### HasLastExportStatus

`func (o *IPurchaseOrderShipmentByFilter) HasLastExportStatus() bool`

HasLastExportStatus returns a boolean if a field has been set.

### GetLines

`func (o *IPurchaseOrderShipmentByFilter) GetLines() []IPurchaseOrderShipmentLineByFilter`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *IPurchaseOrderShipmentByFilter) GetLinesOk() (*[]IPurchaseOrderShipmentLineByFilter, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *IPurchaseOrderShipmentByFilter) SetLines(v []IPurchaseOrderShipmentLineByFilter)`

SetLines sets Lines field to given value.

### HasLines

`func (o *IPurchaseOrderShipmentByFilter) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *IPurchaseOrderShipmentByFilter) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *IPurchaseOrderShipmentByFilter) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


