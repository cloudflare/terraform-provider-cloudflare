package zero_trust_access_mtls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	cf "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

const testCertificate = `-----BEGIN CERTIFICATE-----
MIIDwTCCAqmgAwIBAgIURXiQAGaonddgViImcT1C433iszYwDQYJKoZIhvcNAQEL
BQAwcDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJh
bmNpc2NvMRcwFQYDVQQKDA5NaWdyYXRpb24gVGVzdDEjMCEGA1UEAwwabWlncmF0
aW9uLXRlc3QuZXhhbXBsZS5jb20wHhcNMjUwODIxMTIzNDUzWhcNMjYwODIxMTIz
NDUzWjBwMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBG
cmFuY2lzY28xFzAVBgNVBAoMDk1pZ3JhdGlvbiBUZXN0MSMwIQYDVQQDDBptaWdy
YXRpb24tdGVzdC5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAKBp1D0iXRxCQ09OLnyQDLXNsaCedYXdbThsMSqyTh3qtkM/sAokCGKP
/Mlo9DIXWl4hy/qtNAmXPd7UKF1u5V9aQpuVYHN2cNue9Bjzvdqxux8ii7zN+bGK
3U/gaaFqEWpyQJjKlz9t9H7gUtPcNQErz093F2Cid8uh7JJU2Hml8AyW9HiVKYmS
lWnUE0YABvl1HIos6+KJrMX2ZBvjkF8OvCmE6JkGnV1nH5o7BQ3xQc/s687JilIy
IJHyW5YmDBZ40N3bVrVb2PtTw6b7VROxb+OVE8KORKJ6ITLDEsR1lEIdA6/ECZc9
x0nUTXC3rZksL1VRrTVBHKTL3BTl0e8CAwEAAaNTMFEwHQYDVR0OBBYEFEKStRM/
Wn1H/2weP9+cuZKAwTopMB8GA1UdIwQYMBaAFEKStRM/Wn1H/2weP9+cuZKAwTop
MA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAHcMU1ZwiOjDhdO7
j9QA6zDNvikKmzRdi209VF2OYr4lLNuiXtoziVuhsbXd1ga24UuDhRbasrJmFKFO
V2IyQljeGwCS7TjllHS2hkTZvAWyPjXVPS6fhWAoV1B6I5CoeaJpeaLqlObXDKn6
O8Qm3kagbBtquCqlZfTBVMH6Jg4yPNAXSqqtWsm8imKk+p67DTGMP9Di3m3YZG32
Tl2SqcL0tzrQkkuxj6k0QilowkLYtOPwE8gvSv1YbDfPVyhLaNTSzIsJU+gCsnCE
qAY7EfWG4Vg2shFGLErpkI8/4S2F3ddYMObqwI0w4sIfLqXfTSuPWkhi1AvN/rzq
Et51cI4=
-----END CERTIFICATE-----`

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

// waitBetweenTests adds a delay to prevent API conflicts between tests
func waitBetweenTests(t *testing.T, isZone bool) {
	t.Helper()
	c := cloudflare.NewClient()
	retry := 0
	listParams := zero_trust.AccessCertificateListParams{}
	if isZone {
		listParams.ZoneID = cloudflare.F(os.Getenv("CLOUDFLARE_ZONE_ID"))
	} else {
		listParams.AccountID = cloudflare.F(os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
	}
	for retry < 5 {
		res, err := c.ZeroTrust.Access.Certificates.List(context.Background(), listParams)
		if err != nil {
			retry++
			continue
		}
		if len(res.Result) == 0 {
			return
		}
		time.Sleep(3 * time.Second)
		retry++
		if os.Getenv("TF_LOG") == "DEBUG" {
			t.Logf("Waiting for list to return empty results to prevent API conflicts. Retry number: %d", retry)
		}
	}

}

// TestMigrateZeroTrustAccessMTLSCertificate_Basic tests basic migration from v4 to v5
// The test starts with v4 resource name (cloudflare_access_mutual_tls_certificate) and
// the migration tool renames it to v5 (cloudflare_zero_trust_access_mtls_certificate)
func TestMigrateZeroTrustAccessMTLSCertificate_Basic(t *testing.T) {
	waitBetweenTests(t, false)
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

	identifier := &cf.ResourceContainer{
		Type:       "account",
		Identifier: accountID,
	}
	v4Config := testAccessMutualTLSCertificateMigrationBasic(rnd, identifier, testCertificate, domain)

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
	waitBetweenTests(t, false)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd
	tmpDir := t.TempDir()

	v4Config := testAccessMutualTLSCertificateMigrationZoneScoped(rnd, zoneID, testCertificate)

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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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
