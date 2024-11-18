package api_token

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var tokenID = "ab577ebc55ec0425cfdce9cdbfe45cbe"

func TestAccCloudflareAPITokenData(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("data.cloudflare_api_token.%s", "all")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "name"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "status"),
					resource.TestCheckResourceAttr(name, "status", "active"),
					resource.TestCheckResourceAttr(name, "token_id", tokenID),
				),
			},
		},
	})
}

func testAccCloudflareZoneConfigBasic() string {
	return fmt.Sprintf(`
data "cloudflare_api_token" "all" {
	token_id = "%s"
}
`, tokenID)
}
