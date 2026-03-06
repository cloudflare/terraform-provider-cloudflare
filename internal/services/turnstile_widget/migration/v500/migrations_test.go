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

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_full.tf
var v4FullConfig string

//go:embed testdata/v5_full.tf
var v5FullConfig string

//go:embed testdata/v4_sorted_domains.tf
var v4SortedDomainsConfig string

//go:embed testdata/v5_sorted_domains.tf
var v5SortedDomainsConfig string

// TestMigrateTurnstileWidget_V4ToV5_Basic tests migration of a basic turnstile_widget
// with minimal configuration. This verifies:
// 1. domains field transforms from Set to List
// 2. All required fields are preserved (account_id, name, mode)
// 3. Migration works from both v4 and v5
func TestMigrateTurnstileWidget_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (has GetSchemaVersion, will create version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
				firstStep = resource.TestStep{
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
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("mode"),
								knownvalue.StringExact("managed"),
							),
							// Verify domains transformed to List
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains").AtSliceIndex(0),
								knownvalue.StringExact("example.com"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateTurnstileWidget_V4ToV5_AllOptionalFields tests migration with all optional
// fields configured. This verifies:
// 1. All optional fields are preserved (region, bot_fight_mode, offlabel)
// 2. Multiple domains are handled correctly
// 3. Different mode values work (invisible)
func TestMigrateTurnstileWidget_V4ToV5_AllOptionalFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4FullConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5FullConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-full-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("mode"),
								knownvalue.StringExact("invisible"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("region"),
								knownvalue.StringExact("world"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("bot_fight_mode"),
								knownvalue.Bool(false),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("offlabel"),
								knownvalue.Bool(false),
							),
							// Verify multiple domains
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains"),
								knownvalue.ListSizeExact(2),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateTurnstileWidget_V4ToV5_DomainsSorted tests that domains are
// alphabetically sorted after migration. This is CRITICAL because:
// 1. v4 uses SetAttribute (unordered)
// 2. v5 uses ListAttribute (ordered)
// 3. The Cloudflare API returns domains alphabetically
// 4. Without sorting, every terraform apply would show drift
func TestMigrateTurnstileWidget_V4ToV5_DomainsSorted(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4SortedDomainsConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5SortedDomainsConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-sorted-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify domains are in alphabetical order
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains"),
								knownvalue.ListSizeExact(3),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains").AtSliceIndex(0),
								knownvalue.StringExact("aaa.example.com"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains").AtSliceIndex(1),
								knownvalue.StringExact("mmm.example.com"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_turnstile_widget."+rnd,
								tfjsonpath.New("domains").AtSliceIndex(2),
								knownvalue.StringExact("zzz.example.com"),
							),
						},
					),
				},
			})
		})
	}
}
