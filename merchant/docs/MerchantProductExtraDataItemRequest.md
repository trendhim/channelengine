# MerchantProductExtraDataItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Name of the extra data field. | 
**Value** | Pointer to **NullableString** | Value of the extra data field. | [optional] 
**Type** | Pointer to [**ExtraDataType**](ExtraDataType.md) |  | [optional] 
**IsPublic** | Pointer to **bool** | Add this field to the export of the product feed to the channel. | [optional] 
**LanguageIsoCode** | Pointer to **NullableString** | The 2-letter ISO code of the extra data | [optional] 

## Methods

### NewMerchantProductExtraDataItemRequest

`func NewMerchantProductExtraDataItemRequest(key string, ) *MerchantProductExtraDataItemRequest`

NewMerchantProductExtraDataItemRequest instantiates a new MerchantProductExtraDataItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductExtraDataItemRequestWithDefaults

`func NewMerchantProductExtraDataItemRequestWithDefaults() *MerchantProductExtraDataItemRequest`

NewMerchantProductExtraDataItemRequestWithDefaults instantiates a new MerchantProductExtraDataItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MerchantProductExtraDataItemRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MerchantProductExtraDataItemRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MerchantProductExtraDataItemRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *MerchantProductExtraDataItemRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MerchantProductExtraDataItemRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MerchantProductExtraDataItemRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MerchantProductExtraDataItemRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MerchantProductExtraDataItemRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MerchantProductExtraDataItemRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *MerchantProductExtraDataItemRequest) GetType() ExtraDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantProductExtraDataItemRequest) GetTypeOk() (*ExtraDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantProductExtraDataItemRequest) SetType(v ExtraDataType)`

SetType sets Type field to given value.

### HasType

`func (o *MerchantProductExtraDataItemRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsPublic

`func (o *MerchantProductExtraDataItemRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *MerchantProductExtraDataItemRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *MerchantProductExtraDataItemRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *MerchantProductExtraDataItemRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLanguageIsoCode

`func (o *MerchantProductExtraDataItemRequest) GetLanguageIsoCode() string`

GetLanguageIsoCode returns the LanguageIsoCode field if non-nil, zero value otherwise.

### GetLanguageIsoCodeOk

`func (o *MerchantProductExtraDataItemRequest) GetLanguageIsoCodeOk() (*string, bool)`

GetLanguageIsoCodeOk returns a tuple with the LanguageIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageIsoCode

`func (o *MerchantProductExtraDataItemRequest) SetLanguageIsoCode(v string)`

SetLanguageIsoCode sets LanguageIsoCode field to given value.

### HasLanguageIsoCode

`func (o *MerchantProductExtraDataItemRequest) HasLanguageIsoCode() bool`

HasLanguageIsoCode returns a boolean if a field has been set.

### SetLanguageIsoCodeNil

`func (o *MerchantProductExtraDataItemRequest) SetLanguageIsoCodeNil(b bool)`

 SetLanguageIsoCodeNil sets the value for LanguageIsoCode to be an explicit nil

### UnsetLanguageIsoCode
`func (o *MerchantProductExtraDataItemRequest) UnsetLanguageIsoCode()`

UnsetLanguageIsoCode ensures that no value is present for LanguageIsoCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


