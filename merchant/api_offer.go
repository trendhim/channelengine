/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// OfferAPIService OfferAPI service
type OfferAPIService service

type OfferAPIOfferGetStockRequest struct {
	ctx context.Context
	ApiService *OfferAPIService
	stockLocationIds *[]int32
	skus *[]string
	pageIndex *int32
	pageSize *int32
}

// The ChannelEngine id of the stock location(s).
func (r OfferAPIOfferGetStockRequest) StockLocationIds(stockLocationIds []int32) OfferAPIOfferGetStockRequest {
	r.stockLocationIds = &stockLocationIds
	return r
}

// List of your products&#39; sku&#39;s.
func (r OfferAPIOfferGetStockRequest) Skus(skus []string) OfferAPIOfferGetStockRequest {
	r.skus = &skus
	return r
}

// A page index to get the items (starts from 0)
func (r OfferAPIOfferGetStockRequest) PageIndex(pageIndex int32) OfferAPIOfferGetStockRequest {
	r.pageIndex = &pageIndex
	return r
}

// Number of items to return (default 100)
func (r OfferAPIOfferGetStockRequest) PageSize(pageSize int32) OfferAPIOfferGetStockRequest {
	r.pageSize = &pageSize
	return r
}

func (r OfferAPIOfferGetStockRequest) Execute() (*CollectionOfMerchantOfferGetStockResponse, *http.Response, error) {
	return r.ApiService.OfferGetStockExecute(r)
}

/*
OfferGetStock Gets product stock across all warehouses

Gets the stock available in the warehouses. The warehouses must be set up as stock locations on ChannelEngine.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OfferAPIOfferGetStockRequest
*/
func (a *OfferAPIService) OfferGetStock(ctx context.Context) OfferAPIOfferGetStockRequest {
	return OfferAPIOfferGetStockRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfMerchantOfferGetStockResponse
func (a *OfferAPIService) OfferGetStockExecute(r OfferAPIOfferGetStockRequest) (*CollectionOfMerchantOfferGetStockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfMerchantOfferGetStockResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferAPIService.OfferGetStock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/offer/stock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockLocationIds == nil {
		return localVarReturnValue, nil, reportError("stockLocationIds is required and must be specified")
	}

	if r.skus != nil {
		t := *r.skus
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "skus", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "skus", t, "multi")
		}
	}
	{
		t := *r.stockLocationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "stockLocationIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "stockLocationIds", t, "multi")
		}
	}
	if r.pageIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OfferAPIOfferStockPriceUpdateRequest struct {
	ctx context.Context
	ApiService *OfferAPIService
	merchantStockPriceUpdateRequest *[]MerchantStockPriceUpdateRequest
}

// References to the products that should be updated, and the new values&lt;br /&gt;for the stock or price fields. It is possible to supply only one of the two fields&lt;br /&gt;or both.
func (r OfferAPIOfferStockPriceUpdateRequest) MerchantStockPriceUpdateRequest(merchantStockPriceUpdateRequest []MerchantStockPriceUpdateRequest) OfferAPIOfferStockPriceUpdateRequest {
	r.merchantStockPriceUpdateRequest = &merchantStockPriceUpdateRequest
	return r
}

func (r OfferAPIOfferStockPriceUpdateRequest) Execute() (*SingleOfDictionaryOfStringAndListOfString, *http.Response, error) {
	return r.ApiService.OfferStockPriceUpdateExecute(r)
}

/*
OfferStockPriceUpdate Updates stock and price

Updates product stock and price.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OfferAPIOfferStockPriceUpdateRequest
*/
func (a *OfferAPIService) OfferStockPriceUpdate(ctx context.Context) OfferAPIOfferStockPriceUpdateRequest {
	return OfferAPIOfferStockPriceUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SingleOfDictionaryOfStringAndListOfString
func (a *OfferAPIService) OfferStockPriceUpdateExecute(r OfferAPIOfferStockPriceUpdateRequest) (*SingleOfDictionaryOfStringAndListOfString, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SingleOfDictionaryOfStringAndListOfString
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferAPIService.OfferStockPriceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/offer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.merchantStockPriceUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("merchantStockPriceUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.merchantStockPriceUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OfferAPIOfferStockUpdateRequest struct {
	ctx context.Context
	ApiService *OfferAPIService
	merchantOfferStockUpdateRequest *[]MerchantOfferStockUpdateRequest
}

// References to the new values for the stock fields.
func (r OfferAPIOfferStockUpdateRequest) MerchantOfferStockUpdateRequest(merchantOfferStockUpdateRequest []MerchantOfferStockUpdateRequest) OfferAPIOfferStockUpdateRequest {
	r.merchantOfferStockUpdateRequest = &merchantOfferStockUpdateRequest
	return r
}

func (r OfferAPIOfferStockUpdateRequest) Execute() (*SingleOfDictionaryOfStringAndListOfString, *http.Response, error) {
	return r.ApiService.OfferStockUpdateExecute(r)
}

/*
OfferStockUpdate Updates stock

Updates product stock.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OfferAPIOfferStockUpdateRequest
*/
func (a *OfferAPIService) OfferStockUpdate(ctx context.Context) OfferAPIOfferStockUpdateRequest {
	return OfferAPIOfferStockUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SingleOfDictionaryOfStringAndListOfString
func (a *OfferAPIService) OfferStockUpdateExecute(r OfferAPIOfferStockUpdateRequest) (*SingleOfDictionaryOfStringAndListOfString, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SingleOfDictionaryOfStringAndListOfString
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferAPIService.OfferStockUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/offer/stock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.merchantOfferStockUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("merchantOfferStockUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.merchantOfferStockUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
