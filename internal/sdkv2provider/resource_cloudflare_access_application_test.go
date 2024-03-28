package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_application", &resource.Sweeper{
		Name: "cloudflare_access_application",
		F:    testSweepCloudflareAccessApplications,
	})
}

func testSweepCloudflareAccessApplications(r string) error {
	ctx := context.Background()

	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	// Zone level Access Applications.
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneAccessApps, _, err := client.ListAccessApplications(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListAccessApplicationsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zone level Access Applications: %s", err))
	}

	if len(zoneAccessApps) == 0 {
		log.Print("[DEBUG] No Cloudflare zone level Access Applications to sweep")
		return nil
	}

	for _, accessApp := range zoneAccessApps {
		if err := client.DeleteAccessApplication(context.Background(), cloudflare.ZoneIdentifier(zoneID), accessApp.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zone level Access Application %s", accessApp.ID))
		}
	}

	// Account level Access Applications.
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountAccessApps, _, err := client.ListAccessApplications(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessApplicationsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch account level Access Applications: %s", err))
	}

	if len(accountAccessApps) == 0 {
		log.Print("[DEBUG] No Cloudflare account level Access Applications to sweep")
		return nil
	}

	for _, accessApp := range accountAccessApps {
		if err := client.DeleteAccessApplication(context.Background(), cloudflare.AccountIdentifier(accountID), accessApp.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete account level Access Application %s", accessApp.ID))
		}
	}

	return nil
}

var (
	zoneID = os.Getenv("CLOUDFLARE_ZONE_ID")
	domain = os.Getenv("CLOUDFLARE_DOMAIN")
)

