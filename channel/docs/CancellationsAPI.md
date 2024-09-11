# \CancellationsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancellationChannelBulkCCreate**](CancellationsAPI.md#CancellationChannelBulkCCreate) | **Post** /v2/cancellations/channel/bulk | Creates multiple cancellations
[**CancellationIndex**](CancellationsAPI.md#CancellationIndex) | **Get** /v2/cancellations | Gets cancelations
[**CancellationSingleChannelCreate**](CancellationsAPI.md#CancellationSingleChannelCreate) | **Post** /v2/cancellations/channel | Creates a cancelation



## CancellationChannelBulkCCreate

> ApiResponse CancellationChannelBulkCCreate(ctx).BulkChannelCreateCancellationsRequest(bulkChannelCreateCancellationsRequest).Execute()

Creates multiple cancellations



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
	bulkChannelCreateCancellationsRequest := *openapiclient.NewBulkChannelCreateCancellationsRequest() // BulkChannelCreateCancellationsRequest | Bulk request with specific order identifier, cancelation reason, and order line info for cancelation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CancellationsAPI.CancellationChannelBulkCCreate(context.Background()).BulkChannelCreateCancellationsRequest(bulkChannelCreateCancellationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancellationsAPI.CancellationChannelBulkCCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancellationChannelBulkCCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CancellationsAPI.CancellationChannelBulkCCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationChannelBulkCCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkChannelCreateCancellationsRequest** | [**BulkChannelCreateCancellationsRequest**](BulkChannelCreateCancellationsRequest.md) | Bulk request with specific order identifier, cancelation reason, and order line info for cancelation. | 

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


## CancellationIndex

> CollectionOfChannelCancellationResponse CancellationIndex(ctx).CreatedSince(createdSince).Execute()

Gets cancelations



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
	createdSince := time.Now() // time.Time | The date from which you will get all created cancellations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CancellationsAPI.CancellationIndex(context.Background()).CreatedSince(createdSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancellationsAPI.CancellationIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancellationIndex`: CollectionOfChannelCancellationResponse
	fmt.Fprintf(os.Stdout, "Response from `CancellationsAPI.CancellationIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdSince** | **time.Time** | The date from which you will get all created cancellations | 

### Return type

[**CollectionOfChannelCancellationResponse**](CollectionOfChannelCancellationResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancellationSingleChannelCreate

> ApiResponse CancellationSingleChannelCreate(ctx).SingleChannelCreateCancellationRequest(singleChannelCreateCancellationRequest).Execute()

Creates a cancelation



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
	singleChannelCreateCancellationRequest := *openapiclient.NewSingleChannelCreateCancellationRequest() // SingleChannelCreateCancellationRequest | Single request with specific order identifier, cancelation reason, and order line info for cancelation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CancellationsAPI.CancellationSingleChannelCreate(context.Background()).SingleChannelCreateCancellationRequest(singleChannelCreateCancellationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancellationsAPI.CancellationSingleChannelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancellationSingleChannelCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CancellationsAPI.CancellationSingleChannelCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationSingleChannelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleChannelCreateCancellationRequest** | [**SingleChannelCreateCancellationRequest**](SingleChannelCreateCancellationRequest.md) | Single request with specific order identifier, cancelation reason, and order line info for cancelation. | 

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

