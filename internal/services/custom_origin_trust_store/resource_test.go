package custom_origin_trust_store_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/acm"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func init() {
	resource.AddTestSweepers("cloudflare_custom_origin_trust_store", &resource.Sweeper{
		Name: "cloudflare_custom_origin_trust_store",
		F:    testSweepCloudflareCustomOriginTrustStores,
	})
}

func testSweepCloudflareCustomOriginTrustStores(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if zoneID == "" {
		tflog.Info(ctx, "CLOUDFLARE_ZONE_ID not set, skipping custom origin trust store sweep")
		return nil
	}

	certs, err := client.ACM.CustomTrustStore.List(ctx, acm.CustomTrustStoreListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list custom origin trust store certificates: %s", err))
		return nil
	}

	for _, cert := range certs.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting custom origin trust store certificate: %s", cert.ID))
		_, err := client.ACM.CustomTrustStore.Delete(ctx, cert.ID, acm.CustomTrustStoreDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete custom origin trust store certificate %s: %s", cert.ID, err))
		}
	}

	return nil
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccCloudflareCustomOriginTrustStore_FullLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_custom_origin_trust_store.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	cert, err := generateSelfSignedCert(rnd)
	if err != nil {
		t.Fatalf("failed to generate self-signed certificate: %v", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomOriginTrustStoreDestroy,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCloudflareCustomOriginTrustStoreConfig(rnd, zoneID, cert),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("issuer"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("signature"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("uploaded_on"), knownvalue.NotNull()),
				},
			},
			// Re-apply same config to verify no drift
			{
				Config: testAccCloudflareCustomOriginTrustStoreConfig(rnd, zoneID, cert),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateIdFunc: testAccCloudflareCustomOriginTrustStoreImportStateIdFunc(name),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareCustomOriginTrustStoreDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_origin_trust_store" {
			continue
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		certID := rs.Primary.ID

		cert, err := client.ACM.CustomTrustStore.Get(
			context.Background(),
			certID,
			acm.CustomTrustStoreGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)

		// If GET fails (404 or other error), consider it destroyed
		if err != nil {
			continue
		}

		// If we can still retrieve it, verify it's in a deleted state
		status := cert.Status
		if status != "deleted" && status != "pending_deletion" {
			return fmt.Errorf("custom origin trust store certificate %s still exists with status %s, expected deleted or pending_deletion after destroy", certID, string(status))
		}
	}

	return nil
}

func testAccCloudflareCustomOriginTrustStoreImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		certID := rs.Primary.ID

		return fmt.Sprintf("%s/%s", zoneID, certID), nil
	}
}

func generateSelfSignedCert(testName string) (string, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
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
		return "", fmt.Errorf("failed to create certificate: %w", err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	return string(certPEM), nil
}

func testAccCloudflareCustomOriginTrustStoreConfig(rnd, zoneID, cert string) string {
	return acctest.LoadTestCase("customorigintruststorelifecycle.tf", rnd, zoneID, cert)
}

// TestAccCloudflareCustomOriginTrustStore_NoTrailingNewline tests that
// the provider properly handles certificates without trailing newlines and doesn't
// detect drift. This verifies the normalization logic for the certificate field.
func TestAccCloudflareCustomOriginTrustStore_NoTrailingNewline(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_custom_origin_trust_store.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	cert, err := generateSelfSignedCert(rnd)
	if err != nil {
		t.Fatalf("failed to generate self-signed certificate: %v", err)
	}

	// Strip trailing newline to test drift behavior
	certNoTrailingNewline := strings.TrimRight(cert, "\n")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomOriginTrustStoreDestroy,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCloudflareCustomOriginTrustStoreNoTrailingNewlineConfig(rnd, zoneID, certNoTrailingNewline),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			// Re-apply same config to verify no drift
			{
				Config: testAccCloudflareCustomOriginTrustStoreNoTrailingNewlineConfig(rnd, zoneID, certNoTrailingNewline),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Mocks hardcoded certificates that end without trailing new line /n
func testAccCloudflareCustomOriginTrustStoreNoTrailingNewlineConfig(rnd, zoneID, cert string) string {
	escapedCert := strings.ReplaceAll(cert, "\n", `\n`)
	return fmt.Sprintf(`
resource "cloudflare_custom_origin_trust_store" "%s" {
  zone_id     = "%s"
  certificate = "%s"
}
`, rnd, zoneID, escapedCert)
}
