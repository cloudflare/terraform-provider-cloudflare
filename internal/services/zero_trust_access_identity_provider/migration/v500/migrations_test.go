package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Key changes: cloudflare_access_identity_provider → cloudflare_zero_trust_access_identity_provider
// - Config block to attribute conversion
// - scim_config block to attribute conversion
// - idp_public_cert → idp_public_certs (string → list)
// - api_token removed from config
// - group_member_deprovision removed from scim_config

// Embed migration test configuration files

//go:embed testdata/v4_onetimepin.tf
var v4OnetimepinConfig string

//go:embed testdata/v5_onetimepin.tf
var v5OnetimepinConfig string

//go:embed testdata/v4_onetimepin_zone.tf
var v4OnetimepinZoneConfig string

//go:embed testdata/v5_onetimepin_zone.tf
var v5OnetimepinZoneConfig string

//go:embed testdata/v4_github_oauth.tf
var v4GithubOauthConfig string

//go:embed testdata/v5_github_oauth.tf
var v5GithubOauthConfig string

//go:embed testdata/v4_azuread_scim.tf
var v4AzureadScimConfig string

//go:embed testdata/v5_azuread_scim.tf
var v5AzureadScimConfig string

//go:embed testdata/v4_linkedin_oauth.tf
var v4LinkedinOauthConfig string

//go:embed testdata/v5_linkedin_oauth.tf
var v5LinkedinOauthConfig string

//go:embed testdata/v4_azuread_basic.tf
var v4AzureadBasicConfig string

//go:embed testdata/v5_azuread_basic.tf
var v5AzureadBasicConfig string

//go:embed testdata/v4_saml.tf
var v4SamlConfig string

//go:embed testdata/v5_saml.tf
var v5SamlConfig string

// TestMigrateAccessIdentityProvider_OneTimePin_Basic tests basic OneTimePin migration
func TestMigrateAccessIdentityProvider_OneTimePin_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4OnetimepinConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5OnetimepinConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("onetimepin")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
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
		})
	}
}

// TestMigrateAccessIdentityProvider_OneTimePin_Zone tests zone-scoped OneTimePin migration
func TestMigrateAccessIdentityProvider_OneTimePin_Zone(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4OnetimepinZoneConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5OnetimepinZoneConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
}

// TestMigrateAccessIdentityProvider_GitHub_OAuth tests GitHub OAuth migration
func TestMigrateAccessIdentityProvider_GitHub_OAuth(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4GithubOauthConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5GithubOauthConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
}

// TestMigrateAccessIdentityProvider_AzureAD_SCIM tests AzureAD with SCIM migration
func TestMigrateAccessIdentityProvider_AzureAD_SCIM(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4AzureadScimConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5AzureadScimConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"client_id":     knownvalue.StringExact("test"),
							"client_secret": knownvalue.StringExact("test"),
							"directory_id":  knownvalue.StringExact("directory"),
							"redirect_url":  knownvalue.NotNull(),
						})),
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
		})
	}
}

// TestMigrateAccessIdentityProvider_SAML tests SAML with idp_public_cert → idp_public_certs migration
func TestMigrateAccessIdentityProvider_SAML(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5SamlConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saml")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"issuer_url":     knownvalue.StringExact("jumpcloud"),
							"sso_target_url": knownvalue.StringExact("https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"),
							"attributes": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("email"),
								knownvalue.StringExact("username"),
							}),
							"sign_request": knownvalue.Bool(true),
							"idp_public_certs": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UEC.....GF/Q2/MHadws97cZguTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"),
							}),
							"redirect_url": knownvalue.NotNull(),
						})),
					}),
					{
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
					},
				},
			})
		})
	}
}

// TestMigrateAccessIdentityProvider_LinkedIn_OAuth tests LinkedIn OAuth migration
func TestMigrateAccessIdentityProvider_LinkedIn_OAuth(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4LinkedinOauthConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5LinkedinOauthConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
}

// TestMigrateAccessIdentityProvider_AzureAD_Basic tests basic AzureAD migration (no SCIM)
func TestMigrateAccessIdentityProvider_AzureAD_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4AzureadBasicConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5AzureadBasicConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
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
		})
	}
}

// TestMigrateAccessIdentityProvider_SAML_V4_CertRename tests SAML migration from v4 where
// idp_public_cert (write-only in v4 API) gets renamed to idp_public_certs.
// Since the v4 API doesn't return idp_public_cert in reads, the migrated state has idp_public_certs=nil,
// causing a legitimate Update action when applied with the v5 config that specifies the cert.
func TestMigrateAccessIdentityProvider_SAML_V4_CertRename(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4SamlConfig, rnd, accountID)

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
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: testConfig,
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, testConfig, tmpDir)
					acctest.RunMigrationV2Command(t, testConfig, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saml")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"issuer_url":     knownvalue.StringExact("jumpcloud"),
						"sso_target_url": knownvalue.StringExact("https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"),
						"sign_request":   knownvalue.Bool(true),
						"idp_public_certs": knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UEC.....GF/Q2/MHadws97cZguTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"),
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
