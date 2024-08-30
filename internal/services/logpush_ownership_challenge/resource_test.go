package logpush_ownership_challenge_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareLogpushOwnershipChallenge(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_logpush_ownership_challenge." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	destinationConf := `gs://cf-terraform-provider-acct-test/ownership_challenges`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
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
	return acctest.LoadTestCase("cloudflarelogpushownershipchallengeconfig.tf", resourceID, zoneID, destinationConf)
}
