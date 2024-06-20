# CreatePurchaseOrderShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierShipmentNo** | Pointer to **NullableString** | The field is also known as PRO number is a unique number assigned by the carrier.  It is used to identify and track the shipment that goes out for delivery.  This field is mandatory for US, CA, MX shipment confirmations of Amazon Vendor | [optional] 
**CarrierName** | Pointer to **NullableString** | Name of the carrier | [optional] 
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

### NewCreatePurchaseOrderShipment

`func NewCreatePurchaseOrderShipment() *CreatePurchaseOrderShipment`

NewCreatePurchaseOrderShipment instantiates a new CreatePurchaseOrderShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePurchaseOrderShipmentWithDefaults

`func NewCreatePurchaseOrderShipmentWithDefaults() *CreatePurchaseOrderShipment`

NewCreatePurchaseOrderShipmentWithDefaults instantiates a new CreatePurchaseOrderShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierShipmentNo

`func (o *CreatePurchaseOrderShipment) GetCarrierShipmentNo() string`

GetCarrierShipmentNo returns the CarrierShipmentNo field if non-nil, zero value otherwise.

### GetCarrierShipmentNoOk

`func (o *CreatePurchaseOrderShipment) GetCarrierShipmentNoOk() (*string, bool)`

GetCarrierShipmentNoOk returns a tuple with the CarrierShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierShipmentNo

`func (o *CreatePurchaseOrderShipment) SetCarrierShipmentNo(v string)`

SetCarrierShipmentNo sets CarrierShipmentNo field to given value.

### HasCarrierShipmentNo

`func (o *CreatePurchaseOrderShipment) HasCarrierShipmentNo() bool`

HasCarrierShipmentNo returns a boolean if a field has been set.

### SetCarrierShipmentNoNil

`func (o *CreatePurchaseOrderShipment) SetCarrierShipmentNoNil(b bool)`

 SetCarrierShipmentNoNil sets the value for CarrierShipmentNo to be an explicit nil

### UnsetCarrierShipmentNo
`func (o *CreatePurchaseOrderShipment) UnsetCarrierShipmentNo()`

UnsetCarrierShipmentNo ensures that no value is present for CarrierShipmentNo, not even an explicit nil
### GetCarrierName

`func (o *CreatePurchaseOrderShipment) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *CreatePurchaseOrderShipment) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *CreatePurchaseOrderShipment) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *CreatePurchaseOrderShipment) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### SetCarrierNameNil

`func (o *CreatePurchaseOrderShipment) SetCarrierNameNil(b bool)`

 SetCarrierNameNil sets the value for CarrierName to be an explicit nil

### UnsetCarrierName
`func (o *CreatePurchaseOrderShipment) UnsetCarrierName()`

UnsetCarrierName ensures that no value is present for CarrierName, not even an explicit nil
### GetMerchantShipmentNo

`func (o *CreatePurchaseOrderShipment) GetMerchantShipmentNo() string`

GetMerchantShipmentNo returns the MerchantShipmentNo field if non-nil, zero value otherwise.

### GetMerchantShipmentNoOk

`func (o *CreatePurchaseOrderShipment) GetMerchantShipmentNoOk() (*string, bool)`

GetMerchantShipmentNoOk returns a tuple with the MerchantShipmentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShipmentNo

`func (o *CreatePurchaseOrderShipment) SetMerchantShipmentNo(v string)`

SetMerchantShipmentNo sets MerchantShipmentNo field to given value.

### HasMerchantShipmentNo

`func (o *CreatePurchaseOrderShipment) HasMerchantShipmentNo() bool`

HasMerchantShipmentNo returns a boolean if a field has been set.

### SetMerchantShipmentNoNil

`func (o *CreatePurchaseOrderShipment) SetMerchantShipmentNoNil(b bool)`

 SetMerchantShipmentNoNil sets the value for MerchantShipmentNo to be an explicit nil

### UnsetMerchantShipmentNo
`func (o *CreatePurchaseOrderShipment) UnsetMerchantShipmentNo()`

UnsetMerchantShipmentNo ensures that no value is present for MerchantShipmentNo, not even an explicit nil
### GetShipmentType

`func (o *CreatePurchaseOrderShipment) GetShipmentType() ShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *CreatePurchaseOrderShipment) GetShipmentTypeOk() (*ShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *CreatePurchaseOrderShipment) SetShipmentType(v ShipmentType)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *CreatePurchaseOrderShipment) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetShippedDate

