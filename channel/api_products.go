/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ProductsAPIService ProductsAPI service
type ProductsAPIService service

type ProductsAPIProductAcknowledgeDataChangesRequest struct {
	ctx context.Context
	ApiService *ProductsAPIService
	channelProcessedChangesRequest *ChannelProcessedChangesRequest
}

// The merchant references of the products which have been&lt;br /&gt; successfully created, updated or deleted (after a call to &#39;GetDataChanges&#39;).
func (r ProductsAPIProductAcknowledgeDataChangesRequest) ChannelProcessedChangesRequest(channelProcessedChangesRequest ChannelProcessedChangesRequest) ProductsAPIProductAcknowledgeDataChangesRequest {
	r.channelProcessedChangesRequest = &channelProcessedChangesRequest
	return r
}

func (r ProductsAPIProductAcknowledgeDataChangesRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.ProductAcknowledgeDataChangesExecute(r)
}

/*
ProductAcknowledgeDataChanges Acknowledge product data changes

This endpoint should be called after a call to GET 'v2/products/data'. After a call to<br />this endpoint ChannelEngine 'knows' which products are known to the channel (and the last<br />time they have been updated) and is therefore able to only return the products<br />that really have changed since the last call to POST 'v2/products/data'.<br />The supplied ChannelProductNo will be saved<br />in our database and is supposed to be unique, the 'Updated' and 'Removed'<br />fields consist of ChannelReferences which are sent in a previous call<br />within the field 'Created'.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductsAPIProductAcknowledgeDataChangesRequest
*/
func (a *ProductsAPIService) ProductAcknowledgeDataChanges(ctx context.Context) ProductsAPIProductAcknowledgeDataChangesRequest {
	return ProductsAPIProductAcknowledgeDataChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *ProductsAPIService) ProductAcknowledgeDataChangesExecute(r ProductsAPIProductAcknowledgeDataChangesRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductAcknowledgeDataChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/products/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.channelProcessedChangesRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ProductsAPIProductAcknowledgeOfferChangesRequest struct {
	ctx context.Context
	ApiService *ProductsAPIService
	keyIsMpn *bool
	requestBody *[]string
}

// If set to true, changes has to be a list of merchant references instead of channel references.
func (r ProductsAPIProductAcknowledgeOfferChangesRequest) KeyIsMpn(keyIsMpn bool) ProductsAPIProductAcknowledgeOfferChangesRequest {
	r.keyIsMpn = &keyIsMpn
	return r
}

// The channel references of the updated products.
func (r ProductsAPIProductAcknowledgeOfferChangesRequest) RequestBody(requestBody []string) ProductsAPIProductAcknowledgeOfferChangesRequest {
	r.requestBody = &requestBody
	return r
}

func (r ProductsAPIProductAcknowledgeOfferChangesRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.ProductAcknowledgeOfferChangesExecute(r)
}

/*
ProductAcknowledgeOfferChanges Acknowledge product offer changes

After a call to GET 'v2/products/offers' this endpoint should be called with the<br />ChannelProductNo of the products that are successfully updated.<br />Please see 'v2/products/data' and 'v2/products/data' for documentation.<br />In advanced cases, the MerchantProductNo is used for this.<br />In that case, bool keyIsMpn should be true.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductsAPIProductAcknowledgeOfferChangesRequest
*/
func (a *ProductsAPIService) ProductAcknowledgeOfferChanges(ctx context.Context) ProductsAPIProductAcknowledgeOfferChangesRequest {
	return ProductsAPIProductAcknowledgeOfferChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *ProductsAPIService) ProductAcknowledgeOfferChangesExecute(r ProductsAPIProductAcknowledgeOfferChangesRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductAcknowledgeOfferChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/products/offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.keyIsMpn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyIsMpn", r.keyIsMpn, "")
	} else {
		var defaultValue bool = false
		r.keyIsMpn = &defaultValue
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
	localVarPostBody = r.requestBody
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

type ProductsAPIProductGetDataChangesRequest struct {
	ctx context.Context
	ApiService *ProductsAPIService
	maxCount *int32
	stripHtml *bool
	page *int32
	pageSize *int32
}

// Optional - limit the amount of products returned for each field&lt;br /&gt; (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number.
func (r ProductsAPIProductGetDataChangesRequest) MaxCount(maxCount int32) ProductsAPIProductGetDataChangesRequest {
	r.maxCount = &maxCount
	return r
}

// Optional - strips html by default on all fields
func (r ProductsAPIProductGetDataChangesRequest) StripHtml(stripHtml bool) ProductsAPIProductGetDataChangesRequest {
	r.stripHtml = &stripHtml
	return r
}

// Optional - for default is first page returned
func (r ProductsAPIProductGetDataChangesRequest) Page(page int32) ProductsAPIProductGetDataChangesRequest {
	r.page = &page
	return r
}

// Optional - amount of products returned, if not provided return all products
func (r ProductsAPIProductGetDataChangesRequest) PageSize(pageSize int32) ProductsAPIProductGetDataChangesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ProductsAPIProductGetDataChangesRequest) Execute() (*CollectionChangesOfChannelProductChangesResponse, *http.Response, error) {
	return r.ApiService.ProductGetDataChangesExecute(r)
}

/*
ProductGetDataChanges Gets product data changes

Get all products which have changes since the post http call to POST 'v2/products/data'.<br />The response contains the products which should be created, updated or removed from the channel.<br />After the products have been received and processed successfully 'v2/products/data' should<br />be called with the merchant references of the products which have been processed (see POST 'v2/products/data').<br />ChannelEngine will save this information to make sure that the next call to GET 'v2/products/data' only returns the<br />products that really have changes (and therefore should be created, updated or deleted).<br />A channel willing to integrate with channelengine should therefore only do the following things:<br /> 1. Periodically poll 'v2/products/data' for changes.<br /> 2. If any products are returned, save, update or remove these products.<br /> 3. Send the merchant references of the products that have succesfully been processed<br /> in step 2 to POST 'v2/products/data'.<br /><br />These three simple steps will make sure that the products on the channel will be synchronized<br />with the products on ChannelEngine. ChannelEngine will use the API key to determine the customer<br />whose products should be returned. Note that child products are only returned if their parent product<br />has been acknowledged in a previous transaction. This way ChannelEngine knows the value of<br />'ChannelParentReference'.<br /><br />For self-integrated channels do not use "retry data export for all (including deleted)".

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductsAPIProductGetDataChangesRequest
*/
func (a *ProductsAPIService) ProductGetDataChanges(ctx context.Context) ProductsAPIProductGetDataChangesRequest {
	return ProductsAPIProductGetDataChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionChangesOfChannelProductChangesResponse
func (a *ProductsAPIService) ProductGetDataChangesExecute(r ProductsAPIProductGetDataChangesRequest) (*CollectionChangesOfChannelProductChangesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionChangesOfChannelProductChangesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductGetDataChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/products/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxCount", r.maxCount, "")
	}
	if r.stripHtml != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stripHtml", r.stripHtml, "")
	} else {
		var defaultValue bool = true
		r.stripHtml = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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

type ProductsAPIProductGetDataChangesFullRequest struct {
	ctx context.Context
	ApiService *ProductsAPIService
	productType *DataChangesProductType
	maxCount *int32
	stripHtml *bool
	page *int32
	pageSize *int32
}

// Optional - Type of products
func (r ProductsAPIProductGetDataChangesFullRequest) ProductType(productType DataChangesProductType) ProductsAPIProductGetDataChangesFullRequest {
	r.productType = &productType
	return r
}

// Optional - limit the amount of products returned for each field&lt;br /&gt; (ToBeCreated, ToBeUpdated, ToBeRemoved) to this number.
func (r ProductsAPIProductGetDataChangesFullRequest) MaxCount(maxCount int32) ProductsAPIProductGetDataChangesFullRequest {
	r.maxCount = &maxCount
	return r
}

// Optional - strips html by default on all fields
func (r ProductsAPIProductGetDataChangesFullRequest) StripHtml(stripHtml bool) ProductsAPIProductGetDataChangesFullRequest {
	r.stripHtml = &stripHtml
	return r
}

// Optional - for default is first page returned
func (r ProductsAPIProductGetDataChangesFullRequest) Page(page int32) ProductsAPIProductGetDataChangesFullRequest {
	r.page = &page
	return r
}

// Optional - amount of products returned, if not provided return all products
func (r ProductsAPIProductGetDataChangesFullRequest) PageSize(pageSize int32) ProductsAPIProductGetDataChangesFullRequest {
	r.pageSize = &pageSize
	return r
}

func (r ProductsAPIProductGetDataChangesFullRequest) Execute() (*CollectionChangesOfChannelProductChangesResponse, *http.Response, error) {
	return r.ApiService.ProductGetDataChangesFullExecute(r)
}

/*
ProductGetDataChangesFull Gets product data changes with an optional product type filter. If you select product type products will be filtered by it.  If you won't pass product type you will get products with types: CHILD, PARENT, GRANDPARENT, SINGLE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductsAPIProductGetDataChangesFullRequest
*/
func (a *ProductsAPIService) ProductGetDataChangesFull(ctx context.Context) ProductsAPIProductGetDataChangesFullRequest {
	return ProductsAPIProductGetDataChangesFullRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionChangesOfChannelProductChangesResponse
func (a *ProductsAPIService) ProductGetDataChangesFullExecute(r ProductsAPIProductGetDataChangesFullRequest) (*CollectionChangesOfChannelProductChangesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionChangesOfChannelProductChangesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductGetDataChangesFull")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/products/data/full"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productType", r.productType, "")
	}
	if r.maxCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxCount", r.maxCount, "")
	}
	if r.stripHtml != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stripHtml", r.stripHtml, "")
	} else {
		var defaultValue bool = true
		r.stripHtml = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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

type ProductsAPIProductGetOfferChangesRequest struct {
	ctx context.Context
	ApiService *ProductsAPIService
	page *int32
	pageSize *int32
}

// Optional - for default is first page returned. Not recomended could return the same offers in different pages.
func (r ProductsAPIProductGetOfferChangesRequest) Page(page int32) ProductsAPIProductGetOfferChangesRequest {
	r.page = &page
	return r
}

// Optional - amount of products returned, if not provided return all products. Not recomended could return the same offers in different pages.
func (r ProductsAPIProductGetOfferChangesRequest) PageSize(pageSize int32) ProductsAPIProductGetOfferChangesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ProductsAPIProductGetOfferChangesRequest) Execute() (*CollectionOfChannelOfferResponse, *http.Response, error) {
	return r.ApiService.ProductGetOfferChangesExecute(r)
}

/*
ProductGetOfferChanges Gets product offer changes

GET 'v2/products/offers' and POST 'v2/products/offers' closely resemble GET 'v2/products/data' and POST 'v2/products/data'. See the 'v2/products/data'<br />endpoints for documentation. The difference between both endpoints is that 'v2/products/offers' only returns Price and Stock updates and can (and should)<br />therefore be called more often to keep this information (which might change frequently) up to date.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductsAPIProductGetOfferChangesRequest
*/
func (a *ProductsAPIService) ProductGetOfferChanges(ctx context.Context) ProductsAPIProductGetOfferChangesRequest {
	return ProductsAPIProductGetOfferChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfChannelOfferResponse
func (a *ProductsAPIService) ProductGetOfferChangesExecute(r ProductsAPIProductGetOfferChangesRequest) (*CollectionOfChannelOfferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfChannelOfferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductGetOfferChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/products/offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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
