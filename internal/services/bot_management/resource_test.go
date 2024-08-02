package bot_management_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareBotManagement_SBFM(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	sbfmConfig := cloudflare.BotManagement{
		EnableJS:                     cloudflare.BoolPtr(true),
		SBFMDefinitelyAutomated:      cloudflare.StringPtr("managed_challenge"),
		SBFMLikelyAutomated:          cloudflare.StringPtr("block"),
		SBFMVerifiedBots:             cloudflare.StringPtr("allow"),
		SBFMStaticResourceProtection: cloudflare.BoolPtr(false),
		OptimizeWordpress:            cloudflare.BoolPtr(true),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementSBFM(rnd, zoneID, sbfmConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_definitely_automated", "managed_challenge"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_likely_automated", "block"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_verified_bots", "allow"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_static_resource_protection", "false"),
					resource.TestCheckResourceAttr(resourceID, "optimize_wordpress", "true"),
				),
			},
			// {
			// 	ResourceName:      resourceID,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccCloudflareBotManagement_Unentitled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	bmEntConfig := cloudflare.BotManagement{
		EnableJS:             cloudflare.BoolPtr(true),
		SuppressSessionScore: cloudflare.BoolPtr(false),
		AutoUpdateModel:      cloudflare.BoolPtr(false),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementEntSubscription(rnd, zoneID, bmEntConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "suppress_session_score", "false"),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "false"),
				),
				ExpectError: regexp.MustCompile("zone not entitled to disable"),
			},
		},
	})
}

func testCloudflareBotManagementSBFM(resourceName, rnd string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsbfm.tf", resourceName, rnd,
		*bm.EnableJS, *bm.SBFMDefinitelyAutomated,
		*bm.SBFMLikelyAutomated, *bm.SBFMVerifiedBots,
		*bm.SBFMStaticResourceProtection, *bm.OptimizeWordpress)
}

func testCloudflareBotManagementEntSubscription(resourceName, rnd string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagemententsubscription.tf", resourceName, rnd, *bm.EnableJS, *bm.SuppressSessionScore, false)
}
