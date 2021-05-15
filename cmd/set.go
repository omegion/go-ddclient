package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/omegion/go-ddclient/internal/ip"
	"github.com/omegion/go-ddclient/internal/provider"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupAddCommand(cmd *cobra.Command) {
	cmd.Flags().String("record", "", "Record Name")

	if err := cmd.MarkFlagRequired("record"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("zone", "", "Zone Name")

	if err := cmd.MarkFlagRequired("zone"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("dns-provider", "", "DNS Provider")

	if err := cmd.MarkFlagRequired("dns-provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("ip-provider", "google", "IP Provider")
	cmd.Flags().Int("interval", 1, "Interval in Minutes")
	cmd.Flags().Bool("daemon", false, "Daemon")
}

//nolint:funlen // some commands can be longer.
// Set sets DNS record with given provider and parameters.
func Set() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets DNS record to current IP address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			recordName, _ := cmd.Flags().GetString("record")
			zoneName, _ := cmd.Flags().GetString("zone")
			DNSProvider, _ := cmd.Flags().GetString("dns-provider")
			IPProviderName, _ := cmd.Flags().GetString("ip-provider")
			interval, _ := cmd.Flags().GetInt("interval")
			daemon, _ := cmd.Flags().GetBool("daemon")

			ctx := context.Background()

			IPProvider, err := decideIPProvider(IPProviderName)
			if err != nil {
				return err
			}

			api, err := decideDNSProvider(DNSProvider)
			if err != nil {
				return err
			}

			ipAddress := ip.IP{
				Client:   &http.Client{},
				Provider: IPProvider,
			}

			if daemon {
				ticker := time.NewTicker(time.Duration(interval) * time.Minute)
				quit := make(chan struct{})
				go func() {
					for {
						select {
						case <-ticker.C:
							record, setErr := setRecord(ctx, recordName, zoneName, ipAddress, api)
							if err != nil {
								log.Fatal(setErr)
								close(quit)
							}

							log.Infoln(fmt.Sprintf("Record %s set to %s in zone %s.", record.Name, record.Value, record.Zone.Name))
						case <-quit:
							ticker.Stop()
							return
						}
					}
				}()
				select {}
			}

			record, err := setRecord(ctx, recordName, zoneName, ipAddress, api)
			if err != nil {
				return err
			}

			log.Infoln(fmt.Sprintf("Record %s set to %s in zone %s.", record.Name, record.Value, record.Zone.Name))

			return nil
		},
	}

	setupAddCommand(cmd)

	return cmd
}

func decideDNSProvider(name string) (provider.API, error) {
	//nolint:gocritic // will be extended soon.
	switch name {
	case "cloudflare":
		return provider.SetupCloudflareAPI()
	}

	return provider.CloudflareAPI{}, &provider.NotSupported{Name: name}
}

func decideIPProvider(name string) (ip.Provider, error) {
	providers := ip.AllProviders()
	if prv, ok := providers[name]; ok {
		return prv, nil
	}

	return ip.GoogleIPProvider{}, &ip.NotSupported{Name: name}
}

func setRecord(
	ctx context.Context,
	recordName,
	zoneName string,
	ipAddress ip.IP,
	api provider.API,
) (provider.DNSRecord, error) {
	err := ipAddress.Check()
	if err != nil {
		return provider.DNSRecord{}, err
	}

	record := provider.DNSRecord{
		Name:  recordName,
		Value: ipAddress.Address.String(),
		Zone: provider.DNSZone{
			Name: zoneName,
		},
	}

	err = api.SetRecord(ctx, record)
	if err != nil {
		return provider.DNSRecord{}, err
	}

	return record, nil
}
