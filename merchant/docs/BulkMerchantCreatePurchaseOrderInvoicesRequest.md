# BulkMerchantCreatePurchaseOrderInvoicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to [**[]MerchantPurchaseOrderInvoice**](MerchantPurchaseOrderInvoice.md) |  | [optional] 
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBulkMerchantCreatePurchaseOrderInvoicesRequest

`func NewBulkMerchantCreatePurchaseOrderInvoicesRequest() *BulkMerchantCreatePurchaseOrderInvoicesRequest`

NewBulkMerchantCreatePurchaseOrderInvoicesRequest instantiates a new BulkMerchantCreatePurchaseOrderInvoicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMerchantCreatePurchaseOrderInvoicesRequestWithDefaults

`func NewBulkMerchantCreatePurchaseOrderInvoicesRequestWithDefaults() *BulkMerchantCreatePurchaseOrderInvoicesRequest`

NewBulkMerchantCreatePurchaseOrderInvoicesRequestWithDefaults instantiates a new BulkMerchantCreatePurchaseOrderInvoicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) GetModel() []MerchantPurchaseOrderInvoice`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) GetModelOk() (*[]MerchantPurchaseOrderInvoice, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) SetModel(v []MerchantPurchaseOrderInvoice)`

SetModel sets Model field to given value.

### HasModel

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetChannelId

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *BulkMerchantCreatePurchaseOrderInvoicesRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


