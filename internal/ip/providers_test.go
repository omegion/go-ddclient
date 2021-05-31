package ip_test

import (
	"testing"

	"github.com/omegion/go-ddclient/internal/ip"

	"github.com/stretchr/testify/assert"
)

func TestAllProviders(t *testing.T) {
	expectedProviders := map[string]string{
		"google": ip.GoogleURL,
	}

	for _, prov := range ip.AllProviders() {
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
