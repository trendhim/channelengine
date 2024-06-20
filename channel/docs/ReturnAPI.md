# \ReturnAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnDeclareForChannel**](ReturnAPI.md#ReturnDeclareForChannel) | **Post** /v2/returns/channel | Creates return
[**ReturnGetDeclaredByMerchant**](ReturnAPI.md#ReturnGetDeclaredByMerchant) | **Get** /v2/returns/channel | Gets returns



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
	resp, r, err := apiClient.ReturnAPI.ReturnDeclareForChannel(context.Background()).ChannelReturnRequest(channelReturnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnDeclareForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnDeclareForChannel`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnDeclareForChannel`: %v\n", resp)
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
	resp, r, err := apiClient.ReturnAPI.ReturnGetDeclaredByMerchant(context.Background()).Statuses(statuses).Reasons(reasons).FromDate(fromDate).ToDate(toDate).IsAcknowledged(isAcknowledged).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnGetDeclaredByMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnGetDeclaredByMerchant`: CollectionOfChannelReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnGetDeclaredByMerchant`: %v\n", resp)
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

