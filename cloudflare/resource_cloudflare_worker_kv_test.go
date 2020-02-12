package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareWorkersKV_Basic(t *testing.T) {
	t.Parallel()
	var kvPair cloudflare.WorkersKVPair
	namespaceID := generateRandomResourceName()
	key := generateRandomResourceName()
	id := fmt.Sprintf("%s_%s", namespaceID, key)
	resourceName := "cloudflare_workers_kv." + id
	value := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKV(resourceName, namespaceID, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(namespaceID, key, &kvPair),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
		},
	})
}

func testAccCloudflareWorkersKVDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_kv" {
			continue
		}

		namespaceID := rs.Primary.Attributes["namespace_id"]
		key := rs.Primary.Attributes["key"]

		_, err := client.ReadWorkersKV(context.Background(), namespaceID, key)

		if err == nil {
			return fmt.Errorf("workers kv pair still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersKV(rName string, namespaceID string, key string, value string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv" "%[1]s" {
	namespace_id = "%[2]s
	key = "%[3]s"
	value = "%[4]s"
}`, rName, namespaceID, key, value)
}

func testAccCheckCloudflareWorkersKVExists(namespaceID string, key string, kv *cloudflare.WorkersKVPair) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		value, err := client.ReadWorkersKV(context.Background(), namespaceID, key)
		if err != nil {
			return err
		}

		if value == nil {
			return fmt.Errorf("workers kv key %s not found in namespace %s", key, namespaceID)
		}

		return nil
	}
}
