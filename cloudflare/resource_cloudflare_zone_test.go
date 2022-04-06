package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareZone_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "true"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPlan(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPartialSetup(rnd, "foo.net", "true", "false", "free"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithExplicitFullSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "true", "false", "free"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", "true", "false"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", "true", "false"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(rnd, "żółw.cfapi.net", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config:   testZoneConfig(rnd, "xn--w-uga1v8h.cfapi.net", "true", "false"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithExplicitFullSetup(rnd, fmt.Sprintf("%s.cfapi.net", rnd), "false", "false", "enterprise"),
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

func testZoneConfig(resourceID, zoneName, paused, jumpStart string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
				}`, resourceID, zoneName, paused, jumpStart)
}

func testZoneConfigWithPlan(resourceID, zoneName, paused, jumpStart, plan string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
				}`, resourceID, zoneName, paused, jumpStart, plan)
}

func TestAccCloudflareZone_SetType(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_zone." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithExplicitFullSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise"),
			},
			{
				Config: testZoneConfigWithPartialSetup(rnd, fmt.Sprintf("%s.%s", rnd, zoneName), "true", "false", "enterprise"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
		},
	})
}

func testZoneConfigWithPartialSetup(resourceID, zoneName, paused, jumpStart, plan string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
					type = "partial"
				}`, resourceID, zoneName, paused, jumpStart, plan)
}

func testZoneConfigWithExplicitFullSetup(resourceID, zoneName, paused, jumpStart, plan string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
					type = "full"
				}`, resourceID, zoneName, paused, jumpStart, plan)
}
