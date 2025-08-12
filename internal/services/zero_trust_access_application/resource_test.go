package zero_trust_access_application_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_application", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_application",
		F:    testSweepCloudflareAccessApplications,
	})
}

func testSweepCloudflareAccessApplications(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
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
	zoneID    = os.Getenv("CLOUDFLARE_ZONE_ID")
	domain    = os.Getenv("CLOUDFLARE_DOMAIN")
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestAccCloudflareAccessApplication_BasicZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
					resource.TestCheckResourceAttr(name, "saas_app.%", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCloudflareAccessApplicationConfigBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_BasicAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
				),
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCloudflareAccessApplicationConfigBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigHttpBasic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	idpName := fmt.Sprintf("cloudflare_zero_trust_access_identity_provider.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "scim_config.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.remote_uri", "scim.com"),
					resource.TestCheckResourceAttrPair(name, "scim_config.idp_uid", idpName, "id"),
					resource.TestCheckResourceAttr(name, "scim_config.deactivate_on_delete", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scheme", "httpbasic"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.user", "test"),
					resource.TestCheckResourceAttrSet(name, "scim_config.authentication.password"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.schema", "urn:ietf:params:scim:schemas:core:2.0:User"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.filter", "title pr or userType eq \"Intern\""),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.transform_jsonata", "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.create", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.update", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.delete", "true"),
				),
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_UpdateSCIMConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	idpName := fmt.Sprintf("cloudflare_zero_trust_access_identity_provider.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "scim_config.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.remote_uri", "scim.com"),
					resource.TestCheckResourceAttrPair(name, "scim_config.idp_uid", idpName, "id"),
					resource.TestCheckResourceAttr(name, "scim_config.deactivate_on_delete", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scheme", "httpbasic"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.user", "test"),
					resource.TestCheckResourceAttrSet(name, "scim_config.authentication.password"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.schema", "urn:ietf:params:scim:schemas:core:2.0:User"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.filter", "title pr or userType eq \"Intern\""),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.transform_jsonata", "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.create", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.update", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.delete", "true"),
				),
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				PlanOnly: true,
			},
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerTokenNoMappings(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "scim_config.enabled", "false"),
					resource.TestCheckResourceAttr(name, "scim_config.remote_uri", "scim2.com"),
					resource.TestCheckResourceAttrPair(name, "scim_config.idp_uid", idpName, "id"),
					resource.TestCheckResourceAttr(name, "scim_config.deactivate_on_delete", "false"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scheme", "oauthbearertoken"),
					resource.TestCheckResourceAttrSet(name, "scim_config.authentication.token"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.#", "0"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerTokenNoMappings(rnd, accountID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigInvalidMappingSchema(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareAccessApplicationSCIMConfigInvalidMappingSchema(rnd, accountID, domain),
				ExpectError: regexp.MustCompile(`.*invalid SCIM schema in mappings.*`),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigHttpBasicMissingRequired(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareAccessApplicationSCIMConfigHttpBasicMissingRequired(rnd, accountID, domain),
				ExpectError: regexp.MustCompile(`.*password is a required field.*`),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigOAuthBearerToken(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	idpName := fmt.Sprintf("cloudflare_zero_trust_access_identity_provider.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerToken(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "scim_config.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.remote_uri", "scim.com"),
					resource.TestCheckResourceAttrPair(name, "scim_config.idp_uid", idpName, "id"),
					resource.TestCheckResourceAttr(name, "scim_config.deactivate_on_delete", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scheme", "oauthbearertoken"),
					resource.TestCheckResourceAttrSet(name, "scim_config.authentication.token"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.schema", "urn:ietf:params:scim:schemas:core:2.0:User"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.filter", "title pr or userType eq \"Intern\""),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.transform_jsonata", "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.create", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.update", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.delete", "true"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerToken(rnd, accountID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigOAuth2(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	idpName := fmt.Sprintf("cloudflare_zero_trust_access_identity_provider.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuth2(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "scim_config.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.remote_uri", "scim.com"),
					resource.TestCheckResourceAttrPair(name, "scim_config.idp_uid", idpName, "id"),
					resource.TestCheckResourceAttr(name, "scim_config.deactivate_on_delete", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scheme", "oauth2"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.client_id", "beepboop"),
					resource.TestCheckResourceAttrSet(name, "scim_config.authentication.client_secret"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.authorization_url", "https://www.authorization.com"),
					resource.TestCheckTypeSetElemAttr(name, "scim_config.authentication.scopes.*", "read"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.scopes.#", "1"),
					resource.TestCheckResourceAttr(name, "scim_config.authentication.token_url", "https://www.token.com"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.schema", "urn:ietf:params:scim:schemas:core:2.0:User"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.filter", "title pr or userType eq \"Intern\""),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.transform_jsonata", "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.create", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.update", "true"),
					resource.TestCheckResourceAttr(name, "scim_config.mappings.0.operations.delete", "true"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidOAuth2(rnd, accountID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSCIMConfigOAuth2MissingRequired(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareAccessApplicationSCIMConfigOAuth2MissingRequired(rnd, accountID, domain),
				ExpectError: regexp.MustCompile(`.*token_url is a required field.*`),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithCORS(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.allowed_methods.#", "3"),
					resource.TestCheckResourceAttr(name, "cors_headers.allowed_origins.#", "1"),
					resource.TestCheckResourceAttr(name, "cors_headers.max_age", "10"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSAMLSaas(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "saas"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "saas_app.sp_entity_id", "saas-app.example"),
					resource.TestCheckResourceAttr(name, "saas_app.consumer_service_url", "https://saas-app.example/sso/saml/consume"),
					resource.TestCheckResourceAttr(name, "saas_app.name_id_format", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.default_relay_state", "https://saas-app.example"),
					resource.TestCheckResourceAttr(name, "saas_app.name_id_transform_jsonata", "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
					resource.TestCheckResourceAttr(name, "saas_app.saml_attribute_transform_jsonata", "$ ~>| groups | {'group_name': name} |"),

					resource.TestCheckResourceAttrSet(name, "saas_app.idp_entity_id"),
					resource.TestCheckResourceAttrSet(name, "saas_app.public_key"),
					resource.TestCheckResourceAttrSet(name, "saas_app.sso_endpoint"),

					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.#", "2"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.name", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.name_format", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.source.name", "user_email"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.source.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.friendly_name", "Rank"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.required", "true"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSAMLSaas_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_application." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "type", "saas"),
		resource.TestCheckResourceAttr(name, "session_duration", "24h"),
		resource.TestCheckResourceAttr(name, "saas_app.sp_entity_id", "saas-app.example"),
		resource.TestCheckResourceAttr(name, "saas_app.consumer_service_url", "https://saas-app.example/sso/saml/consume"),
		resource.TestCheckResourceAttr(name, "saas_app.name_id_format", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.default_relay_state", "https://saas-app.example"),
		resource.TestCheckResourceAttr(name, "saas_app.name_id_transform_jsonata", "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
		resource.TestCheckResourceAttr(name, "saas_app.saml_attribute_transform_jsonata", "$ ~>| groups | {'group_name': name} |"),

		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.#", "2"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.name", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.name_format", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.0.source.name", "user_email"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.source.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.friendly_name", "Rank"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_attributes.1.required", "true"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				Check:  checkFn,
			},
			{
				ImportState:         true,
				ImportStateVerify:   true,
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("accounts/%s/", accountID),
				Check:               checkFn,
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithOIDCSaas(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "saas"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "saas_app.auth_type", "oidc"),
					resource.TestCheckResourceAttr(name, "saas_app.redirect_uris.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.redirect_uris.0", "https://saas-app.example/sso/oauth2/callback"),
					resource.TestCheckResourceAttr(name, "saas_app.grant_types.#", "2"),
					resource.TestCheckResourceAttr(name, "saas_app.grant_types.0", "authorization_code"),
					resource.TestCheckResourceAttr(name, "saas_app.grant_types.1", "hybrid"),
					resource.TestCheckResourceAttr(name, "saas_app.scopes.0", "openid"),
					resource.TestCheckResourceAttr(name, "saas_app.scopes.1", "email"),
					resource.TestCheckResourceAttr(name, "saas_app.scopes.2", "profile"),
					resource.TestCheckResourceAttr(name, "saas_app.scopes.3", "groups"),
					resource.TestCheckResourceAttr(name, "saas_app.app_launcher_url", "https://saas-app.example/sso/login"),
					resource.TestCheckResourceAttr(name, "saas_app.group_filter_regex", ".*"),
					resource.TestCheckResourceAttr(name, "saas_app.allow_pkce_without_client_secret", "false"),
					resource.TestCheckResourceAttr(name, "saas_app.refresh_token_options.lifetime", "1h"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_claims.#", "1"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.scope", "profile"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.required", "true"),
					resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.source.name", "rank"),
					resource.TestCheckResourceAttr(name, "saas_app.hybrid_and_implicit_options.return_access_token_from_authorization_endpoint", "true"),
					resource.TestCheckResourceAttr(name, "saas_app.hybrid_and_implicit_options.return_id_token_from_authorization_endpoint", "true"),
					resource.TestCheckResourceAttrSet(name, "saas_app.client_secret"),
					resource.TestCheckResourceAttrSet(name, "saas_app.public_key"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithOIDCSaas_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_application." + rnd

	checkFn := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "type", "saas"),
		resource.TestCheckResourceAttr(name, "session_duration", "24h"),
		resource.TestCheckResourceAttr(name, "saas_app.auth_type", "oidc"),
		resource.TestCheckResourceAttr(name, "saas_app.redirect_uris.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.redirect_uris.0", "https://saas-app.example/sso/oauth2/callback"),
		resource.TestCheckResourceAttr(name, "saas_app.grant_types.#", "2"),
		resource.TestCheckResourceAttr(name, "saas_app.grant_types.0", "authorization_code"),
		resource.TestCheckResourceAttr(name, "saas_app.grant_types.1", "hybrid"),
		resource.TestCheckResourceAttr(name, "saas_app.scopes.0", "openid"),
		resource.TestCheckResourceAttr(name, "saas_app.scopes.1", "email"),
		resource.TestCheckResourceAttr(name, "saas_app.scopes.2", "profile"),
		resource.TestCheckResourceAttr(name, "saas_app.scopes.3", "groups"),
		resource.TestCheckResourceAttr(name, "saas_app.app_launcher_url", "https://saas-app.example/sso/login"),
		resource.TestCheckResourceAttr(name, "saas_app.group_filter_regex", ".*"),
		resource.TestCheckResourceAttr(name, "saas_app.allow_pkce_without_client_secret", "false"),
		resource.TestCheckResourceAttr(name, "saas_app.refresh_token_options.lifetime", "1h"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_claims.#", "1"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.scope", "profile"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.required", "true"),
		resource.TestCheckResourceAttr(name, "saas_app.custom_claims.0.source.name", "rank"),
		resource.TestCheckResourceAttr(name, "saas_app.hybrid_and_implicit_options.return_access_token_from_authorization_endpoint", "true"),
		resource.TestCheckResourceAttr(name, "saas_app.hybrid_and_implicit_options.return_id_token_from_authorization_endpoint", "true"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				Check:  checkFn,
			},
			{
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"saas_app.client_secret", "saas_app.allow_pkce_without_client_secret"},
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				Check:                   checkFn,
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAutoRedirectToIdentity(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithAutoRedirectToIdentity(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithEnableBindingCookie(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithEnableBindingCookie(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithCustomDenyFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithCustomDenyFields(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithADefinedIdp(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithADefinedIdp(rnd, zoneID, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithMultipleIdpsReordered(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	idp1 := utils.GenerateRandomResourceName()
	idp2 := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp1, idp2),
			},
			{
				Config: testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp2, idp1),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp2, idp1),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithHttpOnlyCookieAttribute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttribute(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithHTTPOnlyCookieAttributeSetToFalse(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttributeSetToFalse(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSameSiteCookieAttribute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigSameSiteCookieAttribute(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithLogoURL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigLogoURL(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSkipInterstitial(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigSkipInterstitial(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAppLauncherVisible(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithAppLauncherVisible(rnd, zoneID, domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithSelfHostedDomains(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationWithSelfHostedDomains(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "self_hosted_domains.#", "2"),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "cors_headers.#", "0"),
					resource.TestCheckResourceAttr(name, "sass_app.#", "0"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationWithSelfHostedDomains(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithDefinedTags(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
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
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithADefinedTag(rnd, zoneID, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithLegacyPolicies(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithLegacyPolicies(rnd, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "policies.#", "3"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithLegacyPolicies(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithReusablePolicies(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithReusablePolicies(rnd, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "type", "self_hosted"),
					resource.TestCheckResourceAttr(name, "policies.#", "2"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationConfigWithReusablePolicies(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithReusablePolicies_InvalidPrecedence(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareAccessApplicationConfigWithReusablePoliciesInvalidPrecedence(rnd, domain, accountID),
				ExpectError: regexp.MustCompile(`Attribute policies\[0].precedence value must be at least 1, got: 0`),
			},
		},
	})
}

func TestAccCloudflareAccessApplication_WithAppLauncherCustomization(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessApplicationWithAppLauncherCustomizationFields(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "type", "app_launcher"),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "header_bg_color", "#000000"),
					resource.TestCheckResourceAttr(name, "bg_color", "#000000"),
					resource.TestCheckResourceAttr(name, "app_launcher_logo_url", "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					resource.TestCheckResourceAttr(name, "landing_page_design.title", "title"),
					resource.TestCheckResourceAttr(name, "landing_page_design.message", "message"),
					resource.TestCheckResourceAttr(name, "landing_page_design.button_color", "#000000"),
					resource.TestCheckResourceAttr(name, "landing_page_design.button_text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "landing_page_design.image_url", "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					resource.TestCheckResourceAttr(name, "footer_links.0.name", "footer link"),
					resource.TestCheckResourceAttr(name, "footer_links.0.url", "https://www.cloudflare.com"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessApplicationWithAppLauncherCustomizationFields(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_Infrastructure(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationInfrastructure(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckNoResourceAttr(name, "session_duration"),
					resource.TestCheckResourceAttr(name, "type", "infrastructure"),
					resource.TestCheckResourceAttr(name, "target_criteria.#", "1"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.port", "22"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.protocol", "SSH"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.target_attributes.hostname.#", "1"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.target_attributes.hostname.0", rnd),
					resource.TestCheckResourceAttr(name, "policies.#", "1"),
					resource.TestCheckResourceAttr(name, "policies.0.connection_rules.ssh.usernames.#", "1"),
					resource.TestCheckResourceAttr(name, "policies.0.connection_rules.ssh.usernames.0", "root"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationInfrastructure(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_RDP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	appDomain := fmt.Sprintf("%[1]s.%[2]s", rnd, domain)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationRDP(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "type", "rdp"),
					resource.TestCheckResourceAttr(name, "domain", appDomain),
					resource.TestCheckResourceAttr(name, "destinations.#", "1"),
					resource.TestCheckResourceAttr(name, "destinations.0.uri", appDomain),
					resource.TestCheckResourceAttr(name, "target_criteria.#", "1"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.port", "3389"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.protocol", "RDP"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.target_attributes.hostname.#", "1"),
					resource.TestCheckResourceAttr(name, "target_criteria.0.target_attributes.hostname.0", rnd),
					resource.TestCheckResourceAttr(name, "policies.#", "1"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessApplicationRDP(rnd, accountID, domain),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareAccessApplicationConfigBasic(rnd string, domain string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessapplicationconfigbasic.tf", rnd, domain, identifier.Type, identifier.Identifier)
}

func testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithcors.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationInfrastructure(rnd, accID string) string {
	return acctest.LoadTestCase("accessapplicationconfiginfrastructure.tf", rnd, accID)
}

func testAccCloudflareAccessApplicationRDP(rnd, accID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigrdp.tf", rnd, accID, domain)
}

func testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithsamlsaas.tf", rnd, accountID)
}

func testAccCloudflareAccessApplicationConfigWithOIDCSaas(rnd, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithoidcsaas.tf", rnd, accountID)
}

func testAccCloudflareAccessApplicationConfigWithAutoRedirectToIdentity(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithautoredirecttoidentity.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithEnableBindingCookie(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithenablebindingcookie.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithCustomDenyFields(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithcustomdenyfields.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithADefinedIdp(rnd, zoneID, domain string, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithadefinedidp.tf", rnd, zoneID, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigWithMultipleIdps(rnd, zoneID, domain, accountID, idp1, idp2 string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithmultipleidps.tf", rnd, zoneID, domain, accountID, idp1, idp2)
}

func testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttribute(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithhttponlycookieattribute.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithHTTPOnlyCookieAttributeSetToFalse(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithhttponlycookieattributesettofalse.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigSameSiteCookieAttribute(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigsamesitecookieattribute.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigSkipInterstitial(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigskipinterstitial.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigWithAppLauncherVisible(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithapplaunchervisible.tf", rnd, zoneID, domain)
}

func testAccCloudflareAccessApplicationConfigLogoURL(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("accessapplicationconfiglogourl.tf", rnd, zoneID, domain)
}

func testAccessApplicationWithAppLauncherCustomizationFields(rnd, accountID string) string {
	return acctest.LoadTestCase("accessapplicationwithapplaunchercustomizationfields.tf", rnd, accountID)
}

func testAccCloudflareAccessApplicationWithSelfHostedDomains(rnd string, domain string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessapplicationwithselfhosteddomains.tf", rnd, domain, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessApplicationDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client for destroy check: %s", clientErr))
		return clientErr
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_application" {
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_application." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	updatedName := fmt.Sprintf("%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
			{
				// Ensures no diff on last plan
				Config:   testAccessApplicationWithZoneIDUpdated(rnd, zone, zoneID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithMissingCORSMethods(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithMissingCORSMethods(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(`No attribute specified when one \(and only one\) of\s+\[cors_headers\.allow_all_methods\.<\.allowed_methods\] is required`),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithMissingCORSOrigins(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithMissingCORSOrigins(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(`No attribute specified when one \(and only one\) of\s+\[cors_headers\.allow_all_origins\.<\.allowed_origins\] is required`),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithInvalidSessionDuration(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithInvalidSessionDuration(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(`"session_duration" only supports .*`),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithInvalidPrivateDestination(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithInvalidPrivateDestination(rnd, accountID),
				ExpectError: regexp.MustCompile(`"destinations\[0]\.(hostname|port_range)" can only be set if "<\.type" is one of: "private"`),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithDestinations(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	publicDomain := fmt.Sprintf("d1.%[1]s.%[2]s", rnd, domain)
	privateDomain := fmt.Sprintf("%[1]s.private", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessApplicationWithDestinations(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "destinations.#", "2"),
					resource.TestCheckResourceAttr(name, "destinations.0.type", "private"),
					resource.TestCheckResourceAttr(name, "destinations.0.hostname", privateDomain),
					resource.TestCheckResourceAttr(name, "destinations.1.type", "public"),
					resource.TestCheckResourceAttr(name, "destinations.1.uri", publicDomain),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessApplicationWithDestinations(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplicationMisconfiguredCORSCredentialsAllowingAllOrigins(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationMisconfiguredCORSAllowAllOriginsWithCredentials(rnd, zone, zoneID),
				ExpectError: regexp.MustCompile(`Attribute "cors_headers.allow_all_origins" cannot be specified when\s+"cors_headers.allow_credentials" is specified`),
			},
		},
	})
}

func TestAccCloudflareAccessApplicationWithInvalidSaas(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accoundID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccessApplicationWithInvalidSaas(rnd, accoundID),
				ExpectError: regexp.MustCompile("\"saas_app\" has to be set if \"type\" is one of: \"saas\", \"dash_sso\""),
			},
		},
	})
}

func testAccessApplicationWithZoneID(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationwithzoneid.tf", resourceID, zone, zoneID)
}

func testAccessApplicationWithZoneIDUpdated(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationwithzoneidupdated.tf", resourceID, zone, zoneID)
}

func testAccessApplicationWithMissingCORSMethods(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationwithmissingcorsmethods.tf", resourceID, zone, zoneID)
}

func testAccessApplicationWithMissingCORSOrigins(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationwithmissingcorsorigins.tf", resourceID, zone, zoneID)
}

func testAccessApplicationWithInvalidSessionDuration(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationwithinvalidsessionduration.tf", resourceID, zone, zoneID)
}

func testAccessApplicationWithInvalidPrivateDestination(resourceID, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfiginvalidprivatedestination.tf", resourceID, accountID)
}

func testAccessApplicationWithDestinations(rnd string, domain string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessapplicationconfigwithdestinations.tf", rnd, domain, identifier.Type, identifier.Identifier)
}

func testAccessApplicationMisconfiguredCORSAllowAllOriginsWithCredentials(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accessapplicationmisconfiguredcorsallowalloriginswithcredentials.tf", resourceID, zone, zoneID)
}

func testAccCloudflareAccessApplicationConfigWithADefinedTag(rnd, zoneID, domain string, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithadefinedtag.tf", rnd, zoneID, domain, accountID)
}

func testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfigvalidhttpbasic.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerTokenNoMappings(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfigvalidoauthbearertokennomappings.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerToken(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfigvalidoauthbearertoken.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigValidOAuth2(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfigvalidoauth2.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigOAuth2MissingRequired(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfigoauth2missingrequired.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigHttpBasicMissingRequired(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfighttpbasicmissingrequired.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationSCIMConfigInvalidMappingSchema(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("accessapplicationscimconfiginvalidmappingschema.tf", rnd, accountID, domain)
}

func testAccCloudflareAccessApplicationConfigWithLegacyPolicies(rnd, domain string, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithlegacypolicies.tf", rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigWithReusablePolicies(rnd, domain string, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithreusablepolicies.tf", rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigWithReusablePoliciesInvalidPrecedence(rnd, domain string, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithreusablepolicies_invalid_precedence.tf", rnd, domain, accountID)
}

func testAccessApplicationWithInvalidSaas(resourceID, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithinvalidsaas.tf", resourceID, accountID)
}
