# \ReportAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportCreateSettlementsReport**](ReportAPI.md#ReportCreateSettlementsReport) | **Post** /v2/reports/settlements | Creates a settlement report
[**ReportGetReport**](ReportAPI.md#ReportGetReport) | **Get** /v2/reports/{reportId} | Gets a settlement report
[**ReportGetStatus**](ReportAPI.md#ReportGetStatus) | **Get** /v2/reports/{reportId}/status | Gets the status of a settlement report



## ReportCreateSettlementsReport

> MerchantCreateReportResponse ReportCreateSettlementsReport(ctx).MerchantCreateSettlementsReportRequest(merchantCreateSettlementsReportRequest).Execute()

Creates a settlement report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
    merchantCreateSettlementsReportRequest := *openapiclient.NewMerchantCreateSettlementsReportRequest([]int32{int32(123)}, openapiclient.ReportType("SUMMARY")) // MerchantCreateSettlementsReportRequest | To provide settlementIds and type of report SUMMARY or DETAILED.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportAPI.ReportCreateSettlementsReport(context.Background()).MerchantCreateSettlementsReportRequest(merchantCreateSettlementsReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ReportCreateSettlementsReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportCreateSettlementsReport`: MerchantCreateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ReportCreateSettlementsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportCreateSettlementsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantCreateSettlementsReportRequest** | [**MerchantCreateSettlementsReportRequest**](MerchantCreateSettlementsReportRequest.md) | To provide settlementIds and type of report SUMMARY or DETAILED. | 

### Return type

[**MerchantCreateReportResponse**](MerchantCreateReportResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportGetReport

> *os.File ReportGetReport(ctx, reportId).Execute()

Gets a settlement report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
    reportId := "reportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportAPI.ReportGetReport(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ReportGetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportGetReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ReportGetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportGetStatus

> MerchantGetReportStatusResponse ReportGetStatus(ctx, reportId).Execute()

Gets the status of a settlement report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
    reportId := "reportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportAPI.ReportGetStatus(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.ReportGetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportGetStatus`: MerchantGetReportStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportAPI.ReportGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantGetReportStatusResponse**](MerchantGetReportStatusResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

