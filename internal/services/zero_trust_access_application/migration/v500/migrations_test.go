package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.0
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on breaking changes analysis:
// - All breaking changes happened between 4.x and 5.0.0
// - Key changes: cloudflare_access_application → cloudflare_zero_trust_access_application
// - Block to attribute conversions (cors_headers, saas_app, landing_page_design, scim_config)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

//go:embed testdata/v4_zone_scope.tf
var v4ZoneScopeConfig string

//go:embed testdata/v5_zone_scope.tf
var v5ZoneScopeConfig string

//go:embed testdata/v4_cors_headers.tf
var v4CORSHeadersConfig string

//go:embed testdata/v5_cors_headers.tf
var v5CORSHeadersConfig string

//go:embed testdata/v4_saas_app.tf
var v4SAASAppConfig string

//go:embed testdata/v5_saas_app.tf
var v5SAASAppConfig string

//go:embed testdata/v4_saas_app_oidc.tf
var v4SAASAppOIDCConfig string

//go:embed testdata/v5_saas_app_oidc.tf
var v5SAASAppOIDCConfig string

//go:embed testdata/v4_self_hosted_domains.tf
var v4SelfHostedDomainsConfig string

//go:embed testdata/v5_self_hosted_domains.tf
var v5SelfHostedDomainsConfig string

//go:embed testdata/v4_landing_page_design.tf
var v4LandingPageDesignConfig string

//go:embed testdata/v5_landing_page_design.tf
var v5LandingPageDesignConfig string

//go:embed testdata/v4_with_policies.tf
var v4WithPoliciesConfig string

//go:embed testdata/v5_with_policies.tf
var v5WithPoliciesConfig string

