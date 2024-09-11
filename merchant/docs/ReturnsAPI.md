# \ReturnsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAcknowledge**](ReturnsAPI.md#ReturnAcknowledge) | **Post** /v2.1/returns/merchant/acknowledge | [CLOSED BETA] Acknowledge a return
[**ReturnAcknowledge_0**](ReturnsAPI.md#ReturnAcknowledge_0) | **Post** /v2/returns/merchant/acknowledge | Acknowledges a return
[**ReturnCreate**](ReturnsAPI.md#ReturnCreate) | **Post** /v2.1/returns/merchant | [CLOSED BETA] Create a return
[**ReturnDeclareForMerchant**](ReturnsAPI.md#ReturnDeclareForMerchant) | **Post** /v2/returns/merchant | Creates merchant return
[**ReturnGet**](ReturnsAPI.md#ReturnGet) | **Get** /v2.1/returns/merchant/{identifier} | [CLOSED BETA] Get return by identifier
[**ReturnGetByFilter**](ReturnsAPI.md#ReturnGetByFilter) | **Get** /v2.1/returns/merchant | [CLOSED BETA] Get returns by filter
[**ReturnGetByMerchantOrderNo**](ReturnsAPI.md#ReturnGetByMerchantOrderNo) | **Get** /v2/returns/merchant/{merchantOrderNo} | Gets a return
[**ReturnGetDeclaredByChannel**](ReturnsAPI.md#ReturnGetDeclaredByChannel) | **Get** /v2/returns/merchant | Gets marketplace returns
[**ReturnGetReturns**](ReturnsAPI.md#ReturnGetReturns) | **Get** /v2/returns | Gets returns by filter
[**ReturnGetUnhandled**](ReturnsAPI.md#ReturnGetUnhandled) | **Get** /v2/returns/merchant/new | Gets unhandled returns
[**ReturnHandle**](ReturnsAPI.md#ReturnHandle) | **Post** /v2.1/returns/merchant/handle | [CLOSED BETA] Handle a return
[**ReturnUpdateForMerchant**](ReturnsAPI.md#ReturnUpdateForMerchant) | **Put** /v2/returns | Marks returns as received



## ReturnAcknowledge

> ApiResponse ReturnAcknowledge(ctx).SingleMerchantAcknowledgeReturnRequest(singleMerchantAcknowledgeReturnRequest).Execute()

[CLOSED BETA] Acknowledge a return



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
	singleMerchantAcknowledgeReturnRequest := *openapiclient.NewSingleMerchantAcknowledgeReturnRequest() // SingleMerchantAcknowledgeReturnRequest | The return to acknowledge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnAcknowledge(context.Background()).SingleMerchantAcknowledgeReturnRequest(singleMerchantAcknowledgeReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnAcknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnAcknowledge`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnAcknowledge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantAcknowledgeReturnRequest** | [**SingleMerchantAcknowledgeReturnRequest**](SingleMerchantAcknowledgeReturnRequest.md) | The return to acknowledge | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAcknowledge_0

> ApiResponse ReturnAcknowledge_0(ctx).MerchantReturnAcknowledgeRequest(merchantReturnAcknowledgeRequest).Execute()

Acknowledges a return



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
	merchantReturnAcknowledgeRequest := *openapiclient.NewMerchantReturnAcknowledgeRequest("MerchantReturnNo_example") // MerchantReturnAcknowledgeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnAcknowledge_0(context.Background()).MerchantReturnAcknowledgeRequest(merchantReturnAcknowledgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnAcknowledge_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnAcknowledge_0`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnAcknowledge_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAcknowledge_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantReturnAcknowledgeRequest** | [**MerchantReturnAcknowledgeRequest**](MerchantReturnAcknowledgeRequest.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCreate

> ApiResponse ReturnCreate(ctx).SingleMerchantCreateReturnRequest(singleMerchantCreateReturnRequest).Execute()

[CLOSED BETA] Create a return



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
	singleMerchantCreateReturnRequest := *openapiclient.NewSingleMerchantCreateReturnRequest() // SingleMerchantCreateReturnRequest | The return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnCreate(context.Background()).SingleMerchantCreateReturnRequest(singleMerchantCreateReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantCreateReturnRequest** | [**SingleMerchantCreateReturnRequest**](SingleMerchantCreateReturnRequest.md) | The return | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnDeclareForMerchant

> ApiResponse ReturnDeclareForMerchant(ctx).MerchantReturnRequest(merchantReturnRequest).Execute()

Creates merchant return



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
	merchantReturnRequest := *openapiclient.NewMerchantReturnRequest("MerchantOrderNo_example", "MerchantReturnNo_example", []openapiclient.MerchantReturnLineRequest{*openapiclient.NewMerchantReturnLineRequest("MerchantProductNo_example", int32(123))}) // MerchantReturnRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnDeclareForMerchant(context.Background()).MerchantReturnRequest(merchantReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnDeclareForMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnDeclareForMerchant`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnDeclareForMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnDeclareForMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantReturnRequest** | [**MerchantReturnRequest**](MerchantReturnRequest.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGet

> SingleOfIReturn ReturnGet(ctx, identifier).Type_(type_).Execute()

[CLOSED BETA] Get return by identifier



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
	identifier := "identifier_example" // string | The identifier to search for
	type_ := openapiclient.ReturnIdentifier("RETURN_ID") // ReturnIdentifier | Specify whether to search by ID, Merchant Return No or Channel Return No (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGet(context.Background(), identifier).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGet`: SingleOfIReturn
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The identifier to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ReturnIdentifier**](ReturnIdentifier.md) | Specify whether to search by ID, Merchant Return No or Channel Return No | 

### Return type

[**SingleOfIReturn**](SingleOfIReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetByFilter

> SingleOfIReturn ReturnGetByFilter(ctx).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ChannelExportStatusStatuses(channelExportStatusStatuses).ChannelExportStatusMaxNumberOfExportAttempts(channelExportStatusMaxNumberOfExportAttempts).Reasons(reasons).CreatedDateRangeFromDate(createdDateRangeFromDate).CreatedDateRangeToDate(createdDateRangeToDate).Statuses(statuses).ChannelIds(channelIds).Search(search).IsAcknowledgedByMerchant(isAcknowledgedByMerchant).IsAcknowledgedByChannel(isAcknowledgedByChannel).FulfillmentType(fulfillmentType).CreatorType(creatorType).ExternalBatchNos(externalBatchNos).PageIndex(pageIndex).PageSize(pageSize).Execute()

[CLOSED BETA] Get returns by filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
	identifiersIdentifierType := openapiclient.ReturnByFilterIdentifier("RETURN_ID") // ReturnByFilterIdentifier | The type of identifier: which identifier to filter on (optional)
	identifiersModels := []string{"Inner_example"} // []string | The value (of the selected type) to filter on (optional)
	channelExportStatusStatuses := []openapiclient.ChannelExportStatus{openapiclient.ChannelExportStatus("AWAITING_EXPORT")} // []ChannelExportStatus |  (optional)
	channelExportStatusMaxNumberOfExportAttempts := int32(56) // int32 |  (optional)
	reasons := []openapiclient.ModuleReturnReason{openapiclient.ModuleReturnReason("PRODUCT_DEFECT")} // []ModuleReturnReason |  (optional)
	createdDateRangeFromDate := time.Now() // time.Time |  (optional)
	createdDateRangeToDate := time.Now() // time.Time |  (optional)
	statuses := []openapiclient.ModuleReturnStatus{openapiclient.ModuleReturnStatus("NEW")} // []ModuleReturnStatus |  (optional)
	channelIds := []int32{int32(123)} // []int32 |  (optional)
	search := "search_example" // string |  (optional)
	isAcknowledgedByMerchant := true // bool |  (optional)
	isAcknowledgedByChannel := true // bool |  (optional)
	fulfillmentType := openapiclient.ModuleFulfillmentType("ALL") // ModuleFulfillmentType |  (optional)
	creatorType := openapiclient.CreatorType("MERCHANT") // CreatorType |  (optional)
	externalBatchNos := []string{"Inner_example"} // []string |  (optional)
	pageIndex := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetByFilter(context.Background()).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ChannelExportStatusStatuses(channelExportStatusStatuses).ChannelExportStatusMaxNumberOfExportAttempts(channelExportStatusMaxNumberOfExportAttempts).Reasons(reasons).CreatedDateRangeFromDate(createdDateRangeFromDate).CreatedDateRangeToDate(createdDateRangeToDate).Statuses(statuses).ChannelIds(channelIds).Search(search).IsAcknowledgedByMerchant(isAcknowledgedByMerchant).IsAcknowledgedByChannel(isAcknowledgedByChannel).FulfillmentType(fulfillmentType).CreatorType(creatorType).ExternalBatchNos(externalBatchNos).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetByFilter`: SingleOfIReturn
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifiersIdentifierType** | [**ReturnByFilterIdentifier**](ReturnByFilterIdentifier.md) | The type of identifier: which identifier to filter on | 
 **identifiersModels** | **[]string** | The value (of the selected type) to filter on | 
 **channelExportStatusStatuses** | [**[]ChannelExportStatus**](ChannelExportStatus.md) |  | 
 **channelExportStatusMaxNumberOfExportAttempts** | **int32** |  | 
 **reasons** | [**[]ModuleReturnReason**](ModuleReturnReason.md) |  | 
 **createdDateRangeFromDate** | **time.Time** |  | 
 **createdDateRangeToDate** | **time.Time** |  | 
 **statuses** | [**[]ModuleReturnStatus**](ModuleReturnStatus.md) |  | 
 **channelIds** | **[]int32** |  | 
 **search** | **string** |  | 
 **isAcknowledgedByMerchant** | **bool** |  | 
 **isAcknowledgedByChannel** | **bool** |  | 
 **fulfillmentType** | [**ModuleFulfillmentType**](ModuleFulfillmentType.md) |  | 
 **creatorType** | [**CreatorType**](CreatorType.md) |  | 
 **externalBatchNos** | **[]string** |  | 
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**SingleOfIReturn**](SingleOfIReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetByMerchantOrderNo

> CollectionOfMerchantSingleOrderReturnResponse ReturnGetByMerchantOrderNo(ctx, merchantOrderNo).Execute()

Gets a return



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
	merchantOrderNo := "merchantOrderNo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetByMerchantOrderNo(context.Background(), merchantOrderNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetByMerchantOrderNo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetByMerchantOrderNo`: CollectionOfMerchantSingleOrderReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetByMerchantOrderNo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetByMerchantOrderNoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionOfMerchantSingleOrderReturnResponse**](CollectionOfMerchantSingleOrderReturnResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetDeclaredByChannel

> CollectionOfMerchantReturnResponse ReturnGetDeclaredByChannel(ctx).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()

Gets marketplace returns



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
	channelIds := []int32{int32(123)} // []int32 | Filter on Channel IDs (optional)
	merchantOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the merchant. (optional)
	channelOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the channel. (optional)
	fulfillmentType := openapiclient.FulfillmentType("ALL") // FulfillmentType | Filter on the fulfillment type of the order. (optional)
	statuses := []openapiclient.ReturnStatus{openapiclient.ReturnStatus("IN_PROGRESS")} // []ReturnStatus | Return status(es) to filter on. (optional)
	reasons := []openapiclient.ReturnReason{openapiclient.ReturnReason("PRODUCT_DEFECT")} // []ReturnReason | Return reason(s) to filter on. (optional)
	fromDate := time.Now() // time.Time | Filter on the creation date, starting from this date. This date is inclusive. (optional)
	toDate := time.Now() // time.Time | Filter on the creation date, until this date. This date is exclusive. (optional)
	isAcknowledged := true // bool | Filters based on acknowledgements (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetDeclaredByChannel(context.Background()).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetDeclaredByChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetDeclaredByChannel`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetDeclaredByChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetDeclaredByChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **[]int32** | Filter on Channel IDs | 
 **merchantOrderNos** | **[]string** | Filter on unique order reference used by the merchant. | 
 **channelOrderNos** | **[]string** | Filter on unique order reference used by the channel. | 
 **fulfillmentType** | [**FulfillmentType**](FulfillmentType.md) | Filter on the fulfillment type of the order. | 
 **statuses** | [**[]ReturnStatus**](ReturnStatus.md) | Return status(es) to filter on. | 
 **reasons** | [**[]ReturnReason**](ReturnReason.md) | Return reason(s) to filter on. | 
 **fromDate** | **time.Time** | Filter on the creation date, starting from this date. This date is inclusive. | 
 **toDate** | **time.Time** | Filter on the creation date, until this date. This date is exclusive. | 
 **isAcknowledged** | **bool** | Filters based on acknowledgements | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantReturnResponse**](CollectionOfMerchantReturnResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetReturns

> CollectionOfMerchantReturnResponse ReturnGetReturns(ctx).CreatorType(creatorType).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()

Gets returns by filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/trendhim/channelengine/merchant"
)

func main() {
	creatorType := openapiclient.CreatorFilter("ONLY_FROM_MERCHANT") // CreatorFilter | Filter on the return's creator. Default is MIXED. (optional)
	channelIds := []int32{int32(123)} // []int32 | Filter on Channel IDs (optional)
	merchantOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the merchant. (optional)
	channelOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the channel. (optional)
	fulfillmentType := openapiclient.FulfillmentType("ALL") // FulfillmentType | Filter on the fulfillment type of the order. (optional)
	statuses := []openapiclient.ReturnStatus{openapiclient.ReturnStatus("IN_PROGRESS")} // []ReturnStatus | Return status(es) to filter on. (optional)
	reasons := []openapiclient.ReturnReason{openapiclient.ReturnReason("PRODUCT_DEFECT")} // []ReturnReason | Return reason(s) to filter on. (optional)
	fromDate := time.Now() // time.Time | Filter on the creation date, starting from this date. This date is inclusive. (optional)
	toDate := time.Now() // time.Time | Filter on the creation date, until this date. This date is exclusive. (optional)
	isAcknowledged := true // bool | Filters based on acknowledgements (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetReturns(context.Background()).CreatorType(creatorType).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetReturns`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetReturns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorType** | [**CreatorFilter**](CreatorFilter.md) | Filter on the return&#39;s creator. Default is MIXED. | 
 **channelIds** | **[]int32** | Filter on Channel IDs | 
 **merchantOrderNos** | **[]string** | Filter on unique order reference used by the merchant. | 
 **channelOrderNos** | **[]string** | Filter on unique order reference used by the channel. | 
 **fulfillmentType** | [**FulfillmentType**](FulfillmentType.md) | Filter on the fulfillment type of the order. | 
 **statuses** | [**[]ReturnStatus**](ReturnStatus.md) | Return status(es) to filter on. | 
 **reasons** | [**[]ReturnReason**](ReturnReason.md) | Return reason(s) to filter on. | 
 **fromDate** | **time.Time** | Filter on the creation date, starting from this date. This date is inclusive. | 
 **toDate** | **time.Time** | Filter on the creation date, until this date. This date is exclusive. | 
 **isAcknowledged** | **bool** | Filters based on acknowledgements | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantReturnResponse**](CollectionOfMerchantReturnResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetUnhandled

> CollectionOfMerchantReturnResponse ReturnGetUnhandled(ctx).ChannelIds(channelIds).Page(page).Execute()

Gets unhandled returns



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
	channelIds := []int32{int32(123)} // []int32 | Filter on Channel IDs (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetUnhandled(context.Background()).ChannelIds(channelIds).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetUnhandled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetUnhandled`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetUnhandled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetUnhandledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **[]int32** | Filter on Channel IDs | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantReturnResponse**](CollectionOfMerchantReturnResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnHandle

> ApiResponse ReturnHandle(ctx).SingleMerchantHandleReturnRequest(singleMerchantHandleReturnRequest).Execute()

[CLOSED BETA] Handle a return



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
	singleMerchantHandleReturnRequest := *openapiclient.NewSingleMerchantHandleReturnRequest() // SingleMerchantHandleReturnRequest | The return to handle (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnHandle(context.Background()).SingleMerchantHandleReturnRequest(singleMerchantHandleReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnHandle`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnHandle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantHandleReturnRequest** | [**SingleMerchantHandleReturnRequest**](SingleMerchantHandleReturnRequest.md) | The return to handle | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnUpdateForMerchant

> ApiResponse ReturnUpdateForMerchant(ctx).MerchantReturnUpdateRequest(merchantReturnUpdateRequest).Execute()

Marks returns as received



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
	merchantReturnUpdateRequest := *openapiclient.NewMerchantReturnUpdateRequest(int32(123), []openapiclient.MerchantReturnLineUpdateRequest{*openapiclient.NewMerchantReturnLineUpdateRequest("MerchantProductNo_example", int32(123), int32(123))}) // MerchantReturnUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnUpdateForMerchant(context.Background()).MerchantReturnUpdateRequest(merchantReturnUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnUpdateForMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnUpdateForMerchant`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnUpdateForMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnUpdateForMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantReturnUpdateRequest** | [**MerchantReturnUpdateRequest**](MerchantReturnUpdateRequest.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

