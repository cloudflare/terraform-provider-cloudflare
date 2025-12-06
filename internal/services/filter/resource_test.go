package filter_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}


func init() {
	resource.AddTestSweepers("cloudflare_filter", &resource.Sweeper{
		Name: "cloudflare_filter",
		F:    testSweepCloudflareFilterSweeper,
	})
}

func testSweepCloudflareFilterSweeper(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s",clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping filters sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}
	filters, _, filtersErr := client.Filters(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.FilterListParams{})

	if filtersErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare filters: %s",filtersErr))
		return filtersErr
	}
	if len(filters) == 0 {
		tflog.Info(ctx, "No Cloudflare filters to sweep")
		return nil
	}

	for _, filter := range filters {
		// Use standard filtering helper on the description field
		if !utils.ShouldSweepResource(filter.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting filter: %s (zone: %s)", filter.ID, zoneID))
		err := client.DeleteFilter(ctx, cloudflare.ZoneIdentifier(zoneID), filter.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete filter %s: %s", filter.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted filter: %s", filter.ID))
	}

	return nil
}

func TestAccFilterSimple(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`
	filterUnquoted := `(http.request.uri.path ~ ".*wp-login-` + rnd + `.php" or http.request.uri.path ~ ".*xmlrpc.php") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	return acctest.LoadTestCase("filterconfig.tf", resourceID, zoneID, paused, description, expression)
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
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(multiLineFilter, rnd, zoneID, "true", "multi-line filter",
					"\t\nhttp.request.method in {\"PUT\" \"DELETE\"} and\nhttp.request.uri.path eq \"/\"  \n"),
			},
		},
	})
}

func TestAccFilterHTMLEntity(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	filter := `(http.host eq \"` + domain + `\")`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testFilterWithHTMLEntityConfig(rnd, zoneID, "true", "this is a 'test'", filter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is a 'test'"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testFilterWithHTMLEntityConfig(resourceID, zoneID, paused, description, expression string) string {
	return acctest.LoadTestCase("filterwithhtmlentityconfig.tf", resourceID, zoneID, paused, description, expression)
}
