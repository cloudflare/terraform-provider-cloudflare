package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFilterSimple(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_filter." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`
	filterUnquoted := `(http.request.uri.path ~ ".*wp-login-` + rnd + `.php" or http.request.uri.path ~ ".*xmlrpc.php") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testFilterConfig(rnd, zone, "true", "this is notes", filterQuoted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "expression", filterUnquoted),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
				),
			},
		},
	})
}

func testFilterConfig(resourceID, zoneName, paused, description, expression string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		`, resourceID, zoneName, paused, description, expression)
}

const multiLineFilter = `
resource "cloudflare_filter" "%[1]s" {
	zone = "%[2]s"
	paused = "%[3]s"
	description = "%[4]s"
	expression = <<EOF
%[5]s
EOF
}
`

func TestAccFilterWhitespace(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(multiLineFilter, rnd, zone, "true", "multi-line filter",
					"\t\nhttp.request.method in {\"PUT\" \"DELETE\"} and\nhttp.request.uri.path eq \"/\"  \n"),
			},
		},
	})
}
