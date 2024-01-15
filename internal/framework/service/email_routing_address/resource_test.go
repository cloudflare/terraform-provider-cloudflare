package email_routing_address_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_address", &resource.Sweeper{
		Name: "cloudflare_email_routing_address",
		F: func(region string) error {
			client, err := acctest.SharedClient()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			emails, _, err := client.ListEmailRoutingDestinationAddresses(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListEmailRoutingAddressParameters{})
			if err != nil {
				return fmt.Errorf("failed to fetch email routing destination addresses: %w", err)
			}

			for _, email := range emails {
				_, err := client.DeleteEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(accountID), email.Tag)
				if err != nil {
					return fmt.Errorf("failed to delete email routing destination address %q: %w", email.Email, err)
				}
			}

			return nil
		},
	})
}

// Uncomment to run email routing address test cases.
//
// See: https://github.com/hashicorp/terraform-plugin-testing/issues/85 why this
// isn't possible with the current service that doesn't allow immediate
// deletions.
//
// func TestAccCloudflareEmailRoutingAddress_Basic(t *testing.T) {
// 	rnd := utils.GenerateRandomResourceName()
// 	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
// 	resourceName := "cloudflare_email_routing_address." + rnd

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckCloudflareEmailRoutingAddress(rnd, accountID),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrSet(resourceName, "tag"),
// 					resource.TestCheckResourceAttr(resourceName, "email", rnd+"@example.com"),
// 				),
// 			},
// 			{
// 				ResourceName:        resourceName,
// 				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
// 				ImportState:         true,
// 				ImportStateVerify:   true,
// 			},
// 		},
// 	})
// }

func testAccCheckCloudflareEmailRoutingAddress(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_email_routing_address" "%[1]s" {
    account_id = "%[2]s"
    email      = "%[1]s@example.com"
  }`, rnd, accountID)
}
