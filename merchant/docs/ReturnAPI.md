# \ReturnAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAcknowledge**](ReturnAPI.md#ReturnAcknowledge) | **Post** /v2/returns/merchant/acknowledge | Acknowledges a return
[**ReturnDeclareForMerchant**](ReturnAPI.md#ReturnDeclareForMerchant) | **Post** /v2/returns/merchant | Creates merchant return
[**ReturnGetByMerchantOrderNo**](ReturnAPI.md#ReturnGetByMerchantOrderNo) | **Get** /v2/returns/merchant/{merchantOrderNo} | Gets a return
[**ReturnGetDeclaredByChannel**](ReturnAPI.md#ReturnGetDeclaredByChannel) | **Get** /v2/returns/merchant | Gets marketplace returns
[**ReturnGetReturns**](ReturnAPI.md#ReturnGetReturns) | **Get** /v2/returns | Gets returns by filter
[**ReturnGetUnhandled**](ReturnAPI.md#ReturnGetUnhandled) | **Get** /v2/returns/merchant/new | Gets unhandled returns
[**ReturnUpdateForMerchant**](ReturnAPI.md#ReturnUpdateForMerchant) | **Put** /v2/returns | Marks returns as received



## ReturnAcknowledge

> ApiResponse ReturnAcknowledge(ctx).MerchantReturnAcknowledgeRequest(merchantReturnAcknowledgeRequest).Execute()

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
	resp, r, err := apiClient.ReturnAPI.ReturnAcknowledge(context.Background()).MerchantReturnAcknowledgeRequest(merchantReturnAcknowledgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnAcknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnAcknowledge`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnAcknowledge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAcknowledgeRequest struct via the builder pattern


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
	resp, r, err := apiClient.ReturnAPI.ReturnDeclareForMerchant(context.Background()).MerchantReturnRequest(merchantReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnDeclareForMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnDeclareForMerchant`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnDeclareForMerchant`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnGetByMerchantOrderNo(context.Background(), merchantOrderNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnGetByMerchantOrderNo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetByMerchantOrderNo`: CollectionOfMerchantSingleOrderReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnGetByMerchantOrderNo`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnGetDeclaredByChannel(context.Background()).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnGetDeclaredByChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetDeclaredByChannel`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnGetDeclaredByChannel`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnGetReturns(context.Background()).CreatorType(creatorType).ChannelIds(channelIds).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).FulfillmentType(fulfillmentType).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnGetReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetReturns`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnGetReturns`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnGetUnhandled(context.Background()).ChannelIds(channelIds).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnGetUnhandled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetUnhandled`: CollectionOfMerchantReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnGetUnhandled`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnUpdateForMerchant(context.Background()).MerchantReturnUpdateRequest(merchantReturnUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnUpdateForMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnUpdateForMerchant`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnUpdateForMerchant`: %v\n", resp)
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

