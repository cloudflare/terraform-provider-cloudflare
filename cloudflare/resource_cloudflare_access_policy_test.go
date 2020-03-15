package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAccessPolicyServiceToken(t *testing.T) {
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
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyServiceTokenConfig(rnd, zone, zoneID, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "include.0.service_token.#", "1"),
				),
			},
		},
	})
}

func TestAccAccessPolicyAnyServiceToken(t *testing.T) {
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
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAnyServiceTokenConfig(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
				),
			},
		},
	})
}

func testAccessPolicyServiceTokenConfig(resourceID, zone, zoneID, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_application" "%[1]s" {
			name    = "%[1]s"
			zone_id = "%[3]s"
			domain  = "%[1]s.%[2]s"
		}

		resource "cloudflare_access_service_token" "%[1]s" {
			account_id = "%[4]s"
			name       = "%[1]s"
		}

		resource "cloudflare_access_policy" "%[1]s" {
			application_id = "${cloudflare_access_application.%[1]s.id}"
			name           = "%[1]s"
			zone_id        = "%[3]s"
			decision       = "non_identity"
			precedence     = "10"

			include {
				service_token = ["${cloudflare_access_service_token.%[1]s.id}"]
			}
		}

	`, resourceID, zone, zoneID, accountID)
}

func testAccessPolicyAnyServiceTokenConfig(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_application" "%[1]s" {
			name    = "%[1]s"
			zone_id = "%[3]s"
			domain  = "%[1]s.%[2]s"
		}

		resource "cloudflare_access_policy" "%[1]s" {
			application_id = "${cloudflare_access_application.%[1]s.id}"
			name           = "%[1]s"
			zone_id        = "%[3]s"
			decision       = "non_identity"
			precedence     = "10"

			include {
				any_valid_service_token = true
			}
		}

	`, resourceID, zone, zoneID)
}
