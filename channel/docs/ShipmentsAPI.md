# \ShipmentsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShipmentIndex**](ShipmentsAPI.md#ShipmentIndex) | **Get** /v2/shipments | Gets shipments
[**ShipmentShippingLabel**](ShipmentsAPI.md#ShipmentShippingLabel) | **Get** /v2/orders/{merchantShipmentNo}/shippinglabel | Gets a shipping label



## ShipmentIndex

> CollectionOfChannelShipmentResponse ShipmentIndex(ctx).FromDate(fromDate).ToDate(toDate).ChannelOrderNos(channelOrderNos).Page(page).Execute()

Gets shipments



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
	fromDate := time.Now() // time.Time | Filter on the creation date, starting from this date. This date is inclusive. (optional)
	toDate := time.Now() // time.Time | Filter on the creation date, until this date. This date is exclusive. (optional)
	channelOrderNos := []string{"Inner_example"} // []string | Filter on the unique references (ids) as used by the channel. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentsAPI.ShipmentIndex(context.Background()).FromDate(fromDate).ToDate(toDate).ChannelOrderNos(channelOrderNos).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsAPI.ShipmentIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShipmentIndex`: CollectionOfChannelShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ShipmentsAPI.ShipmentIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShipmentIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **time.Time** | Filter on the creation date, starting from this date. This date is inclusive. | 
 **toDate** | **time.Time** | Filter on the creation date, until this date. This date is exclusive. | 
 **channelOrderNos** | **[]string** | Filter on the unique references (ids) as used by the channel. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfChannelShipmentResponse**](CollectionOfChannelShipmentResponse.md)

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
	openapiclient "github.com/trendhim/channelengine/channel"
)

func main() {
	merchantShipmentNo := "merchantShipmentNo_example" // string | The unique shipment reference as used by the merchant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentsAPI.ShipmentShippingLabel(context.Background(), merchantShipmentNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsAPI.ShipmentShippingLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShipmentShippingLabel`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ShipmentsAPI.ShipmentShippingLabel`: %v\n", resp)
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

