package api_token_permission_groups_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareApiTokenPermissionGroupsDataSource_Basic(t *testing.T) {
	resourceName := "data.cloudflare_api_token_permission_groups_list.dns_read"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApiTokenPermissionGroupsDataSourceConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "DNS Read"),
					resource.TestCheckResourceAttr(resourceName, "scope", "com.cloudflare.api.account.zone"),
				),
			},
		},
	})
}

func testAccApiTokenPermissionGroupsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "cloudflare_api_token_permission_groups_list" "dns_read" {
			name  = "DNS Read"
			scope = "com.cloudflare.api.account.zone"
		}`,
	)
}
