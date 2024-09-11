# SingleMerchantAcknowledgeRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**IdentifierType** | Pointer to [**RefundIdentifier**](RefundIdentifier.md) |  | [optional] 
**Model** | Pointer to [**MerchantAcknowledgeRefund**](MerchantAcknowledgeRefund.md) |  | [optional] 

## Methods

### NewSingleMerchantAcknowledgeRefundRequest

`func NewSingleMerchantAcknowledgeRefundRequest() *SingleMerchantAcknowledgeRefundRequest`

NewSingleMerchantAcknowledgeRefundRequest instantiates a new SingleMerchantAcknowledgeRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantAcknowledgeRefundRequestWithDefaults

`func NewSingleMerchantAcknowledgeRefundRequestWithDefaults() *SingleMerchantAcknowledgeRefundRequest`

NewSingleMerchantAcknowledgeRefundRequestWithDefaults instantiates a new SingleMerchantAcknowledgeRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantAcknowledgeRefundRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantAcknowledgeRefundRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantAcknowledgeRefundRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantAcknowledgeRefundRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantAcknowledgeRefundRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantAcknowledgeRefundRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetIdentifierType

`func (o *SingleMerchantAcknowledgeRefundRequest) GetIdentifierType() RefundIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantAcknowledgeRefundRequest) GetIdentifierTypeOk() (*RefundIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantAcknowledgeRefundRequest) SetIdentifierType(v RefundIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantAcknowledgeRefundRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantAcknowledgeRefundRequest) GetModel() MerchantAcknowledgeRefund`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantAcknowledgeRefundRequest) GetModelOk() (*MerchantAcknowledgeRefund, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantAcknowledgeRefundRequest) SetModel(v MerchantAcknowledgeRefund)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantAcknowledgeRefundRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


