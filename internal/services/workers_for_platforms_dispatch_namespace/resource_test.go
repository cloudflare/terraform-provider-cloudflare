package workers_for_platforms_dispatch_namespace_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkersForPlatformsDispatchNamespace(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersForPlatformsDispatchNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("workersforplatformsnamespacemanagement.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				},
				Check: testAccCheckCloudflareWorkersForPlatformsDispatchNamespaceExists(name, accountID, rnd),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccCheckCloudflareWorkersForPlatformsDispatchNamespaceDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_for_platforms_dispatch_namespace" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		namespaceName := rs.Primary.Attributes["name"]

		_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Get(
			context.Background(),
			namespaceName,
			workers_for_platforms.DispatchNamespaceGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err == nil {
			return fmt.Errorf("dispatch namespace %s still exists", namespaceName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersForPlatformsDispatchNamespaceExists(resourceName, accountID, namespaceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no dispatch namespace ID is set")
		}

		client := acctest.SharedClient()
		_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Get(
			context.Background(),
			namespaceName,
			workers_for_platforms.DispatchNamespaceGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err != nil {
			return fmt.Errorf("dispatch namespace not found: %s", err)
		}

		return nil
	}
}
