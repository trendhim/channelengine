# MerchantWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Events** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) |  | [optional] 

## Methods

### NewMerchantWebhookResponse

`func NewMerchantWebhookResponse() *MerchantWebhookResponse`

NewMerchantWebhookResponse instantiates a new MerchantWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWebhookResponseWithDefaults

`func NewMerchantWebhookResponseWithDefaults() *MerchantWebhookResponse`

NewMerchantWebhookResponseWithDefaults instantiates a new MerchantWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantWebhookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantWebhookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantWebhookResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantWebhookResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantWebhookResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantWebhookResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *MerchantWebhookResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MerchantWebhookResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MerchantWebhookResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MerchantWebhookResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MerchantWebhookResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MerchantWebhookResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIsActive

`func (o *MerchantWebhookResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MerchantWebhookResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MerchantWebhookResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MerchantWebhookResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEvents

`func (o *MerchantWebhookResponse) GetEvents() []WebhookEventType`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MerchantWebhookResponse) GetEventsOk() (*[]WebhookEventType, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MerchantWebhookResponse) SetEvents(v []WebhookEventType)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MerchantWebhookResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *MerchantWebhookResponse) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *MerchantWebhookResponse) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


