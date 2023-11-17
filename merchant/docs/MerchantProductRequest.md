# MerchantProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentMerchantProductNo** | Pointer to **NullableString** | If this product is a different version of another  product (for example, all fields are the same except  size), then this field should contain  the &#39;MerchantProductNo&#39; of the parent. The parent  should already exist (or be present between the products  in the content of the API call, it does not matter whether  the parent is behind the child in the list). | [optional] 
**ParentMerchantProductNo2** | Pointer to **NullableString** | If this product is a different version of another  product (for example, all fields are the same except  color) and itself is a parent with child products (e.g. of sizes),  then this field should contain the &#39;MerchantProductNo&#39; of the grandparent. The grandparent  should already exist (or be present between the products  in the content of the API call, it does not matter whether  the grandparent is behind the child in the list).  When you set this field, the ParentMerchantProductNo should be left empty.                Use this field in case of three level product hierarchy,  e.g. model - color - size.  This is required for channels like Otto. | [optional] 
**ExtraData** | Pointer to [**[]MerchantProductExtraDataItemRequest**](MerchantProductExtraDataItemRequest.md) | An optional list of key-value pairs containing  extra data about this product. This data can be  sent to channels or used for filtering products. | [optional] 
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
**CategoryTrail** | Pointer to **NullableString** | The category to which this product belongs.  Please supply this field in the following format:  &#39;maincategory &gt; category &gt; subcategory&#39;  For example:  &#39;vehicles &gt; bikes &gt; mountainbike&#39;. | [optional] 

## Methods

### NewMerchantProductRequest

`func NewMerchantProductRequest(merchantProductNo string, ) *MerchantProductRequest`

NewMerchantProductRequest instantiates a new MerchantProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductRequestWithDefaults

`func NewMerchantProductRequestWithDefaults() *MerchantProductRequest`

NewMerchantProductRequestWithDefaults instantiates a new MerchantProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentMerchantProductNo

`func (o *MerchantProductRequest) GetParentMerchantProductNo() string`

GetParentMerchantProductNo returns the ParentMerchantProductNo field if non-nil, zero value otherwise.

### GetParentMerchantProductNoOk

`func (o *MerchantProductRequest) GetParentMerchantProductNoOk() (*string, bool)`

GetParentMerchantProductNoOk returns a tuple with the ParentMerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMerchantProductNo

`func (o *MerchantProductRequest) SetParentMerchantProductNo(v string)`

SetParentMerchantProductNo sets ParentMerchantProductNo field to given value.

### HasParentMerchantProductNo

`func (o *MerchantProductRequest) HasParentMerchantProductNo() bool`

HasParentMerchantProductNo returns a boolean if a field has been set.

### SetParentMerchantProductNoNil

`func (o *MerchantProductRequest) SetParentMerchantProductNoNil(b bool)`

 SetParentMerchantProductNoNil sets the value for ParentMerchantProductNo to be an explicit nil

### UnsetParentMerchantProductNo
`func (o *MerchantProductRequest) UnsetParentMerchantProductNo()`

UnsetParentMerchantProductNo ensures that no value is present for ParentMerchantProductNo, not even an explicit nil
### GetParentMerchantProductNo2

`func (o *MerchantProductRequest) GetParentMerchantProductNo2() string`

GetParentMerchantProductNo2 returns the ParentMerchantProductNo2 field if non-nil, zero value otherwise.

### GetParentMerchantProductNo2Ok

`func (o *MerchantProductRequest) GetParentMerchantProductNo2Ok() (*string, bool)`

GetParentMerchantProductNo2Ok returns a tuple with the ParentMerchantProductNo2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMerchantProductNo2

`func (o *MerchantProductRequest) SetParentMerchantProductNo2(v string)`

SetParentMerchantProductNo2 sets ParentMerchantProductNo2 field to given value.

### HasParentMerchantProductNo2

`func (o *MerchantProductRequest) HasParentMerchantProductNo2() bool`

HasParentMerchantProductNo2 returns a boolean if a field has been set.

### SetParentMerchantProductNo2Nil

