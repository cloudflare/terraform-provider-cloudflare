package zero_trust_device_posture_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestAccCloudflareDevicePostureRules_DataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("data.cloudflare_zero_trust_device_posture_rules.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	return acctest.LoadTestCase("deviceposturerulesconfig.tf", name, accountID)
}
