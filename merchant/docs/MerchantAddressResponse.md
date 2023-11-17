# MerchantAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **NullableString** | The first address line, use this field if address validation is disabled in ChannelEngine. | [optional] 
**Line2** | Pointer to **NullableString** | The second address line, use this field if address validation is disabled in ChannelEngine. | [optional] 
**Line3** | Pointer to **NullableString** | The third address line, use this field if address validation is disabled in ChannelEngine. | [optional] 
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

### NewMerchantAddressResponse

`func NewMerchantAddressResponse() *MerchantAddressResponse`

NewMerchantAddressResponse instantiates a new MerchantAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAddressResponseWithDefaults

`func NewMerchantAddressResponseWithDefaults() *MerchantAddressResponse`

NewMerchantAddressResponseWithDefaults instantiates a new MerchantAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *MerchantAddressResponse) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *MerchantAddressResponse) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *MerchantAddressResponse) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *MerchantAddressResponse) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *MerchantAddressResponse) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *MerchantAddressResponse) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *MerchantAddressResponse) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *MerchantAddressResponse) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *MerchantAddressResponse) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *MerchantAddressResponse) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *MerchantAddressResponse) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *MerchantAddressResponse) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *MerchantAddressResponse) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *MerchantAddressResponse) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *MerchantAddressResponse) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *MerchantAddressResponse) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *MerchantAddressResponse) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *MerchantAddressResponse) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetGender

`func (o *MerchantAddressResponse) GetGender() Gender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *MerchantAddressResponse) GetGenderOk() (*Gender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *MerchantAddressResponse) SetGender(v Gender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *MerchantAddressResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetCompanyName

`func (o *MerchantAddressResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MerchantAddressResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *MerchantAddressResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *MerchantAddressResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *MerchantAddressResponse) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *MerchantAddressResponse) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetFirstName

`func (o *MerchantAddressResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MerchantAddressResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MerchantAddressResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MerchantAddressResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MerchantAddressResponse) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MerchantAddressResponse) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MerchantAddressResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MerchantAddressResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MerchantAddressResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MerchantAddressResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MerchantAddressResponse) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MerchantAddressResponse) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetStreetName

`func (o *MerchantAddressResponse) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *MerchantAddressResponse) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *MerchantAddressResponse) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *MerchantAddressResponse) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *MerchantAddressResponse) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *MerchantAddressResponse) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetHouseNr

`func (o *MerchantAddressResponse) GetHouseNr() string`

GetHouseNr returns the HouseNr field if non-nil, zero value otherwise.

### GetHouseNrOk

`func (o *MerchantAddressResponse) GetHouseNrOk() (*string, bool)`

GetHouseNrOk returns a tuple with the HouseNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNr

`func (o *MerchantAddressResponse) SetHouseNr(v string)`

SetHouseNr sets HouseNr field to given value.

### HasHouseNr

`func (o *MerchantAddressResponse) HasHouseNr() bool`

HasHouseNr returns a boolean if a field has been set.

### SetHouseNrNil

`func (o *MerchantAddressResponse) SetHouseNrNil(b bool)`

 SetHouseNrNil sets the value for HouseNr to be an explicit nil

### UnsetHouseNr
`func (o *MerchantAddressResponse) UnsetHouseNr()`

UnsetHouseNr ensures that no value is present for HouseNr, not even an explicit nil
### GetHouseNrAddition

`func (o *MerchantAddressResponse) GetHouseNrAddition() string`

GetHouseNrAddition returns the HouseNrAddition field if non-nil, zero value otherwise.

### GetHouseNrAdditionOk

`func (o *MerchantAddressResponse) GetHouseNrAdditionOk() (*string, bool)`

GetHouseNrAdditionOk returns a tuple with the HouseNrAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNrAddition

`func (o *MerchantAddressResponse) SetHouseNrAddition(v string)`

SetHouseNrAddition sets HouseNrAddition field to given value.

### HasHouseNrAddition

`func (o *MerchantAddressResponse) HasHouseNrAddition() bool`

HasHouseNrAddition returns a boolean if a field has been set.

### SetHouseNrAdditionNil

`func (o *MerchantAddressResponse) SetHouseNrAdditionNil(b bool)`

 SetHouseNrAdditionNil sets the value for HouseNrAddition to be an explicit nil

### UnsetHouseNrAddition
`func (o *MerchantAddressResponse) UnsetHouseNrAddition()`

UnsetHouseNrAddition ensures that no value is present for HouseNrAddition, not even an explicit nil
### GetZipCode

`func (o *MerchantAddressResponse) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *MerchantAddressResponse) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *MerchantAddressResponse) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *MerchantAddressResponse) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *MerchantAddressResponse) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *MerchantAddressResponse) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCity

`func (o *MerchantAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MerchantAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MerchantAddressResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MerchantAddressResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MerchantAddressResponse) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MerchantAddressResponse) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *MerchantAddressResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MerchantAddressResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MerchantAddressResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MerchantAddressResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MerchantAddressResponse) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MerchantAddressResponse) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCountryIso

`func (o *MerchantAddressResponse) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *MerchantAddressResponse) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *MerchantAddressResponse) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *MerchantAddressResponse) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### SetCountryIsoNil

`func (o *MerchantAddressResponse) SetCountryIsoNil(b bool)`

 SetCountryIsoNil sets the value for CountryIso to be an explicit nil

### UnsetCountryIso
`func (o *MerchantAddressResponse) UnsetCountryIso()`

UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
### GetOriginal

`func (o *MerchantAddressResponse) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *MerchantAddressResponse) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *MerchantAddressResponse) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *MerchantAddressResponse) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### SetOriginalNil

`func (o *MerchantAddressResponse) SetOriginalNil(b bool)`

 SetOriginalNil sets the value for Original to be an explicit nil

### UnsetOriginal
`func (o *MerchantAddressResponse) UnsetOriginal()`

UnsetOriginal ensures that no value is present for Original, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


