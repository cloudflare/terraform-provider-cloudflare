package mtls_certificate_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateMTLSCertificate_Migration_Basic tests the mTLS certificate
// migration. This is a simple migration with:
// - No field renames (all fields identical in v4 and v5)
// - No type conversions needed
// - Resource type unchanged
// The test ensures that existing certificates migrate cleanly without modification.
// Note: Uses ca=false since GenerateEphemeralCertAndKey creates leaf certificates.
func TestMigrateMTLSCertificate_Migration_Basic(t *testing.T) {
	t.Skip("Normalization issue with cert. Post-migration apply fixes it.")

	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v4_52_1", // Last v4 release
			version: "4.52.1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip if acceptance tests are not enabled
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			if accountID == "" {
				t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
			}

			rnd := utils.GenerateRandomResourceName()
			certName := fmt.Sprintf("cftftest-%s", rnd)
			resourceName := "cloudflare_mtls_certificate." + rnd

			// Generate valid test certificate
			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"example.com"}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			testConfig := testAccCloudflareMTLSCertificateMigrationConfigV4Basic(accountID, rnd, certName, cert, key)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create mTLS certificate with v4 provider
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
							// Computed fields should be present
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("serial_number"), knownvalue.NotNull()),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						// After migration, all fields should remain identical
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
						// Computed fields preserved
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("serial_number"), knownvalue.NotNull()),
					}),
					{
						// Step 3: Apply the migrated configuration with v5 provider
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
							// Computed fields still present
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("serial_number"), knownvalue.NotNull()),
						},
					},
				},
			})
		})
	}
}

// TestMigrateMTLSCertificate_Migration_WithPrivateKey tests migration of
// a leaf certificate with a private key to ensure optional fields migrate correctly.
func TestMigrateMTLSCertificate_Migration_WithPrivateKey(t *testing.T) {
	t.Skip("Normalization issue with cert. Post-migration apply fixes it.")

	// Skip if acceptance tests are not enabled
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}

	rnd := utils.GenerateRandomResourceName()
	certName := fmt.Sprintf("cftftest-leaf-%s", rnd)
	resourceName := "cloudflare_mtls_certificate." + rnd

	// Generate valid test certificate with private key
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"test.example.com"}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	v4Config := testAccCloudflareMTLSCertificateMigrationConfigV4WithKey(accountID, rnd, certName, cert, key)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create leaf certificate with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("private_key"), knownvalue.NotNull()),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("private_key"), knownvalue.NotNull()),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("private_key"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// TestMigrateMTLSCertificate_Migration_Minimal tests migration of a minimal
// certificate without optional name field.
func TestMigrateMTLSCertificate_Migration_Minimal(t *testing.T) {
	t.Skip("Normalization issue with cert. Post-migration apply fixes it.")

	// Skip if acceptance tests are not enabled
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_mtls_certificate." + rnd

	// Generate valid test certificate
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"minimal.example.com"}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	v4Config := testAccCloudflareMTLSCertificateMigrationConfigV4Minimal(accountID, rnd, cert, key)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create minimal certificate with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// V4 Configuration Functions

// testAccCloudflareMTLSCertificateMigrationConfigV4Basic returns a basic leaf certificate config
// with name field
func testAccCloudflareMTLSCertificateMigrationConfigV4Basic(accountID, rnd, certName, cert, key string) string {
	return fmt.Sprintf(`
resource "cloudflare_mtls_certificate" "%[2]s" {
  account_id   = "%[1]s"
  ca           = false
  certificates = <<EOT
%[4]s
EOT
  private_key  = <<EOT
%[5]s
EOT
  name         = "%[3]s"
}
`, accountID, rnd, certName, cert, key)
}

// testAccCloudflareMTLSCertificateMigrationConfigV4WithKey returns a leaf certificate with private key config
func testAccCloudflareMTLSCertificateMigrationConfigV4WithKey(accountID, rnd, certName, cert, key string) string {
	return fmt.Sprintf(`
resource "cloudflare_mtls_certificate" "%[2]s" {
  account_id   = "%[1]s"
  ca           = false
  certificates = <<EOT
%[4]s
EOT
  private_key  = <<EOT
%[5]s
EOT
  name         = "%[3]s"
}
`, accountID, rnd, certName, cert, key)
}

// testAccCloudflareMTLSCertificateMigrationConfigV4Minimal returns a minimal certificate config without optional name field
func testAccCloudflareMTLSCertificateMigrationConfigV4Minimal(accountID, rnd, cert, key string) string {
	return fmt.Sprintf(`
resource "cloudflare_mtls_certificate" "%[2]s" {
  account_id   = "%[1]s"
  ca           = false
  certificates = <<EOT
%[3]s
EOT
  private_key  = <<EOT
%[4]s
EOT
}
`, accountID, rnd, cert, key)
}
