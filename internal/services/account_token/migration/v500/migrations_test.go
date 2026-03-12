package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// Embed migration test configuration templates.
// v0 templates use map-based resources (pre-v500 schema).
// v500 templates use jsonencode-based resources (current schema).

//go:embed testdata/v0_basic.tf
var v0BasicConfig string

//go:embed testdata/v500_basic.tf
var v500BasicConfig string

//go:embed testdata/v0_with_condition.tf
var v0WithConditionConfig string

//go:embed testdata/v500_with_condition.tf
var v500WithConditionConfig string

//go:embed testdata/v0_with_ttl.tf
var v0WithTTLConfig string

//go:embed testdata/v500_with_ttl.tf
var v500WithTTLConfig string

//go:embed testdata/v0_with_expires.tf
var v0WithExpiresConfig string

//go:embed testdata/v500_with_expires.tf
var v500WithExpiresConfig string

//go:embed testdata/v0_multi_policy.tf
var v0MultiPolicyConfig string

//go:embed testdata/v500_multi_policy.tf
var v500MultiPolicyConfig string

//go:embed testdata/v0_complex_resources.tf
var v0ComplexResourcesConfig string

//go:embed testdata/v500_complex_resources.tf
var v500ComplexResourcesConfig string

//go:embed testdata/v0_nested_resources.tf
var v0NestedResourcesConfig string

//go:embed testdata/v500_nested_resources.tf
var v500NestedResourcesConfig string

//go:embed testdata/v0_basic_named.tf
var v0BasicNamedConfig string

//go:embed testdata/v500_basic_named.tf
var v500BasicNamedConfig string

//go:embed testdata/v0_mixed_flat.tf
var v0MixedFlatConfig string

//go:embed testdata/v500_mixed_both.tf
var v500MixedBothConfig string

// TestMigrateAccountTokenFromV5MapToJSON tests migration from v5 account_token
// with resources as map to latest with resources as JSON string.
func TestMigrateAccountTokenFromV5MapToJSON(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v0WithConditionConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500WithConditionConfig, rnd, accountID)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config: v5Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("account_id"),
						knownvalue.StringExact(accountID),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
						knownvalue.StringExact("allow"),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0),
						knownvalue.StringExact("192.0.2.1/32"),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5ComplexResources tests migration with complex nested resources.
func TestMigrateAccountTokenFromV5ComplexResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v0ComplexResourcesConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500ComplexResourcesConfig, rnd, accountID)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config:      v5Config,
				ExpectError: regexp.MustCompile(`string\s+required`),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_ComplexNestedResources tests that v5.10 can't handle
// nested resources but the latest version with jsonencode can.
func TestMigrateAccountTokenFromV5_ComplexNestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	v5NestedConfig := fmt.Sprintf(v0NestedResourcesConfig, rnd, accountID)
	latestNestedConfig := fmt.Sprintf(v500NestedResourcesConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "5.10.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config:      v5NestedConfig,
				ExpectError: regexp.MustCompile(`string\s+required`),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestNestedConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_10_BothResourceFormats tests migration with both
// flat and nested resource formats.
func TestMigrateAccountTokenFromV5_10_BothResourceFormats(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	v510Config := fmt.Sprintf(v0MixedFlatConfig, rnd, accountID)
	latestMixedConfig := fmt.Sprintf(v500MixedBothConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "5.10.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v510Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestMixedConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("%s-mixed", rnd)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5WithTTL tests migration with token TTL settings.
func TestMigrateAccountTokenFromV5WithTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	notBefore := "2025-01-01T00:00:00Z"
	expiresOn := time.Now().UTC().AddDate(0, 0, 30).Format(time.RFC3339)

	v5Config := fmt.Sprintf(v0WithTTLConfig, rnd, accountID, notBefore, expiresOn)
	latestConfig := fmt.Sprintf(v500WithTTLConfig, rnd, accountID, notBefore, expiresOn)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config: v5Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("not_before"),
						knownvalue.StringExact(notBefore),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact(expiresOn),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5NestedResources tests migration with the new
// nested resources capability.
func TestMigrateAccountTokenFromV5NestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v0BasicConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500NestedResourcesConfig, rnd, accountID)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config: v5Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5RemovedFields tests that computed fields are
// properly removed during migration (policy.id, permission_groups.meta/name).
func TestMigrateAccountTokenFromV5RemovedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v0BasicConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500BasicConfig, rnd, accountID)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config: v5Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_10 tests migration from v5.10.0 (earliest stable
// version for account_token).
func TestMigrateAccountTokenFromV5_10(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v510Config := fmt.Sprintf(v0BasicNamedConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500BasicNamedConfig, rnd, accountID)

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
						VersionConstraint: "5.10.0",
					},
				},
				Config: v510Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("%s-v510", rnd)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_11 tests migration from v5.11.0.
func TestMigrateAccountTokenFromV5_11(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	expiresOn := time.Now().Add(365 * 24 * time.Hour).UTC().Format("2006-01-02T15:04:05Z")

	v5Config := fmt.Sprintf(v0WithExpiresConfig, rnd, accountID, expiresOn)
	latestConfig := fmt.Sprintf(v500WithExpiresConfig, rnd, accountID, expiresOn)

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
						VersionConstraint: "5.11.0",
					},
				},
				Config: v5Config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact(expiresOn),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_12 tests migration from v5.12.0 (last version
// before the change) with multiple policies.
func TestMigrateAccountTokenFromV5_12(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v0MultiPolicyConfig, rnd, accountID)
	latestConfig := fmt.Sprintf(v500MultiPolicyConfig, rnd, accountID)

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
						VersionConstraint: "5.12.0",
					},
				},
				Config: v5Config,
				// v5.12 still has policy ordering issues that cause drift
				ExpectNonEmptyPlan: true,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}
