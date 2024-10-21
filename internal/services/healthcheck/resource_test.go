package healthcheck_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareHealthcheckTCPExists(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "tcp_config.port", "80"),
					resource.TestCheckResourceAttr(name, "tcp_config.method", "connection_established"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckTCPUpdate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "name", rnd),
				),
			},
			{
				PreConfig: func() {
					initialID = healthcheck.ID
				},
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd+"-updated", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					func(state *terraform.State) error {
						if initialID != healthcheck.ID {
							return fmt.Errorf("wanted update but healthcheck got recreated (id changed %q -> %q)", initialID, healthcheck.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "name", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckHTTPExists(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)
	var healthcheck cloudflare.Healthcheck

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckHTTP(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHealthcheckExists(name, zoneID, &healthcheck),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "http_config.port", "80"),
					resource.TestCheckResourceAttr(name, "http_config.method", "GET"),
				),
			},
		},
	})
}

func TestAccCloudflareHealthcheckMissingRequired(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Healthcheck ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundHealthcheck, err := client.Healthcheck(context.Background(), zoneID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*load = foundHealthcheck

		return nil
	}
}

func testAccCheckCloudflareHealthcheckTCP(zoneID, name, ID string) string {
	return acctest.LoadTestCase("healthchecktcp.tf", zoneID, name, ID)
}

func testAccCheckCloudflareHealthcheckHTTP(zoneID, ID string) string {
	return acctest.LoadTestCase("healthcheckhttp.tf", zoneID, ID)
}

func testAccCheckHealthcheckConfigMissingRequired(zoneID, ID string) string {
	return acctest.LoadTestCase("acccheckhealthcheckconfigmissingrequired.tf", zoneID, ID)
}
