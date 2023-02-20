package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareTeamsList_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_list.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "SERIAL"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "items.#", "2"),
					resource.TestCheckResourceAttr(name, "items.0", "asdf-1234"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsList_LottaListItems(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_list.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBigItemCount(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "SERIAL"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "items.#", "1000"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsList_Reordered(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBasic(rnd, accountID),
			},
			{
				Config: testAccCloudflareTeamsListConfigReorderedItems(rnd, accountID),
			},
		},
	})
}

func testAccCloudflareTeamsListConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "My description"
	type        = "SERIAL"
	items       = ["asdf-1234", "asdf-5678"]
}
`, rnd, accountID)
}

func testAccCloudflareTeamsListConfigBigItemCount(rnd, accountID string) string {
	items := []string{}
	for i := 0; i < 1000; i++ {
		items = append(items, `"example-`+strconv.Itoa(i)+`"`)
	}

	return fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "My description"
	type        = "SERIAL"
	items       = [%[3]s]
}
`, rnd, accountID, strings.Join(items, ","))
}

func testAccCloudflareTeamsListConfigReorderedItems(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "My description"
	type        = "SERIAL"
	items       = ["asdf-5678", "asdf-1234"]
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsListDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_teams_list" {
			continue
		}

		identifier := cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey])
		_, err := client.GetTeamsList(context.Background(), identifier, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Teams List still exists")
		}
	}

	return nil
}