`func (o *MerchantProductRequest) SetParentMerchantProductNo2Nil(b bool)`

 SetParentMerchantProductNo2Nil sets the value for ParentMerchantProductNo2 to be an explicit nil

### UnsetParentMerchantProductNo2
`func (o *MerchantProductRequest) UnsetParentMerchantProductNo2()`

UnsetParentMerchantProductNo2 ensures that no value is present for ParentMerchantProductNo2, not even an explicit nil
### GetExtraData

`func (o *MerchantProductRequest) GetExtraData() []MerchantProductExtraDataItemRequest`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantProductRequest) GetExtraDataOk() (*[]MerchantProductExtraDataItemRequest, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantProductRequest) SetExtraData(v []MerchantProductExtraDataItemRequest)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantProductRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantProductRequest) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantProductRequest) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetName

`func (o *MerchantProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantProductRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantProductRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantProductRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantProductRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *MerchantProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MerchantProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MerchantProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MerchantProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MerchantProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MerchantProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBrand

`func (o *MerchantProductRequest) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *MerchantProductRequest) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *MerchantProductRequest) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *MerchantProductRequest) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *MerchantProductRequest) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *MerchantProductRequest) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetSize

`func (o *MerchantProductRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MerchantProductRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MerchantProductRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MerchantProductRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MerchantProductRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MerchantProductRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetColor

`func (o *MerchantProductRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MerchantProductRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MerchantProductRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MerchantProductRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MerchantProductRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MerchantProductRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetEan

`func (o *MerchantProductRequest) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *MerchantProductRequest) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *MerchantProductRequest) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *MerchantProductRequest) HasEan() bool`

HasEan returns a boolean if a field has been set.

### SetEanNil

`func (o *MerchantProductRequest) SetEanNil(b bool)`

 SetEanNil sets the value for Ean to be an explicit nil

### UnsetEan
`func (o *MerchantProductRequest) UnsetEan()`

UnsetEan ensures that no value is present for Ean, not even an explicit nil
### GetManufacturerProductNumber

`func (o *MerchantProductRequest) GetManufacturerProductNumber() string`

GetManufacturerProductNumber returns the ManufacturerProductNumber field if non-nil, zero value otherwise.

### GetManufacturerProductNumberOk

`func (o *MerchantProductRequest) GetManufacturerProductNumberOk() (*string, bool)`

GetManufacturerProductNumberOk returns a tuple with the ManufacturerProductNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerProductNumber

`func (o *MerchantProductRequest) SetManufacturerProductNumber(v string)`

SetManufacturerProductNumber sets ManufacturerProductNumber field to given value.

### HasManufacturerProductNumber

`func (o *MerchantProductRequest) HasManufacturerProductNumber() bool`

HasManufacturerProductNumber returns a boolean if a field has been set.

### SetManufacturerProductNumberNil

`func (o *MerchantProductRequest) SetManufacturerProductNumberNil(b bool)`

 SetManufacturerProductNumberNil sets the value for ManufacturerProductNumber to be an explicit nil

### UnsetManufacturerProductNumber
`func (o *MerchantProductRequest) UnsetManufacturerProductNumber()`

UnsetManufacturerProductNumber ensures that no value is present for ManufacturerProductNumber, not even an explicit nil
### GetMerchantProductNo

`func (o *MerchantProductRequest) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantProductRequest) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantProductRequest) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.


### GetStock

`func (o *MerchantProductRequest) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *MerchantProductRequest) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *MerchantProductRequest) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *MerchantProductRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetPrice

`func (o *MerchantProductRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MerchantProductRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MerchantProductRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MerchantProductRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMinPrice

`func (o *MerchantProductRequest) GetMinPrice() float32`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *MerchantProductRequest) GetMinPriceOk() (*float32, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *MerchantProductRequest) SetMinPrice(v float32)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *MerchantProductRequest) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### SetMinPriceNil

`func (o *MerchantProductRequest) SetMinPriceNil(b bool)`

 SetMinPriceNil sets the value for MinPrice to be an explicit nil

### UnsetMinPrice
`func (o *MerchantProductRequest) UnsetMinPrice()`

UnsetMinPrice ensures that no value is present for MinPrice, not even an explicit nil
### GetMaxPrice

`func (o *MerchantProductRequest) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *MerchantProductRequest) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *MerchantProductRequest) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *MerchantProductRequest) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### SetMaxPriceNil

