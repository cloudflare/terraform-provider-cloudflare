package logpush_ownership_challenge_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareLogpushOwnershipChallenge(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_logpush_ownership_challenge." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	destinationConf := `gs://cf-terraform-provider-acct-test/ownership_challenges`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushOwnershipChallengeConfig(rnd, zoneID, destinationConf),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "destination_conf", destinationConf),
					resource.TestCheckResourceAttrSet(name, "ownership_challenge_filename"),
				),
			},
		},
	})
}

func testCloudflareLogpushOwnershipChallengeConfig(resourceID, zoneID, destinationConf string) string {
	return fmt.Sprintf(`
		resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
		  zone_id = "%[2]s"
		  destination_conf = "%[3]s"
		}
		`, resourceID, zoneID, destinationConf)
}
