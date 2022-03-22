package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareAccessBookmarkBasic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_bookmark.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessBookmarkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessBookmarkConfigBasic(rnd, domain, AccessIdentifier{Type: ZoneType, Value: zoneID}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessBookmarkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessBookmarkConfigBasic(rnd, domain, AccessIdentifier{Type: AccountType, Value: accountID}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
		},
	})
}

func testAccCloudflareAccessBookmarkConfigBasic(rnd string, domain string, identifier AccessIdentifier) string {
	return fmt.Sprintf(`
resource "cloudflare_access_bookmark" "%[1]s" {
  %[3]s_id                  = "%[4]s"
  name                 = "%[1]s"
  domain               = "%[1]s.%[3]s"
  logo_url             = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
  app_launcher_visible = true
}
`, rnd, domain, identifier.Type, identifier.Value)
}

func testAccCheckCloudflareAccessBookmarkDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_bookmark" {
			continue
		}

		_, err := client.AccessBookmark(context.Background(), rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("AccessBookmark still exists")
		}

		_, err = client.AccessBookmark(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("AccessBookmark still exists")
		}
	}

	return nil
}

func TestAccCloudflareAccessBookmarkWithZoneID(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := "cloudflare_access_bookmark." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	updatedName := fmt.Sprintf("%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccessBookmarkWithZoneID(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
				),
			},
			{
				Config: testAccessBookmarkWithZoneIDUpdated(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", updatedName),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
				),
			},
		},
	})
}

func testAccessBookmarkWithZoneID(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_bookmark" "%[1]s" {
      name     = "%[1]s"
      zone_id  = "%[3]s"
      domain   = "%[1]s.%[2]s"
			logo_url = "https://image.com/img"
    }
  `, resourceID, zone, zoneID)
}

func testAccessBookmarkWithZoneIDUpdated(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_bookmark" "%[1]s" {
      name     = "%[1]s-updated"
      zone_id  = "%[3]s"
      domain   = "%[1]s.%[2]s"
			logo_url = "https://image.com/img"
    }
  `, resourceID, zone, zoneID)
}