`func (o *MerchantProductRequest) SetMaxPriceNil(b bool)`

 SetMaxPriceNil sets the value for MaxPrice to be an explicit nil

### UnsetMaxPrice
`func (o *MerchantProductRequest) UnsetMaxPrice()`

UnsetMaxPrice ensures that no value is present for MaxPrice, not even an explicit nil
### GetMSRP

`func (o *MerchantProductRequest) GetMSRP() float32`

GetMSRP returns the MSRP field if non-nil, zero value otherwise.

### GetMSRPOk

`func (o *MerchantProductRequest) GetMSRPOk() (*float32, bool)`

GetMSRPOk returns a tuple with the MSRP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMSRP

`func (o *MerchantProductRequest) SetMSRP(v float32)`

SetMSRP sets MSRP field to given value.

### HasMSRP

`func (o *MerchantProductRequest) HasMSRP() bool`

HasMSRP returns a boolean if a field has been set.

### SetMSRPNil

`func (o *MerchantProductRequest) SetMSRPNil(b bool)`

 SetMSRPNil sets the value for MSRP to be an explicit nil

### UnsetMSRP
`func (o *MerchantProductRequest) UnsetMSRP()`

UnsetMSRP ensures that no value is present for MSRP, not even an explicit nil
### GetPurchasePrice

`func (o *MerchantProductRequest) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *MerchantProductRequest) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *MerchantProductRequest) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *MerchantProductRequest) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *MerchantProductRequest) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *MerchantProductRequest) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetVatRateType

`func (o *MerchantProductRequest) GetVatRateType() VatRateType`

GetVatRateType returns the VatRateType field if non-nil, zero value otherwise.

### GetVatRateTypeOk

`func (o *MerchantProductRequest) GetVatRateTypeOk() (*VatRateType, bool)`

GetVatRateTypeOk returns a tuple with the VatRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRateType

`func (o *MerchantProductRequest) SetVatRateType(v VatRateType)`

SetVatRateType sets VatRateType field to given value.

### HasVatRateType

`func (o *MerchantProductRequest) HasVatRateType() bool`

HasVatRateType returns a boolean if a field has been set.

### GetShippingCost

`func (o *MerchantProductRequest) GetShippingCost() float32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *MerchantProductRequest) GetShippingCostOk() (*float32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *MerchantProductRequest) SetShippingCost(v float32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *MerchantProductRequest) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### SetShippingCostNil

`func (o *MerchantProductRequest) SetShippingCostNil(b bool)`

 SetShippingCostNil sets the value for ShippingCost to be an explicit nil

### UnsetShippingCost
`func (o *MerchantProductRequest) UnsetShippingCost()`

UnsetShippingCost ensures that no value is present for ShippingCost, not even an explicit nil
### GetShippingTime

`func (o *MerchantProductRequest) GetShippingTime() string`

GetShippingTime returns the ShippingTime field if non-nil, zero value otherwise.

### GetShippingTimeOk

`func (o *MerchantProductRequest) GetShippingTimeOk() (*string, bool)`

GetShippingTimeOk returns a tuple with the ShippingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTime

`func (o *MerchantProductRequest) SetShippingTime(v string)`

SetShippingTime sets ShippingTime field to given value.

### HasShippingTime

`func (o *MerchantProductRequest) HasShippingTime() bool`

HasShippingTime returns a boolean if a field has been set.

### SetShippingTimeNil

`func (o *MerchantProductRequest) SetShippingTimeNil(b bool)`

 SetShippingTimeNil sets the value for ShippingTime to be an explicit nil

### UnsetShippingTime
`func (o *MerchantProductRequest) UnsetShippingTime()`

UnsetShippingTime ensures that no value is present for ShippingTime, not even an explicit nil
### GetUrl

