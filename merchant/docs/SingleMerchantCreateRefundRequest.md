# SingleMerchantCreateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**LineIdentifierType** | Pointer to [**OrderLineIdentifier**](OrderLineIdentifier.md) |  | [optional] 
**IdentifierType** | Pointer to [**OrderIdentifier**](OrderIdentifier.md) |  | [optional] 
**Model** | Pointer to [**MerchantCreateRefund**](MerchantCreateRefund.md) |  | [optional] 

## Methods

### NewSingleMerchantCreateRefundRequest

`func NewSingleMerchantCreateRefundRequest() *SingleMerchantCreateRefundRequest`

NewSingleMerchantCreateRefundRequest instantiates a new SingleMerchantCreateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantCreateRefundRequestWithDefaults

`func NewSingleMerchantCreateRefundRequestWithDefaults() *SingleMerchantCreateRefundRequest`

NewSingleMerchantCreateRefundRequestWithDefaults instantiates a new SingleMerchantCreateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantCreateRefundRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantCreateRefundRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantCreateRefundRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantCreateRefundRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantCreateRefundRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantCreateRefundRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetLineIdentifierType

`func (o *SingleMerchantCreateRefundRequest) GetLineIdentifierType() OrderLineIdentifier`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleMerchantCreateRefundRequest) GetLineIdentifierTypeOk() (*OrderLineIdentifier, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleMerchantCreateRefundRequest) SetLineIdentifierType(v OrderLineIdentifier)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleMerchantCreateRefundRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleMerchantCreateRefundRequest) GetIdentifierType() OrderIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantCreateRefundRequest) GetIdentifierTypeOk() (*OrderIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantCreateRefundRequest) SetIdentifierType(v OrderIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantCreateRefundRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantCreateRefundRequest) GetModel() MerchantCreateRefund`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantCreateRefundRequest) GetModelOk() (*MerchantCreateRefund, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantCreateRefundRequest) SetModel(v MerchantCreateRefund)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantCreateRefundRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


