# ChannelOrderLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**OrderStatusView**](OrderStatusView.md) |  | [optional] 
**IsFulfillmentByMarketplace** | Pointer to **bool** | Is the order fulfilled by the marketplace (amazon: FBA, bol: LVB, etc.)?. | [optional] 
**Gtin** | Pointer to **NullableString** | Either the GTIN (EAN, ISBN, UPC etc) provided by the channel, or the the GTIN belonging to the MerchantProductNo in ChannelEngine. | [optional] 
**Description** | Pointer to **NullableString** | The product description (or title) provided by the channel. | [optional] 
**UnitVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the value of a single unit of the ordered product  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**LineTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order line (quantity * unit price) including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**LineVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of the order line (quantity * unit price)  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**OriginalUnitPriceInclVat** | Pointer to **NullableFloat32** | The value of a single unit of the ordered product including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalUnitVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the value of a single unit of the ordered product  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalLineTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order line (quantity * unit price) including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalLineVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of the order line (quantity * unit price)  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalFeeFixed** | Pointer to **float32** | A percentage fee that is charged by the Channel for this orderline.  This fee rate is based on the currency of client  This field is optional, send 0 if not applicable. | [optional] 
**BundleProductMerchantProductNo** | Pointer to **NullableString** | If the product is ordered part of a bundle, this field contains the MerchantProductNo of  the product bundle. | [optional] 
**ChannelProductNo** | **string** | The unique product reference used by the channel. | 
**MerchantProductNo** | Pointer to **NullableString** | The unique product reference used by the merchant. | [optional] 
**Quantity** | **int32** | The number of items of the product. | 
**CancellationRequestedQuantity** | Pointer to **int32** | The number of items for which cancellation was requested by the customer.  Some channels allow a customer to cancel an order until it has been shipped.  When an order has already been acknowledged in ChannelEngine, it can only be cancelled by creating a cancellation.  Use this field to check whether it is still possible to cancel the order. If this is the case, submit a cancellation to ChannelEngine. | [optional] 
**UnitPriceInclVat** | **float32** | The value of a single unit of the ordered product including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | 
**FeeFixed** | Pointer to **float32** | A fixed fee that is charged by the Channel for this orderline.  This fee rate is based on the currency of the Channel  This field is optional, send 0 if not applicable. | [optional] 
**FeeRate** | Pointer to **float32** | A percentage fee that is charged by the Channel for this orderline.  This field is optional, send 0 if not applicable. | [optional] 
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**ExactDeliveryDate** | Pointer to **NullableTime** | Exact delivery date from channels, empty if channels not support this value | [optional] 
**ExpectedDeliveryDate** | Pointer to **NullableTime** | Expected delivery date from channels, empty if channels not support this value | [optional] 
**LatestDeliveryDate** | Pointer to **NullableTime** | Latest delivery date from channels, empty if channels not support this value | [optional] 
**ExpectedShipmentDate** | Pointer to **NullableTime** | Expected shipment date from channels, empty if channels not support this value | [optional] 
**LatestShipmentDate** | Pointer to **NullableTime** | Latest shipment date from channels, empty if channels not support this value | [optional] 

## Methods

### NewChannelOrderLineResponse

`func NewChannelOrderLineResponse(channelProductNo string, quantity int32, unitPriceInclVat float32, ) *ChannelOrderLineResponse`

NewChannelOrderLineResponse instantiates a new ChannelOrderLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOrderLineResponseWithDefaults

`func NewChannelOrderLineResponseWithDefaults() *ChannelOrderLineResponse`

