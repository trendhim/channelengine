# SingleChannelProcessRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | Pointer to [**RefundIdentifier**](RefundIdentifier.md) |  | [optional] 
**Model** | Pointer to [**ChannelProcessRefund**](ChannelProcessRefund.md) |  | [optional] 

## Methods

### NewSingleChannelProcessRefundRequest

`func NewSingleChannelProcessRefundRequest() *SingleChannelProcessRefundRequest`

NewSingleChannelProcessRefundRequest instantiates a new SingleChannelProcessRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleChannelProcessRefundRequestWithDefaults

`func NewSingleChannelProcessRefundRequestWithDefaults() *SingleChannelProcessRefundRequest`

NewSingleChannelProcessRefundRequestWithDefaults instantiates a new SingleChannelProcessRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *SingleChannelProcessRefundRequest) GetIdentifierType() RefundIdentifier`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleChannelProcessRefundRequest) GetIdentifierTypeOk() (*RefundIdentifier, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleChannelProcessRefundRequest) SetIdentifierType(v RefundIdentifier)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleChannelProcessRefundRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleChannelProcessRefundRequest) GetModel() ChannelProcessRefund`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleChannelProcessRefundRequest) GetModelOk() (*ChannelProcessRefund, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleChannelProcessRefundRequest) SetModel(v ChannelProcessRefund)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleChannelProcessRefundRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


