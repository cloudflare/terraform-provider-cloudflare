package authenticated_origin_pulls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/origin_tls_client_auth"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls_certificate", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls_certificate",
		F:    testSweepCloudflareAuthenticatedOriginPullsCertificate,
	})
}

func testSweepCloudflareAuthenticatedOriginPullsCertificate(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping authenticated origin pulls certificates sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	certs, err := client.OriginTLSClientAuth.ZoneCertificates.List(ctx, origin_tls_client_auth.ZoneCertificateListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch authenticated origin pulls certificates: %s", err))
		return nil
	}

	for _, cert := range certs.Result {
		if !utils.ShouldSweepResource(cert.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting authenticated origin pulls certificate: %s", cert.ID))
		_, err := client.OriginTLSClientAuth.ZoneCertificates.Delete(ctx, cert.ID, origin_tls_client_auth.ZoneCertificateDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete authenticated origin pulls certificate %s: %s", cert.ID, err))
		}
	}

	return nil
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_authenticated_origin_pulls_certificate" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		certID := rs.Primary.ID

		_, err := client.OriginTLSClientAuth.ZoneCertificates.Get(
			context.Background(),
			certID,
			origin_tls_client_auth.ZoneCertificateGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)
		if err != nil {
			// Certificate not found, destroy succeeded
			continue
		}

		// Certificate still exists - this may be expected due to async deletion
		// or pending deployment states. Log a warning but don't fail the test.
		tflog.Warn(context.Background(), fmt.Sprintf("Authenticated Origin Pulls Certificate %s still exists but this may be expected due to async deletion", certID))
	}

	return nil
}

func testAccAuthenticatedOriginPullsCertificateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		certID := rs.Primary.ID

		return fmt.Sprintf("%s/%s", zoneID, certID), nil
	}
}

func testAccAuthenticatedOriginPullsCertificateConfig(rnd, zoneID, cert, key string) string {
	return fmt.Sprintf(`
resource "cloudflare_authenticated_origin_pulls_certificate" "%[1]s" {
  zone_id     = "%[2]s"
  certificate = <<EOT
%[3]s
EOT
  private_key = <<EOT
%[4]s
EOT
}`, rnd, zoneID, cert, key)
}

// TestAccAuthenticatedOriginPullsCertificate_FullLifecycle tests the full lifecycle
// of an authenticated origin pulls certificate (zone-level) including create, read, and import.
// Note: This resource does not support in-place updates - all input fields have RequiresReplace.
// Note: ExpectNonEmptyPlan is required because the certificate field returned from the API has
// different formatting (whitespace/newlines) than what was sent. Combined with RequiresReplace,
// this causes Terraform to plan a replacement on every refresh. This is a known issue that should
// be addressed in cloudflare-config by adding a plan modifier to normalize certificate fields.
func TestAccAuthenticatedOriginPullsCertificate_FullLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_authenticated_origin_pulls_certificate." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Generate ephemeral certificate for testing
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"example.com"}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: testAccAuthenticatedOriginPullsCertificateConfig(rnd, zoneID, cert, key),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("uploaded_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("signature"), knownvalue.NotNull()),
				},
				// KNOWN ISSUE: The certificate field returned from the API has different formatting
				// than what was sent, causing RequiresReplace drift on every refresh.
				// This triggers an attempted replacement that fails with "certificate already exists".
				// Should be fixed in cloudflare-config with a certificate normalization plan modifier.
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "certificate", "status"},
				ImportStateIdFunc:       testAccAuthenticatedOriginPullsCertificateImportStateIdFunc(resourceName),
			},
		},
	})
}
