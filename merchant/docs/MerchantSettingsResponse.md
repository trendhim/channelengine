# MerchantSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**DefaultCultureCode** | Pointer to **NullableString** |  | [optional] 
**Settings** | Pointer to [**SettingsResponse**](SettingsResponse.md) |  | [optional] 
**Vat** | Pointer to [**[]VatSettingsResponse**](VatSettingsResponse.md) |  | [optional] 

## Methods

### NewMerchantSettingsResponse

`func NewMerchantSettingsResponse() *MerchantSettingsResponse`

NewMerchantSettingsResponse instantiates a new MerchantSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSettingsResponseWithDefaults

`func NewMerchantSettingsResponseWithDefaults() *MerchantSettingsResponse`

NewMerchantSettingsResponseWithDefaults instantiates a new MerchantSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantSettingsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantSettingsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantSettingsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantSettingsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantSettingsResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantSettingsResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompanyName

`func (o *MerchantSettingsResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MerchantSettingsResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *MerchantSettingsResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *MerchantSettingsResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *MerchantSettingsResponse) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *MerchantSettingsResponse) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCurrencyCode

`func (o *MerchantSettingsResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MerchantSettingsResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MerchantSettingsResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MerchantSettingsResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *MerchantSettingsResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *MerchantSettingsResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDefaultCultureCode

`func (o *MerchantSettingsResponse) GetDefaultCultureCode() string`

GetDefaultCultureCode returns the DefaultCultureCode field if non-nil, zero value otherwise.

### GetDefaultCultureCodeOk

`func (o *MerchantSettingsResponse) GetDefaultCultureCodeOk() (*string, bool)`

GetDefaultCultureCodeOk returns a tuple with the DefaultCultureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCultureCode

`func (o *MerchantSettingsResponse) SetDefaultCultureCode(v string)`

SetDefaultCultureCode sets DefaultCultureCode field to given value.

### HasDefaultCultureCode

`func (o *MerchantSettingsResponse) HasDefaultCultureCode() bool`

HasDefaultCultureCode returns a boolean if a field has been set.

### SetDefaultCultureCodeNil

`func (o *MerchantSettingsResponse) SetDefaultCultureCodeNil(b bool)`

 SetDefaultCultureCodeNil sets the value for DefaultCultureCode to be an explicit nil

### UnsetDefaultCultureCode
`func (o *MerchantSettingsResponse) UnsetDefaultCultureCode()`

UnsetDefaultCultureCode ensures that no value is present for DefaultCultureCode, not even an explicit nil
### GetSettings

`func (o *MerchantSettingsResponse) GetSettings() SettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MerchantSettingsResponse) GetSettingsOk() (*SettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MerchantSettingsResponse) SetSettings(v SettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MerchantSettingsResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetVat

`func (o *MerchantSettingsResponse) GetVat() []VatSettingsResponse`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *MerchantSettingsResponse) GetVatOk() (*[]VatSettingsResponse, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *MerchantSettingsResponse) SetVat(v []VatSettingsResponse)`

SetVat sets Vat field to given value.

### HasVat

`func (o *MerchantSettingsResponse) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *MerchantSettingsResponse) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *MerchantSettingsResponse) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


