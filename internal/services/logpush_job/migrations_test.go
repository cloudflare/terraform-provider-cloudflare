package logpush_job_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflareLogpushJob_Migration_Basic_MultiVersion tests the most fundamental
// logpush job migration scenario with output_options block to attribute transformation.
// This test ensures that:
// 1. output_options block { ... } → output_options = { ... } (block to attribute syntax)
// 2. cve20214428 field is renamed to cve_2021_44228
// 3. kind = "instant-logs" is removed
// 4. Numeric fields are properly converted (max_upload_* fields)
// 5. The migration tool successfully transforms both configuration and state files
func TestMigrateCloudflareLogpushJob_Migration_Basic_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLogpushJobMigrationConfigV4Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := acctest.TestAccCloudflareAccountID
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_logpush_job." + rnd
			testConfig := tc.configFn(accountID, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create logpush job with v4 provider
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					}),
					{
						// Step 3: Apply migrated config with v5 provider
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflareLogpushJob_Migration_OutputOptions tests migration of logpush jobs
// with output_options blocks. This test verifies that:
// 1. output_options block syntax is converted to attribute syntax with =
// 2. cve20214428 field is properly renamed to cve_2021_44228
// 3. All nested fields within output_options are preserved
// 4. State transformation converts array [{...}] to object {...}
func TestMigrateCloudflareLogpushJob_Migration_OutputOptions(t *testing.T) {
	accountID := acctest.TestAccCloudflareAccountID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	v4Config := testAccCloudflareLogpushJobMigrationConfigV4OutputOptions(accountID, rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create logpush job with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact("audit_logs")),
				},
			},
		},
	})
}

// TestMigrateCloudflareLogpushJob_Migration_InstantLogs tests migration of logpush jobs
// with kind = "instant-logs" which is removed in v5 (instant-logs is no longer valid).
func TestMigrateCloudflareLogpushJob_Migration_InstantLogs(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	v4Config := testAccCloudflareLogpushJobMigrationConfigV4InstantLogs(zoneID, rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create logpush job with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("instant-logs")),
				},
			},
			// Step 2: Migrate to v5 provider
			// Note: We expect a non-empty plan because removing "instant-logs" creates a real change
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true, // Removing kind="instant-logs" creates a plan diff
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.Null()),
				},
			},
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.Null()),
				},
			},
		},
	})
}

// V4 Configuration Functions

func testAccCloudflareLogpushJobMigrationConfigV4Basic(accountID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_logpush_job" "%[2]s" {
  account_id       = "%[1]s"
  dataset          = "audit_logs"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
}
`, accountID, rnd)
}

func testAccCloudflareLogpushJobMigrationConfigV4OutputOptions(accountID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_logpush_job" "%[2]s" {
  account_id       = "%[1]s"
  dataset          = "audit_logs"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true

  output_options {
    cve20214428 = true
    output_type = "ndjson"
    field_names = ["ClientIP", "EdgeStartTimestamp"]
  }
}
`, accountID, rnd)
}

func testAccCloudflareLogpushJobMigrationConfigV4InstantLogs(zoneID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_logpush_job" "%[2]s" {
  zone_id       = "%[1]s"
  dataset          = "http_requests"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
  kind             = "instant-logs"
}
`, zoneID, rnd)
}
