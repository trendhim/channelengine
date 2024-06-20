# ChannelCancellation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCancellationNo** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to [**CancelReason**](CancelReason.md) |  | [optional] 
**IsForcedCancellation** | Pointer to **bool** |  | [optional] 
**Lines** | Pointer to [**[]ChannelCancellationLine**](ChannelCancellationLine.md) |  | [optional] 

## Methods

### NewChannelCancellation

`func NewChannelCancellation() *ChannelCancellation`

NewChannelCancellation instantiates a new ChannelCancellation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelCancellationWithDefaults

`func NewChannelCancellationWithDefaults() *ChannelCancellation`

NewChannelCancellationWithDefaults instantiates a new ChannelCancellation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCancellationNo

`func (o *ChannelCancellation) GetChannelCancellationNo() string`

GetChannelCancellationNo returns the ChannelCancellationNo field if non-nil, zero value otherwise.

### GetChannelCancellationNoOk

`func (o *ChannelCancellation) GetChannelCancellationNoOk() (*string, bool)`

GetChannelCancellationNoOk returns a tuple with the ChannelCancellationNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCancellationNo

`func (o *ChannelCancellation) SetChannelCancellationNo(v string)`

SetChannelCancellationNo sets ChannelCancellationNo field to given value.

### HasChannelCancellationNo

`func (o *ChannelCancellation) HasChannelCancellationNo() bool`

HasChannelCancellationNo returns a boolean if a field has been set.

### SetChannelCancellationNoNil

`func (o *ChannelCancellation) SetChannelCancellationNoNil(b bool)`

 SetChannelCancellationNoNil sets the value for ChannelCancellationNo to be an explicit nil

### UnsetChannelCancellationNo
`func (o *ChannelCancellation) UnsetChannelCancellationNo()`

UnsetChannelCancellationNo ensures that no value is present for ChannelCancellationNo, not even an explicit nil
### GetIdentifier

`func (o *ChannelCancellation) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ChannelCancellation) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ChannelCancellation) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ChannelCancellation) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *ChannelCancellation) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *ChannelCancellation) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetReason

`func (o *ChannelCancellation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChannelCancellation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChannelCancellation) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChannelCancellation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ChannelCancellation) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ChannelCancellation) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *ChannelCancellation) GetReasonCode() CancelReason`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ChannelCancellation) GetReasonCodeOk() (*CancelReason, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ChannelCancellation) SetReasonCode(v CancelReason)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ChannelCancellation) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetIsForcedCancellation

`func (o *ChannelCancellation) GetIsForcedCancellation() bool`

GetIsForcedCancellation returns the IsForcedCancellation field if non-nil, zero value otherwise.

### GetIsForcedCancellationOk

`func (o *ChannelCancellation) GetIsForcedCancellationOk() (*bool, bool)`

GetIsForcedCancellationOk returns a tuple with the IsForcedCancellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcedCancellation

`func (o *ChannelCancellation) SetIsForcedCancellation(v bool)`

SetIsForcedCancellation sets IsForcedCancellation field to given value.

### HasIsForcedCancellation

`func (o *ChannelCancellation) HasIsForcedCancellation() bool`

HasIsForcedCancellation returns a boolean if a field has been set.

### GetLines

`func (o *ChannelCancellation) GetLines() []ChannelCancellationLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ChannelCancellation) GetLinesOk() (*[]ChannelCancellationLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ChannelCancellation) SetLines(v []ChannelCancellationLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *ChannelCancellation) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *ChannelCancellation) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *ChannelCancellation) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


