# ChannelProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | An unique identifier which ChannelEngine uses to identify the product.  Needed in the call &#39;POST /v2/products/data&#39;. | [optional] 
**ChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel for the product. | [optional] 
**ParentChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel for the parent product. | [optional] 
**GrandparentChannelProductNo** | Pointer to **NullableString** | The unique product reference used by the Channel for the grandparent product. | [optional] 
**MappedFields** | Pointer to **map[string]string** | A channel can require certain fields to be available. The channel  can provide ChannelEngine with this fields after which the merchants  will be able to fill in this field using custom conditions in ChannelEngine. | [optional] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**ExtraData** | Pointer to [**[]ChannelProductExtraDataItemResponse**](ChannelProductExtraDataItemResponse.md) | An optional list of key-value pairs containing  extra data about this product. This data can be  sent to channels or used for filtering products. | [optional] 
**Name** | Pointer to **NullableString** | The name of the product. | [optional] 
**Description** | Pointer to **NullableString** | A description of the product. Can contain these HTML tags:  div, span, pre, p, br, hr, hgroup, h1, h2, h3, h4, h5, h6, ul, ol, li, dl, dt, dd, strong, em, b, i, u, img, a, abbr, address, blockquote, area, audio, video, caption, table, tbody, td, tfoot, th, thead, tr. | [optional] 
**Brand** | Pointer to **NullableString** | The brand of the product. | [optional] 
**Size** | Pointer to **NullableString** | Optional. The size of the product (variant). E.g. fashion size (S-XL, 46-56, etc), width of the watch, etc.. | [optional] 
**Color** | Pointer to **NullableString** | Optional. The color of the product (variant). | [optional] 
**Ean** | Pointer to **NullableString** | The EAN of GTIN of the product. | [optional] 
**ManufacturerProductNumber** | Pointer to **NullableString** | The unique product reference used by the manufacturer/vendor of the product. | [optional] 
**MerchantProductNo** | **string** | A unique identifier of the product. (sku). | 
**Stock** | Pointer to **int32** | The number of items in stock. | [optional] 
**Price** | Pointer to **float32** | Price, including VAT. | [optional] 
**MinPrice** | Pointer to **NullableFloat32** | Min price, including VAT. | [optional] 
**MaxPrice** | Pointer to **NullableFloat32** | Max price, including VAT. | [optional] 
**MSRP** | Pointer to **NullableFloat32** | Manufacturer&#39;s suggested retail price. | [optional] 
**PurchasePrice** | Pointer to **NullableFloat32** | Optional. The purchase price of the product. Useful for repricing. | [optional] 
**VatRateType** | Pointer to [**VatRateType**](VatRateType.md) |  | [optional] 
**ShippingCost** | Pointer to **NullableFloat32** | Shipping cost of the product. | [optional] 
**ShippingTime** | Pointer to **NullableString** | A textual representation of the shippingtime.  For example, in Dutch: &#39;Op werkdagen voor 22:00 uur besteld, morgen in huis&#39;. | [optional] 
**Url** | Pointer to **NullableString** | A URL pointing to the merchant&#39;s webpage  which displays this product. | [optional] 
**ImageUrl** | Pointer to **NullableString** | A URL at which an image of this product  can be found. | [optional] 
**ExtraImageUrl1** | Pointer to **NullableString** | Url to an additional image of product (1). | [optional] 
**ExtraImageUrl2** | Pointer to **NullableString** | Url to an additional image of product (2). | [optional] 
**ExtraImageUrl3** | Pointer to **NullableString** | Url to an additional image of product (3). | [optional] 
**ExtraImageUrl4** | Pointer to **NullableString** | Url to an additional image of product (4). | [optional] 
**ExtraImageUrl5** | Pointer to **NullableString** | Url to an additional image of product (5). | [optional] 
**ExtraImageUrl6** | Pointer to **NullableString** | Url to an additional image of product (6). | [optional] 
**ExtraImageUrl7** | Pointer to **NullableString** | Url to an additional image of product (7). | [optional] 
**ExtraImageUrl8** | Pointer to **NullableString** | Url to an additional image of product (8). | [optional] 
**ExtraImageUrl9** | Pointer to **NullableString** | Url to an additional image of product (9). | [optional] 
**IsFrozen** | Pointer to **bool** | Specifies whether Product is disabled on all channels. | [optional] 
**CategoryTrail** | Pointer to **NullableString** | The category to which this product belongs.  Please supply this field in the following format:  &#39;maincategory &gt; category &gt; subcategory&#39;  For example:  &#39;vehicles &gt; bikes &gt; mountainbike&#39;. | [optional] 

