package snippets_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
	resourceName := "cloudflare_snippets." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSnippetsConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.StringExact("export { async function fetch(request, env) { return new Response('Hello World!'); } }")),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareSnippetsConfigUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.StringExact("export { async function fetch(request, env) { return new Response('Hello Updated World!'); } }")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata.main_module"), knownvalue.StringExact("main.js")),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
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
