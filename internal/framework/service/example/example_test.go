package example

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_turnstile_widget", &resource.Sweeper{
		Name: "cloudflare_turnstile_widget",
		F: func(region string) error {
			// Clean up left over resources from either manually touching the UI
			// or leaked test results.

			return nil
		},
	})
}

// func TestAccExampleResource(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { provider.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read testing
// 			{
// 				Config: testAccExampleResourceConfig("one"),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "one"),
// 					resource.TestCheckResourceAttr("scaffolding_example.test", "id", "example-id"),
// 				),
// 			},
// 			// ImportState testing
// 			{
// 				ResourceName:      "scaffolding_example.test",
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				// This is not normally necessary, but is here because this
// 				// example code does not have an actual upstream service.
// 				// Once the Read method is able to refresh information from
// 				// the upstream service, this can be removed.
// 				ImportStateVerifyIgnore: []string{"configurable_attribute"},
// 			},
// 			// Update and Read testing
// 			{
// 				Config: testAccExampleResourceConfig("two"),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "two"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func testAccExampleResourceConfig(configurableAttribute string) string {
// 	return fmt.Sprintf(`
// resource "scaffolding_example" "test" {
//   configurable_attribute = %[1]q
// }
// `, configurableAttribute)
// }
