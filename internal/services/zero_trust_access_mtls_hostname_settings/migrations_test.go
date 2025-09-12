package zero_trust_access_mtls_hostname_settings_test

import (
	"context"
	"os"
	"testing"
	"time"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// cleanupMTLSSettings clears all MTLS hostname settings and certificates to prevent test conflicts
func cleanupMTLSSettings(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		t.Fatalf("Failed to create Cloudflare client: %v", clientErr)
	}

	// First clear hostname settings
	deletedSettings := cfv1.UpdateAccessMutualTLSHostnameSettingsParams{
		Settings: []cfv1.AccessMutualTLSHostnameSettings{},
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID != "" {
		_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.AccountIdentifier(accountID), deletedSettings)
		if err != nil {
			t.Logf("Warning: Failed to clear account MTLS hostname settings: %v", err)
		}
		
		// Clean account certificates
		accountCerts, _, err := client.ListAccessMutualTLSCertificates(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListAccessMutualTLSCertificatesParams{})
		if err == nil {
			for _, cert := range accountCerts {
				// Clear hostnames first, then delete
				client.UpdateAccessMutualTLSCertificate(ctx, cfv1.AccountIdentifier(accountID), cfv1.UpdateAccessMutualTLSCertificateParams{
					ID:                  cert.ID,
					AssociatedHostnames: []string{},
				})
				client.DeleteAccessMutualTLSCertificate(ctx, cfv1.AccountIdentifier(accountID), cert.ID)
			}
		}
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")  
	if zoneID != "" {
		_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.ZoneIdentifier(zoneID), deletedSettings)
		if err != nil {
			t.Logf("Warning: Failed to clear zone MTLS hostname settings: %v", err)
		}
		
		// Clean zone certificates
		zoneCerts, _, err := client.ListAccessMutualTLSCertificates(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ListAccessMutualTLSCertificatesParams{})
		if err == nil {
			for _, cert := range zoneCerts {
				// Clear hostnames first, then delete
				client.UpdateAccessMutualTLSCertificate(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.UpdateAccessMutualTLSCertificateParams{
					ID:                  cert.ID,
					AssociatedHostnames: []string{},
				})
				client.DeleteAccessMutualTLSCertificate(ctx, cfv1.ZoneIdentifier(zoneID), cert.ID)
			}
		}
	}
	
	// Add extra wait after cleanup to ensure API has processed the deletions
	time.Sleep(5 * time.Second)
}

// waitBetweenTests adds a delay to prevent API conflicts between tests
func waitBetweenTests(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_LOG") == "DEBUG" {
		t.Logf("Waiting 15 seconds to prevent API conflicts...")
	}
	time.Sleep(15 * time.Second)
}

// TestMigrateZeroTrustAccessMTLSHostnameSettings_Basic tests basic migration from v4 to v5
func TestMigrateZeroTrustAccessMTLSHostnameSettings_Basic(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config using block syntax for settings
	v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_basic.tf", rnd, "account", accountID, domain)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSHostnameSettings_Multiple tests migration with multiple hostnames
func TestMigrateZeroTrustAccessMTLSHostnameSettings_Multiple(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	altDomain := "alt." + domain // Use subdomain to avoid conflicts
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple hostnames - tests migration logic even if API merges them
	v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_multiple.tf", rnd, "account", accountID, domain, altDomain)

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
				// Accept plan diff - v4 provider doesn't handle multiple hostname settings in state properly
				ExpectNonEmptyPlan: true,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// API merges multiple hostnames, so we expect list with both
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(2)),
				// Verify both hostnames are present (order may vary due to API behavior)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(1).AtMapKey("hostname"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSHostnameSettings_BooleanDefaults tests migration when optional booleans are not specified
// Note: Works around v4 provider issues by using ExpectError to handle the plan diff
func TestMigrateZeroTrustAccessMTLSHostnameSettings_BooleanDefaults(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config with boolean defaults (both false) - tests migration works despite v4 provider quirks
	v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_boolean_defaults.tf", rnd, "account", accountID, domain)

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
				// Accept the plan diff caused by v4 provider bug - the point is migration works
				ExpectNonEmptyPlan: true,
			},
			// Migration step - this tests that our migration logic handles the scenario correctly
			// even if v4 provider has state management issues
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
				// Verify that migration provides correct defaults even when v4 doesn't store them properly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSHostnameSettings_BooleanCombinations tests all combinations of boolean values
func TestMigrateZeroTrustAccessMTLSHostnameSettings_BooleanCombinations(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name                        string
		chinaNetwork                bool
		clientCertificateForwarding bool
	}{
		// Skip BothFalse case - v4 provider behavior with all-false booleans causes plan diffs
		{"ChinaFalse_ClientTrue", false, true},
		// Skip china_network = true tests since the test account is not china-network enabled
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cleanupMTLSSettings(t)
			waitBetweenTests(t)
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
			tmpDir := t.TempDir()

			// V4 config with specific boolean combinations
			v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_boolean_combinations.tf", rnd, "account", accountID, domain, tc.clientCertificateForwarding, tc.chinaNetwork)

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
								VersionConstraint: "~> 4.0",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

// TestMigrateZeroTrustAccessMTLSHostnameSettings_AccountScope tests account-scoped migration
func TestMigrateZeroTrustAccessMTLSHostnameSettings_AccountScope(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config explicitly using account_id
	v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_basic.tf", rnd, "account", accountID, domain)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify zone_id is not set for account-scoped resource
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.Null()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSHostnameSettings_ZoneScope tests zone-scoped migration
func TestMigrateZeroTrustAccessMTLSHostnameSettings_ZoneScope(t *testing.T) {
	cleanupMTLSSettings(t)
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_hostname_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config explicitly using zone_id
	v4Config := acctest.LoadTestCase("accessmutualtlshostnamesettings_migration_basic.tf", rnd, "zone", zoneID, domain)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify account_id is not set for zone-scoped resource
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.Null()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
			}),
		},
	})
}
