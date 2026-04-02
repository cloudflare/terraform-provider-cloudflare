package v500_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_email_domain.tf
var v4EmailDomainConfig string

//go:embed testdata/v4_complex.tf
var v4ComplexConfig string

//go:embed testdata/v4_array_explosion.tf
var v4ArrayExplosionConfig string

//go:embed testdata/v4_decision.tf
var v4DecisionConfig string

//go:embed testdata/v4_boolean.tf
var v4BooleanConfig string

//go:embed testdata/v4_service_token.tf
var v4ServiceTokenConfig string

//go:embed testdata/v4_unsupported.tf
var v4UnsupportedConfig string

//go:embed testdata/v4_connection_rules.tf
var v4ConnectionRulesConfig string

//go:embed testdata/v4_connection_rules_email.tf
var v4ConnectionRulesEmailConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v5_complex.tf
var v5ComplexConfig string

//go:embed testdata/v5_decision.tf
var v5DecisionConfig string

//go:embed testdata/v5_condition_types.tf
var v5ConditionTypesConfig string

// TestMigrateZeroTrustAccessPolicyMigrationFromV4Basic tests basic migration from v4 to v5
func TestMigrateZeroTrustAccessPolicyMigrationFromV4Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify email transformation: v4 email list -> v5 single nested object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4Complex tests migration with complex conditions
func TestMigrateZeroTrustAccessPolicyMigrationFromV4Complex(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4ComplexConfig, rnd, accountID)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_prompt"), knownvalue.StringExact("Why do you need access?")),
				// Verify approval_group -> approval_groups transformation
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("approvals_needed"), knownvalue.Float64Exact(2.0)),
				// Verify complex condition transformations - multiple objects for multiple values
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(8)), // email(2) + email_domain(2) + ip(2) + everyone + any_valid_service_token
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(3)), // email(1) + geo(2)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)), // certificate
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4OAuthProviders tests array explosion and attribute transformations
func TestMigrateZeroTrustAccessPolicyMigrationFromV4OAuthProviders(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4ArrayExplosionConfig, rnd, accountID)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify array explosion: email array (2 rules) = 2 include rules
				// Also verify exclude has 1 rule
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4DecisionTypes tests all decision types
func TestMigrateZeroTrustAccessPolicyMigrationFromV4DecisionTypes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		decision string
		name     string
	}{
		{"allow", "Allow"},
		{"deny", "Deny"},
		{"non_identity", "NonIdentity"},
		{"bypass", "Bypass"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Decision_%s", tc.name), func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(v4DecisionConfig, rnd, accountID, tc.decision)

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
								VersionConstraint: "~> 4.52.1",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact(tc.decision)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4OptionalBooleans tests boolean to object transformations
func TestMigrateZeroTrustAccessPolicyMigrationFromV4OptionalBooleans(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		boolField string
		testName  string
	}{
		{"everyone", "Everyone"},
		{"certificate", "Certificate"},
		{"any_valid_service_token", "AnyValidServiceToken"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Boolean_%s", tc.testName), func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(v4BooleanConfig, rnd, accountID, tc.boolField)

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
								VersionConstraint: "~> 4.52.1",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
						// Verify boolean -> empty object transformation
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey(tc.boolField), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4UnsupportedFeatures tests basic v4 to v5 migration functionality
func TestMigrateZeroTrustAccessPolicyMigrationFromV4UnsupportedFeatures(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4UnsupportedConfig, rnd, accountID)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify that basic migration works correctly:
				// - Basic attributes are preserved
				// - Include/exclude rules are properly transformed
				// - Session duration is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4ServiceTokens tests service token transformations
func TestMigrateZeroTrustAccessPolicyMigrationFromV4ServiceTokens(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4ServiceTokenConfig, rnd, accountID)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify any_valid_service_token transformation: boolean -> empty nested object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("any_valid_service_token"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyMigrationFromV4RemovedAttributes tests handling of deprecated attributes
func TestMigrateZeroTrustAccessPolicyMigrationFromV4RemovedAttributes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4UnsupportedConfig, rnd, accountID)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Note: Deprecated attributes (application_id, precedence, zone_id, etc.) are removed by state transformation
				// but we can't easily test their absence with current statecheck functions
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyFromV5_12 tests migration from v5.12.0 to latest
func TestMigrateZeroTrustAccessPolicyFromV5_12(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd

	config := fmt.Sprintf(v5BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.12 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.12.0",
					},
				},
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessPolicyFromV5_14 tests migration from v5.14.0 to latest
func TestMigrateZeroTrustAccessPolicyFromV5_14(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd

	config := fmt.Sprintf(v5BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.14 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.14.0",
					},
				},
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessPolicyFromV5_15 tests migration from v5.15.0 to latest
func TestMigrateZeroTrustAccessPolicyFromV5_15(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd

	config := fmt.Sprintf(v5BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.15 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.15.0",
					},
				},
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessPolicyFromV5_12_Complex tests v5.12 to latest with complex conditions
func TestMigrateZeroTrustAccessPolicyFromV5_12_Complex(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd

	config := fmt.Sprintf(v5ComplexConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.12 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.12.0",
					},
				},
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessPolicyFromV5_14_WithDecisionTypes tests v5.14 with different decision types
func TestMigrateZeroTrustAccessPolicyFromV5_14_WithDecisionTypes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		decision string
		name     string
	}{
		{"allow", "Allow"},
		{"deny", "Deny"},
		{"non_identity", "NonIdentity"},
		{"bypass", "Bypass"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Decision_%s", tc.name), func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd

			config := fmt.Sprintf(v5DecisionConfig, rnd, accountID, tc.decision)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: "5.14.0",
							},
						},
						Config: config,
					},
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						Config:                   config,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact(tc.decision)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicyFromV5_15_WithConditionTypes tests v5.15 with various condition types
func TestMigrateZeroTrustAccessPolicyFromV5_15_WithConditionTypes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd

	config := fmt.Sprintf(v5ConditionTypesConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.15 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.15.0",
					},
				},
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(6)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessPolicyEmailDomainTransformation tests that email_domain
// is correctly migrated from v4 list syntax to v5 object syntax.
// Reproduces TKT-005: email_domain = ["domain"] → email_domain = { domain = "domain" }
// Reported by research team (terraform-cfaccounts MR !7756).
func TestMigrateZeroTrustAccessPolicyEmailDomainTransformation(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	if domain == "" {
		t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
	}
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4EmailDomainConfig, rnd, accountID, domain)

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
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify email_domain transformation: list → object with domain field
				// v4: email_domain = ["example.com"]
				// v5: email_domain = { domain = "example.com" }
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"),
					knownvalue.StringExact(domain),
				),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyConnectionRules tests that connection_rules
