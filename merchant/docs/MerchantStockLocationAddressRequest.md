# MerchantStockLocationAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryIso** | **string** |  | 
**StreetName** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**HouseNr** | Pointer to **NullableString** |  | [optional] 
**HouseNrAddition** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMerchantStockLocationAddressRequest

`func NewMerchantStockLocationAddressRequest(countryIso string, ) *MerchantStockLocationAddressRequest`

NewMerchantStockLocationAddressRequest instantiates a new MerchantStockLocationAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantStockLocationAddressRequestWithDefaults

`func NewMerchantStockLocationAddressRequestWithDefaults() *MerchantStockLocationAddressRequest`

NewMerchantStockLocationAddressRequestWithDefaults instantiates a new MerchantStockLocationAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryIso

`func (o *MerchantStockLocationAddressRequest) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *MerchantStockLocationAddressRequest) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *MerchantStockLocationAddressRequest) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.


### GetStreetName

`func (o *MerchantStockLocationAddressRequest) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *MerchantStockLocationAddressRequest) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *MerchantStockLocationAddressRequest) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *MerchantStockLocationAddressRequest) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *MerchantStockLocationAddressRequest) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *MerchantStockLocationAddressRequest) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetZipCode

`func (o *MerchantStockLocationAddressRequest) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *MerchantStockLocationAddressRequest) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *MerchantStockLocationAddressRequest) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *MerchantStockLocationAddressRequest) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *MerchantStockLocationAddressRequest) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *MerchantStockLocationAddressRequest) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetHouseNr

`func (o *MerchantStockLocationAddressRequest) GetHouseNr() string`

GetHouseNr returns the HouseNr field if non-nil, zero value otherwise.

### GetHouseNrOk

`func (o *MerchantStockLocationAddressRequest) GetHouseNrOk() (*string, bool)`

GetHouseNrOk returns a tuple with the HouseNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNr

`func (o *MerchantStockLocationAddressRequest) SetHouseNr(v string)`

SetHouseNr sets HouseNr field to given value.

### HasHouseNr

`func (o *MerchantStockLocationAddressRequest) HasHouseNr() bool`

HasHouseNr returns a boolean if a field has been set.

### SetHouseNrNil

`func (o *MerchantStockLocationAddressRequest) SetHouseNrNil(b bool)`

 SetHouseNrNil sets the value for HouseNr to be an explicit nil

### UnsetHouseNr
`func (o *MerchantStockLocationAddressRequest) UnsetHouseNr()`

UnsetHouseNr ensures that no value is present for HouseNr, not even an explicit nil
### GetHouseNrAddition

`func (o *MerchantStockLocationAddressRequest) GetHouseNrAddition() string`

GetHouseNrAddition returns the HouseNrAddition field if non-nil, zero value otherwise.

### GetHouseNrAdditionOk

`func (o *MerchantStockLocationAddressRequest) GetHouseNrAdditionOk() (*string, bool)`

GetHouseNrAdditionOk returns a tuple with the HouseNrAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNrAddition

`func (o *MerchantStockLocationAddressRequest) SetHouseNrAddition(v string)`

SetHouseNrAddition sets HouseNrAddition field to given value.

### HasHouseNrAddition

`func (o *MerchantStockLocationAddressRequest) HasHouseNrAddition() bool`

HasHouseNrAddition returns a boolean if a field has been set.

### SetHouseNrAdditionNil

`func (o *MerchantStockLocationAddressRequest) SetHouseNrAdditionNil(b bool)`

 SetHouseNrAdditionNil sets the value for HouseNrAddition to be an explicit nil

### UnsetHouseNrAddition
`func (o *MerchantStockLocationAddressRequest) UnsetHouseNrAddition()`

UnsetHouseNrAddition ensures that no value is present for HouseNrAddition, not even an explicit nil
### GetCity

`func (o *MerchantStockLocationAddressRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MerchantStockLocationAddressRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MerchantStockLocationAddressRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MerchantStockLocationAddressRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MerchantStockLocationAddressRequest) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MerchantStockLocationAddressRequest) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *MerchantStockLocationAddressRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MerchantStockLocationAddressRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MerchantStockLocationAddressRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MerchantStockLocationAddressRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MerchantStockLocationAddressRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MerchantStockLocationAddressRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