`func (o *MerchantProductRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MerchantProductRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MerchantProductRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MerchantProductRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MerchantProductRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MerchantProductRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageUrl

`func (o *MerchantProductRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *MerchantProductRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *MerchantProductRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *MerchantProductRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *MerchantProductRequest) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *MerchantProductRequest) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetExtraImageUrl1

`func (o *MerchantProductRequest) GetExtraImageUrl1() string`

GetExtraImageUrl1 returns the ExtraImageUrl1 field if non-nil, zero value otherwise.

### GetExtraImageUrl1Ok

`func (o *MerchantProductRequest) GetExtraImageUrl1Ok() (*string, bool)`

GetExtraImageUrl1Ok returns a tuple with the ExtraImageUrl1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl1

`func (o *MerchantProductRequest) SetExtraImageUrl1(v string)`

SetExtraImageUrl1 sets ExtraImageUrl1 field to given value.

### HasExtraImageUrl1

`func (o *MerchantProductRequest) HasExtraImageUrl1() bool`

HasExtraImageUrl1 returns a boolean if a field has been set.

### SetExtraImageUrl1Nil

`func (o *MerchantProductRequest) SetExtraImageUrl1Nil(b bool)`

 SetExtraImageUrl1Nil sets the value for ExtraImageUrl1 to be an explicit nil

### UnsetExtraImageUrl1
`func (o *MerchantProductRequest) UnsetExtraImageUrl1()`

UnsetExtraImageUrl1 ensures that no value is present for ExtraImageUrl1, not even an explicit nil
### GetExtraImageUrl2

`func (o *MerchantProductRequest) GetExtraImageUrl2() string`

GetExtraImageUrl2 returns the ExtraImageUrl2 field if non-nil, zero value otherwise.

### GetExtraImageUrl2Ok

`func (o *MerchantProductRequest) GetExtraImageUrl2Ok() (*string, bool)`

GetExtraImageUrl2Ok returns a tuple with the ExtraImageUrl2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl2

`func (o *MerchantProductRequest) SetExtraImageUrl2(v string)`

SetExtraImageUrl2 sets ExtraImageUrl2 field to given value.

### HasExtraImageUrl2

`func (o *MerchantProductRequest) HasExtraImageUrl2() bool`

HasExtraImageUrl2 returns a boolean if a field has been set.

### SetExtraImageUrl2Nil

`func (o *MerchantProductRequest) SetExtraImageUrl2Nil(b bool)`

 SetExtraImageUrl2Nil sets the value for ExtraImageUrl2 to be an explicit nil

### UnsetExtraImageUrl2
`func (o *MerchantProductRequest) UnsetExtraImageUrl2()`

UnsetExtraImageUrl2 ensures that no value is present for ExtraImageUrl2, not even an explicit nil
### GetExtraImageUrl3

`func (o *MerchantProductRequest) GetExtraImageUrl3() string`

GetExtraImageUrl3 returns the ExtraImageUrl3 field if non-nil, zero value otherwise.

### GetExtraImageUrl3Ok

`func (o *MerchantProductRequest) GetExtraImageUrl3Ok() (*string, bool)`

GetExtraImageUrl3Ok returns a tuple with the ExtraImageUrl3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl3

`func (o *MerchantProductRequest) SetExtraImageUrl3(v string)`

SetExtraImageUrl3 sets ExtraImageUrl3 field to given value.

### HasExtraImageUrl3

`func (o *MerchantProductRequest) HasExtraImageUrl3() bool`

HasExtraImageUrl3 returns a boolean if a field has been set.

### SetExtraImageUrl3Nil

`func (o *MerchantProductRequest) SetExtraImageUrl3Nil(b bool)`

 SetExtraImageUrl3Nil sets the value for ExtraImageUrl3 to be an explicit nil

### UnsetExtraImageUrl3
`func (o *MerchantProductRequest) UnsetExtraImageUrl3()`

UnsetExtraImageUrl3 ensures that no value is present for ExtraImageUrl3, not even an explicit nil
### GetExtraImageUrl4

`func (o *MerchantProductRequest) GetExtraImageUrl4() string`

GetExtraImageUrl4 returns the ExtraImageUrl4 field if non-nil, zero value otherwise.

### GetExtraImageUrl4Ok

