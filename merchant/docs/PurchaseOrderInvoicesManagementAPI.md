# \PurchaseOrderInvoicesManagementAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](PurchaseOrderInvoicesManagementAPI.md#CreateInvoice) | **Post** /v2/purchase-orders/invoice | Creates a purchase order invoice
[**CreateInvoices**](PurchaseOrderInvoicesManagementAPI.md#CreateInvoices) | **Post** /v2/purchase-orders/invoice/bulk | Creates a purchase order invoices in a bulk



## CreateInvoice

> ApiResponse CreateInvoice(ctx).SingleMerchantCreatePurchaseOrderInvoiceRequest(singleMerchantCreatePurchaseOrderInvoiceRequest).Execute()

Creates a purchase order invoice



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
	singleMerchantCreatePurchaseOrderInvoiceRequest := *openapiclient.NewSingleMerchantCreatePurchaseOrderInvoiceRequest() // SingleMerchantCreatePurchaseOrderInvoiceRequest | Model for purchase order invoice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderInvoicesManagementAPI.CreateInvoice(context.Background()).SingleMerchantCreatePurchaseOrderInvoiceRequest(singleMerchantCreatePurchaseOrderInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderInvoicesManagementAPI.CreateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoice`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderInvoicesManagementAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantCreatePurchaseOrderInvoiceRequest** | [**SingleMerchantCreatePurchaseOrderInvoiceRequest**](SingleMerchantCreatePurchaseOrderInvoiceRequest.md) | Model for purchase order invoice. | 

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


## CreateInvoices

> ApiResponse CreateInvoices(ctx).BulkMerchantCreatePurchaseOrderInvoicesRequest(bulkMerchantCreatePurchaseOrderInvoicesRequest).Execute()

Creates a purchase order invoices in a bulk



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
	bulkMerchantCreatePurchaseOrderInvoicesRequest := *openapiclient.NewBulkMerchantCreatePurchaseOrderInvoicesRequest() // BulkMerchantCreatePurchaseOrderInvoicesRequest | Model for purchase order invoices. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderInvoicesManagementAPI.CreateInvoices(context.Background()).BulkMerchantCreatePurchaseOrderInvoicesRequest(bulkMerchantCreatePurchaseOrderInvoicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderInvoicesManagementAPI.CreateInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoices`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderInvoicesManagementAPI.CreateInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkMerchantCreatePurchaseOrderInvoicesRequest** | [**BulkMerchantCreatePurchaseOrderInvoicesRequest**](BulkMerchantCreatePurchaseOrderInvoicesRequest.md) | Model for purchase order invoices. | 

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

