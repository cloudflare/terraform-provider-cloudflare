package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAPIToken_Basic(t *testing.T) {

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
				),
			},
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd+"-updated", permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccAPIToken_AllowDeny(t *testing.T) {

	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, false),
			},
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, true),
			},
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, false),
			},
		},
	})
}

func TestAccAPIToken_DoesNotSetConditions(t *testing.T) {

	rnd := generateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithoutCondition(resourceName, rnd, permissionID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[2]s"

		policy {
			effect = "allow"
			permission_groups = [ "%[3]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}
	}
`, resourceName, rnd, permissionID)
}

func TestAccAPIToken_SetIndividualCondition(t *testing.T) {

	rnd := generateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithIndividualCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithIndividualCondition(rnd string, permissionID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[1]s"

		policy {
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}

		condition {
			request_ip {
				in = ["192.0.2.1/32"]
			}
		}
	}
`, rnd, permissionID)
}

func TestAccAPIToken_SetAllCondition(t *testing.T) {

	rnd := generateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithAllCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.not_in.0", "198.51.100.1/32"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithAllCondition(rnd string, permissionID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[1]s"

		policy {
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}

		condition {
			request_ip {
				in     = ["192.0.2.1/32"]
				not_in = ["198.51.100.1/32"]
			}
		}
	}
`, rnd, permissionID)
}

func testAPITokenConfigAllowDeny(resourceID, permissionID, zoneID string, allowAllZonesExceptOne bool) string {
	var add string
	if allowAllZonesExceptOne {
		add = fmt.Sprintf(`
    		policy {
			  effect = "deny"
			  permission_groups = [
			    "%[1]s",
			  ]
			  resources = {
			    "com.cloudflare.api.account.zone.%[2]s" = "*"
			  }
	    	}
	  `, permissionID, zoneID)
	}

	return fmt.Sprintf(`
		resource "cloudflare_api_token" "%[1]s" {
		  name = "%[1]s"

		  policy {
			effect = "allow"
			permission_groups = [
			  "%[2]s",
			]
			resources = {
			  "com.cloudflare.api.account.zone.*" = "*"
			}
		  }
		  %[3]s
		}
		`, resourceID, permissionID, add)
}

func TestAccAPIToken_TokenTTL(t *testing.T) {

	rnd := generateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithTTL(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr(name, "expires_on", "2032-01-01T00:00:00Z"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithTTL(rnd string, permissionID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[1]s"

		policy {
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}

		not_before = "2018-07-01T05:20:00Z"
		expires_on = "2032-01-01T00:00:00Z"
	}
`, rnd, permissionID)
}
