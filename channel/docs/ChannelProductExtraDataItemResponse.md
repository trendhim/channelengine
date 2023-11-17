# ChannelProductExtraDataItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Name of the extra data field. | 
**Value** | Pointer to **NullableString** | Value of the extra data field. | [optional] 
**Type** | Pointer to [**ExtraDataType**](ExtraDataType.md) |  | [optional] 
**IsPublic** | Pointer to **bool** | Add this field to the export of the product feed to the channel. | [optional] 
**LanguageIsoCode** | Pointer to **NullableString** | The 2-letter ISO code of the extra data | [optional] 

## Methods

### NewChannelProductExtraDataItemResponse

`func NewChannelProductExtraDataItemResponse(key string, ) *ChannelProductExtraDataItemResponse`

NewChannelProductExtraDataItemResponse instantiates a new ChannelProductExtraDataItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProductExtraDataItemResponseWithDefaults

`func NewChannelProductExtraDataItemResponseWithDefaults() *ChannelProductExtraDataItemResponse`

NewChannelProductExtraDataItemResponseWithDefaults instantiates a new ChannelProductExtraDataItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ChannelProductExtraDataItemResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ChannelProductExtraDataItemResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ChannelProductExtraDataItemResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ChannelProductExtraDataItemResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ChannelProductExtraDataItemResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ChannelProductExtraDataItemResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ChannelProductExtraDataItemResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ChannelProductExtraDataItemResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ChannelProductExtraDataItemResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *ChannelProductExtraDataItemResponse) GetType() ExtraDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelProductExtraDataItemResponse) GetTypeOk() (*ExtraDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelProductExtraDataItemResponse) SetType(v ExtraDataType)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelProductExtraDataItemResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsPublic

`func (o *ChannelProductExtraDataItemResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ChannelProductExtraDataItemResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ChannelProductExtraDataItemResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ChannelProductExtraDataItemResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLanguageIsoCode

`func (o *ChannelProductExtraDataItemResponse) GetLanguageIsoCode() string`

GetLanguageIsoCode returns the LanguageIsoCode field if non-nil, zero value otherwise.

### GetLanguageIsoCodeOk

`func (o *ChannelProductExtraDataItemResponse) GetLanguageIsoCodeOk() (*string, bool)`

GetLanguageIsoCodeOk returns a tuple with the LanguageIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageIsoCode

`func (o *ChannelProductExtraDataItemResponse) SetLanguageIsoCode(v string)`

SetLanguageIsoCode sets LanguageIsoCode field to given value.

### HasLanguageIsoCode

`func (o *ChannelProductExtraDataItemResponse) HasLanguageIsoCode() bool`

HasLanguageIsoCode returns a boolean if a field has been set.

### SetLanguageIsoCodeNil

`func (o *ChannelProductExtraDataItemResponse) SetLanguageIsoCodeNil(b bool)`

 SetLanguageIsoCodeNil sets the value for LanguageIsoCode to be an explicit nil

### UnsetLanguageIsoCode
`func (o *ChannelProductExtraDataItemResponse) UnsetLanguageIsoCode()`

UnsetLanguageIsoCode ensures that no value is present for LanguageIsoCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


