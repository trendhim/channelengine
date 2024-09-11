# \TargetsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetsCreateTargets**](TargetsAPI.md#TargetsCreateTargets) | **Post** /v2/targets | Creates multiple targets
[**TargetsDeleteTargets**](TargetsAPI.md#TargetsDeleteTargets) | **Delete** /v2/targets | Deletes multiple targets
[**TargetsEditTargets**](TargetsAPI.md#TargetsEditTargets) | **Put** /v2/targets | Edits multiple targets



## TargetsCreateTargets

> SingleOfListOfTargetResponseVm TargetsCreateTargets(ctx).CreateEditTargetRequest(createEditTargetRequest).Execute()

Creates multiple targets



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
	createEditTargetRequest := []openapiclient.CreateEditTargetRequest{*openapiclient.NewCreateEditTargetRequest()} // []CreateEditTargetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.TargetsCreateTargets(context.Background()).CreateEditTargetRequest(createEditTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.TargetsCreateTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TargetsCreateTargets`: SingleOfListOfTargetResponseVm
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.TargetsCreateTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTargetsCreateTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEditTargetRequest** | [**[]CreateEditTargetRequest**](CreateEditTargetRequest.md) |  | 

### Return type

[**SingleOfListOfTargetResponseVm**](SingleOfListOfTargetResponseVm.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TargetsDeleteTargets

> SingleOfDeleteTargetsResponse TargetsDeleteTargets(ctx).DeleteTargetRequest(deleteTargetRequest).Execute()

Deletes multiple targets



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
	deleteTargetRequest := []openapiclient.DeleteTargetRequest{*openapiclient.NewDeleteTargetRequest()} // []DeleteTargetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.TargetsDeleteTargets(context.Background()).DeleteTargetRequest(deleteTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.TargetsDeleteTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TargetsDeleteTargets`: SingleOfDeleteTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.TargetsDeleteTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTargetsDeleteTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTargetRequest** | [**[]DeleteTargetRequest**](DeleteTargetRequest.md) |  | 

### Return type

[**SingleOfDeleteTargetsResponse**](SingleOfDeleteTargetsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TargetsEditTargets

> SingleOfListOfTargetResponseVm TargetsEditTargets(ctx).CreateEditTargetRequest(createEditTargetRequest).Execute()

Edits multiple targets



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
	createEditTargetRequest := []openapiclient.CreateEditTargetRequest{*openapiclient.NewCreateEditTargetRequest()} // []CreateEditTargetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.TargetsEditTargets(context.Background()).CreateEditTargetRequest(createEditTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.TargetsEditTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TargetsEditTargets`: SingleOfListOfTargetResponseVm
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.TargetsEditTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTargetsEditTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEditTargetRequest** | [**[]CreateEditTargetRequest**](CreateEditTargetRequest.md) |  | 

### Return type

[**SingleOfListOfTargetResponseVm**](SingleOfListOfTargetResponseVm.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

