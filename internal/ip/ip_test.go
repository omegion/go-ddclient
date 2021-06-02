package ip_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/go-ddclient/internal/ip"
	"github.com/omegion/go-ddclient/internal/ip/mocks"
)

func TestCloudflareAPI_SetRecord_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockHTTPClient(ctrl)
	googleProvider := ip.NewGoogleIPProvider()

	currentIPAddress := "8.8.8.8"

	resp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(currentIPAddress))),
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

func TestCloudflareAPI_SetRecord_Update_Get_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockHTTPClient(ctrl)
	googleProvider := ip.NewGoogleIPProvider()

	currentIPAddress := "8.8.8.8"

	resp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(currentIPAddress))),
	}

	client.EXPECT().Get(googleProvider.GetURL().String()).Return(&resp, errors.New("custom error"))

	ipAddress := ip.IP{
		Client:   client,
		Provider: googleProvider,
	}

	err := ipAddress.Check()

	assert.EqualError(t, err, "custom error")
}

func TestCloudflareAPI_SetRecord_Update_ExtractIP_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockHTTPClient(ctrl)
	googleProvider := ip.NewGoogleIPProvider()

	currentIPAddress := "wrong8.8.8.8"

	resp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(currentIPAddress))),
	}

	client.EXPECT().Get(googleProvider.GetURL().String()).Return(&resp, nil)

	ipAddress := ip.IP{
		Client:   client,
		Provider: googleProvider,
	}

	err := ipAddress.Check()

	assert.EqualError(t, err, "invalid CIDR address: wrong8.8.8.8/24")
}
