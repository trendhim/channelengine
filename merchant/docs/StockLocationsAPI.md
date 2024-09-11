# \StockLocationsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StockLocationCreate**](StockLocationsAPI.md#StockLocationCreate) | **Post** /v2/stocklocations | Creates a stock location
[**StockLocationIndex**](StockLocationsAPI.md#StockLocationIndex) | **Get** /v2/stocklocations | Gets stock locations



## StockLocationCreate

> ApiResponse StockLocationCreate(ctx).MerchantStockLocationCreateRequest(merchantStockLocationCreateRequest).Execute()

Creates a stock location



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
	merchantStockLocationCreateRequest := *openapiclient.NewMerchantStockLocationCreateRequest("Name_example") // MerchantStockLocationCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StockLocationsAPI.StockLocationCreate(context.Background()).MerchantStockLocationCreateRequest(merchantStockLocationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsAPI.StockLocationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StockLocationCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StockLocationsAPI.StockLocationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStockLocationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantStockLocationCreateRequest** | [**MerchantStockLocationCreateRequest**](MerchantStockLocationCreateRequest.md) |  | 

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


## StockLocationIndex

> CollectionOfMerchantStockLocationWithCountryIsoResponse StockLocationIndex(ctx).Execute()

Gets stock locations



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
	resp, r, err := apiClient.StockLocationsAPI.StockLocationIndex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockLocationsAPI.StockLocationIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StockLocationIndex`: CollectionOfMerchantStockLocationWithCountryIsoResponse
	fmt.Fprintf(os.Stdout, "Response from `StockLocationsAPI.StockLocationIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStockLocationIndexRequest struct via the builder pattern


### Return type

[**CollectionOfMerchantStockLocationWithCountryIsoResponse**](CollectionOfMerchantStockLocationWithCountryIsoResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

