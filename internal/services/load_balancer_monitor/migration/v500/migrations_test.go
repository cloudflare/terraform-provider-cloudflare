package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_headers.tf
var v4WithHeadersConfig string

//go:embed testdata/v5_with_headers.tf
var v5WithHeadersConfig string

//go:embed testdata/v4_full.tf
var v4FullConfig string

//go:embed testdata/v5_full.tf
var v5FullConfig string

// TestMigrateLoadBalancerMonitor_V4ToV5_Basic tests basic field migrations with minimal config.
// This validates default value additions for fields that were optional without defaults in v4.
func TestMigrateLoadBalancerMonitor_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Critical identifiers
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("http"),
							),

							// Verify default values were added by migration
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("description"),
								knownvalue.StringExact(""),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("expected_body"),
								knownvalue.StringExact(""),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("expected_codes"),
								knownvalue.StringExact("200"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("probe_zone"),
								knownvalue.StringExact(""),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("allow_insecure"),
								knownvalue.Bool(false),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("follow_redirects"),
								knownvalue.Bool(false),
							),

							// Verify computed fields exist
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("created_on"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancerMonitor_V4ToV5_WithHeaders tests the complex header transformation.
// This validates the Set (nested) → Map transformation, which is the primary challenge.
func TestMigrateLoadBalancerMonitor_V4ToV5_WithHeaders(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, hostname string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, hostname string) string {
				return fmt.Sprintf(v4WithHeadersConfig, rnd, accountID, hostname)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, hostname string) string {
				return fmt.Sprintf(v5WithHeadersConfig, rnd, accountID, hostname)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			hostname := "health-check." + zoneName
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, hostname)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify basic fields
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("https"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("method"),
								knownvalue.StringExact("GET"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("path"),
								knownvalue.StringExact("/health"),
							),

							// Verify header transformation: Set → Map
							// v4: header { header = "Host" values = ["hostname"] }
							// v5: header = { "Host" = ["hostname"] }
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("header").AtMapKey("Host").AtSliceIndex(0),
								knownvalue.StringExact(hostname),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancerMonitor_V4ToV5_Full tests migration with all optional fields populated.
// This validates comprehensive field coverage and ensures no data loss.
func TestMigrateLoadBalancerMonitor_V4ToV5_Full(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, probeZone, hostname string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, probeZone, hostname string) string {
				return fmt.Sprintf(v4FullConfig, rnd, accountID, probeZone, hostname)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, probeZone, hostname string) string {
				return fmt.Sprintf(v5FullConfig, rnd, accountID, probeZone, hostname)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			hostname := "api." + zoneName
			probeZone := zoneName
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, probeZone, hostname)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify all configured values are preserved
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("https"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("method"),
								knownvalue.StringExact("GET"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("path"),
								knownvalue.StringExact("/api/health"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("port"),
								knownvalue.Int64Exact(8443),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.Int64Exact(60),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("retries"),
								knownvalue.Int64Exact(3),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("timeout"),
								knownvalue.Int64Exact(10),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("allow_insecure"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("follow_redirects"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("expected_codes"),
								knownvalue.StringExact("2xx"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("expected_body"),
								knownvalue.StringExact("OK"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("description"),
								knownvalue.StringExact("Production API health check"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("consecutive_up"),
								knownvalue.Int64Exact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("consecutive_down"),
								knownvalue.Int64Exact(3),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("probe_zone"),
								knownvalue.StringExact(probeZone),
							),

							// Verify header transformation
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer_monitor."+rnd,
								tfjsonpath.New("header").AtMapKey("Host").AtSliceIndex(0),
								knownvalue.StringExact(hostname),
							),
						},
					),
				},
			})
		})
	}
}
