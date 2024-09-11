# \OrdersAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderAcknowledge**](OrdersAPI.md#OrderAcknowledge) | **Post** /v2/orders/acknowledge | Acknowledges orders
[**OrderGetByFilter**](OrdersAPI.md#OrderGetByFilter) | **Get** /v2/orders | Gets orders by filter
[**OrderGetNew**](OrdersAPI.md#OrderGetNew) | **Get** /v2/orders/new | Gets new orders
[**OrderInvoice**](OrdersAPI.md#OrderInvoice) | **Get** /v2/orders/{merchantOrderNo}/invoice | Generates an order invoice
[**OrderPackingSlip**](OrdersAPI.md#OrderPackingSlip) | **Get** /v2/orders/{merchantOrderNo}/packingslip | Generates a packing slip
[**OrderUpdate**](OrdersAPI.md#OrderUpdate) | **Put** /v2/orders/comment | Updates an order comment
[**OrderUploadInvoice**](OrdersAPI.md#OrderUploadInvoice) | **Post** /v2/orders/{merchantOrderNo}/invoice | Uploads an order invoice
[**OrderUploadInvoiceAsString**](OrdersAPI.md#OrderUploadInvoiceAsString) | **Post** /v2/orders/{merchantOrderNo}/invoice-base64 | Uploads an order invoice PDF from Base64 string.



## OrderAcknowledge

> ApiResponse OrderAcknowledge(ctx).MerchantOrderAcknowledgementRequest(merchantOrderAcknowledgementRequest).Execute()

Acknowledges orders



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
	merchantOrderAcknowledgementRequest := *openapiclient.NewMerchantOrderAcknowledgementRequest("MerchantOrderNo_example", int32(123)) // MerchantOrderAcknowledgementRequest | Relations between the id's returned by ChannelEngine and the references which the merchant uses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderAcknowledge(context.Background()).MerchantOrderAcknowledgementRequest(merchantOrderAcknowledgementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderAcknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderAcknowledge`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderAcknowledge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantOrderAcknowledgementRequest** | [**MerchantOrderAcknowledgementRequest**](MerchantOrderAcknowledgementRequest.md) | Relations between the id&#39;s returned by ChannelEngine and the references which the merchant uses. | 

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


## OrderGetByFilter

> CollectionOfMerchantOrderResponse OrderGetByFilter(ctx).Statuses(statuses).EmailAddresses(emailAddresses).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).CommercialOrderNos(commercialOrderNos).FromDate(fromDate).ToDate(toDate).FromCreatedAtDate(fromCreatedAtDate).ToCreatedAtDate(toCreatedAtDate).ExcludeMarketplaceFulfilledOrdersAndLines(excludeMarketplaceFulfilledOrdersAndLines).FulfillmentType(fulfillmentType).OnlyWithCancellationRequests(onlyWithCancellationRequests).ChannelIds(channelIds).StockLocationIds(stockLocationIds).IsAcknowledged(isAcknowledged).FromUpdatedAtDate(fromUpdatedAtDate).ToUpdatedAtDate(toUpdatedAtDate).FromAcknowledgedDate(fromAcknowledgedDate).ToAcknowledgedDate(toAcknowledgedDate).FromClosedAtDate(fromClosedAtDate).ToClosedAtDate(toClosedAtDate).Page(page).Execute()

Gets orders by filter



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
	statuses := []openapiclient.OrderStatusView{openapiclient.OrderStatusView("IN_PROGRESS")} // []OrderStatusView | Order status(es) to filter on. AWAITING_PAYMENT orders will be excluded if it is not included in this Statuses filter. (optional)
	emailAddresses := []string{"Inner_example"} // []string | Client emailaddresses to filter on. (optional)
	merchantOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the merchant. (optional)
	channelOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the channel. (optional)
	commercialOrderNos := []string{"Inner_example"} // []string | Filter on commercial order numbers. (optional)
	fromDate := time.Now() // time.Time | Filter on the order date, starting from this date. This date is inclusive.<br />The order date is based on the data we got from a channel. (optional)
	toDate := time.Now() // time.Time | Filter on the order date, until this date. This date is exclusive.<br />The order date is based on the data we got from a channel. (optional)
	fromCreatedAtDate := time.Now() // time.Time | Filter on the created at date in ChannelEngine, starting from this date. This date is inclusive.<br />The created date is set on the date and time when the order is created. (optional)
	toCreatedAtDate := time.Now() // time.Time | Filter on the created at date in ChannelEngine, until this date. This date is exclusive.<br />The created date is set on the date and time when the order is created. (optional)
	excludeMarketplaceFulfilledOrdersAndLines := true // bool | Exclude order (lines) fulfilled by the marketplace (amazon:FBA, bol:LVB, etc.) (optional)
	fulfillmentType := openapiclient.FulfillmentType("ALL") // FulfillmentType | Filter orders on fulfillment type. This will include all orders lines, even if they are partially fulfilled by the marketplace.<br />To exclude orders and lines that are fulfilled by the marketplace from the response, set ExcludeMarketplaceFulfilledOrdersAndLines to true. (optional)
	onlyWithCancellationRequests := true // bool | Filter on orders containing cancellation requests.<br />Some channels allow a customer to cancel an order until it has been shipped.<br />When an order has already been acknowledged in ChannelEngine, it can only be cancelled by creating a cancellation. (optional)
	channelIds := []int32{int32(123)} // []int32 | Filter orders on channel(s). (optional)
	stockLocationIds := []int32{int32(123)} // []int32 | Filter on stock locations (optional)
	isAcknowledged := true // bool | Filter on acknowledged value (optional)
	fromUpdatedAtDate := time.Now() // time.Time | Filter on the order update date, starting from this date. This date is inclusive. (optional)
	toUpdatedAtDate := time.Now() // time.Time | Filter on the order update date, unitl this date. This date is exclusive. (optional)
	fromAcknowledgedDate := time.Now() // time.Time | Filter on the order acknowledged date, starting from this date. This date is inclusive. (optional)
	toAcknowledgedDate := time.Now() // time.Time | Filter on the order acknowledged date, until this date. This date is exclusive. (optional)
	fromClosedAtDate := time.Now() // time.Time | Filter on the order ClosedAt date, starting from this date. This date is inclusive. (optional)
	toClosedAtDate := time.Now() // time.Time | Filter on the order ClosedAt date, until this date. This date is exclusive. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderGetByFilter(context.Background()).Statuses(statuses).EmailAddresses(emailAddresses).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).CommercialOrderNos(commercialOrderNos).FromDate(fromDate).ToDate(toDate).FromCreatedAtDate(fromCreatedAtDate).ToCreatedAtDate(toCreatedAtDate).ExcludeMarketplaceFulfilledOrdersAndLines(excludeMarketplaceFulfilledOrdersAndLines).FulfillmentType(fulfillmentType).OnlyWithCancellationRequests(onlyWithCancellationRequests).ChannelIds(channelIds).StockLocationIds(stockLocationIds).IsAcknowledged(isAcknowledged).FromUpdatedAtDate(fromUpdatedAtDate).ToUpdatedAtDate(toUpdatedAtDate).FromAcknowledgedDate(fromAcknowledgedDate).ToAcknowledgedDate(toAcknowledgedDate).FromClosedAtDate(fromClosedAtDate).ToClosedAtDate(toClosedAtDate).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderGetByFilter`: CollectionOfMerchantOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | [**[]OrderStatusView**](OrderStatusView.md) | Order status(es) to filter on. AWAITING_PAYMENT orders will be excluded if it is not included in this Statuses filter. | 
 **emailAddresses** | **[]string** | Client emailaddresses to filter on. | 
 **merchantOrderNos** | **[]string** | Filter on unique order reference used by the merchant. | 
 **channelOrderNos** | **[]string** | Filter on unique order reference used by the channel. | 
 **commercialOrderNos** | **[]string** | Filter on commercial order numbers. | 
 **fromDate** | **time.Time** | Filter on the order date, starting from this date. This date is inclusive.&lt;br /&gt;The order date is based on the data we got from a channel. | 
 **toDate** | **time.Time** | Filter on the order date, until this date. This date is exclusive.&lt;br /&gt;The order date is based on the data we got from a channel. | 
 **fromCreatedAtDate** | **time.Time** | Filter on the created at date in ChannelEngine, starting from this date. This date is inclusive.&lt;br /&gt;The created date is set on the date and time when the order is created. | 
 **toCreatedAtDate** | **time.Time** | Filter on the created at date in ChannelEngine, until this date. This date is exclusive.&lt;br /&gt;The created date is set on the date and time when the order is created. | 
 **excludeMarketplaceFulfilledOrdersAndLines** | **bool** | Exclude order (lines) fulfilled by the marketplace (amazon:FBA, bol:LVB, etc.) | 
 **fulfillmentType** | [**FulfillmentType**](FulfillmentType.md) | Filter orders on fulfillment type. This will include all orders lines, even if they are partially fulfilled by the marketplace.&lt;br /&gt;To exclude orders and lines that are fulfilled by the marketplace from the response, set ExcludeMarketplaceFulfilledOrdersAndLines to true. | 
 **onlyWithCancellationRequests** | **bool** | Filter on orders containing cancellation requests.&lt;br /&gt;Some channels allow a customer to cancel an order until it has been shipped.&lt;br /&gt;When an order has already been acknowledged in ChannelEngine, it can only be cancelled by creating a cancellation. | 
 **channelIds** | **[]int32** | Filter orders on channel(s). | 
 **stockLocationIds** | **[]int32** | Filter on stock locations | 
 **isAcknowledged** | **bool** | Filter on acknowledged value | 
 **fromUpdatedAtDate** | **time.Time** | Filter on the order update date, starting from this date. This date is inclusive. | 
 **toUpdatedAtDate** | **time.Time** | Filter on the order update date, unitl this date. This date is exclusive. | 
 **fromAcknowledgedDate** | **time.Time** | Filter on the order acknowledged date, starting from this date. This date is inclusive. | 
 **toAcknowledgedDate** | **time.Time** | Filter on the order acknowledged date, until this date. This date is exclusive. | 
 **fromClosedAtDate** | **time.Time** | Filter on the order ClosedAt date, starting from this date. This date is inclusive. | 
 **toClosedAtDate** | **time.Time** | Filter on the order ClosedAt date, until this date. This date is exclusive. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantOrderResponse**](CollectionOfMerchantOrderResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetNew

> CollectionOfMerchantOrderResponse OrderGetNew(ctx).StockLocationId(stockLocationId).Execute()

Gets new orders



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
	stockLocationId := int32(56) // int32 | The ChannelEngine id of the stock location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderGetNew(context.Background()).StockLocationId(stockLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderGetNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderGetNew`: CollectionOfMerchantOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderGetNew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLocationId** | **int32** | The ChannelEngine id of the stock location. | 

### Return type

[**CollectionOfMerchantOrderResponse**](CollectionOfMerchantOrderResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderInvoice

> *os.File OrderInvoice(ctx, merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()

Generates an order invoice



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
	merchantOrderNo := "merchantOrderNo_example" // string | The unique order reference as used by the merchant.
	useCustomerCulture := true // bool | Generate the invoice in the billing address' country's language. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderInvoice(context.Background(), merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderInvoice`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** | The unique order reference as used by the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useCustomerCulture** | **bool** | Generate the invoice in the billing address&#39; country&#39;s language. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPackingSlip

> *os.File OrderPackingSlip(ctx, merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()

Generates a packing slip



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
	merchantOrderNo := "merchantOrderNo_example" // string | The unique order reference as used by the merchant.
	useCustomerCulture := true // bool | Generate the invoice in the billing address' country's language. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderPackingSlip(context.Background(), merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderPackingSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderPackingSlip`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderPackingSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** | The unique order reference as used by the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderPackingSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useCustomerCulture** | **bool** | Generate the invoice in the billing address&#39; country&#39;s language. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderUpdate

> ApiResponse OrderUpdate(ctx).MerchantOrderCommentUpdateRequest(merchantOrderCommentUpdateRequest).Execute()

Updates an order comment



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
	merchantOrderCommentUpdateRequest := *openapiclient.NewMerchantOrderCommentUpdateRequest("MerchantComment_example") // MerchantOrderCommentUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderUpdate(context.Background()).MerchantOrderCommentUpdateRequest(merchantOrderCommentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderUpdate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantOrderCommentUpdateRequest** | [**MerchantOrderCommentUpdateRequest**](MerchantOrderCommentUpdateRequest.md) |  | 

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


## OrderUploadInvoice

> ApiResponse OrderUploadInvoice(ctx, merchantOrderNo).Invoice(invoice).InvoiceNumber(invoiceNumber).Execute()

Uploads an order invoice



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
	merchantOrderNo := "merchantOrderNo_example" // string | The unique order reference as used by the merchant.
	invoice := os.NewFile(1234, "some_file") // *os.File | PDF invoice file up to 1 MB with additional data.
	invoiceNumber := "invoiceNumber_example" // string | The invoice number used in the invoice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderUploadInvoice(context.Background(), merchantOrderNo).Invoice(invoice).InvoiceNumber(invoiceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderUploadInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderUploadInvoice`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderUploadInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** | The unique order reference as used by the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderUploadInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoice** | ***os.File** | PDF invoice file up to 1 MB with additional data. | 
 **invoiceNumber** | **string** | The invoice number used in the invoice. | 

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


## OrderUploadInvoiceAsString

> ApiResponse OrderUploadInvoiceAsString(ctx, merchantOrderNo).MerchantInvoiceUploadRequest(merchantInvoiceUploadRequest).Execute()

Uploads an order invoice PDF from Base64 string.



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
	merchantOrderNo := "merchantOrderNo_example" // string | 
	merchantInvoiceUploadRequest := *openapiclient.NewMerchantInvoiceUploadRequest("InvoiceContent_example", "InvoiceNumber_example") // MerchantInvoiceUploadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderUploadInvoiceAsString(context.Background(), merchantOrderNo).MerchantInvoiceUploadRequest(merchantInvoiceUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderUploadInvoiceAsString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderUploadInvoiceAsString`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderUploadInvoiceAsString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderUploadInvoiceAsStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantInvoiceUploadRequest** | [**MerchantInvoiceUploadRequest**](MerchantInvoiceUploadRequest.md) |  | 

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

