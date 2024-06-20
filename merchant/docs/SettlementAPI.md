# \SettlementAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettlementGetByFilter**](SettlementAPI.md#SettlementGetByFilter) | **Get** /v2/settlements | Gets settlements
[**SettlementUploadSettlement**](SettlementAPI.md#SettlementUploadSettlement) | **Post** /v2/settlements/upload | Imports a settlement file.



## SettlementGetByFilter

> CollectionOfMerchantSettlementReportsResponse SettlementGetByFilter(ctx).UnExportedOnly(unExportedOnly).SettlementIds(settlementIds).ChannelSettlementNos(channelSettlementNos).ChannelIds(channelIds).FromStartDate(fromStartDate).ToStartDate(toStartDate).FromEndDate(fromEndDate).ToEndDate(toEndDate).FromCreateDate(fromCreateDate).ToCreateDate(toCreateDate).FromUpdateDate(fromUpdateDate).ToUpdateDate(toUpdateDate).Page(page).Execute()

Gets settlements



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
	unExportedOnly := true // bool | Filter on settlements that have not been exported before. (optional)
	settlementIds := []int32{int32(123)} // []int32 | Filter on settlement IDs. (optional)
	channelSettlementNos := []string{"Inner_example"} // []string | Filter on channel settlement nos. (optional)
	channelIds := []int32{int32(123)} // []int32 | Filter on channel id list. (optional)
	fromStartDate := time.Now() // time.Time | Filter on the start date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	toStartDate := time.Now() // time.Time | Filter on the start date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	fromEndDate := time.Now() // time.Time | Filter on the end date of the settlement in ChannelEngine, starting from this date. This date is inclusive. (optional)
	toEndDate := time.Now() // time.Time | Filter on the end date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	fromCreateDate := time.Now() // time.Time | Filter on the create date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	toCreateDate := time.Now() // time.Time | Filter on the create date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	fromUpdateDate := time.Now() // time.Time | Filter on the update date of the settlement in ChannelEngine, starting from this date. This date is inclusive. (optional)
	toUpdateDate := time.Now() // time.Time | Filter on the update date of the settlement in ChannelEngine, until this date. This date is exclusive. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.SettlementGetByFilter(context.Background()).UnExportedOnly(unExportedOnly).SettlementIds(settlementIds).ChannelSettlementNos(channelSettlementNos).ChannelIds(channelIds).FromStartDate(fromStartDate).ToStartDate(toStartDate).FromEndDate(fromEndDate).ToEndDate(toEndDate).FromCreateDate(fromCreateDate).ToCreateDate(toCreateDate).FromUpdateDate(fromUpdateDate).ToUpdateDate(toUpdateDate).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettlementGetByFilter`: CollectionOfMerchantSettlementReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.SettlementGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unExportedOnly** | **bool** | Filter on settlements that have not been exported before. | 
 **settlementIds** | **[]int32** | Filter on settlement IDs. | 
 **channelSettlementNos** | **[]string** | Filter on channel settlement nos. | 
 **channelIds** | **[]int32** | Filter on channel id list. | 
 **fromStartDate** | **time.Time** | Filter on the start date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **toStartDate** | **time.Time** | Filter on the start date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **fromEndDate** | **time.Time** | Filter on the end date of the settlement in ChannelEngine, starting from this date. This date is inclusive. | 
 **toEndDate** | **time.Time** | Filter on the end date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **fromCreateDate** | **time.Time** | Filter on the create date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **toCreateDate** | **time.Time** | Filter on the create date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **fromUpdateDate** | **time.Time** | Filter on the update date of the settlement in ChannelEngine, starting from this date. This date is inclusive. | 
 **toUpdateDate** | **time.Time** | Filter on the update date of the settlement in ChannelEngine, until this date. This date is exclusive. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantSettlementReportsResponse**](CollectionOfMerchantSettlementReportsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettlementUploadSettlement

> ApiResponse SettlementUploadSettlement(ctx).Settlement(settlement).ChannelId(channelId).Execute()

Imports a settlement file.



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
	settlement := os.NewFile(1234, "some_file") // *os.File | Settlement file up to 1 MB with additional data.  Format should be the one that the channel expects.
	channelId := int32(56) // int32 | The channel ID for which the settlement is for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.SettlementUploadSettlement(context.Background()).Settlement(settlement).ChannelId(channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementUploadSettlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettlementUploadSettlement`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.SettlementUploadSettlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementUploadSettlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settlement** | ***os.File** | Settlement file up to 1 MB with additional data.  Format should be the one that the channel expects. | 
 **channelId** | **int32** | The channel ID for which the settlement is for. | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

