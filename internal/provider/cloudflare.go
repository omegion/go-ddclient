package provider

import (
	"context"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/cloudflare_mock.go -package=mocks github.com/omegion/go-ddclient/internal/provider CloudflareAPIInterface
// CloudflareAPIInterface is an interface for Cloudflare api for mocking.
type CloudflareAPIInterface interface {
	ZoneIDByName(zoneName string) (string, error)
	DNSRecords(ctx context.Context, zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
	UpdateDNSRecord(ctx context.Context, zoneID, recordID string, rr cloudflare.DNSRecord) error
	CreateDNSRecord(ctx context.Context, zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error)
}

// CloudflareAPI is API entrypoint for Cloudflare.
type CloudflareAPI struct {
	client CloudflareAPIInterface
}

// SetupCloudflareAPI setups CloudflareAPI.
func SetupCloudflareAPI() (API, error) {
	client, err := cloudflare.NewWithAPIToken(os.Getenv("CF_API_KEY"))
	if err != nil {
		return CloudflareAPI{}, err
	}

	return NewCloudflareAPI(client), nil
}

// NewCloudflareAPI is a factory for NewCloudflareAPI.
func NewCloudflareAPI(client CloudflareAPIInterface) *CloudflareAPI {
	return &CloudflareAPI{client: client}
}

// SetRecord sets a record in Cloudflare.
func (a CloudflareAPI) SetRecord(ctx context.Context, record DNSRecord) error {
	zoneID, err := a.client.ZoneIDByName(record.Zone.Name)
	if err != nil {
		return err
	}

	cRecord := cloudflare.DNSRecord{
		Type: "A",
		Name: record.Name,
		TTL:  0,
	}

	existsRecords, err := a.client.DNSRecords(ctx, zoneID, cRecord)
	if err != nil {
		return err
	}

	cRecord.Content = record.Value

	if len(existsRecords) > 0 {
		err = a.client.UpdateDNSRecord(ctx, zoneID, existsRecords[0].ID, cRecord)
		if err != nil {
			return err
		}

		return nil
	}

	_, err = a.client.CreateDNSRecord(ctx, zoneID, cRecord)
	if err != nil {
		return err
	}

	return nil
}
