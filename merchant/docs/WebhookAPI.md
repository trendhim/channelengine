# \WebhookAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksCreate**](WebhookAPI.md#WebhooksCreate) | **Post** /v2/webhooks | Creates a webhook
[**WebhooksDelete**](WebhookAPI.md#WebhooksDelete) | **Delete** /v2/webhooks/{webhookName} | Deletes a webhook
[**WebhooksGetAll**](WebhookAPI.md#WebhooksGetAll) | **Get** /v2/webhooks | Gets webhooks
[**WebhooksUpdate**](WebhookAPI.md#WebhooksUpdate) | **Put** /v2/webhooks | Updates a webhook



## WebhooksCreate

> ApiResponse WebhooksCreate(ctx).MerchantWebhookRequest(merchantWebhookRequest).Execute()

Creates a webhook



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
	merchantWebhookRequest := *openapiclient.NewMerchantWebhookRequest("Name_example", "Url_example", []openapiclient.WebhookEventType{openapiclient.WebhookEventType("ORDERS_CREATE")}) // MerchantWebhookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhooksCreate(context.Background()).MerchantWebhookRequest(merchantWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhooksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhooksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantWebhookRequest** | [**MerchantWebhookRequest**](MerchantWebhookRequest.md) |  | 

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


## WebhooksDelete

> ApiResponse WebhooksDelete(ctx, webhookName).Execute()

Deletes a webhook



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
	webhookName := "webhookName_example" // string | The unique name of a webhook you want to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhooksDelete(context.Background(), webhookName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhooksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksDelete`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhooksDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookName** | **string** | The unique name of a webhook you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksGetAll

> CollectionOfMerchantWebhookResponse WebhooksGetAll(ctx).Execute()

Gets webhooks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhooksGetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhooksGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksGetAll`: CollectionOfMerchantWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhooksGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetAllRequest struct via the builder pattern


### Return type

[**CollectionOfMerchantWebhookResponse**](CollectionOfMerchantWebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksUpdate

> ApiResponse WebhooksUpdate(ctx).MerchantWebhookRequest(merchantWebhookRequest).Execute()

Updates a webhook



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
	merchantWebhookRequest := *openapiclient.NewMerchantWebhookRequest("Name_example", "Url_example", []openapiclient.WebhookEventType{openapiclient.WebhookEventType("ORDERS_CREATE")}) // MerchantWebhookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhooksUpdate(context.Background()).MerchantWebhookRequest(merchantWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhooksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksUpdate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhooksUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantWebhookRequest** | [**MerchantWebhookRequest**](MerchantWebhookRequest.md) |  | 

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

