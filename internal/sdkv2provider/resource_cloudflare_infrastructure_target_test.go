package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareInfrastructureTargetCreateUpdate(t *testing.T) {
	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_infrastructure_target.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareInfrastructureTargetDestroy,
		Steps: []resource.TestStep{
			{
				// Create resource configuration
				Config: testAccCloudflareInfrastructureTargetCreate(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", rnd),
					resource.TestCheckResourceAttr(name, "ip.ipv4.ip_addr", "250.26.29.250"),
					resource.TestCheckNoResourceAttr(name, "ip.ipv6"),
				),
			},
			{
				// Update resource configuration
				Config: testAccCloudflareInfrastructureTargetUpdate(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "ip.ipv4.ip_addr", "250.26.29.250"),
					resource.TestCheckResourceAttr(name, "ip.ipv6.ip_addr", "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					resource.TestCheckResourceAttr(name, "ip.ipv6.virtual_network_id", "01920a8c-dc14-7bb2-b67b-14c858494a54"),
				),
			},
		},
	})
}

func testAccCloudflareInfrastructureTargetCreate(accID, hostname string) string {
	return fmt.Sprintf(`
	resource "cloudflare_zero_trust_infrastructure_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s"
	ip = {
		ipv4 = {
           ip_addr = "250.26.29.250",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        },
	}`, accID, hostname)
}

func testAccCloudflareInfrastructureTargetUpdate(accID, hostname string) string {
	return fmt.Sprintf(`
	resource "cloudflare_zero_trust_infrastructure_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s-updated"
	ip = {
		ipv4 = {
           ip_addr = "250.26.29.250",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        },
		ipv6 = {
           ip_addr = "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        }
	}`, accID, hostname)
}

func testAccCheckCloudflareInfrastructureTargetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_infrastructure_target" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		targetID := rs.Primary.ID
		client := testAccProvider.Meta().(*cloudflare.API)
		target, err := client.GetInfrastructureTarget(context.Background(), cloudflare.AccountIdentifier(accountID), targetID)

		if err == nil {
			return fmt.Errorf("infrastructure target with ID %s still exists", target.ID)
		}

	}

	return nil
}
