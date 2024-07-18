package user_agent_blocking_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareUserAgentBlockingRule_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the UA
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_user_agent_blocking_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareUserAgentBlockingRule(rnd, zoneID, "js_challenge"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "mode", "js_challenge"),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ua"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "Mozilla"),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
		CheckDestroy: testAccCheckCloudflareUserAgentBlockingRulesDestroy,
	})
}

func testAccCloudflareUserAgentBlockingRule(rnd, zoneID, mode string) string {
	return acctest.LoadTestCase("useragentblockingrule.tf", rnd, zoneID, mode)
}

func testAccCheckCloudflareUserAgentBlockingRulesDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_user_agent_blocking_rule" {
			continue
		}

		_, err := client.UserAgentRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("User Agent Blocking Rule still exists")
		}
	}

	return nil
}
