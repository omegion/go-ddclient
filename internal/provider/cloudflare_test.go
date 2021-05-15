package provider_test

import (
	"context"
	"testing"

	"github.com/omegion/go-ddclient/internal/provider"
	"github.com/omegion/go-ddclient/internal/provider/mocks"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	zoneID     = "12345678"
	zoneName   = "example.com"
	recordID   = "123456"
	recordName = "test.example.com"
)

func TestCloudflareAPI_SetRecord_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	api := mocks.NewMockCloudflareAPIInterface(ctrl)

	var existingRecords []cloudflare.DNSRecord

	var createdRecord cloudflare.DNSRecordResponse

	api.EXPECT().ZoneIDByName(zoneName).Return(zoneID, nil)
	api.EXPECT().DNSRecords(ctx, zoneID, gomock.Any()).Return(existingRecords, nil)
	api.EXPECT().CreateDNSRecord(ctx, zoneID, gomock.Any()).Return(&createdRecord, nil)

	record := provider.DNSRecord{
		Name: recordName,
		Zone: provider.DNSZone{
			Name: zoneName,
		},
	}

	cloudflareAPI := provider.NewCloudflareAPI(api)
	err := cloudflareAPI.SetRecord(ctx, record)

	assert.Equal(t, nil, err)
}

func TestCloudflareAPI_SetRecord_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	api := mocks.NewMockCloudflareAPIInterface(ctrl)

	existingRecords := []cloudflare.DNSRecord{
		{
			ID: recordID,
		},
	}

	api.EXPECT().ZoneIDByName(zoneName).Return(zoneID, nil)
	api.EXPECT().DNSRecords(ctx, zoneID, gomock.Any()).Return(existingRecords, nil)
	api.EXPECT().UpdateDNSRecord(ctx, zoneID, recordID, gomock.Any()).Return(nil)

	record := provider.DNSRecord{
		Name:  recordName,
		Value: "",
		Zone: provider.DNSZone{
			Name: zoneName,
		},
	}

	cloudflareAPI := provider.NewCloudflareAPI(api)
	err := cloudflareAPI.SetRecord(ctx, record)

	assert.Equal(t, nil, err)
}
