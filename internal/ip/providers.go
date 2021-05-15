package ip

import (
	"fmt"
	"net"
	"net/url"

	log "github.com/sirupsen/logrus"
)

const (
	google = "https://domains.google.com/checkip"
)

// Provider is an interface for IP providers.
type Provider interface {
	ExtractIP(data []byte) (net.IP, error)
	GetName() string
	GetURL() *url.URL
}

// AllProviders returns all supported IP providers.
func AllProviders() map[string]Provider {
	providers := make(map[string]Provider)

	providers["google"] = NewGoogleIPProvider()

	log.Debugln(fmt.Sprintf("%d IP provider(s) are loaded.", len(providers)))

	return providers
}

// GoogleIPProvider is IP provider.
type GoogleIPProvider struct {
	Name string
	URL  *url.URL
}

// NewGoogleIPProvider instantiate an IP provider.
func NewGoogleIPProvider() *GoogleIPProvider {
	providerURL, _ := url.Parse(google)

	return &GoogleIPProvider{
		Name: "google",
		URL:  providerURL,
	}
}

// GetName gets the provider name.
func (p GoogleIPProvider) GetName() string {
	return p.Name
}

// GetURL gets the provider url.
func (p GoogleIPProvider) GetURL() *url.URL {
	return p.URL
}

// ExtractIP extracts IP address from provider data.
func (p GoogleIPProvider) ExtractIP(data []byte) (net.IP, error) {
	cidr := fmt.Sprintf("%s/24", data)

	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return net.IP{}, err
	}

	return ip, nil
}
