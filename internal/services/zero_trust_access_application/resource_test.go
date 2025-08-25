package zero_trust_access_application_test

import (
	"context"
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
	"github.com/pkg/errors"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_auth_401_redirect"), knownvalue.Bool(false)),

					// destinations and self_hosted_domains should be populated from the API
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destinations"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destinations").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("public")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destinations").AtSliceIndex(0).AtMapKey("uri"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("self_hosted_domains"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("zones/%s/", zoneID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains"},
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
	resourceName := name

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains"},
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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("remote_uri"), knownvalue.StringExact("scim.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("deactivate_on_delete"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scheme"), knownvalue.StringExact("httpbasic")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("user"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("password"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("schema"), knownvalue.StringExact("urn:ietf:params:scim:schemas:core:2.0:User")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("filter"), knownvalue.StringExact("title pr or userType eq \"Intern\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("transform_jsonata"), knownvalue.StringExact("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("create"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("update"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("delete"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "scim_config.authentication.password", "auto_redirect_to_identity"},
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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("remote_uri"), knownvalue.StringExact("scim.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("deactivate_on_delete"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scheme"), knownvalue.StringExact("httpbasic")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("user"), knownvalue.StringExact("test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("password"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("schema"), knownvalue.StringExact("urn:ietf:params:scim:schemas:core:2.0:User")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("filter"), knownvalue.StringExact("title pr or userType eq \"Intern\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("transform_jsonata"), knownvalue.StringExact("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("create"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("update"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("delete"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "scim_config.authentication.password", "auto_redirect_to_identity"},
			},
			{
				// Ensures no diff on second plan
				Config:   testAccCloudflareAccessApplicationSCIMConfigValidHttpBasic(rnd, accountID, domain),
				PlanOnly: true,
			},
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerTokenNoMappings(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("remote_uri"), knownvalue.StringExact("scim2.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("deactivate_on_delete"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scheme"), knownvalue.StringExact("oauthbearertoken")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("token"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings"), knownvalue.Null()),
				},
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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuthBearerToken(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("remote_uri"), knownvalue.StringExact("scim.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("deactivate_on_delete"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scheme"), knownvalue.StringExact("oauthbearertoken")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("token"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("schema"), knownvalue.StringExact("urn:ietf:params:scim:schemas:core:2.0:User")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("filter"), knownvalue.StringExact("title pr or userType eq \"Intern\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("transform_jsonata"), knownvalue.StringExact("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("create"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("update"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("delete"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "scim_config.authentication.token", "auto_redirect_to_identity"},
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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationSCIMConfigValidOAuth2(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("remote_uri"), knownvalue.StringExact("scim.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("deactivate_on_delete"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scheme"), knownvalue.StringExact("oauth2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("client_id"), knownvalue.StringExact("beepboop")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("client_secret"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("authorization_url"), knownvalue.StringExact("https://www.authorization.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("scopes"), knownvalue.SetSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("authentication").AtMapKey("token_url"), knownvalue.StringExact("https://www.token.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("schema"), knownvalue.StringExact("urn:ietf:params:scim:schemas:core:2.0:User")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("filter"), knownvalue.StringExact("title pr or userType eq \"Intern\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("transform_jsonata"), knownvalue.StringExact("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("create"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("update"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scim_config").AtMapKey("mappings").AtSliceIndex(0).AtMapKey("operations").AtMapKey("delete"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "scim_config.authentication.client_secret", "auto_redirect_to_identity"},
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
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithCORS(rnd, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_methods"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_origins"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("max_age"), knownvalue.Int64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("zones/%s/", zoneID),
				ImportStateVerifyIgnore: []string{"destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains"},
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
	resourceName := name
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("sp_entity_id"), knownvalue.StringExact("saas-app.example")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("consumer_service_url"), knownvalue.StringExact("https://saas-app.example/sso/saml/consume")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("name_id_format"), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("default_relay_state"), knownvalue.StringExact("https://saas-app.example")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("name_id_transform_jsonata"), knownvalue.StringExact("$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("saml_attribute_transform_jsonata"), knownvalue.StringExact("$ ~>| groups | {'group_name': name} |")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("idp_entity_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("public_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("sso_endpoint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("name_format"), knownvalue.StringExact("urn:oasis:names:tc:SAML:2.0:attrname-format:basic")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("source").AtMapKey("name"), knownvalue.StringExact("user_email")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("rank")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("source").AtMapKey("name"), knownvalue.StringExact("rank")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("friendly_name"), knownvalue.StringExact("Rank")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("required"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains"},
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
	resourceName := name

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("saas")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("sp_entity_id"), knownvalue.StringExact("saas-app.example")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("consumer_service_url"), knownvalue.StringExact("https://saas-app.example/sso/saml/consume")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("name_id_format"), knownvalue.StringExact("email")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("default_relay_state"), knownvalue.StringExact("https://saas-app.example")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("name_id_transform_jsonata"), knownvalue.StringExact("$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("saml_attribute_transform_jsonata"), knownvalue.StringExact("$ ~>| groups | {'group_name': name} |")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("email")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("name_format"), knownvalue.StringExact("urn:oasis:names:tc:SAML:2.0:attrname-format:basic")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(0).AtMapKey("source").AtMapKey("name"), knownvalue.StringExact("user_email")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("rank")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("source").AtMapKey("name"), knownvalue.StringExact("rank")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("friendly_name"), knownvalue.StringExact("Rank")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saas_app").AtMapKey("custom_attributes").AtSliceIndex(1).AtMapKey("required"), knownvalue.Bool(true)),
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:            testAccCloudflareAccessApplicationConfigWithSAMLSaas(rnd, accountID),
				ConfigStateChecks: stateChecks,
			},
			{
				ImportState:         true,
				ImportStateVerify:   true,
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("accounts/%s/", accountID),
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
	resourceName := name
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
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "tags", "auto_redirect_to_identity"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}

					policiesCount := s[0].Attributes["policies.#"]
					if policiesCount != "2" {
						return fmt.Errorf("expected 2 policies, got %s", policiesCount)
					}

					if s[0].Attributes["policies.0.id"] == "" {
						return fmt.Errorf("expected policy ID to be preserved")
					}
					if s[0].Attributes["policies.1.id"] == "" {
						return fmt.Errorf("expected policy ID to be preserved")
					}

					if _, ok := s[0].Attributes["policies.0.name"]; ok {
						return fmt.Errorf("expected policy name to be nullified")
					}
					if _, ok := s[0].Attributes["policies.0.decision"]; ok {
						return fmt.Errorf("expected policy decision to be nullified")
					}
					if _, ok := s[0].Attributes["policies.0.include.#"]; ok {
						return fmt.Errorf("expected policy include to be nullified")
					}

					if _, ok := s[0].Attributes["skip_interstitial"]; ok {
						return fmt.Errorf("expected skip_interstitial to be nullified")
					}
					if _, ok := s[0].Attributes["allow_iframe"]; ok {
						return fmt.Errorf("expected allow_iframe to be nullified")
					}
					if _, ok := s[0].Attributes["path_cookie_attribute"]; ok {
						return fmt.Errorf("expected path_cookie_attribute to be nullified")
					}

					return nil
				},
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

func TestAccCloudflareAccessApplication_WarpInvalid(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accoundID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareAccessApplicationWarpInvalid(rnd, accoundID),
				ExpectError: regexp.MustCompile(`"allow_authenticate_via_warp" can only be set if "type" is one of:\s"self_hosted", "ssh", "vnc", "rdp", "saas", "dash_sso"`),
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

func testAccCloudflareAccessApplicationWarpInvalid(rnd, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwarpinvalid.tf", rnd, accountID)
}

func TestAccCloudflareAccessApplication_BooleanFieldsPersistence(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigBooleanFields(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("skip_interstitial"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("path_cookie_attribute"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allow_credentials"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_headers"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_methods"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cors_headers").AtMapKey("allowed_origins"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				// Ensures no diff on second plan - this is the key test for boolean persistence issues
				Config:   testAccCloudflareAccessApplicationConfigBooleanFields(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_AllowIframeFalsePersistence(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigAllowIframeFalse(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Bool(false)),
				},
			},
			{
				// Test that omitting allow_iframe doesn't cause a diff when API returns false
				Config: testAccCloudflareAccessApplicationConfigAllowIframeOmitted(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Should be normalized to null without causing a diff
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Null()),
				},
			},
			{
				// Ensures no diff on subsequent plan
				Config:   testAccCloudflareAccessApplicationConfigAllowIframeOmitted(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessApplication_BooleanFieldTransitions(t *testing.T) {
	t.Skip("Account-level WARP setting keep gets toggled off")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				// Start with boolean fields set to true
				Config: testAccCloudflareAccessApplicationConfigBooleanFieldsTrue(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("skip_interstitial"), knownvalue.Bool(true)),
				},
			},
			{
				// Change to false
				Config: testAccCloudflareAccessApplicationConfigBooleanFieldsFalse(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("skip_interstitial"), knownvalue.Bool(false)),
				},
			},
			{
				// Remove the boolean fields entirely (should not cause drift)
				Config: testAccCloudflareAccessApplicationConfigBooleanFieldsOmitted(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Should be normalized to null without drift
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_iframe"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("skip_interstitial"), knownvalue.Null()),
				},
			},
			{
				// Ensures no diff on final plan
				Config:   testAccCloudflareAccessApplicationConfigBooleanFieldsOmitted(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareAccessApplicationConfigBooleanFields(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id                     = "%[3]s"
  name                           = "%[1]s"
  domain                         = "%[1]s.%[2]s"
  type                           = "self_hosted"
  session_duration               = "24h"
  allow_iframe                   = false
  skip_interstitial              = false
  allow_authenticate_via_warp    = false
  path_cookie_attribute          = false
  cors_headers = {
    allowed_headers    = ["x-custom-header"]
    allowed_methods    = ["GET", "POST"]
    allowed_origins    = ["https://example.com"]
    allow_credentials  = false
    max_age            = 300
  }
}
`, rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigAllowIframeFalse(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[3]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[2]s"
  type             = "self_hosted"
  session_duration = "24h"
  allow_iframe     = false
}
`, rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigAllowIframeOmitted(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[3]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[2]s"
  type             = "self_hosted"
  session_duration = "24h"
}
`, rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigBooleanFieldsTrue(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id                     = "%[3]s"
  name                           = "%[1]s"
  domain                         = "%[1]s.%[2]s"
  type                           = "self_hosted"
  session_duration               = "24h"
  allow_iframe                   = true
  skip_interstitial              = true
  allow_authenticate_via_warp    = true
  path_cookie_attribute          = true
}
`, rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigBooleanFieldsFalse(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id                     = "%[3]s"
  name                           = "%[1]s"
  domain                         = "%[1]s.%[2]s"
  type                           = "self_hosted"
  session_duration               = "24h"
  allow_iframe                   = false
  skip_interstitial              = false
  allow_authenticate_via_warp    = false
  path_cookie_attribute          = false
}
`, rnd, domain, accountID)
}

func testAccCloudflareAccessApplicationConfigBooleanFieldsOmitted(rnd, domain, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[3]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[2]s"
  type             = "self_hosted"
  session_duration = "24h"
}
`, rnd, domain, accountID)
}

func TestAccCloudflareAccessApplication_TagsOrderIgnored(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_application.%s", rnd)
	resourceName := name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessApplicationConfigWithTagsOrdering(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("self_hosted")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tags"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("ccc"),
						knownvalue.StringExact("aaa"),
						knownvalue.StringExact("bbb"),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"service_auth_401_redirect", "destinations", "enable_binding_cookie", "options_preflight_bypass", "self_hosted_domains", "tags", "auto_redirect_to_identity"},
			},
			{
				Config: testAccCloudflareAccessApplicationConfigWithTagsOrdering(rnd, domain, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tags"), knownvalue.ListSizeExact(3)),
				},
			},
			{
				Config:   testAccCloudflareAccessApplicationConfigWithTagsOrdering(rnd, domain, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareAccessApplicationConfigWithTagsOrdering(rnd, domain, accountID string) string {
	return acctest.LoadTestCase("accessapplicationconfigwithtagsordering.tf", rnd, domain, accountID)
}
