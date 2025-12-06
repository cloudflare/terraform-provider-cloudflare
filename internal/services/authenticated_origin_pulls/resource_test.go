package authenticated_origin_pulls_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls",
		F:    testSweepCloudflareAuthenticatedOriginPulls,
	})
}

func testSweepCloudflareAuthenticatedOriginPulls(r string) error {
	ctx := context.Background()
	// Authenticated Origin Pulls is a zone-level security setting (enabled/disabled).
	// It can be configured globally, per-zone, or per-hostname, but it's still a
	// singleton setting rather than creating multiple resources.
	// No sweeping required.
	tflog.Info(ctx, "Authenticated Origin Pulls doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflareAuthenticatedOriginPullsGlobal(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into correct test setup for reproducibility..")

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsGlobalConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAuthenticatedOriginPullsPerZone(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into correct test setup for reproducibility.")

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsPerZoneConfig(zoneID, rnd, "per-zone"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAuthenticatedOriginPullsPerHostname(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into correct test setup for reproducibility.")

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	hostname := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsConfig(zoneID, rnd, "per-hostname", hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAuthenticatedOriginPullsGlobalConfig(zoneID, name string) string {
	return acctest.LoadTestCase("authenticatedoriginpullsglobalconfig.tf", zoneID, name)
}

func testAccCheckCloudflareAuthenticatedOriginPullsPerZoneConfig(zoneID, name, aopType string) string {
	return acctest.LoadTestCase("authenticatedoriginpullsperzoneconfig.tf", name, zoneID, aopType)
}

func testAccCheckCloudflareAuthenticatedOriginPullsConfig(zoneID, name, aopType, hostname string) string {
	return acctest.LoadTestCase("authenticatedoriginpullsconfig.tf", name, zoneID, aopType, name+"."+hostname)
}
