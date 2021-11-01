package cloudflare

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_filter", &resource.Sweeper{
		Name: "cloudflare_filter",
		F:    testSweepCloudflareFilterSweeper,
	})
}

func testSweepCloudflareFilterSweeper(r string) error {
	client, clientErr := sharedClient()
	if clientErr != nil {
		log.Printf("[ERROR] Failed to create Cloudflare client: %s", clientErr)
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	filters, filtersErr := client.Filters(context.Background(), zoneID, cloudflare.PaginationOptions{})

	if filtersErr != nil {
		log.Printf("[ERROR] Failed to fetch Cloudflare filters: %s", filtersErr)
	}

	for _, filter := range filters {
		err := client.DeleteFilter(context.Background(), zoneID, filter.ID)

		if err != nil {
			log.Printf("[ERROR] Failed to delete Cloudflare filter (%s) in zone ID: %s", filter.ID, zoneID)
		}
	}

	return nil
}

func TestAccCloudflareFilter_Simple(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`
	filterUnquoted := `(http.request.uri.path ~ ".*wp-login-` + rnd + `.php" or http.request.uri.path ~ ".*xmlrpc.php") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testFilterConfig(rnd, zoneID, "true", "this is notes", filterQuoted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "expression", filterUnquoted),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareFilter_InvalidExpressionWithAPIToken(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN to confirm users depending on API
	// tokens can also work with the filter validation.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testFilterInvalidExpression(rnd, zoneID),
				ExpectError: regexp.MustCompile("filter expression is invalid"),
			},
		},
	})
}

func TestAccCloudflareFilter_InvalidExpressionWithAPIKey(t *testing.T) {
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testFilterInvalidExpression(rnd, zoneID),
				ExpectError: regexp.MustCompile("filter expression is invalid"),
			},
		},
	})
}

func testFilterInvalidExpression(resourceID, zoneID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "true"
		  description = "example"
		  expression = "(foo eq \"bar\")"
		}
		`, resourceID, zoneID)
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

func TestAccCloudflareFilter_Whitespace(t *testing.T) {
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(multiLineFilter, rnd, zoneID, "true", "multi-line filter",
					"\t\nhttp.request.method in {\"PUT\" \"DELETE\"} and\nhttp.request.uri.path eq \"/\"  \n"),
			},
		},
	})
}

func TestAccCloudflareFilter_HTMLEntity(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filter := `(http.host eq \"` + domain + `\")`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
