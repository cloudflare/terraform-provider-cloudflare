package byo_ip_prefix_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
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
			testAccPreCheckBYOIPPrefix(t)
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
	return fmt.Sprintf(`
  resource "cloudflare_byo_ip_prefix" "%[3]s" {
	  prefix_id = "%[1]s"
	  description = "%[2]s"
  }`, prefixID, description, name)
}
