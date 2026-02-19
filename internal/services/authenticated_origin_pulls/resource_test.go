package authenticated_origin_pulls_test

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
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls",
		F:    testSweepCloudflareAuthenticatedOriginPulls,
	})
}

func testSweepCloudflareAuthenticatedOriginPulls(r string) error {
	ctx := context.Background()
	// Authenticated Origin Pulls hostname associations are voided by setting enabled: null.
	// The test framework handles cleanup via CheckDestroy. Sweeping is informational only.
	tflog.Info(ctx, "Authenticated Origin Pulls hostname associations are cleaned up via CheckDestroy")
	return nil
}

func testAccCheckCloudflareAuthenticatedOriginPullsDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_authenticated_origin_pulls" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		hostname := rs.Primary.Attributes["hostname"]

		if hostname == "" {
			continue
		}

		result, err := client.OriginTLSClientAuth.Hostnames.Get(
			context.Background(),
			hostname,
			origin_tls_client_auth.HostnameGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)
		if err != nil {
			// Hostname association not found, destroy succeeded
			continue
		}

		// Check if the association was voided (enabled should be false or association removed)
		// After deletion, the API may return the hostname with enabled=false or not return it at all
		if result.Enabled {
			return fmt.Errorf("authenticated origin pulls hostname association %s still enabled", hostname)
		}
	}

	return nil
}

func testAccAuthenticatedOriginPullsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		hostname := rs.Primary.Attributes["hostname"]

		return fmt.Sprintf("%s/%s", zoneID, hostname), nil
	}
}

func testAccAuthenticatedOriginPullsConfig(rnd, zoneID, cert, key, hostname string, enabled bool) string {
	return fmt.Sprintf(`
resource "cloudflare_authenticated_origin_pulls_hostname_certificate" "%[1]s" {
  zone_id     = "%[2]s"
  certificate = <<EOT
%[3]s
EOT
  private_key = <<EOT
%[4]s
EOT
}

resource "cloudflare_authenticated_origin_pulls" "%[1]s" {
  zone_id = "%[2]s"
  config = [{
    cert_id  = cloudflare_authenticated_origin_pulls_hostname_certificate.%[1]s.id
    enabled  = %[6]t
    hostname = "%[5]s"
  }]
}`, rnd, zoneID, cert, key, hostname, enabled)
}

// TestAccAuthenticatedOriginPulls_FullLifecycle tests the full lifecycle of an
// authenticated origin pulls hostname association including create, update, read, and import.
//
// This resource associates a hostname certificate with a specific hostname to enable
// authenticated origin pulls for that hostname. The API endpoint
// PUT /zones/{zone_id}/origin_tls_client_auth/hostnames returns an array of hostname
// associations, which is handled by custom array response processing.
//
// Lifecycle:
// 1. Create: Associate certificate with hostname (enabled=true)
// 2. Update: Change enabled from true to false
// 3. Import: Verify import works with zone_id/hostname format
// 4. Delete: Voiding the association (enabled=null) happens automatically during destroy
func TestAccAuthenticatedOriginPulls_FullLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_authenticated_origin_pulls." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// Generate ephemeral certificate for testing
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAuthenticatedOriginPullsDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with enabled = true
			{
				Config: testAccAuthenticatedOriginPullsConfig(rnd, zoneID, cert, key, domain, true),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cert_id"), knownvalue.NotNull()),
				},
			},
			// Step 2: Update - change enabled to false
			{
				Config: testAccAuthenticatedOriginPullsConfig(rnd, zoneID, cert, key, domain, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				},
			},
			// Step 3: Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config", "created_at", "updated_at", "status"},
				ImportStateIdFunc:       testAccAuthenticatedOriginPullsImportStateIdFunc(resourceName),
			},
		},
	})
}
