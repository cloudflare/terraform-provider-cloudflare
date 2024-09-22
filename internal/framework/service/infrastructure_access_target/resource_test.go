package infrastructure_access_target_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_infrastructure_access_target", &resource.Sweeper{
		Name: "cloudflare_infrastructure_access_target",
		F: func(region string) error {
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			// Retrieve all targets created under the current test account
			targets, _, err := client.ListInfrastructureAccessTargets(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.TargetListParams{})
			if err != nil {
				return fmt.Errorf("failed to fetch rulesets: %w", err)
			}

			// Delete each target
			for _, target := range targets {
				err := client.DeleteInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountID), target.ID)
				if err != nil {
					return fmt.Errorf("failed to delete ruleset %q: %w", target.ID, err)
				}
			}

			return nil
		},
	})
}

func TestAccCfInfrastructureAccessTargetCreateUpdate(t *testing.T) {
	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_infrastructure_access_target.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create resource configuration
				Config: testAccCfInfrastructureAccessTargetCreate(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname", rnd),
					resource.TestCheckResourceAttr(resourceName, "ip.ipv4.ip_addr", "250.26.29.250"),
					resource.TestCheckNoResourceAttr(resourceName, "ip.ipv6"),
				),
			},
			{
				// Update resource configuration
				Config: testAccCfInfrastructureAccessTargetUpdate(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, "ip.ipv4.ip_addr", "250.26.29.250"),
					resource.TestCheckResourceAttr(resourceName, "ip.ipv6.ip_addr", "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					resource.TestCheckResourceAttr(resourceName, "ip.ipv6.virtual_network_id", "01920a8c-dc14-7bb2-b67b-14c858494a54"),
				),
			},
		},
	})
}

func testAccCfInfrastructureAccessTargetCreate(accID, hostname string) string {
	return fmt.Sprintf(`
	resource "cloudflare_infrastructure_access_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s"
	ip = {
		ipv4 = {
           ip_addr = "250.26.29.250",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        },
	}`, accID, hostname)
}

func testAccCfInfrastructureAccessTargetUpdate(accID, hostname string) string {
	return fmt.Sprintf(`
	resource "cloudflare_infrastructure_access_target" "%[2]s" {
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