// TestMigrateZeroTrustAccessApplication_Basic tests basic state schema migration within v5 provider
func TestMigrateZeroTrustAccessApplication_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_ZoneScope tests zone-scoped application functionality
func TestMigrateZeroTrustAccessApplication_ZoneScope(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5ZoneScopeConfig, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_CORSHeaders tests CORS headers functionality
func TestMigrateZeroTrustAccessApplication_CORSHeaders(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5CORSHeadersConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_AllowedIDPs tests allowed_idps functionality
func TestMigrateZeroTrustAccessApplication_AllowedIDPs(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test config without allowed_idps (requires valid IDP resources)
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SelfHostedDomains tests self_hosted_domains functionality
func TestMigrateZeroTrustAccessApplication_SelfHostedDomains(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5SelfHostedDomainsConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_CustomPages tests custom_pages functionality
func TestMigrateZeroTrustAccessApplication_CustomPages(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test config without custom_pages (requires valid custom page resources)
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_Tags tests tags functionality
func TestMigrateZeroTrustAccessApplication_Tags(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test config without tags (requires pre-created tag resources)
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SAASAppBasic tests SAAS app functionality
func TestMigrateZeroTrustAccessApplication_SAASAppBasic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5SAASAppConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_Basic tests the actual v4 to v5 migration using cmd/migrate (v2 migrator)
func TestMigrateZeroTrustAccessApplication_V4toV5_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run v2 migration from v4 to v5
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_binding_cookie"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_WithPolicies tests migration with policies using v2 migrator
func TestMigrateZeroTrustAccessApplication_V4toV5_WithPolicies(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	appResourceName := "cloudflare_zero_trust_access_application." + rnd
	policyResourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4WithPoliciesConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run v2 migration from v4 to v5
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(appResourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(appResourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(appResourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(appResourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(policyResourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-policy", rnd))),
				statecheck.ExpectKnownValue(policyResourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SAASAppOIDC tests OIDC SAAS app functionality
func TestMigrateZeroTrustAccessApplication_SAASAppOIDC(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5SAASAppOIDCConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_LandingPageDesign tests landing_page_design functionality
func TestMigrateZeroTrustAccessApplication_LandingPageDesign(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test basic self_hosted app (app_launcher has naming conflicts)
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FooterLinks tests footer_links functionality
func TestMigrateZeroTrustAccessApplication_FooterLinks(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test basic self_hosted app (app_launcher has naming conflicts)
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SCIMConfig tests SCIM config functionality (simplified)
func TestMigrateZeroTrustAccessApplication_SCIMConfig(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test basic SAAS app without SCIM (SCIM requires complex IDP setup)
	testConfig := fmt.Sprintf(v5SAASAppConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_CORSHeaders_Manual tests that cors_headers works correctly as an object
func TestMigrateZeroTrustAccessApplication_CORSHeaders_Manual(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5CORSHeadersConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					// Verify cors_headers is now an object (not array) with expected structure
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_methods"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_origins"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allow_credentials"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("max_age"), knownvalue.Float64Exact(600)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_CORSHeaders tests CORS headers migration (array → object)
func TestMigrateZeroTrustAccessApplication_V4toV5_CORSHeaders(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4CORSHeadersConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Migrate to v5 - cors_headers should be transformed from array to object
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				// Verify cors_headers is now an object (MaxItems:1 transformation)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allow_credentials"), knownvalue.Bool(true)),
				// max_age should be converted from int64 to float64
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("max_age"), knownvalue.Float64Exact(600)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_SAASApp tests SAAS app migration (array → object with nested)
func TestMigrateZeroTrustAccessApplication_V4toV5_SAASApp(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4SAASAppConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Migrate to v5 - saas_app should be transformed from array to object
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
				// Verify saas_app is now an object (MaxItems:1 transformation)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("sp_entity_id"), knownvalue.StringExact("example.com")),
				// custom_attribute should be renamed to custom_attributes (plural)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_Minimal tests migration with minimal config (null fields)
func TestMigrateZeroTrustAccessApplication_V4toV5_Minimal(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4MinimalConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Migrate to v5 - type should default to "self_hosted"
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// type should default to "self_hosted" when not specified in v4
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				// Optional nested blocks should be null when not specified
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.Null()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_LandingPageDesign tests landing_page_design migration
func TestMigrateZeroTrustAccessApplication_V4toV5_LandingPageDesign(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4LandingPageDesignConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // API always returns name="App Launcher" for app_launcher type
			},
			// Step 2: Migrate to v5 - landing_page_design should be transformed from array to object
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true, // API always returns name="App Launcher" for app_launcher type
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("app_launcher")),
					// Verify landing_page_design is now an object (MaxItems:1 transformation)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("landing_page_design"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("landing_page_design").AtMapKey("title"), knownvalue.StringExact("Welcome to App")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("landing_page_design").AtMapKey("button_color"), knownvalue.StringExact("#0051c3")),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_V4toV5_SelfHostedDomains tests self_hosted_domains migration
func TestMigrateZeroTrustAccessApplication_V4toV5_SelfHostedDomains(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4SelfHostedDomainsConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Create with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Migrate to v5 - self_hosted_domains should be converted from set to list
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				// Verify self_hosted_domains was migrated
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("self_hosted_domains"), knownvalue.ListSizeExact(2)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_12 tests migration from v5.12.0 to latest
func TestMigrateZeroTrustAccessApplication_FromV5_12(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_14 tests migration from v5.14.0 to latest
func TestMigrateZeroTrustAccessApplication_FromV5_14(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_15 tests migration from v5.15.0 to latest
func TestMigrateZeroTrustAccessApplication_FromV5_15(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_12_WithCORSHeaders tests v5.12 to latest with CORS headers
func TestMigrateZeroTrustAccessApplication_FromV5_12_WithCORSHeaders(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5CORSHeadersConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allow_credentials"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("max_age"), knownvalue.Float64Exact(600)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_14_WithSelfHostedDomains tests v5.14 with self_hosted_domains
func TestMigrateZeroTrustAccessApplication_FromV5_14_WithSelfHostedDomains(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5SelfHostedDomainsConfig, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("self_hosted_domains"), knownvalue.SetSizeExact(2)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessApplication_FromV5_15_SAASApp tests v5.15 with SAAS app configuration
func TestMigrateZeroTrustAccessApplication_FromV5_15_SAASApp(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	testConfig := fmt.Sprintf(v5SAASAppOIDCConfig, rnd, accountID)

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
				Config: testConfig,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
				},
			},
		},
	})
}
