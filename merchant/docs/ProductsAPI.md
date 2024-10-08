# \ProductsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductBulkDelete**](ProductsAPI.md#ProductBulkDelete) | **Post** /v2/products/bulkdelete | Deletes products
[**ProductBulkPatch**](ProductsAPI.md#ProductBulkPatch) | **Patch** /v2/products | Updates products attributes
[**ProductBulkPatchExtraDataItems**](ProductsAPI.md#ProductBulkPatchExtraDataItems) | **Patch** /v2/products/extra-data/bulk | Adds, updates, or deletes custom attributes
[**ProductCreate**](ProductsAPI.md#ProductCreate) | **Post** /v2/products | Updates or creates products
[**ProductDelete**](ProductsAPI.md#ProductDelete) | **Delete** /v2/products/{merchantProductNo} | Deletes a product
[**ProductFreeze**](ProductsAPI.md#ProductFreeze) | **Post** /v2/products/freeze | Updates selected products and sets them either to frozen or not-frozen status.
[**ProductGetByFilter**](ProductsAPI.md#ProductGetByFilter) | **Get** /v2/products | Gets products
[**ProductGetByMerchantProductNo**](ProductsAPI.md#ProductGetByMerchantProductNo) | **Get** /v2/products/{merchantProductNo} | Gets a product
[**ProductPatch**](ProductsAPI.md#ProductPatch) | **Patch** /v2/products/{merchantProductNo} | Updates product attributes
[**ProductPatchExtraDataItems**](ProductsAPI.md#ProductPatchExtraDataItems) | **Patch** /v2/products/extra-data | Adds, updates, or deletes a custom attribute



## ProductBulkDelete

> ApiResponse ProductBulkDelete(ctx).RequestBody(requestBody).Execute()

Deletes products



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
	requestBody := []string{"Property_example"} // []string | The list of MerchantProductNo of the products you wish to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductBulkDelete(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductBulkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBulkDelete`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductBulkDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | The list of MerchantProductNo of the products you wish to delete. | 

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


## ProductBulkPatch

> SingleOfProductCreationResult ProductBulkPatch(ctx).PatchMerchantProductDto(patchMerchantProductDto).Execute()

Updates products attributes



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
	patchMerchantProductDto := *openapiclient.NewPatchMerchantProductDto() // PatchMerchantProductDto | 1) PropertiesToUpdate: Fields to update<br />2) MerchantProductRequestModels: Products to be updated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductBulkPatch(context.Background()).PatchMerchantProductDto(patchMerchantProductDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductBulkPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBulkPatch`: SingleOfProductCreationResult
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductBulkPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchMerchantProductDto** | [**PatchMerchantProductDto**](PatchMerchantProductDto.md) | 1) PropertiesToUpdate: Fields to update&lt;br /&gt;2) MerchantProductRequestModels: Products to be updated | 

### Return type

[**SingleOfProductCreationResult**](SingleOfProductCreationResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBulkPatchExtraDataItems

> SingleOfProductCreationResult ProductBulkPatchExtraDataItems(ctx).MerchantProductExtraDataRequest(merchantProductExtraDataRequest).Execute()

Adds, updates, or deletes custom attributes



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
	merchantProductExtraDataRequest := []openapiclient.MerchantProductExtraDataRequest{*openapiclient.NewMerchantProductExtraDataRequest("MerchantProductNo_example", []openapiclient.ProductExtraDataRequest{*openapiclient.NewProductExtraDataRequest("Op_example", "Key_example")})} // []MerchantProductExtraDataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductBulkPatchExtraDataItems(context.Background()).MerchantProductExtraDataRequest(merchantProductExtraDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductBulkPatchExtraDataItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBulkPatchExtraDataItems`: SingleOfProductCreationResult
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductBulkPatchExtraDataItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkPatchExtraDataItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantProductExtraDataRequest** | [**[]MerchantProductExtraDataRequest**](MerchantProductExtraDataRequest.md) |  | 

### Return type

[**SingleOfProductCreationResult**](SingleOfProductCreationResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreate

> SingleOfProductCreationResult ProductCreate(ctx).MerchantProductRequest(merchantProductRequest).IgnoreStock(ignoreStock).IgnorePrice(ignorePrice).Execute()

Updates or creates products



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
	merchantProductRequest := []openapiclient.MerchantProductRequest{*openapiclient.NewMerchantProductRequest("MerchantProductNo_example")} // []MerchantProductRequest | 
	ignoreStock := true // bool |  (optional) (default to false)
	ignorePrice := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCreate(context.Background()).MerchantProductRequest(merchantProductRequest).IgnoreStock(ignoreStock).IgnorePrice(ignorePrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCreate`: SingleOfProductCreationResult
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantProductRequest** | [**[]MerchantProductRequest**](MerchantProductRequest.md) |  | 
 **ignoreStock** | **bool** |  | [default to false]
 **ignorePrice** | **bool** |  | [default to false]

### Return type

[**SingleOfProductCreationResult**](SingleOfProductCreationResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDelete

> ApiResponse ProductDelete(ctx, merchantProductNo).Execute()

Deletes a product



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
	merchantProductNo := "merchantProductNo_example" // string | The MerchantProductNo of the product you wish to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductDelete(context.Background(), merchantProductNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDelete`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantProductNo** | **string** | The MerchantProductNo of the product you wish to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteRequest struct via the builder pattern


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


## ProductFreeze

> SingleOfApiResponse ProductFreeze(ctx).FreezeProductRequest(freezeProductRequest).Execute()

Updates selected products and sets them either to frozen or not-frozen status.



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
	freezeProductRequest := []openapiclient.FreezeProductRequest{*openapiclient.NewFreezeProductRequest("MerchantProductNo_example", "Reason_example", openapiclient.FreezingActionRequest("FREEZE"))} // []FreezeProductRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductFreeze(context.Background()).FreezeProductRequest(freezeProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductFreeze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductFreeze`: SingleOfApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductFreeze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductFreezeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freezeProductRequest** | [**[]FreezeProductRequest**](FreezeProductRequest.md) |  | 

### Return type

[**SingleOfApiResponse**](SingleOfApiResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetByFilter

> CollectionOfMerchantProductResponse ProductGetByFilter(ctx).Search(search).EanList(eanList).MerchantProductNoList(merchantProductNoList).Page(page).Execute()

Gets products



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
	search := "search_example" // string | Search product(s) by Name, MerchantProductNo, Ean or Brand<br />This search is applied to the result after applying the other filters. (optional)
	eanList := []string{"Inner_example"} // []string | Search products by submitting a list of EAN's. (optional)
	merchantProductNoList := []string{"Inner_example"} // []string | Search products by submitting a list of MerchantProductNo's. (optional)
	page := int32(56) // int32 | The page to filter on. Starts at 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductGetByFilter(context.Background()).Search(search).EanList(eanList).MerchantProductNoList(merchantProductNoList).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductGetByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductGetByFilter`: CollectionOfMerchantProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductGetByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search product(s) by Name, MerchantProductNo, Ean or Brand&lt;br /&gt;This search is applied to the result after applying the other filters. | 
 **eanList** | **[]string** | Search products by submitting a list of EAN&#39;s. | 
 **merchantProductNoList** | **[]string** | Search products by submitting a list of MerchantProductNo&#39;s. | 
 **page** | **int32** | The page to filter on. Starts at 1. | 

### Return type

[**CollectionOfMerchantProductResponse**](CollectionOfMerchantProductResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetByMerchantProductNo

> SingleOfMerchantProductResponse ProductGetByMerchantProductNo(ctx, merchantProductNo).Execute()

Gets a product



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
	merchantProductNo := "merchantProductNo_example" // string | The unique product reference used by the Merchant (sku).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductGetByMerchantProductNo(context.Background(), merchantProductNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductGetByMerchantProductNo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductGetByMerchantProductNo`: SingleOfMerchantProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductGetByMerchantProductNo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantProductNo** | **string** | The unique product reference used by the Merchant (sku). | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductGetByMerchantProductNoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleOfMerchantProductResponse**](SingleOfMerchantProductResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductPatch

> SingleOfProductCreationResult ProductPatch(ctx, merchantProductNo).Operation(operation).Execute()

Updates product attributes



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
	merchantProductNo := "merchantProductNo_example" // string | The MerchantProductNo of the product you wish to patch
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation | The JsonPatchDocument providing the operations you wish to perform on the product. <br /> Value contains the value you wish to set on the property you're updating (used with operations \"add\" and \"replace\").<br /> Path contains the path to the property you're updating (e.g. Description). Every property in the model used for creation an updating can be used.<br /> Op contains the operation you wish to perform (\"add\",\"replace\",\"remove\").<br /> From is only used when using the \"move\" operation. It refers to the source path of the value to be moved. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductPatch(context.Background(), merchantProductNo).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductPatch`: SingleOfProductCreationResult
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantProductNo** | **string** | The MerchantProductNo of the product you wish to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operation** | [**[]Operation**](Operation.md) | The JsonPatchDocument providing the operations you wish to perform on the product. &lt;br /&gt; Value contains the value you wish to set on the property you&#39;re updating (used with operations \&quot;add\&quot; and \&quot;replace\&quot;).&lt;br /&gt; Path contains the path to the property you&#39;re updating (e.g. Description). Every property in the model used for creation an updating can be used.&lt;br /&gt; Op contains the operation you wish to perform (\&quot;add\&quot;,\&quot;replace\&quot;,\&quot;remove\&quot;).&lt;br /&gt; From is only used when using the \&quot;move\&quot; operation. It refers to the source path of the value to be moved. | 

### Return type

[**SingleOfProductCreationResult**](SingleOfProductCreationResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductPatchExtraDataItems

> SingleOfProductCreationResult ProductPatchExtraDataItems(ctx).MerchantProductExtraDataRequest(merchantProductExtraDataRequest).Execute()

Adds, updates, or deletes a custom attribute



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
	merchantProductExtraDataRequest := *openapiclient.NewMerchantProductExtraDataRequest("MerchantProductNo_example", []openapiclient.ProductExtraDataRequest{*openapiclient.NewProductExtraDataRequest("Op_example", "Key_example")}) // MerchantProductExtraDataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductPatchExtraDataItems(context.Background()).MerchantProductExtraDataRequest(merchantProductExtraDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductPatchExtraDataItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductPatchExtraDataItems`: SingleOfProductCreationResult
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductPatchExtraDataItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductPatchExtraDataItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantProductExtraDataRequest** | [**MerchantProductExtraDataRequest**](MerchantProductExtraDataRequest.md) |  | 

### Return type

[**SingleOfProductCreationResult**](SingleOfProductCreationResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

