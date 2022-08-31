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

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSingleEntry(resourceID, rnd, cloudflare.AuthIdCharacteristics{Name: "test-header", Type: "header"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "zone_id", rnd),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
				),
			},
		},
	})
}

func testAccCloudflareAPIShieldSingleEntry(resourceName string, rnd string, authChar cloudflare.AuthIdCharacteristics) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_token" "%[1]s" {
		zone_id = "%s"
		auth_id_characteristics {
			name = "%s"
			type = "%s"
		}
	}
`, resourceName, rnd, authChar.Name, authChar.Type)
}
