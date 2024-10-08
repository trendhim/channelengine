/*
ChannelEngine Merchant API

Testing ProductsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package merchant

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/trendhim/channelengine/merchant"
)

func Test_merchant_ProductsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductsAPIService ProductBulkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductBulkDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductBulkPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductBulkPatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductBulkPatchExtraDataItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductBulkPatchExtraDataItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductsAPI.ProductDelete(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductFreeze", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductFreeze(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductGetByFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductGetByFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductGetByMerchantProductNo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductsAPI.ProductGetByMerchantProductNo(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductsAPI.ProductPatch(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductsAPIService ProductPatchExtraDataItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductsAPI.ProductPatchExtraDataItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
