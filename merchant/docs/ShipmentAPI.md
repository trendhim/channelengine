# \ShipmentAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShipmentCreate**](ShipmentAPI.md#ShipmentCreate) | **Post** /v2/shipments | Creates shipments
[**ShipmentCreateForChannelMethod**](ShipmentAPI.md#ShipmentCreateForChannelMethod) | **Post** /v2/shipments/channelmethod | Creates a shipment and initiates shipping label generation
[**ShipmentGetShipmentLabelCarriers**](ShipmentAPI.md#ShipmentGetShipmentLabelCarriers) | **Post** /v2/carriers/{merchantOrderNo} | Gets carriers providing shipping labels
[**ShipmentIndex**](ShipmentAPI.md#ShipmentIndex) | **Get** /v2/shipments/merchant | Gets shipments by filter
[**ShipmentShippingLabel**](ShipmentAPI.md#ShipmentShippingLabel) | **Get** /v2/orders/{merchantShipmentNo}/shippinglabel | Gets a shipping label
[**ShipmentUpdate**](ShipmentAPI.md#ShipmentUpdate) | **Put** /v2/shipments/{merchantShipmentNo} | Updates a shipment



## ShipmentCreate

> ApiResponse ShipmentCreate(ctx).MerchantShipmentRequest(merchantShipmentRequest).Execute()

Creates shipments



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
    merchantShipmentRequest := *openapiclient.NewMerchantShipmentRequest("MerchantShipmentNo_example", "MerchantOrderNo_example", []openapiclient.MerchantShipmentLineRequest{*openapiclient.NewMerchantShipmentLineRequest("MerchantProductNo_example", int32(123))}) // MerchantShipmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentCreate(context.Background()).MerchantShipmentRequest(merchantShipmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShipmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantShipmentRequest** | [**MerchantShipmentRequest**](MerchantShipmentRequest.md) |  | 

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


## ShipmentCreateForChannelMethod

> ApiResponse ShipmentCreateForChannelMethod(ctx).MerchantChannelLabelShipmentRequest(merchantChannelLabelShipmentRequest).Execute()

Creates a shipment and initiates shipping label generation



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
    merchantChannelLabelShipmentRequest := *openapiclient.NewMerchantChannelLabelShipmentRequest(*openapiclient.NewMerchantShipmentPackageDimensionsRequest(), *openapiclient.NewMerchantShipmentPackageWeightRequest(), "ChannelMethodCode_example", "MerchantShipmentNo_example", "MerchantOrderNo_example", []openapiclient.MerchantShipmentLineRequest{*openapiclient.NewMerchantShipmentLineRequest("MerchantProductNo_example", int32(123))}) // MerchantChannelLabelShipmentRequest | The shipment to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentCreateForChannelMethod(context.Background()).MerchantChannelLabelShipmentRequest(merchantChannelLabelShipmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentCreateForChannelMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentCreateForChannelMethod`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentCreateForChannelMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShipmentCreateForChannelMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantChannelLabelShipmentRequest** | [**MerchantChannelLabelShipmentRequest**](MerchantChannelLabelShipmentRequest.md) | The shipment to create | 

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


## ShipmentGetShipmentLabelCarriers

> CollectionOfMerchantShipmentLabelCarrierResponse ShipmentGetShipmentLabelCarriers(ctx, merchantOrderNo).MerchantShipmentLabelCarrierRequest(merchantShipmentLabelCarrierRequest).Execute()

Gets carriers providing shipping labels



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
    merchantOrderNo := "merchantOrderNo_example" // string | The merchant's order reference.
    merchantShipmentLabelCarrierRequest := *openapiclient.NewMerchantShipmentLabelCarrierRequest([]openapiclient.MerchantShipmentLineRequest{*openapiclient.NewMerchantShipmentLineRequest("MerchantProductNo_example", int32(123))}, *openapiclient.NewMerchantShipmentPackageDimensionsRequest(), *openapiclient.NewMerchantShipmentPackageWeightRequest()) // MerchantShipmentLabelCarrierRequest | The parcel information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentGetShipmentLabelCarriers(context.Background(), merchantOrderNo).MerchantShipmentLabelCarrierRequest(merchantShipmentLabelCarrierRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentGetShipmentLabelCarriers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentGetShipmentLabelCarriers`: CollectionOfMerchantShipmentLabelCarrierResponse
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentGetShipmentLabelCarriers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantOrderNo** | **string** | The merchant&#39;s order reference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentGetShipmentLabelCarriersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantShipmentLabelCarrierRequest** | [**MerchantShipmentLabelCarrierRequest**](MerchantShipmentLabelCarrierRequest.md) | The parcel information | 

### Return type

[**CollectionOfMerchantShipmentLabelCarrierResponse**](CollectionOfMerchantShipmentLabelCarrierResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentIndex

> CollectionOfMerchantShipmentResponse ShipmentIndex(ctx).MerchantShipmentNos(merchantShipmentNos).MerchantOrderNos(merchantOrderNos).Method(method).ShippedFromCountryCodes(shippedFromCountryCodes).FromShipmentDate(fromShipmentDate).ToShipmentDate(toShipmentDate).FromCreateDate(fromCreateDate).ToCreateDate(toCreateDate).FromUpdateDate(fromUpdateDate).ToUpdateDate(toUpdateDate).FulfillmentType(fulfillmentType).ChannelShipmentNos(channelShipmentNos).ChannelOrderNos(channelOrderNos).Page(page).Execute()

Gets shipments by filter



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
    merchantShipmentNos := []string{"Inner_example"} // []string | Filter on the unique references (ids) as used by the merchant. (optional)
    merchantOrderNos := []string{"Inner_example"} // []string | Filter on the unique references (ids) of order as used by the merchant. (optional)
    method := "method_example" // string | Filter on the shipping method. (optional)
    shippedFromCountryCodes := []string{"Inner_example"} // []string | 2-digit Country code (optional)
    fromShipmentDate := time.Now() // time.Time | Filter on the shipment date, starting from this date. This date is inclusive. (optional)
    toShipmentDate := time.Now() // time.Time | Filter on the shipment date, until this date. This date is exclusive. (optional)
    fromCreateDate := time.Now() // time.Time | Filter on the create date of the shipment in ChannelEngine, starting from this date. This date is inclusive. (optional)
    toCreateDate := time.Now() // time.Time | Filter on the create date of the shipment in ChannelEngine, until this date. This date is exclusive. (optional)
    fromUpdateDate := time.Now() // time.Time | Filter on the update date of the shipment in ChannelEngine, starting from this date. This date is inclusive. (optional)
    toUpdateDate := time.Now() // time.Time | Filter on the update date of the shipment in ChannelEngine, until this date. This date is exclusive. (optional)
    fulfillmentType := openapiclient.ShipmentFulfillmentType("ALL") // ShipmentFulfillmentType | Filter on the fulfillment type of the shipment. (optional)
    channelShipmentNos := []string{"Inner_example"} // []string | Filter on the unique references (ids) as used by the channel. (optional)
    channelOrderNos := []string{"Inner_example"} // []string | Filter on the unique references (ids) of order as used by the channel. (optional)
    page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentIndex(context.Background()).MerchantShipmentNos(merchantShipmentNos).MerchantOrderNos(merchantOrderNos).Method(method).ShippedFromCountryCodes(shippedFromCountryCodes).FromShipmentDate(fromShipmentDate).ToShipmentDate(toShipmentDate).FromCreateDate(fromCreateDate).ToCreateDate(toCreateDate).FromUpdateDate(fromUpdateDate).ToUpdateDate(toUpdateDate).FulfillmentType(fulfillmentType).ChannelShipmentNos(channelShipmentNos).ChannelOrderNos(channelOrderNos).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentIndex`: CollectionOfMerchantShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShipmentIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantShipmentNos** | **[]string** | Filter on the unique references (ids) as used by the merchant. | 
 **merchantOrderNos** | **[]string** | Filter on the unique references (ids) of order as used by the merchant. | 
 **method** | **string** | Filter on the shipping method. | 
 **shippedFromCountryCodes** | **[]string** | 2-digit Country code | 
 **fromShipmentDate** | **time.Time** | Filter on the shipment date, starting from this date. This date is inclusive. | 
 **toShipmentDate** | **time.Time** | Filter on the shipment date, until this date. This date is exclusive. | 
 **fromCreateDate** | **time.Time** | Filter on the create date of the shipment in ChannelEngine, starting from this date. This date is inclusive. | 
 **toCreateDate** | **time.Time** | Filter on the create date of the shipment in ChannelEngine, until this date. This date is exclusive. | 
 **fromUpdateDate** | **time.Time** | Filter on the update date of the shipment in ChannelEngine, starting from this date. This date is inclusive. | 
 **toUpdateDate** | **time.Time** | Filter on the update date of the shipment in ChannelEngine, until this date. This date is exclusive. | 
 **fulfillmentType** | [**ShipmentFulfillmentType**](ShipmentFulfillmentType.md) | Filter on the fulfillment type of the shipment. | 
 **channelShipmentNos** | **[]string** | Filter on the unique references (ids) as used by the channel. | 
 **channelOrderNos** | **[]string** | Filter on the unique references (ids) of order as used by the channel. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantShipmentResponse**](CollectionOfMerchantShipmentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentShippingLabel

> *os.File ShipmentShippingLabel(ctx, merchantShipmentNo).Execute()

Gets a shipping label



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
    merchantShipmentNo := "merchantShipmentNo_example" // string | The unique shipment reference as used by the merchant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentShippingLabel(context.Background(), merchantShipmentNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentShippingLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentShippingLabel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentShippingLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantShipmentNo** | **string** | The unique shipment reference as used by the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentShippingLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.shippingLabel, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentUpdate

> ApiResponse ShipmentUpdate(ctx, merchantShipmentNo).MerchantShipmentTrackingRequest(merchantShipmentTrackingRequest).Execute()

Updates a shipment



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
    merchantShipmentNo := "merchantShipmentNo_example" // string | The merchant's shipment reference.
    merchantShipmentTrackingRequest := *openapiclient.NewMerchantShipmentTrackingRequest("Method_example", "TrackTraceNo_example") // MerchantShipmentTrackingRequest | The updated tracking information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentAPI.ShipmentUpdate(context.Background(), merchantShipmentNo).MerchantShipmentTrackingRequest(merchantShipmentTrackingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentUpdate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantShipmentNo** | **string** | The merchant&#39;s shipment reference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantShipmentTrackingRequest** | [**MerchantShipmentTrackingRequest**](MerchantShipmentTrackingRequest.md) | The updated tracking information. | 

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

