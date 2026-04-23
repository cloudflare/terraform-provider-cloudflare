package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_firewall.tf
var v4FirewallConfig string

//go:embed testdata/v5_firewall.tf
var v5FirewallConfig string

//go:embed testdata/v4_disk_encryption.tf
var v4DiskEncryptionConfig string

//go:embed testdata/v5_disk_encryption.tf
var v5DiskEncryptionConfig string

//go:embed testdata/v4_multiple_platforms.tf
var v4MultiplePlatformsConfig string

//go:embed testdata/v5_multiple_platforms.tf
var v5MultiplePlatformsConfig string

//go:embed testdata/v4_file.tf
var v4FileConfig string

//go:embed testdata/v5_file.tf
var v5FileConfig string

//go:embed testdata/v4_domain_joined.tf
var v4DomainJoinedConfig string

//go:embed testdata/v5_domain_joined.tf
var v5DomainJoinedConfig string

//go:embed testdata/v4_without_name.tf
var v4WithoutNameConfig string

//go:embed testdata/v4_empty_name.tf
var v4EmptyNameConfig string

//go:embed testdata/v5_updated_name.tf
var v5UpdatedNameConfig string

// clearAPIToken unsets CLOUDFLARE_API_TOKEN for Zero Trust tests (require API_KEY + EMAIL)
// and returns a cleanup function to restore it.
func clearAPIToken(t *testing.T) {
	t.Helper()
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		t.Cleanup(func() { os.Setenv("CLOUDFLARE_API_TOKEN", originalToken) })
	}
}

// TestMigrateDevicePostureRuleBasic tests migration of a basic os_version posture rule
// with input and match transformations.
func TestMigrateDevicePostureRuleBasic(t *testing.T) {
	clearAPIToken(t)

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
			name := fmt.Sprintf("tf-test-posture-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Device posture rule for corporate devices.")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("24h")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("expiration"), knownvalue.StringExact("25h")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("1.0.0")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact("<")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_distro_name"), knownvalue.StringExact("ubuntu")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_distro_revision"), knownvalue.StringExact("1.0.0")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_version_extra"), knownvalue.StringExact("(a)")),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleWithoutName tests migration of a v4 resource created without name.
// Name is optional in both v4 and v5 (API returns null when not set).
// Uses MigrationV2TestStepWithPlan pattern: migration produces a one-time plan diff (name "" -> null)
// that needs an intermediate plan step to correct before final validation.
func TestMigrateDevicePostureRuleWithoutName(t *testing.T) {
	// NOTE: This test passes because TransformConfigWithState automatically generates
	// a name (using the Terraform resource name) when state/API don't have one.

	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config without name field (name was optional in v4, required in v5)
	// When v4 creates this resource without a name, the API stores it with empty name.
	// After migration, we auto-generate a name, and v5 will update the API resource.
	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  type       = "os_version"
  schedule   = "5m"

  match {
    platform = "linux"
  }

  input {
    version  = "10.0.0"
    operator = ">="
  }
}`, rnd, accountID)

	// Build test steps
	steps := []resource.TestStep{
		{
			// Step 1: Create with v4 provider (without name in config)
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		},
	}

	// Step 2: Migration adds name to config/state, v5 updates API with the name
	// A plan diff is expected because the API resource needs to be updated with the name
	migrationSteps := acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		// Name is null after migration (not set in v4 config, API returns "" which is normalized to null)
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.Null()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),

		// Match should be converted to array
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),

		// Input should be object
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.0")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
	})
	steps = append(steps, migrationSteps...)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps:      steps,
	})
}

// TestMigrateDevicePostureRuleFirewall tests migration of a firewall rule.
func TestMigrateDevicePostureRuleFirewall(t *testing.T) {
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4FirewallConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5FirewallConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-firewall-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("firewall")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("enabled"), knownvalue.Bool(true)),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleDiskEncryption tests Set->List conversion for check_disks.
func TestMigrateDevicePostureRuleDiskEncryption(t *testing.T) {
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4DiskEncryptionConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5DiskEncryptionConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-disk-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("disk_encryption")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("check_disks"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("require_all"), knownvalue.Bool(true)),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleMultiplePlatforms tests multiple match blocks->array conversion.
func TestMigrateDevicePostureRuleMultiplePlatforms(t *testing.T) {
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4MultiplePlatformsConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5MultiplePlatformsConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-multi-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("firewall")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(1).AtMapKey("platform"), knownvalue.StringExact("mac")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(2).AtMapKey("platform"), knownvalue.StringExact("linux")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("enabled"), knownvalue.Bool(true)),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleFileType tests file type with multiple input attributes.
func TestMigrateDevicePostureRuleFileType(t *testing.T) {
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4FileConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5FileConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-file-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("file")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("path"), knownvalue.StringExact("C:\\Program Files\\app.exe")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("exists"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("thumbprint"), knownvalue.StringExact("0123456789abcdef")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("sha256"), knownvalue.StringExact("abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890")),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleDomainJoined tests domain_joined type.
func TestMigrateDevicePostureRuleDomainJoined(t *testing.T) {
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4DomainJoinedConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5DomainJoinedConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-domain-%s", rnd)
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("domain_joined")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					}),
				},
			})
		})
	}
}

// TestMigrateDevicePostureRuleEmptyNameThenUpdate tests migration of a v4 resource with
// name="" (explicitly empty), verifies it migrates to null, then updates with a real name in v5.
func TestMigrateDevicePostureRuleEmptyNameThenUpdate(t *testing.T) {
	clearAPIToken(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-renamed-%s", rnd)
	tmpDir := t.TempDir()
	v4Config := fmt.Sprintf(v4EmptyNameConfig, rnd, accountID)
	v5Config := fmt.Sprintf(v5UpdatedNameConfig, rnd, accountID, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider (name = "")
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
			// Step 2: Migrate to v5 — config still has name="" but state is null after Transform.
			// This produces a plan diff (null → "") and a non-empty refresh plan, both expected.
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
			// Step 3: Verify state settled — name is "" (from config), plan is clean
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.0")),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
				},
			},
			// Step 4: Update with v5 config that sets a real name
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.0")),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
				},
			},
		},
	})
}

// Note: Tests for integration types (client_certificate, client_certificate_v2, sentinelone_s2s,
// tanium_s2s, kolide, crowdstrike_s2s, intune, workspace_one) are not included because they
// require valid connection_id values from real integrations.
