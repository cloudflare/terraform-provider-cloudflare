package user_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareUserDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "cloudflare_user" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_user.test", "id"),
					resource.TestCheckResourceAttrSet("data.cloudflare_user.test", "email"),
					resource.TestCheckResourceAttrSet("data.cloudflare_user.test", "username"),
				),
			},
		},
	})
}
