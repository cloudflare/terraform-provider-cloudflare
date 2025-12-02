package address_map_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_address_map", &resource.Sweeper{
		Name: "cloudflare_address_map",
		F:    testSweepCloudflareAddressMaps,
	})
}

func testSweepCloudflareAddressMaps(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping address maps sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	addressMaps, err := client.ListAddressMaps(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListAddressMapsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch address maps: %s", err))
		return fmt.Errorf("failed to fetch address maps: %w", err)
	}

	if len(addressMaps) == 0 {
		tflog.Info(ctx, "No address maps to sweep")
		return nil
	}

	for _, addressMap := range addressMaps {
		// Use standard filtering helper on the description field
		if addressMap.Description == nil || !utils.ShouldSweepResource(*addressMap.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting address map: %s (account: %s)", addressMap.ID, accountID))
		err := client.DeleteAddressMap(ctx, cloudflare.AccountIdentifier(accountID), addressMap.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete address map %s: %s", addressMap.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted address map: %s", addressMap.ID))
	}

	return nil
}

func TestAccCloudflareAddressMap(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Pending permission fixes for IP delegation.")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_address_map.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	altZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: generateCloudflareAddressMapConfig(rnd, accountID, nil, nil, false, nil, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "default_sni", ""),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "ips.#", "0"),
					resource.TestCheckResourceAttr(name, "memberships.#", "0"),
				),
			},
			{
				Config: generateCloudflareAddressMapConfig(
					rnd,
					accountID,
					cloudflare.StringPtr(rnd),
					cloudflare.StringPtr("*."+domain),
					true,
					[]string{"199.212.90.4", "199.212.90.5"},
					[]cloudflare.AddressMapMembershipContainer{
						{Identifier: zoneID, Kind: cloudflare.AddressMapMembershipZone},
						{Identifier: altZoneID, Kind: cloudflare.AddressMapMembershipZone},
					}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "default_sni", "*."+domain),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "ips.#", "2"),
					resource.TestCheckResourceAttr(name, "ips.0.ip", "199.212.90.4"),
					resource.TestCheckResourceAttr(name, "ips.1.ip", "199.212.90.5"),
					resource.TestCheckResourceAttr(name, "memberships.#", "2"),
					resource.TestCheckResourceAttr(name, "memberships.0.identifier", zoneID),
					resource.TestCheckResourceAttr(name, "memberships.0.kind", "zone"),
					resource.TestCheckResourceAttr(name, "memberships.1.identifier", altZoneID),
					resource.TestCheckResourceAttr(name, "memberships.1.kind", "zone"),
				),
			},
			{
				Config: generateCloudflareAddressMapConfig(rnd, accountID, cloudflare.StringPtr(""), cloudflare.StringPtr("*."+domain), false, nil, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "default_sni", "*."+domain),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "ips.#", "0"),
					resource.TestCheckResourceAttr(name, "memberships.#", "0"),
				),
			},
		},
	})
}

func generateCloudflareAddressMapConfig(rnd, accountId string, desc, sni *string, enabled bool, ips []string, memberships []cloudflare.AddressMapMembershipContainer) string {
	descFragment := "# Description"
	if desc != nil {
		descFragment = fmt.Sprintf("description = %q", *desc)
	}
	sniFragment := "# SNI"
	if sni != nil {
		sniFragment = fmt.Sprintf("default_sni = %q", *sni)
	}
	ipsFragment := "# IPs"
	if len(ips) > 0 {
		for _, ip := range ips {
			ipsFragment += fmt.Sprintf("\n  ips { ip = %q }", ip)
		}
	}
	membershipsFragment := "# Memberships"
	if len(memberships) > 0 {
		for _, membership := range memberships {
			membershipsFragment += fmt.Sprintf("\n  memberships {\n identifier = %q\n  kind = %q\n}", membership.Identifier, membership.Kind)
		}
	}

	return acctest.LoadTestCase("generatecloudflareaddressmapconfig.tf", rnd, accountId, enabled, descFragment, sniFragment, ipsFragment, membershipsFragment)
}
