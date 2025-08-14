package regional_hostname_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
	client := acctest.SharedClient()

	// Get all regional hostnames for the test zone
	hostnames, err := client.Addressing.RegionalHostnames.List(context.Background(), addressing.RegionalHostnameListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		return fmt.Errorf("failed to list regional hostnames: %w", err)
	}

	for _, hostname := range hostnames.Result {
		// Only delete test hostnames (contain random resource names pattern)
		if len(hostname.Hostname) >= 10 {
			_, err := client.Addressing.RegionalHostnames.Delete(context.Background(), hostname.Hostname, addressing.RegionalHostnameDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				return fmt.Errorf("failed to delete regional hostname %s: %w", hostname.Hostname, err)
			}
		}
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

func TestAccCloudflareRegionalHostname_Migration_TimeoutsRemoval(t *testing.T) {
	// This test verifies that the migration tool properly removes the timeouts block
	// when upgrading from v4 to v5, since v5 provider no longer supports timeouts
	// configuration for regional_hostname resources.

	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_regional_hostname." + rnd
	tmpDir := t.TempDir()

	// V4 config with timeouts block
	v4Config := testRegionalHostnameV4ConfigWithTimeouts(rnd, zoneName, "ca")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider including timeouts
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1", // Use exact v4 version
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration from v4 to current version
			// This will run the migration tool which should remove the timeouts block
			// and the state upgrade function will handle the state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
			}),
			{
				// Step 3: Apply migrated config with current provider
				// Should succeed without any timeouts configuration
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					// Verify routing has default value (handled by state upgrader)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				},
			},
			{
				// Step 4: Import verification to ensure resource is properly accessible
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateIdFunc:        testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore:  []string{"created_on"}, // Computed field may vary
			},
		},
	})
}

func testRegionalHostnameV4ConfigWithTimeouts(name string, zoneName, regionKey string) string {
	// Use a random subdomain to avoid conflicts with existing regional hostnames
	hostname := fmt.Sprintf("%s.%s", name, zoneName)
	return acctest.LoadTestCase("regionalhostname_v4_with_timeouts.tf", name, zoneID, hostname, regionKey)
}
