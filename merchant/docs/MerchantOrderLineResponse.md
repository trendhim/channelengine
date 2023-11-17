# MerchantOrderLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**OrderStatusView**](OrderStatusView.md) |  | [optional] 
**IsFulfillmentByMarketplace** | Pointer to **bool** | Is the order fulfilled by the marketplace (amazon: FBA, bol: LVB, etc.)?. | [optional] 
**Gtin** | Pointer to **NullableString** | Either the GTIN (EAN, ISBN, UPC etc) provided by the channel, or the the GTIN belonging to the MerchantProductNo in ChannelEngine. | [optional] 
**Description** | Pointer to **NullableString** | The product description (or title) provided by the channel. | [optional] 
**StockLocation** | Pointer to [**MerchantStockLocationResponse**](MerchantStockLocationResponse.md) |  | [optional] 
**UnitVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the value of a single unit of the ordered product  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**LineTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order line (quantity * unit price) including VAT  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**LineVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of the order line (quantity * unit price)  (in the shop&#39;s base currency calculated using the exchange rate at the time of ordering). | [optional] 
**OriginalUnitPriceInclVat** | Pointer to **NullableFloat32** | The value of a single unit of the ordered product including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalUnitVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the value of a single unit of the ordered product  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalLineTotalInclVat** | Pointer to **NullableFloat32** | The total value of the order line (quantity * unit price) including VAT  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalLineVat** | Pointer to **NullableFloat32** | The total amount of VAT charged over the total value of the order line (quantity * unit price)  (in the currency in which the order was paid for, see CurrencyCode). | [optional] 
**OriginalFeeFixed** | Pointer to **float32** | A percentage fee that is charged by the Channel for this orderline.  This fee rate is based on the currency of client  This field is optional, send 0 if not applicable. | [optional] 
**BundleProductMerchantProductNo** | Pointer to **NullableString** | If the product is ordered part of a bundle, this field contains the MerchantProductNo of  the product bundle. | [optional] 
**JurisCode** | Pointer to **NullableString** | State assigned code identifying the jurisdiction. | [optional] 
**JurisName** | Pointer to **NullableString** | Name of a tax jurisdiction. | [optional] 
**VatRate** | Pointer to **float32** | VAT rate of the orderline. | [optional] 
**UnitPriceExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**LineTotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalUnitPriceExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**OriginalLineTotalExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**ExtraData** | Pointer to [**[]MerchantOrderLineExtraDataResponse**](MerchantOrderLineExtraDataResponse.md) |  | [optional] 
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

### NewMerchantOrderLineResponse

`func NewMerchantOrderLineResponse(channelProductNo string, quantity int32, unitPriceInclVat float32, ) *MerchantOrderLineResponse`

NewMerchantOrderLineResponse instantiates a new MerchantOrderLineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantOrderLineResponseWithDefaults

`func NewMerchantOrderLineResponseWithDefaults() *MerchantOrderLineResponse`

NewMerchantOrderLineResponseWithDefaults instantiates a new MerchantOrderLineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MerchantOrderLineResponse) GetStatus() OrderStatusView`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantOrderLineResponse) GetStatusOk() (*OrderStatusView, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantOrderLineResponse) SetStatus(v OrderStatusView)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MerchantOrderLineResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsFulfillmentByMarketplace

`func (o *MerchantOrderLineResponse) GetIsFulfillmentByMarketplace() bool`

GetIsFulfillmentByMarketplace returns the IsFulfillmentByMarketplace field if non-nil, zero value otherwise.

### GetIsFulfillmentByMarketplaceOk

`func (o *MerchantOrderLineResponse) GetIsFulfillmentByMarketplaceOk() (*bool, bool)`

GetIsFulfillmentByMarketplaceOk returns a tuple with the IsFulfillmentByMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFulfillmentByMarketplace

`func (o *MerchantOrderLineResponse) SetIsFulfillmentByMarketplace(v bool)`

SetIsFulfillmentByMarketplace sets IsFulfillmentByMarketplace field to given value.

### HasIsFulfillmentByMarketplace

`func (o *MerchantOrderLineResponse) HasIsFulfillmentByMarketplace() bool`

HasIsFulfillmentByMarketplace returns a boolean if a field has been set.

### GetGtin

`func (o *MerchantOrderLineResponse) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *MerchantOrderLineResponse) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *MerchantOrderLineResponse) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *MerchantOrderLineResponse) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### SetGtinNil

