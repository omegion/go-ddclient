package ip_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/go-ddclient/internal/ip"
)

func TestAllProviders(t *testing.T) {
	expectedProviders := map[string]string{
		"google": ip.GoogleURL,
	}

	for name := range expectedProviders {
		prov, err := ip.GetProvider(name)
		assert.NoError(t, err)

		expectedIPAddress := []byte("8.8.8.8")

		ipAddress, err := prov.ExtractIP(expectedIPAddress)

		assert.NoError(t, err)
		assert.Equal(t, expectedProviders[prov.GetName()], prov.GetURL().String())
		assert.Equal(t, "8.8.8.8", ipAddress.String())

		malformedExpectedIPAddress := []byte("wrong8.8.8.8")

		_, err = prov.ExtractIP(malformedExpectedIPAddress)

		assert.EqualError(t, err, "invalid CIDR address: wrong8.8.8.8/24")
	}
}

func TestAllProviders_Failure(t *testing.T) {
	provider, err := ip.GetProvider("unknown-provider")

	assert.EqualError(t, err, "Provider unknown-provider not supported.")
	assert.Equal(t, nil, provider)
}
