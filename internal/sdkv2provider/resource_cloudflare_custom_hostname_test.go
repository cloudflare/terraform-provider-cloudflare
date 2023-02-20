package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

	client, clientErr := sharedClient()
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
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.method", "txt"),
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
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl {
    method = "txt"
  }
}
`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WaitForActive(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWaitForActive(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.method", "txt"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ssl_pending_validation", "true"),
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
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl {
    method = "txt"
  }
  wait_for_ssl_pending_validation = true
}
`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithCustomOriginServer(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomOriginServer(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "custom_origin_server", fmt.Sprintf("origin.%s.terraform.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(resourceName, "custom_origin_sni", fmt.Sprintf("origin.%s.terraform.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.method", "txt"),
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
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  custom_origin_server = "origin.%[2]s.terraform.cfapi.net"
	custom_origin_sni = "origin.%[2]s.terraform.cfapi.net"
  ssl {
    method = "txt"
  }
}

resource "cloudflare_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "origin.%[2]s.terraform.cfapi.net"
  value   = "example.com"
  type    = "CNAME"
  ttl     = 3600
}`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithHTTPValidation(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithHTTPValidation(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.method", "http"),
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
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl {
    method = "http"
  }
}
`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithCustomSSLSettings(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomSSLSettings(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.http2", "off"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.min_tls_version", "1.2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.ciphers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.certificate_authority", "digicert"),
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
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl {
    method = "http"
    settings {
      http2 = "off"
      min_tls_version = "1.2"
      ciphers = [
        "ECDHE-RSA-AES128-GCM-SHA256",
        "AES128-SHA"
      ]
      early_hints = "off"
    }
  }
}
`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_Update(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomSSLSettings(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.http2", "off"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.min_tls_version", "1.2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.ciphers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.early_hints", "off"),
				),
			},
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomSSLSettingsUpdated(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.http2", "off"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.min_tls_version", "1.1"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.ciphers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.settings.0.early_hints", "off"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameWithCustomSSLSettingsUpdated(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl {
    method = "http"
    settings {
      http2 = "off"
      min_tls_version = "1.1"
      ciphers = [
        "ECDHE-RSA-AES128-GCM-SHA256",
        "AES128-SHA"
      ]
      early_hints = "off"
    }
  }
}
`, zoneID, rnd, domain)
}

func testAccCheckCloudflareCustomHostnameWithNoSSL(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
}
`, zoneID, rnd, domain)
}

func TestAccCloudflareCustomHostname_WithNoSSL(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithNoSSL(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareCustomHostname_UpdatingZoneForcesNewResource(t *testing.T) {
	t.Parallel()

	var before, after cloudflare.CustomHostname
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	altZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	altDomain := os.Getenv("CLOUDFLARE_ALT_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAltZoneID(t)
			testAccPreCheckAltDomain(t)
		},
		ProviderFactories: providerFactories,
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

func TestAccCloudflareCustomHostname_Import(t *testing.T) {
	t.Parallel()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd, domain),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateVerifyIgnore: []string{
					"ssl.#",
					"ssl.0.certificate_authority",
					"ssl.0.validation_records",
					"ssl.0.validation_errors",
					"ssl.0.custom_certificate",
					"ssl.0.custom_key",
					"ssl.0.method",
					"ssl.0.status",
					"ssl.0.type",
					"ssl.0.wildcard",
					"wait_for_ssl_pending_validation",
				},
			},
		},
	})
}

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

		client := testAccProvider.Meta().(*cloudflare.API)
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
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameWithCustomMetadata(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.method", "txt"),
					resource.TestCheckResourceAttr(resourceName, "ssl.0.wildcard", "true"),
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
	return fmt.Sprintf(`
	resource "cloudflare_custom_hostname" "%[2]s" {
		zone_id = "%[1]s"
		hostname = "%[2]s.%[3]s"
		ssl {
			method = "txt"
			wildcard = true
		}
		custom_metadata = {
			"customer_id" = 12345
			"redirect_to_https" = true
			"security_tag" = "low"
		}
	}
	`, zoneID, rnd, domain)
}
