# SingleMerchantCreatePurchaseOrderInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to [**MerchantPurchaseOrderInvoice**](MerchantPurchaseOrderInvoice.md) |  | [optional] 
**Model** | Pointer to [**[]MerchantPurchaseOrderInvoice**](MerchantPurchaseOrderInvoice.md) |  | [optional] 
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSingleMerchantCreatePurchaseOrderInvoiceRequest

`func NewSingleMerchantCreatePurchaseOrderInvoiceRequest() *SingleMerchantCreatePurchaseOrderInvoiceRequest`

NewSingleMerchantCreatePurchaseOrderInvoiceRequest instantiates a new SingleMerchantCreatePurchaseOrderInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMerchantCreatePurchaseOrderInvoiceRequestWithDefaults

`func NewSingleMerchantCreatePurchaseOrderInvoiceRequestWithDefaults() *SingleMerchantCreatePurchaseOrderInvoiceRequest`

NewSingleMerchantCreatePurchaseOrderInvoiceRequestWithDefaults instantiates a new SingleMerchantCreatePurchaseOrderInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetModels() MerchantPurchaseOrderInvoice`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetModelsOk() (*MerchantPurchaseOrderInvoice, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) SetModels(v MerchantPurchaseOrderInvoice)`

SetModels sets Models field to given value.

### HasModels

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetModel

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetModel() []MerchantPurchaseOrderInvoice`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetModelOk() (*[]MerchantPurchaseOrderInvoice, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) SetModel(v []MerchantPurchaseOrderInvoice)`

SetModel sets Model field to given value.

### HasModel

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetChannelId

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SingleMerchantCreatePurchaseOrderInvoiceRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


