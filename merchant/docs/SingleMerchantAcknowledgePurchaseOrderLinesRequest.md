# SingleMerchantAcknowledgePurchaseOrderLinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**LineIdentifierType** | Pointer to [**PurchaseOrderLineIdentifierType**](PurchaseOrderLineIdentifierType.md) |  | [optional] 
**IdentifierType** | Pointer to [**PurchaseOrderIdentifierType**](PurchaseOrderIdentifierType.md) |  | [optional] 
**Model** | Pointer to [**MerchantAcknowledgePurchaseOrder**](MerchantAcknowledgePurchaseOrder.md) |  | [optional] 

## Methods

### NewSingleMerchantAcknowledgePurchaseOrderLinesRequest

`func NewSingleMerchantAcknowledgePurchaseOrderLinesRequest() *SingleMerchantAcknowledgePurchaseOrderLinesRequest`

NewSingleMerchantAcknowledgePurchaseOrderLinesRequest instantiates a new SingleMerchantAcknowledgePurchaseOrderLinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantAcknowledgePurchaseOrderLinesRequestWithDefaults

`func NewSingleMerchantAcknowledgePurchaseOrderLinesRequestWithDefaults() *SingleMerchantAcknowledgePurchaseOrderLinesRequest`

NewSingleMerchantAcknowledgePurchaseOrderLinesRequestWithDefaults instantiates a new SingleMerchantAcknowledgePurchaseOrderLinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetLineIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetLineIdentifierType() PurchaseOrderLineIdentifierType`

GetLineIdentifierType returns the LineIdentifierType field if non-nil, zero value otherwise.

### GetLineIdentifierTypeOk

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetLineIdentifierTypeOk() (*PurchaseOrderLineIdentifierType, bool)`

GetLineIdentifierTypeOk returns a tuple with the LineIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) SetLineIdentifierType(v PurchaseOrderLineIdentifierType)`

SetLineIdentifierType sets LineIdentifierType field to given value.

### HasLineIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) HasLineIdentifierType() bool`

HasLineIdentifierType returns a boolean if a field has been set.

### GetIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetIdentifierType() PurchaseOrderIdentifierType`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetIdentifierTypeOk() (*PurchaseOrderIdentifierType, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) SetIdentifierType(v PurchaseOrderIdentifierType)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetModel() MerchantAcknowledgePurchaseOrder`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) GetModelOk() (*MerchantAcknowledgePurchaseOrder, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) SetModel(v MerchantAcknowledgePurchaseOrder)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantAcknowledgePurchaseOrderLinesRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


