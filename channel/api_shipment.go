/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
	"os"
)


// ShipmentAPIService ShipmentAPI service
type ShipmentAPIService service

type ShipmentAPIShipmentIndexRequest struct {
	ctx context.Context
	ApiService *ShipmentAPIService
	fromDate *time.Time
	toDate *time.Time
	channelOrderNos *[]string
	page *int32
}

// Filter on the creation date, starting from this date. This date is inclusive.
func (r ShipmentAPIShipmentIndexRequest) FromDate(fromDate time.Time) ShipmentAPIShipmentIndexRequest {
	r.fromDate = &fromDate
	return r
}

// Filter on the creation date, until this date. This date is exclusive.
func (r ShipmentAPIShipmentIndexRequest) ToDate(toDate time.Time) ShipmentAPIShipmentIndexRequest {
	r.toDate = &toDate
	return r
}

// Filter on the unique references (ids) as used by the channel.
func (r ShipmentAPIShipmentIndexRequest) ChannelOrderNos(channelOrderNos []string) ShipmentAPIShipmentIndexRequest {
	r.channelOrderNos = &channelOrderNos
	return r
}

// The page to filter on. Starts at 1.
func (r ShipmentAPIShipmentIndexRequest) Page(page int32) ShipmentAPIShipmentIndexRequest {
	r.page = &page
	return r
}

func (r ShipmentAPIShipmentIndexRequest) Execute() (*CollectionOfChannelShipmentResponse, *http.Response, error) {
	return r.ApiService.ShipmentIndexExecute(r)
}

/*
ShipmentIndex Gets shipments

Gets all shipments created since the supplied date with status CLOSED.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShipmentAPIShipmentIndexRequest
*/
func (a *ShipmentAPIService) ShipmentIndex(ctx context.Context) ShipmentAPIShipmentIndexRequest {
	return ShipmentAPIShipmentIndexRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfChannelShipmentResponse
func (a *ShipmentAPIService) ShipmentIndexExecute(r ShipmentAPIShipmentIndexRequest) (*CollectionOfChannelShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfChannelShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentAPIService.ShipmentIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/shipments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "")
	}
	if r.channelOrderNos != nil {
		t := *r.channelOrderNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "channelOrderNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "channelOrderNos", t, "multi")
		}
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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

type ShipmentAPIShipmentShippingLabelRequest struct {
	ctx context.Context
	ApiService *ShipmentAPIService
	merchantShipmentNo string
}

func (r ShipmentAPIShipmentShippingLabelRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.ShipmentShippingLabelExecute(r)
}

/*
ShipmentShippingLabel Gets a shipping label

 Downloads the shipping label for the shipment.<br /> <br /> **NB:** it may take some time between the creation of the shipment and the availability of the label.<br />A "404 not found" error might indicate that the label is not available yet.<br />A "410 gone" the shipping label is not available anymore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantShipmentNo The unique shipment reference as used by the merchant.
 @return ShipmentAPIShipmentShippingLabelRequest
*/
func (a *ShipmentAPIService) ShipmentShippingLabel(ctx context.Context, merchantShipmentNo string) ShipmentAPIShipmentShippingLabelRequest {
	return ShipmentAPIShipmentShippingLabelRequest{
		ApiService: a,
		ctx: ctx,
		merchantShipmentNo: merchantShipmentNo,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ShipmentAPIService) ShipmentShippingLabelExecute(r ShipmentAPIShipmentShippingLabelRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentAPIService.ShipmentShippingLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/orders/{merchantShipmentNo}/shippinglabel"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantShipmentNo"+"}", url.PathEscape(parameterValueToString(r.merchantShipmentNo, "merchantShipmentNo")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.shippingLabel", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
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
