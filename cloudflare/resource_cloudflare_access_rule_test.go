package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccessRuleASN(t *testing.T) {
	name := "cloudflare_access_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(name, "challenge", "this is notes", "asn", "AS112"),
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

func testAccessRuleAccountConfig(resourceID, mode, notes, target, value string) string {
	return fmt.Sprintf(`
				resource "cloudflare_access_rule" "%[1]s" {
					notes = "%[3]s"
					mode = "%[2]s"
					configuration {
					  target = "%[4]s"
					  value = "%[5]s"
					}
				}`, resourceID, mode, notes, target, value)
}
