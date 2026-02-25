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

// Embed test config files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_optional.tf
var v4WithOptionalConfig string

//go:embed testdata/v5_with_optional.tf
var v5WithOptionalConfig string

//go:embed testdata/v4_removed_fields.tf
var v4RemovedFieldsConfig string

//go:embed testdata/v4_complete.tf
var v4CompleteConfig string

// TestMigrateZone_Basic tests migration of the core field transformations:
//   - zone → name (field rename)
//   - account_id (flat string) → account.id (nested object)
func TestMigrateZone_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneName, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneName, accountID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneName, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneName, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneName, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			resourceName := "cloudflare_zone." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneName, accountID)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// zone → name (field rename)
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
						// account_id → account.id (flat string → nested object)
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
						// Sanity: resource was created
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
						// Computed list: name_servers exercises types.List → customfield.List conversion
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name_servers"), knownvalue.NotNull()),
						// Computed nested objects: API repopulates after migration nullifies
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateZone_WithOptionalFields tests migration of pass-through fields:
//   - paused (Bool) — unchanged across versions
//   - type (String) — unchanged across versions
func TestMigrateZone_WithOptionalFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneName, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneName, accountID string) string {
				return fmt.Sprintf(v4WithOptionalConfig, rnd, zoneName, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneName, accountID string) string {
				return fmt.Sprintf(v5WithOptionalConfig, rnd, zoneName, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			resourceName := "cloudflare_zone." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneName, accountID)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
						// pass-through: paused preserved
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
						// pass-through: type preserved
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
					}),
				},
			})
		})
	}
}

// TestMigrateZone_RemovedFields tests that v4-only fields (jump_start, plan) are dropped cleanly
// during migration without causing errors. This is a v4-only test since jump_start does not
// exist in the v5 schema and plan changes from a configurable string to a computed nested object.
func TestMigrateZone_RemovedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	v4Config := fmt.Sprintf(v4RemovedFieldsConfig, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				// Core transformations still work with removed fields present in v4 state
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// plan: v4 configurable string dropped, v5 computed nested object repopulated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZone_Complete tests migration with all v4 configurable attributes including
// fields that are removed (jump_start, plan) and pass-through fields with non-default values
// (paused=true). This is v4-only since plan and jump_start don't exist in v5, and
// paused=true causes a non-empty refresh plan with the v5 provider (API creates as false,
// requires a separate update).
func TestMigrateZone_Complete(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	v4Config := fmt.Sprintf(v4CompleteConfig, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				// Core transformations
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				// pass-through: paused=true (non-default value)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(true)),
				// pass-through: type preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
				// Removed fields don't break migration; computed fields repopulated
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name_servers"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan"), knownvalue.NotNull()),
			}),
		},
	})
}