`func (o *MerchantOrderLineResponse) SetGtinNil(b bool)`

 SetGtinNil sets the value for Gtin to be an explicit nil

### UnsetGtin
`func (o *MerchantOrderLineResponse) UnsetGtin()`

UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
### GetDescription

`func (o *MerchantOrderLineResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MerchantOrderLineResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MerchantOrderLineResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MerchantOrderLineResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MerchantOrderLineResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MerchantOrderLineResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStockLocation

`func (o *MerchantOrderLineResponse) GetStockLocation() MerchantStockLocationResponse`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *MerchantOrderLineResponse) GetStockLocationOk() (*MerchantStockLocationResponse, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *MerchantOrderLineResponse) SetStockLocation(v MerchantStockLocationResponse)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *MerchantOrderLineResponse) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetUnitVat

`func (o *MerchantOrderLineResponse) GetUnitVat() float32`

GetUnitVat returns the UnitVat field if non-nil, zero value otherwise.

### GetUnitVatOk

`func (o *MerchantOrderLineResponse) GetUnitVatOk() (*float32, bool)`

GetUnitVatOk returns a tuple with the UnitVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitVat

`func (o *MerchantOrderLineResponse) SetUnitVat(v float32)`

SetUnitVat sets UnitVat field to given value.

### HasUnitVat

`func (o *MerchantOrderLineResponse) HasUnitVat() bool`

HasUnitVat returns a boolean if a field has been set.

### SetUnitVatNil

`func (o *MerchantOrderLineResponse) SetUnitVatNil(b bool)`

 SetUnitVatNil sets the value for UnitVat to be an explicit nil

### UnsetUnitVat
`func (o *MerchantOrderLineResponse) UnsetUnitVat()`

UnsetUnitVat ensures that no value is present for UnitVat, not even an explicit nil
### GetLineTotalInclVat

`func (o *MerchantOrderLineResponse) GetLineTotalInclVat() float32`

GetLineTotalInclVat returns the LineTotalInclVat field if non-nil, zero value otherwise.

### GetLineTotalInclVatOk

`func (o *MerchantOrderLineResponse) GetLineTotalInclVatOk() (*float32, bool)`

GetLineTotalInclVatOk returns a tuple with the LineTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalInclVat

`func (o *MerchantOrderLineResponse) SetLineTotalInclVat(v float32)`

SetLineTotalInclVat sets LineTotalInclVat field to given value.

### HasLineTotalInclVat

`func (o *MerchantOrderLineResponse) HasLineTotalInclVat() bool`

HasLineTotalInclVat returns a boolean if a field has been set.

### SetLineTotalInclVatNil

`func (o *MerchantOrderLineResponse) SetLineTotalInclVatNil(b bool)`

 SetLineTotalInclVatNil sets the value for LineTotalInclVat to be an explicit nil

### UnsetLineTotalInclVat
`func (o *MerchantOrderLineResponse) UnsetLineTotalInclVat()`

UnsetLineTotalInclVat ensures that no value is present for LineTotalInclVat, not even an explicit nil
### GetLineVat

`func (o *MerchantOrderLineResponse) GetLineVat() float32`

GetLineVat returns the LineVat field if non-nil, zero value otherwise.

### GetLineVatOk

`func (o *MerchantOrderLineResponse) GetLineVatOk() (*float32, bool)`

GetLineVatOk returns a tuple with the LineVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineVat

`func (o *MerchantOrderLineResponse) SetLineVat(v float32)`

SetLineVat sets LineVat field to given value.

### HasLineVat

`func (o *MerchantOrderLineResponse) HasLineVat() bool`

HasLineVat returns a boolean if a field has been set.

### SetLineVatNil

`func (o *MerchantOrderLineResponse) SetLineVatNil(b bool)`

 SetLineVatNil sets the value for LineVat to be an explicit nil

### UnsetLineVat
`func (o *MerchantOrderLineResponse) UnsetLineVat()`

UnsetLineVat ensures that no value is present for LineVat, not even an explicit nil
### GetOriginalUnitPriceInclVat

`func (o *MerchantOrderLineResponse) GetOriginalUnitPriceInclVat() float32`

GetOriginalUnitPriceInclVat returns the OriginalUnitPriceInclVat field if non-nil, zero value otherwise.

### GetOriginalUnitPriceInclVatOk

`func (o *MerchantOrderLineResponse) GetOriginalUnitPriceInclVatOk() (*float32, bool)`

GetOriginalUnitPriceInclVatOk returns a tuple with the OriginalUnitPriceInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitPriceInclVat

`func (o *MerchantOrderLineResponse) SetOriginalUnitPriceInclVat(v float32)`

SetOriginalUnitPriceInclVat sets OriginalUnitPriceInclVat field to given value.

### HasOriginalUnitPriceInclVat

`func (o *MerchantOrderLineResponse) HasOriginalUnitPriceInclVat() bool`

HasOriginalUnitPriceInclVat returns a boolean if a field has been set.

### SetOriginalUnitPriceInclVatNil

`func (o *MerchantOrderLineResponse) SetOriginalUnitPriceInclVatNil(b bool)`

 SetOriginalUnitPriceInclVatNil sets the value for OriginalUnitPriceInclVat to be an explicit nil

### UnsetOriginalUnitPriceInclVat
`func (o *MerchantOrderLineResponse) UnsetOriginalUnitPriceInclVat()`

UnsetOriginalUnitPriceInclVat ensures that no value is present for OriginalUnitPriceInclVat, not even an explicit nil
### GetOriginalUnitVat

`func (o *MerchantOrderLineResponse) GetOriginalUnitVat() float32`

GetOriginalUnitVat returns the OriginalUnitVat field if non-nil, zero value otherwise.

### GetOriginalUnitVatOk

`func (o *MerchantOrderLineResponse) GetOriginalUnitVatOk() (*float32, bool)`

GetOriginalUnitVatOk returns a tuple with the OriginalUnitVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitVat

`func (o *MerchantOrderLineResponse) SetOriginalUnitVat(v float32)`

SetOriginalUnitVat sets OriginalUnitVat field to given value.

### HasOriginalUnitVat

`func (o *MerchantOrderLineResponse) HasOriginalUnitVat() bool`

HasOriginalUnitVat returns a boolean if a field has been set.

### SetOriginalUnitVatNil

`func (o *MerchantOrderLineResponse) SetOriginalUnitVatNil(b bool)`

 SetOriginalUnitVatNil sets the value for OriginalUnitVat to be an explicit nil

### UnsetOriginalUnitVat
`func (o *MerchantOrderLineResponse) UnsetOriginalUnitVat()`

UnsetOriginalUnitVat ensures that no value is present for OriginalUnitVat, not even an explicit nil
### GetOriginalLineTotalInclVat

`func (o *MerchantOrderLineResponse) GetOriginalLineTotalInclVat() float32`

GetOriginalLineTotalInclVat returns the OriginalLineTotalInclVat field if non-nil, zero value otherwise.

### GetOriginalLineTotalInclVatOk

`func (o *MerchantOrderLineResponse) GetOriginalLineTotalInclVatOk() (*float32, bool)`

GetOriginalLineTotalInclVatOk returns a tuple with the OriginalLineTotalInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineTotalInclVat

`func (o *MerchantOrderLineResponse) SetOriginalLineTotalInclVat(v float32)`

SetOriginalLineTotalInclVat sets OriginalLineTotalInclVat field to given value.

### HasOriginalLineTotalInclVat

`func (o *MerchantOrderLineResponse) HasOriginalLineTotalInclVat() bool`

HasOriginalLineTotalInclVat returns a boolean if a field has been set.

### SetOriginalLineTotalInclVatNil

`func (o *MerchantOrderLineResponse) SetOriginalLineTotalInclVatNil(b bool)`

 SetOriginalLineTotalInclVatNil sets the value for OriginalLineTotalInclVat to be an explicit nil

### UnsetOriginalLineTotalInclVat
`func (o *MerchantOrderLineResponse) UnsetOriginalLineTotalInclVat()`

UnsetOriginalLineTotalInclVat ensures that no value is present for OriginalLineTotalInclVat, not even an explicit nil
### GetOriginalLineVat

`func (o *MerchantOrderLineResponse) GetOriginalLineVat() float32`

GetOriginalLineVat returns the OriginalLineVat field if non-nil, zero value otherwise.

### GetOriginalLineVatOk

`func (o *MerchantOrderLineResponse) GetOriginalLineVatOk() (*float32, bool)`

GetOriginalLineVatOk returns a tuple with the OriginalLineVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineVat

`func (o *MerchantOrderLineResponse) SetOriginalLineVat(v float32)`

SetOriginalLineVat sets OriginalLineVat field to given value.

### HasOriginalLineVat

`func (o *MerchantOrderLineResponse) HasOriginalLineVat() bool`

HasOriginalLineVat returns a boolean if a field has been set.

### SetOriginalLineVatNil

`func (o *MerchantOrderLineResponse) SetOriginalLineVatNil(b bool)`

 SetOriginalLineVatNil sets the value for OriginalLineVat to be an explicit nil

### UnsetOriginalLineVat
`func (o *MerchantOrderLineResponse) UnsetOriginalLineVat()`

UnsetOriginalLineVat ensures that no value is present for OriginalLineVat, not even an explicit nil
### GetOriginalFeeFixed

`func (o *MerchantOrderLineResponse) GetOriginalFeeFixed() float32`

GetOriginalFeeFixed returns the OriginalFeeFixed field if non-nil, zero value otherwise.

### GetOriginalFeeFixedOk

`func (o *MerchantOrderLineResponse) GetOriginalFeeFixedOk() (*float32, bool)`

GetOriginalFeeFixedOk returns a tuple with the OriginalFeeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFeeFixed

`func (o *MerchantOrderLineResponse) SetOriginalFeeFixed(v float32)`

SetOriginalFeeFixed sets OriginalFeeFixed field to given value.

### HasOriginalFeeFixed

`func (o *MerchantOrderLineResponse) HasOriginalFeeFixed() bool`

HasOriginalFeeFixed returns a boolean if a field has been set.

### GetBundleProductMerchantProductNo

`func (o *MerchantOrderLineResponse) GetBundleProductMerchantProductNo() string`

GetBundleProductMerchantProductNo returns the BundleProductMerchantProductNo field if non-nil, zero value otherwise.

### GetBundleProductMerchantProductNoOk

`func (o *MerchantOrderLineResponse) GetBundleProductMerchantProductNoOk() (*string, bool)`

GetBundleProductMerchantProductNoOk returns a tuple with the BundleProductMerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleProductMerchantProductNo

`func (o *MerchantOrderLineResponse) SetBundleProductMerchantProductNo(v string)`

SetBundleProductMerchantProductNo sets BundleProductMerchantProductNo field to given value.

### HasBundleProductMerchantProductNo

`func (o *MerchantOrderLineResponse) HasBundleProductMerchantProductNo() bool`

HasBundleProductMerchantProductNo returns a boolean if a field has been set.

### SetBundleProductMerchantProductNoNil

`func (o *MerchantOrderLineResponse) SetBundleProductMerchantProductNoNil(b bool)`

 SetBundleProductMerchantProductNoNil sets the value for BundleProductMerchantProductNo to be an explicit nil

### UnsetBundleProductMerchantProductNo
`func (o *MerchantOrderLineResponse) UnsetBundleProductMerchantProductNo()`

UnsetBundleProductMerchantProductNo ensures that no value is present for BundleProductMerchantProductNo, not even an explicit nil
### GetJurisCode

`func (o *MerchantOrderLineResponse) GetJurisCode() string`

GetJurisCode returns the JurisCode field if non-nil, zero value otherwise.

### GetJurisCodeOk

`func (o *MerchantOrderLineResponse) GetJurisCodeOk() (*string, bool)`

GetJurisCodeOk returns a tuple with the JurisCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisCode

`func (o *MerchantOrderLineResponse) SetJurisCode(v string)`

SetJurisCode sets JurisCode field to given value.

### HasJurisCode

`func (o *MerchantOrderLineResponse) HasJurisCode() bool`

HasJurisCode returns a boolean if a field has been set.

### SetJurisCodeNil

`func (o *MerchantOrderLineResponse) SetJurisCodeNil(b bool)`

 SetJurisCodeNil sets the value for JurisCode to be an explicit nil

### UnsetJurisCode
`func (o *MerchantOrderLineResponse) UnsetJurisCode()`

UnsetJurisCode ensures that no value is present for JurisCode, not even an explicit nil
### GetJurisName

`func (o *MerchantOrderLineResponse) GetJurisName() string`

GetJurisName returns the JurisName field if non-nil, zero value otherwise.

### GetJurisNameOk

`func (o *MerchantOrderLineResponse) GetJurisNameOk() (*string, bool)`

GetJurisNameOk returns a tuple with the JurisName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisName

`func (o *MerchantOrderLineResponse) SetJurisName(v string)`

SetJurisName sets JurisName field to given value.

### HasJurisName

`func (o *MerchantOrderLineResponse) HasJurisName() bool`

HasJurisName returns a boolean if a field has been set.

### SetJurisNameNil

`func (o *MerchantOrderLineResponse) SetJurisNameNil(b bool)`

 SetJurisNameNil sets the value for JurisName to be an explicit nil

### UnsetJurisName
`func (o *MerchantOrderLineResponse) UnsetJurisName()`

UnsetJurisName ensures that no value is present for JurisName, not even an explicit nil
### GetVatRate

`func (o *MerchantOrderLineResponse) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *MerchantOrderLineResponse) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *MerchantOrderLineResponse) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *MerchantOrderLineResponse) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetUnitPriceExclVat

