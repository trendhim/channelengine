/*
ChannelEngine Channel API

Testing ReturnAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package channel

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func Test_channel_ReturnAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReturnAPIService ReturnDeclareForChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnAPI.ReturnDeclareForChannel(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnAPIService ReturnGetDeclaredByMerchant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnAPI.ReturnGetDeclaredByMerchant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
