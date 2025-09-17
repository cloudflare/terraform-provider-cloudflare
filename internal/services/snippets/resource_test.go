package snippets_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_snippets", &resource.Sweeper{
		Name: "cloudflare_snippets",
		F:    testSweepCloudflareSnippets,
	})
}

func testSweepCloudflareSnippets(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if zoneID == "" {
		// Skip sweeping if no zone ID is set
		return nil
	}

	// List all snippets in the zone
	list, err := client.Snippets.List(ctx, snippets.SnippetListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		return fmt.Errorf("failed to list snippets: %w", err)
	}

	// Delete all snippets in the test zone
	// Note: In a test environment, we assume all snippets can be deleted
	for list != nil {
		for _, snippet := range list.Result {
			_, err := client.Snippets.Delete(ctx, snippet.SnippetName, snippets.SnippetDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				// Log but continue sweeping other snippets
				continue
			}
		}

		list, err = list.GetNextPage()
		if err != nil {
			break
		}
	}

	return nil
}

func TestAccCloudflareSnippets_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareSnippetsConfig(rnd, zoneID),
				ExpectError: regexp.MustCompile("use 'cloudflare_snippet' instead of 'cloudflare_snippets'"),
			},
		},
	})
}

func testAccCheckCloudflareSnippetsDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_snippets" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		snippetName := rs.Primary.Attributes["snippet_name"]

		_, err := client.Snippets.Get(context.Background(), snippetName, snippets.SnippetGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err == nil {
			return fmt.Errorf("snippet still exists")
		}
	}

	return nil
}

func testAccCloudflareSnippetsConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareSnippetsConfigUpdate(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, zoneID)
}
