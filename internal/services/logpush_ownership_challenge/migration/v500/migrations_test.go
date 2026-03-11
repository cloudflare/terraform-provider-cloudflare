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

// Embed test configs
//
//go:embed testdata/v4_basic_zone.tf
var v4BasicZoneConfig string

//go:embed testdata/v5_basic_zone.tf
var v5BasicZoneConfig string

//go:embed testdata/v4_basic_account.tf
var v4BasicAccountConfig string

//go:embed testdata/v5_basic_account.tf
var v5BasicAccountConfig string

// TestMigrateLogpushOwnershipChallenge_V4ToV5_BasicZone tests the v4→v5 migration
// for a zone-scoped ownership challenge.
//
// This verifies:
// 1. zone_id and destination_conf are preserved (direct copies)
// 2. ownership_challenge_filename is renamed to filename in v5 state (non-null)
// 3. filename, message, valid are available as computed fields after API call
func TestMigrateLogpushOwnershipChallenge_V4ToV5_BasicZone(t *testing.T) {
	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicZoneConfig, rnd, zoneID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicZoneConfig, rnd, zoneID)
			},
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
							// Validate direct-copy fields (zone_id and destination_conf)
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("destination_conf"),
								knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges"),
							),
							// Validate filename is preserved from v4 ownership_challenge_filename (renamed, not dropped)
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("filename"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLogpushOwnershipChallenge_V4ToV5_BasicAccount tests the v4→v5 migration
// for an account-scoped ownership challenge.
//
// This verifies:
// 1. account_id and destination_conf are preserved (direct copies)
// 2. ownership_challenge_filename is renamed to filename in v5 state (non-null)
// 3. filename, message, valid are available as computed fields after API call
func TestMigrateLogpushOwnershipChallenge_V4ToV5_BasicAccount(t *testing.T) {
	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BasicAccountConfig, rnd, accountID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicAccountConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
							// Validate direct-copy fields (account_id and destination_conf)
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("destination_conf"),
								knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges"),
							),
							// Validate filename is preserved from v4 ownership_challenge_filename (renamed, not dropped)
							statecheck.ExpectKnownValue(
								"cloudflare_logpush_ownership_challenge."+rnd,
								tfjsonpath.New("filename"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}
