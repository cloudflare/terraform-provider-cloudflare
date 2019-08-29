package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAccessRuleASN(t *testing.T) {
	name := "cloudflare_access_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig("challenge", "this is notes", "asn", "AS112"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
		},
	})
}

func testAccessRuleAccountConfig(mode, notes, target, value string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_rule" "test" {
  notes = "%[2]s"
  mode = "%[1]s"
  configuration =	 {
    target = "%[3]s"
    value = "%[4]s"
  }
}`, mode, notes, target, value)
}
