package sdkv2provider

import (
	"fmt"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessApplicationDataSourceAccount_Name(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccount_Name(accountID, rnd, domain),
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

func testAccCheckCloudflareAccessApplicationAccount_Name(accountID, name, domain string) string {
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

func TestAccCloudflareAccessApplicationDataSourceAccount_Domain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccount_Domain(accountID, rnd, domain),
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

func testAccCheckCloudflareAccessApplicationAccount_Domain(accountID, name, domain string) string {
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

func TestAccCloudflareAccessApplicationDataSourceZone_Name(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZone_Name(zoneID, rnd, domain),
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

func testAccCheckCloudflareAccessApplicationZone_Name(zoneID, name, domain string) string {
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

func TestAccCloudflareAccessApplicationDataSource_Zone_Domain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZone_Domain(zoneID, rnd, domain),
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

func testAccCheckCloudflareAccessApplicationZone_Domain(zoneID, name, domain string) string {
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
