package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

func TestAccCloudflareDevicePostureRules_DataSource(t *testing.T) {
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("data.cloudflare_device_posture_rules.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRulesConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttrSet(name, "rules.0.id"),
					resource.TestCheckResourceAttr(name, "rules.0.name", rnd),
					resource.TestCheckResourceAttr(name, "rules.0.type", "file"),
					resource.TestCheckResourceAttr(name, "rules.0.description", "check for /dev/random"),
					resource.TestCheckResourceAttr(name, "rules.0.schedule", "1h"),
					resource.TestCheckResourceAttr(name, "rules.0.expiration", ""),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureRulesConfig(name, accountID string) string {
	return fmt.Sprintf(`resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name        = "%[1]s"
  type        = "file"
  description = "check for /dev/random"
  schedule    = "1h"

  match {
    platform = "linux"
  }

  input {
    path = "/dev/random"
  }
}

data "cloudflare_device_posture_rules" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  depends_on = [cloudflare_device_posture_rule.%[1]s]
}
`, name, accountID)
}
