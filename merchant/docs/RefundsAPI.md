# \RefundsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundAcknowledge**](RefundsAPI.md#RefundAcknowledge) | **Post** /v2.1/refunds/merchant/acknowledge | [CLOSED BETA] Acknowledge a refund
[**RefundCreate**](RefundsAPI.md#RefundCreate) | **Post** /v2.1/refunds/merchant | [CLOSED BETA] Create a refund
[**RefundGet**](RefundsAPI.md#RefundGet) | **Get** /v2.1/refunds/merchant/{identifier} | [CLOSED BETA] Get refund by identifier
[**RefundGetByFilter**](RefundsAPI.md#RefundGetByFilter) | **Get** /v2.1/refunds/merchant | [CLOSED BETA] Get refunds by filter



## RefundAcknowledge

> ApiResponse RefundAcknowledge(ctx).SingleMerchantAcknowledgeRefundRequest(singleMerchantAcknowledgeRefundRequest).Execute()

[CLOSED BETA] Acknowledge a refund



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
	singleMerchantAcknowledgeRefundRequest := *openapiclient.NewSingleMerchantAcknowledgeRefundRequest() // SingleMerchantAcknowledgeRefundRequest | The refund to acknowledge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundAcknowledge(context.Background()).SingleMerchantAcknowledgeRefundRequest(singleMerchantAcknowledgeRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundAcknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundAcknowledge`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundAcknowledge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantAcknowledgeRefundRequest** | [**SingleMerchantAcknowledgeRefundRequest**](SingleMerchantAcknowledgeRefundRequest.md) | The refund to acknowledge | 

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


## RefundCreate

> ApiResponse RefundCreate(ctx).SingleMerchantCreateRefundRequest(singleMerchantCreateRefundRequest).Execute()

[CLOSED BETA] Create a refund



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
	singleMerchantCreateRefundRequest := *openapiclient.NewSingleMerchantCreateRefundRequest() // SingleMerchantCreateRefundRequest | The refund (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundCreate(context.Background()).SingleMerchantCreateRefundRequest(singleMerchantCreateRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleMerchantCreateRefundRequest** | [**SingleMerchantCreateRefundRequest**](SingleMerchantCreateRefundRequest.md) | The refund | 

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


## RefundGet

> SingleOfIRefund RefundGet(ctx, identifier).Type_(type_).Execute()

[CLOSED BETA] Get refund by identifier



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
	identifier := "identifier_example" // string | The identifier to search for
	type_ := openapiclient.RefundIdentifier("REFUND_ID") // RefundIdentifier | Specify whether to search by ID, Merchant Refund No or Channel Refund No (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundGet(context.Background(), identifier).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundGet`: SingleOfIRefund
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The identifier to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**RefundIdentifier**](RefundIdentifier.md) | Specify whether to search by ID, Merchant Refund No or Channel Refund No | 

### Return type

[**SingleOfIRefund**](SingleOfIRefund.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundGetByFilter

> SingleOfIRefund RefundGetByFilter(ctx).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ChannelExportStatusStatuses(channelExportStatusStatuses).ChannelExportStatusMaxNumberOfExportAttempts(channelExportStatusMaxNumberOfExportAttempts).Reasons(reasons).CreatedDateRangeFromDate(createdDateRangeFromDate).CreatedDateRangeToDate(createdDateRangeToDate).ChannelIds(channelIds).Search(search).IsAcknowledgedByMerchant(isAcknowledgedByMerchant).IsAcknowledgedByChannel(isAcknowledgedByChannel).FulfillmentType(fulfillmentType).CreatorType(creatorType).ExternalBatchNos(externalBatchNos).PageIndex(pageIndex).PageSize(pageSize).Execute()

[CLOSED BETA] Get refunds by filter



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
	identifiersIdentifierType := openapiclient.RefundByFilterIdentifier("REFUND_ID") // RefundByFilterIdentifier | The type of identifier: which identifier to filter on (optional)
	identifiersModels := []string{"Inner_example"} // []string | The value (of the selected type) to filter on (optional)
	channelExportStatusStatuses := []openapiclient.ChannelExportStatus{openapiclient.ChannelExportStatus("AWAITING_EXPORT")} // []ChannelExportStatus |  (optional)
	channelExportStatusMaxNumberOfExportAttempts := int32(56) // int32 |  (optional)
	reasons := []openapiclient.RefundReason{openapiclient.RefundReason("PRODUCT_DEFECT")} // []RefundReason |  (optional)
	createdDateRangeFromDate := time.Now() // time.Time |  (optional)
	createdDateRangeToDate := time.Now() // time.Time |  (optional)
	channelIds := []int32{int32(123)} // []int32 |  (optional)
	search := "search_example" // string |  (optional)
	isAcknowledgedByMerchant := true // bool |  (optional)
	isAcknowledgedByChannel := true // bool |  (optional)
	fulfillmentType := openapiclient.ModuleFulfillmentType("ALL") // ModuleFulfillmentType |  (optional)
	creatorType := openapiclient.CreatorType("MERCHANT") // CreatorType |  (optional)
	externalBatchNos := []string{"Inner_example"} // []string |  (optional)
	pageIndex := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundGetByFilter(context.Background()).IdentifiersIdentifierType(identifiersIdentifierType).IdentifiersModels(identifiersModels).ChannelExportStatusStatuses(channelExportStatusStatuses).ChannelExportStatusMaxNumberOfExportAttempts(channelExportStatusMaxNumberOfExportAttempts).Reasons(reasons).CreatedDateRangeFromDate(createdDateRangeFromDate).CreatedDateRangeToDate(createdDateRangeToDate).ChannelIds(channelIds).Search(search).IsAcknowledgedByMerchant(isAcknowledgedByMerchant).IsAcknowledgedByChannel(isAcknowledgedByChannel).FulfillmentType(fulfillmentType).CreatorType(creatorType).ExternalBatchNos(externalBatchNos).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundGetByFilter`: SingleOfIRefund
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifiersIdentifierType** | [**RefundByFilterIdentifier**](RefundByFilterIdentifier.md) | The type of identifier: which identifier to filter on | 
 **identifiersModels** | **[]string** | The value (of the selected type) to filter on | 
 **channelExportStatusStatuses** | [**[]ChannelExportStatus**](ChannelExportStatus.md) |  | 
 **channelExportStatusMaxNumberOfExportAttempts** | **int32** |  | 
 **reasons** | [**[]RefundReason**](RefundReason.md) |  | 
 **createdDateRangeFromDate** | **time.Time** |  | 
 **createdDateRangeToDate** | **time.Time** |  | 
 **channelIds** | **[]int32** |  | 
 **search** | **string** |  | 
 **isAcknowledgedByMerchant** | **bool** |  | 
 **isAcknowledgedByChannel** | **bool** |  | 
 **fulfillmentType** | [**ModuleFulfillmentType**](ModuleFulfillmentType.md) |  | 
 **creatorType** | [**CreatorType**](CreatorType.md) |  | 
 **externalBatchNos** | **[]string** |  | 
 **pageIndex** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**SingleOfIRefund**](SingleOfIRefund.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

