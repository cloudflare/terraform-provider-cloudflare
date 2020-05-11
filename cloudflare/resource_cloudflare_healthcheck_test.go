package cloudflare

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareHealthcheckTCPExists(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, "tcp-test-healthcheck", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "port", "80"),
					resource.TestCheckResourceAttr(name, "method", "connection_established"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckTCPUpdate(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, "tcp-test-healthcheck-update", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "name", "tcp-test-healthcheck-update"),
				),
			},
			{
				PreConfig: func() {
					initialID = healthcheck.ID
				},
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, "tcp-test-healthcheck-updated", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					func(state *terraform.State) error {
						if initialID != healthcheck.ID {
							return fmt.Errorf("wanted update but healthcheck got recreated (id changed %q -> %q)",
								initialID, healthcheck.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "name", "tcp-test-healthcheck-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckHTTPExists(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckHTTP(zoneID, "http-test-healthcheck", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "port", "80"),
					resource.TestCheckResourceAttr(name, "method", "GET"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckMissingRequired(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckHealthcheckConfigMissingRequired(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("The argument \"name\" is required, but no definition was found.")),
			},
		},
	})
}

func testAccCheckCloudflareHealthcheckExists(n string, zoneID string, load *cloudflare.Healthcheck) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Healthcheck ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundHealthcheck, err := client.Healthcheck(zoneID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*load = foundHealthcheck

		return nil
	}
}

func testAccCheckCloudflareHealthcheckTCP(zoneID, name, ID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_healthcheck" "%[3]s" {
    zone_id = "%[1]s"
    name = "%[2]s"
    address = "example.com"
    type = "TCP"
    method = "connection_established"
    port = 80
  }`, zoneID, name, ID)
}

func testAccCheckCloudflareHealthcheckHTTP(zoneID, name, ID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_healthcheck" "%[3]s" {
    zone_id = "%[1]s"
    name = "%[2]s"
    address = "example.com"
    type = "HTTP"
    expected_codes = [
      "200"
    ]
  }`, zoneID, name, ID)
}

func testAccCheckHealthcheckConfigMissingRequired(zoneID, ID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_healthcheck" "%[2]s" {
    zone_id = "%[1]s"
    description = "Example health check description"
  }`, zoneID, ID)
}
