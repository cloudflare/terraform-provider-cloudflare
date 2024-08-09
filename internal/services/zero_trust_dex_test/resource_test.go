package zero_trust_dex_test_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDeviceDexTest_Traceroute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_test.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_test.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsTracerouteIpv4(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_test.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsHttp(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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
	return acctest.LoadTestCase("devicedextestshttp.tf", rnd, accountID)
}

func testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedexteststraceroute.tf", rnd, accountID)
}

func testAccCloudflareDeviceDexTestsTracerouteIpv4(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedexteststracerouteipv4.tf", rnd, accountID)
}
