package ip_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/omegion/go-ddclient/internal/ip"
	"github.com/omegion/go-ddclient/internal/ip/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCloudflareAPI_SetRecord_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockHTTPClient(ctrl)
	googleProvider := ip.NewGoogleIPProvider()

	currentIPAddress := "8.8.8.8"
	r := ioutil.NopCloser(bytes.NewReader([]byte(currentIPAddress)))

	resp := http.Response{
		StatusCode: 200,
		Body:       r,
	}

	client.EXPECT().Get(googleProvider.GetURL().String()).Return(&resp, nil)

	ipAddress := ip.IP{
		Client:   client,
		Provider: googleProvider,
	}

	err := ipAddress.Check()

	assert.Equal(t, nil, err)
	assert.Equal(t, currentIPAddress, ipAddress.Address.String())
}
