# MerchantAcknowledgePurchaseOrderLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderLineIdentifier** | Pointer to **NullableString** |  | [optional] 
**AcknowledgementCode** | Pointer to [**PurchaseOrderAcknowledgementCode**](PurchaseOrderAcknowledgementCode.md) |  | [optional] 
**AcknowledgedQuantity** | Pointer to **int32** |  | [optional] 
**RejectionReason** | Pointer to [**PurchaseOrderRejectionReason**](PurchaseOrderRejectionReason.md) |  | [optional] 
**ScheduledShipDate** | Pointer to **NullableTime** |  | [optional] 
**ScheduledDeliveryDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewMerchantAcknowledgePurchaseOrderLine

`func NewMerchantAcknowledgePurchaseOrderLine() *MerchantAcknowledgePurchaseOrderLine`

NewMerchantAcknowledgePurchaseOrderLine instantiates a new MerchantAcknowledgePurchaseOrderLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAcknowledgePurchaseOrderLineWithDefaults

`func NewMerchantAcknowledgePurchaseOrderLineWithDefaults() *MerchantAcknowledgePurchaseOrderLine`

NewMerchantAcknowledgePurchaseOrderLineWithDefaults instantiates a new MerchantAcknowledgePurchaseOrderLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderLineIdentifier

`func (o *MerchantAcknowledgePurchaseOrderLine) GetOrderLineIdentifier() string`

GetOrderLineIdentifier returns the OrderLineIdentifier field if non-nil, zero value otherwise.

### GetOrderLineIdentifierOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetOrderLineIdentifierOk() (*string, bool)`

GetOrderLineIdentifierOk returns a tuple with the OrderLineIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineIdentifier

`func (o *MerchantAcknowledgePurchaseOrderLine) SetOrderLineIdentifier(v string)`

SetOrderLineIdentifier sets OrderLineIdentifier field to given value.

### HasOrderLineIdentifier

`func (o *MerchantAcknowledgePurchaseOrderLine) HasOrderLineIdentifier() bool`

HasOrderLineIdentifier returns a boolean if a field has been set.

### SetOrderLineIdentifierNil

`func (o *MerchantAcknowledgePurchaseOrderLine) SetOrderLineIdentifierNil(b bool)`

 SetOrderLineIdentifierNil sets the value for OrderLineIdentifier to be an explicit nil

### UnsetOrderLineIdentifier
`func (o *MerchantAcknowledgePurchaseOrderLine) UnsetOrderLineIdentifier()`

UnsetOrderLineIdentifier ensures that no value is present for OrderLineIdentifier, not even an explicit nil
### GetAcknowledgementCode

`func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgementCode() PurchaseOrderAcknowledgementCode`

GetAcknowledgementCode returns the AcknowledgementCode field if non-nil, zero value otherwise.

### GetAcknowledgementCodeOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgementCodeOk() (*PurchaseOrderAcknowledgementCode, bool)`

GetAcknowledgementCodeOk returns a tuple with the AcknowledgementCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementCode

`func (o *MerchantAcknowledgePurchaseOrderLine) SetAcknowledgementCode(v PurchaseOrderAcknowledgementCode)`

SetAcknowledgementCode sets AcknowledgementCode field to given value.

### HasAcknowledgementCode

`func (o *MerchantAcknowledgePurchaseOrderLine) HasAcknowledgementCode() bool`

HasAcknowledgementCode returns a boolean if a field has been set.

### GetAcknowledgedQuantity

`func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgedQuantity() int32`

GetAcknowledgedQuantity returns the AcknowledgedQuantity field if non-nil, zero value otherwise.

### GetAcknowledgedQuantityOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetAcknowledgedQuantityOk() (*int32, bool)`

GetAcknowledgedQuantityOk returns a tuple with the AcknowledgedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedQuantity

`func (o *MerchantAcknowledgePurchaseOrderLine) SetAcknowledgedQuantity(v int32)`

SetAcknowledgedQuantity sets AcknowledgedQuantity field to given value.

### HasAcknowledgedQuantity

`func (o *MerchantAcknowledgePurchaseOrderLine) HasAcknowledgedQuantity() bool`

HasAcknowledgedQuantity returns a boolean if a field has been set.

### GetRejectionReason

`func (o *MerchantAcknowledgePurchaseOrderLine) GetRejectionReason() PurchaseOrderRejectionReason`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetRejectionReasonOk() (*PurchaseOrderRejectionReason, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *MerchantAcknowledgePurchaseOrderLine) SetRejectionReason(v PurchaseOrderRejectionReason)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *MerchantAcknowledgePurchaseOrderLine) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetScheduledShipDate

`func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledShipDate() time.Time`

GetScheduledShipDate returns the ScheduledShipDate field if non-nil, zero value otherwise.

### GetScheduledShipDateOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledShipDateOk() (*time.Time, bool)`

GetScheduledShipDateOk returns a tuple with the ScheduledShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledShipDate

`func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledShipDate(v time.Time)`

SetScheduledShipDate sets ScheduledShipDate field to given value.

### HasScheduledShipDate

`func (o *MerchantAcknowledgePurchaseOrderLine) HasScheduledShipDate() bool`

HasScheduledShipDate returns a boolean if a field has been set.

### SetScheduledShipDateNil

`func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledShipDateNil(b bool)`

 SetScheduledShipDateNil sets the value for ScheduledShipDate to be an explicit nil

### UnsetScheduledShipDate
`func (o *MerchantAcknowledgePurchaseOrderLine) UnsetScheduledShipDate()`

UnsetScheduledShipDate ensures that no value is present for ScheduledShipDate, not even an explicit nil
### GetScheduledDeliveryDate

`func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledDeliveryDate() time.Time`

GetScheduledDeliveryDate returns the ScheduledDeliveryDate field if non-nil, zero value otherwise.

### GetScheduledDeliveryDateOk

`func (o *MerchantAcknowledgePurchaseOrderLine) GetScheduledDeliveryDateOk() (*time.Time, bool)`

GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryDate

`func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledDeliveryDate(v time.Time)`

SetScheduledDeliveryDate sets ScheduledDeliveryDate field to given value.

### HasScheduledDeliveryDate

`func (o *MerchantAcknowledgePurchaseOrderLine) HasScheduledDeliveryDate() bool`

HasScheduledDeliveryDate returns a boolean if a field has been set.

### SetScheduledDeliveryDateNil

`func (o *MerchantAcknowledgePurchaseOrderLine) SetScheduledDeliveryDateNil(b bool)`

 SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil

### UnsetScheduledDeliveryDate
`func (o *MerchantAcknowledgePurchaseOrderLine) UnsetScheduledDeliveryDate()`

UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


