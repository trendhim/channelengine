# \ReturnsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnCreate**](ReturnsAPI.md#ReturnCreate) | **Post** /v2.1/returns/channel | [CLOSED BETA] Create a return
[**ReturnDeclareForChannel**](ReturnsAPI.md#ReturnDeclareForChannel) | **Post** /v2/returns/channel | Creates return
[**ReturnGet**](ReturnsAPI.md#ReturnGet) | **Get** /v2.1/returns/channel/{identifier} | [CLOSED BETA] Get return by identifier
[**ReturnGetByFilter**](ReturnsAPI.md#ReturnGetByFilter) | **Get** /v2.1/returns/channel | [CLOSED BETA] Get returns by filter
[**ReturnGetDeclaredByMerchant**](ReturnsAPI.md#ReturnGetDeclaredByMerchant) | **Get** /v2/returns/channel | Gets returns
[**ReturnProcess**](ReturnsAPI.md#ReturnProcess) | **Post** /v2.1/returns/channel/process | [CLOSED BETA] Process a return



## ReturnCreate

> ApiResponse ReturnCreate(ctx).SingleChannelCreateReturnRequest(singleChannelCreateReturnRequest).Execute()

[CLOSED BETA] Create a return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	singleChannelCreateReturnRequest := *openapiclient.NewSingleChannelCreateReturnRequest() // SingleChannelCreateReturnRequest | The return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnCreate(context.Background()).SingleChannelCreateReturnRequest(singleChannelCreateReturnRequest).Execute()
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
 **singleChannelCreateReturnRequest** | [**SingleChannelCreateReturnRequest**](SingleChannelCreateReturnRequest.md) | The return | 

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


## ReturnDeclareForChannel

> ApiResponse ReturnDeclareForChannel(ctx).ChannelReturnRequest(channelReturnRequest).Execute()

Creates return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	channelReturnRequest := *openapiclient.NewChannelReturnRequest("ChannelReference_example", []openapiclient.ChannelReturnLineRequest{*openapiclient.NewChannelReturnLineRequest(int32(123))}) // ChannelReturnRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnDeclareForChannel(context.Background()).ChannelReturnRequest(channelReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnDeclareForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnDeclareForChannel`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnDeclareForChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnDeclareForChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelReturnRequest** | [**ChannelReturnRequest**](ChannelReturnRequest.md) |  | 

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
	openapiclient "github.com/trendhim/channelengine/channel"
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

> SingleOfIReturn ReturnGetByFilter(ctx).PageIndex(pageIndex).PageSize(pageSize).ChannelReturnsFilterRequest(channelReturnsFilterRequest).Execute()

[CLOSED BETA] Get returns by filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	pageIndex := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	channelReturnsFilterRequest := *openapiclient.NewChannelReturnsFilterRequest() // ChannelReturnsFilterRequest | The filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetByFilter(context.Background()).PageIndex(pageIndex).PageSize(pageSize).ChannelReturnsFilterRequest(channelReturnsFilterRequest).Execute()
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
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **channelReturnsFilterRequest** | [**ChannelReturnsFilterRequest**](ChannelReturnsFilterRequest.md) | The filter | 

### Return type

[**SingleOfIReturn**](SingleOfIReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnGetDeclaredByMerchant

> CollectionOfChannelReturnResponse ReturnGetDeclaredByMerchant(ctx).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()

Gets returns



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	statuses := []openapiclient.ReturnStatus{openapiclient.ReturnStatus("IN_PROGRESS")} // []ReturnStatus | Return status(es) to filter on. (optional)
	reasons := []openapiclient.ReturnReason{openapiclient.ReturnReason("PRODUCT_DEFECT")} // []ReturnReason | Return reason(s) to filter on. (optional)
	fromDate := time.Now() // time.Time | Filter on the creation date, starting from this date. This date is inclusive. (optional)
	toDate := time.Now() // time.Time | Filter on the creation date, until this date. This date is exclusive. (optional)
	isAcknowledged := true // bool | Filters based on acknowledgements (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnGetDeclaredByMerchant(context.Background()).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnGetDeclaredByMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetDeclaredByMerchant`: CollectionOfChannelReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnGetDeclaredByMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetDeclaredByMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | [**[]ReturnStatus**](ReturnStatus.md) | Return status(es) to filter on. | 
 **reasons** | [**[]ReturnReason**](ReturnReason.md) | Return reason(s) to filter on. | 
 **fromDate** | **time.Time** | Filter on the creation date, starting from this date. This date is inclusive. | 
 **toDate** | **time.Time** | Filter on the creation date, until this date. This date is exclusive. | 
 **isAcknowledged** | **bool** | Filters based on acknowledgements | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfChannelReturnResponse**](CollectionOfChannelReturnResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnProcess

> ApiResponse ReturnProcess(ctx).SingleChannelProcessReturnRequest(singleChannelProcessReturnRequest).Execute()

[CLOSED BETA] Process a return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	singleChannelProcessReturnRequest := *openapiclient.NewSingleChannelProcessReturnRequest() // SingleChannelProcessReturnRequest | The return to acknowledge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.ReturnProcess(context.Background()).SingleChannelProcessReturnRequest(singleChannelProcessReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.ReturnProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnProcess`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.ReturnProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleChannelProcessReturnRequest** | [**SingleChannelProcessReturnRequest**](SingleChannelProcessReturnRequest.md) | The return to acknowledge | 

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

