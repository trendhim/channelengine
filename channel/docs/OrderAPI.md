# \OrderAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderCreate**](OrderAPI.md#OrderCreate) | **Post** /v2/orders | Creates order
[**OrderInvoice**](OrderAPI.md#OrderInvoice) | **Get** /v2/orders/{merchantOrderNo}/invoice | Generates an order invoice
[**OrderPackingSlip**](OrderAPI.md#OrderPackingSlip) | **Get** /v2/orders/{merchantOrderNo}/packingslip | Generates a packing slip



## OrderCreate

> ApiResponse OrderCreate(ctx).ChannelOrderRequest(channelOrderRequest).Execute()

Creates order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
    channelOrderRequest := *openapiclient.NewChannelOrderRequest(*openapiclient.NewChannelAddressRequest(), *openapiclient.NewChannelAddressRequest(), "ChannelOrderNo_example", []openapiclient.ChannelOrderLineRequest{*openapiclient.NewChannelOrderLineRequest("ChannelProductNo_example", int32(123), float32(123))}, "Email_example", float32(123), "CurrencyCode_example", time.Now()) // ChannelOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.OrderCreate(context.Background()).ChannelOrderRequest(channelOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelOrderRequest** | [**ChannelOrderRequest**](ChannelOrderRequest.md) |  | 

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
    openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
    merchantOrderNo := "merchantOrderNo_example" // string | The unique order reference as used by the merchant.
    useCustomerCulture := true // bool | Generate the invoice in the billing address' country's language. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.OrderInvoice(context.Background(), merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderInvoice`: %v\n", resp)
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
    openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
    merchantOrderNo := "merchantOrderNo_example" // string | The unique order reference as used by the merchant.
    useCustomerCulture := true // bool | Generate the invoice in the billing address' country's language. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAPI.OrderPackingSlip(context.Background(), merchantOrderNo).UseCustomerCulture(useCustomerCulture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderPackingSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderPackingSlip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderPackingSlip`: %v\n", resp)
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

