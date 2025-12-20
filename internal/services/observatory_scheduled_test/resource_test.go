package observatory_scheduled_test_test

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cfv3 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/speed"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_observatory_scheduled_test", &resource.Sweeper{
		Name: "cloudflare_observatory_scheduled_test",
		F:    testSweepCloudflareObservatoryScheduledTests,
	})
}

func testSweepCloudflareObservatoryScheduledTests(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping observatory scheduled tests sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	pages, err := client.Speed.Pages.List(ctx, speed.PageListParams{
		ZoneID: cfv3.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch observatory scheduled tests: %s", err))
		return fmt.Errorf("failed to fetch observatory scheduled tests: %w", err)
	}

	if len(pages.Result) == 0 {
		tflog.Info(ctx, "No observatory scheduled tests to sweep")
		return nil
	}

	for _, page := range pages.Result {
		// Use standard filtering helper on the URL
		if !utils.ShouldSweepResource(page.URL) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting observatory scheduled test: %s (region: %s, zone: %s)", page.URL, page.Region.Label, zoneID))
		_, err := client.Speed.Schedule.Delete(ctx, page.URL, speed.ScheduleDeleteParams{
			ZoneID: cfv3.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete observatory scheduled test %s: %s", page.URL, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted observatory scheduled test: %s", page.URL))
	}

	return nil
}

func TestAccCloudflareObservatoryScheduledTest_Basic(t *testing.T) {
	t.Skip("needs to be fixed by service team");
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_observatory_scheduled_test.%s", rnd)
	url1 := url.PathEscape(domain + "/" + rnd)
	url2 := url.PathEscape(domain + "/" + rnd + "-updated")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create + Read
			{
				Config: testAccCloudflareObservatoryScheduledTestConfig(rnd, zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New(consts.ZoneIDSchemaKey),
						knownvalue.StringExact(zoneID),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("url"),
						knownvalue.StringExact(url1),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("id"),
						knownvalue.StringExact(url1),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("schedule").AtMapKey("region"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("schedule").AtMapKey("frequency"),
						knownvalue.NotNull(),
					),
				},
			},
			// Step 2: Update + Read
			{
				Config: testAccCloudflareObservatoryScheduledTestConfig(rnd, zoneID, domain, rnd+"-updated"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New(consts.ZoneIDSchemaKey),
						knownvalue.StringExact(zoneID),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("url"),
						knownvalue.StringExact(url2),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("id"),
						knownvalue.StringExact(url2),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("schedule").AtMapKey("region"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("schedule").AtMapKey("frequency"),
						knownvalue.NotNull(),
					),
				},
			},
			// Step 3: Import
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/%s", zoneID, url2),
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

func testAccCloudflareObservatoryScheduledTestConfig(resourceName, zoneID, domain, path string) string {
	return fmt.Sprintf(`
resource "cloudflare_observatory_scheduled_test" "%[1]s" {
  zone_id = %[2]q
  url     = urlencode("%[3]s/%[4]s")
}
`, resourceName, zoneID, domain, path)
}