`func (o *CreatePurchaseOrderShipment) GetShippedDate() time.Time`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *CreatePurchaseOrderShipment) GetShippedDateOk() (*time.Time, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *CreatePurchaseOrderShipment) SetShippedDate(v time.Time)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *CreatePurchaseOrderShipment) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *CreatePurchaseOrderShipment) GetEstimatedDeliveryDate() time.Time`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *CreatePurchaseOrderShipment) GetEstimatedDeliveryDateOk() (*time.Time, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *CreatePurchaseOrderShipment) SetEstimatedDeliveryDate(v time.Time)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *CreatePurchaseOrderShipment) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetSellingPartyId

`func (o *CreatePurchaseOrderShipment) GetSellingPartyId() string`

GetSellingPartyId returns the SellingPartyId field if non-nil, zero value otherwise.

### GetSellingPartyIdOk

`func (o *CreatePurchaseOrderShipment) GetSellingPartyIdOk() (*string, bool)`

GetSellingPartyIdOk returns a tuple with the SellingPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPartyId

`func (o *CreatePurchaseOrderShipment) SetSellingPartyId(v string)`

SetSellingPartyId sets SellingPartyId field to given value.

### HasSellingPartyId

`func (o *CreatePurchaseOrderShipment) HasSellingPartyId() bool`

HasSellingPartyId returns a boolean if a field has been set.

### SetSellingPartyIdNil

`func (o *CreatePurchaseOrderShipment) SetSellingPartyIdNil(b bool)`

 SetSellingPartyIdNil sets the value for SellingPartyId to be an explicit nil

### UnsetSellingPartyId
`func (o *CreatePurchaseOrderShipment) UnsetSellingPartyId()`

UnsetSellingPartyId ensures that no value is present for SellingPartyId, not even an explicit nil
### GetShipToPartyId

`func (o *CreatePurchaseOrderShipment) GetShipToPartyId() string`

GetShipToPartyId returns the ShipToPartyId field if non-nil, zero value otherwise.

### GetShipToPartyIdOk

`func (o *CreatePurchaseOrderShipment) GetShipToPartyIdOk() (*string, bool)`

GetShipToPartyIdOk returns a tuple with the ShipToPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToPartyId

`func (o *CreatePurchaseOrderShipment) SetShipToPartyId(v string)`

SetShipToPartyId sets ShipToPartyId field to given value.

### HasShipToPartyId

`func (o *CreatePurchaseOrderShipment) HasShipToPartyId() bool`

HasShipToPartyId returns a boolean if a field has been set.

### SetShipToPartyIdNil

`func (o *CreatePurchaseOrderShipment) SetShipToPartyIdNil(b bool)`

 SetShipToPartyIdNil sets the value for ShipToPartyId to be an explicit nil

### UnsetShipToPartyId
`func (o *CreatePurchaseOrderShipment) UnsetShipToPartyId()`

UnsetShipToPartyId ensures that no value is present for ShipToPartyId, not even an explicit nil
### GetBillOfLadingNumber

`func (o *CreatePurchaseOrderShipment) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *CreatePurchaseOrderShipment) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *CreatePurchaseOrderShipment) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *CreatePurchaseOrderShipment) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *CreatePurchaseOrderShipment) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *CreatePurchaseOrderShipment) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetShipmentWeightUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasure() WeightUnitOfMeasure`

GetShipmentWeightUnitOfMeasure returns the ShipmentWeightUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentWeightUnitOfMeasureOk

`func (o *CreatePurchaseOrderShipment) GetShipmentWeightUnitOfMeasureOk() (*WeightUnitOfMeasure, bool)`

GetShipmentWeightUnitOfMeasureOk returns a tuple with the ShipmentWeightUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeightUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) SetShipmentWeightUnitOfMeasure(v WeightUnitOfMeasure)`

