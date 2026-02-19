package v500_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	_ "embed"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_zone_scoped.tf
var v4ZoneScopedConfig string

//go:embed testdata/v5_zone_scoped.tf
var v5ZoneScopedConfig string

// clearAPIToken unsets CLOUDFLARE_API_TOKEN for Access tests (require API_KEY + EMAIL).
func clearAPIToken(t *testing.T) {
	t.Helper()
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		t.Cleanup(func() { os.Setenv("CLOUDFLARE_API_TOKEN", originalToken) })
	}
}

func generateTestCertificate(testName string) (string, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:      []string{"US"},
			Province:     []string{"CA"},
			Locality:     []string{"San Francisco"},
			Organization: []string{fmt.Sprintf("Test-%s", testName)},
			CommonName:   fmt.Sprintf("%s-test.example.com", testName),
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		return "", err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	return string(certPEM), nil
}

func waitForCertificateCleanup(t *testing.T, isZone bool) {
	t.Helper()

	if os.Getenv("TF_ACC") == "" {
		return
	}

	client, err := acctest.SharedV1Client()
	if err != nil {
		t.Fatalf("Failed to create Cloudflare client: %s", err)
	}

	maxWait := 30 * time.Second
	pollInterval := 2 * time.Second
	deadline := time.Now().Add(maxWait)

	for time.Now().Before(deadline) {
		var certs []cloudflare.AccessMutualTLSCertificate
		var listErr error

		if isZone {
			certs, _, listErr = client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.ZoneIdentifier(os.Getenv("CLOUDFLARE_ZONE_ID")), cloudflare.ListAccessMutualTLSCertificatesParams{})
		} else {
			certs, _, listErr = client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.ListAccessMutualTLSCertificatesParams{})
		}

		if listErr != nil {
			time.Sleep(pollInterval)
			continue
		}

		if len(certs) == 0 {
			return
		}

		time.Sleep(pollInterval)
	}
}

// TestMigrateAccessMTLSCertificateBasic tests basic migration with account_id.
func TestMigrateAccessMTLSCertificateBasic(t *testing.T) {
	waitForCertificateCleanup(t, false)
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, cert string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, cert string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, name, cert)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, cert string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, name, cert)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-mtls-%s", rnd)
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd

			testCert, err := generateTestCertificate(fmt.Sprintf("basic-%s", rnd))
			if err != nil {
				t.Fatalf("Failed to generate test certificate: %v", err)
			}

			testConfig := tc.configFn(rnd, accountID, name, testCert)
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateAccessMTLSCertificateZoneScoped tests zone-scoped resource migration.
func TestMigrateAccessMTLSCertificateZoneScoped(t *testing.T) {
	waitForCertificateCleanup(t, true)
	clearAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name, cert string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, name, cert string) string {
				return fmt.Sprintf(v4ZoneScopedConfig, rnd, zoneID, name, cert)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, name, cert string) string {
				return fmt.Sprintf(v5ZoneScopedConfig, rnd, zoneID, name, cert)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-mtls-zone-%s", rnd)
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_mtls_certificate." + rnd

			testCert, err := generateTestCertificate(fmt.Sprintf("zone-%s", rnd))
			if err != nil {
				t.Fatalf("Failed to generate test certificate: %v", err)
			}

			testConfig := tc.configFn(rnd, zoneID, name, testCert)
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}
