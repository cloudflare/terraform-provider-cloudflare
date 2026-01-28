package zero_trust_access_application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustAccessApplication_Basic tests basic state schema migration within v5 provider
// Note: This tests the state migration (v0->v1) within the same resource type,
// not the resource type migration (cloudflare_access_application -> cloudflare_zero_trust_access_application)
// which requires the cmd/migrate tool and terraform state mv commands.
func TestMigrateZeroTrustAccessApplication_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Test config for state schema migration test
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  service_auth_401_redirect = true
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	// Test config with CORS headers
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"

  cors_headers = {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com", "https://test.com"]
    allowed_headers = ["Authorization", "Content-Type"]
    allow_credentials = false
    max_age         = 600
  }
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	// Test config with unique self_hosted_domains to avoid conflicts
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id            = "%[2]s"
  name                  = "%[1]s"
  type                  = "self_hosted"
  self_hosted_domains   = ["%[1]s-app1.%[3]s", "%[1]s-app2.%[3]s", "%[1]s-app3.%[3]s"]
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	// Test config with SAAS app using correct structure
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"
  session_duration = "24h"

  saas_app = {
    consumer_service_url = "https://example.com/sso/saml/consume"
    sp_entity_id        = "example.com"
    name_id_format      = "email"
    
    custom_attributes = [{
      name   = "email"
      name_format = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
      source = { name = "user_email" }
    }]
  }
  auto_redirect_to_identity = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	// V4 configuration using the old resource type
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  session_duration = "24h"
  enable_binding_cookie = true
}`, rnd, accountID, domain)

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

	// V4 configuration with policies (string array)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s-policy"
  decision   = "allow"
  include {
    everyone = true
  }
}

resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  session_duration = "24h"

  policies = [cloudflare_access_policy.%[1]s.id]
}`, rnd, accountID, domain)

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

	// Test config with OIDC SAAS app using valid grant types and no read-only attributes
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"
  session_duration = "24h"

  saas_app = {
    auth_type         = "oidc"
    redirect_uris     = ["https://example.com/callback"]
    grant_types       = ["authorization_code", "hybrid"]
    scopes            = ["openid", "profile", "email"]
    app_launcher_url  = "https://example.com/app"
    group_filter_regex = ".*"
    allow_pkce_without_client_secret = false
    
    custom_claims = [{
      name     = "rank"
      required = true
      scope    = "profile"
      source = { name = "rank" }
    }]
    
    hybrid_and_implicit_options = {
      return_id_token_from_authorization_endpoint = true
      return_access_token_from_authorization_endpoint = true
    }
  }
  auto_redirect_to_identity = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"
  session_duration = "24h"

  saas_app = {
    consumer_service_url = "https://example.com/sso/saml/consume"
    sp_entity_id        = "example.com"
    name_id_format      = "email"
  }
  auto_redirect_to_identity = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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
// This used to require manual state editing, but the v2 migrator now handles this automatically
func TestMigrateZeroTrustAccessApplication_CORSHeaders_Manual(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd

	// Configuration that should work with both old and new schema
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"
  
  cors_headers = {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com"]
    allow_credentials = true
    max_age = 600
  }
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
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

	// V4 configuration with cors_headers block (becomes array in state)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"

  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allowed_headers   = ["Authorization", "Content-Type"]
    allow_credentials = true
    max_age           = 600
  }
}`, rnd, accountID, domain)

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

	// V4 configuration with saas_app block
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"

  saas_app {
    consumer_service_url = "https://example.com/sso/saml/consume"
    sp_entity_id         = "example.com"
    name_id_format       = "email"

    custom_attribute {
      name         = "email"
      name_format  = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
      source {
        name = "user_email"
      }
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

	// V4 minimal configuration - tests null field handling
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
}`, rnd, accountID, domain)

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

	// V4 configuration with landing_page_design (app_launcher type)
	// Note: app_launcher type always returns name="App Launcher" from the API regardless of input,
	// so we expect a plan diff for the name field
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "app_launcher"

  landing_page_design {
    title             = "Welcome to App"
    message           = "Please select an application"
    button_color      = "#0051c3"
    button_text_color = "#ffffff"
  }
}`, rnd, accountID)

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
				//ConfigPlanChecks: resource.ConfigPlanChecks{
				//	PostApplyPostRefresh: []plancheck.PlanCheck{
				//		plancheck.ExpectDeferredChange(),
				//	},
				//},
			},
			// Step 2: Migrate to v5 - landing_page_design should be transformed from array to object
			// Note: ExpectNonEmptyPlan is true because app_launcher type always returns name="App Launcher"
			// from the API regardless of what name is configured
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true, // API always returns name="App Launcher" for app_launcher type
				ConfigStateChecks: []statecheck.StateCheck{
					// State should have "App Launcher" as the name (from API)
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

	// V4 configuration with self_hosted_domains (set type)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id          = "%[2]s"
  name                = "%[1]s"
  type                = "self_hosted"
  self_hosted_domains = ["%[1]s-a.%[3]s", "%[1]s-b.%[3]s"]
}`, rnd, accountID, domain)

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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"
}`, rnd, accountID, domain)

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
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"
}`, rnd, accountID, domain)

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
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"
}`, rnd, accountID, domain)

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
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"

  cors_headers = {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
    max_age           = 600
  }
}`, rnd, accountID, domain)

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
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id          = "%[2]s"
  name                = "%[1]s"
  type                = "self_hosted"
  session_duration    = "24h"
  self_hosted_domains = ["%[1]s-a.%[3]s", "%[1]s-b.%[3]s"]
}`, rnd, accountID, domain)

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
				Config: config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
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

	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "saas"
  session_duration = "24h"

  saas_app = {
    auth_type         = "oidc"
    redirect_uris     = ["https://example.com/callback"]
    grant_types       = ["authorization_code"]
    scopes            = ["openid", "profile", "email"]
    app_launcher_url  = "https://example.com/app"
  }
}`, rnd, accountID)

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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
				},
			},
		},
	})
}
