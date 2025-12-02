package custom_hostname_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
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
	resource.AddTestSweepers("cloudflare_custom_hostname", &resource.Sweeper{
		Name: "cloudflare_custom_hostname",
		F:    testSweepCloudflareCustomHostnames,
	})
}

func TestAccCloudflareCustomHostname_LetsEncryptCA(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCA(zoneID, rnd, domain, "lets_encrypt"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.certificate_authority", "lets_encrypt"),
				),
			},
		},
	})
}


func TestAccCloudflareCustomHostname_InvalidHostname(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareCustomHostnameInvalidHostname(zoneID, rnd),
				ExpectError: regexp.MustCompile("Invalid custom hostname|cannot contain spaces|cannot contain any special characters"),
			},
		},
	})
}

func TestAccCloudflareCustomHostname_InvalidImportID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: "invalid-import-id",
				ExpectError:   regexp.MustCompile("invalid ID|expected format"),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: zoneID, // Missing hostname ID
				ExpectError:   regexp.MustCompile("invalid ID|expected format"),
			},
		},
	})
}

func TestAccCloudflareCustomHostname_SSLCertificateTransition(t *testing.T) {
	// This test reproduces GitHub Issue #3012: SSL changing from Cloudflare managed cert to custom certificate fails
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{rnd + "." + domain}, expiry)
	if err != nil {
		t.Error(err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with Cloudflare-managed certificate
				Config: testAccCheckCloudflareCustomHostnameSSLTransitionStep1(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttr(resourceName, "ssl.type", "dv"),
					// Verify it's using Cloudflare-managed certificate (no custom cert/key)
					resource.TestCheckNoResourceAttr(resourceName, "ssl.custom_certificate"),
					resource.TestCheckNoResourceAttr(resourceName, "ssl.custom_key"),
				),
			},
			{
				// Step 2: Transition to custom certificate
				Config: testAccCheckCloudflareCustomHostnameSSLTransitionStep2(zoneID, rnd, domain, cert, key),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						// Test changing certificate authority
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssl").AtMapKey("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttr(resourceName, "ssl.type", "dv"),
					resource.TestCheckResourceAttr(resourceName, "ssl.certificate_authority", "lets_encrypt"),
				),
			},
		},
	})
}

func TestAccCloudflareCustomHostname_TLS13(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithTLSVersion(zoneID, rnd, domain, "1.3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.min_tls_version", "1.3"),
				),
			},
		},
	})
}

func testSweepCloudflareCustomHostnames(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping custom hostnames sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	hostnames, _, hostnamesErr := client.CustomHostnames(ctx, zoneID, 1, cloudflare.CustomHostname{})
	if hostnamesErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare custom hostnames: %s", hostnamesErr))
		return hostnamesErr
	}

	if len(hostnames) == 0 {
		tflog.Info(ctx, "No Cloudflare custom hostnames to sweep")
		return nil
	}

	for _, hostname := range hostnames {
		if !utils.ShouldSweepResource(hostname.Hostname) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting custom hostname: %s (%s)", hostname.Hostname, hostname.ID))
		err := client.DeleteCustomHostname(ctx, zoneID, hostname.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete custom hostname %s (%s): %s", hostname.Hostname, hostname.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted custom hostname: %s (%s)", hostname.Hostname, hostname.ID))
	}

	return nil
}

func TestAccCloudflareCustomHostname_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", zoneID, s.RootModule().Resources[resourceName].Primary.ID), nil
				},
				ImportStateVerifyIgnore: []string{
					"ssl.validation_records",
					"ssl.validation_errors",
					"ssl.wildcard",
					"ssl.certificate_authority",
					"ownership_verification",
					"ownership_verification_http",
					"created_at",
					"status",
					"verification_errors",
					"wait_for_ssl_pending_validation",
				},
			},
		},
	})
}

func TestAccCloudflareCustomHostname_WithCertificate(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	expiry := time.Now().Add(time.Hour * 1)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{rnd + "." + domain}, expiry)
	if err != nil {
		t.Error(err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCertificate(zoneID, rnd, domain, cert, key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "http"),
					resource.TestCheckResourceAttr(resourceName, "ssl.type", "dv"),
					resource.TestCheckResourceAttrSet(resourceName, "ssl.custom_certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "ssl.custom_key"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamebasic.tf", zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WaitForActive(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWaitForActive(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWaitForActive(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewaitforactive.tf", zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithCustomOriginServer(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomOriginServer(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "custom_origin_server", fmt.Sprintf("origin.%s.terraform.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(resourceName, "custom_origin_sni", fmt.Sprintf("origin.%s.terraform.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWithCustomOriginServer(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithcustomoriginserver.tf", zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithHTTPValidation(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithHTTPValidation(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "http"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWithHTTPValidation(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithhttpvalidation.tf", zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithCustomSSLSettings(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomSSLSettings(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.http2", "off"),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.min_tls_version", "1.2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.ciphers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.certificate_authority", "google"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWithCustomSSLSettings(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithcustomsslsettings.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareCustomHostnameWithCustomSSLSettingsUpdated(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithcustomsslsettingsupdated.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareCustomHostnameWithNoSSL(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithnossl.tf", zoneID, rnd, domain)
}

// SSL is required parameter, per the schema
//
// func TestAccCloudflareCustomHostname_WithNoSSL(t *testing.T) {
// 	t.Parallel()
// 	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
// 	domain := os.Getenv("CLOUDFLARE_DOMAIN")
// 	rnd := utils.GenerateRandomResourceName()
// 	resourceName := "cloudflare_custom_hostname." + rnd
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckCloudflareCustomHostnameWithNoSSL(zoneID, rnd, domain),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
// 					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
// 					resource.TestCheckResourceAttr(resourceName, "ssl.#", "0"),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccCloudflareCustomHostname_UpdatingZoneForcesNewResource(t *testing.T) {

	var before, after cloudflare.CustomHostname
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	altZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	altDomain := os.Getenv("CLOUDFLARE_ALT_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AlternateZoneID(t)
			acctest.TestAccPreCheck_AlternateDomain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &before),
				),
			},
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(altZoneID, rnd, altDomain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(altZoneID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, altDomain))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(altZoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, altDomain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &after),
					testAccCheckCloudflareCustomHostnameRecreated(&before, &after),
				),
			},
		},
	})
}

