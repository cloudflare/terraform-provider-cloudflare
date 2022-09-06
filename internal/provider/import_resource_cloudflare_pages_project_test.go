package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflarePagesProject_Import(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_pages_project.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDirectUpload(rnd, accountID),
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_overwrite"},
			},
		},
	})
}
