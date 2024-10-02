package observatory_scheduled_test_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareObservatoryScheduledTest_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_observatory_scheduled_test.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCloudflareWaitingRoomDestroy,
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
	return acctest.LoadTestCase("observatoryscheduledtest.tf", resourceName, zoneID, domain)
}
