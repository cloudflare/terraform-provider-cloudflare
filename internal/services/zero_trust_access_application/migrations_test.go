package zero_trust_access_application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// Table-driven test for basic application types
func TestAccCloudflareZeroTrustAccessApplication_Migration_BasicTypes(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name           string
		appType        string
		versions       []string
		accountScoped  bool
		requiresDomain bool
	}{
		{
			name:           "SelfHostedAccount",
			appType:        "self_hosted",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  true,
			requiresDomain: true,
		},
		{
			name:           "SelfHostedZone",
			appType:        "self_hosted",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  false,
			requiresDomain: true,
		},
		{
			name:           "SSH",
			appType:        "ssh",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  true,
			requiresDomain: true,
		},
		{
			name:           "VNC",
			appType:        "vnc",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  true,
			requiresDomain: true,
		},
		{
			name:           "AppLauncher",
			appType:        "app_launcher",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  true,
			requiresDomain: false,
		},
		{
			name:           "Bookmark",
			appType:        "bookmark",
			versions:       []string{"4.52.1", "5.2.0", "5.7.1"},
			accountScoped:  true,
			requiresDomain: true,
		},
		{
			name:           "RDP",
			appType:        "rdp",
			versions:       []string{"5.2.0", "5.7.1"}, // RDP was added in v5
			accountScoped:  true,
			requiresDomain: true,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		for _, version := range tc.versions {
			version := version // capture range variable
			testName := fmt.Sprintf("%s_from_%s", tc.name, version)

			t.Run(testName, func(t *testing.T) {

				rnd := utils.GenerateRandomResourceName()
				resourceName := "cloudflare_zero_trust_access_application." + rnd
				tmpDir := t.TempDir()

				// Determine the resource name based on version
				var tfResourceName string
				if version == "4.52.1" {
					tfResourceName = "cloudflare_access_application"
				} else {
					tfResourceName = "cloudflare_zero_trust_access_application"
				}

				var v4Config string
				var checks []statecheck.StateCheck

				if tc.accountScoped {
					if tc.requiresDomain {
						v4Config = testAccCloudflareAccessApplicationBasic(tfResourceName, rnd, accountID, domain, tc.appType)
						checks = []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact(tc.appType)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						}
					} else {
						v4Config = testAccCloudflareAccessApplicationAppLauncher(tfResourceName, rnd, accountID)
						checks = []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact(tc.appType)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						}
					}
				} else {
					v4Config = testAccCloudflareAccessApplicationZone(tfResourceName, rnd, zoneID, domain, tc.appType)
					checks = []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact(tc.appType)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					}
				}

				// Determine provider version constraint
				var providerVersion string
				var providerSource string
				if version == "4.52.1" {
					providerVersion = "= 4.52.1"
					providerSource = "cloudflare/cloudflare"
				} else {
					providerVersion = fmt.Sprintf("= %s", version)
					providerSource = "cloudflare/cloudflare"
				}

				resource.Test(t, resource.TestCase{
					PreCheck: func() {
						acctest.TestAccPreCheck(t)
						if tc.accountScoped {
							acctest.TestAccPreCheck_AccountID(t)
						}
						if !tc.accountScoped {
							acctest.TestAccPreCheck_ZoneID(t)
						}
						if tc.requiresDomain {
							acctest.TestAccPreCheck_Domain(t)
						}
					},
					CheckDestroy: testAccCheckCloudflareAccessApplicationDestroy,
					WorkingDir:   tmpDir,
					Steps: []resource.TestStep{
						{
							// Step 1: Create with specific version provider
							ExternalProviders: map[string]resource.ExternalProvider{
								"cloudflare": {
									VersionConstraint: providerVersion,
									Source:            providerSource,
								},
							},
							Config: v4Config,
							// Known drift in 5.2.0
							ExpectNonEmptyPlan: version == "5.2.0",
						},
						// Step 2: Run migration and verify plan has no changes with state checks
						acctest.MigrationTestStep(t, v4Config, tmpDir, version, checks),
						/*
							{
								ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
								RefreshState: true,
								PreConfig: func() {
									configPath := fmt.Sprintf("%s/test_migration.tf", tmpDir)
									configContent, err := os.ReadFile(configPath)
									if err != nil {
										t.Logf("Error reading config file: %v", err)
									} else {
										t.Logf("Config file contents:\n%s", string(configContent))
									}
								},
							},
						*/

						// TODO do we need import tests?
					},
				})
			})
		}
	}
}

