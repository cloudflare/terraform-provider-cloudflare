package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_output_options.tf
var v4OutputOptionsConfig string

//go:embed testdata/v5_output_options.tf
var v5OutputOptionsConfig string

//go:embed testdata/v4_instant_logs.tf
var v4InstantLogsConfig string

//go:embed testdata/v5_instant_logs.tf
var v5InstantLogsConfig string

// TestMigrateLogpushJobBasic tests basic logpush_job migration with output_options.
// This test verifies:
// 1. output_options block { ... } → output_options = { ... } (block to attribute syntax)
// 2. cve20214428 field is renamed to cve_2021_44228
// 3. Empty string fields (filter, logpull_options, name) converted to null
// 4. v4 schema defaults preserved in output_options
func TestMigrateLogpushJobBasic(t *testing.T) {
	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, name) },
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn:       func(rnd, accountID, name string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-logpush-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				step1 = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify identity fields
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("dataset"),
								knownvalue.StringExact("audit_logs"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLogpushJobOutputOptions tests migration of logpush_job with all output_options fields.
// This test verifies:
// 1. All output_options nested fields are preserved
// 2. cve20214428 field is renamed to cve_2021_44228
// 3. field_names list is preserved
// 4. sample_rate type conversion works correctly
func TestMigrateLogpushJobOutputOptions(t *testing.T) {
	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4OutputOptionsConfig, rnd, accountID, name) },
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn:       func(rnd, accountID, name string) string { return fmt.Sprintf(v5OutputOptionsConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-logpush-options-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				step1 = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify identity fields
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("dataset"),
								knownvalue.StringExact("audit_logs"),
							),
							// Verify ALL output_options fields are preserved
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("batch_prefix"),
								knownvalue.StringExact("batch-start-"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("batch_suffix"),
								knownvalue.StringExact("-batch-end"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("cve_2021_44228"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("field_delimiter"),
								knownvalue.StringExact("|"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("output_type"),
								knownvalue.StringExact("ndjson"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("record_delimiter"),
								knownvalue.StringExact("\\n"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("record_prefix"),
								knownvalue.StringExact("["),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("record_suffix"),
								knownvalue.StringExact("]\\n"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("record_template"),
								knownvalue.StringExact("{{.ClientIP}}"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("sample_rate"),
								knownvalue.Float64Exact(0.5),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("timestamp_format"),
								knownvalue.StringExact("rfc3339"),
							),
							// Verify field_names list is preserved (all 3 elements)
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("field_names").AtSliceIndex(0),
								knownvalue.StringExact("ClientIP"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("field_names").AtSliceIndex(1),
								knownvalue.StringExact("EdgeStartTimestamp"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("output_options").AtMapKey("field_names").AtSliceIndex(2),
								knownvalue.StringExact("ClientRequestMethod"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLogpushJobInstantLogs tests migration of logpush_job with kind="instant-logs".
// This test verifies:
// 1. kind="instant-logs" is converted to kind="" (instant-logs no longer valid in v5)
// 2. State transformation properly handles this field removal
func TestMigrateLogpushJobInstantLogs(t *testing.T) {
	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4InstantLogsConfig, rnd, zoneID) },
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn:       func(rnd, zoneID string) string { return fmt.Sprintf(v5InstantLogsConfig, rnd, zoneID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				step1 = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify identity fields
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("dataset"),
								knownvalue.StringExact("http_requests"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(true),
							),
							// Verify kind is converted to empty string
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_job."+rnd,
								tfjsonpath.New("kind"),
								knownvalue.StringExact(""),
							),
						},
					),
				},
			})
		})
	}
}
