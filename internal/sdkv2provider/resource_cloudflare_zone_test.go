package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZone_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_BasicWithJumpStartEnabled(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "true", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
					resource.TestCheckResourceAttr(name, "jump_start", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_WithPlan(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPlan(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_PartialSetup(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, "foo.net", "true", "false", "free", accountID, "partial"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "foo.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_FullSetup(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free", accountID, "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccZoneWithUnicodeIsStoredAsUnicode(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccZoneWithoutUnicodeIsStoredAsUnicode(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccZonePerformsUnicodeComparison(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", "true", "false", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config:   testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", "true", "false", accountID),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_WithEnterprisePlan(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "false", "false", "enterprise", accountID, "full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.cfapi.net", rnd)),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDEnterprise),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZone_Secondary(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithTypeSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise", accountID, "secondary"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", fmt.Sprintf("%s.%s", rnd, zoneName)),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDEnterprise),
					resource.TestCheckResourceAttr(name, "type", "secondary"),
				),
			},
		},
	})
}

func testZoneConfig(resourceID, zoneName, paused, jumpStart, accountID string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					account_id = "%[5]s"
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
				}`, resourceID, zoneName, paused, jumpStart, accountID)
}

func testZoneConfigWithPlan(resourceID, zoneName, paused, jumpStart, plan, accountID string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					account_id = "%[6]s"
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
				}`, resourceID, zoneName, paused, jumpStart, plan, accountID)
}

func TestAccCloudflareZone_SetType(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					account_id = "%[6]s"
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
					type = "%[7]s"
				}`, resourceID, zoneName, paused, jumpStart, plan, accountID, zoneType)
}
