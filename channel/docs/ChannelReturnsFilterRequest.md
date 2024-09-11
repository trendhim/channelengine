# ChannelReturnsFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | Pointer to [**BulkReturnIdentifierRequest**](BulkReturnIdentifierRequest.md) |  | [optional] 
**ChannelExportStatus** | Pointer to [**ChannelExportStatusRequest**](ChannelExportStatusRequest.md) |  | [optional] 
**Reasons** | Pointer to [**[]ModuleReturnReason**](ModuleReturnReason.md) |  | [optional] 
**CreatedDateRange** | Pointer to [**DateRangeRequest**](DateRangeRequest.md) |  | [optional] 
**Statuses** | Pointer to [**[]ModuleReturnStatus**](ModuleReturnStatus.md) |  | [optional] 
**ChannelIds** | Pointer to **[]int32** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**IsAcknowledgedByMerchant** | Pointer to **NullableBool** |  | [optional] 
**IsAcknowledgedByChannel** | Pointer to **NullableBool** |  | [optional] 
**FulfillmentType** | Pointer to [**ModuleFulfillmentType**](ModuleFulfillmentType.md) |  | [optional] 
**CreatorType** | Pointer to [**CreatorType**](CreatorType.md) |  | [optional] 
**ExternalBatchNos** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChannelReturnsFilterRequest

`func NewChannelReturnsFilterRequest() *ChannelReturnsFilterRequest`

NewChannelReturnsFilterRequest instantiates a new ChannelReturnsFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelReturnsFilterRequestWithDefaults

`func NewChannelReturnsFilterRequestWithDefaults() *ChannelReturnsFilterRequest`

NewChannelReturnsFilterRequestWithDefaults instantiates a new ChannelReturnsFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *ChannelReturnsFilterRequest) GetIdentifiers() BulkReturnIdentifierRequest`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ChannelReturnsFilterRequest) GetIdentifiersOk() (*BulkReturnIdentifierRequest, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ChannelReturnsFilterRequest) SetIdentifiers(v BulkReturnIdentifierRequest)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ChannelReturnsFilterRequest) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetChannelExportStatus

`func (o *ChannelReturnsFilterRequest) GetChannelExportStatus() ChannelExportStatusRequest`

GetChannelExportStatus returns the ChannelExportStatus field if non-nil, zero value otherwise.

### GetChannelExportStatusOk

`func (o *ChannelReturnsFilterRequest) GetChannelExportStatusOk() (*ChannelExportStatusRequest, bool)`

GetChannelExportStatusOk returns a tuple with the ChannelExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelExportStatus

`func (o *ChannelReturnsFilterRequest) SetChannelExportStatus(v ChannelExportStatusRequest)`

SetChannelExportStatus sets ChannelExportStatus field to given value.

### HasChannelExportStatus

`func (o *ChannelReturnsFilterRequest) HasChannelExportStatus() bool`

HasChannelExportStatus returns a boolean if a field has been set.

### GetReasons

`func (o *ChannelReturnsFilterRequest) GetReasons() []ModuleReturnReason`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ChannelReturnsFilterRequest) GetReasonsOk() (*[]ModuleReturnReason, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ChannelReturnsFilterRequest) SetReasons(v []ModuleReturnReason)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ChannelReturnsFilterRequest) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *ChannelReturnsFilterRequest) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *ChannelReturnsFilterRequest) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil
### GetCreatedDateRange

`func (o *ChannelReturnsFilterRequest) GetCreatedDateRange() DateRangeRequest`

GetCreatedDateRange returns the CreatedDateRange field if non-nil, zero value otherwise.

### GetCreatedDateRangeOk

`func (o *ChannelReturnsFilterRequest) GetCreatedDateRangeOk() (*DateRangeRequest, bool)`

GetCreatedDateRangeOk returns a tuple with the CreatedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateRange

`func (o *ChannelReturnsFilterRequest) SetCreatedDateRange(v DateRangeRequest)`

SetCreatedDateRange sets CreatedDateRange field to given value.

### HasCreatedDateRange

`func (o *ChannelReturnsFilterRequest) HasCreatedDateRange() bool`

HasCreatedDateRange returns a boolean if a field has been set.

### GetStatuses

`func (o *ChannelReturnsFilterRequest) GetStatuses() []ModuleReturnStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ChannelReturnsFilterRequest) GetStatusesOk() (*[]ModuleReturnStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ChannelReturnsFilterRequest) SetStatuses(v []ModuleReturnStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ChannelReturnsFilterRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *ChannelReturnsFilterRequest) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *ChannelReturnsFilterRequest) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetChannelIds

`func (o *ChannelReturnsFilterRequest) GetChannelIds() []int32`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *ChannelReturnsFilterRequest) GetChannelIdsOk() (*[]int32, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *ChannelReturnsFilterRequest) SetChannelIds(v []int32)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *ChannelReturnsFilterRequest) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *ChannelReturnsFilterRequest) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *ChannelReturnsFilterRequest) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil
### GetSearch

