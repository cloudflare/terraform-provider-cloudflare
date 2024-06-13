package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAddressMap(t *testing.T) {
	skipForDefaultAccount(t, "Pending permission fixes for IP delegation.")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_address_map.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	altZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
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

	return fmt.Sprintf(`
resource "cloudflare_address_map" "%[1]s" {
  account_id  = "%[2]s"
  enabled = %t
  %[4]s
  %[5]s
  %[6]s
  %[7]s
}
`, rnd, accountId, enabled, descFragment, sniFragment, ipsFragment, membershipsFragment)
}