## Methods

### NewChannelProductResponse

`func NewChannelProductResponse(merchantProductNo string, ) *ChannelProductResponse`

NewChannelProductResponse instantiates a new ChannelProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProductResponseWithDefaults

`func NewChannelProductResponseWithDefaults() *ChannelProductResponse`

NewChannelProductResponseWithDefaults instantiates a new ChannelProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelProductResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelProductResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelProductResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelProductResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelProductNo

`func (o *ChannelProductResponse) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelProductResponse) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelProductResponse) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.

### HasChannelProductNo

`func (o *ChannelProductResponse) HasChannelProductNo() bool`

HasChannelProductNo returns a boolean if a field has been set.

### SetChannelProductNoNil

`func (o *ChannelProductResponse) SetChannelProductNoNil(b bool)`

 SetChannelProductNoNil sets the value for ChannelProductNo to be an explicit nil

### UnsetChannelProductNo
`func (o *ChannelProductResponse) UnsetChannelProductNo()`

UnsetChannelProductNo ensures that no value is present for ChannelProductNo, not even an explicit nil
### GetParentChannelProductNo

`func (o *ChannelProductResponse) GetParentChannelProductNo() string`

GetParentChannelProductNo returns the ParentChannelProductNo field if non-nil, zero value otherwise.

### GetParentChannelProductNoOk

`func (o *ChannelProductResponse) GetParentChannelProductNoOk() (*string, bool)`

GetParentChannelProductNoOk returns a tuple with the ParentChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChannelProductNo

`func (o *ChannelProductResponse) SetParentChannelProductNo(v string)`

SetParentChannelProductNo sets ParentChannelProductNo field to given value.

### HasParentChannelProductNo

`func (o *ChannelProductResponse) HasParentChannelProductNo() bool`

HasParentChannelProductNo returns a boolean if a field has been set.

### SetParentChannelProductNoNil

`func (o *ChannelProductResponse) SetParentChannelProductNoNil(b bool)`

 SetParentChannelProductNoNil sets the value for ParentChannelProductNo to be an explicit nil

### UnsetParentChannelProductNo
`func (o *ChannelProductResponse) UnsetParentChannelProductNo()`

UnsetParentChannelProductNo ensures that no value is present for ParentChannelProductNo, not even an explicit nil
### GetGrandparentChannelProductNo

`func (o *ChannelProductResponse) GetGrandparentChannelProductNo() string`

GetGrandparentChannelProductNo returns the GrandparentChannelProductNo field if non-nil, zero value otherwise.

### GetGrandparentChannelProductNoOk

`func (o *ChannelProductResponse) GetGrandparentChannelProductNoOk() (*string, bool)`

GetGrandparentChannelProductNoOk returns a tuple with the GrandparentChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandparentChannelProductNo

`func (o *ChannelProductResponse) SetGrandparentChannelProductNo(v string)`

SetGrandparentChannelProductNo sets GrandparentChannelProductNo field to given value.

### HasGrandparentChannelProductNo

`func (o *ChannelProductResponse) HasGrandparentChannelProductNo() bool`

HasGrandparentChannelProductNo returns a boolean if a field has been set.

### SetGrandparentChannelProductNoNil

`func (o *ChannelProductResponse) SetGrandparentChannelProductNoNil(b bool)`

 SetGrandparentChannelProductNoNil sets the value for GrandparentChannelProductNo to be an explicit nil

### UnsetGrandparentChannelProductNo
`func (o *ChannelProductResponse) UnsetGrandparentChannelProductNo()`

UnsetGrandparentChannelProductNo ensures that no value is present for GrandparentChannelProductNo, not even an explicit nil
### GetMappedFields

`func (o *ChannelProductResponse) GetMappedFields() map[string]string`

GetMappedFields returns the MappedFields field if non-nil, zero value otherwise.

### GetMappedFieldsOk

`func (o *ChannelProductResponse) GetMappedFieldsOk() (*map[string]string, bool)`

GetMappedFieldsOk returns a tuple with the MappedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedFields

`func (o *ChannelProductResponse) SetMappedFields(v map[string]string)`

SetMappedFields sets MappedFields field to given value.

### HasMappedFields

`func (o *ChannelProductResponse) HasMappedFields() bool`

HasMappedFields returns a boolean if a field has been set.

### SetMappedFieldsNil

`func (o *ChannelProductResponse) SetMappedFieldsNil(b bool)`

 SetMappedFieldsNil sets the value for MappedFields to be an explicit nil

### UnsetMappedFields
`func (o *ChannelProductResponse) UnsetMappedFields()`

UnsetMappedFields ensures that no value is present for MappedFields, not even an explicit nil
### GetProductType

`func (o *ChannelProductResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ChannelProductResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ChannelProductResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ChannelProductResponse) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetExtraData

