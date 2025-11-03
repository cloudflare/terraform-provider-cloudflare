package zero_trust_access_mtls_certificate_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	cf "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)


// Helper functions for loading test configurations
func testAccessMutualTLSCertificateMigrationBasic(rnd string, identifier *cf.ResourceContainer, cert string, domain string) string {
	// Convert literal \n to actual newlines for proper certificate format
	processedCert := fmt.Sprintf("<<EOT\n%s\nEOT", strings.ReplaceAll(cert, "\\n", "\n"))
	return acctest.LoadTestCase("accessmutualtlscertificate_migration_basic.tf", rnd, identifier.Type, identifier.Identifier, processedCert, domain)
}

func testAccessMutualTLSCertificateMigrationZoneScoped(rnd string, zoneID string, cert string) string {
	processedCert := fmt.Sprintf("<<EOT\n%s\nEOT", strings.ReplaceAll(cert, "\\n", "\n"))
	return acctest.LoadTestCase("accessmutualtlscertificate_migration_zone_scoped.tf", rnd, zoneID, processedCert)
}


// TestMigrateZeroTrustAccessMTLSCertificate_Basic tests basic migration from v4 to v5
// The test starts with v4 resource name (cloudflare_access_mutual_tls_certificate) and
// the migration tool renames it to v5 (cloudflare_zero_trust_access_mtls_certificate)
func TestMigrateZeroTrustAccessMTLSCertificate_Basic(t *testing.T) {
	t.Skip(`Skipping due to consistent conflicts: "message": "access.api.error.conflict: certificate has active associations"`)
	waitForCertificateCleanup(t, false)
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd
	tmpDir := t.TempDir()

	testCert, err := generateUniqueTestCertificate(fmt.Sprintf("basic-%s", rnd))
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	identifier := &cf.ResourceContainer{
		Type:       "account",
		Identifier: accountID,
	}
	v4Config := testAccessMutualTLSCertificateMigrationBasic(rnd, identifier, testCert, domain)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.SetSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
				// New computed field in v5
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSCertificate_ZoneScoped tests zone-scoped resource migration
func TestMigrateZeroTrustAccessMTLSCertificate_ZoneScoped(t *testing.T) {
	waitForCertificateCleanup(t, true)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd
	tmpDir := t.TempDir()

	testCert, err := generateUniqueTestCertificate(fmt.Sprintf("zone-%s", rnd))
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	v4Config := testAccessMutualTLSCertificateMigrationZoneScoped(rnd, zoneID, testCert)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
				// Note: associated_hostnames might be nil or empty set in zone-scoped resources
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
			}),
		},
	})
}
