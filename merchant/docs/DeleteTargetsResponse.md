# DeleteTargetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedTargets** | Pointer to [**[]DeleteTargetResponseVm**](DeleteTargetResponseVm.md) |  | [optional] 
**NotExistingTargets** | Pointer to [**[]DeleteTargetResponseVm**](DeleteTargetResponseVm.md) |  | [optional] 

## Methods

### NewDeleteTargetsResponse

`func NewDeleteTargetsResponse() *DeleteTargetsResponse`

NewDeleteTargetsResponse instantiates a new DeleteTargetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTargetsResponseWithDefaults

`func NewDeleteTargetsResponseWithDefaults() *DeleteTargetsResponse`

NewDeleteTargetsResponseWithDefaults instantiates a new DeleteTargetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedTargets

`func (o *DeleteTargetsResponse) GetDeletedTargets() []DeleteTargetResponseVm`

GetDeletedTargets returns the DeletedTargets field if non-nil, zero value otherwise.

### GetDeletedTargetsOk

`func (o *DeleteTargetsResponse) GetDeletedTargetsOk() (*[]DeleteTargetResponseVm, bool)`

GetDeletedTargetsOk returns a tuple with the DeletedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTargets

`func (o *DeleteTargetsResponse) SetDeletedTargets(v []DeleteTargetResponseVm)`

SetDeletedTargets sets DeletedTargets field to given value.

### HasDeletedTargets

`func (o *DeleteTargetsResponse) HasDeletedTargets() bool`

HasDeletedTargets returns a boolean if a field has been set.

### SetDeletedTargetsNil

`func (o *DeleteTargetsResponse) SetDeletedTargetsNil(b bool)`

 SetDeletedTargetsNil sets the value for DeletedTargets to be an explicit nil

### UnsetDeletedTargets
`func (o *DeleteTargetsResponse) UnsetDeletedTargets()`

UnsetDeletedTargets ensures that no value is present for DeletedTargets, not even an explicit nil
### GetNotExistingTargets

`func (o *DeleteTargetsResponse) GetNotExistingTargets() []DeleteTargetResponseVm`

GetNotExistingTargets returns the NotExistingTargets field if non-nil, zero value otherwise.

### GetNotExistingTargetsOk

`func (o *DeleteTargetsResponse) GetNotExistingTargetsOk() (*[]DeleteTargetResponseVm, bool)`

GetNotExistingTargetsOk returns a tuple with the NotExistingTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotExistingTargets

`func (o *DeleteTargetsResponse) SetNotExistingTargets(v []DeleteTargetResponseVm)`

SetNotExistingTargets sets NotExistingTargets field to given value.

### HasNotExistingTargets

`func (o *DeleteTargetsResponse) HasNotExistingTargets() bool`

HasNotExistingTargets returns a boolean if a field has been set.

### SetNotExistingTargetsNil

`func (o *DeleteTargetsResponse) SetNotExistingTargetsNil(b bool)`

 SetNotExistingTargetsNil sets the value for NotExistingTargets to be an explicit nil

### UnsetNotExistingTargets
`func (o *DeleteTargetsResponse) UnsetNotExistingTargets()`

UnsetNotExistingTargets ensures that no value is present for NotExistingTargets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


