# \ProductBundleAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductBundleGetByFilter**](ProductBundleAPI.md#ProductBundleGetByFilter) | **Get** /v2/productbundles | Gets product bundles



## ProductBundleGetByFilter

> CollectionOfMerchantProductBundleResponse ProductBundleGetByFilter(ctx).Search(search).EanList(eanList).MerchantProductNoList(merchantProductNoList).Page(page).Execute()

Gets product bundles



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
    search := "search_example" // string | Search product(s) by Name, MerchantProductNo, Ean or Brand<br />This search is applied to the result after applying the other filters. (optional)
    eanList := []string{"Inner_example"} // []string | Search products by submitting a list of EAN's. (optional)
    merchantProductNoList := []string{"Inner_example"} // []string | Search products by submitting a list of MerchantProductNo's. (optional)
    page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductBundleAPI.ProductBundleGetByFilter(context.Background()).Search(search).EanList(eanList).MerchantProductNoList(merchantProductNoList).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductBundleAPI.ProductBundleGetByFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductBundleGetByFilter`: CollectionOfMerchantProductBundleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductBundleAPI.ProductBundleGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBundleGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search product(s) by Name, MerchantProductNo, Ean or Brand&lt;br /&gt;This search is applied to the result after applying the other filters. | 
 **eanList** | **[]string** | Search products by submitting a list of EAN&#39;s. | 
 **merchantProductNoList** | **[]string** | Search products by submitting a list of MerchantProductNo&#39;s. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantProductBundleResponse**](CollectionOfMerchantProductBundleResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

