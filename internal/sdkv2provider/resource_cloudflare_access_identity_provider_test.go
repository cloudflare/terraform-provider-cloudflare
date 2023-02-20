package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_identity_provider", &resource.Sweeper{
		Name: "cloudflare_access_identity_provider",
		F:    testSweepCloudflareAccessIdentityProviders,
	})
}

func testSweepCloudflareAccessIdentityProviders(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accessIDPs, accessIDPsErr := client.AccessIdentityProviders(context.Background(), accountID)
	if accessIDPsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Access Identity Providers: %s", accessIDPsErr))
	}

	if len(accessIDPs) == 0 {
		log.Print("[DEBUG] No Access Identity Providers to sweep")
		return nil
	}

	for _, idp := range accessIDPs {
		tflog.Info(ctx, fmt.Sprintf("Deleting Access Identity Provider ID: %s", idp.ID))
		_, err := client.DeleteAccessIdentityProvider(context.Background(), accountID, idp.ID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Access Identity Provider (%s): %s", idp.ID, err))
		}
	}

	return nil
}

func TestAccCloudflareAccessIdentityProvider_OneTimePin(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the OTP Access
	// endpoint does not yet support the API tokens for updates and it results in
	// state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, AccessIdentifier{Type: AccountType, Value: accountID}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "onetimepin"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, AccessIdentifier{Type: ZoneType, Value: zoneID}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "onetimepin"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OAuth(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "github"),
					resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.client_secret"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OAuthWithUpdate(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "github"),
					resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.client_secret"),
				),
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuthUpdatedName(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "github"),
					resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.client_secret"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SAML(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderSAML(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "saml"),
					resource.TestCheckResourceAttr(resourceName, "config.0.issuer_url", "jumpcloud"),
					resource.TestCheckResourceAttr(resourceName, "config.0.sso_target_url", "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"),
					resource.TestCheckResourceAttr(resourceName, "config.0.sign_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "config.0.attributes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "config.0.attributes.0", "email"),
					resource.TestCheckResourceAttr(resourceName, "config.0.attributes.1", "username"),
					resource.TestCheckResourceAttrSet(resourceName, "config.0.idp_public_cert"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderOneTimePin(name string, identifier AccessIdentifier) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  %[2]s_id = "%[3]s"
  name     = "%[1]s"
  type     = "onetimepin"
}`, name, identifier.Type, identifier.Value)
}

func testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, name string) string {
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

func testAccCheckCloudflareAccessIdentityProviderOAuthUpdatedName(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s-updated"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
  }
}`, accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderSAML(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "saml"
  config {
    issuer_url = "jumpcloud"
    sso_target_url = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes = [ "email", "username" ]
    sign_request = false
    idp_public_cert = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEG\nA1UEC…..GF/Q2/MHadws97cZg\nuTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
	}
}`, accountID, name)
}
