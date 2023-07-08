package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessApplicationDataSource_AccountName(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountName(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountName(accountID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_AccountDomain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountDomain(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountDomain(accountID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		domain = "%[1]s.%[3]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneName(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneName(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneName(zoneID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, zoneID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneDomain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		domain = "%[1]s.%[3]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, zoneID, domain)
}
