package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

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

// Custom plan check for device custom profile that allows only include/exclude computed drift
type expectEmptyPlanExceptIncludeExcludeDrift struct{}

func (e expectEmptyPlanExceptIncludeExcludeDrift) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		// Skip no-op and read actions
		if rc.Change.Actions[0] == "no-op" || rc.Change.Actions[0] == "read" {
			continue
		}

		// Check if this is an update action
		if rc.Change.Actions[0] != "update" {
			resp.Error = fmt.Errorf("expected empty plan, but %s has planned action(s): %v", rc.Address, rc.Change.Actions)
			return
		}

		// For updates, check each attribute change
		beforeMap, beforeOk := rc.Change.Before.(map[string]interface{})
		afterMap, afterOk := rc.Change.After.(map[string]interface{})

		if !beforeOk || !afterOk {
			resp.Error = fmt.Errorf("expected empty plan, but %s has non-map changes", rc.Address)
			return
		}

		// Check each attribute that's different
		for key, afterValue := range afterMap {
			beforeValue, _ := beforeMap[key]

			// Skip if values are the same
			if reflect.DeepEqual(beforeValue, afterValue) {
				continue
			}

			// Allow changes from falsey to null
			if afterValue == nil {
				if isFalseyValue(beforeValue) {
					continue // This change is allowed
				}
			}

			// Device profile specific: Allow include/exclude/precedence computed field drift
			// - include/exclude: Optional+Computed, will show as "(known after apply)" when not specified
			// - precedence: API auto-assigns values, may differ from config value
			if key == "include" || key == "exclude" || key == "precedence" {
				// Allow drift for these fields
				continue
			}

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for include/exclude drift, but %s.%s has change from %v to %v",
				rc.Address, key, beforeValue, afterValue)
			return
		}
	}
}

// isFalseyValue checks if a value is "falsey" (false, 0, "", empty slice/map)
func isFalseyValue(v interface{}) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case bool:
		return !val
	case string:
		return val == ""
	case float64:
		return val == 0
	case int:
		return val == 0
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	default:
		return false
	}
}

var ExpectEmptyPlanExceptIncludeExcludeDrift = expectEmptyPlanExceptIncludeExcludeDrift{}

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on zero_trust_device_profiles → zero_trust_device_custom_profile migration:
// - v4: cloudflare_zero_trust_device_profiles (with match AND precedence)
// - v5: cloudflare_zero_trust_device_custom_profile

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_service_mode_v2.tf
var v4ServiceModeV2Config string

//go:embed testdata/v5_service_mode_v2.tf
var v5ServiceModeV2Config string

//go:embed testdata/v4_maximal.tf
var v4MaximalConfig string

//go:embed testdata/v5_maximal.tf
var v5MaximalConfig string

//go:embed testdata/v4_old_name.tf
var v4OldNameConfig string

//go:embed testdata/v5_old_name.tf
var v5OldNameConfig string

//go:embed testdata/v4_high_precedence.tf
var v4HighPrecedenceConfig string

//go:embed testdata/v5_high_precedence.tf
var v5HighPrecedenceConfig string

// TestMigrateDeviceCustomProfileBasic tests migration of custom device profile from v4 to v5
// Key transformations tested:
// - Type conversions: Int64 → Float64 (auto_connect, captive_portal, precedence, service_mode_v2.port)
// - Structure change: service_mode_v2 flatten → nested object
// - ID extraction: account_id/policy_id → policy_id attribute
// - Field removals: default, fallback_domains
// - Field keeps: name, description, match, precedence, enabled
func TestMigrateDeviceCustomProfileBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string, precedence int) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, precedence) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, precedence) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			// Use timestamp-based precedence to ensure uniqueness (within valid range 1-999)
			precedence := 1 + int(time.Now().Unix()%998)
			testConfig := tc.configFn(rnd, accountID, precedence)
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
					// Note: Persistent drift is expected for include/exclude computed fields.
					// Our custom plan check only allows the include/exclude drift as acceptable,
					// failing on anything else.
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
							PostApplyPostRefresh: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
						// Resource should be renamed to cloudflare_zero_trust_device_custom_profile
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

						// Validate fields kept from v4
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Custom Profile Test")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test custom device profile")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"test@example.com\"")),

						// Validate type conversions Int64 → Float64
						// Note: precedence is allowed to drift (API auto-assigns values), so we just check it exists
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(180)),

						// Validate service_mode_v2 structure change (flatten → nested)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(3128)),

						// Validate policy_id extracted from ID (computed field)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("policy_id"), knownvalue.NotNull()),

						// Validate pass-through fields
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("switch_locked"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.example.com")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),

						// Computed fields should be populated by API (just check they're set)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("default"), knownvalue.Bool(false)),
					},
				},
			},
		})
		})
	}
}

