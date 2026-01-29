package leaked_credential_check_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/leaked_credential_checks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func init() {
	resource.AddTestSweepers("cloudflare_leaked_credential_check", &resource.Sweeper{
		Name: "cloudflare_leaked_credential_check",
		F:    testSweepCloudflareLeakedCredentialCheck,
	})
}

func TestAccCloudflareLeakedCredentialsCheck_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create + Read
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			// Step 2: Update + Read
			{
				Config: testAccCloudflareLeakedCredentialsCheckDisabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestAccCloudflareLeakedCredentialsCheck_StateConsistency(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestAccCloudflareLeakedCredentialsCheck_Import(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the resource
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			// Step 2: Import the resource using zone_id
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        zoneID,
				ImportStateVerifyIdentifierAttribute: consts.ZoneIDSchemaKey,
			},
		},
	})
}

// testSweepCloudflareLeakedCredentialCheck resets the leaked credential check setting to disabled.
//
// This sweeper:
// - Resets the zone-level leaked credential check setting to disabled (default state)
// - Note: This is a zone-level singleton setting, not a traditional resource with instances
// - Tests may enable this setting, so the sweeper ensures it's reset to default after test runs
//
// Run with: go test ./internal/services/leaked_credential_check/ -v -sweep=all
//
// Requires:
// - CLOUDFLARE_ZONE_ID (zone ID to reset the setting for)
// - CLOUDFLARE_EMAIL + CLOUDFLARE_API_KEY or CLOUDFLARE_API_TOKEN
func testSweepCloudflareLeakedCredentialCheck(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping leaked credential check sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Checking leaked credential check setting for zone: %s", zoneID))

	// Get current setting
	setting, err := client.LeakedCredentialChecks.Get(ctx, leaked_credential_checks.LeakedCredentialCheckGetParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to get leaked credential check setting: %s", err))
		// Don't fail the sweep - setting may not exist or zone may not support it
		return nil
	}

	// If already disabled, nothing to do
	if !setting.Enabled {
		tflog.Info(ctx, fmt.Sprintf("Leaked credential check already disabled for zone: %s", zoneID))
		return nil
	}

	// Reset to disabled state
	tflog.Info(ctx, fmt.Sprintf("Resetting leaked credential check to disabled for zone: %s", zoneID))
	_, err = client.LeakedCredentialChecks.New(ctx, leaked_credential_checks.LeakedCredentialCheckNewParams{
		ZoneID:  cloudflare.F(zoneID),
		Enabled: cloudflare.F(false),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to reset leaked credential check: %s", err))
		// Continue anyway - don't fail the sweep
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Successfully reset leaked credential check to disabled for zone: %s", zoneID))
	return nil
}

// Helper functions to load test case configurations
func testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, name string) string {
	return acctest.LoadTestCase("enabled.tf", zoneID, name)
}

func testAccCloudflareLeakedCredentialsCheckDisabled(zoneID, name string) string {
	return acctest.LoadTestCase("disabled.tf", zoneID, name)
}
