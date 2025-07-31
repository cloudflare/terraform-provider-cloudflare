package zero_trust_tunnel_warp_connector_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWARPConnectorCreateBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func TestAccWARPConnectorUpdateName(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	name1 := fmt.Sprintf("%s_1", rnd)
	name2 := fmt.Sprintf("%s_2", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func testAccCheckWARPConnectorBasic(accID, name string) string {
	return acctest.LoadTestCase("warp_connector_basic.tf", accID, name)
}

func testAccCheckWARPConnectorUpdateName(accID, resourceName, name string) string {
	return acctest.LoadTestCase("warp_connector_update_name.tf", accID, resourceName, name)
}
