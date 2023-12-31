/*
ChannelEngine Merchant API

Testing ShipmentAPIService

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

func Test_merchant_ShipmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShipmentAPIService ShipmentCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentCreateForChannelMethod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentCreateForChannelMethod(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentGetShipmentLabelCarriers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantOrderNo string

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentGetShipmentLabelCarriers(context.Background(), merchantOrderNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentShippingLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantShipmentNo string

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentShippingLabel(context.Background(), merchantShipmentNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantShipmentNo string

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentUpdate(context.Background(), merchantShipmentNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
