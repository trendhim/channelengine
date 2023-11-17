# ProductExtraDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Key** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProductExtraDataRequest

`func NewProductExtraDataRequest(op string, key string, ) *ProductExtraDataRequest`

NewProductExtraDataRequest instantiates a new ProductExtraDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductExtraDataRequestWithDefaults

`func NewProductExtraDataRequestWithDefaults() *ProductExtraDataRequest`

NewProductExtraDataRequestWithDefaults instantiates a new ProductExtraDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ProductExtraDataRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ProductExtraDataRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ProductExtraDataRequest) SetOp(v string)`

SetOp sets Op field to given value.


### GetKey

`func (o *ProductExtraDataRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProductExtraDataRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProductExtraDataRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ProductExtraDataRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductExtraDataRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductExtraDataRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductExtraDataRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProductExtraDataRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProductExtraDataRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


