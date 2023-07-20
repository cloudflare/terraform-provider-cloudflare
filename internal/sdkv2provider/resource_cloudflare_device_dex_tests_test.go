package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDeviceDexTest_Traceroute(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_dex_test.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "interval", "0h30m0s"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "data.0.host", "dash.cloudflare.com"),
					resource.TestCheckResourceAttr(name, "data.0.kind", "traceroute"),
				),
			},
		},
	})
}

func TestAccCloudflareDeviceDexTest_TracerouteIPv4(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_dex_test.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsTracerouteIpv4(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "interval", "0h30m0s"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "data.0.host", "1.1.1.1"),
					resource.TestCheckResourceAttr(name, "data.0.kind", "traceroute"),
				),
			},
		},
	})
}

func TestAccCloudflareDeviceDexTest_HTTP(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_dex_test.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsHttp(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "interval", "0h30m0s"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "data.0.host", "https://dash.cloudflare.com/home"),
					resource.TestCheckResourceAttr(name, "data.0.kind", "http"),
					resource.TestCheckResourceAttr(name, "data.0.method", "GET"),
				),
			},
		},
	})
}

func testAccCloudflareDeviceDexTestsHttp(accountID, rnd string) string {
	return fmt.Sprintf(`
	resource "cloudflare_device_dex_test" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "%[1]s"
		interval = "0h30m0s"
		enabled = true
		data {
			host = "https://dash.cloudflare.com/home"
			kind = "http"
			method = "GET"
		}
	}
	`, rnd, accountID)
}

func testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd string) string {
	return fmt.Sprintf(`
	resource "cloudflare_device_dex_test" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "%[1]s"
		interval = "0h30m0s"
		enabled = true
		data {
			host = "dash.cloudflare.com"
			kind = "traceroute"
		}
	}
	`, rnd, accountID)
}

func testAccCloudflareDeviceDexTestsTracerouteIpv4(accountID, rnd string) string {
	return fmt.Sprintf(`
	resource "cloudflare_device_dex_test" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "%[1]s"
		interval = "0h30m0s"
		enabled = true
		data {
			host = "1.1.1.1"
			kind = "traceroute"
		}
	}
	`, rnd, accountID)
}