`func (o *ChannelProductResponse) GetExtraData() []ChannelProductExtraDataItemResponse`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelProductResponse) GetExtraDataOk() (*[]ChannelProductExtraDataItemResponse, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelProductResponse) SetExtraData(v []ChannelProductExtraDataItemResponse)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelProductResponse) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelProductResponse) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelProductResponse) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetName

`func (o *ChannelProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelProductResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelProductResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChannelProductResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChannelProductResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ChannelProductResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelProductResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelProductResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelProductResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ChannelProductResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ChannelProductResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBrand

`func (o *ChannelProductResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ChannelProductResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ChannelProductResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ChannelProductResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ChannelProductResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ChannelProductResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetSize

`func (o *ChannelProductResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ChannelProductResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ChannelProductResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ChannelProductResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ChannelProductResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ChannelProductResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetColor

`func (o *ChannelProductResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChannelProductResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChannelProductResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ChannelProductResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ChannelProductResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ChannelProductResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetEan

`func (o *ChannelProductResponse) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ChannelProductResponse) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ChannelProductResponse) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ChannelProductResponse) HasEan() bool`

HasEan returns a boolean if a field has been set.

### SetEanNil

`func (o *ChannelProductResponse) SetEanNil(b bool)`

 SetEanNil sets the value for Ean to be an explicit nil

### UnsetEan
`func (o *ChannelProductResponse) UnsetEan()`

UnsetEan ensures that no value is present for Ean, not even an explicit nil
### GetManufacturerProductNumber

`func (o *ChannelProductResponse) GetManufacturerProductNumber() string`

GetManufacturerProductNumber returns the ManufacturerProductNumber field if non-nil, zero value otherwise.

### GetManufacturerProductNumberOk

`func (o *ChannelProductResponse) GetManufacturerProductNumberOk() (*string, bool)`

GetManufacturerProductNumberOk returns a tuple with the ManufacturerProductNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerProductNumber

`func (o *ChannelProductResponse) SetManufacturerProductNumber(v string)`

SetManufacturerProductNumber sets ManufacturerProductNumber field to given value.

### HasManufacturerProductNumber

`func (o *ChannelProductResponse) HasManufacturerProductNumber() bool`

HasManufacturerProductNumber returns a boolean if a field has been set.

### SetManufacturerProductNumberNil

`func (o *ChannelProductResponse) SetManufacturerProductNumberNil(b bool)`

 SetManufacturerProductNumberNil sets the value for ManufacturerProductNumber to be an explicit nil

### UnsetManufacturerProductNumber
`func (o *ChannelProductResponse) UnsetManufacturerProductNumber()`

UnsetManufacturerProductNumber ensures that no value is present for ManufacturerProductNumber, not even an explicit nil
### GetMerchantProductNo

`func (o *ChannelProductResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *ChannelProductResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *ChannelProductResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetStock

`func (o *ChannelProductResponse) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ChannelProductResponse) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ChannelProductResponse) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ChannelProductResponse) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetPrice

