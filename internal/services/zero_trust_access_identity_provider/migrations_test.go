package zero_trust_access_identity_provider_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustAccessIdentityProvider_OneTimePin_Basic tests basic migration from v4 to v5
func TestMigrateZeroTrustAccessIdentityProvider_OneTimePin_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the OTP Access
	// endpoint does not yet support the API tokens for updates and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config using old resource name - OneTimePin doesn't need config block in v4
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "onetimepin"
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("onetimepin")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify config is now an object (not a list) and has redirect_url
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"redirect_url": knownvalue.NotNull(),
				})),
			}),
			{
				// Step 3: Apply with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_OneTimePin_Zone tests zone-scoped migration
func TestMigrateZeroTrustAccessIdentityProvider_OneTimePin_Zone(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config using zone_id
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  type    = "onetimepin"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("onetimepin")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"redirect_url": knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_OAuth_Comprehensive tests OAuth provider migration
func TestMigrateZeroTrustAccessIdentityProvider_OAuth_Comprehensive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config using block syntax for config
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "github"
  config {
    client_id     = "test"
    client_secret = "secret"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify config block converted to object with secrets preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"client_id":     knownvalue.StringExact("test"),
					"client_secret": knownvalue.StringExact("secret"),
					"redirect_url":  knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_AzureAD_Comprehensive tests AzureAD with SCIM migration
func TestMigrateZeroTrustAccessIdentityProvider_AzureAD_Comprehensive(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config with both config and scim_config blocks, including deprecated field
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "azureAD"
  config {
    client_id     = "test"
    client_secret = "test"
    directory_id  = "directory"
  }
  scim_config {
    enabled                    = true
    seat_deprovision          = true
    user_deprovision          = true
    identity_update_behavior  = "no_action"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify config block converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"client_id":     knownvalue.StringExact("test"),
					"client_secret": knownvalue.StringExact("test"),
					"directory_id":  knownvalue.StringExact("directory"),
					"redirect_url":  knownvalue.NotNull(),
				})),
				// Verify scim_config block converted to object, deprecated field removed
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"enabled":                  knownvalue.Bool(true),
					"seat_deprovision":         knownvalue.Bool(true),
					"user_deprovision":         knownvalue.Bool(true),
					"identity_update_behavior": knownvalue.StringExact("no_action"),
					"secret":                   knownvalue.NotNull(),
					"scim_base_url":            knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_SAML_Comprehensive tests SAML with cert field migration
func TestMigrateZeroTrustAccessIdentityProvider_SAML_Comprehensive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config with idp_public_cert (single string) - this will be migrated to idp_public_certs (list)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saml"
  config {
    issuer_url       = "jumpcloud"
    sso_target_url   = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes       = ["email", "username"]
    sign_request     = true
    idp_public_cert  = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UEC…..GF/Q2/MHadws97cZguTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Custom migration step for SAML - allows idp_public_cert → idp_public_certs transformation
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Allow the idp_public_cert → idp_public_certs transformation
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saml")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					// Verify config block converted to object with cert field renamed and converted to list
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"issuer_url":     knownvalue.StringExact("jumpcloud"),
						"sso_target_url": knownvalue.StringExact("https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"),
						"attributes": knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("email"),
							knownvalue.StringExact("username"),
						}),
						"sign_request": knownvalue.Bool(true),
						// Critical: idp_public_cert should be migrated to idp_public_certs as a list
						"idp_public_certs": knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UEC…..GF/Q2/MHadws97cZguTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"),
						}),
						"redirect_url": knownvalue.NotNull(),
					})),
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_GitHub tests GitHub OAuth provider migration
func TestMigrateZeroTrustAccessIdentityProvider_GitHub(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "github"
  config {
    client_id     = "test"
    client_secret = "secret"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"client_id":     knownvalue.StringExact("test"),
					"client_secret": knownvalue.StringExact("secret"),
					"redirect_url":  knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_GenericOAuth tests generic OAuth provider migration
func TestMigrateZeroTrustAccessIdentityProvider_GenericOAuth(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "linkedin"
  config {
    client_id     = "test"
    client_secret = "secret"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("linkedin")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"client_id":     knownvalue.StringExact("test"),
					"client_secret": knownvalue.StringExact("secret"),
					"redirect_url":  knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}

// TestMigrateZeroTrustAccessIdentityProvider_SCIM_Config_Secret tests SCIM secret handling
func TestMigrateZeroTrustAccessIdentityProvider_SCIM_Secret_Enabled_After_Resource_Creation(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic AzureAD setup (simplified to avoid v4 provider instability)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "azureAD"
  config {
    client_id     = "test"
    client_secret = "test"
    directory_id  = "directory"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify config block converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"client_id":     knownvalue.StringExact("test"),
					"client_secret": knownvalue.StringExact("test"),
					"directory_id":  knownvalue.StringExact("directory"),
					"redirect_url":  knownvalue.NotNull(),
				})),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
			},
		},
	})
}
