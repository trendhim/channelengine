# UpdatePurchaseOrderShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantShipmentNo** | Pointer to **NullableString** | The number the merchant uses to identify this PO shipment | [optional] 
**ShipmentType** | Pointer to [**ShipmentType**](ShipmentType.md) |  | [optional] 
**ShippedDate** | Pointer to **time.Time** | When the shipment will be/was shipped | [optional] 
**EstimatedDeliveryDate** | Pointer to **time.Time** | Estimated delivery time in the channel&#39;s warehouse | [optional] 
**SellingPartyId** | Pointer to **NullableString** | The merchant&#39;s identifying &#39;selling party number&#39; at the channel | [optional] 
**ShipToPartyId** | Pointer to **NullableString** | The destination&#39;s &#39;ship to party&#39; number at the channel | [optional] 
**BillOfLadingNumber** | Pointer to **NullableString** | Bill Of Lading (BOL) number is the unique number assigned by the vendor. The BOL present in the Shipment Confirmation message ideally matches the paper BOL provided with the shipment, but that is no must. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field. | [optional] 
**ShipmentWeightUnitOfMeasure** | Pointer to [**WeightUnitOfMeasure**](WeightUnitOfMeasure.md) |  | [optional] 
**ShipmentWeight** | Pointer to **NullableFloat32** | The shipment&#39;s weight | [optional] 
**ShipmentVolumeUnitOfMeasure** | Pointer to [**VolumeUnitOfMeasure**](VolumeUnitOfMeasure.md) |  | [optional] 
**ShipmentVolume** | Pointer to **NullableFloat32** | The shipment&#39;s volume | [optional] 
**Lines** | Pointer to [**[]ChangePurchaseOrderShipmentLine**](ChangePurchaseOrderShipmentLine.md) | Shipment information for each shipped product | [optional] 

## Methods

### NewUpdatePurchaseOrderShipment

`func NewUpdatePurchaseOrderShipment() *UpdatePurchaseOrderShipment`

NewUpdatePurchaseOrderShipment instantiates a new UpdatePurchaseOrderShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePurchaseOrderShipmentWithDefaults

`func NewUpdatePurchaseOrderShipmentWithDefaults() *UpdatePurchaseOrderShipment`

NewUpdatePurchaseOrderShipmentWithDefaults instantiates a new UpdatePurchaseOrderShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantShipmentNo

`func (o *UpdatePurchaseOrderShipment) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *UpdatePurchaseOrderShipment) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *UpdatePurchaseOrderShipment) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.

### HasMerchantShipmentNo

`func (o *UpdatePurchaseOrderShipment) HasMerchantShipmentNo() bool`

HasMerchantShipmentNo returns a boolean if a field has been set.

### SetMerchantShipmentNoNil

`func (o *UpdatePurchaseOrderShipment) SetMerchantShipmentNoNil(b bool)`

 SetMerchantShipmentNoNil sets the value for MerchantShipmentNo to be an explicit nil

### UnsetMerchantShipmentNo
`func (o *UpdatePurchaseOrderShipment) UnsetMerchantShipmentNo()`

UnsetMerchantShipmentNo ensures that no value is present for MerchantShipmentNo, not even an explicit nil
### GetShipmentType

`func (o *UpdatePurchaseOrderShipment) GetShipmentType() ShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *UpdatePurchaseOrderShipment) GetShipmentTypeOk() (*ShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *UpdatePurchaseOrderShipment) SetShipmentType(v ShipmentType)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *UpdatePurchaseOrderShipment) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetShippedDate

