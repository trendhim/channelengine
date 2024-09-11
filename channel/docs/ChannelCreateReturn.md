# ChannelCreateReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderIdentifier** | Pointer to **NullableString** |  | [optional] 
**ChannelReturnNo** | Pointer to **NullableString** |  | [optional] 
**Handled** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to [**ModuleReturnReason**](ModuleReturnReason.md) |  | [optional] 
**ReturnDate** | Pointer to **time.Time** |  | [optional] 
**ExtraData** | Pointer to **map[string]string** |  | [optional] 
**Lines** | Pointer to [**[]ChannelCreateReturnLine**](ChannelCreateReturnLine.md) |  | [optional] 

## Methods

### NewChannelCreateReturn

`func NewChannelCreateReturn() *ChannelCreateReturn`

NewChannelCreateReturn instantiates a new ChannelCreateReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelCreateReturnWithDefaults

`func NewChannelCreateReturnWithDefaults() *ChannelCreateReturn`

NewChannelCreateReturnWithDefaults instantiates a new ChannelCreateReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderIdentifier

`func (o *ChannelCreateReturn) GetOrderIdentifier() string`

GetOrderIdentifier returns the OrderIdentifier field if non-nil, zero value otherwise.

### GetOrderIdentifierOk

`func (o *ChannelCreateReturn) GetOrderIdentifierOk() (*string, bool)`

GetOrderIdentifierOk returns a tuple with the OrderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIdentifier

`func (o *ChannelCreateReturn) SetOrderIdentifier(v string)`

SetOrderIdentifier sets OrderIdentifier field to given value.

### HasOrderIdentifier

`func (o *ChannelCreateReturn) HasOrderIdentifier() bool`

HasOrderIdentifier returns a boolean if a field has been set.

### SetOrderIdentifierNil

`func (o *ChannelCreateReturn) SetOrderIdentifierNil(b bool)`

 SetOrderIdentifierNil sets the value for OrderIdentifier to be an explicit nil

### UnsetOrderIdentifier
`func (o *ChannelCreateReturn) UnsetOrderIdentifier()`

UnsetOrderIdentifier ensures that no value is present for OrderIdentifier, not even an explicit nil
### GetChannelReturnNo

`func (o *ChannelCreateReturn) GetChannelReturnNo() string`

GetChannelReturnNo returns the ChannelReturnNo field if non-nil, zero value otherwise.

### GetChannelReturnNoOk

`func (o *ChannelCreateReturn) GetChannelReturnNoOk() (*string, bool)`

GetChannelReturnNoOk returns a tuple with the ChannelReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReturnNo

`func (o *ChannelCreateReturn) SetChannelReturnNo(v string)`

SetChannelReturnNo sets ChannelReturnNo field to given value.

### HasChannelReturnNo

`func (o *ChannelCreateReturn) HasChannelReturnNo() bool`

HasChannelReturnNo returns a boolean if a field has been set.

### SetChannelReturnNoNil

`func (o *ChannelCreateReturn) SetChannelReturnNoNil(b bool)`

 SetChannelReturnNoNil sets the value for ChannelReturnNo to be an explicit nil

### UnsetChannelReturnNo
`func (o *ChannelCreateReturn) UnsetChannelReturnNo()`

UnsetChannelReturnNo ensures that no value is present for ChannelReturnNo, not even an explicit nil
### GetHandled

`func (o *ChannelCreateReturn) GetHandled() bool`

GetHandled returns the Handled field if non-nil, zero value otherwise.

### GetHandledOk

`func (o *ChannelCreateReturn) GetHandledOk() (*bool, bool)`

GetHandledOk returns a tuple with the Handled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandled

`func (o *ChannelCreateReturn) SetHandled(v bool)`

SetHandled sets Handled field to given value.

### HasHandled

`func (o *ChannelCreateReturn) HasHandled() bool`

HasHandled returns a boolean if a field has been set.

### GetReason

`func (o *ChannelCreateReturn) GetReason() ModuleReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChannelCreateReturn) GetReasonOk() (*ModuleReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChannelCreateReturn) SetReason(v ModuleReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChannelCreateReturn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReturnDate

`func (o *ChannelCreateReturn) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *ChannelCreateReturn) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *ChannelCreateReturn) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *ChannelCreateReturn) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### GetExtraData

`func (o *ChannelCreateReturn) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ChannelCreateReturn) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ChannelCreateReturn) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ChannelCreateReturn) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *ChannelCreateReturn) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *ChannelCreateReturn) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetLines

`func (o *ChannelCreateReturn) GetLines() []ChannelCreateReturnLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelCreateReturn) GetLinesOk() (*[]ChannelCreateReturnLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelCreateReturn) SetLines(v []ChannelCreateReturnLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *ChannelCreateReturn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *ChannelCreateReturn) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *ChannelCreateReturn) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