SetShipmentWeightUnitOfMeasure sets ShipmentWeightUnitOfMeasure field to given value.

### HasShipmentWeightUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) HasShipmentWeightUnitOfMeasure() bool`

HasShipmentWeightUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentWeight

`func (o *CreatePurchaseOrderShipment) GetShipmentWeight() float32`

GetShipmentWeight returns the ShipmentWeight field if non-nil, zero value otherwise.

### GetShipmentWeightOk

`func (o *CreatePurchaseOrderShipment) GetShipmentWeightOk() (*float32, bool)`

GetShipmentWeightOk returns a tuple with the ShipmentWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentWeight

`func (o *CreatePurchaseOrderShipment) SetShipmentWeight(v float32)`

SetShipmentWeight sets ShipmentWeight field to given value.

### HasShipmentWeight

`func (o *CreatePurchaseOrderShipment) HasShipmentWeight() bool`

HasShipmentWeight returns a boolean if a field has been set.

### SetShipmentWeightNil

`func (o *CreatePurchaseOrderShipment) SetShipmentWeightNil(b bool)`

 SetShipmentWeightNil sets the value for ShipmentWeight to be an explicit nil

### UnsetShipmentWeight
`func (o *CreatePurchaseOrderShipment) UnsetShipmentWeight()`

UnsetShipmentWeight ensures that no value is present for ShipmentWeight, not even an explicit nil
### GetShipmentVolumeUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasure() VolumeUnitOfMeasure`

GetShipmentVolumeUnitOfMeasure returns the ShipmentVolumeUnitOfMeasure field if non-nil, zero value otherwise.

### GetShipmentVolumeUnitOfMeasureOk

`func (o *CreatePurchaseOrderShipment) GetShipmentVolumeUnitOfMeasureOk() (*VolumeUnitOfMeasure, bool)`

GetShipmentVolumeUnitOfMeasureOk returns a tuple with the ShipmentVolumeUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolumeUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) SetShipmentVolumeUnitOfMeasure(v VolumeUnitOfMeasure)`

SetShipmentVolumeUnitOfMeasure sets ShipmentVolumeUnitOfMeasure field to given value.

### HasShipmentVolumeUnitOfMeasure

`func (o *CreatePurchaseOrderShipment) HasShipmentVolumeUnitOfMeasure() bool`

HasShipmentVolumeUnitOfMeasure returns a boolean if a field has been set.

### GetShipmentVolume

`func (o *CreatePurchaseOrderShipment) GetShipmentVolume() float32`

GetShipmentVolume returns the ShipmentVolume field if non-nil, zero value otherwise.

### GetShipmentVolumeOk

`func (o *CreatePurchaseOrderShipment) GetShipmentVolumeOk() (*float32, bool)`

GetShipmentVolumeOk returns a tuple with the ShipmentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentVolume

`func (o *CreatePurchaseOrderShipment) SetShipmentVolume(v float32)`

SetShipmentVolume sets ShipmentVolume field to given value.

### HasShipmentVolume

`func (o *CreatePurchaseOrderShipment) HasShipmentVolume() bool`

HasShipmentVolume returns a boolean if a field has been set.

### SetShipmentVolumeNil

`func (o *CreatePurchaseOrderShipment) SetShipmentVolumeNil(b bool)`

 SetShipmentVolumeNil sets the value for ShipmentVolume to be an explicit nil

### UnsetShipmentVolume
`func (o *CreatePurchaseOrderShipment) UnsetShipmentVolume()`

UnsetShipmentVolume ensures that no value is present for ShipmentVolume, not even an explicit nil
### GetLines

`func (o *CreatePurchaseOrderShipment) GetLines() []ChangePurchaseOrderShipmentLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *CreatePurchaseOrderShipment) GetLinesOk() (*[]ChangePurchaseOrderShipmentLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *CreatePurchaseOrderShipment) SetLines(v []ChangePurchaseOrderShipmentLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *CreatePurchaseOrderShipment) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *CreatePurchaseOrderShipment) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *CreatePurchaseOrderShipment) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


