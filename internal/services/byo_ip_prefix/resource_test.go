package byo_ip_prefix_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareBYOIPPrefix(t *testing.T) {
	t.Parallel()
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_byo_ip_prefix.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_BYOIPPrefix(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareBYOIPPrefixConfig(prefixID, "BYOIP Prefix Description old", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "description", "BYOIP Prefix Description old"),
				),
			},
			{
				Config: testAccCheckCloudflareBYOIPPrefixConfig(prefixID, "BYOIP Prefix Description new", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "description", "BYOIP Prefix Description new"),
				),
			},
		},
	})
}

func testAccCheckCloudflareBYOIPPrefixConfig(prefixID, description, name string) string {
	return acctest.LoadTestCase("byoipprefixconfig.tf", prefixID, description, name)
}
