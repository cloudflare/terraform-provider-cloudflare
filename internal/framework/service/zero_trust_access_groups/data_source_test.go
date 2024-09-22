package zero_trust_access_groups_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessGroups_DataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareZeroTrustAccessGroupsDataSourceConfig(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_zero_trust_access_groups.this", "account_id"),
					resource.TestCheckResourceAttrSet("data.cloudflare_zero_trust_access_groups.this", "groups.#"),
				),
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustAccessGroupsDataSourceConfig(accountID string) string {
	return fmt.Sprintf(`
data "cloudflare_zero_trust_access_groups" "this" {
    account_id = "%s"
}
`, accountID)
}
