package oauth_scope_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareOAuthScopes_Datasource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_oauth_scopes.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareOAuthScopesDataSourceConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "result.0.id"),
					resource.TestCheckResourceAttrSet(name, "result.0.name"),
					testAccCheckOAuthScopesContains(name, "user-details.read"),
					testAccCheckOAuthScopesContains(name, "teams.read"),
				),
			},
		},
	})
}

func testAccCheckOAuthScopesContains(resourceName, scopeID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		count, err := strconv.Atoi(rs.Primary.Attributes["result.#"])
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			if rs.Primary.Attributes[fmt.Sprintf("result.%d.id", i)] == scopeID {
				return nil
			}
		}

		return fmt.Errorf("OAuth scope %q not found in %s", scopeID, resourceName)
	}
}

func testAccCloudflareOAuthScopesDataSourceConfig(resourceID string) string {
	return acctest.LoadTestCase("data_source.tf", resourceID)
}
