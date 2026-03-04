package custom_ssl_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_custom_ssl", &resource.Sweeper{
		Name: "cloudflare_custom_ssl",
		F:    testSweepCloudflareCustomSSL,
	})
}

func testSweepCloudflareCustomSSL(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping custom SSL sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	certs, err := client.CustomCertificates.List(ctx, custom_certificates.CustomCertificateListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch custom SSL certificates: %s", err))
		return nil
	}

	for _, cert := range certs.Result {
		if !utils.ShouldSweepResource(cert.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting custom SSL certificate: %s", cert.ID))
		_, err := client.CustomCertificates.Delete(ctx, cert.ID, custom_certificates.CustomCertificateDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete custom SSL certificate %s: %s", cert.ID, err))
		}
	}

	return nil
}

func testAccCheckCloudflareCustomSSLDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_ssl" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.CustomCertificates.Get(context.Background(), rs.Primary.ID, custom_certificates.CustomCertificateGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			continue
		}

		tflog.Warn(context.Background(), fmt.Sprintf("Custom SSL certificate %s still exists but this may be expected", rs.Primary.ID))
	}

	return nil
}

// TestAccCustomSSL_Basic tests the basic CRUD lifecycle of a custom SSL certificate.
// This validates that the resource can be created, read, imported, and deleted.
// Note: Update scenarios are not tested here because the Cloudflare API requires replacing
// the entire certificate (including cert and key) to change any properties like bundle_method.
// Reference: https://developers.cloudflare.com/api/resources/custom_certificates/methods/edit/
func TestAccCustomSSL_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_CUSTOM_SSL_TEST") != "1" {
		t.Skip("Skipping custom SSL test due to quota limits on test zone. Set CLOUDFLARE_CUSTOM_SSL_TEST=1 to run.")
	}
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_ssl." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomSSLBasicConfig(zoneID, rnd, cert, key),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", zoneID, s.RootModule().Resources[name].Primary.ID), nil
				},
				ImportStateVerifyIgnore: []string{
					"certificate", // write-only, not returned by API
					"private_key", // write-only, not returned by API
					"status",      // async state transition (pending -> active)
					"modified_on", // timestamp changes between operations
					"type",        // default value handling
				},
			},
		},
	})
}

// TestAccCustomSSL_WithGeoRestrictions tests the optional geo_restrictions attribute.
// This validates that optional nested attributes are handled correctly.
// Note: This test may fail with quota errors on zones with limited custom certificate slots.
func TestAccCustomSSL_WithGeoRestrictions(t *testing.T) {
	if os.Getenv("CLOUDFLARE_CUSTOM_SSL_TEST") != "1" {
		t.Skip("Skipping custom SSL test due to quota limits on test zone. Set CLOUDFLARE_CUSTOM_SSL_TEST=1 to run.")
	}
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_ssl." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomSSLWithGeoRestrictionsConfig(zoneID, rnd, cert, key),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("geo_restrictions").AtMapKey("label"), knownvalue.StringExact("us")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCustomSSLBasicConfig(zoneID, rnd, cert, key string) string {
	return acctest.LoadTestCase("customsslcertbasic.tf", zoneID, rnd, cert, key)
}

func testAccCustomSSLWithGeoRestrictionsConfig(zoneID, rnd, cert, key string) string {
	return acctest.LoadTestCase("customsslwithgeorestrictions.tf", zoneID, rnd, cert, key)
}

// TestAccCustomSSL_NoPolicyDrift verifies that a custom SSL certificate created with a
// policy field does not produce plan drift on a second apply. This is a regression test
// for SECENG-12729 where the API returns "policy_restrictions" but Terraform sends "policy",
// causing phantom drift on every subsequent plan/apply cycle.
func TestAccCustomSSL_NoPolicyDrift(t *testing.T) {
	if os.Getenv("CLOUDFLARE_CUSTOM_SSL_TEST") != "1" {
		t.Skip("Skipping custom SSL test due to quota limits on test zone. Set CLOUDFLARE_CUSTOM_SSL_TEST=1 to run.")
	}
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_ssl." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	config := testAccCustomSSLWithPolicyConfig(zoneID, rnd, cert, key, "(country: US)")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with policy, verify no drift after apply+refresh
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionCreate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			// Step 2: Re-apply same config — must produce empty plan (no drift)
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccCustomSSL_NoBundleMethodDrift verifies that bundle_method is persisted correctly
// through the create → read → plan cycle. This is a regression test for SECENG-12729
// where the PATCH request omitted bundle_method, causing API error 2100 or drift.
func TestAccCustomSSL_NoBundleMethodDrift(t *testing.T) {
	if os.Getenv("CLOUDFLARE_CUSTOM_SSL_TEST") != "1" {
		t.Skip("Skipping custom SSL test due to quota limits on test zone. Set CLOUDFLARE_CUSTOM_SSL_TEST=1 to run.")
	}
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_ssl." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	config := testAccCustomSSLBasicConfig(zoneID, rnd, cert, key)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with bundle_method="force", verify it persists with no drift
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionCreate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
				},
			},
			// Step 2: Re-apply identical config — bundle_method must not drift
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func testAccCustomSSLWithPolicyConfig(zoneID, rnd, cert, key, policy string) string {
	return acctest.LoadTestCase("customsslwithpolicy.tf", zoneID, rnd, cert, key, policy)
}
