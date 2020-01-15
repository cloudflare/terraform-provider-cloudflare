package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareCustomHostname_Basic(t *testing.T) {
	t.Parallel()
	var customHostname cloudflare.CustomHostname
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &customHostname),
					resource.TestCheckResourceAttr(
						resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", "thisis@cool.com"),
					resource.TestCheckResourceAttr(
						resourceName, "ssl.method", "http"),
				),
			},
		},
	})
}

func TestAccCloudflareCustomHostname_Advanced(t *testing.T) {
	t.Parallel()
	var customHostname cloudflare.CustomHostname
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameAdvanced(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomHostnameExists(resourceName, &customHostname),
					resource.TestCheckResourceAttr(
						resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(
						resourceName, "custom_origin_server", "serverone.myinfra.com"),
					resource.TestCheckResourceAttr(
						resourceName, "ssl.method", "http"),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", "thisis@cool.com"),
					resource.TestCheckResourceAttr(
						resourceName, "ssl.custom_hostname_settings.tls13", "on"),
					resource.TestCheckResourceAttr(
						resourceName, "ssl.custom_hostname_settings.http2", "on"),
					resource.TestCheckResourceAttr(
						resourceName, "ssl.custom_hostname_settings.mintlsversion", "1.1"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameBasic(zoneID string, rName string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "thisis@cool.com"
  ssl = {
    method = "http"
  }
}`, zoneID, rName)
}

func testAccCheckCloudflareCustomHostnameAdvanced(zoneID string, rName string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "thisis@cool.com"
  custom_origin_server = "serverone.myinfra.com"
  ssl {
	method = "http"
	custom_hostname_settings {
		http2 = "on"
		tls13 = "on"
		mintlsversion = "1.1"
		ciphers = ["ECDHE-RSA-AES128-GCM-SHA256","AES128-SHA"]
	}
  }
}`, zoneID, rName)
}

func testAccCheckCloudflareCustomHostnameDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname" {
			continue
		}

		err := client.DeleteCustomHostname(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("custom hostname still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareCustomHostnameExists(n string, customHostname *cloudflare.CustomHostname) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundCustomHostname, err := client.CustomHostname(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundCustomHostname.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}

		*customHostname = foundCustomHostname

		return nil
	}
}
