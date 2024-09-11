# ChannelOrderLineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProductNo** | **string** | The unique product reference used by the channel. | 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the merchant. | [optional] 
**Quantity** | **int32** | The number of items of the product. | 
**CancellationRequestedQuantity** | Pointer to **int32** | The number of items for which cancellation was requested by the customer.  Some channels allow a customer to cancel an order until it has been shipped.  When an order has already been acknowledged in ChannelEngine, it can only be cancelled by creating a cancellation.  Use this field to check whether it is still possible to cancel the order. If this is the case, submit a cancellation to ChannelEngine. | [optional] 
**UnitPriceInclVat** | **float32** | The value of a single unit of the ordered product including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | 
**FeeFixed** | Pointer to **float32** | A fixed fee that is charged by the Channel for this orderline.  This fee rate is based on the currency of the Channel  This field is optional, send 0 if not applicable. | [optional] 
**FeeRate** | Pointer to **float32** | A percentage fee that is charged by the Channel for this orderline.  This field is optional, send 0 if not applicable. | [optional] 
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**ExactDeliveryDate** | Pointer to **NullableTime** | The exact date when the order (line) must be delivered. Delivery cannot occur before or after this date.  This can be empty if the channel does not provide this info. | [optional] 
**ExpectedDeliveryDate** | Pointer to **NullableTime** | The date when the order should be delivered. This field must always contain a value, either the exact date or the latest possible date.  If this value is not set, the system defaults to using the order date plus two days as the expected delivery date.  This field is used for backwards compatibility. | [optional] 
**LatestDeliveryDate** | Pointer to **NullableTime** | The latest possible date when the order (line) must be delivered. Delivery can occur on or before this date, but not after.  This can be empty if the channel does not provide this info. | [optional] 
**ExactShipmentDate** | Pointer to **NullableTime** | The exact date when the order (line) must be shipped. Shipment cannot occur before or after this date.  This can be empty if the channel does not provide this info. | [optional] 
**ExpectedShipmentDate** | Pointer to **NullableTime** | The date when when the order (line) should be shipped.  This can be empty if the channel does not provide this info. | [optional] 
**LatestShipmentDate** | Pointer to **NullableTime** | The latest possible date when the order (line) must be shipped. Shipment can occur on or before this date, but not after.  This can be empty if the channel does not provide this info. | [optional] 

## Methods

### NewChannelOrderLineRequest

`func NewChannelOrderLineRequest(channelProductNo string, quantity int32, unitPriceInclVat float32, ) *ChannelOrderLineRequest`

NewChannelOrderLineRequest instantiates a new ChannelOrderLineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOrderLineRequestWithDefaults

`func NewChannelOrderLineRequestWithDefaults() *ChannelOrderLineRequest`

NewChannelOrderLineRequestWithDefaults instantiates a new ChannelOrderLineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProductNo

`func (o *ChannelOrderLineRequest) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelOrderLineRequest) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelOrderLineRequest) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *ChannelOrderLineRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelOrderLineRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelOrderLineRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelOrderLineRequest) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelOrderLineRequest) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelOrderLineRequest) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *ChannelOrderLineRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelOrderLineRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelOrderLineRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCancellationRequestedQuantity

`func (o *ChannelOrderLineRequest) GetCancellationRequestedQuantity() int32`

GetCancellationRequestedQuantity returns the CancellationRequestedQuantity field if non-nil, zero value otherwise.

### GetCancellationRequestedQuantityOk

`func (o *ChannelOrderLineRequest) GetCancellationRequestedQuantityOk() (*int32, bool)`

GetCancellationRequestedQuantityOk returns a tuple with the CancellationRequestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequestedQuantity

`func (o *ChannelOrderLineRequest) SetCancellationRequestedQuantity(v int32)`

SetCancellationRequestedQuantity sets CancellationRequestedQuantity field to given value.

### HasCancellationRequestedQuantity

`func (o *ChannelOrderLineRequest) HasCancellationRequestedQuantity() bool`

HasCancellationRequestedQuantity returns a boolean if a field has been set.

### GetUnitPriceInclVat

`func (o *ChannelOrderLineRequest) GetUnitPriceInclVat() float32`

GetUnitPriceInclVat returns the UnitPriceInclVat field if non-nil, zero value otherwise.

### GetUnitPriceInclVatOk

`func (o *ChannelOrderLineRequest) GetUnitPriceInclVatOk() (*float32, bool)`

GetUnitPriceInclVatOk returns a tuple with the UnitPriceInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceInclVat

`func (o *ChannelOrderLineRequest) SetUnitPriceInclVat(v float32)`

SetUnitPriceInclVat sets UnitPriceInclVat field to given value.


### GetFeeFixed

`func (o *ChannelOrderLineRequest) GetFeeFixed() float32`

GetFeeFixed returns the FeeFixed field if non-nil, zero value otherwise.

### GetFeeFixedOk

`func (o *ChannelOrderLineRequest) GetFeeFixedOk() (*float32, bool)`

GetFeeFixedOk returns a tuple with the FeeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeFixed

`func (o *ChannelOrderLineRequest) SetFeeFixed(v float32)`

SetFeeFixed sets FeeFixed field to given value.

### HasFeeFixed

`func (o *ChannelOrderLineRequest) HasFeeFixed() bool`

HasFeeFixed returns a boolean if a field has been set.

### GetFeeRate

`func (o *ChannelOrderLineRequest) GetFeeRate() float32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ChannelOrderLineRequest) GetFeeRateOk() (*float32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ChannelOrderLineRequest) SetFeeRate(v float32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *ChannelOrderLineRequest) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetCondition

