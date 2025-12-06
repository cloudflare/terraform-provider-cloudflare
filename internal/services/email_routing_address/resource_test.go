package email_routing_address_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_address", &resource.Sweeper{
		Name: "cloudflare_email_routing_address",
		F: func(region string) error {
			client := acctest.SharedClient()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			ctx := context.Background()

			if accountID == "" {
				tflog.Info(ctx, "Skipping email routing addresses sweep: CLOUDFLARE_ACCOUNT_ID not set")
				return nil
			}

			addresses, err := client.EmailRouting.Addresses.List(ctx, email_routing.AddressListParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch email routing addresses: %s", err))
				return fmt.Errorf("failed to fetch email routing destination addresses: %w", err)
			}

			addressList := addresses.Result
			if len(addressList) == 0 {
				tflog.Info(ctx, "No email routing addresses to sweep")
				return nil
			}

			for _, address := range addressList {
				if !utils.ShouldSweepResource(address.Email) {
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleting email routing address: %s (%s) (account: %s)", address.Email, address.Tag, accountID))
				_, err := client.EmailRouting.Addresses.Delete(ctx, address.Tag, email_routing.AddressDeleteParams{
					AccountID: cloudflare.F(accountID),
				})
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete email routing address %s: %s", address.Email, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted email routing address: %s", address.Tag))
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
	return acctest.LoadTestCase("emailroutingaddress.tf", rnd, accountID)
}
