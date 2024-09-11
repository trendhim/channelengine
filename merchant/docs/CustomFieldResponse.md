# CustomFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**IsUsed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomFieldResponse

`func NewCustomFieldResponse() *CustomFieldResponse`

NewCustomFieldResponse instantiates a new CustomFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldResponseWithDefaults

`func NewCustomFieldResponseWithDefaults() *CustomFieldResponse`

NewCustomFieldResponseWithDefaults instantiates a new CustomFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CustomFieldResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomFieldResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomFieldResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomFieldResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CustomFieldResponse) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CustomFieldResponse) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetIsPublic

`func (o *CustomFieldResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CustomFieldResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CustomFieldResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CustomFieldResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsUsed

`func (o *CustomFieldResponse) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *CustomFieldResponse) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *CustomFieldResponse) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.

### HasIsUsed

`func (o *CustomFieldResponse) HasIsUsed() bool`

HasIsUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


