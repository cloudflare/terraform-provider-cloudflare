package zero_trust_access_identity_provider_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_identity_provider", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_identity_provider",
		F:    testSweepCloudflareAccessIdentityProviders,
	})
}

func testSweepCloudflareAccessIdentityProviders(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accessIDPs, _, accessIDPsErr := client.ListAccessIdentityProviders(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessIdentityProvidersParams{})
	if accessIDPsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Access Identity Providers: %s", accessIDPsErr))
	}

	if len(accessIDPs) == 0 {
		log.Print("[DEBUG] No Access Identity Providers to sweep")
		return nil
	}

	for _, idp := range accessIDPs {
		tflog.Info(ctx, fmt.Sprintf("Deleting Access Identity Provider ID: %s", idp.ID))
		_, err := client.DeleteAccessIdentityProvider(context.Background(), cloudflare.AccountIdentifier(accountID), idp.ID)

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

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "onetimepin"),
					resource.TestCheckResourceAttrWith(resourceName, "config.0.redirect_url", func(value string) error {
						if !strings.HasSuffix(value, ".cloudflareaccess.com/cdn-cgi/access/callback") {
							return fmt.Errorf("expected redirect_url to be a Cloudflare Access URL, got %s", value)
						}
						return nil
					}),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "onetimepin"),
					resource.TestCheckResourceAttrWith(resourceName, "config.0.redirect_url", func(value string) error {
						if !strings.HasSuffix(value, ".cloudflareaccess.com/cdn-cgi/access/callback") {
							return fmt.Errorf("expected redirect_url to be a Cloudflare Access URL, got %s", value)
						}
						return nil
					}),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OAuth(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

func TestAccCloudflareAccessIdentityProvider_AzureAD(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Pending investigation into automating Azure IDP.")

	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "azureAD"),
					resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
					resource.TestCheckResourceAttr(resourceName, "config.0.directory_id", "directory"),
					resource.TestCheckResourceAttr(resourceName, "config.0.condtional_access_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scim_config.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scim_config.0.user_deprovision", "true"),
					resource.TestCheckResourceAttr(resourceName, "scim_config.0.seat_deprovision", "true"),
					resource.TestCheckResourceAttr(resourceName, "scim_config.0.group_member_deprovision", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OAuth_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(resourceName, "name", rnd),
		resource.TestCheckResourceAttr(resourceName, "type", "github"),
		resource.TestCheckResourceAttr(resourceName, "config.0.client_id", "test"),
		resource.TestCheckResourceAttrSet(resourceName, "config.0.client_secret"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				Check:  checkFn,
			},
			// {
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// 	ResourceName:        resourceName,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			// 	Check:               checkFn,
			// },
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SCIM_Config_Secret(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrWith(resourceName, "scim_config.0.secret", func(value string) error {
			if value == "" {
				return errors.New("secret is empty")
			}

			if strings.Contains(value, "*") {
				return errors.New("secret was redacted")
			}

			return nil
		}),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				Check:  checkFn,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				Check:  checkFn,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SCIM_Secret_Enabled_After_Resource_Creation(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrWith(resourceName, "scim_config.0.secret", func(value string) error {
			if value == "" {
				return errors.New("secret is empty")
			}

			if strings.Contains(value, "*") {
				return errors.New("secret was redacted")
			}

			return nil
		}),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADNoSCIM(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scim_config.0.secret", ""),
				),
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				Check:  checkFn,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				Check:  checkFn,
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderOneTimePin(name string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessidentityprovideronetimepin.tf", name, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovideroauth.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderOAuthUpdatedName(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovideroauthupdatedname.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderSAML(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovidersaml.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityproviderazuread.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityproviderazureadupdated.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderAzureADNoSCIM(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityproviderazureadnoscim.tf", accountID, name)
}
