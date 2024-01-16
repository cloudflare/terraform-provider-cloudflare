package d1_test

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
	resource.AddTestSweepers("cloudflare_d1_database", &resource.Sweeper{
		Name: "cloudflare_d1_database",
		F: func(region string) error {
			client, err := acctest.SharedClient()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			databases, _, err := client.ListD1Databases(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListD1DatabasesParams{})
			if err != nil {
				return fmt.Errorf("failed to fetch R2 buckets: %w", err)
			}

			for _, database := range databases {
				// hardcoded D1 identifier until we can solve the cyclic import
				// issues and automatically create this resource.
				if database.UUID == "ce8b95dc-b376-4ff8-9b9e-1801ed6d745d" {
					continue
				}

				err := client.DeleteD1Database(ctx, cloudflare.AccountIdentifier(accountID), database.UUID)
				if err != nil {
					return fmt.Errorf("failed to delete D1 database %q: %w", database.Name, err)
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareD1Database_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_d1_database." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareD1DatabaseBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "version", "beta"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareD1DatabaseBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_d1_database" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
  }`, rnd, accountID)
}
