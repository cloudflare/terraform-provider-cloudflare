package zero_trust_device_posture_rule_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareDevicePostureRules_DataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("data.cloudflare_zero_trust_device_posture_rules.%s", rnd)

	var lastAddedResourceIDX string
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRulesConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						var err error
						rs, ok := state.RootModule().Resources[name]
						if !ok {
							err = fmt.Errorf("resource not found: %s", name)
							return err
						}
						result, ok := rs.Primary.Attributes["result.#"]
						if !ok {
							err = fmt.Errorf("result attribute not found in resource %s", name)
							return err
						}
						lastAddedResourceIDX = getResultIndex(result)
						resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID)
						resource.TestCheckResourceAttr(name, fmt.Sprintf("result.%s.name", lastAddedResourceIDX), rnd)
						resource.TestCheckResourceAttrSet(name, fmt.Sprintf("result.%s.id", lastAddedResourceIDX))
						resource.TestCheckResourceAttr(name, fmt.Sprintf("result.%s.type", lastAddedResourceIDX), "file")
						resource.TestCheckResourceAttr(name, fmt.Sprintf("result.%s.description", lastAddedResourceIDX), "check for /dev/random")
						resource.TestCheckResourceAttr(name, fmt.Sprintf("result.%s.schedule", lastAddedResourceIDX), "1h")
						resource.TestCheckNoResourceAttr(name, fmt.Sprintf("result.%s.expiration", lastAddedResourceIDX))
						return nil
					},
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureRulesConfig(name, accountID string) string {
	return acctest.LoadTestCase("deviceposturerulesconfig.tf", name, accountID)
}

func getResultIndex(idx string) string {
	length, err := strconv.Atoi(idx)
	if err != nil {
		return ""
	}
	return strconv.Itoa(length - 1)
}
