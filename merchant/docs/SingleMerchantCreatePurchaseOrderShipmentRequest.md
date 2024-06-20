# SingleMerchantCreatePurchaseOrderShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **int32** | The identifier of the channel | [optional] 
**Model** | Pointer to [**CreatePurchaseOrderShipment**](CreatePurchaseOrderShipment.md) |  | [optional] 

## Methods

### NewSingleMerchantCreatePurchaseOrderShipmentRequest

`func NewSingleMerchantCreatePurchaseOrderShipmentRequest() *SingleMerchantCreatePurchaseOrderShipmentRequest`

NewSingleMerchantCreatePurchaseOrderShipmentRequest instantiates a new SingleMerchantCreatePurchaseOrderShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantCreatePurchaseOrderShipmentRequestWithDefaults

`func NewSingleMerchantCreatePurchaseOrderShipmentRequestWithDefaults() *SingleMerchantCreatePurchaseOrderShipmentRequest`

NewSingleMerchantCreatePurchaseOrderShipmentRequestWithDefaults instantiates a new SingleMerchantCreatePurchaseOrderShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) GetModel() CreatePurchaseOrderShipment`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) GetModelOk() (*CreatePurchaseOrderShipment, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) SetModel(v CreatePurchaseOrderShipment)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantCreatePurchaseOrderShipmentRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


