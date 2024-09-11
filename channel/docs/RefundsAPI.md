# \RefundsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundGet**](RefundsAPI.md#RefundGet) | **Get** /v2.1/refunds/channel/{identifier} | [CLOSED BETA] Get refund by identifier
[**RefundGetByFilter**](RefundsAPI.md#RefundGetByFilter) | **Get** /v2.1/refunds/channel | [CLOSED BETA] Get refunds by filter
[**RefundProcess**](RefundsAPI.md#RefundProcess) | **Post** /v2.1/refunds/channel/process | [CLOSED BETA] Process a refund



## RefundGet

> SingleOfIRefund RefundGet(ctx, identifier).Type_(type_).Execute()

[CLOSED BETA] Get refund by identifier



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
	type_ := openapiclient.RefundIdentifier("REFUND_ID") // RefundIdentifier | Specify whether to search by ID, Merchant Refund No or Channel Refund No (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundGet(context.Background(), identifier).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundGet`: SingleOfIRefund
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The identifier to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**RefundIdentifier**](RefundIdentifier.md) | Specify whether to search by ID, Merchant Refund No or Channel Refund No | 

### Return type

[**SingleOfIRefund**](SingleOfIRefund.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundGetByFilter

> SingleOfIRefund RefundGetByFilter(ctx).PageIndex(pageIndex).PageSize(pageSize).ChannelRefundsFilterRequest(channelRefundsFilterRequest).Execute()

[CLOSED BETA] Get refunds by filter



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
	channelRefundsFilterRequest := *openapiclient.NewChannelRefundsFilterRequest() // ChannelRefundsFilterRequest | The filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundGetByFilter(context.Background()).PageIndex(pageIndex).PageSize(pageSize).ChannelRefundsFilterRequest(channelRefundsFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundGetByFilter`: SingleOfIRefund
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **channelRefundsFilterRequest** | [**ChannelRefundsFilterRequest**](ChannelRefundsFilterRequest.md) | The filter | 

### Return type

[**SingleOfIRefund**](SingleOfIRefund.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundProcess

> ApiResponse RefundProcess(ctx).SingleChannelProcessRefundRequest(singleChannelProcessRefundRequest).Execute()

[CLOSED BETA] Process a refund



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
	singleChannelProcessRefundRequest := *openapiclient.NewSingleChannelProcessRefundRequest() // SingleChannelProcessRefundRequest | The refund to acknowledge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundProcess(context.Background()).SingleChannelProcessRefundRequest(singleChannelProcessRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundProcess`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleChannelProcessRefundRequest** | [**SingleChannelProcessRefundRequest**](SingleChannelProcessRefundRequest.md) | The refund to acknowledge | 

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