NewChannelOrderLineResponseWithDefaults instantiates a new ChannelOrderLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChannelOrderLineResponse) GetStatus() OrderStatusView`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChannelOrderLineResponse) GetStatusOk() (*OrderStatusView, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChannelOrderLineResponse) SetStatus(v OrderStatusView)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChannelOrderLineResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsFulfillmentByMarketplace

`func (o *ChannelOrderLineResponse) GetIsFulfillmentByMarketplace() bool`

GetIsFulfillmentByMarketplace returns the IsFulfillmentByMarketplace field if non-nil, zero value otherwise.

### GetIsFulfillmentByMarketplaceOk

`func (o *ChannelOrderLineResponse) GetIsFulfillmentByMarketplaceOk() (*bool, bool)`

GetIsFulfillmentByMarketplaceOk returns a tuple with the IsFulfillmentByMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFulfillmentByMarketplace

`func (o *ChannelOrderLineResponse) SetIsFulfillmentByMarketplace(v bool)`

SetIsFulfillmentByMarketplace sets IsFulfillmentByMarketplace field to given value.

### HasIsFulfillmentByMarketplace

`func (o *ChannelOrderLineResponse) HasIsFulfillmentByMarketplace() bool`

HasIsFulfillmentByMarketplace returns a boolean if a field has been set.

### GetGtin

`func (o *ChannelOrderLineResponse) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ChannelOrderLineResponse) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ChannelOrderLineResponse) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ChannelOrderLineResponse) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### SetGtinNil

`func (o *ChannelOrderLineResponse) SetGtinNil(b bool)`

 SetGtinNil sets the value for Gtin to be an explicit nil

### UnsetGtin
`func (o *ChannelOrderLineResponse) UnsetGtin()`

UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
### GetDescription