`func (o *ChannelOrderLineRequest) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ChannelOrderLineRequest) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ChannelOrderLineRequest) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ChannelOrderLineRequest) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetExactDeliveryDate

`func (o *ChannelOrderLineRequest) GetExactDeliveryDate() time.Time`

GetExactDeliveryDate returns the ExactDeliveryDate field if non-nil, zero value otherwise.

### GetExactDeliveryDateOk

`func (o *ChannelOrderLineRequest) GetExactDeliveryDateOk() (*time.Time, bool)`

GetExactDeliveryDateOk returns a tuple with the ExactDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactDeliveryDate

`func (o *ChannelOrderLineRequest) SetExactDeliveryDate(v time.Time)`

SetExactDeliveryDate sets ExactDeliveryDate field to given value.

### HasExactDeliveryDate

`func (o *ChannelOrderLineRequest) HasExactDeliveryDate() bool`

HasExactDeliveryDate returns a boolean if a field has been set.

### SetExactDeliveryDateNil

`func (o *ChannelOrderLineRequest) SetExactDeliveryDateNil(b bool)`

 SetExactDeliveryDateNil sets the value for ExactDeliveryDate to be an explicit nil

### UnsetExactDeliveryDate
`func (o *ChannelOrderLineRequest) UnsetExactDeliveryDate()`

UnsetExactDeliveryDate ensures that no value is present for ExactDeliveryDate, not even an explicit nil
### GetExpectedDeliveryDate

`func (o *ChannelOrderLineRequest) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *ChannelOrderLineRequest) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *ChannelOrderLineRequest) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *ChannelOrderLineRequest) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### SetExpectedDeliveryDateNil

`func (o *ChannelOrderLineRequest) SetExpectedDeliveryDateNil(b bool)`

 SetExpectedDeliveryDateNil sets the value for ExpectedDeliveryDate to be an explicit nil

### UnsetExpectedDeliveryDate
`func (o *ChannelOrderLineRequest) UnsetExpectedDeliveryDate()`

UnsetExpectedDeliveryDate ensures that no value is present for ExpectedDeliveryDate, not even an explicit nil
### GetLatestDeliveryDate

`func (o *ChannelOrderLineRequest) GetLatestDeliveryDate() time.Time`

GetLatestDeliveryDate returns the LatestDeliveryDate field if non-nil, zero value otherwise.

### GetLatestDeliveryDateOk

`func (o *ChannelOrderLineRequest) GetLatestDeliveryDateOk() (*time.Time, bool)`

GetLatestDeliveryDateOk returns a tuple with the LatestDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDate

`func (o *ChannelOrderLineRequest) SetLatestDeliveryDate(v time.Time)`

SetLatestDeliveryDate sets LatestDeliveryDate field to given value.

### HasLatestDeliveryDate

`func (o *ChannelOrderLineRequest) HasLatestDeliveryDate() bool`

HasLatestDeliveryDate returns a boolean if a field has been set.

