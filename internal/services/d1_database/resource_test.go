package d1_database_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/d1"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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

func testAccCheckCloudflareD1DatabaseDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_d1_database" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.D1.Database.Get(
			context.Background(),
			rs.Primary.ID,
			d1.DatabaseGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("d1 database %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// TestAccCloudflareD1Database_Basic verifies basic create and read of a D1
// database with only the required attributes (account_id, name).
//
// Skip: read_replication is Optional-only in the current schema, but the API
// always returns it in responses, causing refresh drift. Blocked on BUGS-2016
// codegen fix (PR #7183).
func TestAccCloudflareD1Database_Basic(t *testing.T) {
	t.Skip("BUGS-2016: read_replication drift -- waiting for codegen fix on next (PR #7183)")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_d1_database." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareD1DatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareD1DatabaseBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "version", "production"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("uuid"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("version"), knownvalue.StringExact("production")),
				},
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

// TestAccCloudflareD1Database_ReadReplication verifies creating a D1 database
// with read_replication set to "disabled", then updating it to "auto", and
// back to "disabled".
func TestAccCloudflareD1Database_ReadReplication(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_d1_database." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareD1DatabaseDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with read_replication disabled
			{
				Config: testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "read_replication.mode", "disabled"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("read_replication").AtMapKey("mode"), knownvalue.StringExact("disabled")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"file_size"},
			},
			// Step 2: Update to auto
			{
				Config: testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, "auto"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "read_replication.mode", "auto"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("read_replication").AtMapKey("mode"), knownvalue.StringExact("auto")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"file_size"},
			},
			// Step 3: Update back to disabled
			{
				Config: testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, "disabled"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "read_replication.mode", "disabled"),
				),
			},
		},
	})
}

// TestAccCloudflareD1Database_NoDrift verifies that re-applying the same
// config produces an empty plan (no perpetual drift).
//
// Skip: read_replication is Optional-only in the current schema, but the API
// always returns it in responses, causing refresh drift. Blocked on BUGS-2016
// codegen fix (PR #7183).
func TestAccCloudflareD1Database_NoDrift(t *testing.T) {
	t.Skip("BUGS-2016: read_replication drift -- waiting for codegen fix on next (PR #7183)")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_d1_database." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareD1DatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareD1DatabaseBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
				),
			},
			// Re-apply the same config and verify no changes are planned
			{
				Config:             testAccCheckCloudflareD1DatabaseBasic(rnd, accountID),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccCloudflareD1Database_ReadReplicationNoDrift verifies that a database
// created with read_replication does not show perpetual drift. This is the
// regression test for BUGS-2016.
//
// Skip: The codegen fix (x-stainless-terraform-configurability: computed_optional
// on read-replication-details-for-request) is on origin/generated but has not
// synced to next yet. Remove this skip once codegen sync PR #7183 merges and
// read_replication is marked Computed+Optional in schema.go.
func TestAccCloudflareD1Database_ReadReplicationNoDrift(t *testing.T) {
	t.Skip("BUGS-2016: waiting for codegen fix to land on next (PR #7183)")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_d1_database." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareD1DatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "read_replication.mode", "disabled"),
				),
			},
			// Re-apply the same config and verify no drift
			{
				Config:             testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, "disabled"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// Config helpers

func testAccCheckCloudflareD1DatabaseBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("d1databasebasic.tf", rnd, accountID)
}

func testAccCheckCloudflareD1DatabaseReadReplication(rnd, accountID, mode string) string {
	return acctest.LoadTestCase("d1databasereadreplication.tf", rnd, accountID, mode)
}