`func (o *ChannelOrderLineResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelOrderLineResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelOrderLineResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelOrderLineResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ChannelOrderLineResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ChannelOrderLineResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUnitVat

`func (o *ChannelOrderLineResponse) GetUnitVat() float32`

GetUnitVat returns the UnitVat field if non-nil, zero value otherwise.

### GetUnitVatOk

`func (o *ChannelOrderLineResponse) GetUnitVatOk() (*float32, bool)`

GetUnitVatOk returns a tuple with the UnitVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitVat

`func (o *ChannelOrderLineResponse) SetUnitVat(v float32)`

SetUnitVat sets UnitVat field to given value.

### HasUnitVat

`func (o *ChannelOrderLineResponse) HasUnitVat() bool`

HasUnitVat returns a boolean if a field has been set.

### SetUnitVatNil

`func (o *ChannelOrderLineResponse) SetUnitVatNil(b bool)`

 SetUnitVatNil sets the value for UnitVat to be an explicit nil

### UnsetUnitVat
`func (o *ChannelOrderLineResponse) UnsetUnitVat()`

UnsetUnitVat ensures that no value is present for UnitVat, not even an explicit nil
### GetLineTotalInclVat

`func (o *ChannelOrderLineResponse) GetLineTotalInclVat() float32`

GetLineTotalInclVat returns the LineTotalInclVat field if non-nil, zero value otherwise.

### GetLineTotalInclVatOk

`func (o *ChannelOrderLineResponse) GetLineTotalInclVatOk() (*float32, bool)`

GetLineTotalInclVatOk returns a tuple with the LineTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalInclVat

`func (o *ChannelOrderLineResponse) SetLineTotalInclVat(v float32)`

SetLineTotalInclVat sets LineTotalInclVat field to given value.

### HasLineTotalInclVat

`func (o *ChannelOrderLineResponse) HasLineTotalInclVat() bool`

HasLineTotalInclVat returns a boolean if a field has been set.

### SetLineTotalInclVatNil

`func (o *ChannelOrderLineResponse) SetLineTotalInclVatNil(b bool)`

 SetLineTotalInclVatNil sets the value for LineTotalInclVat to be an explicit nil

### UnsetLineTotalInclVat
`func (o *ChannelOrderLineResponse) UnsetLineTotalInclVat()`

UnsetLineTotalInclVat ensures that no value is present for LineTotalInclVat, not even an explicit nil
### GetLineVat

`func (o *ChannelOrderLineResponse) GetLineVat() float32`

GetLineVat returns the LineVat field if non-nil, zero value otherwise.

### GetLineVatOk

`func (o *ChannelOrderLineResponse) GetLineVatOk() (*float32, bool)`

GetLineVatOk returns a tuple with the LineVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineVat

`func (o *ChannelOrderLineResponse) SetLineVat(v float32)`

SetLineVat sets LineVat field to given value.

### HasLineVat

`func (o *ChannelOrderLineResponse) HasLineVat() bool`

HasLineVat returns a boolean if a field has been set.

### SetLineVatNil

`func (o *ChannelOrderLineResponse) SetLineVatNil(b bool)`

 SetLineVatNil sets the value for LineVat to be an explicit nil

### UnsetLineVat
`func (o *ChannelOrderLineResponse) UnsetLineVat()`

UnsetLineVat ensures that no value is present for LineVat, not even an explicit nil
### GetOriginalUnitPriceInclVat

`func (o *ChannelOrderLineResponse) GetOriginalUnitPriceInclVat() float32`

GetOriginalUnitPriceInclVat returns the OriginalUnitPriceInclVat field if non-nil, zero value otherwise.

### GetOriginalUnitPriceInclVatOk

`func (o *ChannelOrderLineResponse) GetOriginalUnitPriceInclVatOk() (*float32, bool)`

GetOriginalUnitPriceInclVatOk returns a tuple with the OriginalUnitPriceInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitPriceInclVat

`func (o *ChannelOrderLineResponse) SetOriginalUnitPriceInclVat(v float32)`

SetOriginalUnitPriceInclVat sets OriginalUnitPriceInclVat field to given value.

### HasOriginalUnitPriceInclVat

`func (o *ChannelOrderLineResponse) HasOriginalUnitPriceInclVat() bool`

HasOriginalUnitPriceInclVat returns a boolean if a field has been set.

### SetOriginalUnitPriceInclVatNil

`func (o *ChannelOrderLineResponse) SetOriginalUnitPriceInclVatNil(b bool)`

 SetOriginalUnitPriceInclVatNil sets the value for OriginalUnitPriceInclVat to be an explicit nil

### UnsetOriginalUnitPriceInclVat
`func (o *ChannelOrderLineResponse) UnsetOriginalUnitPriceInclVat()`

UnsetOriginalUnitPriceInclVat ensures that no value is present for OriginalUnitPriceInclVat, not even an explicit nil
### GetOriginalUnitVat

`func (o *ChannelOrderLineResponse) GetOriginalUnitVat() float32`

GetOriginalUnitVat returns the OriginalUnitVat field if non-nil, zero value otherwise.

### GetOriginalUnitVatOk

`func (o *ChannelOrderLineResponse) GetOriginalUnitVatOk() (*float32, bool)`

GetOriginalUnitVatOk returns a tuple with the OriginalUnitVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitVat

`func (o *ChannelOrderLineResponse) SetOriginalUnitVat(v float32)`

SetOriginalUnitVat sets OriginalUnitVat field to given value.

### HasOriginalUnitVat

`func (o *ChannelOrderLineResponse) HasOriginalUnitVat() bool`

HasOriginalUnitVat returns a boolean if a field has been set.

### SetOriginalUnitVatNil

`func (o *ChannelOrderLineResponse) SetOriginalUnitVatNil(b bool)`

 SetOriginalUnitVatNil sets the value for OriginalUnitVat to be an explicit nil

### UnsetOriginalUnitVat
`func (o *ChannelOrderLineResponse) UnsetOriginalUnitVat()`

UnsetOriginalUnitVat ensures that no value is present for OriginalUnitVat, not even an explicit nil
### GetOriginalLineTotalInclVat

`func (o *ChannelOrderLineResponse) GetOriginalLineTotalInclVat() float32`

GetOriginalLineTotalInclVat returns the OriginalLineTotalInclVat field if non-nil, zero value otherwise.

### GetOriginalLineTotalInclVatOk

`func (o *ChannelOrderLineResponse) GetOriginalLineTotalInclVatOk() (*float32, bool)`

GetOriginalLineTotalInclVatOk returns a tuple with the OriginalLineTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineTotalInclVat

`func (o *ChannelOrderLineResponse) SetOriginalLineTotalInclVat(v float32)`

SetOriginalLineTotalInclVat sets OriginalLineTotalInclVat field to given value.

### HasOriginalLineTotalInclVat

`func (o *ChannelOrderLineResponse) HasOriginalLineTotalInclVat() bool`

HasOriginalLineTotalInclVat returns a boolean if a field has been set.

### SetOriginalLineTotalInclVatNil

`func (o *ChannelOrderLineResponse) SetOriginalLineTotalInclVatNil(b bool)`

 SetOriginalLineTotalInclVatNil sets the value for OriginalLineTotalInclVat to be an explicit nil

### UnsetOriginalLineTotalInclVat
`func (o *ChannelOrderLineResponse) UnsetOriginalLineTotalInclVat()`

UnsetOriginalLineTotalInclVat ensures that no value is present for OriginalLineTotalInclVat, not even an explicit nil
### GetOriginalLineVat

`func (o *ChannelOrderLineResponse) GetOriginalLineVat() float32`

GetOriginalLineVat returns the OriginalLineVat field if non-nil, zero value otherwise.

### GetOriginalLineVatOk

`func (o *ChannelOrderLineResponse) GetOriginalLineVatOk() (*float32, bool)`

GetOriginalLineVatOk returns a tuple with the OriginalLineVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineVat

`func (o *ChannelOrderLineResponse) SetOriginalLineVat(v float32)`

SetOriginalLineVat sets OriginalLineVat field to given value.

### HasOriginalLineVat

`func (o *ChannelOrderLineResponse) HasOriginalLineVat() bool`

HasOriginalLineVat returns a boolean if a field has been set.

### SetOriginalLineVatNil

`func (o *ChannelOrderLineResponse) SetOriginalLineVatNil(b bool)`

 SetOriginalLineVatNil sets the value for OriginalLineVat to be an explicit nil

### UnsetOriginalLineVat
`func (o *ChannelOrderLineResponse) UnsetOriginalLineVat()`

UnsetOriginalLineVat ensures that no value is present for OriginalLineVat, not even an explicit nil
### GetOriginalFeeFixed

`func (o *ChannelOrderLineResponse) GetOriginalFeeFixed() float32`

GetOriginalFeeFixed returns the OriginalFeeFixed field if non-nil, zero value otherwise.

### GetOriginalFeeFixedOk

`func (o *ChannelOrderLineResponse) GetOriginalFeeFixedOk() (*float32, bool)`

GetOriginalFeeFixedOk returns a tuple with the OriginalFeeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFeeFixed

`func (o *ChannelOrderLineResponse) SetOriginalFeeFixed(v float32)`

SetOriginalFeeFixed sets OriginalFeeFixed field to given value.

### HasOriginalFeeFixed

`func (o *ChannelOrderLineResponse) HasOriginalFeeFixed() bool`

HasOriginalFeeFixed returns a boolean if a field has been set.

### GetBundleProductMerchantProductNo

`func (o *ChannelOrderLineResponse) GetBundleProductMerchantProductNo() string`

GetBundleProductMerchantProductNo returns the BundleProductMerchantProductNo field if non-nil, zero value otherwise.

### GetBundleProductMerchantProductNoOk

`func (o *ChannelOrderLineResponse) GetBundleProductMerchantProductNoOk() (*string, bool)`

GetBundleProductMerchantProductNoOk returns a tuple with the BundleProductMerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleProductMerchantProductNo

`func (o *ChannelOrderLineResponse) SetBundleProductMerchantProductNo(v string)`

SetBundleProductMerchantProductNo sets BundleProductMerchantProductNo field to given value.

### HasBundleProductMerchantProductNo

`func (o *ChannelOrderLineResponse) HasBundleProductMerchantProductNo() bool`

HasBundleProductMerchantProductNo returns a boolean if a field has been set.

### SetBundleProductMerchantProductNoNil

`func (o *ChannelOrderLineResponse) SetBundleProductMerchantProductNoNil(b bool)`

 SetBundleProductMerchantProductNoNil sets the value for BundleProductMerchantProductNo to be an explicit nil

### UnsetBundleProductMerchantProductNo
`func (o *ChannelOrderLineResponse) UnsetBundleProductMerchantProductNo()`

UnsetBundleProductMerchantProductNo ensures that no value is present for BundleProductMerchantProductNo, not even an explicit nil
### GetChannelProductNo

`func (o *ChannelOrderLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelOrderLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelOrderLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *ChannelOrderLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelOrderLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelOrderLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *ChannelOrderLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *ChannelOrderLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *ChannelOrderLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *ChannelOrderLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChannelOrderLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChannelOrderLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCancellationRequestedQuantity

`func (o *ChannelOrderLineResponse) GetCancellationRequestedQuantity() int32`

GetCancellationRequestedQuantity returns the CancellationRequestedQuantity field if non-nil, zero value otherwise.

### GetCancellationRequestedQuantityOk

`func (o *ChannelOrderLineResponse) GetCancellationRequestedQuantityOk() (*int32, bool)`

GetCancellationRequestedQuantityOk returns a tuple with the CancellationRequestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequestedQuantity

`func (o *ChannelOrderLineResponse) SetCancellationRequestedQuantity(v int32)`

SetCancellationRequestedQuantity sets CancellationRequestedQuantity field to given value.

### HasCancellationRequestedQuantity

`func (o *ChannelOrderLineResponse) HasCancellationRequestedQuantity() bool`

HasCancellationRequestedQuantity returns a boolean if a field has been set.

### GetUnitPriceInclVat

`func (o *ChannelOrderLineResponse) GetUnitPriceInclVat() float32`

GetUnitPriceInclVat returns the UnitPriceInclVat field if non-nil, zero value otherwise.

### GetUnitPriceInclVatOk

`func (o *ChannelOrderLineResponse) GetUnitPriceInclVatOk() (*float32, bool)`

GetUnitPriceInclVatOk returns a tuple with the UnitPriceInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceInclVat

`func (o *ChannelOrderLineResponse) SetUnitPriceInclVat(v float32)`

SetUnitPriceInclVat sets UnitPriceInclVat field to given value.


### GetFeeFixed

`func (o *ChannelOrderLineResponse) GetFeeFixed() float32`

GetFeeFixed returns the FeeFixed field if non-nil, zero value otherwise.

### GetFeeFixedOk

`func (o *ChannelOrderLineResponse) GetFeeFixedOk() (*float32, bool)`

GetFeeFixedOk returns a tuple with the FeeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeFixed

`func (o *ChannelOrderLineResponse) SetFeeFixed(v float32)`

SetFeeFixed sets FeeFixed field to given value.

### HasFeeFixed

`func (o *ChannelOrderLineResponse) HasFeeFixed() bool`

HasFeeFixed returns a boolean if a field has been set.

### GetFeeRate

`func (o *ChannelOrderLineResponse) GetFeeRate() float32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ChannelOrderLineResponse) GetFeeRateOk() (*float32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ChannelOrderLineResponse) SetFeeRate(v float32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *ChannelOrderLineResponse) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetCondition

`func (o *ChannelOrderLineResponse) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ChannelOrderLineResponse) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ChannelOrderLineResponse) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ChannelOrderLineResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetExactDeliveryDate

`func (o *ChannelOrderLineResponse) GetExactDeliveryDate() time.Time`

GetExactDeliveryDate returns the ExactDeliveryDate field if non-nil, zero value otherwise.

### GetExactDeliveryDateOk

`func (o *ChannelOrderLineResponse) GetExactDeliveryDateOk() (*time.Time, bool)`

GetExactDeliveryDateOk returns a tuple with the ExactDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactDeliveryDate

`func (o *ChannelOrderLineResponse) SetExactDeliveryDate(v time.Time)`

SetExactDeliveryDate sets ExactDeliveryDate field to given value.

### HasExactDeliveryDate

`func (o *ChannelOrderLineResponse) HasExactDeliveryDate() bool`

HasExactDeliveryDate returns a boolean if a field has been set.

### SetExactDeliveryDateNil

`func (o *ChannelOrderLineResponse) SetExactDeliveryDateNil(b bool)`

 SetExactDeliveryDateNil sets the value for ExactDeliveryDate to be an explicit nil

### UnsetExactDeliveryDate
`func (o *ChannelOrderLineResponse) UnsetExactDeliveryDate()`

UnsetExactDeliveryDate ensures that no value is present for ExactDeliveryDate, not even an explicit nil
### GetExpectedDeliveryDate

`func (o *ChannelOrderLineResponse) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *ChannelOrderLineResponse) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *ChannelOrderLineResponse) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *ChannelOrderLineResponse) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### SetExpectedDeliveryDateNil

`func (o *ChannelOrderLineResponse) SetExpectedDeliveryDateNil(b bool)`

 SetExpectedDeliveryDateNil sets the value for ExpectedDeliveryDate to be an explicit nil

### UnsetExpectedDeliveryDate
`func (o *ChannelOrderLineResponse) UnsetExpectedDeliveryDate()`

UnsetExpectedDeliveryDate ensures that no value is present for ExpectedDeliveryDate, not even an explicit nil
### GetLatestDeliveryDate

`func (o *ChannelOrderLineResponse) GetLatestDeliveryDate() time.Time`

GetLatestDeliveryDate returns the LatestDeliveryDate field if non-nil, zero value otherwise.

### GetLatestDeliveryDateOk

`func (o *ChannelOrderLineResponse) GetLatestDeliveryDateOk() (*time.Time, bool)`

GetLatestDeliveryDateOk returns a tuple with the LatestDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDate

`func (o *ChannelOrderLineResponse) SetLatestDeliveryDate(v time.Time)`

SetLatestDeliveryDate sets LatestDeliveryDate field to given value.

### HasLatestDeliveryDate

`func (o *ChannelOrderLineResponse) HasLatestDeliveryDate() bool`

HasLatestDeliveryDate returns a boolean if a field has been set.

### SetLatestDeliveryDateNil

`func (o *ChannelOrderLineResponse) SetLatestDeliveryDateNil(b bool)`

 SetLatestDeliveryDateNil sets the value for LatestDeliveryDate to be an explicit nil

### UnsetLatestDeliveryDate
`func (o *ChannelOrderLineResponse) UnsetLatestDeliveryDate()`

UnsetLatestDeliveryDate ensures that no value is present for LatestDeliveryDate, not even an explicit nil
### GetExpectedShipmentDate

`func (o *ChannelOrderLineResponse) GetExpectedShipmentDate() time.Time`

GetExpectedShipmentDate returns the ExpectedShipmentDate field if non-nil, zero value otherwise.

### GetExpectedShipmentDateOk

`func (o *ChannelOrderLineResponse) GetExpectedShipmentDateOk() (*time.Time, bool)`

GetExpectedShipmentDateOk returns a tuple with the ExpectedShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShipmentDate

`func (o *ChannelOrderLineResponse) SetExpectedShipmentDate(v time.Time)`

SetExpectedShipmentDate sets ExpectedShipmentDate field to given value.

### HasExpectedShipmentDate

`func (o *ChannelOrderLineResponse) HasExpectedShipmentDate() bool`

HasExpectedShipmentDate returns a boolean if a field has been set.

### SetExpectedShipmentDateNil

`func (o *ChannelOrderLineResponse) SetExpectedShipmentDateNil(b bool)`

 SetExpectedShipmentDateNil sets the value for ExpectedShipmentDate to be an explicit nil

### UnsetExpectedShipmentDate
`func (o *ChannelOrderLineResponse) UnsetExpectedShipmentDate()`

UnsetExpectedShipmentDate ensures that no value is present for ExpectedShipmentDate, not even an explicit nil
### GetLatestShipmentDate

`func (o *ChannelOrderLineResponse) GetLatestShipmentDate() time.Time`

GetLatestShipmentDate returns the LatestShipmentDate field if non-nil, zero value otherwise.

### GetLatestShipmentDateOk

`func (o *ChannelOrderLineResponse) GetLatestShipmentDateOk() (*time.Time, bool)`

GetLatestShipmentDateOk returns a tuple with the LatestShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestShipmentDate

`func (o *ChannelOrderLineResponse) SetLatestShipmentDate(v time.Time)`

SetLatestShipmentDate sets LatestShipmentDate field to given value.

### HasLatestShipmentDate

`func (o *ChannelOrderLineResponse) HasLatestShipmentDate() bool`

HasLatestShipmentDate returns a boolean if a field has been set.

### SetLatestShipmentDateNil

`func (o *ChannelOrderLineResponse) SetLatestShipmentDateNil(b bool)`

 SetLatestShipmentDateNil sets the value for LatestShipmentDate to be an explicit nil

### UnsetLatestShipmentDate
`func (o *ChannelOrderLineResponse) UnsetLatestShipmentDate()`

UnsetLatestShipmentDate ensures that no value is present for LatestShipmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