`func (o *MerchantOrderLineResponse) GetUnitPriceExclVat() float32`

GetUnitPriceExclVat returns the UnitPriceExclVat field if non-nil, zero value otherwise.

### GetUnitPriceExclVatOk

`func (o *MerchantOrderLineResponse) GetUnitPriceExclVatOk() (*float32, bool)`

GetUnitPriceExclVatOk returns a tuple with the UnitPriceExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceExclVat

`func (o *MerchantOrderLineResponse) SetUnitPriceExclVat(v float32)`

SetUnitPriceExclVat sets UnitPriceExclVat field to given value.

### HasUnitPriceExclVat

`func (o *MerchantOrderLineResponse) HasUnitPriceExclVat() bool`

HasUnitPriceExclVat returns a boolean if a field has been set.

### SetUnitPriceExclVatNil

`func (o *MerchantOrderLineResponse) SetUnitPriceExclVatNil(b bool)`

 SetUnitPriceExclVatNil sets the value for UnitPriceExclVat to be an explicit nil

### UnsetUnitPriceExclVat
`func (o *MerchantOrderLineResponse) UnsetUnitPriceExclVat()`

UnsetUnitPriceExclVat ensures that no value is present for UnitPriceExclVat, not even an explicit nil
### GetLineTotalExclVat

