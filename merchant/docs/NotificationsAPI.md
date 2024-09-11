# \NotificationsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationIndex**](NotificationsAPI.md#NotificationIndex) | **Get** /v2/notifications | Gets notifications



## NotificationIndex

> CollectionOfMerchantNotificationResponse NotificationIndex(ctx).FromDate(fromDate).ToDate(toDate).Types(types).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).MerchantReturnNos(merchantReturnNos).ChannelReturnNos(channelReturnNos).MerchantShipmentNos(merchantShipmentNos).Page(page).Execute()

Gets notifications



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
	fromDate := time.Now() // time.Time | Filter on the notification date, starting from this date. This date is inclusive. (optional)
	toDate := time.Now() // time.Time | Filter on the notification date, until this date. This date is exclusive. (optional)
	types := []openapiclient.NotificationType{openapiclient.NotificationType("CHANNEL_ORDER_ANONYMIZED_BY_REQUEST")} // []NotificationType | Notification type(s) to filter on. (optional)
	merchantOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the merchant. (optional)
	channelOrderNos := []string{"Inner_example"} // []string | Filter on unique order reference used by the channel. (optional)
	merchantReturnNos := []string{"Inner_example"} // []string | Filter on unique return reference used by the merchant. (optional)
	channelReturnNos := []string{"Inner_example"} // []string | Filter on unique return reference used by the channel. (optional)
	merchantShipmentNos := []string{"Inner_example"} // []string | Filter on unique shipment reference used by the merchant. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.NotificationIndex(context.Background()).FromDate(fromDate).ToDate(toDate).Types(types).MerchantOrderNos(merchantOrderNos).ChannelOrderNos(channelOrderNos).MerchantReturnNos(merchantReturnNos).ChannelReturnNos(channelReturnNos).MerchantShipmentNos(merchantShipmentNos).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.NotificationIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationIndex`: CollectionOfMerchantNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.NotificationIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **time.Time** | Filter on the notification date, starting from this date. This date is inclusive. | 
 **toDate** | **time.Time** | Filter on the notification date, until this date. This date is exclusive. | 
 **types** | [**[]NotificationType**](NotificationType.md) | Notification type(s) to filter on. | 
 **merchantOrderNos** | **[]string** | Filter on unique order reference used by the merchant. | 
 **channelOrderNos** | **[]string** | Filter on unique order reference used by the channel. | 
 **merchantReturnNos** | **[]string** | Filter on unique return reference used by the merchant. | 
 **channelReturnNos** | **[]string** | Filter on unique return reference used by the channel. | 
 **merchantShipmentNos** | **[]string** | Filter on unique shipment reference used by the merchant. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantNotificationResponse**](CollectionOfMerchantNotificationResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

