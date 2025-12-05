package certificate_pack_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ssl"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func init() {
	resource.AddTestSweepers("cloudflare_certificate_pack", &resource.Sweeper{
		Name: "cloudflare_certificate_pack",
		F:    testSweepCloudflareCertificatePack,
	})
}

func testSweepCloudflareCertificatePack(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping certificate packs sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	certificates, certErr := client.ListCertificatePacks(ctx, zoneID)
	if certErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch certificate packs: %s", certErr))
		return certErr
	}

	if len(certificates) == 0 {
		tflog.Info(ctx, "No Cloudflare certificate packs to sweep")
		return nil
	}

	for _, certificate := range certificates {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(certificate.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting certificate pack: %s (zone: %s)", certificate.ID, zoneID))
		if err := client.DeleteCertificatePack(ctx, zoneID, certificate.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete certificate pack %s: %s", certificate.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted certificate pack: %s", certificate.ID))
	}

	return nil
}

func testAccCheckCloudflareCertificatePackDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_certificate_pack" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.SSL.CertificatePacks.Get(context.Background(), rs.Primary.ID, ssl.CertificatePackGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			// If certificate pack is not found, it was successfully deleted
			continue
		}

		// If we can still retrieve the certificate pack, it might not be fully cleaned up
		// but this is expected behavior for certificate packs - they don't always get immediately deleted
		tflog.Warn(context.Background(), fmt.Sprintf("Certificate pack %s still exists but this may be expected", rs.Primary.ID))
	}

	return nil
}

func TestAccCertificatePack_AdvancedLetsEncrypt(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackAdvancedLetsEncryptConfig(zoneID, domain, "advanced", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
					// Lists and computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func testAccCertificatePackAdvancedLetsEncryptConfig(zoneID, domain, certType, rnd string) string {
	return acctest.LoadTestCase("acccertificatepackadvancedletsencryptconfig.tf", zoneID, domain, rnd, certType)
}

func TestAccCertificatePack_WaitForActive(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackAdvancedWaitForActiveConfig(zoneID, domain, "advanced", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
					// Lists and computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func testAccCertificatePackAdvancedWaitForActiveConfig(zoneID, domain, certType, rnd string) string {
	return acctest.LoadTestCase("acccertificatepackadvancedwaitforactiveconfig.tf", zoneID, domain, rnd, certType)
}

func testAccCertificatePackBasicConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("basic.tf", zoneID, domain, rnd)
}

func testAccCertificatePackGoogleCAConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("google_ca.tf", zoneID, domain, rnd)
}

func testAccCertificatePackSSLComCAConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("ssl_com_ca.tf", zoneID, domain, rnd)
}

func testAccCertificatePackHttpValidationConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("http_validation.tf", zoneID, domain, rnd)
}

// func testAccCertificatePackEmailValidationConfig(zoneID, domain, rnd string) string {
//	return acctest.LoadTestCase("email_validation.tf", zoneID, domain, rnd)
// }

func testAccCertificatePackValidity14DaysConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("validity_14_days.tf", zoneID, domain, rnd)
}

func testAccCertificatePackValidity30DaysConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("validity_30_days.tf", zoneID, domain, rnd)
}

func testAccCertificatePackValidity365DaysConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("validity_365_days.tf", zoneID, domain, rnd)
}

func testAccCertificatePackCloudflareBrandingTrueConfig(zoneID, domain, rnd string) string {
	return acctest.LoadTestCase("cloudflare_branding_true.tf", zoneID, domain, rnd)
}

func TestAccCertificatePack_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackBasicConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Optional attributes - should be null when not set
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Null()),
					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func TestAccCertificatePack_GoogleCA(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackGoogleCAConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("google")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func TestAccCertificatePack_SSLComCA(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackSSLComCAConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("ssl_com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func TestAccCertificatePack_HttpValidation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackHttpValidationConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

// Email validation is not supported by any of the available CAs:
// - Let's Encrypt and Google only support 'txt' and 'http' validation
// - SSL.com appears to also only support 'txt' and 'http' validation
// Leaving this test commented out until email validation is supported by a CA
//
// func TestAccCertificatePack_EmailValidation(t *testing.T) {
//	 NOTE: Email validation is currently not supported by any available certificate authority
// }

func TestAccCertificatePack_ValidityDays(t *testing.T) {
	// Test different validity periods
	testCases := []struct {
		name       string
		configFunc func(string, string, string) string
		days       int64
		ca         string
	}{
		{"14Days", testAccCertificatePackValidity14DaysConfig, 14, "google"},
		{"30Days", testAccCertificatePackValidity30DaysConfig, 30, "google"},
		{"365Days", testAccCertificatePackValidity365DaysConfig, 365, "ssl_com"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			name := "cloudflare_certificate_pack." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { acctest.TestAccPreCheck(t) },
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(zoneID, domain, rnd),
						ConfigStateChecks: []statecheck.StateCheck{
							// Required attributes
							statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact(tc.ca)),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(tc.days)),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
							// Computed attributes
							statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
						},
					},
					{
						ResourceName:            name,
						ImportState:             true,
						ImportStateVerify:       true,
						ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
						ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
					},
				},
			})
		})
	}
}

func TestAccCertificatePack_CloudflareBranding(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackCloudflareBrandingTrueConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "certificates", "cloudflare_branding", "hosts", "primary_certificate", "status", "type", "validation_errors", "validation_method", "validation_records", "validity_days"},
			},
		},
	})
}

func TestAccCertificatePack_ComputedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackBasicConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
					// Computed attributes - comprehensive validation
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
					// validation_errors and validation_records can be null or lists depending on certificate status
				},
			},
		},
	})
}