`func (o *MerchantOrderLineResponse) GetLineTotalExclVat() float32`

GetLineTotalExclVat returns the LineTotalExclVat field if non-nil, zero value otherwise.

### GetLineTotalExclVatOk

`func (o *MerchantOrderLineResponse) GetLineTotalExclVatOk() (*float32, bool)`

GetLineTotalExclVatOk returns a tuple with the LineTotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalExclVat

`func (o *MerchantOrderLineResponse) SetLineTotalExclVat(v float32)`

SetLineTotalExclVat sets LineTotalExclVat field to given value.

### HasLineTotalExclVat

`func (o *MerchantOrderLineResponse) HasLineTotalExclVat() bool`

HasLineTotalExclVat returns a boolean if a field has been set.

### SetLineTotalExclVatNil

`func (o *MerchantOrderLineResponse) SetLineTotalExclVatNil(b bool)`

 SetLineTotalExclVatNil sets the value for LineTotalExclVat to be an explicit nil

### UnsetLineTotalExclVat
`func (o *MerchantOrderLineResponse) UnsetLineTotalExclVat()`

UnsetLineTotalExclVat ensures that no value is present for LineTotalExclVat, not even an explicit nil
### GetOriginalUnitPriceExclVat

`func (o *MerchantOrderLineResponse) GetOriginalUnitPriceExclVat() float32`

