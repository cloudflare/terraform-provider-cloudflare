package d1_database_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_d1_database", &resource.Sweeper{
		Name: "cloudflare_d1_database",
		F: func(region string) error {
			ctx := context.Background()
			client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", err))
				return fmt.Errorf("error establishing client: %w", err)
			}

			if accountID == "" {
				tflog.Info(ctx, "Skipping D1 databases sweep: CLOUDFLARE_ACCOUNT_ID not set")
				return nil
			}

			databases, _, err := client.ListD1Databases(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListD1DatabasesParams{})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch D1 databases: %s", err))
				return fmt.Errorf("failed to fetch D1 databases: %w", err)
			}

			if len(databases) == 0 {
				tflog.Info(ctx, "No D1 databases to sweep")
				return nil
			}

			for _, database := range databases {
				// hardcoded D1 identifier until we can solve the cyclic import
				// issues and automatically create this resource.
				if database.UUID == "ce8b95dc-b376-4ff8-9b9e-1801ed6d745d" {
					continue
				}

				if !utils.ShouldSweepResource(database.Name) {
					continue
				}

				tflog.Info(ctx, fmt.Sprintf("Deleting D1 database: %s (%s) (account: %s)", database.Name, database.UUID, accountID))
				err := client.DeleteD1Database(ctx, cfv1.AccountIdentifier(accountID), database.UUID)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete D1 database %s: %s", database.Name, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted D1 database: %s", database.UUID))
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
					resource.TestCheckResourceAttr(resourceName, "version", "production"),
				),
			},
			// {
			// 	ResourceName:        resourceName,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func TestAccCloudflareD1Database_ReadReplicationRejected(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID),
				ExpectError: regexp.MustCompile("Invalid Attribute Configuration"),
			},
		},
	})
}

func testAccCheckCloudflareD1DatabaseBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("d1databasebasic.tf", rnd, accountID)
}

func testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID string) string {
	return acctest.LoadTestCase("d1databasebasicreadreplication.tf", rnd, accountID)
}

func TestAccCloudflareD1Database_ReadReplicationUpdate(t *testing.T) {
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
				),
			},
			{
				Config: testAccCheckCloudflareD1DatabaseWithReadReplication(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "read_replication.mode", "auto"),
				),
			},
		},
	})
}

func testAccCheckCloudflareD1DatabaseWithReadReplication(rnd, accountID string) string {
	return acctest.LoadTestCase("d1databasewithreadreplication.tf", rnd, accountID)
}
