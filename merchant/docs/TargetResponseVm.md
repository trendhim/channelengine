# TargetResponseVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableInt32** |  | [optional] 
**TargetInclVat** | Pointer to **NullableFloat32** |  | [optional] 
**TargetExclVat** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 

## Methods

### NewTargetResponseVm

`func NewTargetResponseVm() *TargetResponseVm`

NewTargetResponseVm instantiates a new TargetResponseVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResponseVmWithDefaults

`func NewTargetResponseVmWithDefaults() *TargetResponseVm`

NewTargetResponseVmWithDefaults instantiates a new TargetResponseVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *TargetResponseVm) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *TargetResponseVm) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *TargetResponseVm) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *TargetResponseVm) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *TargetResponseVm) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *TargetResponseVm) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetTargetInclVat

`func (o *TargetResponseVm) GetTargetInclVat() float32`

GetTargetInclVat returns the TargetInclVat field if non-nil, zero value otherwise.

### GetTargetInclVatOk

`func (o *TargetResponseVm) GetTargetInclVatOk() (*float32, bool)`

GetTargetInclVatOk returns a tuple with the TargetInclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInclVat

`func (o *TargetResponseVm) SetTargetInclVat(v float32)`

SetTargetInclVat sets TargetInclVat field to given value.

### HasTargetInclVat

`func (o *TargetResponseVm) HasTargetInclVat() bool`

HasTargetInclVat returns a boolean if a field has been set.

### SetTargetInclVatNil

`func (o *TargetResponseVm) SetTargetInclVatNil(b bool)`

 SetTargetInclVatNil sets the value for TargetInclVat to be an explicit nil

### UnsetTargetInclVat
`func (o *TargetResponseVm) UnsetTargetInclVat()`

UnsetTargetInclVat ensures that no value is present for TargetInclVat, not even an explicit nil
### GetTargetExclVat

`func (o *TargetResponseVm) GetTargetExclVat() float32`

GetTargetExclVat returns the TargetExclVat field if non-nil, zero value otherwise.

### GetTargetExclVatOk

`func (o *TargetResponseVm) GetTargetExclVatOk() (*float32, bool)`

GetTargetExclVatOk returns a tuple with the TargetExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExclVat

`func (o *TargetResponseVm) SetTargetExclVat(v float32)`

SetTargetExclVat sets TargetExclVat field to given value.

### HasTargetExclVat

`func (o *TargetResponseVm) HasTargetExclVat() bool`

HasTargetExclVat returns a boolean if a field has been set.

### SetTargetExclVatNil

`func (o *TargetResponseVm) SetTargetExclVatNil(b bool)`

 SetTargetExclVatNil sets the value for TargetExclVat to be an explicit nil

### UnsetTargetExclVat
`func (o *TargetResponseVm) UnsetTargetExclVat()`

UnsetTargetExclVat ensures that no value is present for TargetExclVat, not even an explicit nil
### GetCreatedAt

`func (o *TargetResponseVm) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TargetResponseVm) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TargetResponseVm) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TargetResponseVm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TargetResponseVm) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TargetResponseVm) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TargetResponseVm) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TargetResponseVm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMonth

`func (o *TargetResponseVm) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *TargetResponseVm) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *TargetResponseVm) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *TargetResponseVm) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *TargetResponseVm) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TargetResponseVm) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TargetResponseVm) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *TargetResponseVm) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