GetOriginalUnitPriceExclVat returns the OriginalUnitPriceExclVat field if non-nil, zero value otherwise.

### GetOriginalUnitPriceExclVatOk

`func (o *MerchantOrderLineResponse) GetOriginalUnitPriceExclVatOk() (*float32, bool)`

GetOriginalUnitPriceExclVatOk returns a tuple with the OriginalUnitPriceExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnitPriceExclVat

`func (o *MerchantOrderLineResponse) SetOriginalUnitPriceExclVat(v float32)`

SetOriginalUnitPriceExclVat sets OriginalUnitPriceExclVat field to given value.

### HasOriginalUnitPriceExclVat

`func (o *MerchantOrderLineResponse) HasOriginalUnitPriceExclVat() bool`

HasOriginalUnitPriceExclVat returns a boolean if a field has been set.

### SetOriginalUnitPriceExclVatNil

`func (o *MerchantOrderLineResponse) SetOriginalUnitPriceExclVatNil(b bool)`

 SetOriginalUnitPriceExclVatNil sets the value for OriginalUnitPriceExclVat to be an explicit nil

### UnsetOriginalUnitPriceExclVat
`func (o *MerchantOrderLineResponse) UnsetOriginalUnitPriceExclVat()`

UnsetOriginalUnitPriceExclVat ensures that no value is present for OriginalUnitPriceExclVat, not even an explicit nil
### GetOriginalLineTotalExclVat