// func TestAccCloudflareCustomHostname_Import(t *testing.T) {
// 	t.Parallel()
//
// 	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
// 	domain := os.Getenv("CLOUDFLARE_DOMAIN")
// 	rnd := utils.GenerateRandomResourceName()
// 	resourceName := "cloudflare_custom_hostname." + rnd
//
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
// 			},
// 			{
// 				ResourceName:        resourceName,
// 				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
// 				ImportState:         true,
// 				ImportStateVerify:   true,
// 				ImportStateVerifyIgnore: []string{
// 					"ssl.#",
// 					"ssl.certificate_authority",
// 					"ssl.validation_records",
// 					"ssl.validation_errors",
// 					"ssl.custom_certificate",
// 					"ssl.custom_key",
// 					"ssl.method",
// 					"ssl.status",
// 					"ssl.type",
// 					"ssl.wildcard",
// 					"wait_for_ssl_pending_validation",
// 				},
// 			},
// 		},
// 	})
// }

func testAccCheckCloudflareCustomHostnameRecreated(before, after *cloudflare.CustomHostname) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of CustomHostname Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflareCustomHostnameExists(n string, customHostname *cloudflare.CustomHostname) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No CustomHostname ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundCustomHostname, err := client.CustomHostname(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundCustomHostname.ID != rs.Primary.ID {
			return fmt.Errorf("CustomHostname not found")
		}

		*customHostname = foundCustomHostname

		return nil
	}
}

func TestAccCloudflareCustomHostname_WithCustomMetadata(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomMetadata(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.method", "txt"),
					resource.TestCheckResourceAttr(resourceName, "ssl.wildcard", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.value"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.type"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification.name"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_url"),
					resource.TestCheckResourceAttrSet(resourceName, "ownership_verification_http.http_body"),
					resource.TestCheckResourceAttr(resourceName, "custom_metadata.customer_id", "12345"),
					resource.TestCheckResourceAttr(resourceName, "custom_metadata.redirect_to_https", "true"),
					resource.TestCheckResourceAttr(resourceName, "custom_metadata.security_tag", "low"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWithCustomMetadata(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("customhostnamewithcustommetadata.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareCustomHostnameWithCA(zoneID, rnd, domain, ca string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%s" {
	zone_id = "%s"
	hostname = "%s.%s"
	ssl = {
		method = "txt"
		type = "dv"
		certificate_authority = "%s"
	}
	
	lifecycle {
		ignore_changes = [
			created_at,
			ownership_verification,
			ownership_verification_http,
			ssl.wildcard,
			status,
			verification_errors
		]
	}
}
`, rnd, zoneID, rnd, domain, ca)
}


func testAccCheckCloudflareCustomHostnameWithTLSVersion(zoneID, rnd, domain, tlsVersion string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%s" {
	zone_id = "%s"
	hostname = "%s.%s"
	ssl = {
		method = "txt"
		type = "dv"
		settings = {
			min_tls_version = "%s"
		}
	}
	
	lifecycle {
		ignore_changes = [
			created_at,
			ownership_verification,
			ownership_verification_http,
			ssl.certificate_authority,
			ssl.wildcard,
			status,
			verification_errors
		]
	}
}
`, rnd, zoneID, rnd, domain, tlsVersion)
}

func testAccCheckCloudflareCustomHostnameSSLTransitionStep1(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%s" {
	zone_id = "%s"
	hostname = "%s.%s"
	ssl = {
		method = "txt"
		type = "dv"
		# Using Cloudflare-managed certificate (no custom cert/key)
	}
	
	lifecycle {
		ignore_changes = [
			created_at,
			ownership_verification,
			ownership_verification_http,
			ssl.certificate_authority,
			ssl.wildcard,
			status,
			verification_errors
		]
	}
}
`, rnd, zoneID, rnd, domain)
}

func testAccCheckCloudflareCustomHostnameSSLTransitionStep2(zoneID, rnd, domain, cert, key string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%s" {
	zone_id = "%s"
	hostname = "%s.%s"
	ssl = {
		method = "txt"
		type = "dv"
		# Test changing certificate_authority instead of adding custom cert
		certificate_authority = "lets_encrypt"
	}
	
	lifecycle {
		ignore_changes = [
			created_at,
			ownership_verification,
			ownership_verification_http,
			ssl.wildcard,
			status,
			verification_errors
		]
	}
}
`, rnd, zoneID, rnd, domain)
}

func testAccCloudflareCustomHostnameDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		return fmt.Errorf("failed to create Cloudflare client: %s", clientErr)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.CustomHostname(context.Background(), zoneID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("custom hostname still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareCustomHostnameInvalidHostname(zoneID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%s" {
	zone_id = "%s"
	hostname = "invalid hostname with spaces.terraform.cfapi.net"
	ssl = {
		method = "txt"
		type = "dv"
	}
}
`, rnd, zoneID)
}

func testAccCheckCloudflareCustomHostnameWithCertificate(zoneID, rnd, domain, cert, key string) string {
	return acctest.LoadTestCase("customhostnamewithcertificate.tf", zoneID, rnd, domain, cert, key)
}