// Test basic migration without http_only_cookie_attribute set
func TestAccCloudflareZeroTrustAccessApplication_Migration_BasicWithoutHTTPOnly(t *testing.T) {
	t.Skip("Skipping test without http_only_cookie_attribute - needs investigation")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	versions := []string{"4.52.1", "5.2.0", "5.7.1"}

	for _, version := range versions {
		version := version
		testName := fmt.Sprintf("BasicNoHTTPOnly_from_%s", version)

		t.Run(testName, func(t *testing.T) {

			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_application." + rnd
			tmpDir := t.TempDir()

			// Determine the resource name based on version
			var tfResourceName string
			if version == "4.52.1" {
				tfResourceName = "cloudflare_access_application"
			} else {
				tfResourceName = "cloudflare_zero_trust_access_application"
			}

			// Config without http_only_cookie_attribute
			v4Config := fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  account_id       = "%[3]s"
  name             = "%[2]s"
  domain           = "%[2]s.%[4]s"
  type             = "self_hosted"
  session_duration = "24h"
}`, tfResourceName, rnd, accountID, domain)

			var providerVersion string
			if version == "4.52.1" {
				providerVersion = "= 4.52.1"
			} else {
				providerVersion = fmt.Sprintf("= %s", version)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareAccessApplicationDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: providerVersion,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					}),
				},
			})
		})
	}
}

// Test SAAS application migration
func TestAccCloudflareZeroTrustAccessApplication_Migration_SAAS(t *testing.T) {
	versions := []string{"4.52.1", "5.2.0", "5.7.1"}

	for _, version := range versions {
		version := version
		testName := fmt.Sprintf("SAAS_from_%s", version)

		t.Run(testName, func(t *testing.T) {

			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_application." + rnd
			tmpDir := t.TempDir()

			// Determine the resource name based on version
			var tfResourceName string
			if version == "4.52.1" {
				tfResourceName = "cloudflare_access_application"
			} else {
				tfResourceName = "cloudflare_zero_trust_access_application"
			}

			v4Config := testAccCloudflareAccessApplicationBasicSaas(tfResourceName, rnd, accountID)

			var providerVersion string
			if version == "4.52.1" {
				providerVersion = "= 4.52.1"
			} else {
				providerVersion = fmt.Sprintf("= %s", version)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				CheckDestroy: testAccCheckCloudflareAccessApplicationDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: providerVersion,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// Test CORS headers migration (block to object)
func TestAccCloudflareZeroTrustAccessApplication_Migration_CORS(t *testing.T) {
	versions := []string{"4.52.1", "5.2.0", "5.7.1"}

	for _, version := range versions {
		version := version
		testName := fmt.Sprintf("CORS_from_%s", version)

		t.Run(testName, func(t *testing.T) {

			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_application." + rnd
			tmpDir := t.TempDir()

			// Determine the resource name based on version
			var tfResourceName string
			if version == "4.52.1" {
				tfResourceName = "cloudflare_access_application"
			} else {
				tfResourceName = "cloudflare_zero_trust_access_application"
			}

			v4Config := testAccCloudflareAccessApplicationWithCORS(tfResourceName, rnd, accountID, domain)

			var providerVersion string
			if version == "4.52.1" {
				providerVersion = "= 4.52.1"
			} else {
				providerVersion = fmt.Sprintf("= %s", version)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareAccessApplicationDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: providerVersion,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// Test helper functions for configurations
func testAccCloudflareAccessApplicationBasic(resourceName, rnd, accountID, domain, appType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  account_id       = "%[3]s"
  name             = "%[2]s"
  domain           = "%[2]s.%[4]s"
  type             = "%[5]s"
  session_duration = "24h"
  http_only_cookie_attribute = false
}`, resourceName, rnd, accountID, domain, appType)
}

func testAccCloudflareAccessApplicationZone(resourceName, rnd, zoneID, domain, appType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  zone_id          = "%[3]s"
  name             = "%[2]s"
  domain           = "%[2]s.%[4]s"
  type             = "%[5]s"
  session_duration = "24h"
}`, resourceName, rnd, zoneID, domain, appType)
}

func testAccCloudflareAccessApplicationBasicSaas(resourceName, rnd, accountID string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  account_id = "%[3]s"
  name       = "%[2]s"
  type       = "saas"
  
  saas_app {
    consumer_service_url = "https://example.com"
    sp_entity_id         = "https://example.com"
    name_id_format       = "email"
    sso_endpoint         = "https://example.com/sso"
  }
}`, resourceName, rnd, accountID)
}

