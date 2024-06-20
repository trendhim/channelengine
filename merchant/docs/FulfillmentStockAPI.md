# \FulfillmentStockAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FulfillmentStockGetFulfillementStockWithStockLocations**](FulfillmentStockAPI.md#FulfillmentStockGetFulfillementStockWithStockLocations) | **Get** /v2/fulfillmentstock | Gets product stock across all warehouses with stock locations



## FulfillmentStockGetFulfillementStockWithStockLocations

> CollectionOfMerchantFulfillmentStockStockLocationsResponse FulfillmentStockGetFulfillementStockWithStockLocations(ctx).MerchantProductNos(merchantProductNos).PageIndex(pageIndex).PageSize(pageSize).Execute()

Gets product stock across all warehouses with stock locations



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
	merchantProductNos := []string{"Inner_example"} // []string | List of your products' MerchantProductNo's. (optional)
	pageIndex := int32(56) // int32 | A page index to get the items (starts from 0) (optional)
	pageSize := int32(56) // int32 | Number of items to return (default 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FulfillmentStockAPI.FulfillmentStockGetFulfillementStockWithStockLocations(context.Background()).MerchantProductNos(merchantProductNos).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentStockAPI.FulfillmentStockGetFulfillementStockWithStockLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FulfillmentStockGetFulfillementStockWithStockLocations`: CollectionOfMerchantFulfillmentStockStockLocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `FulfillmentStockAPI.FulfillmentStockGetFulfillementStockWithStockLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFulfillmentStockGetFulfillementStockWithStockLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantProductNos** | **[]string** | List of your products&#39; MerchantProductNo&#39;s. | 
 **pageIndex** | **int32** | A page index to get the items (starts from 0) | 
 **pageSize** | **int32** | Number of items to return (default 100) | 

### Return type

[**CollectionOfMerchantFulfillmentStockStockLocationsResponse**](CollectionOfMerchantFulfillmentStockStockLocationsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

