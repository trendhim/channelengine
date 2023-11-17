# ChannelGlobalChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **NullableString** | The country code of the Global Channel. | [optional] 
**GlobalChannelId** | Pointer to **NullableInt32** | The ID of the Global Channel. | [optional] 
**Channels** | Pointer to [**[]ChannelChannelResponse**](ChannelChannelResponse.md) | The status of the instances. | [optional] 
**LanguageCode** | Pointer to **NullableString** | The language code of the Global Channel. | [optional] 
**GlobalChannelName** | Pointer to **NullableString** | The name of the Global Channel. | [optional] 

## Methods

### NewChannelGlobalChannelResponse

`func NewChannelGlobalChannelResponse() *ChannelGlobalChannelResponse`

NewChannelGlobalChannelResponse instantiates a new ChannelGlobalChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelGlobalChannelResponseWithDefaults

`func NewChannelGlobalChannelResponseWithDefaults() *ChannelGlobalChannelResponse`

NewChannelGlobalChannelResponseWithDefaults instantiates a new ChannelGlobalChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *ChannelGlobalChannelResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ChannelGlobalChannelResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ChannelGlobalChannelResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ChannelGlobalChannelResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *ChannelGlobalChannelResponse) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *ChannelGlobalChannelResponse) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetGlobalChannelId

`func (o *ChannelGlobalChannelResponse) GetGlobalChannelId() int32`

GetGlobalChannelId returns the GlobalChannelId field if non-nil, zero value otherwise.

### GetGlobalChannelIdOk

`func (o *ChannelGlobalChannelResponse) GetGlobalChannelIdOk() (*int32, bool)`

GetGlobalChannelIdOk returns a tuple with the GlobalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelId

`func (o *ChannelGlobalChannelResponse) SetGlobalChannelId(v int32)`

SetGlobalChannelId sets GlobalChannelId field to given value.

### HasGlobalChannelId

`func (o *ChannelGlobalChannelResponse) HasGlobalChannelId() bool`

HasGlobalChannelId returns a boolean if a field has been set.

### SetGlobalChannelIdNil

`func (o *ChannelGlobalChannelResponse) SetGlobalChannelIdNil(b bool)`

 SetGlobalChannelIdNil sets the value for GlobalChannelId to be an explicit nil

### UnsetGlobalChannelId
`func (o *ChannelGlobalChannelResponse) UnsetGlobalChannelId()`

UnsetGlobalChannelId ensures that no value is present for GlobalChannelId, not even an explicit nil
### GetChannels

`func (o *ChannelGlobalChannelResponse) GetChannels() []ChannelChannelResponse`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ChannelGlobalChannelResponse) GetChannelsOk() (*[]ChannelChannelResponse, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ChannelGlobalChannelResponse) SetChannels(v []ChannelChannelResponse)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ChannelGlobalChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *ChannelGlobalChannelResponse) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *ChannelGlobalChannelResponse) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetLanguageCode

`func (o *ChannelGlobalChannelResponse) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *ChannelGlobalChannelResponse) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *ChannelGlobalChannelResponse) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *ChannelGlobalChannelResponse) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *ChannelGlobalChannelResponse) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *ChannelGlobalChannelResponse) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetGlobalChannelName

`func (o *ChannelGlobalChannelResponse) GetGlobalChannelName() string`

GetGlobalChannelName returns the GlobalChannelName field if non-nil, zero value otherwise.

### GetGlobalChannelNameOk

`func (o *ChannelGlobalChannelResponse) GetGlobalChannelNameOk() (*string, bool)`

GetGlobalChannelNameOk returns a tuple with the GlobalChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChannelName

`func (o *ChannelGlobalChannelResponse) SetGlobalChannelName(v string)`

SetGlobalChannelName sets GlobalChannelName field to given value.

### HasGlobalChannelName

`func (o *ChannelGlobalChannelResponse) HasGlobalChannelName() bool`

HasGlobalChannelName returns a boolean if a field has been set.

### SetGlobalChannelNameNil

`func (o *ChannelGlobalChannelResponse) SetGlobalChannelNameNil(b bool)`

 SetGlobalChannelNameNil sets the value for GlobalChannelName to be an explicit nil

### UnsetGlobalChannelName
`func (o *ChannelGlobalChannelResponse) UnsetGlobalChannelName()`

UnsetGlobalChannelName ensures that no value is present for GlobalChannelName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


