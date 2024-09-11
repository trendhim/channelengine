# \CompetitionPricesAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompetitionPricesGetBuyBoxPrices**](CompetitionPricesAPI.md#CompetitionPricesGetBuyBoxPrices) | **Get** /v2/competitionprices/buyboxprices | Gets the price from the buy box winner



## CompetitionPricesGetBuyBoxPrices

> CollectionOfMerchantProductWithBuyBoxPrice CompetitionPricesGetBuyBoxPrices(ctx).ChannelId(channelId).GtinList(gtinList).SkuList(skuList).Page(page).Execute()

Gets the price from the buy box winner



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
	channelId := int32(56) // int32 | The id of the channel (optional)
	gtinList := []string{"Inner_example"} // []string | Search products by submitting a list of GTIN's. (optional)<br />Max. 2000. (optional)
	skuList := []string{"Inner_example"} // []string | Search products by submitting a list of Sku's. (optional)<br />Max. 2000. If GtinList is already set, this one is ignored. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompetitionPricesAPI.CompetitionPricesGetBuyBoxPrices(context.Background()).ChannelId(channelId).GtinList(gtinList).SkuList(skuList).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompetitionPricesAPI.CompetitionPricesGetBuyBoxPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompetitionPricesGetBuyBoxPrices`: CollectionOfMerchantProductWithBuyBoxPrice
	fmt.Fprintf(os.Stdout, "Response from `CompetitionPricesAPI.CompetitionPricesGetBuyBoxPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompetitionPricesGetBuyBoxPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **int32** | The id of the channel | 
 **gtinList** | **[]string** | Search products by submitting a list of GTIN&#39;s. (optional)&lt;br /&gt;Max. 2000. | 
 **skuList** | **[]string** | Search products by submitting a list of Sku&#39;s. (optional)&lt;br /&gt;Max. 2000. If GtinList is already set, this one is ignored. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantProductWithBuyBoxPrice**](CollectionOfMerchantProductWithBuyBoxPrice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

