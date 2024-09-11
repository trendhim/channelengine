# MerchantCreateReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderIdentifier** | Pointer to **NullableString** |  | [optional] 
**MerchantReturnNo** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to [**ModuleReturnReason**](ModuleReturnReason.md) |  | [optional] 
**MerchantComment** | Pointer to **NullableString** |  | [optional] 
**ReturnDate** | Pointer to **time.Time** |  | [optional] 
**ExtraData** | Pointer to **map[string]string** |  | [optional] 
**Lines** | Pointer to [**[]MerchantCreateReturnLine**](MerchantCreateReturnLine.md) |  | [optional] 

## Methods

### NewMerchantCreateReturn

`func NewMerchantCreateReturn() *MerchantCreateReturn`

NewMerchantCreateReturn instantiates a new MerchantCreateReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateReturnWithDefaults

`func NewMerchantCreateReturnWithDefaults() *MerchantCreateReturn`

NewMerchantCreateReturnWithDefaults instantiates a new MerchantCreateReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderIdentifier

`func (o *MerchantCreateReturn) GetOrderIdentifier() string`

GetOrderIdentifier returns the OrderIdentifier field if non-nil, zero value otherwise.

### GetOrderIdentifierOk

`func (o *MerchantCreateReturn) GetOrderIdentifierOk() (*string, bool)`

GetOrderIdentifierOk returns a tuple with the OrderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIdentifier

`func (o *MerchantCreateReturn) SetOrderIdentifier(v string)`

SetOrderIdentifier sets OrderIdentifier field to given value.

### HasOrderIdentifier

`func (o *MerchantCreateReturn) HasOrderIdentifier() bool`

HasOrderIdentifier returns a boolean if a field has been set.

### SetOrderIdentifierNil

`func (o *MerchantCreateReturn) SetOrderIdentifierNil(b bool)`

 SetOrderIdentifierNil sets the value for OrderIdentifier to be an explicit nil

### UnsetOrderIdentifier
`func (o *MerchantCreateReturn) UnsetOrderIdentifier()`

UnsetOrderIdentifier ensures that no value is present for OrderIdentifier, not even an explicit nil
### GetMerchantReturnNo

`func (o *MerchantCreateReturn) GetMerchantReturnNo() string`

GetMerchantReturnNo returns the MerchantReturnNo field if non-nil, zero value otherwise.

### GetMerchantReturnNoOk

`func (o *MerchantCreateReturn) GetMerchantReturnNoOk() (*string, bool)`

GetMerchantReturnNoOk returns a tuple with the MerchantReturnNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReturnNo

`func (o *MerchantCreateReturn) SetMerchantReturnNo(v string)`

SetMerchantReturnNo sets MerchantReturnNo field to given value.

### HasMerchantReturnNo

`func (o *MerchantCreateReturn) HasMerchantReturnNo() bool`

HasMerchantReturnNo returns a boolean if a field has been set.

### SetMerchantReturnNoNil

`func (o *MerchantCreateReturn) SetMerchantReturnNoNil(b bool)`

 SetMerchantReturnNoNil sets the value for MerchantReturnNo to be an explicit nil

### UnsetMerchantReturnNo
`func (o *MerchantCreateReturn) UnsetMerchantReturnNo()`

UnsetMerchantReturnNo ensures that no value is present for MerchantReturnNo, not even an explicit nil
### GetReason

`func (o *MerchantCreateReturn) GetReason() ModuleReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantCreateReturn) GetReasonOk() (*ModuleReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantCreateReturn) SetReason(v ModuleReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantCreateReturn) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMerchantComment

`func (o *MerchantCreateReturn) GetMerchantComment() string`

GetMerchantComment returns the MerchantComment field if non-nil, zero value otherwise.

### GetMerchantCommentOk

`func (o *MerchantCreateReturn) GetMerchantCommentOk() (*string, bool)`

GetMerchantCommentOk returns a tuple with the MerchantComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantComment

`func (o *MerchantCreateReturn) SetMerchantComment(v string)`

SetMerchantComment sets MerchantComment field to given value.

### HasMerchantComment

`func (o *MerchantCreateReturn) HasMerchantComment() bool`

HasMerchantComment returns a boolean if a field has been set.

### SetMerchantCommentNil

`func (o *MerchantCreateReturn) SetMerchantCommentNil(b bool)`

 SetMerchantCommentNil sets the value for MerchantComment to be an explicit nil

### UnsetMerchantComment
`func (o *MerchantCreateReturn) UnsetMerchantComment()`

UnsetMerchantComment ensures that no value is present for MerchantComment, not even an explicit nil
### GetReturnDate

`func (o *MerchantCreateReturn) GetReturnDate() time.Time`

GetReturnDate returns the ReturnDate field if non-nil, zero value otherwise.

### GetReturnDateOk

`func (o *MerchantCreateReturn) GetReturnDateOk() (*time.Time, bool)`

GetReturnDateOk returns a tuple with the ReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDate

`func (o *MerchantCreateReturn) SetReturnDate(v time.Time)`

SetReturnDate sets ReturnDate field to given value.

### HasReturnDate

`func (o *MerchantCreateReturn) HasReturnDate() bool`

HasReturnDate returns a boolean if a field has been set.

### GetExtraData

`func (o *MerchantCreateReturn) GetExtraData() map[string]string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *MerchantCreateReturn) GetExtraDataOk() (*map[string]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *MerchantCreateReturn) SetExtraData(v map[string]string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *MerchantCreateReturn) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *MerchantCreateReturn) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *MerchantCreateReturn) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetLines

`func (o *MerchantCreateReturn) GetLines() []MerchantCreateReturnLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MerchantCreateReturn) GetLinesOk() (*[]MerchantCreateReturnLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MerchantCreateReturn) SetLines(v []MerchantCreateReturnLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *MerchantCreateReturn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *MerchantCreateReturn) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *MerchantCreateReturn) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


