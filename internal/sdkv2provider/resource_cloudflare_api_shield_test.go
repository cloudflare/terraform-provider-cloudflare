package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAPIShield_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
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
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "test-header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "header"),
				),
			},
		},
	})
}

func TestAccAPIShield_EmptyAuthIdCharacteristics(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldEmptyAuthIdCharacteristics(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "0"),
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

func testAccCloudflareAPIShieldEmptyAuthIdCharacteristics(resourceName, rnd string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield" "%[1]s" {
		zone_id = "%[2]s"
	}
	`, resourceName, rnd)
}