`func (o *MerchantOrderLineResponse) GetOriginalLineTotalExclVat() float32`

GetOriginalLineTotalExclVat returns the OriginalLineTotalExclVat field if non-nil, zero value otherwise.

### GetOriginalLineTotalExclVatOk

`func (o *MerchantOrderLineResponse) GetOriginalLineTotalExclVatOk() (*float32, bool)`

GetOriginalLineTotalExclVatOk returns a tuple with the OriginalLineTotalExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLineTotalExclVat

`func (o *MerchantOrderLineResponse) SetOriginalLineTotalExclVat(v float32)`

SetOriginalLineTotalExclVat sets OriginalLineTotalExclVat field to given value.

### HasOriginalLineTotalExclVat

`func (o *MerchantOrderLineResponse) HasOriginalLineTotalExclVat() bool`

HasOriginalLineTotalExclVat returns a boolean if a field has been set.

### SetOriginalLineTotalExclVatNil

`func (o *MerchantOrderLineResponse) SetOriginalLineTotalExclVatNil(b bool)`

 SetOriginalLineTotalExclVatNil sets the value for OriginalLineTotalExclVat to be an explicit nil

### UnsetOriginalLineTotalExclVat
`func (o *MerchantOrderLineResponse) UnsetOriginalLineTotalExclVat()`

UnsetOriginalLineTotalExclVat ensures that no value is present for OriginalLineTotalExclVat, not even an explicit nil
### GetExtraData

