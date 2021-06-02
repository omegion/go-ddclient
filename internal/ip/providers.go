package ip

import (
	"fmt"
	"net"
	"net/url"
)

const (
	// GoogleURL is Google provider url for ip checks.
	GoogleURL = "https://domains.google.com/checkip"
)

// Provider is an interface for IP providers.
type Provider interface {
	ExtractIP(data []byte) (net.IP, error)
	GetName() string
	GetURL() *url.URL
}

// GetProvider returns an IP provider with given name.
func GetProvider(name string) (Provider, error) {
	switch name {
	case "google":
		return NewGoogleIPProvider(), nil
	default:
		return nil, &NotSupported{Name: name}
	}
}

// GoogleIPProvider is IP provider.
type GoogleIPProvider struct {
	Name string
	URL  *url.URL
}

// NewGoogleIPProvider instantiate an IP provider.
func NewGoogleIPProvider() *GoogleIPProvider {
	providerURL, _ := url.Parse(GoogleURL)

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
