# SingleMerchantUpdatePurchaseOrderShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **int32** | The identifier of the channel | [optional] 
**Model** | Pointer to [**UpdatePurchaseOrderShipment**](UpdatePurchaseOrderShipment.md) |  | [optional] 

## Methods

### NewSingleMerchantUpdatePurchaseOrderShipmentRequest

`func NewSingleMerchantUpdatePurchaseOrderShipmentRequest() *SingleMerchantUpdatePurchaseOrderShipmentRequest`

NewSingleMerchantUpdatePurchaseOrderShipmentRequest instantiates a new SingleMerchantUpdatePurchaseOrderShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantUpdatePurchaseOrderShipmentRequestWithDefaults

`func NewSingleMerchantUpdatePurchaseOrderShipmentRequestWithDefaults() *SingleMerchantUpdatePurchaseOrderShipmentRequest`

NewSingleMerchantUpdatePurchaseOrderShipmentRequestWithDefaults instantiates a new SingleMerchantUpdatePurchaseOrderShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetModel() UpdatePurchaseOrderShipment`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) GetModelOk() (*UpdatePurchaseOrderShipment, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) SetModel(v UpdatePurchaseOrderShipment)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantUpdatePurchaseOrderShipmentRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


