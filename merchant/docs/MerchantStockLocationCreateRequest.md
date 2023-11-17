# MerchantStockLocationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsDefault** | Pointer to **bool** |  | [optional] 
**FallBackToDefault** | Pointer to **bool** | If false: only use fulfillment by channel, else (also) use merchant fulfillment. | [optional] 
**Address** | Pointer to [**MerchantStockLocationAddressRequest**](MerchantStockLocationAddressRequest.md) |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMerchantStockLocationCreateRequest

`func NewMerchantStockLocationCreateRequest(name string, ) *MerchantStockLocationCreateRequest`

NewMerchantStockLocationCreateRequest instantiates a new MerchantStockLocationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantStockLocationCreateRequestWithDefaults

`func NewMerchantStockLocationCreateRequestWithDefaults() *MerchantStockLocationCreateRequest`

NewMerchantStockLocationCreateRequestWithDefaults instantiates a new MerchantStockLocationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantStockLocationCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantStockLocationCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantStockLocationCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *MerchantStockLocationCreateRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MerchantStockLocationCreateRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MerchantStockLocationCreateRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MerchantStockLocationCreateRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetFallBackToDefault

`func (o *MerchantStockLocationCreateRequest) GetFallBackToDefault() bool`

GetFallBackToDefault returns the FallBackToDefault field if non-nil, zero value otherwise.

### GetFallBackToDefaultOk

`func (o *MerchantStockLocationCreateRequest) GetFallBackToDefaultOk() (*bool, bool)`

GetFallBackToDefaultOk returns a tuple with the FallBackToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallBackToDefault

`func (o *MerchantStockLocationCreateRequest) SetFallBackToDefault(v bool)`

SetFallBackToDefault sets FallBackToDefault field to given value.

### HasFallBackToDefault

`func (o *MerchantStockLocationCreateRequest) HasFallBackToDefault() bool`

HasFallBackToDefault returns a boolean if a field has been set.

### GetAddress

`func (o *MerchantStockLocationCreateRequest) GetAddress() MerchantStockLocationAddressRequest`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MerchantStockLocationCreateRequest) GetAddressOk() (*MerchantStockLocationAddressRequest, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MerchantStockLocationCreateRequest) SetAddress(v MerchantStockLocationAddressRequest)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MerchantStockLocationCreateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MerchantStockLocationCreateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MerchantStockLocationCreateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MerchantStockLocationCreateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MerchantStockLocationCreateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *MerchantStockLocationCreateRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *MerchantStockLocationCreateRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


