package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_filter", &resource.Sweeper{
		Name: "cloudflare_filter",
		F:    testSweepCloudflareFilterSweeper,
	})
}

func testSweepCloudflareFilterSweeper(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	filters, _, filtersErr := client.Filters(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.FilterListParams{})

	if filtersErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare filters: %s", filtersErr))
	}

	for _, filter := range filters {
		err := client.DeleteFilter(context.Background(), cloudflare.ZoneIdentifier(zoneID), filter.ID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare filter (%s) in zone ID: %s", filter.ID, zoneID))
		}
	}

	return nil
}

func TestAccFilterSimple(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`
	filterUnquoted := `(http.request.uri.path ~ ".*wp-login-` + rnd + `.php" or http.request.uri.path ~ ".*xmlrpc.php") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testFilterConfig(rnd, zoneID, "true", "this is notes", filterQuoted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "expression", filterUnquoted),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func testFilterConfig(resourceID, zoneID, paused, description, expression string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		`, resourceID, zoneID, paused, description, expression)
}

const multiLineFilter = `
resource "cloudflare_filter" "%[1]s" {
	zone_id = "%[2]s"
	paused = "%[3]s"
	description = "%[4]s"
	expression = <<EOF
%[5]s
EOF
}
`

func TestAccFilterWhitespace(t *testing.T) {
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(multiLineFilter, rnd, zoneID, "true", "multi-line filter",
					"\t\nhttp.request.method in {\"PUT\" \"DELETE\"} and\nhttp.request.uri.path eq \"/\"  \n"),
			},
		},
	})
}

func TestAccFilterHTMLEntity(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filter := `(http.host eq \"` + domain + `\")`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testFilterWithHTMLEntityConfig(rnd, zoneID, "true", "this is a 'test'", filter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is a \u0026#39;test\u0026#39;"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testFilterWithHTMLEntityConfig(resourceID, zoneID, paused, description, expression string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		`, resourceID, zoneID, paused, description, expression)
}
