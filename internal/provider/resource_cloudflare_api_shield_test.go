package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAPIShield_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSingleEntry(rnd, zoneID, cloudflare.AuthIdCharacteristics{Name: "test-header", Type: "header"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "test-header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "header"),
				),
			},
		},
	})
}

func testAccCloudflareAPIShieldSingleEntry(resourceName, rnd string, authChar cloudflare.AuthIdCharacteristics) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield" "%[1]s" {
		zone_id = "%[2]s"
		auth_id_characteristics {
			name = "%[3]s"
			type = "%[4]s"
		}
	}
`, resourceName, rnd, authChar.Name, authChar.Type)
}
