package zero_trust_access_application_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessApplicationDataSource_AccountName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_zero_trust_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountName(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountName(accountID, name, domain string) string {
	return acctest.LoadTestCase("data_accessapplicationaccountname.tf", name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_AccountDomain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_zero_trust_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountDomain(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountDomain(accountID, name, domain string) string {
	return acctest.LoadTestCase("data_accessapplicationaccountdomain.tf", name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_zero_trust_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneName(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneName(zoneID, name, domain string) string {
	return acctest.LoadTestCase("data_accessapplicationzonename.tf", name, zoneID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneDomain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_zero_trust_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, name, domain string) string {
	return acctest.LoadTestCase("data_accessapplicationzonedomain.tf", name, zoneID, domain)
}
