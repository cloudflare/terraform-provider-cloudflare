package cloudflare

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudflareIPRanges(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudflareIPRangesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareIPRanges("data.cloudflare_ip_ranges.testing"),
				),
			},
		},
	})
}

func testAccCloudflareIPRanges(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			cidrBlockSizeV4 int
			cidrBlockSizeV6 int
			err             error
		)

		// v4
		if cidrBlockSizeV4, err = strconv.Atoi(a["cidr_blocks_v4.#"]); err != nil {
			return err
		}

		if cidrBlockSizeV4 < 10 {
			return fmt.Errorf("cidr_blocks_v4 seem suspiciously low: %d", cidrBlockSizeV4)
		}

		var cidrBlocksV4 sort.StringSlice = make([]string, cidrBlockSizeV4)

		for i := range make([]string, cidrBlockSizeV4) {

			block := a[fmt.Sprintf("cidr_blocks_v4.%d", i)]

			if _, _, err := net.ParseCIDR(block); err != nil {
				return fmt.Errorf("malformed CIDR block %s: %s", block, err)
			}

			cidrBlocksV4[i] = block

		}

		if !sort.IsSorted(cidrBlocksV4) {
			return fmt.Errorf("unexpected order of cidr_blocks: %s", cidrBlocksV4)
		}

		// v6
		if cidrBlockSizeV6, err = strconv.Atoi(a["cidr_blocks_v6.#"]); err != nil {
			return err
		}

		if cidrBlockSizeV6 < 5 {
			return fmt.Errorf("cidr_blocks_v6 seem suspiciously low: %d", cidrBlockSizeV6)
		}

		var cidrBlocksV6 sort.StringSlice = make([]string, cidrBlockSizeV6)

		for i := range make([]string, cidrBlockSizeV6) {

			block := a[fmt.Sprintf("cidr_blocks_v6.%d", i)]

			if _, _, err := net.ParseCIDR(block); err != nil {
				return fmt.Errorf("malformed CIDR block %s: %s", block, err)
			}

			cidrBlocksV6[i] = block

		}

		if !sort.IsSorted(cidrBlocksV6) {
			return fmt.Errorf("unexpected order of cidr_blocks: %s", cidrBlocksV6)
		}

		return nil
	}
}

const testAccCloudflareIPRangesConfig = `
data "cloudflare_ip_ranges" "testing" {}
`
