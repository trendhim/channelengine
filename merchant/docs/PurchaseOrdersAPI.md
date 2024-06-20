# \PurchaseOrdersAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Acknowledge**](PurchaseOrdersAPI.md#Acknowledge) | **Post** /v2/purchase-orders/lines/acknowledge | Acknowledges lines of a purchase order
[**Create**](PurchaseOrdersAPI.md#Create) | **Post** /v2/purchase-orders/shipments | Create a purchase order shipment.
[**GetByFilter**](PurchaseOrdersAPI.md#GetByFilter) | **Get** /v2/purchase-orders/shipments/merchant | Gets purchase order shipments by filter
[**GetByFilter_0**](PurchaseOrdersAPI.md#GetByFilter_0) | **Get** /v2/purchase-orders | Gets purchase orders by filter
[**Update**](PurchaseOrdersAPI.md#Update) | **Put** /v2/purchase-orders/shipments | Update a purchase order shipment.



## Acknowledge

> ApiResponse Acknowledge(ctx).SingleMerchantAcknowledgePurchaseOrderLinesRequest(singleMerchantAcknowledgePurchaseOrderLinesRequest).Execute()

Acknowledges lines of a purchase order



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
	singleMerchantAcknowledgePurchaseOrderLinesRequest := *openapiclient.NewSingleMerchantAcknowledgePurchaseOrderLinesRequest() // SingleMerchantAcknowledgePurchaseOrderLinesRequest | Model for purchase order and lines data to be acknowledged with the channel. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersAPI.Acknowledge(context.Background()).SingleMerchantAcknowledgePurchaseOrderLinesRequest(singleMerchantAcknowledgePurchaseOrderLinesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersAPI.Acknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Acknowledge`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersAPI.Acknowledge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantAcknowledgePurchaseOrderLinesRequest** | [**SingleMerchantAcknowledgePurchaseOrderLinesRequest**](SingleMerchantAcknowledgePurchaseOrderLinesRequest.md) | Model for purchase order and lines data to be acknowledged with the channel. | 

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


## Create

> ApiResponse Create(ctx).SingleMerchantCreatePurchaseOrderShipmentRequest(singleMerchantCreatePurchaseOrderShipmentRequest).Execute()

Create a purchase order shipment.



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
	singleMerchantCreatePurchaseOrderShipmentRequest := *openapiclient.NewSingleMerchantCreatePurchaseOrderShipmentRequest() // SingleMerchantCreatePurchaseOrderShipmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersAPI.Create(context.Background()).SingleMerchantCreatePurchaseOrderShipmentRequest(singleMerchantCreatePurchaseOrderShipmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantCreatePurchaseOrderShipmentRequest** | [**SingleMerchantCreatePurchaseOrderShipmentRequest**](SingleMerchantCreatePurchaseOrderShipmentRequest.md) |  | 

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


## GetByFilter

> CollectionOfIPurchaseOrderShipmentByFilter GetByFilter(ctx).ChannelId(channelId).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ShippedDateRangeFromDate(shippedDateRangeFromDate).ShippedDateRangeToDate(shippedDateRangeToDate).CreateDateRangeFromDate(createDateRangeFromDate).CreateDateRangeToDate(createDateRangeToDate).UpdateDateRangeFromDate(updateDateRangeFromDate).UpdateDateRangeToDate(updateDateRangeToDate).BillOfLadingNumber(billOfLadingNumber).CarrierName(carrierName).PageIndex(pageIndex).PageSize(pageSize).Execute()

Gets purchase order shipments by filter



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
	channelId := int32(56) // int32 | The identifier of the channel (optional)
	identifiersIdentifierType := openapiclient.PurchaseOrderShipmentIdentifierTypeValue("MERCHANT_SHIPMENT_NO") // PurchaseOrderShipmentIdentifierTypeValue | The type of identifier: which identifier to filter on (optional)
	identifiersModels := []string{"Inner_example"} // []string | The value (of the selected type) to filter on (optional)
	shippedDateRangeFromDate := time.Now() // time.Time |  (optional)
	shippedDateRangeToDate := time.Now() // time.Time |  (optional)
	createDateRangeFromDate := time.Now() // time.Time |  (optional)
	createDateRangeToDate := time.Now() // time.Time |  (optional)
	updateDateRangeFromDate := time.Now() // time.Time |  (optional)
	updateDateRangeToDate := time.Now() // time.Time |  (optional)
	billOfLadingNumber := "billOfLadingNumber_example" // string | The Bill of Lading number. Multiple shipments can have the same Bill of Lading number (optional)
	carrierName := "carrierName_example" // string | The name of the carrier (optional)
	pageIndex := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersAPI.GetByFilter(context.Background()).ChannelId(channelId).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ShippedDateRangeFromDate(shippedDateRangeFromDate).ShippedDateRangeToDate(shippedDateRangeToDate).CreateDateRangeFromDate(createDateRangeFromDate).CreateDateRangeToDate(createDateRangeToDate).UpdateDateRangeFromDate(updateDateRangeFromDate).UpdateDateRangeToDate(updateDateRangeToDate).BillOfLadingNumber(billOfLadingNumber).CarrierName(carrierName).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersAPI.GetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByFilter`: CollectionOfIPurchaseOrderShipmentByFilter
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersAPI.GetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **int32** | The identifier of the channel | 
 **identifiersIdentifierType** | [**PurchaseOrderShipmentIdentifierTypeValue**](PurchaseOrderShipmentIdentifierTypeValue.md) | The type of identifier: which identifier to filter on | 
 **identifiersModels** | **[]string** | The value (of the selected type) to filter on | 
 **shippedDateRangeFromDate** | **time.Time** |  | 
 **shippedDateRangeToDate** | **time.Time** |  | 
 **createDateRangeFromDate** | **time.Time** |  | 
 **createDateRangeToDate** | **time.Time** |  | 
 **updateDateRangeFromDate** | **time.Time** |  | 
 **updateDateRangeToDate** | **time.Time** |  | 
 **billOfLadingNumber** | **string** | The Bill of Lading number. Multiple shipments can have the same Bill of Lading number | 
 **carrierName** | **string** | The name of the carrier | 
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**CollectionOfIPurchaseOrderShipmentByFilter**](CollectionOfIPurchaseOrderShipmentByFilter.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByFilter_0

> CollectionOfIPurchaseOrderByFilter GetByFilter_0(ctx).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).Statuses(statuses).OrderDateRangeFromDate(orderDateRangeFromDate).OrderDateRangeToDate(orderDateRangeToDate).CreateDateRangeFromDate(createDateRangeFromDate).CreateDateRangeToDate(createDateRangeToDate).UpdateDateRangeFromDate(updateDateRangeFromDate).UpdateDateRangeToDate(updateDateRangeToDate).ChannelIds(channelIds).Type_(type_).PageIndex(pageIndex).PageSize(pageSize).Execute()

Gets purchase orders by filter



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
	identifiersIdentifierType := openapiclient.PurchaseOrderIdentifierType("PURCHASE_ORDER_ID") // PurchaseOrderIdentifierType | The type of identifier: which identifier to filter on (optional)
	identifiersModels := []string{"Inner_example"} // []string | The value (of the selected type) to filter on (optional)
	statuses := []openapiclient.ModulesPurchaseOrderStatus{openapiclient.ModulesPurchaseOrderStatus("NEW")} // []ModulesPurchaseOrderStatus |  (optional)
	orderDateRangeFromDate := time.Now() // time.Time |  (optional)
	orderDateRangeToDate := time.Now() // time.Time |  (optional)
	createDateRangeFromDate := time.Now() // time.Time |  (optional)
	createDateRangeToDate := time.Now() // time.Time |  (optional)
	updateDateRangeFromDate := time.Now() // time.Time |  (optional)
	updateDateRangeToDate := time.Now() // time.Time |  (optional)
	channelIds := []int32{int32(123)} // []int32 |  (optional)
	type_ := openapiclient.ModulesPurchaseOrderType("REGULAR_ORDER") // ModulesPurchaseOrderType |  (optional)
	pageIndex := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersAPI.GetByFilter_0(context.Background()).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).Statuses(statuses).OrderDateRangeFromDate(orderDateRangeFromDate).OrderDateRangeToDate(orderDateRangeToDate).CreateDateRangeFromDate(createDateRangeFromDate).CreateDateRangeToDate(createDateRangeToDate).UpdateDateRangeFromDate(updateDateRangeFromDate).UpdateDateRangeToDate(updateDateRangeToDate).ChannelIds(channelIds).Type_(type_).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersAPI.GetByFilter_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByFilter_0`: CollectionOfIPurchaseOrderByFilter
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersAPI.GetByFilter_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByFilter_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifiersIdentifierType** | [**PurchaseOrderIdentifierType**](PurchaseOrderIdentifierType.md) | The type of identifier: which identifier to filter on | 
 **identifiersModels** | **[]string** | The value (of the selected type) to filter on | 
 **statuses** | [**[]ModulesPurchaseOrderStatus**](ModulesPurchaseOrderStatus.md) |  | 
 **orderDateRangeFromDate** | **time.Time** |  | 
 **orderDateRangeToDate** | **time.Time** |  | 
 **createDateRangeFromDate** | **time.Time** |  | 
 **createDateRangeToDate** | **time.Time** |  | 
 **updateDateRangeFromDate** | **time.Time** |  | 
 **updateDateRangeToDate** | **time.Time** |  | 
 **channelIds** | **[]int32** |  | 
 **type_** | [**ModulesPurchaseOrderType**](ModulesPurchaseOrderType.md) |  | 
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**CollectionOfIPurchaseOrderByFilter**](CollectionOfIPurchaseOrderByFilter.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ApiResponse Update(ctx).SingleMerchantUpdatePurchaseOrderShipmentRequest(singleMerchantUpdatePurchaseOrderShipmentRequest).Execute()

Update a purchase order shipment.

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
	singleMerchantUpdatePurchaseOrderShipmentRequest := *openapiclient.NewSingleMerchantUpdatePurchaseOrderShipmentRequest() // SingleMerchantUpdatePurchaseOrderShipmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersAPI.Update(context.Background()).SingleMerchantUpdatePurchaseOrderShipmentRequest(singleMerchantUpdatePurchaseOrderShipmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersAPI.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantUpdatePurchaseOrderShipmentRequest** | [**SingleMerchantUpdatePurchaseOrderShipmentRequest**](SingleMerchantUpdatePurchaseOrderShipmentRequest.md) |  | 

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

