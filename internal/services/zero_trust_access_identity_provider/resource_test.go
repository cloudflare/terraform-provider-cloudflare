package zero_trust_access_identity_provider_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.AccountIdentifier(accountID)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("onetimepin")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("redirect_url"), knownvalue.StringRegexp(regexp.MustCompile(`\.cloudflareaccess\.com/cdn-cgi/access/callback$`))),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("accounts/%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.AccountIdentifier(accountID)),
				PlanOnly: true,
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.ZoneIdentifier(zoneID)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("onetimepin")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("redirect_url"), knownvalue.StringRegexp(regexp.MustCompile(`\.cloudflareaccess\.com/cdn-cgi/access/callback$`))),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("zones/%s/", zoneID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOneTimePin(rnd, cloudflare.ZoneIdentifier(zoneID)),
				PlanOnly: true,
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				PlanOnly: true,
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret"},
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				PlanOnly: true,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuthUpdatedName(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
				},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOAuthUpdatedName(accountID, rnd),
				PlanOnly: true,
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderSAML(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saml")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("issuer_url"), knownvalue.StringExact("jumpcloud")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sso_target_url"), knownvalue.StringExact("https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes").AtSliceIndex(0), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes").AtSliceIndex(1), knownvalue.StringExact("username")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("idp_public_certs"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.sign_request"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderSAML(accountID, rnd),
				PlanOnly: true,
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("directory_id"), knownvalue.StringExact("directory")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("user_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("seat_deprovision"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.conditional_access_enabled", "scim_config.secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OAuth_Import(t *testing.T) {
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
				},
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				PlanOnly: true,
			},
			{
				Config:            testAccCheckCloudflareAccessIdentityProviderOAuth(accountID, rnd),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// cant import client_secret
					"config.client_secret",
				},
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("accounts/%s/", accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
				},
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SCIM_Config_Secret(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrWith(resourceName, "scim_config.secret", func(value string) error {
			if value == "" {
				return errors.New("secret is empty")
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				Check:  checkFn,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("directory_id"), knownvalue.StringExact("directory")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("user_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("seat_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.conditional_access_enabled", "scim_config.secret"},
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				PlanOnly: true,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				Check:  checkFn,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("directory_id"), knownvalue.StringExact("directory")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("user_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("seat_deprovision"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SCIM_Secret_Enabled_After_Resource_Creation(t *testing.T) {
	t.Skip("TODO: failing due to inconsistent apply caused by secret value")
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrWith(resourceName, "scim_config.secret", func(value string) error {
			if value == "" {
				return errors.New("secret is empty")
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
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADNoSCIM(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.Null()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.conditional_access_enabled"},
			},
			{
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADNoSCIM(accountID, rnd),
				PlanOnly: true,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				Check:  checkFn,
			},
			{
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				PlanOnly: true,
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				Check:  checkFn,
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OneTimePin_ConflictsWithSCIM(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the OTP Access
	// endpoint does not yet support the API tokens for updates and it results in
	// state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareAccessIdentityProviderOneTimePinWithScim(rnd, cloudflare.AccountIdentifier(accountID)),
				ExpectError: regexp.MustCompile(`"scim_config" can not be set if "type" is one of: "onetimepin"`),
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_OIDC_Comprehensive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOAuthMinimal(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOIDCComprehensive(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("oidc")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("auth_url"), knownvalue.StringExact("https://example.com/auth")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("token_url"), knownvalue.StringExact("https://example.com/token")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("certs_url"), knownvalue.StringExact("https://example.com/certs")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(0), knownvalue.StringExact("openid")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(1), knownvalue.StringExact("profile")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(2), knownvalue.StringExact("email")),
				},
			},
			{
				Config:                  testAccCheckCloudflareAccessIdentityProviderOIDCComprehensive(accountID, rnd),
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOIDCComprehensive(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_Okta(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderOkta(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("okta")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("okta_account"), knownvalue.StringExact("https://terraform-test.okta.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("authorization_server_id"), knownvalue.StringExact("default")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderOkta(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_GenericOAuth(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderGenericOAuth(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("oidc")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("secret")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("auth_url"), knownvalue.StringExact("https://example.com/auth")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("token_url"), knownvalue.StringExact("https://example.com/token")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("certs_url"), knownvalue.StringExact("https://example.com/certs")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(0), knownvalue.StringExact("openid")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(1), knownvalue.StringExact("profile")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("scopes").AtSliceIndex(2), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("pkce_enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderGenericOAuth(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SAML_Comprehensive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderSAMLComprehensive(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saml")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("issuer_url"), knownvalue.StringExact("jumpcloud")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sso_target_url"), knownvalue.StringExact("https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes").AtSliceIndex(0), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes").AtSliceIndex(1), knownvalue.StringExact("username")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("attributes").AtSliceIndex(2), knownvalue.StringExact("groups")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("email_attribute_name"), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sign_request"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("idp_public_certs"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("header_attributes"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("header_attributes").AtSliceIndex(0).AtMapKey("attribute_name"), knownvalue.StringExact("department")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("header_attributes").AtSliceIndex(0).AtMapKey("header_name"), knownvalue.StringExact("X-Department")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.sign_request"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderSAMLComprehensive(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_AzureAD_Comprehensive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADComprehensive(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("azureAD")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_secret"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("directory_id"), knownvalue.StringExact("directory")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("support_groups"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("conditional_access_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("prompt"), knownvalue.StringExact("select_account")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("seat_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("user_deprovision"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("identity_update_behavior"), knownvalue.StringExact("automatic")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("scim_base_url"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.conditional_access_enabled", "scim_config.secret"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADComprehensive(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_SCIM_IdentityUpdateBehaviorValues(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	testCases := []struct {
		name     string
		behavior string
	}{
		{name: "NoAction", behavior: "no_action"},
		{name: "Reauth", behavior: "reauth"},
		{name: "Automatic", behavior: "automatic"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccCheckCloudflareAccessIdentityProviderAzureADWithIdentityUpdateBehavior(accountID, rnd, tc.behavior),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("identity_update_behavior"), knownvalue.StringExact(tc.behavior)),
						},
					},
				},
			})
		})
	}
}

func TestAccCloudflareAccessIdentityProvider_PlanModifiers_SCIMSecretPersistence(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				// Update a non-SCIM field to test that secret is preserved in state
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				// Ensures no diff on plan - verifies plan modifier keeps secret from showing as change
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_PlanModifiers_SCIMSecretAfterImport(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				// Import the resource - this will set secret to null in state
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.conditional_access_enabled", "scim_config.secret"},
			},
			{
				// After import, ensure no diff on plan with same config
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				PlanOnly: true,
			},
			{
				// Update a non-SCIM field after import to test the plan modifier fix
				// This should NOT show a diff for scim_config.secret
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact("test2")),
				},
			},
			{
				// Ensures no diff on final plan - this is the critical test
				// that verifies the plan modifier doesn't incorrectly set secret to unknown
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessIdentityProvider_Normalizers_SCIMBaseURLPersistence(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_identity_provider." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderAzureAD(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("scim_base_url"), knownvalue.NotNull()),
				},
			},
			{
				// Update a non-SCIM field to test that scim_base_url is preserved from state
				Config: testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("scim_base_url"), knownvalue.NotNull()),
				},
			},
			{
				// Ensures no diff on plan - verifies normalizer keeps scim_base_url from showing as change
				Config:   testAccCheckCloudflareAccessIdentityProviderAzureADUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderOneTimePin(name string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessidentityprovideronetimepin.tf", name, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessIdentityProviderOneTimePinWithScim(name string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessidentityprovideronetimepinwithscim.tf", name, identifier.Type, identifier.Identifier)
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

func testAccCheckCloudflareAccessIdentityProviderOAuthMinimal(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovideroauthminimal.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderOIDCComprehensive(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovideroidccomprehensive.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderOkta(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityproviderokta.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderGenericOAuth(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovidergenericoauth.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderSAMLComprehensive(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityprovidersamlcomprehensive.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderAzureADComprehensive(accountID, name string) string {
	return acctest.LoadTestCase("accessidentityproviderazurecomprehensive.tf", accountID, name)
}

func testAccCheckCloudflareAccessIdentityProviderAzureADWithIdentityUpdateBehavior(accountID, name, behavior string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name       = "%[2]s"
  type       = "azureAD"
  config = {
    client_id      = "test"
    client_secret  = "test"
    directory_id   = "directory"
    support_groups = true
  }
  scim_config = {
    enabled          = true
    seat_deprovision = true
    user_deprovision = true
    identity_update_behavior = "%[3]s"
  }
}`, accountID, name, behavior)
}

func testAccCheckCloudflareZeroTrustAccessIdentityProviderDestroy(s *terraform.State) error {
	client, _ := acctest.SharedV1Client()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_identity_provider" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]

		var err error
		if accountID != "" {
			_, err = client.GetAccessIdentityProvider(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)
		} else {
			_, err = client.GetAccessIdentityProvider(context.Background(), cloudflare.ZoneIdentifier(zoneID), rs.Primary.ID)
		}

		if err == nil {
			return fmt.Errorf("zero trust access identity provider still exists")
		}
	}

	return nil
}
