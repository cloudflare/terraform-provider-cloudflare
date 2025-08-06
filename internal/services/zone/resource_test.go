package zone_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
				Config: testZoneConfig(rnd, "żółw.cfapi.net", accountID),
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "żółw.cfapi.net"),
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config:   testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", accountID),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "żółw.cfapi.net"),
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
