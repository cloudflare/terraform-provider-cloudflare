package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_identity_provider", &resource.Sweeper{
		Name: "cloudflare_access_identity_provider",
		F:    testSweepCloudflareAccessIdentityProviders,
	})
}

func testSweepCloudflareAccessIdentityProviders(r string) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client, clientErr := sharedClient()
	if clientErr != nil {
		log.Printf("[ERROR] Failed to create Cloudflare client: %s", clientErr)
	}

	accessIDPs, accessIDPsErr := client.ListZones()
	if accessIDPsErr != nil {
		log.Printf("[ERROR] Failed to fetch Access Identity Providers: %s", accessIDPsErr)
	}

	if len(accessIDPs) == 0 {
		log.Print("[DEBUG] No Access Identity Providers to sweep")
		return nil
	}

	for _, idp := range accessIDPs {
		log.Printf("[INFO] Deleting Access Identity Provider ID: %s", idp.ID)
		_, err := client.DeleteAccessIdentityProvider(accountID, idp.ID)

		if err != nil {
			log.Printf("[ERROR] Failed to delete Access Identity Provider (%s): %s", idp.ID, err)
		}
	}

	return nil
}

func TestAccCloudflareAccessIdentityProviderOneTimePin(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "onetimepin"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProviderOAuth(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOPAuth(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "github"),
					resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.client_secret"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderOneTimePin(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "onetimepin"
}`, accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderOPAuth(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
  }
}`, accountID, name)
}
