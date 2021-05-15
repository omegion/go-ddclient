package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/omegion/go-ddclient/internal/provider"

	"github.com/spf13/cobra"
)

// Set sets DNS record with given provider and parameters.
func Set() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets DNS record to current IP address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			recordName, _ := cmd.Flags().GetString("record")
			recordValue, _ := cmd.Flags().GetString("value")
			zoneName, _ := cmd.Flags().GetString("zone")
			providerName, _ := cmd.Flags().GetString("provider")

			ctx := context.Background()

			record := provider.DNSRecord{
				Name:  recordName,
				Value: recordValue,
				Zone: provider.DNSZone{
					Name: zoneName,
				},
			}

			api, err := decideProvider(providerName)
			if err != nil {
				return err
			}

			err = api.SetRecord(ctx, record)
			if err != nil {
				return err
			}

			fmt.Printf("Record %s set to %s in zone %s\n", record.Name, record.Value, record.Zone.Name)

			return nil
		},
	}

	cmd.Flags().String("record", "", "Record Name")

	if err := cmd.MarkFlagRequired("record"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("value", "", "Value")

	if err := cmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("zone", "", "Zone Name")

	if err := cmd.MarkFlagRequired("zone"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("provider", "", "DNS Provider")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func decideProvider(name string) (provider.API, error) {
	//nolint:gocritic // will be extended soon.
	switch name {
	case "cloudflare":
		return provider.SetupCloudflareAPI()
	}

	return provider.CloudflareAPI{}, &provider.NotSupported{Name: name}
}