`func (o *ChannelProductResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ChannelProductResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ChannelProductResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ChannelProductResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMinPrice

`func (o *ChannelProductResponse) GetMinPrice() float32`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *ChannelProductResponse) GetMinPriceOk() (*float32, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *ChannelProductResponse) SetMinPrice(v float32)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *ChannelProductResponse) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### SetMinPriceNil

`func (o *ChannelProductResponse) SetMinPriceNil(b bool)`

 SetMinPriceNil sets the value for MinPrice to be an explicit nil

### UnsetMinPrice
`func (o *ChannelProductResponse) UnsetMinPrice()`

UnsetMinPrice ensures that no value is present for MinPrice, not even an explicit nil
### GetMaxPrice

`func (o *ChannelProductResponse) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *ChannelProductResponse) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *ChannelProductResponse) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *ChannelProductResponse) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### SetMaxPriceNil

`func (o *ChannelProductResponse) SetMaxPriceNil(b bool)`

 SetMaxPriceNil sets the value for MaxPrice to be an explicit nil

### UnsetMaxPrice
`func (o *ChannelProductResponse) UnsetMaxPrice()`

UnsetMaxPrice ensures that no value is present for MaxPrice, not even an explicit nil
### GetMSRP

`func (o *ChannelProductResponse) GetMSRP() float32`

GetMSRP returns the MSRP field if non-nil, zero value otherwise.

### GetMSRPOk

`func (o *ChannelProductResponse) GetMSRPOk() (*float32, bool)`

GetMSRPOk returns a tuple with the MSRP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMSRP

`func (o *ChannelProductResponse) SetMSRP(v float32)`

SetMSRP sets MSRP field to given value.

### HasMSRP

`func (o *ChannelProductResponse) HasMSRP() bool`

HasMSRP returns a boolean if a field has been set.

### SetMSRPNil

`func (o *ChannelProductResponse) SetMSRPNil(b bool)`

 SetMSRPNil sets the value for MSRP to be an explicit nil

### UnsetMSRP
`func (o *ChannelProductResponse) UnsetMSRP()`

UnsetMSRP ensures that no value is present for MSRP, not even an explicit nil
### GetPurchasePrice

`func (o *ChannelProductResponse) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *ChannelProductResponse) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *ChannelProductResponse) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *ChannelProductResponse) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *ChannelProductResponse) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *ChannelProductResponse) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetVatRateType

`func (o *ChannelProductResponse) GetVatRateType() VatRateType`

GetVatRateType returns the VatRateType field if non-nil, zero value otherwise.

### GetVatRateTypeOk

`func (o *ChannelProductResponse) GetVatRateTypeOk() (*VatRateType, bool)`

GetVatRateTypeOk returns a tuple with the VatRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRateType

`func (o *ChannelProductResponse) SetVatRateType(v VatRateType)`

SetVatRateType sets VatRateType field to given value.

### HasVatRateType

`func (o *ChannelProductResponse) HasVatRateType() bool`

HasVatRateType returns a boolean if a field has been set.

### GetShippingCost

`func (o *ChannelProductResponse) GetShippingCost() float32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ChannelProductResponse) GetShippingCostOk() (*float32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ChannelProductResponse) SetShippingCost(v float32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ChannelProductResponse) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### SetShippingCostNil

`func (o *ChannelProductResponse) SetShippingCostNil(b bool)`

 SetShippingCostNil sets the value for ShippingCost to be an explicit nil

### UnsetShippingCost
`func (o *ChannelProductResponse) UnsetShippingCost()`

UnsetShippingCost ensures that no value is present for ShippingCost, not even an explicit nil
### GetShippingTime

`func (o *ChannelProductResponse) GetShippingTime() string`

GetShippingTime returns the ShippingTime field if non-nil, zero value otherwise.

### GetShippingTimeOk

`func (o *ChannelProductResponse) GetShippingTimeOk() (*string, bool)`

GetShippingTimeOk returns a tuple with the ShippingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTime

`func (o *ChannelProductResponse) SetShippingTime(v string)`

SetShippingTime sets ShippingTime field to given value.

### HasShippingTime

`func (o *ChannelProductResponse) HasShippingTime() bool`

HasShippingTime returns a boolean if a field has been set.

### SetShippingTimeNil

`func (o *ChannelProductResponse) SetShippingTimeNil(b bool)`

 SetShippingTimeNil sets the value for ShippingTime to be an explicit nil

### UnsetShippingTime
`func (o *ChannelProductResponse) UnsetShippingTime()`

