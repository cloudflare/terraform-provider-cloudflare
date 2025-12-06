package zone_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func init() {
	resource.AddTestSweepers("cloudflare_zone", &resource.Sweeper{
		Name: "cloudflare_zone",
		F:    testSweepCloudflareZone,
	})
}

// testSweepCloudflareZone removes test zones created during acceptance tests.
//
// This sweeper:
// - Lists all zones in the test account
// - Filters for test zones (prefix: tf-acc-test-, tf-acctest-, test prefixes with terraform.cfapi.net)
// - Deletes matching zones
// - Continues on errors to sweep as many zones as possible
//
// Run with: go test ./internal/services/zone/ -v -sweep=all
//
// Requires:
// - CLOUDFLARE_ACCOUNT_ID (for account-scoped zones)
// - CLOUDFLARE_EMAIL + CLOUDFLARE_API_KEY or CLOUDFLARE_API_TOKEN
func testSweepCloudflareZone(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return nil
	}

	// List all zones in the account
	page, err := client.Zones.List(ctx, zones.ZoneListParams{
		Account: cloudflare.F(zones.ZoneListParamsAccount{
			ID: cloudflare.F(accountID),
		}),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zones: %s",err))
		return err
	}
	if len(page.Result) == 0 {
		tflog.Info(ctx, "No Cloudflare zones to sweep")
		return nil
	}

	for _, zone := range page.Result {
		// Use standard filtering helper (also checked in isTestZone for zone-specific patterns)
		if !utils.ShouldSweepResource(zone.Name) && !isTestZone(zone.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting test zone: %s (%s)", zone.Name, zone.ID))
		_, err := client.Zones.Delete(ctx, zones.ZoneDeleteParams{
			ZoneID: cloudflare.F(zone.ID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zone %s (%s): %s", zone.Name, zone.ID,err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted zone: %s (%s)", zone.Name, zone.ID))
	}

	return nil
}

// isTestZone checks if a zone name matches test naming patterns
func isTestZone(name string) bool {
	// Use standard test resource naming convention
	if utils.ShouldSweepResource(name) {
		return true
	}

	// Match terraform.cfapi.net test domains (zone-specific pattern)
	if strings.HasSuffix(name, ".terraform.cfapi.net") ||
		name == "terraform.cfapi.net" {
		return true
	}

	// Match .cfapi.net domains (from acceptance tests)
	if strings.HasSuffix(name, ".cfapi.net") &&
		!strings.Contains(name, "production") &&
		!strings.Contains(name, "prod-") {
		return true
	}

	return false
}

func TestAccCloudflareZone_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
					// Check development_mode
					resource.TestCheckResourceAttr(name, "development_mode", "0"),
					// Spot-check zone metadata
					resource.TestCheckResourceAttr(name, "meta.phishing_detected", "false"),
					// Check owner information
					resource.TestCheckResourceAttr(name, "owner.id", accountID), resource.TestCheckResourceAttrSet(name, "owner.name"),
					resource.TestCheckResourceAttr(name, "owner.type", "organization"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_PartialSetup(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, accountID, fmt.Sprintf("%s.net", rnd), "partial"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.net", rnd)),
					resource.TestCheckResourceAttr(name, "type", "partial"),
					resource.TestCheckResourceAttrSet(name, "verification_key"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "0"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		}})
}

func TestAccCloudflareZone_FullSetup(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, accountID, fmt.Sprintf("%s.cfapi.net", rnd), "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccZoneWithUnicodeIsStoredAsUnicode(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	unicodeDomain := fmt.Sprintf("żóła.%s.cfapi.net", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, unicodeDomain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", unicodeDomain),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccZoneWithoutUnicodeIsStoredAsUnicode(t *testing.T) {
	t.Skip("unicode translation not working correctly")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := fmt.Sprintf("xn--w-uga1v8h.%s.cfapi.net", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("żółw.%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccZonePerformsUnicodeComparison(t *testing.T) {
	t.Skip("unicode translation not working correctly")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("żółw.%s.cfapi.net", rnd), accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("żółw.%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config:   testZoneConfig(rnd, fmt.Sprintf("xn--w-uga1v8h.%s.cfapi.net", rnd), accountID),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("żółw.%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_WithEnterprisePlanVanityNameServers(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Need working zone_subscription to create enterprise plan before setting vanity name servers")

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeVanityNameServersSetup(rnd, accountID, fmt.Sprintf("%s.%s", rnd, zoneName), "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
					resource.TestCheckResourceAttr(name, "vanity_name_servers.#", "2"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_Secondary(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, accountID, fmt.Sprintf("%s.%s", rnd, zoneName), "secondary"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "secondary"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_SecondaryWithVanityNameServers(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Need working zone_subscription to create enterprise plan before setting vanity name servers")

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeVanityNameServersSetup(rnd, accountID, fmt.Sprintf("%s.%s", rnd, zoneName), "secondary"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "secondary"),
					resource.TestCheckResourceAttr(name, "vanity_name_servers.#", "2"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_TogglePaused(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPaused(rnd, accountID, fmt.Sprintf("%s.cfapi.net", rnd), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config: testZoneConfigWithPaused(rnd, accountID, fmt.Sprintf("%s.cfapi.net", rnd), true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config: testZoneConfigWithPaused(rnd, accountID, fmt.Sprintf("%s.cfapi.net", rnd), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZone_SetType(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, accountID, fmt.Sprintf("%s.%s", rnd, zoneName), "full"),
			},
			{
				Config: testZoneConfigWithTypeSetup(rnd, accountID, fmt.Sprintf("%s.%s", rnd, zoneName), "partial"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testZoneConfig(resourceID, zoneName, accountID string) string {
	return acctest.LoadTestCase("zoneconfig.tf", resourceID, zoneName, accountID)
}

func testZoneConfigWithTypeSetup(resourceID, accountID, zoneName, zoneType string) string {
	return acctest.LoadTestCase("zoneconfigwithtypesetup.tf", resourceID, accountID, zoneName, zoneType)
}

func testZoneConfigWithTypeVanityNameServersSetup(resourceID, accountID, zoneName, zoneType string) string {
	return acctest.LoadTestCase("zoneconfigwithtypevanitynameserverssetup.tf", resourceID, accountID, zoneName, zoneType)
}

func testZoneConfigWithPaused(resourceID, accountID, zoneName string, paused bool) string {
	return acctest.LoadTestCase("zoneconfigwithpaused.tf", resourceID, accountID, zoneName, paused)
}
