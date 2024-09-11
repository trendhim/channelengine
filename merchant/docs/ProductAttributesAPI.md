# \ProductAttributesAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAttributeGroupAddProductExtraData**](ProductAttributesAPI.md#ProductAttributeGroupAddProductExtraData) | **Put** /v2/product-attribute-group/{groupName}/add | Adds custom attributes to a group
[**ProductAttributeGroupCreate**](ProductAttributesAPI.md#ProductAttributeGroupCreate) | **Post** /v2/product-attribute-group | Creates a custom attribute group
[**ProductAttributeGroupDelete**](ProductAttributesAPI.md#ProductAttributeGroupDelete) | **Delete** /v2/product-attribute-group/{groupName} | Deletes a custom attribute group
[**ProductAttributeGroupGetByFilter**](ProductAttributesAPI.md#ProductAttributeGroupGetByFilter) | **Get** /v2/product-attribute-group | Gets custom attribute groups
[**ProductAttributeGroupGetWithChannelsByFilter**](ProductAttributesAPI.md#ProductAttributeGroupGetWithChannelsByFilter) | **Get** /v2/product-attribute-group/linked-channels | Gets custom attribute groups and linked marketplaces
[**ProductAttributeGroupRemoveProductExtraData**](ProductAttributesAPI.md#ProductAttributeGroupRemoveProductExtraData) | **Put** /v2/product-attribute-group/{groupName}/remove | Deletes custom attributes from a group
[**ProductAttributeGroupRenameProductAttributeGroup**](ProductAttributesAPI.md#ProductAttributeGroupRenameProductAttributeGroup) | **Post** /v2/product-attribute-group/rename | Renames custom attribute groups



## ProductAttributeGroupAddProductExtraData

> ApiResponse ProductAttributeGroupAddProductExtraData(ctx, groupName).AddProductExtraDataRequests(addProductExtraDataRequests).Execute()

Adds custom attributes to a group



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
	groupName := "groupName_example" // string | The group name of the product attribute group you wish to add product extra data.
	addProductExtraDataRequests := *openapiclient.NewAddProductExtraDataRequests() // AddProductExtraDataRequests | Product extra data keys to be added. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupAddProductExtraData(context.Background(), groupName).AddProductExtraDataRequests(addProductExtraDataRequests).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupAddProductExtraData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupAddProductExtraData`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupAddProductExtraData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The group name of the product attribute group you wish to add product extra data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupAddProductExtraDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addProductExtraDataRequests** | [**AddProductExtraDataRequests**](AddProductExtraDataRequests.md) | Product extra data keys to be added. | 

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


## ProductAttributeGroupCreate

> ApiResponse ProductAttributeGroupCreate(ctx).ProductAttributeGroupRequest(productAttributeGroupRequest).Execute()

Creates a custom attribute group



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
	productAttributeGroupRequest := []openapiclient.ProductAttributeGroupRequest{*openapiclient.NewProductAttributeGroupRequest("GroupName_example", []string{"ProductExtraDataKeys_example"})} // []ProductAttributeGroupRequest | Collection of product attribute groups with linked product extra data keys.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupCreate(context.Background()).ProductAttributeGroupRequest(productAttributeGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAttributeGroupRequest** | [**[]ProductAttributeGroupRequest**](ProductAttributeGroupRequest.md) | Collection of product attribute groups with linked product extra data keys. | 

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


## ProductAttributeGroupDelete

> ApiResponse ProductAttributeGroupDelete(ctx, groupName).Execute()

Deletes a custom attribute group



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
	groupName := "groupName_example" // string | The group name of the product attribute group you wish to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupDelete(context.Background(), groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupDelete`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The group name of the product attribute group you wish to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAttributeGroupGetByFilter

> CollectionOfMerchantProductAttributeGroupWithProductExtraDataResponse ProductAttributeGroupGetByFilter(ctx).GroupNames(groupNames).Page(page).Execute()

Gets custom attribute groups



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
	groupNames := []string{"Inner_example"} // []string |  (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupGetByFilter(context.Background()).GroupNames(groupNames).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupGetByFilter`: CollectionOfMerchantProductAttributeGroupWithProductExtraDataResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupNames** | **[]string** |  | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantProductAttributeGroupWithProductExtraDataResponse**](CollectionOfMerchantProductAttributeGroupWithProductExtraDataResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAttributeGroupGetWithChannelsByFilter

> CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse ProductAttributeGroupGetWithChannelsByFilter(ctx).GroupNames(groupNames).Page(page).Execute()

Gets custom attribute groups and linked marketplaces



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
	groupNames := []string{"Inner_example"} // []string |  (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupGetWithChannelsByFilter(context.Background()).GroupNames(groupNames).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupGetWithChannelsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupGetWithChannelsByFilter`: CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupGetWithChannelsByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupGetWithChannelsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupNames** | **[]string** |  | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse**](CollectionOfMerchantProductAttributeGroupWithLinkedChannelsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAttributeGroupRemoveProductExtraData

> ApiResponse ProductAttributeGroupRemoveProductExtraData(ctx, groupName).RemoveProductExtraDataRequests(removeProductExtraDataRequests).Execute()

Deletes custom attributes from a group



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
	groupName := "groupName_example" // string | The group name of the product attribute group you wish to remove product extra data.
	removeProductExtraDataRequests := *openapiclient.NewRemoveProductExtraDataRequests() // RemoveProductExtraDataRequests | Product extra data keys to be removed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupRemoveProductExtraData(context.Background(), groupName).RemoveProductExtraDataRequests(removeProductExtraDataRequests).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupRemoveProductExtraData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupRemoveProductExtraData`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupRemoveProductExtraData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The group name of the product attribute group you wish to remove product extra data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupRemoveProductExtraDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeProductExtraDataRequests** | [**RemoveProductExtraDataRequests**](RemoveProductExtraDataRequests.md) | Product extra data keys to be removed. | 

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


## ProductAttributeGroupRenameProductAttributeGroup

> ApiResponse ProductAttributeGroupRenameProductAttributeGroup(ctx).RenameProductAttributeGroupRequests(renameProductAttributeGroupRequests).Execute()

Renames custom attribute groups



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
	renameProductAttributeGroupRequests := []openapiclient.RenameProductAttributeGroupRequests{*openapiclient.NewRenameProductAttributeGroupRequests()} // []RenameProductAttributeGroupRequests | List of old and new product attribute group names. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.ProductAttributeGroupRenameProductAttributeGroup(context.Background()).RenameProductAttributeGroupRequests(renameProductAttributeGroupRequests).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.ProductAttributeGroupRenameProductAttributeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeGroupRenameProductAttributeGroup`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.ProductAttributeGroupRenameProductAttributeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeGroupRenameProductAttributeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameProductAttributeGroupRequests** | [**[]RenameProductAttributeGroupRequests**](RenameProductAttributeGroupRequests.md) | List of old and new product attribute group names. | 

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

