package zero_trust_access_mtls_certificate_test

import (
	"fmt"
	"os"
	"testing"
	"time"

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

// waitBetweenTests adds a delay to prevent API conflicts between tests
func waitBetweenTests(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_LOG") == "DEBUG" {
		t.Logf("Waiting 15 seconds to prevent API conflicts...")
	}
	time.Sleep(15 * time.Second)
}

// TestMigrateZeroTrustAccessMTLSCertificate_Basic tests basic migration from v4 to v5
func TestMigrateZeroTrustAccessMTLSCertificate_Basic(t *testing.T) {
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
	resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd
	tmpDir := t.TempDir()

	// V4 config using old resource name and hardcoded certificate
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
  name                 = "%[1]s"
  account_id           = "%[2]s"
  certificate          = <<EOT
%[4]s
EOT
  associated_hostnames = ["%[1]s.terraform.%[3]s", "%[1]ss.terraform.%[3]s"]
}`, rnd, accountID, domain, testCertificate)

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
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
				// New computed field in v5
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessMTLSCertificate_Minimal tests minimal configuration migration
func TestMigrateZeroTrustAccessMTLSCertificate_Minimal(t *testing.T) {
	waitBetweenTests(t)
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd
	tmpDir := t.TempDir()

	// V4 config with only required fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
  name        = "%[1]s"
  account_id  = "%[2]s"
  certificate = <<EOT
%[3]s
EOT
}`, rnd, accountID, testCertificate)

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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.SetSizeExact(0)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
			}),
		},
	})
}