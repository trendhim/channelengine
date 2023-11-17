# ChannelAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gender** | Pointer to [**Gender**](Gender.md) |  | [optional] 
**CompanyName** | Pointer to **NullableString** | Optional. Company addressed too. | [optional] 
**FirstName** | Pointer to **NullableString** | The first name of the customer. | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the customer (includes the surname prefix [tussenvoegsel] like &#39;de&#39;, &#39;van&#39;, &#39;du&#39;). | [optional] 
**StreetName** | Pointer to **NullableString** | The name of the street (without house number information)  This field might be empty if address validation is disabled in ChannelEngine. | [optional] 
**HouseNr** | Pointer to **NullableString** | The house number  This field might be empty if address validation is disabled in ChannelEngine. | [optional] 
**HouseNrAddition** | Pointer to **NullableString** | Optional. Addition to the house number  If the address is: Groenhazengracht 2c, the address will be:  StreetName: Groenhazengracht  HouseNo: 2  HouseNrAddition: c  This field might be empty if address validation is disabled in ChannelEngine. | [optional] 
**ZipCode** | Pointer to **NullableString** | The zip or postal code. | [optional] 
**City** | Pointer to **NullableString** | The name of the city. | [optional] 
**Region** | Pointer to **NullableString** | Optional. State/province/region. | [optional] 
**CountryIso** | Pointer to **NullableString** | For example: NL, BE, FR. | [optional] 
**Original** | Pointer to **NullableString** | Optional. The address as a single string: use in case the address lines are entered  as single lines and later parsed into street, house number and house number addition. | [optional] 

## Methods

### NewChannelAddressRequest

`func NewChannelAddressRequest() *ChannelAddressRequest`

NewChannelAddressRequest instantiates a new ChannelAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAddressRequestWithDefaults

`func NewChannelAddressRequestWithDefaults() *ChannelAddressRequest`

NewChannelAddressRequestWithDefaults instantiates a new ChannelAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGender

`func (o *ChannelAddressRequest) GetGender() Gender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ChannelAddressRequest) GetGenderOk() (*Gender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ChannelAddressRequest) SetGender(v Gender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ChannelAddressRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetCompanyName

`func (o *ChannelAddressRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ChannelAddressRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ChannelAddressRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ChannelAddressRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *ChannelAddressRequest) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *ChannelAddressRequest) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetFirstName

`func (o *ChannelAddressRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ChannelAddressRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ChannelAddressRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ChannelAddressRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ChannelAddressRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ChannelAddressRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ChannelAddressRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ChannelAddressRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ChannelAddressRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ChannelAddressRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ChannelAddressRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ChannelAddressRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetStreetName

`func (o *ChannelAddressRequest) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *ChannelAddressRequest) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *ChannelAddressRequest) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *ChannelAddressRequest) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *ChannelAddressRequest) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *ChannelAddressRequest) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetHouseNr

`func (o *ChannelAddressRequest) GetHouseNr() string`

GetHouseNr returns the HouseNr field if non-nil, zero value otherwise.

### GetHouseNrOk

`func (o *ChannelAddressRequest) GetHouseNrOk() (*string, bool)`

GetHouseNrOk returns a tuple with the HouseNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNr

`func (o *ChannelAddressRequest) SetHouseNr(v string)`

SetHouseNr sets HouseNr field to given value.

### HasHouseNr

`func (o *ChannelAddressRequest) HasHouseNr() bool`

HasHouseNr returns a boolean if a field has been set.

### SetHouseNrNil

`func (o *ChannelAddressRequest) SetHouseNrNil(b bool)`

 SetHouseNrNil sets the value for HouseNr to be an explicit nil

### UnsetHouseNr
`func (o *ChannelAddressRequest) UnsetHouseNr()`

UnsetHouseNr ensures that no value is present for HouseNr, not even an explicit nil
### GetHouseNrAddition

`func (o *ChannelAddressRequest) GetHouseNrAddition() string`

GetHouseNrAddition returns the HouseNrAddition field if non-nil, zero value otherwise.

### GetHouseNrAdditionOk

`func (o *ChannelAddressRequest) GetHouseNrAdditionOk() (*string, bool)`

GetHouseNrAdditionOk returns a tuple with the HouseNrAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNrAddition

`func (o *ChannelAddressRequest) SetHouseNrAddition(v string)`

SetHouseNrAddition sets HouseNrAddition field to given value.

### HasHouseNrAddition

`func (o *ChannelAddressRequest) HasHouseNrAddition() bool`

HasHouseNrAddition returns a boolean if a field has been set.

### SetHouseNrAdditionNil

`func (o *ChannelAddressRequest) SetHouseNrAdditionNil(b bool)`

 SetHouseNrAdditionNil sets the value for HouseNrAddition to be an explicit nil

### UnsetHouseNrAddition
`func (o *ChannelAddressRequest) UnsetHouseNrAddition()`

UnsetHouseNrAddition ensures that no value is present for HouseNrAddition, not even an explicit nil
### GetZipCode

`func (o *ChannelAddressRequest) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ChannelAddressRequest) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ChannelAddressRequest) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ChannelAddressRequest) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ChannelAddressRequest) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ChannelAddressRequest) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCity

`func (o *ChannelAddressRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ChannelAddressRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ChannelAddressRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ChannelAddressRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ChannelAddressRequest) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ChannelAddressRequest) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *ChannelAddressRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ChannelAddressRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ChannelAddressRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ChannelAddressRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ChannelAddressRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ChannelAddressRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCountryIso

`func (o *ChannelAddressRequest) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *ChannelAddressRequest) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *ChannelAddressRequest) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *ChannelAddressRequest) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### SetCountryIsoNil

`func (o *ChannelAddressRequest) SetCountryIsoNil(b bool)`

 SetCountryIsoNil sets the value for CountryIso to be an explicit nil

### UnsetCountryIso
`func (o *ChannelAddressRequest) UnsetCountryIso()`

UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
### GetOriginal

`func (o *ChannelAddressRequest) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *ChannelAddressRequest) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *ChannelAddressRequest) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *ChannelAddressRequest) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### SetOriginalNil

`func (o *ChannelAddressRequest) SetOriginalNil(b bool)`

 SetOriginalNil sets the value for Original to be an explicit nil

### UnsetOriginal
`func (o *ChannelAddressRequest) UnsetOriginal()`

UnsetOriginal ensures that no value is present for Original, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


