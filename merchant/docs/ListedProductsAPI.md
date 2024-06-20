# \ListedProductsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListedProductGetByFilter**](ListedProductsAPI.md#ListedProductGetByFilter) | **Get** /v2/channels/{channelId}/products | Gets products listed by channel



## ListedProductGetByFilter

> CollectionOfChannelListedProductResponse ListedProductGetByFilter(ctx, channelId).MerchantProductNos(merchantProductNos).Page(page).Execute()

Gets products listed by channel



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
	channelId := int32(56) // int32 | The id of a channel
	merchantProductNos := []string{"Inner_example"} // []string | The unique product references used by the Merchant (SKUs) (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListedProductsAPI.ListedProductGetByFilter(context.Background(), channelId).MerchantProductNos(merchantProductNos).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListedProductsAPI.ListedProductGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListedProductGetByFilter`: CollectionOfChannelListedProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ListedProductsAPI.ListedProductGetByFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int32** | The id of a channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiListedProductGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantProductNos** | **[]string** | The unique product references used by the Merchant (SKUs) | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfChannelListedProductResponse**](CollectionOfChannelListedProductResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