// with application_id and precedence are correctly migrated from v4 to v5.
// Reproduces TKT-007: connection_rules stored as JSON array [] in v4 state
// causes "invalid JSON, expected '{', got '['" error during state upgrade.
// Reported by research team (terraform-cfaccounts MR !7756).
func TestMigrateZeroTrustAccessPolicyConnectionRules(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4ConnectionRulesConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			// With the PriorSchema: nil fix, state upgrade succeeds
			// The key success is that we don't get: "AttributeName("connection_rules"): invalid JSON, expected "{", got "["
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify include was transformed correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessPolicyStateMvScenario tests migration after manual `terraform state mv`.
// This reproduces the research-team issue where users did `terraform state mv` instead of using
// the `moved` block, resulting in cloudflare_zero_trust_access_policy resources with v4-format state.
//
// Error: AttributeName("connection_rules"): invalid JSON, expected "{", got "["
//
// To reproduce the bug (with PriorSchema: &v5Schema):
// 1. Create resource with v4 cloudflare_access_policy
// 2. Rename resource type in state to cloudflare_zero_trust_access_policy (simulating state mv)
// 3. Update config to use cloudflare_zero_trust_access_policy
// 4. Run v5 provider - triggers UpgradeState with v4-format data
// 5. With PriorSchema: &v5Schema, the framework fails to parse connection_rules=[] as object {}
func TestMigrateZeroTrustAccessPolicyStateMvScenario(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4ConnectionRulesEmailConfig, rnd, accountID)
	// v5 config uses cloudflare_zero_trust_access_policy
	v5Config := fmt.Sprintf(`resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  decision         = "allow"
  session_duration = "24h"

  include = [{
    everyone = {}
  }]

  connection_rules = {
    ssh = {
      usernames         = ["root", "admin"]
      allow_email_alias = true
    }
  }
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "~> 4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Simulate `terraform state mv` by renaming resource type in state file
				// This leaves cloudflare_zero_trust_access_policy with v4-format state data
				PreConfig: func() {
					renameResourceTypeInState(t, tmpDir, "cloudflare_access_policy", "cloudflare_zero_trust_access_policy", rnd)
				},
				// Step 3: Run v5 provider with renamed resource
				// This triggers UpgradeState (not MoveState) because resource type already matches
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// renameResourceTypeInState renames a resource type in the terraform state file.
// This simulates what `terraform state mv` does.
func renameResourceTypeInState(t *testing.T, tmpDir string, oldType string, newType string, resourceName string) {
	// Find the terraform working directory (work* subdirectory of tmpDir)
	matches, err := filepath.Glob(filepath.Join(tmpDir, "work*"))
	if err != nil || len(matches) == 0 {
		t.Fatalf("Could not find work directory in %s: %v", tmpDir, err)
	}
	workDir := matches[0]
	stateFile := filepath.Join(workDir, "terraform.tfstate")

	data, err := os.ReadFile(stateFile)
	if err != nil {
		t.Fatalf("Could not read state file %s: %v", stateFile, err)
	}

	// Parse state
	var state struct {
		Version   int `json:"version"`
		Resources []struct {
			Module    string          `json:"module,omitempty"`
			Mode      string          `json:"mode"`
			Type      string          `json:"type"`
			Name      string          `json:"name"`
			Provider  string          `json:"provider"`
			Instances json.RawMessage `json:"instances"`
		} `json:"resources"`
	}
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Could not parse state file: %v", err)
	}

	// Rename resource type
	found := false
	for i := range state.Resources {
		if state.Resources[i].Type == oldType && state.Resources[i].Name == resourceName {
			state.Resources[i].Type = newType
			state.Resources[i].Provider = "provider[\"registry.terraform.io/cloudflare/cloudflare\"]"
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Could not find resource %s.%s in state", oldType, resourceName)
	}

	// Write back
	newData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		t.Fatalf("Could not marshal state: %v", err)
	}

	if err := os.WriteFile(stateFile, newData, 0644); err != nil {
		t.Fatalf("Could not write state file: %v", err)
	}

	t.Logf("Renamed %s.%s to %s.%s in state file", oldType, resourceName, newType, resourceName)
}