`func (o *ChannelReturnsFilterRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ChannelReturnsFilterRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ChannelReturnsFilterRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ChannelReturnsFilterRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ChannelReturnsFilterRequest) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ChannelReturnsFilterRequest) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetIsAcknowledgedByMerchant

`func (o *ChannelReturnsFilterRequest) GetIsAcknowledgedByMerchant() bool`

GetIsAcknowledgedByMerchant returns the IsAcknowledgedByMerchant field if non-nil, zero value otherwise.

### GetIsAcknowledgedByMerchantOk

`func (o *ChannelReturnsFilterRequest) GetIsAcknowledgedByMerchantOk() (*bool, bool)`

GetIsAcknowledgedByMerchantOk returns a tuple with the IsAcknowledgedByMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgedByMerchant

`func (o *ChannelReturnsFilterRequest) SetIsAcknowledgedByMerchant(v bool)`

SetIsAcknowledgedByMerchant sets IsAcknowledgedByMerchant field to given value.

### HasIsAcknowledgedByMerchant

`func (o *ChannelReturnsFilterRequest) HasIsAcknowledgedByMerchant() bool`

HasIsAcknowledgedByMerchant returns a boolean if a field has been set.

### SetIsAcknowledgedByMerchantNil

`func (o *ChannelReturnsFilterRequest) SetIsAcknowledgedByMerchantNil(b bool)`

 SetIsAcknowledgedByMerchantNil sets the value for IsAcknowledgedByMerchant to be an explicit nil

### UnsetIsAcknowledgedByMerchant
`func (o *ChannelReturnsFilterRequest) UnsetIsAcknowledgedByMerchant()`

UnsetIsAcknowledgedByMerchant ensures that no value is present for IsAcknowledgedByMerchant, not even an explicit nil
### GetIsAcknowledgedByChannel

`func (o *ChannelReturnsFilterRequest) GetIsAcknowledgedByChannel() bool`

GetIsAcknowledgedByChannel returns the IsAcknowledgedByChannel field if non-nil, zero value otherwise.

### GetIsAcknowledgedByChannelOk

`func (o *ChannelReturnsFilterRequest) GetIsAcknowledgedByChannelOk() (*bool, bool)`

GetIsAcknowledgedByChannelOk returns a tuple with the IsAcknowledgedByChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgedByChannel

`func (o *ChannelReturnsFilterRequest) SetIsAcknowledgedByChannel(v bool)`

SetIsAcknowledgedByChannel sets IsAcknowledgedByChannel field to given value.

### HasIsAcknowledgedByChannel

`func (o *ChannelReturnsFilterRequest) HasIsAcknowledgedByChannel() bool`

HasIsAcknowledgedByChannel returns a boolean if a field has been set.

### SetIsAcknowledgedByChannelNil

`func (o *ChannelReturnsFilterRequest) SetIsAcknowledgedByChannelNil(b bool)`

 SetIsAcknowledgedByChannelNil sets the value for IsAcknowledgedByChannel to be an explicit nil

### UnsetIsAcknowledgedByChannel
`func (o *ChannelReturnsFilterRequest) UnsetIsAcknowledgedByChannel()`

UnsetIsAcknowledgedByChannel ensures that no value is present for IsAcknowledgedByChannel, not even an explicit nil
### GetFulfillmentType

`func (o *ChannelReturnsFilterRequest) GetFulfillmentType() ModuleFulfillmentType`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *ChannelReturnsFilterRequest) GetFulfillmentTypeOk() (*ModuleFulfillmentType, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *ChannelReturnsFilterRequest) SetFulfillmentType(v ModuleFulfillmentType)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *ChannelReturnsFilterRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetCreatorType

`func (o *ChannelReturnsFilterRequest) GetCreatorType() CreatorType`

GetCreatorType returns the CreatorType field if non-nil, zero value otherwise.

### GetCreatorTypeOk

`func (o *ChannelReturnsFilterRequest) GetCreatorTypeOk() (*CreatorType, bool)`

GetCreatorTypeOk returns a tuple with the CreatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorType

`func (o *ChannelReturnsFilterRequest) SetCreatorType(v CreatorType)`

SetCreatorType sets CreatorType field to given value.

### HasCreatorType

`func (o *ChannelReturnsFilterRequest) HasCreatorType() bool`

HasCreatorType returns a boolean if a field has been set.

### GetExternalBatchNos

`func (o *ChannelReturnsFilterRequest) GetExternalBatchNos() []string`

GetExternalBatchNos returns the ExternalBatchNos field if non-nil, zero value otherwise.

### GetExternalBatchNosOk

`func (o *ChannelReturnsFilterRequest) GetExternalBatchNosOk() (*[]string, bool)`

GetExternalBatchNosOk returns a tuple with the ExternalBatchNos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBatchNos

`func (o *ChannelReturnsFilterRequest) SetExternalBatchNos(v []string)`

SetExternalBatchNos sets ExternalBatchNos field to given value.

### HasExternalBatchNos

`func (o *ChannelReturnsFilterRequest) HasExternalBatchNos() bool`

HasExternalBatchNos returns a boolean if a field has been set.

### SetExternalBatchNosNil

`func (o *ChannelReturnsFilterRequest) SetExternalBatchNosNil(b bool)`

 SetExternalBatchNosNil sets the value for ExternalBatchNos to be an explicit nil

### UnsetExternalBatchNos
`func (o *ChannelReturnsFilterRequest) UnsetExternalBatchNos()`

UnsetExternalBatchNos ensures that no value is present for ExternalBatchNos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


