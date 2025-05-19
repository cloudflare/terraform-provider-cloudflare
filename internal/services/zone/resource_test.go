package zone_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	planIDFree       = "free"
	planIDLite       = "lite"
	planIDPro        = "pro"
	planIDProPlus    = "pro_plus"
	planIDBusiness   = "business"
	planIDEnterprise = "enterprise"

	planIDPartnerFree       = "partners_free"
	planIDPartnerPro        = "partners_pro"
	planIDPartnerBusiness   = "partners_business"
	planIDPartnerEnterprise = "partners_enterprise"
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
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_WithPlan(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPlan(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
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
				Config: testZoneConfigWithTypeSetup(rnd, "foo.net", "true", "false", "free", accountID, "partial"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "foo.net"),
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
		},
	})
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
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free", accountID, "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccZoneWithUnicodeIsStoredAsUnicode(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_WithEnterprisePlan(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "false", "false", "enterprise", accountID, "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_WithEnterprisePlanVanityNameServers(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeVanityNameServersSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "false", "false", "enterprise", accountID, "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "full"),
					resource.TestCheckResourceAttr(name, "vanity_name_servers.#", "2"),
				),
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
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise", accountID, "secondary"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "secondary"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_SecondaryWithVanityNameServers(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeVanityNameServersSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise", accountID, "secondary"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "secondary"),
					resource.TestCheckResourceAttr(name, "vanity_name_servers.#", "2"),
				),
			},
		},
	})
}

func testZoneConfig(resourceID, zoneName, paused, jumpStart, accountID string) string {
	return acctest.LoadTestCase("zoneconfig.tf", resourceID, zoneName, paused, jumpStart, accountID)
}

func testZoneConfigWithPlan(resourceID, zoneName, paused, jumpStart, plan, accountID string) string {
	return acctest.LoadTestCase("zoneconfigwithplan.tf", resourceID, zoneName, paused, jumpStart, plan, accountID)
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
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise", accountID, "full"),
			},
			{
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise", accountID, "partial"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
		},
	})
}

func testZoneConfigWithTypeSetup(resourceID, zoneName, paused, jumpStart, plan, accountID, zoneType string) string {
	return acctest.LoadTestCase("zoneconfigwithtypesetup.tf", resourceID, zoneName, paused, jumpStart, plan, accountID, zoneType)
}

func testZoneConfigWithTypeVanityNameServersSetup(resourceID, zoneName, paused, jumpStart, plan, accountID, zoneType string) string {
	return acctest.LoadTestCase("zoneconfigwithtypevanitynameserverssetup.tf", resourceID, zoneName, paused, jumpStart, plan, accountID, zoneType)
}
