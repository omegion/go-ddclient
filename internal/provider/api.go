package provider

import (
	"context"
)

//go:generate mockgen -destination=mocks/api_mock.go -package=mocks github.com/omegion/go-ddclient/internal/provider API
// API is an interface for all providers.
type API interface {
	SetRecord(ctx context.Context, record DNSRecord) error
}

// DNSZone is zone struct for all providers.
type DNSZone struct {
	Name string
}

// DNSRecord is record struct for all providers.
type DNSRecord struct {
	Name  string
	Value string
	Zone  DNSZone
}

// GetProvider returns an DNS provider with given name.
func GetProvider(name string) (API, error) {
	switch name {
	case "cloudflare":
		return SetupCloudflareAPI()
	default:
		return CloudflareAPI{}, &NotSupported{Name: name}
	}
}
