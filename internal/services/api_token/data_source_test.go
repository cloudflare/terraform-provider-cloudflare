package api_token_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAPITokenData(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	permissionID := "82e64a83756745bbbb1c9c2701bf816b"
	dataSourceName := fmt.Sprintf("data.cloudflare_api_token.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneConfigBasic(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "status"),
					resource.TestCheckResourceAttr(dataSourceName, "status", "active"),
				),
			},
		},
	})
}

func testAccCloudflareZoneConfigBasic(name, permissionID string) string {
	return acctest.LoadTestCase("apitokendatasource.tf", name, permissionID)
}
