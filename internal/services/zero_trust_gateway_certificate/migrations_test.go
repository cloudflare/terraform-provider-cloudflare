package zero_trust_gateway_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// testAccCheckCloudflareZeroTrustGatewayCertificateDestroy verifies the certificate is destroyed
func testAccCheckCloudflareZeroTrustGatewayCertificateDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_gateway_certificate" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		certID := rs.Primary.ID

		_, err := client.ZeroTrust.Gateway.Certificates.Get(context.Background(), certID, zero_trust.GatewayCertificateGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("zero trust gateway certificate still exists: %s", certID)
		}
	}

	return nil
}

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_BasicWithoutValidity tests basic migration
// when validity_period_days is NOT explicitly set in the v4 config (uses v4 default).
// This ensures the migration removes validity_period_days from state since v5 uses no_refresh.
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_BasicWithoutValidity(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_certificate" "%[1]s" {
  account_id      = "%[2]s"
  gateway_managed = true
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareZeroTrustGatewayCertificateDestroy,
		WorkingDir:   tmpDir,
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
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify v4-only fields removed from config
				// Verify computed fields are present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("binding_status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithActivate tests migration
// with the activate field set. This ensures boolean fields are preserved correctly.
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithActivate(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_certificate" "%[1]s" {
  account_id      = "%[2]s"
  gateway_managed = true
  activate        = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareZeroTrustGatewayCertificateDestroy,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify activate field is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("activate"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithExplicitValidity tests migration
// when validity_period_days IS explicitly set in the v4 config.
// This ensures: 1) the value is preserved in v5 state, and 2) Int -> Int64 type conversion works.
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithExplicitValidity(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
	tmpDir := t.TempDir()

	// Test with explicit validity_period_days (3650 days = 10 years)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_certificate" "%[1]s" {
  account_id           = "%[2]s"
  gateway_managed      = true
  validity_period_days = 3650
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareZeroTrustGatewayCertificateDestroy,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify validity_period_days is preserved and converted from Int to Int64 (stored as float64 in JSON)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_period_days"), knownvalue.Float64Exact(3650)),
			}),
		},
	})
}