`func (o *MerchantProductRequest) GetExtraImageUrl4Ok() (*string, bool)`

GetExtraImageUrl4Ok returns a tuple with the ExtraImageUrl4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl4

`func (o *MerchantProductRequest) SetExtraImageUrl4(v string)`

SetExtraImageUrl4 sets ExtraImageUrl4 field to given value.

### HasExtraImageUrl4

`func (o *MerchantProductRequest) HasExtraImageUrl4() bool`

HasExtraImageUrl4 returns a boolean if a field has been set.

### SetExtraImageUrl4Nil

`func (o *MerchantProductRequest) SetExtraImageUrl4Nil(b bool)`

 SetExtraImageUrl4Nil sets the value for ExtraImageUrl4 to be an explicit nil

### UnsetExtraImageUrl4
`func (o *MerchantProductRequest) UnsetExtraImageUrl4()`

UnsetExtraImageUrl4 ensures that no value is present for ExtraImageUrl4, not even an explicit nil
### GetExtraImageUrl5

`func (o *MerchantProductRequest) GetExtraImageUrl5() string`

GetExtraImageUrl5 returns the ExtraImageUrl5 field if non-nil, zero value otherwise.

### GetExtraImageUrl5Ok

`func (o *MerchantProductRequest) GetExtraImageUrl5Ok() (*string, bool)`

GetExtraImageUrl5Ok returns a tuple with the ExtraImageUrl5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl5

`func (o *MerchantProductRequest) SetExtraImageUrl5(v string)`

SetExtraImageUrl5 sets ExtraImageUrl5 field to given value.

### HasExtraImageUrl5

`func (o *MerchantProductRequest) HasExtraImageUrl5() bool`

HasExtraImageUrl5 returns a boolean if a field has been set.

### SetExtraImageUrl5Nil

`func (o *MerchantProductRequest) SetExtraImageUrl5Nil(b bool)`

 SetExtraImageUrl5Nil sets the value for ExtraImageUrl5 to be an explicit nil

### UnsetExtraImageUrl5
`func (o *MerchantProductRequest) UnsetExtraImageUrl5()`

UnsetExtraImageUrl5 ensures that no value is present for ExtraImageUrl5, not even an explicit nil
### GetExtraImageUrl6

`func (o *MerchantProductRequest) GetExtraImageUrl6() string`

GetExtraImageUrl6 returns the ExtraImageUrl6 field if non-nil, zero value otherwise.

### GetExtraImageUrl6Ok

`func (o *MerchantProductRequest) GetExtraImageUrl6Ok() (*string, bool)`

GetExtraImageUrl6Ok returns a tuple with the ExtraImageUrl6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl6

`func (o *MerchantProductRequest) SetExtraImageUrl6(v string)`

SetExtraImageUrl6 sets ExtraImageUrl6 field to given value.

### HasExtraImageUrl6

`func (o *MerchantProductRequest) HasExtraImageUrl6() bool`

HasExtraImageUrl6 returns a boolean if a field has been set.

### SetExtraImageUrl6Nil

`func (o *MerchantProductRequest) SetExtraImageUrl6Nil(b bool)`

 SetExtraImageUrl6Nil sets the value for ExtraImageUrl6 to be an explicit nil

### UnsetExtraImageUrl6
`func (o *MerchantProductRequest) UnsetExtraImageUrl6()`

UnsetExtraImageUrl6 ensures that no value is present for ExtraImageUrl6, not even an explicit nil
### GetExtraImageUrl7

`func (o *MerchantProductRequest) GetExtraImageUrl7() string`

GetExtraImageUrl7 returns the ExtraImageUrl7 field if non-nil, zero value otherwise.

### GetExtraImageUrl7Ok

`func (o *MerchantProductRequest) GetExtraImageUrl7Ok() (*string, bool)`

GetExtraImageUrl7Ok returns a tuple with the ExtraImageUrl7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl7

`func (o *MerchantProductRequest) SetExtraImageUrl7(v string)`

SetExtraImageUrl7 sets ExtraImageUrl7 field to given value.

### HasExtraImageUrl7

