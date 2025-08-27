package zero_trust_access_application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustAccessApplication_BasicV4toV5 tests migration from v4.52.1 to current v5
func TestMigrateZeroTrustAccessApplication_BasicV4toV5(t *testing.T) {
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
	tmpDir := t.TempDir()

	// V4 config - explicitly set computed attributes to avoid drift
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  http_only_cookie_attribute = false
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4.52.1 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_BasicV5toV5 tests migration from v5.7.1 to current v5
func TestMigrateZeroTrustAccessApplication_BasicV5toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// V5 config - explicitly set http_only_cookie_attribute to avoid drift
	v5Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  http_only_cookie_attribute = false
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.7.1 provider (recent stable version)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.7.1",
					},
				},
				Config: v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			// Step 2: Upgrade to current version and verify no changes
			acctest.MigrationTestStep(t, v5Config, tmpDir, "5.7.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_ZoneScopeV4toV5 tests zone-scoped application migration
func TestMigrateZeroTrustAccessApplication_ZoneScopeV4toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
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
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				},
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_ZoneScopeV5toV5 tests zone-scoped application v5 to v5 migration
func TestMigrateZeroTrustAccessApplication_ZoneScopeV5toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(`
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
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0",
					},
				},
				Config: v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				},
			},
			acctest.MigrationTestStep(t, v5Config, tmpDir, "5.0.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_CORSHeadersV4toV5 tests CORS headers migration from v4
func TestMigrateZeroTrustAccessApplication_CORSHeadersV4toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// V4 config with CORS headers - uses block syntax
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"

  cors_headers {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com"]
    allow_credentials = true
    max_age         = 600
  }
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_CORSHeadersV5toV5 tests CORS headers migration within v5
func TestMigrateZeroTrustAccessApplication_CORSHeadersV5toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// V5 config with CORS headers (uses = instead of block)
	v5Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[3]s"
  type         = "self_hosted"

  cors_headers = {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com"]
    allow_credentials = true
    max_age         = 600
  }
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0",
					},
				},
				Config: v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			acctest.MigrationTestStep(t, v5Config, tmpDir, "5.0.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SelfHostedDomainsV4toV5 tests self_hosted_domains migration
func TestMigrateZeroTrustAccessApplication_SelfHostedDomainsV4toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// Test config with unique self_hosted_domains
	v4Config := fmt.Sprintf(`
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
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("self_hosted_domains"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("self_hosted_domains"), knownvalue.ListSizeExact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_SAASAppBasicV4toV5 tests SAAS app migration from v4
func TestMigrateZeroTrustAccessApplication_SAASAppBasicV4toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// V4 config with SAAS app using blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"
  session_duration = "24h"

  saas_app {
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
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_DestinationsV4toV5 tests migration of destinations block to object
func TestMigrateZeroTrustAccessApplication_DestinationsV4toV5(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_application." + rnd
	tmpDir := t.TempDir()

	// V4 config with destinations as blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  http_only_cookie_attribute = false
  
  destinations {
    type = "public"
    uri  = "%[1]s.%[3]s"
  }
}`, rnd, accountID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4.52.1 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destinations"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
				// In v5, destinations should be transformed from blocks to objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destinations"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessApplication_MultiVersion tests migration across multiple version pairs
func TestMigrateZeroTrustAccessApplication_MultiVersion(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// Test cases for different version migrations
	testCases := []struct {
		name        string
		fromVersion string
		description string
	}{
		{
			name:        "V4_Latest_to_Current",
			fromVersion: "4.52.1",
			description: "Migration from v4 latest to current",
		},
		{
			name:        "V5_7_1_to_Current",
			fromVersion: "5.7.1",
			description: "Migration from v5.7.1 to current",
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_application." + rnd
			tmpDir := t.TempDir()

			// Config that should work across all versions
			// Set http_only_cookie_attribute explicitly to avoid drift
			config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  session_duration = "24h"
  auto_redirect_to_identity = false
  http_only_cookie_attribute = false
}`, rnd, accountID, domain)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specified version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.fromVersion,
							},
						},
						Config: config,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						},
					},
					// Step 2: Migrate to current version
					acctest.MigrationTestStep(t, config, tmpDir, tc.fromVersion, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					}),
				},
			})
		})
	}
}
