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
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		tflog.Info(ctx, "Skipping snippets sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	list, err := client.Snippets.List(ctx, snippets.SnippetListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list snippets: %s", err))
		return fmt.Errorf("failed to list snippets: %w", err)
	}

	snippetCount := 0
	for list != nil {
		for _, snippet := range list.Result {
			if !utils.ShouldSweepResource(snippet.SnippetName) {
				continue
			}
			snippetCount++
			tflog.Info(ctx, fmt.Sprintf("Deleting snippet: %s (zone: %s)", snippet.SnippetName, zoneID))
			_, err := client.Snippets.Delete(ctx, snippet.SnippetName, snippets.SnippetDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete snippet %s: %s", snippet.SnippetName, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted snippet: %s", snippet.SnippetName))
		}

		list, err = list.GetNextPage()
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to get next page of snippets: %s", err))
			break
		}
	}

	if snippetCount == 0 {
		tflog.Info(ctx, "No snippets to sweep")
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
