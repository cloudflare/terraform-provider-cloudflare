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
//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

//go:embed testdata/v4_rule_settings.tf
var v4RuleSettingsConfig string

//go:embed testdata/v5_rule_settings.tf
var v5RuleSettingsConfig string

//go:embed testdata/v4_nested.tf
var v4NestedConfig string

//go:embed testdata/v5_nested.tf
var v5NestedConfig string

//go:embed testdata/v4_complex_settings.tf
var v4ComplexSettingsConfig string

//go:embed testdata/v5_complex_settings.tf
var v5ComplexSettingsConfig string

//go:embed testdata/v4_empty_rule_settings.tf
var v4EmptyRuleSettingsConfig string

//go:embed testdata/v5_empty_rule_settings.tf
var v5EmptyRuleSettingsConfig string

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_Minimal tests basic field migrations with dual test cases
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_Minimal(t *testing.T) {
	// Zero Trust resources don't support API token authentication yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MinimalConfig, rnd, accountID, rnd)
			},
		},
		{
			name:           "from_v5", // Tests within v5 (version bump)
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MinimalConfig, rnd, accountID, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
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
					acctest.MigrationV2TestStepForGatewayPolicy(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, false,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-test-minimal-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Minimal policy for migration testing")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
							// Precedence is auto-calculated by API, just verify it exists
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_RuleSettings tests migration with rule_settings and field rename (block_page_reason → block_reason)
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_RuleSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
				return fmt.Sprintf(v4RuleSettingsConfig, rnd, accountID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5RuleSettingsConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
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
			// For from_v5 test case, v5 provider has computed field drift even on initial create
			if tc.useDevProvider {
				step1.ExpectNonEmptyPlan = true
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					acctest.MigrationV2TestStepForGatewayPolicy(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, true,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
							// Rule settings should be converted from block to attribute
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_page_enabled"), knownvalue.Bool(true)),
							// Field rename: block_page_reason → block_reason
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_reason"), knownvalue.StringExact("Access blocked by company policy")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("ip_categories"), knownvalue.Bool(true)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_Nested tests migration with nested blocks and field rename (message → msg)
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_Nested(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
				return fmt.Sprintf(v4NestedConfig, rnd, accountID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5NestedConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
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
			// For from_v5 test case, v5 provider has computed field drift even on initial create
			if tc.useDevProvider {
				step1.ExpectNonEmptyPlan = true
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					acctest.MigrationV2TestStepForGatewayPolicy(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, true,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
							// block_page_enabled should be present
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_page_enabled"), knownvalue.Bool(true)),
							// Nested notification_settings block should be converted to attribute
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
							// Field rename: message → msg
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("Connection blocked")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com/")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_ComplexSettings tests migration with check_session and payload_log
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_ComplexSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
				return fmt.Sprintf(v4ComplexSettingsConfig, rnd, accountID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5ComplexSettingsConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
					// For from_v5 test case, v5 provider has computed field drift on initial create
					// when rule_settings contains add_headers/override_ips
					ExpectNonEmptyPlan: true,
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
					acctest.MigrationV2TestStepForGatewayPolicy(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, true,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("allow")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("traffic"), knownvalue.StringExact("http.request.uri matches \".*api.*\"")),
							// All nested blocks should be converted to attributes
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("enforce"), knownvalue.Bool(true)),
							// Duration should be "24h0m0s"
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("duration"), knownvalue.StringExact("24h0m0s")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("payload_log").AtMapKey("enabled"), knownvalue.Bool(true)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_EmptyRuleSettings tests migration without rule_settings
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_EmptyRuleSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
				return fmt.Sprintf(v4EmptyRuleSettingsConfig, rnd, accountID)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5EmptyRuleSettingsConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
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
					acctest.MigrationV2TestStepForGatewayPolicy(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, false,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
							// Precedence is auto-calculated by API
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}