`func (o *MerchantOrderLineResponse) GetExtraData() []MerchantOrderLineExtraDataResponse`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantOrderLineResponse) GetExtraDataOk() (*[]MerchantOrderLineExtraDataResponse, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantOrderLineResponse) SetExtraData(v []MerchantOrderLineExtraDataResponse)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantOrderLineResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantOrderLineResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantOrderLineResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetChannelProductNo

`func (o *MerchantOrderLineResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *MerchantOrderLineResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *MerchantOrderLineResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.


### GetMerchantProductNo

`func (o *MerchantOrderLineResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantOrderLineResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantOrderLineResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantOrderLineResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantOrderLineResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantOrderLineResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetQuantity

`func (o *MerchantOrderLineResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantOrderLineResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantOrderLineResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCancellationRequestedQuantity

`func (o *MerchantOrderLineResponse) GetCancellationRequestedQuantity() int32`

GetCancellationRequestedQuantity returns the CancellationRequestedQuantity field if non-nil, zero value otherwise.

### GetCancellationRequestedQuantityOk

`func (o *MerchantOrderLineResponse) GetCancellationRequestedQuantityOk() (*int32, bool)`

GetCancellationRequestedQuantityOk returns a tuple with the CancellationRequestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequestedQuantity

`func (o *MerchantOrderLineResponse) SetCancellationRequestedQuantity(v int32)`

SetCancellationRequestedQuantity sets CancellationRequestedQuantity field to given value.

### HasCancellationRequestedQuantity

`func (o *MerchantOrderLineResponse) HasCancellationRequestedQuantity() bool`

HasCancellationRequestedQuantity returns a boolean if a field has been set.

### GetUnitPriceInclVat

`func (o *MerchantOrderLineResponse) GetUnitPriceInclVat() float32`

GetUnitPriceInclVat returns the UnitPriceInclVat field if non-nil, zero value otherwise.

### GetUnitPriceInclVatOk

`func (o *MerchantOrderLineResponse) GetUnitPriceInclVatOk() (*float32, bool)`

GetUnitPriceInclVatOk returns a tuple with the UnitPriceInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceInclVat

`func (o *MerchantOrderLineResponse) SetUnitPriceInclVat(v float32)`

SetUnitPriceInclVat sets UnitPriceInclVat field to given value.


### GetFeeFixed

`func (o *MerchantOrderLineResponse) GetFeeFixed() float32`

GetFeeFixed returns the FeeFixed field if non-nil, zero value otherwise.

### GetFeeFixedOk

`func (o *MerchantOrderLineResponse) GetFeeFixedOk() (*float32, bool)`

GetFeeFixedOk returns a tuple with the FeeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeFixed

`func (o *MerchantOrderLineResponse) SetFeeFixed(v float32)`

SetFeeFixed sets FeeFixed field to given value.

### HasFeeFixed

`func (o *MerchantOrderLineResponse) HasFeeFixed() bool`

HasFeeFixed returns a boolean if a field has been set.

### GetFeeRate

`func (o *MerchantOrderLineResponse) GetFeeRate() float32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *MerchantOrderLineResponse) GetFeeRateOk() (*float32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *MerchantOrderLineResponse) SetFeeRate(v float32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *MerchantOrderLineResponse) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetCondition

`func (o *MerchantOrderLineResponse) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MerchantOrderLineResponse) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MerchantOrderLineResponse) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MerchantOrderLineResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *MerchantOrderLineResponse) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *MerchantOrderLineResponse) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *MerchantOrderLineResponse) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *MerchantOrderLineResponse) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### SetExpectedDeliveryDateNil

`func (o *MerchantOrderLineResponse) SetExpectedDeliveryDateNil(b bool)`

 SetExpectedDeliveryDateNil sets the value for ExpectedDeliveryDate to be an explicit nil

### UnsetExpectedDeliveryDate
`func (o *MerchantOrderLineResponse) UnsetExpectedDeliveryDate()`

UnsetExpectedDeliveryDate ensures that no value is present for ExpectedDeliveryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


