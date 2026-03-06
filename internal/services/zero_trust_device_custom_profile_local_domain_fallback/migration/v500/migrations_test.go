package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"hash/fnv"
	"os"
	"reflect"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Custom plan check that allows precedence drift for device custom profiles
type expectEmptyPlanExceptPrecedenceDrift struct{}

func (e expectEmptyPlanExceptPrecedenceDrift) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
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

			// Device custom profile specific: Allow precedence computed field drift
			// - precedence: API auto-assigns values, may differ from config value
			if key == "precedence" {
				// Allow drift for precedence
				continue
			}

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for precedence drift, but %s.%s has change from %v to %v",
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

var ExpectEmptyPlanExceptPrecedenceDrift = expectEmptyPlanExceptPrecedenceDrift{}

// generateUniquePrecedence generates a unique precedence value based on the random resource name
// to avoid conflicts with existing profiles in the account
func generateUniquePrecedence(rnd string) int {
	h := fnv.New32a()
	h.Write([]byte(rnd))
	// Generate a precedence in range 101-999 to avoid conflicts
	// API requires precedence between 1-999 for custom profiles
	return 101 + int(h.Sum32()%899)
}

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_full.tf
var v4FullConfig string

//go:embed testdata/v5_full.tf
var v5FullConfig string

//go:embed testdata/v4_deprecated_name.tf
var v4DeprecatedNameConfig string

//go:embed testdata/v5_deprecated_name.tf
var v5DeprecatedNameConfig string

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Basic tests basic migration with minimal config
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Basic(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("name"), knownvalue.StringExact("Test Custom Profile Basic")),

		// Verify fallback domain migrated correctly
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("policy_id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
	}

	// Use MigrationV2TestStepWithStateNormalization like tunnel test
	// This handles multi-resource migration with cross-references
	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_BasicFromV5 tests v5 to v5 version bump
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_BasicFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5BasicConfig, rnd, accountID, precedence)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
			},
			{
				// Step 2: Run migration (version bump)
				PreConfig: func() {
					acctest.WriteOutConfig(t, v5Config, tmpDir)
					acctest.RunMigrationV2Command(t, v5Config, tmpDir, sourceVer, targetVer)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Full tests migration with all fields
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Full(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4FullConfig, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("name"), knownvalue.StringExact("Test Custom Profile Full")),

		// Verify fallback domain migrated correctly with all fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("policy_id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
		// Note: Can't easily test exact values due to Set ordering
	}

	// Use MigrationV2TestStepWithStateNormalization like tunnel test
	// This handles multi-resource migration with cross-references
	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_FullFromV5 tests v5 to v5 version bump with all fields
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_FullFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5FullConfig, rnd, accountID, precedence)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
			},
			{
				// Step 2: Run migration (version bump)
				PreConfig: func() {
					acctest.WriteOutConfig(t, v5Config, tmpDir)
					acctest.RunMigrationV2Command(t, v5Config, tmpDir, sourceVer, targetVer)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedName tests migration from both deprecated resource names
// (cloudflare_device_settings_policy → cloudflare_zero_trust_device_custom_profile,
// cloudflare_fallback_domain → cloudflare_zero_trust_device_custom_profile_local_domain_fallback)
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedName(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4DeprecatedNameConfig, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated to new custom profile name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("name"), knownvalue.StringExact("Deprecated Name Profile")),

		// Verify fallback domain migrated to new custom profile fallback domain name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("policy_id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
	}

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider using deprecated resource names
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedNameFromV5 tests v5 to v5 version bump after deprecated name migration
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedNameFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5DeprecatedNameConfig, rnd, accountID, precedence)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
			},
			{
				// Step 2: Run migration (version bump)
				PreConfig: func() {
					acctest.WriteOutConfig(t, v5Config, tmpDir)
					acctest.RunMigrationV2Command(t, v5Config, tmpDir, sourceVer, targetVer)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						ExpectEmptyPlanExceptPrecedenceDrift,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}
