package snippets_test

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
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_snippet", &resource.Sweeper{
		Name: "cloudflare_snippet",
		F:    testSweepCloudflareSnippet,
	})
}

func testSweepCloudflareSnippet(_ string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zone := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zone == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	err := client.DeleteZoneSnippet(context.Background(), cloudflare.ZoneIdentifier(zone), "test_snippet")
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to disable Cloudflare Zone Snippet: %s", err))
	}

	return nil
}

func TestAccCloudflareSnippet(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSnippet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "name", "test_snippet"),
				),
			},
		},
	})
}

func testAccCheckCloudflareSnippet(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_snippet" "%[1]s" {
	zone_id  = "%[2]s"
	name = "test_snippet"
	main_module = "file1.js"
	files {
		name = "file1.js"
		content = "export default {async fetch(request) {return fetch(request)}};"
	}

	files {
		name = "file2.js"
		content = "export default {async fetch(request) {return fetch(request)}};"
	}
  }`, rnd, zoneID)
}
