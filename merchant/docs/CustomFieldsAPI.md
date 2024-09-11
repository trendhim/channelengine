# \CustomFieldsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomFieldsGetCustomFields**](CustomFieldsAPI.md#CustomFieldsGetCustomFields) | **Get** /v2/custom-fields | Gets custom fields



## CustomFieldsGetCustomFields

> CollectionOfCustomFieldResponse CustomFieldsGetCustomFields(ctx).PageIndex(pageIndex).PageSize(pageSize).Execute()

Gets custom fields

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
	pageIndex := int32(56) // int32 | A page index to get the items (starts from 0) (optional)
	pageSize := int32(56) // int32 | Number of items to return (default 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsGetCustomFields(context.Background()).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsGetCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsGetCustomFields`: CollectionOfCustomFieldResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsGetCustomFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsGetCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **int32** | A page index to get the items (starts from 0) | 
 **pageSize** | **int32** | Number of items to return (default 100) | 

### Return type

[**CollectionOfCustomFieldResponse**](CollectionOfCustomFieldResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

