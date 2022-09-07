package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflarePagesProject_Import(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_pages_project.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectBuildConfig(rnd, accountID, rnd, "project_owner", "project_repo"),
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
			},
		},
	})
}
