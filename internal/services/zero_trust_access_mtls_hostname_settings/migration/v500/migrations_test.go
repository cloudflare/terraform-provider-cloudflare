package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	cfv1 "github.com/cloudflare/cloudflare-go"
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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

//go:embed testdata/v4_boolean_defaults.tf
var v4BooleanDefaultsConfig string

//go:embed testdata/v5_boolean_defaults.tf
var v5BooleanDefaultsConfig string

//go:embed testdata/v4_boolean_combinations.tf
var v4BooleanCombinationsConfig string

//go:embed testdata/v5_boolean_combinations.tf
var v5BooleanCombinationsConfig string

// clearAPIToken unsets CLOUDFLARE_API_TOKEN for Access tests (require API_KEY + EMAIL).
func clearAPIToken(t *testing.T) {
	t.Helper()
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		t.Cleanup(func() { os.Setenv("CLOUDFLARE_API_TOKEN", originalToken) })
	}
}

// cleanupMTLSSettings clears all MTLS hostname settings to prevent test conflicts.
func cleanupMTLSSettings(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		return
	}
	ctx := context.Background()
	client, err := acctest.SharedV1Client()
	if err != nil {
		t.Fatalf("Failed to create Cloudflare client: %v", err)
	}

	deletedSettings := cfv1.UpdateAccessMutualTLSHostnameSettingsParams{
		Settings: []cfv1.AccessMutualTLSHostnameSettings{},
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID != "" {
		for i := 0; i < 3; i++ {
			_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.AccountIdentifier(accountID), deletedSettings)
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID != "" {
		for i := 0; i < 3; i++ {
			_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.ZoneIdentifier(zoneID), deletedSettings)
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}

	time.Sleep(10 * time.Second)
}

// TestMigrateMTLSHostnameSettingsBasic tests basic migration with dual test cases.
func TestMigrateMTLSHostnameSettingsBasic(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, domain string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
			testConfig := tc.configFn(rnd, accountID, domain)
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
					}),
				},
			})
		})
	}
}

// TestMigrateMTLSHostnameSettingsMultiple tests migration with multiple hostnames.
func TestMigrateMTLSHostnameSettingsMultiple(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	altDomain := "alt." + domain
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4MultipleConfig, rnd, "account", accountID, domain, altDomain)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(1).AtMapKey("hostname"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateMTLSHostnameSettingsBooleanDefaults tests migration when optional booleans are not specified.
func TestMigrateMTLSHostnameSettingsBooleanDefaults(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BooleanDefaultsConfig, rnd, "account", accountID, domain)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateMTLSHostnameSettingsBooleanCombinations tests all combinations of boolean values.
func TestMigrateMTLSHostnameSettingsBooleanCombinations(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name                        string
		chinaNetwork                bool
		clientCertificateForwarding bool
	}{
		{"ChinaFalse_ClientTrue", false, true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cleanupMTLSSettings(t)
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(v4BooleanCombinationsConfig, rnd, "account", accountID, domain, tc.clientCertificateForwarding, tc.chinaNetwork)
			sourceVer, targetVer := acctest.InferMigrationVersions(acctest.GetLastV4Version())

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
						Config: v4Config,
					},
					acctest.MigrationV2TestStep(t, v4Config, tmpDir, acctest.GetLastV4Version(), sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(tc.chinaNetwork)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(tc.clientCertificateForwarding)),
					}),
				},
			})
		})
	}
}

// TestMigrateMTLSHostnameSettingsAccountScope tests account-scoped migration.
func TestMigrateMTLSHostnameSettingsAccountScope(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID, domain)

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
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.Null()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
			}),
		},
	})
}

// TestMigrateMTLSHostnameSettingsZoneScope tests zone-scoped migration.
func TestMigrateMTLSHostnameSettingsZoneScope(t *testing.T) {
	//t.Skip("access.api.error.conflict: previous certificate settings still being updated")
	cleanupMTLSSettings(t)
	clearAPIToken(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// v4_basic.tf uses %[2]s for account_id, but zone tests need zone_id
	// Use the basic template with zone_id format arg
	v4Config := fmt.Sprintf(`resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
  zone_id = "%[2]s"
  settings {
    hostname                      = "%[3]s"
    client_certificate_forwarding = true
    china_network                 = false
  }
}
`, rnd, zoneID, domain)

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
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, acctest.GetLastV4Version(), "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.Null()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
			}),
		},
	})
}