`func (o *UpdatePurchaseOrderShipment) GetShippedDate() time.Time`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *UpdatePurchaseOrderShipment) GetShippedDateOk() (*time.Time, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *UpdatePurchaseOrderShipment) SetShippedDate(v time.Time)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *UpdatePurchaseOrderShipment) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *UpdatePurchaseOrderShipment) GetEstimatedDeliveryDate() time.Time`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *UpdatePurchaseOrderShipment) GetEstimatedDeliveryDateOk() (*time.Time, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *UpdatePurchaseOrderShipment) SetEstimatedDeliveryDate(v time.Time)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *UpdatePurchaseOrderShipment) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetSellingPartyId

`func (o *UpdatePurchaseOrderShipment) GetSellingPartyId() string`

GetSellingPartyId returns the SellingPartyId field if non-nil, zero value otherwise.

### GetSellingPartyIdOk

`func (o *UpdatePurchaseOrderShipment) GetSellingPartyIdOk() (*string, bool)`

GetSellingPartyIdOk returns a tuple with the SellingPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPartyId

`func (o *UpdatePurchaseOrderShipment) SetSellingPartyId(v string)`

SetSellingPartyId sets SellingPartyId field to given value.

### HasSellingPartyId

`func (o *UpdatePurchaseOrderShipment) HasSellingPartyId() bool`

HasSellingPartyId returns a boolean if a field has been set.

### SetSellingPartyIdNil

`func (o *UpdatePurchaseOrderShipment) SetSellingPartyIdNil(b bool)`

 SetSellingPartyIdNil sets the value for SellingPartyId to be an explicit nil

### UnsetSellingPartyId
`func (o *UpdatePurchaseOrderShipment) UnsetSellingPartyId()`

UnsetSellingPartyId ensures that no value is present for SellingPartyId, not even an explicit nil
### GetShipToPartyId

`func (o *UpdatePurchaseOrderShipment) GetShipToPartyId() string`

GetShipToPartyId returns the ShipToPartyId field if non-nil, zero value otherwise.

### GetShipToPartyIdOk

`func (o *UpdatePurchaseOrderShipment) GetShipToPartyIdOk() (*string, bool)`

GetShipToPartyIdOk returns a tuple with the ShipToPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToPartyId

`func (o *UpdatePurchaseOrderShipment) SetShipToPartyId(v string)`

SetShipToPartyId sets ShipToPartyId field to given value.

### HasShipToPartyId

`func (o *UpdatePurchaseOrderShipment) HasShipToPartyId() bool`

HasShipToPartyId returns a boolean if a field has been set.

### SetShipToPartyIdNil

`func (o *UpdatePurchaseOrderShipment) SetShipToPartyIdNil(b bool)`

 SetShipToPartyIdNil sets the value for ShipToPartyId to be an explicit nil

### UnsetShipToPartyId
`func (o *UpdatePurchaseOrderShipment) UnsetShipToPartyId()`

UnsetShipToPartyId ensures that no value is present for ShipToPartyId, not even an explicit nil
### GetBillOfLadingNumber

`func (o *UpdatePurchaseOrderShipment) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *UpdatePurchaseOrderShipment) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *UpdatePurchaseOrderShipment) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *UpdatePurchaseOrderShipment) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *UpdatePurchaseOrderShipment) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *UpdatePurchaseOrderShipment) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetShipmentWeightUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasure() WeightUnitOfMeasure`

GetShipmentWeightUnitOfMeasure returns the ShipmentWeightUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentWeightUnitOfMeasureOk

`func (o *UpdatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasureOk() (*WeightUnitOfMeasure, bool)`

GetShipmentWeightUnitOfMeasureOk returns a tuple with the ShipmentWeightUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeightUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) SetShipmentWeightUnitOfMeasure(v WeightUnitOfMeasure)`

SetShipmentWeightUnitOfMeasure sets ShipmentWeightUnitOfMeasure field to given value.

### HasShipmentWeightUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) HasShipmentWeightUnitOfMeasure() bool`

HasShipmentWeightUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentWeight

`func (o *UpdatePurchaseOrderShipment) GetShipmentWeight() float32`

GetShipmentWeight returns the ShipmentWeight field if non-nil, zero value otherwise.

### GetShipmentWeightOk

`func (o *UpdatePurchaseOrderShipment) GetShipmentWeightOk() (*float32, bool)`

GetShipmentWeightOk returns a tuple with the ShipmentWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeight

`func (o *UpdatePurchaseOrderShipment) SetShipmentWeight(v float32)`

SetShipmentWeight sets ShipmentWeight field to given value.

### HasShipmentWeight

`func (o *UpdatePurchaseOrderShipment) HasShipmentWeight() bool`

HasShipmentWeight returns a boolean if a field has been set.

### SetShipmentWeightNil

`func (o *UpdatePurchaseOrderShipment) SetShipmentWeightNil(b bool)`

 SetShipmentWeightNil sets the value for ShipmentWeight to be an explicit nil

### UnsetShipmentWeight
`func (o *UpdatePurchaseOrderShipment) UnsetShipmentWeight()`

UnsetShipmentWeight ensures that no value is present for ShipmentWeight, not even an explicit nil
### GetShipmentVolumeUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasure() VolumeUnitOfMeasure`

GetShipmentVolumeUnitOfMeasure returns the ShipmentVolumeUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentVolumeUnitOfMeasureOk

`func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasureOk() (*VolumeUnitOfMeasure, bool)`

GetShipmentVolumeUnitOfMeasureOk returns a tuple with the ShipmentVolumeUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolumeUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) SetShipmentVolumeUnitOfMeasure(v VolumeUnitOfMeasure)`

SetShipmentVolumeUnitOfMeasure sets ShipmentVolumeUnitOfMeasure field to given value.

### HasShipmentVolumeUnitOfMeasure

`func (o *UpdatePurchaseOrderShipment) HasShipmentVolumeUnitOfMeasure() bool`

HasShipmentVolumeUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentVolume

`func (o *UpdatePurchaseOrderShipment) GetShipmentVolume() float32`

GetShipmentVolume returns the ShipmentVolume field if non-nil, zero value otherwise.

### GetShipmentVolumeOk

`func (o *UpdatePurchaseOrderShipment) GetShipmentVolumeOk() (*float32, bool)`

GetShipmentVolumeOk returns a tuple with the ShipmentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolume

`func (o *UpdatePurchaseOrderShipment) SetShipmentVolume(v float32)`

SetShipmentVolume sets ShipmentVolume field to given value.

### HasShipmentVolume

`func (o *UpdatePurchaseOrderShipment) HasShipmentVolume() bool`

HasShipmentVolume returns a boolean if a field has been set.

### SetShipmentVolumeNil

`func (o *UpdatePurchaseOrderShipment) SetShipmentVolumeNil(b bool)`

 SetShipmentVolumeNil sets the value for ShipmentVolume to be an explicit nil

### UnsetShipmentVolume
`func (o *UpdatePurchaseOrderShipment) UnsetShipmentVolume()`

UnsetShipmentVolume ensures that no value is present for ShipmentVolume, not even an explicit nil
### GetLines

`func (o *UpdatePurchaseOrderShipment) GetLines() []ChangePurchaseOrderShipmentLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UpdatePurchaseOrderShipment) GetLinesOk() (*[]ChangePurchaseOrderShipmentLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UpdatePurchaseOrderShipment) SetLines(v []ChangePurchaseOrderShipmentLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UpdatePurchaseOrderShipment) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *UpdatePurchaseOrderShipment) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *UpdatePurchaseOrderShipment) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


