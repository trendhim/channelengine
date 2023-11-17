# ChannelProductReferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ChannelEngine product ID. | [optional] 
**ChannelProductNo** | **string** | The unique product reference used by the Channel. | 

## Methods

### NewChannelProductReferencesRequest

`func NewChannelProductReferencesRequest(channelProductNo string, ) *ChannelProductReferencesRequest`

NewChannelProductReferencesRequest instantiates a new ChannelProductReferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProductReferencesRequestWithDefaults

`func NewChannelProductReferencesRequestWithDefaults() *ChannelProductReferencesRequest`

NewChannelProductReferencesRequestWithDefaults instantiates a new ChannelProductReferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelProductReferencesRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelProductReferencesRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelProductReferencesRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelProductReferencesRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelProductNo

`func (o *ChannelProductReferencesRequest) GetChannelProductNo() string`

GetChannelProductNo returns the ChannelProductNo field if non-nil, zero value otherwise.

### GetChannelProductNoOk

`func (o *ChannelProductReferencesRequest) GetChannelProductNoOk() (*string, bool)`

GetChannelProductNoOk returns a tuple with the ChannelProductNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProductNo

`func (o *ChannelProductReferencesRequest) SetChannelProductNo(v string)`

SetChannelProductNo sets ChannelProductNo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


