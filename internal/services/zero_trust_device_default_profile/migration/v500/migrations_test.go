package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"reflect"
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

// Custom plan check for device default profile that allows only include/exclude computed drift
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

			// Device profile specific: Allow include/exclude computed field drift
			// The include/exclude fields are Optional+Computed and will show as "(known after apply)" on refresh
			// when not specified in config
			if key == "include" || key == "exclude" {
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
// Based on zero_trust_device_profiles → zero_trust_device_default_profile migration:
// - v4: cloudflare_zero_trust_device_profiles (with default=true OR no match+precedence)
// - v5: cloudflare_zero_trust_device_default_profile

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

// TestMigrateDeviceDefaultProfileBasic tests migration of default device profile from v4 to v5
// Key transformations tested:
// - Type conversions: Int64 → Float64 (auto_connect, captive_portal, service_mode_v2.port)
// - Structure change: service_mode_v2 flatten → nested object
// - Field additions: register_interface_ip_with_dns, sccm_vpn_boundary_support
// - Field removals: default, enabled, name, description, match, precedence (not applicable for default profile)
func TestMigrateDeviceDefaultProfileBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
						// Resource should be renamed to cloudflare_zero_trust_device_default_profile
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

						// Validate type conversions Int64 → Float64
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(180)),

						// Validate service_mode_v2 structure change (flatten → nested)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(8080)),

						// Validate field additions with v5 defaults
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("register_interface_ip_with_dns"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("sccm_vpn_boundary_support"), knownvalue.Bool(false)),

						// Validate pass-through fields
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("switch_locked"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.example.com")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),

						// Computed fields should be populated by API (just check they're set)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("default"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					},
				},
			},
		})
		})
	}
}

// TestMigrateDeviceDefaultProfileWithServiceModeV2 tests migration with service_mode_v2 transformation
// Tests: flat fields (service_mode_v2_mode, service_mode_v2_port) → nested object
func TestMigrateDeviceDefaultProfileWithServiceModeV2(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4ServiceModeV2Config, rnd, accountID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5ServiceModeV2Config, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							
							// Verify service_mode_v2 nested object created
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(8080)),
							
							// Verify other fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(15)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateDeviceDefaultProfileMaximal tests migration with all optional fields
// Tests: comprehensive field preservation, type conversions
func TestMigrateDeviceDefaultProfileMaximal(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4MaximalConfig, rnd, accountID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5MaximalConfig, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							
							// Verify boolean fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("switch_locked"), knownvalue.Bool(true)),
							
							// Verify numeric fields (converted to Float64)
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(30)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(600)),
							
							// Verify string fields
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.cf-tf-test.com")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),
							
							// Verify service_mode_v2 nested object
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(443)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateDeviceDefaultProfileOldResourceName tests migration from old resource name
// Tests: cloudflare_device_settings_policy (old name) → cloudflare_zero_trust_device_default_profile
func TestMigrateDeviceDefaultProfileOldResourceName(t *testing.T) {
	// Note: from_v5 test case not applicable - old resource name only existed in v4
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4OldNameConfig, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Use external v4 provider (will create version=0 state)
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
					// Step 2: Run migration and verify state
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
							// Verify resource type changed to new name
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
							statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
						},
					},
				},
			})
		})
	}
}
