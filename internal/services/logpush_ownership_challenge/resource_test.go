package logpush_ownership_challenge_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_logpush_ownership_challenge", &resource.Sweeper{
		Name: "cloudflare_logpush_ownership_challenge",
		F:    testSweepCloudflareLogpushOwnershipChallenges,
	})
}

func testSweepCloudflareLogpushOwnershipChallenges(r string) error {
	ctx := context.Background()
	// Logpush ownership challenges are read-only resources that don't create
	// persistent state that needs cleanup. They just generate ownership verification
	// tokens. No sweeping required.
	tflog.Info(ctx, "Logpush ownership challenges don't require sweeping (read-only resource)")
	return nil
}

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
					resource.TestCheckResourceAttrSet(name, "filename"),
				),
			},
		},
	})
}

func testCloudflareLogpushOwnershipChallengeConfig(resourceID, zoneID, destinationConf string) string {
	return acctest.LoadTestCase("cloudflarelogpushownershipchallengeconfig.tf", resourceID, zoneID, destinationConf)
}
