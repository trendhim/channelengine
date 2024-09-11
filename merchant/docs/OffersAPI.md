# \OffersAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfferGetStock**](OffersAPI.md#OfferGetStock) | **Get** /v2/offer/stock | Gets product stock across all warehouses
[**OfferStockPriceUpdate**](OffersAPI.md#OfferStockPriceUpdate) | **Put** /v2/offer | Updates stock and price
[**OfferStockUpdate**](OffersAPI.md#OfferStockUpdate) | **Put** /v2/offer/stock | Updates stock



## OfferGetStock

> CollectionOfMerchantOfferGetStockResponse OfferGetStock(ctx).StockLocationIds(stockLocationIds).Skus(skus).PageIndex(pageIndex).PageSize(pageSize).Execute()

Gets product stock across all warehouses



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
	stockLocationIds := []int32{int32(123)} // []int32 | The ChannelEngine id of the stock location(s).
	skus := []string{"Inner_example"} // []string | List of your products' sku's. (optional)
	pageIndex := int32(56) // int32 | A page index to get the items (starts from 0) (optional)
	pageSize := int32(56) // int32 | Number of items to return (default 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OffersAPI.OfferGetStock(context.Background()).StockLocationIds(stockLocationIds).Skus(skus).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OffersAPI.OfferGetStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfferGetStock`: CollectionOfMerchantOfferGetStockResponse
	fmt.Fprintf(os.Stdout, "Response from `OffersAPI.OfferGetStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfferGetStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLocationIds** | **[]int32** | The ChannelEngine id of the stock location(s). | 
 **skus** | **[]string** | List of your products&#39; sku&#39;s. | 
 **pageIndex** | **int32** | A page index to get the items (starts from 0) | 
 **pageSize** | **int32** | Number of items to return (default 100) | 

### Return type

[**CollectionOfMerchantOfferGetStockResponse**](CollectionOfMerchantOfferGetStockResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfferStockPriceUpdate

> SingleOfDictionaryOfStringAndListOfString OfferStockPriceUpdate(ctx).MerchantStockPriceUpdateRequest(merchantStockPriceUpdateRequest).Execute()

Updates stock and price



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
	merchantStockPriceUpdateRequest := []openapiclient.MerchantStockPriceUpdateRequest{*openapiclient.NewMerchantStockPriceUpdateRequest("MerchantProductNo_example")} // []MerchantStockPriceUpdateRequest | References to the products that should be updated, and the new values<br />for the stock or price fields. It is possible to supply only one of the two fields<br />or both.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OffersAPI.OfferStockPriceUpdate(context.Background()).MerchantStockPriceUpdateRequest(merchantStockPriceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OffersAPI.OfferStockPriceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfferStockPriceUpdate`: SingleOfDictionaryOfStringAndListOfString
	fmt.Fprintf(os.Stdout, "Response from `OffersAPI.OfferStockPriceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfferStockPriceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantStockPriceUpdateRequest** | [**[]MerchantStockPriceUpdateRequest**](MerchantStockPriceUpdateRequest.md) | References to the products that should be updated, and the new values&lt;br /&gt;for the stock or price fields. It is possible to supply only one of the two fields&lt;br /&gt;or both. | 

### Return type

[**SingleOfDictionaryOfStringAndListOfString**](SingleOfDictionaryOfStringAndListOfString.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfferStockUpdate

> SingleOfDictionaryOfStringAndListOfString OfferStockUpdate(ctx).MerchantOfferStockUpdateRequest(merchantOfferStockUpdateRequest).Execute()

Updates stock



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
	merchantOfferStockUpdateRequest := []openapiclient.MerchantOfferStockUpdateRequest{*openapiclient.NewMerchantOfferStockUpdateRequest("MerchantProductNo_example", []openapiclient.MerchantStockLocationUpdateRequest{*openapiclient.NewMerchantStockLocationUpdateRequest()})} // []MerchantOfferStockUpdateRequest | References to the new values for the stock fields.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OffersAPI.OfferStockUpdate(context.Background()).MerchantOfferStockUpdateRequest(merchantOfferStockUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OffersAPI.OfferStockUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfferStockUpdate`: SingleOfDictionaryOfStringAndListOfString
	fmt.Fprintf(os.Stdout, "Response from `OffersAPI.OfferStockUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfferStockUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantOfferStockUpdateRequest** | [**[]MerchantOfferStockUpdateRequest**](MerchantOfferStockUpdateRequest.md) | References to the new values for the stock fields. | 

### Return type

[**SingleOfDictionaryOfStringAndListOfString**](SingleOfDictionaryOfStringAndListOfString.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

