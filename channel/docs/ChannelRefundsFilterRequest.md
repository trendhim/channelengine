# ChannelRefundsFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | Pointer to [**BulkRefundIdentifierRequest**](BulkRefundIdentifierRequest.md) |  | [optional] 
**ChannelExportStatus** | Pointer to [**ChannelExportStatusRequest**](ChannelExportStatusRequest.md) |  | [optional] 
**Reasons** | Pointer to [**[]RefundReason**](RefundReason.md) |  | [optional] 
**CreatedDateRange** | Pointer to [**DateRangeRequest**](DateRangeRequest.md) |  | [optional] 
**ChannelIds** | Pointer to **[]int32** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**IsAcknowledgedByMerchant** | Pointer to **NullableBool** |  | [optional] 
**IsAcknowledgedByChannel** | Pointer to **NullableBool** |  | [optional] 
**FulfillmentType** | Pointer to [**ModuleFulfillmentType**](ModuleFulfillmentType.md) |  | [optional] 
**CreatorType** | Pointer to [**CreatorType**](CreatorType.md) |  | [optional] 
**ExternalBatchNos** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChannelRefundsFilterRequest

`func NewChannelRefundsFilterRequest() *ChannelRefundsFilterRequest`

NewChannelRefundsFilterRequest instantiates a new ChannelRefundsFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelRefundsFilterRequestWithDefaults

`func NewChannelRefundsFilterRequestWithDefaults() *ChannelRefundsFilterRequest`

NewChannelRefundsFilterRequestWithDefaults instantiates a new ChannelRefundsFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *ChannelRefundsFilterRequest) GetIdentifiers() BulkRefundIdentifierRequest`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ChannelRefundsFilterRequest) GetIdentifiersOk() (*BulkRefundIdentifierRequest, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ChannelRefundsFilterRequest) SetIdentifiers(v BulkRefundIdentifierRequest)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ChannelRefundsFilterRequest) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetChannelExportStatus

`func (o *ChannelRefundsFilterRequest) GetChannelExportStatus() ChannelExportStatusRequest`

GetChannelExportStatus returns the ChannelExportStatus field if non-nil, zero value otherwise.

### GetChannelExportStatusOk

`func (o *ChannelRefundsFilterRequest) GetChannelExportStatusOk() (*ChannelExportStatusRequest, bool)`

GetChannelExportStatusOk returns a tuple with the ChannelExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelExportStatus

`func (o *ChannelRefundsFilterRequest) SetChannelExportStatus(v ChannelExportStatusRequest)`

SetChannelExportStatus sets ChannelExportStatus field to given value.

### HasChannelExportStatus

`func (o *ChannelRefundsFilterRequest) HasChannelExportStatus() bool`

HasChannelExportStatus returns a boolean if a field has been set.

### GetReasons

`func (o *ChannelRefundsFilterRequest) GetReasons() []RefundReason`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ChannelRefundsFilterRequest) GetReasonsOk() (*[]RefundReason, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ChannelRefundsFilterRequest) SetReasons(v []RefundReason)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ChannelRefundsFilterRequest) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *ChannelRefundsFilterRequest) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *ChannelRefundsFilterRequest) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil
### GetCreatedDateRange

`func (o *ChannelRefundsFilterRequest) GetCreatedDateRange() DateRangeRequest`

GetCreatedDateRange returns the CreatedDateRange field if non-nil, zero value otherwise.

### GetCreatedDateRangeOk

`func (o *ChannelRefundsFilterRequest) GetCreatedDateRangeOk() (*DateRangeRequest, bool)`

GetCreatedDateRangeOk returns a tuple with the CreatedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateRange

`func (o *ChannelRefundsFilterRequest) SetCreatedDateRange(v DateRangeRequest)`

SetCreatedDateRange sets CreatedDateRange field to given value.

### HasCreatedDateRange

`func (o *ChannelRefundsFilterRequest) HasCreatedDateRange() bool`

HasCreatedDateRange returns a boolean if a field has been set.

### GetChannelIds

`func (o *ChannelRefundsFilterRequest) GetChannelIds() []int32`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *ChannelRefundsFilterRequest) GetChannelIdsOk() (*[]int32, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *ChannelRefundsFilterRequest) SetChannelIds(v []int32)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *ChannelRefundsFilterRequest) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *ChannelRefundsFilterRequest) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *ChannelRefundsFilterRequest) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil
### GetSearch

`func (o *ChannelRefundsFilterRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ChannelRefundsFilterRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ChannelRefundsFilterRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ChannelRefundsFilterRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ChannelRefundsFilterRequest) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ChannelRefundsFilterRequest) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetIsAcknowledgedByMerchant

`func (o *ChannelRefundsFilterRequest) GetIsAcknowledgedByMerchant() bool`

GetIsAcknowledgedByMerchant returns the IsAcknowledgedByMerchant field if non-nil, zero value otherwise.

### GetIsAcknowledgedByMerchantOk