func testAccCloudflareAccessApplicationAppLauncher(resourceName, rnd, accountID string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  account_id                 = "%[3]s"
  name                       = "%[2]s"
  type                       = "app_launcher"
  app_launcher_visible       = true
  app_launcher_logo_url      = "https://example.com/logo.png"
}`, resourceName, rnd, accountID)
}

func testAccCloudflareAccessApplicationWithCORS(resourceName, rnd, accountID, domain string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  account_id = "%[3]s"
  name       = "%[2]s"
  domain     = "%[2]s.%[4]s"
  type       = "self_hosted"
  http_only_cookie_attribute = false
  
  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
    max_age           = 3600
  }
}`, resourceName, rnd, accountID, domain)
}

// Test migration with standalone policies referenced by ID
func TestAccCloudflareZeroTrustAccessApplication_Migration_StandalonePolicies(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	versions := []string{"4.52.1"}

	for _, version := range versions {
		version := version
		testName := fmt.Sprintf("StandalonePolicies_from_%s", version)

		t.Run(testName, func(t *testing.T) {

			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_application." + rnd
			tmpDir := t.TempDir()

			// Determine the resource name based on version
			var tfResourceName string
			var policyResourceName string
			if version == "4.52.1" {
				tfResourceName = "cloudflare_access_application"
				policyResourceName = "cloudflare_access_policy"
			} else {
				tfResourceName = "cloudflare_zero_trust_access_application"
				policyResourceName = "cloudflare_zero_trust_access_policy"
			}

			// Config with standalone policies referenced by ID
			v4Config := fmt.Sprintf(`
resource "%[3]s" "allow_%[2]s" {
  account_id     = "%[4]s"
  name           = "Allow Policy"
  decision       = "allow"
  include {
    certificate = true
  }
}

resource "%[3]s" "deny_%[2]s" {
  account_id     = "%[4]s"
  name           = "Deny Policy"
  decision       = "deny"
  session_duration = "24h"
  include {
    everyone = true
  }
}

resource "%[1]s" "%[2]s" {
  account_id       = "%[4]s"
  name             = "%[2]s"
  domain           = "%[2]s.%[5]s"
  type             = "self_hosted"
  session_duration = "24h"
  http_only_cookie_attribute = false
  
  policies = [
    %[3]s.allow_%[2]s.id,
    %[3]s.deny_%[2]s.id
  ]
}`, tfResourceName, rnd, policyResourceName, accountID, domain)

			var providerVersion string
			if version == "4.52.1" {
				providerVersion = "= 4.52.1"
			} else {
				providerVersion = fmt.Sprintf("= %s", version)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareAccessApplicationDestroy,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: providerVersion,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						// After migration, policies should be embedded in the application resource
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					}),
				},
			})
		})
	}
}