### SetLatestDeliveryDateNil

`func (o *ChannelOrderLineRequest) SetLatestDeliveryDateNil(b bool)`

 SetLatestDeliveryDateNil sets the value for LatestDeliveryDate to be an explicit nil

### UnsetLatestDeliveryDate
`func (o *ChannelOrderLineRequest) UnsetLatestDeliveryDate()`

UnsetLatestDeliveryDate ensures that no value is present for LatestDeliveryDate, not even an explicit nil
### GetExactShipmentDate

`func (o *ChannelOrderLineRequest) GetExactShipmentDate() time.Time`

GetExactShipmentDate returns the ExactShipmentDate field if non-nil, zero value otherwise.

### GetExactShipmentDateOk

`func (o *ChannelOrderLineRequest) GetExactShipmentDateOk() (*time.Time, bool)`

GetExactShipmentDateOk returns a tuple with the ExactShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactShipmentDate

`func (o *ChannelOrderLineRequest) SetExactShipmentDate(v time.Time)`

SetExactShipmentDate sets ExactShipmentDate field to given value.

### HasExactShipmentDate

`func (o *ChannelOrderLineRequest) HasExactShipmentDate() bool`

HasExactShipmentDate returns a boolean if a field has been set.

### SetExactShipmentDateNil

`func (o *ChannelOrderLineRequest) SetExactShipmentDateNil(b bool)`

 SetExactShipmentDateNil sets the value for ExactShipmentDate to be an explicit nil

### UnsetExactShipmentDate
`func (o *ChannelOrderLineRequest) UnsetExactShipmentDate()`

UnsetExactShipmentDate ensures that no value is present for ExactShipmentDate, not even an explicit nil
### GetExpectedShipmentDate

`func (o *ChannelOrderLineRequest) GetExpectedShipmentDate() time.Time`

GetExpectedShipmentDate returns the ExpectedShipmentDate field if non-nil, zero value otherwise.

### GetExpectedShipmentDateOk

`func (o *ChannelOrderLineRequest) GetExpectedShipmentDateOk() (*time.Time, bool)`

GetExpectedShipmentDateOk returns a tuple with the ExpectedShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShipmentDate

`func (o *ChannelOrderLineRequest) SetExpectedShipmentDate(v time.Time)`

SetExpectedShipmentDate sets ExpectedShipmentDate field to given value.

### HasExpectedShipmentDate

`func (o *ChannelOrderLineRequest) HasExpectedShipmentDate() bool`

HasExpectedShipmentDate returns a boolean if a field has been set.

### SetExpectedShipmentDateNil

`func (o *ChannelOrderLineRequest) SetExpectedShipmentDateNil(b bool)`

 SetExpectedShipmentDateNil sets the value for ExpectedShipmentDate to be an explicit nil

### UnsetExpectedShipmentDate
`func (o *ChannelOrderLineRequest) UnsetExpectedShipmentDate()`

UnsetExpectedShipmentDate ensures that no value is present for ExpectedShipmentDate, not even an explicit nil
### GetLatestShipmentDate

`func (o *ChannelOrderLineRequest) GetLatestShipmentDate() time.Time`

GetLatestShipmentDate returns the LatestShipmentDate field if non-nil, zero value otherwise.

### GetLatestShipmentDateOk

`func (o *ChannelOrderLineRequest) GetLatestShipmentDateOk() (*time.Time, bool)`

GetLatestShipmentDateOk returns a tuple with the LatestShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestShipmentDate

`func (o *ChannelOrderLineRequest) SetLatestShipmentDate(v time.Time)`

SetLatestShipmentDate sets LatestShipmentDate field to given value.

### HasLatestShipmentDate

`func (o *ChannelOrderLineRequest) HasLatestShipmentDate() bool`

HasLatestShipmentDate returns a boolean if a field has been set.

### SetLatestShipmentDateNil

`func (o *ChannelOrderLineRequest) SetLatestShipmentDateNil(b bool)`

 SetLatestShipmentDateNil sets the value for LatestShipmentDate to be an explicit nil

### UnsetLatestShipmentDate
`func (o *ChannelOrderLineRequest) UnsetLatestShipmentDate()`

UnsetLatestShipmentDate ensures that no value is present for LatestShipmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


