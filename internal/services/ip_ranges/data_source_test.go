package ip_ranges_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareIPRange_Networks(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_ip_ranges.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPRangesNetworks(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPRangeCount(name),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPRangesNetworks(rnd string) string {
	return acctest.LoadTestCase("networks.tf", rnd)
}

func testAccCheckCloudflareIPRangeCount(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		jdcloudCidrs, _ := strconv.Atoi(rs.Primary.Attributes["jdcloud_cidrs.#"])
		if jdcloudCidrs < 10 {
			return fmt.Errorf("jdcloud_cidrs size is suspiciously low. should be > 10, got: %d", jdcloudCidrs)
		}

		ipv4Cidrs, _ := strconv.Atoi(rs.Primary.Attributes["ipv4_cidrs.#"])
		if ipv4Cidrs < 10 {
			return fmt.Errorf("ipv4_cidrs size is suspiciously low. should be > 10, got: %d", ipv4Cidrs)
		}

		ipv6Cidrs, _ := strconv.Atoi(rs.Primary.Attributes["ipv6_cidrs.#"])
		if ipv6Cidrs < 5 {
			return fmt.Errorf("ipv6_cidrs size is suspiciously low. should be > 5, got: %d", ipv6Cidrs)
		}

		return nil
	}
}
