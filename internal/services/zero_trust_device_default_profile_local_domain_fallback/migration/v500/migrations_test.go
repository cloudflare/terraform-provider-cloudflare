package v500_test

import (
	_ "embed"
	"fmt"
	"os"
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

//go:embed testdata/v4_null_policy_id.tf
var v4NullPolicyIDConfig string

//go:embed testdata/v5_null_policy_id.tf
var v5NullPolicyIDConfig string

//go:embed testdata/v4_multiple_domains.tf
var v4MultipleDomainsConfig string

//go:embed testdata/v5_multiple_domains.tf
var v5MultipleDomainsConfig string

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Basic tests basic migration with minimal config
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Basic(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		// Verify fallback domain migrated correctly
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
	}

	// Use MigrationV2TestStepWithStateNormalization for consistent migration testing
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_BasicFromV5 tests v5 to v5 version bump
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_BasicFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5BasicConfig, rnd, accountID)
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
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Full tests migration with all fields
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Full(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4FullConfig, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		// Verify fallback domain migrated correctly with all fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
		// Note: Can't easily test exact values due to Set ordering
	}

	// Use MigrationV2TestStepWithStateNormalization for consistent migration testing
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedName tests migration from the deprecated cloudflare_fallback_domain resource name
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedName(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4DeprecatedNameConfig, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
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
				// Step 1: Create with v4 provider using deprecated resource name
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedNameFromV5 tests v5 to v5 version bump with deprecated name migration
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedNameFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5DeprecatedNameConfig, rnd, accountID)
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
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyID tests migration where policy_id = null is treated as the default profile
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyID(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4NullPolicyIDConfig, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.NotNull()),
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
				// Step 1: Create with v4 provider using explicit policy_id = null
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyIDFromV5 tests v5 to v5 version bump where policy_id was null
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyIDFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5NullPolicyIDConfig, rnd, accountID)
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
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomains tests migration with 3 domains (Set → List)
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomains(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4MultipleDomainsConfig, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		// Verify all 3 domains were preserved; exact ordering not checked due to Set → List conversion
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(3)),
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
				// Step 1: Create with v4 provider with 3 domains
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomainsFromV5 tests v5 to v5 version bump with multiple domains
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomainsFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5MultipleDomainsConfig, rnd, accountID)
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
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(3)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_FullFromV5 tests v5 to v5 version bump with all fields
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_FullFromV5(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5FullConfig, rnd, accountID)
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
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}
