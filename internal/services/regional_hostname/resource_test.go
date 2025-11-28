package regional_hostname_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var zoneID = os.Getenv("CLOUDFLARE_ZONE_ID")

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_regional_hostname", &resource.Sweeper{
		Name: "cloudflare_regional_hostname",
		F:    testSweepCloudflareRegionalHostname,
	})
}

func testSweepCloudflareRegionalHostname(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	if zoneID == "" {
		tflog.Info(ctx, "Skipping regional hostnames sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// Get all regional hostnames for the test zone
	hostnames, err := client.Addressing.RegionalHostnames.List(ctx, addressing.RegionalHostnameListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list regional hostnames: %s", err))
		return fmt.Errorf("failed to list regional hostnames: %w", err)
	}

	if len(hostnames.Result) == 0 {
		tflog.Info(ctx, "No regional hostnames to sweep")
		return nil
	}

	for _, hostname := range hostnames.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(hostname.Hostname) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting regional hostname: %s (zone: %s)", hostname.Hostname, zoneID))
		_, err := client.Addressing.RegionalHostnames.Delete(ctx, hostname.Hostname, addressing.RegionalHostnameDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete regional hostname %s: %s", hostname.Hostname, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted regional hostname: %s", hostname.Hostname))
	}

	return nil
}

func TestAccCloudflareRegionalHostname_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_regional_hostname." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname := fmt.Sprintf("%s.%s", rnd, zoneName) // Expected hostname

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRegionalHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testRegionalHostnameBasicConfig(rnd, zoneName, "ca"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "region_key", "ca"),
					resource.TestCheckResourceAttr(resourceName, "routing", "dns"),
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"created_on"},
			},
		},
	})
}

func TestAccCloudflareRegionalHostname_UpdateRegion(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_regional_hostname." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname := fmt.Sprintf("%s.%s", rnd, zoneName) // Expected hostname

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRegionalHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testRegionalHostnameConfig(rnd, zoneName, "ca", "dns"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "region_key", "ca"),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "routing", "dns"),
				),
			},
			{
				Config: testRegionalHostnameConfig(rnd, zoneName, "au", "dns"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "region_key", "au"),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "routing", "dns"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"created_on"},
			},
		},
	})
}

func TestAccCloudflareRegionalHostname_DifferentRegions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_regional_hostname." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname := fmt.Sprintf("%s.%s", rnd, zoneName) // Expected hostname

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRegionalHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testRegionalHostnameConfig(rnd, zoneName, "us", "dns"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "region_key", "us"),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "routing", "dns"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_on"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"created_on"},
			},
		},
	})
}

func testRegionalHostnameConfig(name string, zoneName, regionKey, routing string) string {
	// Use random subdomain to avoid conflicts
	hostname := fmt.Sprintf("%s.%s", name, zoneName)
	return acctest.LoadTestCase("regionalhostnameconfig.tf", name, zoneID, hostname, regionKey, routing)
}

func testRegionalHostnameBasicConfig(name string, zoneName, regionKey string) string {
	// Use random subdomain to avoid conflicts
	hostname := fmt.Sprintf("%s.%s", name, zoneName)
	return acctest.LoadTestCase("regionalhostname_basic.tf", name, zoneID, hostname, regionKey)
}

func testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		hostname := rs.Primary.Attributes["hostname"]

		return fmt.Sprintf("%s/%s", zoneID, hostname), nil
	}
}

func testAccCheckCloudflareRegionalHostnameDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_regional_hostname" {
			continue
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		hostname := rs.Primary.Attributes["hostname"]

		_, err := client.Addressing.RegionalHostnames.Get(
			context.Background(),
			hostname,
			addressing.RegionalHostnameGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)
		if err == nil {
			return fmt.Errorf("regional hostname %s still exists in zone %s", hostname, zoneID)
		}
	}

	return nil
}
