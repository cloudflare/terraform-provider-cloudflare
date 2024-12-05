package snippet_rules_test

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
	resource.AddTestSweepers("cloudflare_snippet_rules", &resource.Sweeper{
		Name: "cloudflare_snippet_rules",
		F:    testSweepCloudflareSnippetRules,
	})
}

func testSweepCloudflareSnippetRules(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zone := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zone == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	_, err := client.UpdateZoneSnippetsRules(context.Background(), cloudflare.ZoneIdentifier(zone), []cloudflare.SnippetRule{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to disable Cloudflare Zone Snippet Rules: %s", err))
	}

	return nil
}

func TestAccCloudflareSnippetRules(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet_rules." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSnippetRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "some description 1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.snippet_name", "test_snippet_0"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "some description 2"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.snippet_name", "test_snippet_1"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "some description 3"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.snippet_name", "test_snippet_2"),
				),
			},
		},
	})
}

func testAccCheckCloudflareSnippetRules(rnd, zoneID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_snippet" "%[1]s" {
		count = 3
		zone_id  = "%[2]s"
		name = "test_snippet_${count.index}"
		main_module = "file1.js"
		files {
			name = "file1.js"
			content = "export default {async fetch(request) {return fetch(request)}};"
		}
	}

  resource "cloudflare_snippet_rules" "%[1]s" {
		zone_id  = "%[2]s"
		rules {
			enabled = true
			expression = "true"
			description = "some description 1"
			snippet_name = "test_snippet_0"
		}

		rules {
			enabled = true
			expression = "true"
			description = "some description 2"
			snippet_name = "test_snippet_1"
		}

		rules {
			enabled = true
			expression = "true"
			description = "some description 3"
			snippet_name = "test_snippet_2"
		}

		depends_on = ["cloudflare_snippet.%[1]s"]
  }`, rnd, zoneID)
}
