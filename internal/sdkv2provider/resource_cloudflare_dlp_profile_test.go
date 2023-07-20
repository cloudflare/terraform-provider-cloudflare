package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "custom profile"),
					resource.TestCheckResourceAttr(name, "type", "custom"),
					resource.TestCheckResourceAttr(name, "allowed_match_count", "0"),
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
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "allowed_match_count", "0"),
					resource.TestCheckResourceAttr(name, "description", "custom profile 2"),
					resource.TestCheckResourceAttr(name, "type", "custom"),

					resource.TestCheckTypeSetElemNestedAttrs(name, "entry.*", map[string]string{
						"name":                 fmt.Sprintf("%s_entry2", rnd),
						"enabled":              "true",
						"pattern.0.regex":      "^3[0-9]",
						"pattern.0.validation": "luhn",
					}),

					resource.TestCheckTypeSetElemNestedAttrs(name, "entry.*", map[string]string{
						"name":                 fmt.Sprintf("%s_entry1", rnd),
						"enabled":              "true",
						"pattern.0.regex":      "^4[0-9]",
						"pattern.0.validation": "luhn",
					}),
				),
			},
		},
	})
}

func TestAccCloudflareDLPProfile_CustomWithAllowedMatchCount(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dlp_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDLPProfileConfigCustomWithAllowedMatchCount(accountID, rnd, "custom profile", 42),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "custom profile"),
					resource.TestCheckResourceAttr(name, "allowed_match_count", "42"),
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

func testAccCloudflareDLPProfileConfigCustom(accountID, rnd, description string) string {
	return fmt.Sprintf(`
resource "cloudflare_dlp_profile" "%[1]s" {
  account_id                = "%[3]s"
  name                      = "%[1]s"
  description               = "%[2]s"
  type                      = "custom"
  allowed_match_count       = 0
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
  account_id                = "%[3]s"
  name                      = "%[1]s"
  description               = "%[2]s"
  allowed_match_count       = 0
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

func testAccCloudflareDLPProfileConfigCustomWithAllowedMatchCount(accountID, rnd, description string, allowedMatchCount uint) string {
	return fmt.Sprintf(`
resource "cloudflare_dlp_profile" "%[1]s" {
  account_id                = "%[3]s"
  name                      = "%[1]s"
  description               = "%[2]s"
  allowed_match_count       = %[4]d
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
`, rnd, description, accountID, allowedMatchCount)
}
