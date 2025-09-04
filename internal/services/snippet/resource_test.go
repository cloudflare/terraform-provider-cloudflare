package snippet_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
	resource.AddTestSweepers("cloudflare_snippet", &resource.Sweeper{
		Name: "cloudflare_snippet",
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

func TestAccCloudflareSnippet_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSnippetConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("main.js"),
							"content": knownvalue.StringExact(`export default {
  async fetch(request) {
    // Get the current timestamp
    const timestamp = Date.now();
    // Convert the timestamp to hexadecimal format
    const hexTimestamp = timestamp.toString(16);
    // Clone the request and add the custom header
    const modifiedRequest = new Request(request, {
        headers: new Headers(request.headers)
    });
    modifiedRequest.headers.set("X-Hex-Timestamp", hexTimestamp);
    // Pass the modified request to the origin
    const response = await fetch(modifiedRequest);
    return response;
  },
}
`),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"main_module": knownvalue.StringExact("main.js"),
					})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareSnippetConfigUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"name": knownvalue.StringExact("main.js"),
							"content": knownvalue.StringExact(`export default {
  async fetch(request) {
    return new Response('Hello, World!');
  }
}
`),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"main_module": knownvalue.StringExact("main.js"),
					})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCheckCloudflareSnippetDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_snippet" {
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

func testAccCloudflareSnippetConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareSnippetConfigUpdate(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, zoneID)
}