UnsetShippingTime ensures that no value is present for ShippingTime, not even an explicit nil
### GetUrl

`func (o *ChannelProductResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelProductResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelProductResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChannelProductResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChannelProductResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChannelProductResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageUrl

`func (o *ChannelProductResponse) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ChannelProductResponse) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ChannelProductResponse) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ChannelProductResponse) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ChannelProductResponse) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ChannelProductResponse) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExtraImageUrl1

`func (o *ChannelProductResponse) GetExtraImageUrl1() string`

GetExtraImageUrl1 returns the ExtraImageUrl1 field if non-nil, zero value otherwise.

### GetExtraImageUrl1Ok

`func (o *ChannelProductResponse) GetExtraImageUrl1Ok() (*string, bool)`

GetExtraImageUrl1Ok returns a tuple with the ExtraImageUrl1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl1

`func (o *ChannelProductResponse) SetExtraImageUrl1(v string)`

SetExtraImageUrl1 sets ExtraImageUrl1 field to given value.

### HasExtraImageUrl1

`func (o *ChannelProductResponse) HasExtraImageUrl1() bool`

HasExtraImageUrl1 returns a boolean if a field has been set.

### SetExtraImageUrl1Nil

`func (o *ChannelProductResponse) SetExtraImageUrl1Nil(b bool)`

 SetExtraImageUrl1Nil sets the value for ExtraImageUrl1 to be an explicit nil

### UnsetExtraImageUrl1
`func (o *ChannelProductResponse) UnsetExtraImageUrl1()`

UnsetExtraImageUrl1 ensures that no value is present for ExtraImageUrl1, not even an explicit nil
### GetExtraImageUrl2

`func (o *ChannelProductResponse) GetExtraImageUrl2() string`

GetExtraImageUrl2 returns the ExtraImageUrl2 field if non-nil, zero value otherwise.

### GetExtraImageUrl2Ok

`func (o *ChannelProductResponse) GetExtraImageUrl2Ok() (*string, bool)`

GetExtraImageUrl2Ok returns a tuple with the ExtraImageUrl2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl2

`func (o *ChannelProductResponse) SetExtraImageUrl2(v string)`

SetExtraImageUrl2 sets ExtraImageUrl2 field to given value.

### HasExtraImageUrl2

`func (o *ChannelProductResponse) HasExtraImageUrl2() bool`

HasExtraImageUrl2 returns a boolean if a field has been set.

### SetExtraImageUrl2Nil

`func (o *ChannelProductResponse) SetExtraImageUrl2Nil(b bool)`

 SetExtraImageUrl2Nil sets the value for ExtraImageUrl2 to be an explicit nil

### UnsetExtraImageUrl2
`func (o *ChannelProductResponse) UnsetExtraImageUrl2()`

UnsetExtraImageUrl2 ensures that no value is present for ExtraImageUrl2, not even an explicit nil
### GetExtraImageUrl3

`func (o *ChannelProductResponse) GetExtraImageUrl3() string`

GetExtraImageUrl3 returns the ExtraImageUrl3 field if non-nil, zero value otherwise.

### GetExtraImageUrl3Ok

`func (o *ChannelProductResponse) GetExtraImageUrl3Ok() (*string, bool)`

GetExtraImageUrl3Ok returns a tuple with the ExtraImageUrl3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl3

`func (o *ChannelProductResponse) SetExtraImageUrl3(v string)`

SetExtraImageUrl3 sets ExtraImageUrl3 field to given value.

### HasExtraImageUrl3

`func (o *ChannelProductResponse) HasExtraImageUrl3() bool`

HasExtraImageUrl3 returns a boolean if a field has been set.

### SetExtraImageUrl3Nil

`func (o *ChannelProductResponse) SetExtraImageUrl3Nil(b bool)`

 SetExtraImageUrl3Nil sets the value for ExtraImageUrl3 to be an explicit nil

### UnsetExtraImageUrl3
`func (o *ChannelProductResponse) UnsetExtraImageUrl3()`

UnsetExtraImageUrl3 ensures that no value is present for ExtraImageUrl3, not even an explicit nil
### GetExtraImageUrl4

`func (o *ChannelProductResponse) GetExtraImageUrl4() string`

