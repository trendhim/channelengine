# MerchantReturnUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnId** | **int32** | The ChannelEngine return ID of the return you would like to update. | 
**Lines** | [**[]MerchantReturnLineUpdateRequest**](MerchantReturnLineUpdateRequest.md) |  | 

## Methods

### NewMerchantReturnUpdateRequest

`func NewMerchantReturnUpdateRequest(returnId int32, lines []MerchantReturnLineUpdateRequest, ) *MerchantReturnUpdateRequest`

NewMerchantReturnUpdateRequest instantiates a new MerchantReturnUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantReturnUpdateRequestWithDefaults

`func NewMerchantReturnUpdateRequestWithDefaults() *MerchantReturnUpdateRequest`

NewMerchantReturnUpdateRequestWithDefaults instantiates a new MerchantReturnUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnId

`func (o *MerchantReturnUpdateRequest) GetReturnId() int32`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *MerchantReturnUpdateRequest) GetReturnIdOk() (*int32, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *MerchantReturnUpdateRequest) SetReturnId(v int32)`

SetReturnId sets ReturnId field to given value.


### GetLines

`func (o *MerchantReturnUpdateRequest) GetLines() []MerchantReturnLineUpdateRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantReturnUpdateRequest) GetLinesOk() (*[]MerchantReturnLineUpdateRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantReturnUpdateRequest) SetLines(v []MerchantReturnLineUpdateRequest)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


