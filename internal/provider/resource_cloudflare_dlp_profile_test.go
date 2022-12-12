package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareDLPProfile_Custom(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dlp_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDLPProfileConfigCustom(accountID, rnd, "custom profile"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "custom profile"),
					resource.TestCheckResourceAttr(name, "type", "custom"),
					resource.TestCheckResourceAttr(name, "entry.0.name", fmt.Sprintf("%s_entry1", rnd)),
					resource.TestCheckResourceAttr(name, "entry.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.regex", "^4[0-9]"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.validation", "luhn"),
				),
			},
		},
	})
}

func TestAccCloudflareDLPProfile_Custom_MultipleEntries(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dlp_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDLPProfileConfigCustomMultipleEntries(accountID, rnd, "custom profile 2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "custom profile 2"),
					resource.TestCheckResourceAttr(name, "type", "custom"),

					resource.TestCheckResourceAttr(name, "entry.0.name", fmt.Sprintf("%s_entry2", rnd)),
					resource.TestCheckResourceAttr(name, "entry.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.regex", "^3[0-9]"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.validation", "luhn"),

					resource.TestCheckResourceAttr(name, "entry.1.name", fmt.Sprintf("%s_entry1", rnd)),
					resource.TestCheckResourceAttr(name, "entry.1.enabled", "true"),
					resource.TestCheckResourceAttr(name, "entry.1.pattern.0.regex", "^4[0-9]"),
					resource.TestCheckResourceAttr(name, "entry.1.pattern.0.validation", "luhn"),
				),
			},
		},
	})
}

func testAccCloudflareDLPProfileConfigCustom(accountID, rnd, description string) string {
	return fmt.Sprintf(`
resource "cloudflare_dlp_profile" "%[1]s" {
  account_id                = "%[3]s"
  name                      = "%[1]s"
  description               = "%[2]s"
  type                      = "custom"
  entry {
	name = "%[1]s_entry1"
	enabled = true
	pattern {
		regex = "^4[0-9]"
		validation = "luhn"
	}
  }
}
`, rnd, description, accountID)
}

func testAccCloudflareDLPProfileConfigCustomMultipleEntries(accountID, rnd, description string) string {
	return fmt.Sprintf(`
resource "cloudflare_dlp_profile" "%[1]s" {
  account_id                  = "%[3]s"
  name                      = "%[1]s"
  description               = "%[2]s"
  type                      = "custom"
  entry {
	name = "%[1]s_entry1"
	enabled = true
	pattern {
		regex = "^4[0-9]"
		validation = "luhn"
	}
  }

  entry {
	name = "%[1]s_entry2"
	enabled = true
	pattern {
		regex = "^3[0-9]"
		validation = "luhn"
	}
  }
}
`, rnd, description, accountID)
}