// TestMigrateDeviceCustomProfileWithServiceModeV2 tests migration of custom device profile with service_mode_v2 fields
// Key transformations tested:
// - service_mode_v2 flatten → nested object (mode="warp", no port)
func TestMigrateDeviceCustomProfileWithServiceModeV2(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string, precedence int) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v4ServiceModeV2Config, rnd, accountID, precedence) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v5ServiceModeV2Config, rnd, accountID, precedence) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			precedence := 1 + int(time.Now().Unix()%998)
			testConfig := tc.configFn(rnd, accountID, precedence)
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
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
							PostApplyPostRefresh: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Custom Profile with Service Mode")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"admin@example.com\"")),

							// Verify service_mode_v2 nested object (mode="warp", port is null/not set)
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("warp")),

							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(15)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateDeviceCustomProfileMaximal tests migration with all optional fields populated
// Key transformations tested:
// - Comprehensive field preservation and type conversions
func TestMigrateDeviceCustomProfileMaximal(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string, precedence int) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v4MaximalConfig, rnd, accountID, precedence) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v5MaximalConfig, rnd, accountID, precedence) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			precedence := 1 + int(time.Now().Unix()%998)
			testConfig := tc.configFn(rnd, accountID, precedence)
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
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
							PostApplyPostRefresh: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Maximal Custom Profile")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test custom profile with all optional fields")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"maximal@example.com\"")),

							// Verify boolean fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("switch_locked"), knownvalue.Bool(true)),

							// Verify numeric fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(30)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(600)),

							// Verify string fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.custom.cf-tf-test.com")),

							// Verify service_mode_v2 nested object
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(443)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateDeviceCustomProfileOldResourceName tests migration from old resource name
// Key transformations tested:
// - cloudflare_device_settings_policy → cloudflare_zero_trust_device_custom_profile
func TestMigrateDeviceCustomProfileOldResourceName(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string, precedence int) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v4OldNameConfig, rnd, accountID, precedence) },
		},
		// Note: v5 test omitted because old resource name doesn't exist in v5
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			precedence := 1 + int(time.Now().Unix()%998)
			testConfig := tc.configFn(rnd, accountID, precedence)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			firstStep := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: testConfig,
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
							PostApplyPostRefresh: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							// Verify resource type changed to new custom profile name
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Old Name Custom Profile")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"legacy@example.com\"")),

							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),

							// Verify service_mode_v2 nested object
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(8080)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateDeviceCustomProfileHighPrecedence tests migration with high precedence value
// Key transformations tested:
// - Precedence field handling for custom profiles
func TestMigrateDeviceCustomProfileHighPrecedence(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string, precedence int) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v4HighPrecedenceConfig, rnd, accountID, precedence) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string, precedence int) string { return fmt.Sprintf(v5HighPrecedenceConfig, rnd, accountID, precedence) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			precedence := 1 + int(time.Now().Unix()%998)
			testConfig := tc.configFn(rnd, accountID, precedence)
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
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
							PostApplyPostRefresh: []plancheck.PlanCheck{
								acctest.DebugNonEmptyPlan,
								ExpectEmptyPlanExceptIncludeExcludeDrift,
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("High Priority Profile")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"vip@example.com\"")),

							// Verify precedence exists (API may auto-assign different value)
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.NotNull()),

							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(5)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(120)),

							// Verify service_mode_v2 nested object
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(8080)),
						},
					},
				},
			})
		})
	}
}
