# \ProductAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAcknowledgeDataChanges**](ProductAPI.md#ProductAcknowledgeDataChanges) | **Post** /v2/products/data | Acknowledge product data changes
[**ProductAcknowledgeOfferChanges**](ProductAPI.md#ProductAcknowledgeOfferChanges) | **Post** /v2/products/offers | Acknowledge product offer changes
[**ProductGetDataChanges**](ProductAPI.md#ProductGetDataChanges) | **Get** /v2/products/data | Gets product data changes
[**ProductGetDataChangesFull**](ProductAPI.md#ProductGetDataChangesFull) | **Get** /v2/products/data/full | Gets product data changes with an optional product type filter. If you select product type products will be filtered by it.  If you won&#39;t pass product type you will get products with types: CHILD, PARENT, GRANDPARENT, SINGLE
[**ProductGetOfferChanges**](ProductAPI.md#ProductGetOfferChanges) | **Get** /v2/products/offers | Gets product offer changes



## ProductAcknowledgeDataChanges

> ApiResponse ProductAcknowledgeDataChanges(ctx).ChannelProcessedChangesRequest(channelProcessedChangesRequest).Execute()

Acknowledge product data changes



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
    channelProcessedChangesRequest := *openapiclient.NewChannelProcessedChangesRequest() // ChannelProcessedChangesRequest | The merchant references of the products which have been<br /> successfully created, updated or deleted (after a call to 'GetDataChanges'). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ProductAcknowledgeDataChanges(context.Background()).ChannelProcessedChangesRequest(channelProcessedChangesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAcknowledgeDataChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductAcknowledgeDataChanges`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAcknowledgeDataChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAcknowledgeDataChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelProcessedChangesRequest** | [**ChannelProcessedChangesRequest**](ChannelProcessedChangesRequest.md) | The merchant references of the products which have been&lt;br /&gt; successfully created, updated or deleted (after a call to &#39;GetDataChanges&#39;). | 

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


## ProductAcknowledgeOfferChanges

> ApiResponse ProductAcknowledgeOfferChanges(ctx).KeyIsMpn(keyIsMpn).RequestBody(requestBody).Execute()

Acknowledge product offer changes



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
    keyIsMpn := true // bool | If set to true, changes has to be a list of merchant references instead of channel references. (optional) (default to false)
    requestBody := []string{"Property_example"} // []string | The channel references of the updated products. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ProductAcknowledgeOfferChanges(context.Background()).KeyIsMpn(keyIsMpn).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAcknowledgeOfferChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductAcknowledgeOfferChanges`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAcknowledgeOfferChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAcknowledgeOfferChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyIsMpn** | **bool** | If set to true, changes has to be a list of merchant references instead of channel references. | [default to false]
 **requestBody** | **[]string** | The channel references of the updated products. | 

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


## ProductGetDataChanges

> CollectionChangesOfChannelProductChangesResponse ProductGetDataChanges(ctx).MaxCount(maxCount).StripHtml(stripHtml).Page(page).PageSize(pageSize).Execute()

Gets product data changes



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
    maxCount := int32(56) // int32 | Optional - limit the amount of products returned for each field<br /> (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number. (optional)
    stripHtml := true // bool | Optional - strips html by default on all fields (optional) (default to true)
    page := int32(56) // int32 | Optional - for default is first page returned (optional) (default to 1)
    pageSize := int32(56) // int32 | Optional - amount of products returned, if not provided return all products (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ProductGetDataChanges(context.Background()).MaxCount(maxCount).StripHtml(stripHtml).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductGetDataChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetDataChanges`: CollectionChangesOfChannelProductChangesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductGetDataChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetDataChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxCount** | **int32** | Optional - limit the amount of products returned for each field&lt;br /&gt; (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number. | 
 **stripHtml** | **bool** | Optional - strips html by default on all fields | [default to true]
 **page** | **int32** | Optional - for default is first page returned | [default to 1]
 **pageSize** | **int32** | Optional - amount of products returned, if not provided return all products | 

### Return type

[**CollectionChangesOfChannelProductChangesResponse**](CollectionChangesOfChannelProductChangesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetDataChangesFull

> CollectionChangesOfChannelProductChangesResponse ProductGetDataChangesFull(ctx).ProductType(productType).MaxCount(maxCount).StripHtml(stripHtml).Page(page).PageSize(pageSize).Execute()

Gets product data changes with an optional product type filter. If you select product type products will be filtered by it.  If you won't pass product type you will get products with types: CHILD, PARENT, GRANDPARENT, SINGLE

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
    productType := openapiclient.DataChangesProductType("SINGLE") // DataChangesProductType | Optional - Type of products (optional)
    maxCount := int32(56) // int32 | Optional - limit the amount of products returned for each field<br /> (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number. (optional)
    stripHtml := true // bool | Optional - strips html by default on all fields (optional) (default to true)
    page := int32(56) // int32 | Optional - for default is first page returned (optional) (default to 1)
    pageSize := int32(56) // int32 | Optional - amount of products returned, if not provided return all products (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ProductGetDataChangesFull(context.Background()).ProductType(productType).MaxCount(maxCount).StripHtml(stripHtml).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductGetDataChangesFull``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetDataChangesFull`: CollectionChangesOfChannelProductChangesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductGetDataChangesFull`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetDataChangesFullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productType** | [**DataChangesProductType**](DataChangesProductType.md) | Optional - Type of products | 
 **maxCount** | **int32** | Optional - limit the amount of products returned for each field&lt;br /&gt; (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number. | 
 **stripHtml** | **bool** | Optional - strips html by default on all fields | [default to true]
 **page** | **int32** | Optional - for default is first page returned | [default to 1]
 **pageSize** | **int32** | Optional - amount of products returned, if not provided return all products | 

### Return type

[**CollectionChangesOfChannelProductChangesResponse**](CollectionChangesOfChannelProductChangesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetOfferChanges

> CollectionOfChannelOfferResponse ProductGetOfferChanges(ctx).Page(page).PageSize(pageSize).Execute()

Gets product offer changes



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
    page := int32(56) // int32 | Optional - for default is first page returned (optional) (default to 1)
    pageSize := int32(56) // int32 | Optional - amount of products returned, if not provided return all products (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ProductGetOfferChanges(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductGetOfferChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetOfferChanges`: CollectionOfChannelOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductGetOfferChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetOfferChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Optional - for default is first page returned | [default to 1]
 **pageSize** | **int32** | Optional - amount of products returned, if not provided return all products | 

### Return type

[**CollectionOfChannelOfferResponse**](CollectionOfChannelOfferResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

