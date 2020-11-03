package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareWorkerSecret_Basic(t *testing.T) {
	t.Parallel()
	var secret cloudflare.WorkersSecret
	script_name := generateRandomResourceName()
	key := generateRandomResourceName()
	value := generateRandomResourceName()
	resourceName := "cloudflare_workers_secret." + script_name

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckAccount(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecret(script_name, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(script_name, key, secret),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
		},
	})
}

func testAccCloudflareWorkerSecretDestroy(s *terraform.State) error {
	 resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		secretResponse, err := client.ListWorkersSecrets(context.Background(), script_name)
		if err != nil {
			return err
		}

		for _, secret := range secretResponse.Result {
			if secret.Name == script_name {
				return nil
			}
		}

		return fmt.Errorf("worker secret key %s not found against worker script %s", key, script_name)
	}
}

func testAccCheckCloudflareWorkerSecret(script_name string, key string, value string) string {
	return fmt.Sprintf(`
	resource "cloudflare_worker_script" "cloudflare_worker_script.%[1]s" {
		name 	= ""cloudflare_worker_script.%[1]s""
	}

	resource "cloudflare_worker_secret" "cloudflare_worker_secret.%[1]s" {
		script_name = "cloudflare_worker_script.%[1]s"
		key 		= "%[2]s"
		value		= "%[3]s"
	}`, script_name, key, value)
}

func testAccCheckCloudflareWorkerSecretExists(script_name string, key string, secret cloudflare.WorkersSecret) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		secretResponse, err := client.ListWorkersSecrets(context.Background(), script_name)
		if err != nil {
			return err
		}

		for _, secret := range secretResponse.Result {
			if secret.Name == script_name {
				return nil
			}
		}

		return fmt.Errorf("worker secret key %s not found against worker script %s", key, script_name)
	}
}
