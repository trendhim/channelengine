# MerchantGetReportStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReportStatus**](ReportStatus.md) |  | [optional] 
**ResourceUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMerchantGetReportStatusResponse

`func NewMerchantGetReportStatusResponse() *MerchantGetReportStatusResponse`

NewMerchantGetReportStatusResponse instantiates a new MerchantGetReportStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGetReportStatusResponseWithDefaults

`func NewMerchantGetReportStatusResponseWithDefaults() *MerchantGetReportStatusResponse`

NewMerchantGetReportStatusResponseWithDefaults instantiates a new MerchantGetReportStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MerchantGetReportStatusResponse) GetStatus() ReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantGetReportStatusResponse) GetStatusOk() (*ReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantGetReportStatusResponse) SetStatus(v ReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MerchantGetReportStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResourceUrl

`func (o *MerchantGetReportStatusResponse) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *MerchantGetReportStatusResponse) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *MerchantGetReportStatusResponse) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *MerchantGetReportStatusResponse) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### SetResourceUrlNil

`func (o *MerchantGetReportStatusResponse) SetResourceUrlNil(b bool)`

 SetResourceUrlNil sets the value for ResourceUrl to be an explicit nil

### UnsetResourceUrl
`func (o *MerchantGetReportStatusResponse) UnsetResourceUrl()`

UnsetResourceUrl ensures that no value is present for ResourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


