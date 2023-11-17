# MerchantWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of a webhook. | 
**Url** | **string** | The callback URL used by a webhook. E.g.: https://test-store.com/callback. | 
**IsActive** | Pointer to **bool** | Determines if a webhook is active, and callbacks should proceed. | [optional] 
**Events** | [**[]WebhookEventType**](WebhookEventType.md) | The events supported by the webhook. | 

## Methods

### NewMerchantWebhookRequest

`func NewMerchantWebhookRequest(name string, url string, events []WebhookEventType, ) *MerchantWebhookRequest`

NewMerchantWebhookRequest instantiates a new MerchantWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWebhookRequestWithDefaults

`func NewMerchantWebhookRequestWithDefaults() *MerchantWebhookRequest`

NewMerchantWebhookRequestWithDefaults instantiates a new MerchantWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *MerchantWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MerchantWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MerchantWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIsActive

`func (o *MerchantWebhookRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MerchantWebhookRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MerchantWebhookRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MerchantWebhookRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEvents

`func (o *MerchantWebhookRequest) GetEvents() []WebhookEventType`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MerchantWebhookRequest) GetEventsOk() (*[]WebhookEventType, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MerchantWebhookRequest) SetEvents(v []WebhookEventType)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


