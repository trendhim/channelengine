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
**ExpectedDeliveryDate** | Pointer to **NullableTime** | Expected delivery date from channels, empty if channels not support this value | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


