package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareZone(t *testing.T) {
	name := "cloudflare_zone.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("test", "example.org", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.org"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config: testZoneConfigWithPlan("test", "example.org", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.org"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "full"),
				),
			},
			{
				Config: testZoneConfigWithPartialSetup("test", "example.org", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.org"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", planIDFree),
					resource.TestCheckResourceAttr(name, "type", "partial"),
				),
			},
			{
				Config: testZoneConfigWithExplicitFullSetup("test", "example.org", "true", "false", "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.org"),
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
