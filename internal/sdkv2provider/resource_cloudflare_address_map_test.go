package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAddressMap(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_address_map.%s", rnd)

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
					cloudflare.StringPtr("Terraform provider test"),
					cloudflare.StringPtr("*.ipam.rocks"),
					true,
					[]string{"1.0.0.2", "1.0.0.3"},
					[]cloudflare.AddressMapMembershipContainer{
						{Identifier: "a4de22ad6ae5f5ab736ec887dc14660d", Kind: cloudflare.AddressMapMembershipZone},
						{Identifier: "c6737c4cd61718ad3c3953680e638959", Kind: cloudflare.AddressMapMembershipZone},
					}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "description", "Terraform provider test"),
					resource.TestCheckResourceAttr(name, "default_sni", "*.ipam.rocks"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "ips.0.ip", "1.0.0.2"),
					resource.TestCheckResourceAttr(name, "ips.1.ip", "1.0.0.3"),
					resource.TestCheckResourceAttr(name, "memberships.0.identifier", "a4de22ad6ae5f5ab736ec887dc14660d"),
					resource.TestCheckResourceAttr(name, "memberships.0.kind", "zone"),
					resource.TestCheckResourceAttr(name, "memberships.1.identifier", "c6737c4cd61718ad3c3953680e638959"),
					resource.TestCheckResourceAttr(name, "memberships.1.kind", "zone"),
				),
			},
			{
				Config: generateCloudflareAddressMapConfig(rnd, accountID, cloudflare.StringPtr(""), cloudflare.StringPtr("*.ipam.rocks"), false, nil, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "default_sni", "*.ipam.rocks"),
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
