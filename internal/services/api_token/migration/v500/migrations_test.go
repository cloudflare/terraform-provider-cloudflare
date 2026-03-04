package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const (
	legacyProviderVersion  = "4.52.1"                // Last stable v4 release
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_condition.tf
var v4WithConditionConfig string

//go:embed testdata/v5_with_condition.tf
var v5WithConditionConfig string

//go:embed testdata/v4_with_ttl.tf
var v4WithTTLConfig string

//go:embed testdata/v5_with_ttl.tf
var v5WithTTLConfig string

//go:embed testdata/v4_deny_effect.tf
var v4DenyEffectConfig string

//go:embed testdata/v5_deny_effect.tf
var v5DenyEffectConfig string

//go:embed testdata/v4_condition_not_in.tf
var v4ConditionNotInConfig string

//go:embed testdata/v5_condition_not_in.tf
var v5ConditionNotInConfig string

//go:embed testdata/v4_multiple_policies.tf
var v4MultiplePoliciesConfig string

//go:embed testdata/v5_multiple_policies.tf
var v5MultiplePoliciesConfig string

//go:embed testdata/v4_multiple_resources.tf
var v4MultipleResourcesConfig string

//go:embed testdata/v5_multiple_resources.tf
var v5MultipleResourcesConfig string

//go:embed testdata/v4_mixed_effects.tf
var v4MixedEffectsConfig string

//go:embed testdata/v5_mixed_effects.tf
var v5MixedEffectsConfig string

// V5→V5 upgrade test configs (early v5 with map resources → latest v5 with jsonencode)

//go:embed testdata/v5_early_basic_with_condition.tf
var v5EarlyBasicWithConditionConfig string

//go:embed testdata/v5_latest_basic_with_condition.tf
var v5LatestBasicWithConditionConfig string

//go:embed testdata/v5_early_basic_map.tf
var v5EarlyBasicMapConfig string

//go:embed testdata/v5_latest_basic_jsonencode.tf
var v5LatestBasicJsonencodeConfig string

//go:embed testdata/v5_early_with_ttl.tf
var v5EarlyWithTTLConfig string

//go:embed testdata/v5_latest_with_ttl.tf
var v5LatestWithTTLConfig string

//go:embed testdata/v5_early_complex_resources.tf
var v5EarlyComplexResourcesConfig string

//go:embed testdata/v5_latest_complex_resources.tf
var v5LatestComplexResourcesConfig string

//go:embed testdata/v5_early_nested_resources_flat.tf
var v5EarlyNestedResourcesFlatConfig string

//go:embed testdata/v5_early_nested_resources_nested.tf
var v5EarlyNestedResourcesNestedConfig string

//go:embed testdata/v5_latest_nested_resources.tf
var v5LatestNestedResourcesConfig string

//go:embed testdata/v5_early_both_formats.tf
var v5EarlyBothFormatsConfig string

//go:embed testdata/v5_latest_both_formats.tf
var v5LatestBothFormatsConfig string

// TestMigrateAPIToken_V4ToV5_Basic tests basic migration of api_token.
// Verifies: policy→policies rename, permission_groups restructure, resources→JSON.
func TestMigrateAPIToken_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4BasicConfig, rnd)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5BasicConfig, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck:   func() { acctest.TestAccPreCheck(t) },
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
								knownvalue.StringExact("allow"),
							),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
								knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
							),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
								knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_WithCondition tests migration with condition block.
// Verifies: condition array→object unwrap, request_ip array→object unwrap.
func TestMigrateAPIToken_V4ToV5_WithCondition(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4WithConditionConfig, rnd)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5WithConditionConfig, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck:   func() { acctest.TestAccPreCheck(t) },
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0),
								knownvalue.StringExact("192.0.2.1/32"),
							),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
								knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_WithTTL tests migration with token TTL settings.
// Verifies: not_before and expires_on timestamp preservation.
func TestMigrateAPIToken_V4ToV5_WithTTL(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4WithTTLConfig, rnd)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5WithTTLConfig, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck:   func() { acctest.TestAccPreCheck(t) },
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2027-01-01T00:00:00Z")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
								knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_DenyEffect tests migration with "deny" effect policy.