GetExtraImageUrl4 returns the ExtraImageUrl4 field if non-nil, zero value otherwise.

### GetExtraImageUrl4Ok

`func (o *ChannelProductResponse) GetExtraImageUrl4Ok() (*string, bool)`

GetExtraImageUrl4Ok returns a tuple with the ExtraImageUrl4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl4

`func (o *ChannelProductResponse) SetExtraImageUrl4(v string)`

SetExtraImageUrl4 sets ExtraImageUrl4 field to given value.

### HasExtraImageUrl4

`func (o *ChannelProductResponse) HasExtraImageUrl4() bool`

HasExtraImageUrl4 returns a boolean if a field has been set.

### SetExtraImageUrl4Nil

`func (o *ChannelProductResponse) SetExtraImageUrl4Nil(b bool)`

 SetExtraImageUrl4Nil sets the value for ExtraImageUrl4 to be an explicit nil

### UnsetExtraImageUrl4
`func (o *ChannelProductResponse) UnsetExtraImageUrl4()`

UnsetExtraImageUrl4 ensures that no value is present for ExtraImageUrl4, not even an explicit nil
### GetExtraImageUrl5

`func (o *ChannelProductResponse) GetExtraImageUrl5() string`

GetExtraImageUrl5 returns the ExtraImageUrl5 field if non-nil, zero value otherwise.

### GetExtraImageUrl5Ok

`func (o *ChannelProductResponse) GetExtraImageUrl5Ok() (*string, bool)`

GetExtraImageUrl5Ok returns a tuple with the ExtraImageUrl5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl5

`func (o *ChannelProductResponse) SetExtraImageUrl5(v string)`

SetExtraImageUrl5 sets ExtraImageUrl5 field to given value.

### HasExtraImageUrl5

`func (o *ChannelProductResponse) HasExtraImageUrl5() bool`

HasExtraImageUrl5 returns a boolean if a field has been set.

### SetExtraImageUrl5Nil

`func (o *ChannelProductResponse) SetExtraImageUrl5Nil(b bool)`

 SetExtraImageUrl5Nil sets the value for ExtraImageUrl5 to be an explicit nil

### UnsetExtraImageUrl5
`func (o *ChannelProductResponse) UnsetExtraImageUrl5()`

UnsetExtraImageUrl5 ensures that no value is present for ExtraImageUrl5, not even an explicit nil
### GetExtraImageUrl6

`func (o *ChannelProductResponse) GetExtraImageUrl6() string`

GetExtraImageUrl6 returns the ExtraImageUrl6 field if non-nil, zero value otherwise.

### GetExtraImageUrl6Ok

`func (o *ChannelProductResponse) GetExtraImageUrl6Ok() (*string, bool)`

GetExtraImageUrl6Ok returns a tuple with the ExtraImageUrl6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl6

`func (o *ChannelProductResponse) SetExtraImageUrl6(v string)`

SetExtraImageUrl6 sets ExtraImageUrl6 field to given value.

### HasExtraImageUrl6

`func (o *ChannelProductResponse) HasExtraImageUrl6() bool`

HasExtraImageUrl6 returns a boolean if a field has been set.

### SetExtraImageUrl6Nil

`func (o *ChannelProductResponse) SetExtraImageUrl6Nil(b bool)`

 SetExtraImageUrl6Nil sets the value for ExtraImageUrl6 to be an explicit nil

### UnsetExtraImageUrl6
`func (o *ChannelProductResponse) UnsetExtraImageUrl6()`

UnsetExtraImageUrl6 ensures that no value is present for ExtraImageUrl6, not even an explicit nil
### GetExtraImageUrl7

`func (o *ChannelProductResponse) GetExtraImageUrl7() string`

GetExtraImageUrl7 returns the ExtraImageUrl7 field if non-nil, zero value otherwise.

### GetExtraImageUrl7Ok

`func (o *ChannelProductResponse) GetExtraImageUrl7Ok() (*string, bool)`

GetExtraImageUrl7Ok returns a tuple with the ExtraImageUrl7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl7

`func (o *ChannelProductResponse) SetExtraImageUrl7(v string)`

SetExtraImageUrl7 sets ExtraImageUrl7 field to given value.

### HasExtraImageUrl7

