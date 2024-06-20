# \CancellationAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancellationCreate**](CancellationAPI.md#CancellationCreate) | **Post** /v2/cancellations | Creates a cancelation
[**CancellationGetForMerchant**](CancellationAPI.md#CancellationGetForMerchant) | **Get** /v2/cancellations/merchant | Gets cancelations



## CancellationCreate

> ApiResponse CancellationCreate(ctx).MerchantCancellationRequest(merchantCancellationRequest).Execute()

Creates a cancelation



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
	merchantCancellationRequest := *openapiclient.NewMerchantCancellationRequest("MerchantCancellationNo_example", "MerchantOrderNo_example", []openapiclient.MerchantCancellationLineRequest{*openapiclient.NewMerchantCancellationLineRequest("MerchantProductNo_example", int32(123))}) // MerchantCancellationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CancellationAPI.CancellationCreate(context.Background()).MerchantCancellationRequest(merchantCancellationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancellationAPI.CancellationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancellationCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CancellationAPI.CancellationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantCancellationRequest** | [**MerchantCancellationRequest**](MerchantCancellationRequest.md) |  | 

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


## CancellationGetForMerchant

> CollectionOfMerchantCancellationResponse CancellationGetForMerchant(ctx).CreatedSince(createdSince).Page(page).Execute()

Gets cancelations



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
	createdSince := time.Now() // time.Time | Filter on the create date of the cancellation in ChannelEngine, starting from this date. This date is inclusive. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CancellationAPI.CancellationGetForMerchant(context.Background()).CreatedSince(createdSince).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancellationAPI.CancellationGetForMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancellationGetForMerchant`: CollectionOfMerchantCancellationResponse
	fmt.Fprintf(os.Stdout, "Response from `CancellationAPI.CancellationGetForMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationGetForMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdSince** | **time.Time** | Filter on the create date of the cancellation in ChannelEngine, starting from this date. This date is inclusive. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantCancellationResponse**](CollectionOfMerchantCancellationResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

