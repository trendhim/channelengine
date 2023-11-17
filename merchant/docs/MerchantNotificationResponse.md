# MerchantNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier used by ChannelEngine. | [optional] 
**Read** | Pointer to **bool** | Indicating whether the notification is already read using the backoffice. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Get the created date time. | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**NotificationType**](NotificationType.md) |  | [optional] 

## Methods

### NewMerchantNotificationResponse

`func NewMerchantNotificationResponse() *MerchantNotificationResponse`

NewMerchantNotificationResponse instantiates a new MerchantNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantNotificationResponseWithDefaults

`func NewMerchantNotificationResponseWithDefaults() *MerchantNotificationResponse`

NewMerchantNotificationResponseWithDefaults instantiates a new MerchantNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantNotificationResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantNotificationResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantNotificationResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantNotificationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRead

`func (o *MerchantNotificationResponse) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *MerchantNotificationResponse) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *MerchantNotificationResponse) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *MerchantNotificationResponse) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MerchantNotificationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantNotificationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantNotificationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantNotificationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *MerchantNotificationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MerchantNotificationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MerchantNotificationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MerchantNotificationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MerchantNotificationResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MerchantNotificationResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSubject

`func (o *MerchantNotificationResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MerchantNotificationResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MerchantNotificationResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MerchantNotificationResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *MerchantNotificationResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *MerchantNotificationResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetCount

`func (o *MerchantNotificationResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MerchantNotificationResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MerchantNotificationResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MerchantNotificationResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *MerchantNotificationResponse) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantNotificationResponse) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantNotificationResponse) SetType(v NotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *MerchantNotificationResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