`func (o *MerchantProductRequest) HasExtraImageUrl7() bool`

HasExtraImageUrl7 returns a boolean if a field has been set.

### SetExtraImageUrl7Nil

`func (o *MerchantProductRequest) SetExtraImageUrl7Nil(b bool)`

 SetExtraImageUrl7Nil sets the value for ExtraImageUrl7 to be an explicit nil

### UnsetExtraImageUrl7
`func (o *MerchantProductRequest) UnsetExtraImageUrl7()`

UnsetExtraImageUrl7 ensures that no value is present for ExtraImageUrl7, not even an explicit nil
### GetExtraImageUrl8

`func (o *MerchantProductRequest) GetExtraImageUrl8() string`

GetExtraImageUrl8 returns the ExtraImageUrl8 field if non-nil, zero value otherwise.

### GetExtraImageUrl8Ok

`func (o *MerchantProductRequest) GetExtraImageUrl8Ok() (*string, bool)`

GetExtraImageUrl8Ok returns a tuple with the ExtraImageUrl8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl8

`func (o *MerchantProductRequest) SetExtraImageUrl8(v string)`

SetExtraImageUrl8 sets ExtraImageUrl8 field to given value.

### HasExtraImageUrl8

`func (o *MerchantProductRequest) HasExtraImageUrl8() bool`

HasExtraImageUrl8 returns a boolean if a field has been set.

### SetExtraImageUrl8Nil

`func (o *MerchantProductRequest) SetExtraImageUrl8Nil(b bool)`

 SetExtraImageUrl8Nil sets the value for ExtraImageUrl8 to be an explicit nil

### UnsetExtraImageUrl8
`func (o *MerchantProductRequest) UnsetExtraImageUrl8()`

UnsetExtraImageUrl8 ensures that no value is present for ExtraImageUrl8, not even an explicit nil
### GetExtraImageUrl9

`func (o *MerchantProductRequest) GetExtraImageUrl9() string`

GetExtraImageUrl9 returns the ExtraImageUrl9 field if non-nil, zero value otherwise.

### GetExtraImageUrl9Ok

`func (o *MerchantProductRequest) GetExtraImageUrl9Ok() (*string, bool)`

GetExtraImageUrl9Ok returns a tuple with the ExtraImageUrl9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraImageUrl9

`func (o *MerchantProductRequest) SetExtraImageUrl9(v string)`

SetExtraImageUrl9 sets ExtraImageUrl9 field to given value.

### HasExtraImageUrl9

`func (o *MerchantProductRequest) HasExtraImageUrl9() bool`

HasExtraImageUrl9 returns a boolean if a field has been set.

### SetExtraImageUrl9Nil

`func (o *MerchantProductRequest) SetExtraImageUrl9Nil(b bool)`

 SetExtraImageUrl9Nil sets the value for ExtraImageUrl9 to be an explicit nil

### UnsetExtraImageUrl9
`func (o *MerchantProductRequest) UnsetExtraImageUrl9()`

UnsetExtraImageUrl9 ensures that no value is present for ExtraImageUrl9, not even an explicit nil
### GetCategoryTrail

`func (o *MerchantProductRequest) GetCategoryTrail() string`

GetCategoryTrail returns the CategoryTrail field if non-nil, zero value otherwise.

### GetCategoryTrailOk

`func (o *MerchantProductRequest) GetCategoryTrailOk() (*string, bool)`

GetCategoryTrailOk returns a tuple with the CategoryTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTrail

`func (o *MerchantProductRequest) SetCategoryTrail(v string)`

SetCategoryTrail sets CategoryTrail field to given value.

### HasCategoryTrail

`func (o *MerchantProductRequest) HasCategoryTrail() bool`

HasCategoryTrail returns a boolean if a field has been set.

### SetCategoryTrailNil

`func (o *MerchantProductRequest) SetCategoryTrailNil(b bool)`

 SetCategoryTrailNil sets the value for CategoryTrail to be an explicit nil

### UnsetCategoryTrail
`func (o *MerchantProductRequest) UnsetCategoryTrail()`

UnsetCategoryTrail ensures that no value is present for CategoryTrail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


