package cloudflare

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareZoneBasic(t *testing.T) {
	name := "cloudflare_zone.tf-acc-basic-zone"
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig(resourceName, "example.cfapi.net", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneWithPlan(t *testing.T) {
	name := "cloudflare_zone.tf-acc-with-plan-zone"
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPlan(resourceName, "example.cfapi.net", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
		},
	})
}

func TestAccCloudflareZonePartialSetup(t *testing.T) {
	name := "cloudflare_zone.tf-acc-partial-setup-zone"
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithPartialSetup(resourceName, "example.cfapi.net", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneFullSetup(t *testing.T) {
	name := "cloudflare_zone.tf-acc-full-setup-zone"
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfigWithExplicitFullSetup(resourceName, "example.cfapi.net", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.cfapi.net"),
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
	name := "cloudflare_zone.tf-acc-unicode-test-1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("tf-acc-unicode-test-1", "żółw.cfapi.net", "true", "false"),
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
	name := "cloudflare_zone.tf-acc-unicode-test-2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("tf-acc-unicode-test-2", "xn--w-uga1v8h.cfapi.net", "true", "false"),
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
	name := "cloudflare_zone.tf-acc-unicode-test-3"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("tf-acc-unicode-test-3", "żółw.cfapi.net", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "żółw.cfapi.net"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config:   testZoneConfig("tf-acc-unicode-test-3", "xn--w-uga1v8h.cfapi.net", "true", "false"),
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

func TestPlanNameFallsBackToEmptyIfUnknown(t *testing.T) {
	type args struct {
		planName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"free website", args{"Free Website"}, planIDFree,
		},
		{
			"enterprise", args{"Enterprise Website"}, planIDEnterprise,
		},
		{
			"undefined or new", args{"New Awesome Plan Website"}, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := planIDForName(tt.args.planName); got != tt.want {
				t.Errorf("planIDForName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanIDFallsBackToEmptyIfUnknown(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"free id", args{planIDFree}, "Free Website"},
		{"enterprise id", args{planIDEnterprise}, "Enterprise Website"},
		{"unknonw id", args{"bogus"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := planNameForID(tt.args.id); got != tt.want {
				t.Errorf("planNameForID() = %v, want %v", got, tt.want)
			}
		})
	}
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