`func (o *ChannelProductResponse) HasExtraImageUrl7() bool`

HasExtraImageUrl7 returns a boolean if a field has been set.

### SetExtraImageUrl7Nil

`func (o *ChannelProductResponse) SetExtraImageUrl7Nil(b bool)`

 SetExtraImageUrl7Nil sets the value for ExtraImageUrl7 to be an explicit nil

### UnsetExtraImageUrl7
`func (o *ChannelProductResponse) UnsetExtraImageUrl7()`

UnsetExtraImageUrl7 ensures that no value is present for ExtraImageUrl7, not even an explicit nil
### GetExtraImageUrl8

`func (o *ChannelProductResponse) GetExtraImageUrl8() string`

GetExtraImageUrl8 returns the ExtraImageUrl8 field if non-nil, zero value otherwise.

### GetExtraImageUrl8Ok

`func (o *ChannelProductResponse) GetExtraImageUrl8Ok() (*string, bool)`

GetExtraImageUrl8Ok returns a tuple with the ExtraImageUrl8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl8

`func (o *ChannelProductResponse) SetExtraImageUrl8(v string)`

SetExtraImageUrl8 sets ExtraImageUrl8 field to given value.

### HasExtraImageUrl8

`func (o *ChannelProductResponse) HasExtraImageUrl8() bool`

HasExtraImageUrl8 returns a boolean if a field has been set.

### SetExtraImageUrl8Nil

`func (o *ChannelProductResponse) SetExtraImageUrl8Nil(b bool)`

 SetExtraImageUrl8Nil sets the value for ExtraImageUrl8 to be an explicit nil

### UnsetExtraImageUrl8
`func (o *ChannelProductResponse) UnsetExtraImageUrl8()`

UnsetExtraImageUrl8 ensures that no value is present for ExtraImageUrl8, not even an explicit nil
### GetExtraImageUrl9

`func (o *ChannelProductResponse) GetExtraImageUrl9() string`

GetExtraImageUrl9 returns the ExtraImageUrl9 field if non-nil, zero value otherwise.

### GetExtraImageUrl9Ok

`func (o *ChannelProductResponse) GetExtraImageUrl9Ok() (*string, bool)`

GetExtraImageUrl9Ok returns a tuple with the ExtraImageUrl9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl9

`func (o *ChannelProductResponse) SetExtraImageUrl9(v string)`

SetExtraImageUrl9 sets ExtraImageUrl9 field to given value.

### HasExtraImageUrl9

`func (o *ChannelProductResponse) HasExtraImageUrl9() bool`

HasExtraImageUrl9 returns a boolean if a field has been set.

### SetExtraImageUrl9Nil

`func (o *ChannelProductResponse) SetExtraImageUrl9Nil(b bool)`

 SetExtraImageUrl9Nil sets the value for ExtraImageUrl9 to be an explicit nil

### UnsetExtraImageUrl9
`func (o *ChannelProductResponse) UnsetExtraImageUrl9()`

UnsetExtraImageUrl9 ensures that no value is present for ExtraImageUrl9, not even an explicit nil
### GetIsFrozen

`func (o *ChannelProductResponse) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *ChannelProductResponse) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *ChannelProductResponse) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *ChannelProductResponse) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.

### GetCategoryTrail

`func (o *ChannelProductResponse) GetCategoryTrail() string`

GetCategoryTrail returns the CategoryTrail field if non-nil, zero value otherwise.

### GetCategoryTrailOk

`func (o *ChannelProductResponse) GetCategoryTrailOk() (*string, bool)`

GetCategoryTrailOk returns a tuple with the CategoryTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTrail

`func (o *ChannelProductResponse) SetCategoryTrail(v string)`

SetCategoryTrail sets CategoryTrail field to given value.

### HasCategoryTrail

`func (o *ChannelProductResponse) HasCategoryTrail() bool`

HasCategoryTrail returns a boolean if a field has been set.

### SetCategoryTrailNil

`func (o *ChannelProductResponse) SetCategoryTrailNil(b bool)`

 SetCategoryTrailNil sets the value for CategoryTrail to be an explicit nil

### UnsetCategoryTrail
`func (o *ChannelProductResponse) UnsetCategoryTrail()`

UnsetCategoryTrail ensures that no value is present for CategoryTrail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