func TestAccCloudflareAccessApplication_BasicZone(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "saas_app.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_BasicAccount(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "sass_app.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithCORS(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "1"),
					resource.TestCheckResourceAttr(name, "cors_headers.0.allowed_methods.#", "3"),
					resource.TestCheckResourceAttr(name, "cors_headers.0.allowed_origins.#", "1"),
					resource.TestCheckResourceAttr(name, "cors_headers.0.max_age", "10"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSAMLSaas(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "saas"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "saas_app.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.0.sp_entity_id", "saas-app.example"),
					resource.TestCheckResourceAttr(name, "saas_app.0.consumer_service_url", "https://saas-app.example/sso/saml/consume"),
					resource.TestCheckResourceAttr(name, "saas_app.0.name_id_format", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.0.default_relay_state", "https://saas-app.example"),
					resource.TestCheckResourceAttr(name, "saas_app.0.name_id_transform_jsonata", "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
					resource.TestCheckResourceAttr(name, "saas_app.0.saml_attribute_transform_jsonata", "$ ~>| groups | {'group_name': name} |"),

					resource.TestCheckResourceAttrSet(name, "saas_app.0.idp_entity_id"),
					resource.TestCheckResourceAttrSet(name, "saas_app.0.public_key"),
					resource.TestCheckResourceAttrSet(name, "saas_app.0.sso_endpoint"),

					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.#", "2"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.name", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.name_format", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.source.0.name", "user_email"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.source.0.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.friendly_name", "Rank"),
					resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.required", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSAMLSaas_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := "cloudflare_access_application." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "type", "saas"),
		resource.TestCheckResourceAttr(name, "session_duration", "24h"),
		resource.TestCheckResourceAttr(name, "saas_app.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.0.sp_entity_id", "saas-app.example"),
		resource.TestCheckResourceAttr(name, "saas_app.0.consumer_service_url", "https://saas-app.example/sso/saml/consume"),
		resource.TestCheckResourceAttr(name, "saas_app.0.name_id_format", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.0.default_relay_state", "https://saas-app.example"),
		resource.TestCheckResourceAttr(name, "saas_app.0.name_id_transform_jsonata", "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
		resource.TestCheckResourceAttr(name, "saas_app.0.saml_attribute_transform_jsonata", "$ ~>| groups | {'group_name': name} |"),

		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.#", "2"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.name", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.name_format", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.0.source.0.name", "user_email"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.source.0.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.friendly_name", "Rank"),
		resource.TestCheckResourceAttr(name, "saas_app.0.custom_attribute.1.required", "true"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				Check:  checkFn,
			},
			{
				ImportState:         true,
				ImportStateVerify:   true,
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				Check:               checkFn,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithOIDCSaas(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "saas"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "saas_app.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.0.auth_type", "oidc"),
					resource.TestCheckResourceAttr(name, "saas_app.0.redirect_uris.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.0.redirect_uris.0", "https://saas-app.example/sso/oauth2/callback"),
					resource.TestCheckResourceAttr(name, "saas_app.0.grant_types.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.0.grant_types.0", "authorization_code"),
					resource.TestCheckResourceAttr(name, "saas_app.0.scopes.#", "4"),
					resource.TestCheckResourceAttr(name, "saas_app.0.scopes.0", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.0.scopes.1", "groups"),
					resource.TestCheckResourceAttr(name, "saas_app.0.scopes.2", "openid"),
					resource.TestCheckResourceAttr(name, "saas_app.0.scopes.3", "profile"),
					resource.TestCheckResourceAttr(name, "saas_app.0.app_launcher_url", "https://saas-app.example/sso/login"),
					resource.TestCheckResourceAttr(name, "saas_app.0.group_filter_regex", ".*"),
					resource.TestCheckResourceAttrSet(name, "saas_app.0.client_secret"),
					resource.TestCheckResourceAttrSet(name, "saas_app.0.public_key"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithOIDCSaas_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := "cloudflare_access_application." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "type", "saas"),
		resource.TestCheckResourceAttr(name, "session_duration", "24h"),
		resource.TestCheckResourceAttr(name, "saas_app.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.0.auth_type", "oidc"),
		resource.TestCheckResourceAttr(name, "saas_app.0.redirect_uris.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.0.redirect_uris.0", "https://saas-app.example/sso/oauth2/callback"),
		resource.TestCheckResourceAttr(name, "saas_app.0.grant_types.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.0.grant_types.0", "authorization_code"),
		resource.TestCheckResourceAttr(name, "saas_app.0.scopes.#", "4"),
		resource.TestCheckResourceAttr(name, "saas_app.0.scopes.0", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.0.scopes.1", "groups"),
		resource.TestCheckResourceAttr(name, "saas_app.0.scopes.2", "openid"),
		resource.TestCheckResourceAttr(name, "saas_app.0.scopes.3", "profile"),
		resource.TestCheckResourceAttr(name, "saas_app.0.app_launcher_url", "https://saas-app.example/sso/login"),
		resource.TestCheckResourceAttr(name, "saas_app.0.group_filter_regex", ".*"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				Check:  checkFn,
			},
			{
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"saas_app.0.client_secret"},
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				Check:                   checkFn,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAutoRedirectToIdentity(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithAutoRedirectToIdentity(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "true"),
					resource.TestCheckResourceAttr(name, "allowed_idps.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithEnableBindingCookie(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithEnableBindingCookie(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "enable_binding_cookie", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithCustomDenyFields(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithCustomDenyFields(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "custom_deny_message", "denied!"),
					resource.TestCheckResourceAttr(name, "custom_deny_url", "https://www.cloudflare.com"),
					resource.TestCheckResourceAttr(name, "custom_non_identity_deny_url", "https://www.blocked.com"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithADefinedIdps(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithADefinedIdp(rnd, zoneID, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "true"),
					resource.TestCheckResourceAttr(name, "allowed_idps.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithMultipleIdpsReordered(t *testing.T) {
	rnd := generateRandomResourceName()
	idp1 := generateRandomResourceName()
	idp2 := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp1, idp2),
			},
			{
				Config: testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp2, idp1),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithHttpOnlyCookieAttribute(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttribute(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "http_only_cookie_attribute", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithHTTPOnlyCookieAttributeSetToFalse(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttributeSetToFalse(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "http_only_cookie_attribute", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSameSiteCookieAttribute(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigSameSiteCookieAttribute(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "same_site_cookie_attribute", "strict"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithLogoURL(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigLogoURL(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "logo_url", "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSkipInterstitial(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigSkipInterstitial(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "skip_interstitial", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAppLauncherVisible(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithAppLauncherVisible(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "app_launcher_visible", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSelfHostedDomains(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationWithSelfHostedDomains(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttrSet(name, "domain"),
					resource.TestCheckResourceAttr(name, "self_hosted_domains.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "sass_app.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithDefinedTags(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithADefinedTag(rnd, zoneID, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAppLauncherCustomization(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{Config: testAccessApplicationWithAppLauncherCustomizationFields(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", "App Launcher"),
					resource.TestCheckResourceAttr(name, "type", "app_launcher"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "header_bg_color", "#000000"),
					resource.TestCheckResourceAttr(name, "bg_color", "#000000"),
					resource.TestCheckResourceAttr(name, "app_launcher_logo_url", "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					resource.TestCheckResourceAttr(name, "landing_page_design.#", "1"),
					resource.TestCheckResourceAttr(name, "landing_page_design.0.title", "title"),
					resource.TestCheckResourceAttr(name, "landing_page_design.0.message", "message"),
					resource.TestCheckResourceAttr(name, "landing_page_design.0.button_color", "#000000"),
					resource.TestCheckResourceAttr(name, "landing_page_design.0.button_text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "landing_page_design.0.image_url", "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					resource.TestCheckResourceAttr(name, "footer_links.#", "1"),
					resource.TestCheckResourceAttr(name, "footer_links.0.name", "footer link"),
					resource.TestCheckResourceAttr(name, "footer_links.0.url", "https://www.cloudflare.com"),
				),
			},
		},
	})
}

func testAccCloudflareAccessApplicationConfigBasic(rnd string, domain string, identifier *cloudflare.ResourceContainer) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  %[3]s_id                  = "%[4]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[2]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false
}
`, rnd, domain, identifier.Type, identifier.Identifier)
}

func testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id          = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"
  cors_headers {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com"]
    allow_credentials = true
    max_age = 10
  }
  auto_redirect_to_identity = false
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "saas"
  session_duration = "24h"
  saas_app {
    consumer_service_url = "https://saas-app.example/sso/saml/consume"
    sp_entity_id  = "saas-app.example"
    name_id_format =  "email"
	default_relay_state = "https://saas-app.example"
	name_id_transform_jsonata = "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"
	saml_attribute_transform_jsonata = "$ ~>| groups | {'group_name': name} |"

	custom_attribute {
		name = "email"
		name_format = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
		source {
			name = "user_email"
		}
	}
	custom_attribute {
		name = "rank"
		required = true
		friendly_name = "Rank"
		source {
			name = "rank"
		}
	}
  }
  auto_redirect_to_identity = false
}
`, rnd, accountID)
}

func testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "saas"
  session_duration = "24h"
  saas_app {
	auth_type = "oidc"
	redirect_uris = ["https://saas-app.example/sso/oauth2/callback"]
	grant_types = ["authorization_code"]
	scopes = ["openid", "email", "profile", "groups"]
	app_launcher_url = "https://saas-app.example/sso/login"
	group_filter_regex = ".*"
  }
  auto_redirect_to_identity = false
}
`, rnd, accountID)
}

func testAccCloudflareAccessApplicationConfigWithAutoRedirectToIdentity(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  type    = "onetimepin"
}

resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = true
  allowed_idps              = [cloudflare_access_identity_provider.%[1]s.id]

  depends_on = ["cloudflare_access_identity_provider.%[1]s"]
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithEnableBindingCookie(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  enable_binding_cookie     = true
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithCustomDenyFields(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  custom_deny_message       = "denied!"
  custom_deny_url           = "https://www.cloudflare.com"
	custom_non_identity_deny_url = "https://www.blocked.com"
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithADefinedIdp(rnd, zoneID, domain string, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "%[4]s"
  name = "%[1]s"
  type = "onetimepin"
}
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = true
  allowed_idps              = [cloudflare_access_identity_provider.%[1]s.id]
}
`, rnd, zoneID, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp1, idp2 string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[5]s" {
  account_id = "%[4]s"
  name = "%[5]s"
  type = "onetimepin"
}

resource "cloudflare_access_identity_provider" "%[6]s" {
  account_id = "%[4]s"
  name = "%[6]s"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
  }
}

resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  allowed_idps              = [
    cloudflare_access_identity_provider.%[5]s.id,
    cloudflare_access_identity_provider.%[6]s.id,
  ]
}
`, rnd, zoneID, domain, accountID, idp1, idp2)
}

func testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttribute(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  http_only_cookie_attribute = true
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttributeSetToFalse(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  http_only_cookie_attribute = false
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigSameSiteCookieAttribute(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  same_site_cookie_attribute = "strict"
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigSkipInterstitial(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  skip_interstitial          = true
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithAppLauncherVisible(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  app_launcher_visible      = true
}
`, rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigLogoURL(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  logo_url          		 = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
}
`, rnd, zoneID, domain)
}

func testAccessApplicationWithAppLauncherCustomizationFields(rnd, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_application" "%[1]s" {
			account_id       = "%[2]s"
			type             = "app_launcher"
			session_duration = "24h"
			app_launcher_visible = false
			app_launcher_logo_url = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
			bg_color = "#000000"
			header_bg_color = "#000000"

			footer_links {
				name = "footer link"
				url = "https://www.cloudflare.com"
			}


			landing_page_design {
				title = "title"
				message = "message"
				button_color = "#000000"
				image_url = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
				button_text_color = "#000000"
			}
	}
	`, rnd, accountID)
}

func testAccCloudflareAccessApplicationWithSelfHostedDomains(rnd string, domain string, identifier *cloudflare.ResourceContainer) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
  %[3]s_id                  = "%[4]s"
  name                      = "%[1]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false
  self_hosted_domains       = [
    "d1.%[1]s.%[2]s",
    "d2.%[1]s.%[2]s"
  ]
}
`, rnd, domain, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessApplicationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_application" {
			continue
		}

		var notFoundError *cloudflare.NotFoundError
		if rs.Primary.Attributes[consts.ZoneIDSchemaKey] != "" {
			_, err := client.GetAccessApplication(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
			if !errors.As(err, &notFoundError) {
				return fmt.Errorf("AccessApplication still exists")
			}
		}

		if rs.Primary.Attributes[consts.AccountIDSchemaKey] != "" {
			_, err := client.GetAccessApplication(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
			if !errors.As(err, &notFoundError) {
				return fmt.Errorf("AccessApplication still exists")
			}
		}

	}

	return nil
}

func TestAccCloudflareAccessApplicationWithZoneID(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_application." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	updatedName := fmt.Sprintf("%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessApplicationWithZoneID(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
			{
				Config: testAccessApplicationWithZoneIDUpdated(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", updatedName),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithMissingCORSMethods(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithMissingCORSMethods(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile("must set allowed_methods or allow_all_methods"),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithMissingCORSOrigins(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithMissingCORSOrigins(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile("must set allowed_origins or allow_all_origins"),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithInvalidSessionDuration(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithInvalidSessionDuration(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta(`"session_duration" only supports "ns", "us" (or "µs"), "ms", "s", "m", or "h" as valid units`)),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationMisconfiguredCORSCredentialsAllowingAllOrigins(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationMisconfiguredCORSAllowAllOriginsWithCredentials(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta(`CORS credentials are not permitted when all origins are allowed`)),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationMisconfiguredCORSCredentialsAllowingWildcardOrigins(t *testing.T) {
	rnd := generateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationMisconfiguredCORSAllowWildcardOriginWithCredentials(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta(`CORS credentials are not permitted when all origins are allowed`)),
			},
		},
	})
}

func testAccessApplicationWithZoneID(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"
    }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationWithZoneIDUpdated(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s-updated"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"
    }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationWithMissingCORSMethods(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s-updated"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"

    cors_headers {
      allow_all_origins = true
    }
  }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationWithMissingCORSOrigins(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s-updated"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"

    cors_headers {
      allow_all_methods = true
    }
  }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationWithInvalidSessionDuration(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name             = "%[1]s-updated"
      zone_id          = "%[3]s"
      domain           = "%[1]s.%[2]s"
      type             = "self_hosted"
      session_duration = "24z"
  }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationMisconfiguredCORSAllowAllOriginsWithCredentials(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name             = "%[1]s-updated"
      zone_id          = "%[3]s"
      domain           = "%[1]s.%[2]s"
      type             = "self_hosted"

      cors_headers {
        allowed_methods = ["GET"]
        allow_all_origins = true
        allow_credentials = true
      }
  }
  `, resourceID, zone, zoneID)
}

func testAccessApplicationMisconfiguredCORSAllowWildcardOriginWithCredentials(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name             = "%[1]s-updated"
      zone_id          = "%[3]s"
      domain           = "%[1]s.%[2]s"
      type             = "self_hosted"

      cors_headers {
        allowed_methods = ["GET"]
        allowed_origins = ["*"]
        allow_credentials = true
      }
  }
  `, resourceID, zone, zoneID)
}

func testAccCloudflareAccessApplicationConfigWithADefinedTag(rnd, zoneID, domain string, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_tag" "%[1]s" {
  account_id = "%[4]s"
  name = "%[1]s"
}
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
 	tags             = [cloudflare_access_tag.%[1]s.id]
}
`, rnd, zoneID, domain, accountID)
}
