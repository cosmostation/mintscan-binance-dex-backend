package client

import (
	"os"
	"testing"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"
	"github.com/stretchr/testify/require"
)

var client *Client

func TestMain(m *testing.M) {
	config := config.ParseConfig()
	client = NewClient(config.Node, config.Market)

	os.Exit(m.Run())
}

func TestGetOrder(t *testing.T) {
	testCases := []struct {
		info    string
		orderID string
	}{
		{"BEP8 Mini Token Order ID", "DDFA476FA637F3061FB31CFC7951367DD220BA84-1098"},
		{"BEP2 Token Order ID", "1468EE412C3ADC9CFF3EF31ADC7EDD288F5E208E-14509857"},
	}

	for _, tc := range testCases {
		order, err := client.GetOrder(tc.orderID)
		require.NoError(t, err)

		require.NotNil(t, order)
	}
}
