# MerchantProductBundleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantProductNo** | Pointer to **NullableString** |  | [optional] 
**Ean** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Parts** | Pointer to [**[]MerchantProductBundlePartResponse**](MerchantProductBundlePartResponse.md) |  | [optional] 

## Methods

### NewMerchantProductBundleResponse

`func NewMerchantProductBundleResponse() *MerchantProductBundleResponse`

NewMerchantProductBundleResponse instantiates a new MerchantProductBundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductBundleResponseWithDefaults

`func NewMerchantProductBundleResponseWithDefaults() *MerchantProductBundleResponse`

NewMerchantProductBundleResponseWithDefaults instantiates a new MerchantProductBundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantProductNo

`func (o *MerchantProductBundleResponse) GetMerchantProductNo() string`

GetMerchantProductNo returns the MerchantProductNo field if non-nil, zero value otherwise.

### GetMerchantProductNoOk

`func (o *MerchantProductBundleResponse) GetMerchantProductNoOk() (*string, bool)`

GetMerchantProductNoOk returns a tuple with the MerchantProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProductNo

`func (o *MerchantProductBundleResponse) SetMerchantProductNo(v string)`

SetMerchantProductNo sets MerchantProductNo field to given value.

### HasMerchantProductNo

`func (o *MerchantProductBundleResponse) HasMerchantProductNo() bool`

HasMerchantProductNo returns a boolean if a field has been set.

### SetMerchantProductNoNil

`func (o *MerchantProductBundleResponse) SetMerchantProductNoNil(b bool)`

 SetMerchantProductNoNil sets the value for MerchantProductNo to be an explicit nil

### UnsetMerchantProductNo
`func (o *MerchantProductBundleResponse) UnsetMerchantProductNo()`

UnsetMerchantProductNo ensures that no value is present for MerchantProductNo, not even an explicit nil
### GetEan

`func (o *MerchantProductBundleResponse) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *MerchantProductBundleResponse) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *MerchantProductBundleResponse) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *MerchantProductBundleResponse) HasEan() bool`

HasEan returns a boolean if a field has been set.

### SetEanNil

`func (o *MerchantProductBundleResponse) SetEanNil(b bool)`

 SetEanNil sets the value for Ean to be an explicit nil

### UnsetEan
`func (o *MerchantProductBundleResponse) UnsetEan()`

UnsetEan ensures that no value is present for Ean, not even an explicit nil
### GetName

`func (o *MerchantProductBundleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantProductBundleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantProductBundleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantProductBundleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantProductBundleResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantProductBundleResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrice

`func (o *MerchantProductBundleResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MerchantProductBundleResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MerchantProductBundleResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MerchantProductBundleResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetParts

`func (o *MerchantProductBundleResponse) GetParts() []MerchantProductBundlePartResponse`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MerchantProductBundleResponse) GetPartsOk() (*[]MerchantProductBundlePartResponse, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MerchantProductBundleResponse) SetParts(v []MerchantProductBundlePartResponse)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MerchantProductBundleResponse) HasParts() bool`

HasParts returns a boolean if a field has been set.

### SetPartsNil

`func (o *MerchantProductBundleResponse) SetPartsNil(b bool)`

 SetPartsNil sets the value for Parts to be an explicit nil

### UnsetParts
`func (o *MerchantProductBundleResponse) UnsetParts()`

UnsetParts ensures that no value is present for Parts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


