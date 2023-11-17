# \ChannelsAPI

All URIs are relative to *https://demo.channelengine.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelPluginsGet**](ChannelsAPI.md#ChannelPluginsGet) | **Get** /v2/channels | Gets channels



## ChannelPluginsGet

> CollectionOfChannelGlobalChannelResponse ChannelPluginsGet(ctx).Execute()

Gets channels



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelsAPI.ChannelPluginsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ChannelPluginsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChannelPluginsGet`: CollectionOfChannelGlobalChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ChannelPluginsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChannelPluginsGetRequest struct via the builder pattern


### Return type

[**CollectionOfChannelGlobalChannelResponse**](CollectionOfChannelGlobalChannelResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

