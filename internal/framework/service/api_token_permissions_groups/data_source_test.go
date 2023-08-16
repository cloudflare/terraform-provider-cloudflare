package api_token_permissions_groups_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareApiTokenPermissionGroups_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// permission groups endpoint does not yet support the API tokens, and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareApiTokenPermissionGroupsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareApiTokenPermissionGroups("data.cloudflare_api_token_permission_groups.some"),
				),
			},
		},
	})
}

func testAccCloudflareApiTokenPermissionGroups(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		permCount, err := strconv.Atoi(a["permissions.%"])
		if err != nil {
			return fmt.Errorf("failed to convert total permission count to integer")
		}

		if permCount < 100 {
			return fmt.Errorf("total API token permission groups size is too small. expected: > 100, got: %d", permCount)
		}

		zonePermCount, err := strconv.Atoi(a["zone.%"])
		if err != nil {
			return fmt.Errorf("failed to convert zone permission count to integer")
		}

		if zonePermCount < 50 {
			return fmt.Errorf("zone API token permission groups size is too small. expected: > 50, got: %d", zonePermCount)
		}

		accountPermCount, err := strconv.Atoi(a["account.%"])
		if err != nil {
			return fmt.Errorf("failed to convert account permission count to integer")
		}

		if accountPermCount < 80 {
			return fmt.Errorf("account API token permission groups size is too small. expected: > 80, got: %d", accountPermCount)
		}

		userPermCount, err := strconv.Atoi(a["user.%"])
		if err != nil {
			return fmt.Errorf("failed to convert user permission count to integer")
		}

		if userPermCount < 5 {
			return fmt.Errorf("user API token permission groups size is too small. expected: > 5, got: %d", userPermCount)
		}
		r2PermCount, err := strconv.Atoi(a["r2.%"])
		if err != nil {
			return fmt.Errorf("failed to convert r2 permission count to integer")
		}
		if r2PermCount < 2 {
			return fmt.Errorf("r2 API token permission groups size is too small. expected: > 2, got: %d", r2PermCount)
		}

		apiTokenReadId, ok := a["permissions.API Tokens Read"]
		if !ok {
			return fmt.Errorf("couldn't get 'API Tokens Read' permission ID")
		}

		apiTokenReadIdShouldBe := "0cc3a61731504c89b99ec1be78b77aa0"

		if apiTokenReadId != apiTokenReadIdShouldBe {
			return fmt.Errorf("ApiTokenPermissionGroups 'API Tokens Read' is '%s', but should be '%s'",
				apiTokenReadId,
				apiTokenReadIdShouldBe,
			)
		}

		return nil
	}
}

const testAccCloudflareApiTokenPermissionGroupsConfig = `
data "cloudflare_api_token_permission_groups" "some" {}
`
