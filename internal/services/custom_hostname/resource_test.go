package custom_hostname_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_custom_hostname", &resource.Sweeper{
		Name: "cloudflare_custom_hostname",
		F:    testSweepCloudflareCustomHostnames,
	})
}

func testSweepCloudflareCustomHostnames(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	hostnames, _, hostnamesErr := client.CustomHostnames(context.Background(), zoneID, 1, cloudflare.CustomHostname{})
	if hostnamesErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare custom hostnames: %s", hostnamesErr))
	}

	if len(hostnames) == 0 {
		log.Print("[DEBUG] No Cloudflare custom hostnames to sweep")
		return nil
	}

	for _, hostname := range hostnames {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare custom hostname: %s", hostname.ID))
		err := client.DeleteCustomHostname(context.Background(), zoneID, hostname.ID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare custom hostname (%s): %s", hostname.Hostname, err))
		}
	}

	return nil
}

func TestAccCloudflareCustomHostname_Basic(t *testing.T) {
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
		},
	})
}

func TestAccCloudflareCustomHostname_WithCertificate(t *testing.T) {
	t.Parallel()
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
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomSSLSettingsUpdated(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.http2", "off"),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.min_tls_version", "1.1"),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.ciphers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.settings.early_hints", "off"),
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
	t.Parallel()

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
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &before),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
				),
			},
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(altZoneID, rnd, altDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &after),
					testAccCheckCloudflareCustomHostnameRecreated(&before, &after),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, altZoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, altDomain)),
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

func testAccCheckCloudflareCustomHostnameWithCertificate(zoneID, rnd, domain, cert, key string) string {
	return acctest.LoadTestCase("customhostnamewithcertificate.tf", zoneID, rnd, domain, cert, key)
}