`func (o *ChannelRefundsFilterRequest) GetIsAcknowledgedByMerchantOk() (*bool, bool)`

GetIsAcknowledgedByMerchantOk returns a tuple with the IsAcknowledgedByMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgedByMerchant

`func (o *ChannelRefundsFilterRequest) SetIsAcknowledgedByMerchant(v bool)`

SetIsAcknowledgedByMerchant sets IsAcknowledgedByMerchant field to given value.

### HasIsAcknowledgedByMerchant

`func (o *ChannelRefundsFilterRequest) HasIsAcknowledgedByMerchant() bool`

HasIsAcknowledgedByMerchant returns a boolean if a field has been set.

### SetIsAcknowledgedByMerchantNil

`func (o *ChannelRefundsFilterRequest) SetIsAcknowledgedByMerchantNil(b bool)`

 SetIsAcknowledgedByMerchantNil sets the value for IsAcknowledgedByMerchant to be an explicit nil

### UnsetIsAcknowledgedByMerchant
`func (o *ChannelRefundsFilterRequest) UnsetIsAcknowledgedByMerchant()`

UnsetIsAcknowledgedByMerchant ensures that no value is present for IsAcknowledgedByMerchant, not even an explicit nil
### GetIsAcknowledgedByChannel

`func (o *ChannelRefundsFilterRequest) GetIsAcknowledgedByChannel() bool`

GetIsAcknowledgedByChannel returns the IsAcknowledgedByChannel field if non-nil, zero value otherwise.

### GetIsAcknowledgedByChannelOk

`func (o *ChannelRefundsFilterRequest) GetIsAcknowledgedByChannelOk() (*bool, bool)`

GetIsAcknowledgedByChannelOk returns a tuple with the IsAcknowledgedByChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgedByChannel

`func (o *ChannelRefundsFilterRequest) SetIsAcknowledgedByChannel(v bool)`

SetIsAcknowledgedByChannel sets IsAcknowledgedByChannel field to given value.

### HasIsAcknowledgedByChannel

`func (o *ChannelRefundsFilterRequest) HasIsAcknowledgedByChannel() bool`

HasIsAcknowledgedByChannel returns a boolean if a field has been set.

### SetIsAcknowledgedByChannelNil

`func (o *ChannelRefundsFilterRequest) SetIsAcknowledgedByChannelNil(b bool)`

 SetIsAcknowledgedByChannelNil sets the value for IsAcknowledgedByChannel to be an explicit nil

### UnsetIsAcknowledgedByChannel
`func (o *ChannelRefundsFilterRequest) UnsetIsAcknowledgedByChannel()`

UnsetIsAcknowledgedByChannel ensures that no value is present for IsAcknowledgedByChannel, not even an explicit nil
### GetFulfillmentType

`func (o *ChannelRefundsFilterRequest) GetFulfillmentType() ModuleFulfillmentType`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *ChannelRefundsFilterRequest) GetFulfillmentTypeOk() (*ModuleFulfillmentType, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *ChannelRefundsFilterRequest) SetFulfillmentType(v ModuleFulfillmentType)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *ChannelRefundsFilterRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetCreatorType

`func (o *ChannelRefundsFilterRequest) GetCreatorType() CreatorType`

GetCreatorType returns the CreatorType field if non-nil, zero value otherwise.

### GetCreatorTypeOk

`func (o *ChannelRefundsFilterRequest) GetCreatorTypeOk() (*CreatorType, bool)`

GetCreatorTypeOk returns a tuple with the CreatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorType

`func (o *ChannelRefundsFilterRequest) SetCreatorType(v CreatorType)`

SetCreatorType sets CreatorType field to given value.

### HasCreatorType

`func (o *ChannelRefundsFilterRequest) HasCreatorType() bool`

HasCreatorType returns a boolean if a field has been set.

### GetExternalBatchNos

`func (o *ChannelRefundsFilterRequest) GetExternalBatchNos() []string`

GetExternalBatchNos returns the ExternalBatchNos field if non-nil, zero value otherwise.

### GetExternalBatchNosOk

`func (o *ChannelRefundsFilterRequest) GetExternalBatchNosOk() (*[]string, bool)`

GetExternalBatchNosOk returns a tuple with the ExternalBatchNos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBatchNos

`func (o *ChannelRefundsFilterRequest) SetExternalBatchNos(v []string)`

SetExternalBatchNos sets ExternalBatchNos field to given value.

### HasExternalBatchNos

`func (o *ChannelRefundsFilterRequest) HasExternalBatchNos() bool`

HasExternalBatchNos returns a boolean if a field has been set.

### SetExternalBatchNosNil

`func (o *ChannelRefundsFilterRequest) SetExternalBatchNosNil(b bool)`

 SetExternalBatchNosNil sets the value for ExternalBatchNos to be an explicit nil

### UnsetExternalBatchNos
`func (o *ChannelRefundsFilterRequest) UnsetExternalBatchNos()`

UnsetExternalBatchNos ensures that no value is present for ExternalBatchNos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


