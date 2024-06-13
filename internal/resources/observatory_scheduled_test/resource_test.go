package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareObservatoryScheduledTest_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_observatory_scheduled_test.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWaitingRoomDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareObservatoryScheduledTest(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "url", domain+"/"+rnd),
					resource.TestCheckResourceAttr(name, "region", "us-central1"),
					resource.TestCheckResourceAttr(name, "frequency", "DAILY"),
				),
			},
		},
	})
}

func testAccCheckCloudflareObservatoryScheduledTestDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_observatory_scheduled_test" {
			continue
		}

		_, err := client.GetObservatoryScheduledPageTest(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), cloudflare.GetObservatoryScheduledPageTestParams{
			URL:    rs.Primary.Attributes["url"],
			Region: rs.Primary.Attributes["region"],
		})
		if err == nil {
			return fmt.Errorf("observatory scheduled test still exists")
		}
	}

	return nil
}

func testAccCloudflareObservatoryScheduledTest(resourceName, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_observatory_scheduled_test" "%[1]s" {
  zone_id   = "%[2]s"
  url       = "%[3]s/%[1]s"
  region    = "us-central1"
  frequency = "DAILY"
}
`, resourceName, zoneID, domain)
}