func TestMigrateAPIToken_V4ToV5_DenyEffect(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4DenyEffectConfig, rnd)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5DenyEffectConfig, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck:   func() { acctest.TestAccPreCheck(t) },
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
								knownvalue.StringExact("deny"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_ConditionNotIn tests migration with condition using only not_in.
func TestMigrateAPIToken_V4ToV5_ConditionNotIn(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4ConditionNotInConfig, rnd)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5ConditionNotInConfig, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck:   func() { acctest.TestAccPreCheck(t) },
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"),
								knownvalue.ListSizeExact(1),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_MultiplePolicies tests migration with multiple policies.
func TestMigrateAPIToken_V4ToV5_MultiplePolicies(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MultiplePoliciesConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MultiplePoliciesConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
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
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
								knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
							),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
								knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_MultipleResources tests migration with multiple resource keys
// in a single policy to ensure proper JSON encoding with multiple entries.
func TestMigrateAPIToken_V4ToV5_MultipleResources(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MultipleResourcesConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MultipleResourcesConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
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
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAPIToken_V4ToV5_MixedEffects tests migration with mixed allow/deny policies.
func TestMigrateAPIToken_V4ToV5_MixedEffects(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MixedEffectsConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MixedEffectsConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
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
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
								knownvalue.StringExact("allow"),
							),
							statecheck.ExpectKnownValue(resourceName,
								tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("effect"),
								knownvalue.StringExact("deny"),
							),
						},
					),
				},
			})
		})
	}
}

// ==================== V5 Internal Upgrade Tests ====================
// These test the production v0→v1 upgrader (early v5 with map resources → current v5 with JSON resources).
// policies as attribute) is incompatible with the v4 source schema (condition as array block, policy as block)
// used by slot 0 in test mode. Both v4 and v5-early share schema_version=0.

func TestMigrateAPITokenFromV5MapToJSON(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with early v5 (schema_version=0, resources as map)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.0.0"}}, Config: fmt.Sprintf(v5EarlyBasicWithConditionConfig, rnd)},
			// Step 2: Stepping stone through v5.18 (applies 0→1 upgrade)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.18.0"}}, Config: fmt.Sprintf(v5LatestBasicWithConditionConfig, rnd)},
			// Step 3: Upgrade to current (applies 1→500 upgrade)
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestBasicWithConditionConfig, rnd), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5ComplexResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with early v5 (schema_version=0, resources as map)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.0.0"}}, Config: fmt.Sprintf(v5EarlyComplexResourcesConfig, rnd, accountID)},
			// Step 2: Stepping stone through v5.18 (applies 0→1 upgrade)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.18.0"}}, Config: fmt.Sprintf(v5LatestComplexResourcesConfig, rnd, accountID)},
			// Step 3: Upgrade to current (applies 1→500 upgrade)
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestComplexResourcesConfig, rnd, accountID), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5WithTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with early v5 (schema_version=0)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.0.0"}}, Config: fmt.Sprintf(v5EarlyWithTTLConfig, rnd)},
			// Step 2: Stepping stone through v5.18 (applies 0→1 upgrade)
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.18.0"}}, Config: fmt.Sprintf(v5LatestWithTTLConfig, rnd)},
			// Step 3: Upgrade to current (applies 1→500 upgrade)
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestWithTTLConfig, rnd), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2027-01-01T00:00:00Z")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5RemovedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.0.0"}}, Config: fmt.Sprintf(v5EarlyBasicMapConfig, rnd)},
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestBasicJsonencodeConfig, rnd), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5_4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.4.0"}}, Config: fmt.Sprintf(v5EarlyBasicMapConfig, rnd), ExpectNonEmptyPlan: true},
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestBasicJsonencodeConfig, rnd), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5_7(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.7.0"}}, Config: fmt.Sprintf(v5EarlyWithTTLConfig, rnd)},
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestWithTTLConfig, rnd), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5_ComplexNestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.7.0"}}, Config: fmt.Sprintf(v5EarlyNestedResourcesNestedConfig, rnd, accountID), ExpectError: regexp.MustCompile(`Incorrect attribute value type`)},
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.7.0"}}, Config: fmt.Sprintf(v5EarlyNestedResourcesFlatConfig, rnd, accountID)},
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestNestedResourcesConfig, rnd, accountID), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID))),
			}},
		},
	})
}

func TestMigrateAPITokenFromV5_4_BothResourceFormats(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_api_token.%s", rnd)
	tmpDir := t.TempDir()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) }, WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{ExternalProviders: map[string]resource.ExternalProvider{"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.4.0"}}, Config: fmt.Sprintf(v5EarlyBothFormatsConfig, rnd, accountID), ExpectNonEmptyPlan: true},
			{ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories, Config: fmt.Sprintf(v5LatestBothFormatsConfig, rnd, accountID), ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"), knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID))),
			}},
		},
	})
}
