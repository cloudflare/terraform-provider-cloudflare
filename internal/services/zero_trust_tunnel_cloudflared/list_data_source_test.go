package zero_trust_tunnel_cloudflared_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunnelDatasource_List(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zero_trust_tunnel_cloudflareds.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelDatasourceList(accID, rnd),
				Check:  resource.TestCheckResourceAttr(name, "result.0.name", rnd),
			},
		},
	})
}

func testAccCheckCloudflareTunnelDatasourceList(accID, name string) string {
	return acctest.LoadTestCase("datasource_list.tf", accID, name)
}
