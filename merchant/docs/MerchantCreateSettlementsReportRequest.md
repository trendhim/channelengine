# MerchantCreateSettlementsReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementIds** | **[]int32** |  | 
**Type** | [**ReportType**](ReportType.md) |  | 

## Methods

### NewMerchantCreateSettlementsReportRequest

`func NewMerchantCreateSettlementsReportRequest(settlementIds []int32, type_ ReportType, ) *MerchantCreateSettlementsReportRequest`

NewMerchantCreateSettlementsReportRequest instantiates a new MerchantCreateSettlementsReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateSettlementsReportRequestWithDefaults

`func NewMerchantCreateSettlementsReportRequestWithDefaults() *MerchantCreateSettlementsReportRequest`

NewMerchantCreateSettlementsReportRequestWithDefaults instantiates a new MerchantCreateSettlementsReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementIds

`func (o *MerchantCreateSettlementsReportRequest) GetSettlementIds() []int32`

GetSettlementIds returns the SettlementIds field if non-nil, zero value otherwise.

### GetSettlementIdsOk

`func (o *MerchantCreateSettlementsReportRequest) GetSettlementIdsOk() (*[]int32, bool)`

GetSettlementIdsOk returns a tuple with the SettlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementIds

`func (o *MerchantCreateSettlementsReportRequest) SetSettlementIds(v []int32)`

SetSettlementIds sets SettlementIds field to given value.


### GetType

`func (o *MerchantCreateSettlementsReportRequest) GetType() ReportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantCreateSettlementsReportRequest) GetTypeOk() (*ReportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantCreateSettlementsReportRequest) SetType(v ReportType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


