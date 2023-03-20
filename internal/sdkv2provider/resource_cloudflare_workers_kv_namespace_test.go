package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWorkersKVNamespace_Basic(t *testing.T) {
	t.Parallel()
	var namespace cloudflare.WorkersKVNamespace
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespace(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, &namespace),
					resource.TestCheckResourceAttr(resourceName, "title", rnd),
				),
			},
		},
	})
}

func testAccCloudflareWorkersKVNamespaceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_kv_namespace" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		resp, _, err := client.ListWorkersKVNamespaces(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersKVNamespacesParams{})
		if err == nil {
			return err
		}

		for _, n := range resp {
			if n.ID == rs.Primary.ID {
				return fmt.Errorf("namespace still exists but should not")
			}
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersKVNamespace(rName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[2]s"
	title = "%[1]s"
}`, rName, accountID)
}

func testAccCheckCloudflareWorkersKVNamespaceExists(title string, namespace *cloudflare.WorkersKVNamespace) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		rs, ok := s.RootModule().Resources["cloudflare_workers_kv_namespace."+title]
		if !ok {
			return fmt.Errorf("not found: %s", title)
		}
		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		resp, _, err := client.ListWorkersKVNamespaces(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersKVNamespacesParams{})
		if err != nil {
			return err
		}

		for _, n := range resp {
			if n.Title == title {
				*namespace = n
				return nil
			}
		}

		return fmt.Errorf("namespace not found")
	}
}
