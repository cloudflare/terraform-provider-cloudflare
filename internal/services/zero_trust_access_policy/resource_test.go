package zero_trust_access_policy_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccessPolicy_ServiceToken(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyServiceTokenConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.service_token.%", "1"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyServiceTokenConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessPolicy_AnyServiceToken(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAnyServiceTokenConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "include.0.any_valid_service_token.%"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyAnyServiceTokenConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyServiceTokenConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyservicetokenconfig.tf", resourceID, zone, accountID)
}

func testAccessPolicyAnyServiceTokenConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyanyservicetokenconfig.tf", resourceID, zone, accountID)
}

func testAccessPolicyWithZoneID(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accesspolicywithzoneid.tf", resourceID, zone, zoneID)
}

func testAccessPolicyWithZoneIDUpdated(resourceID, zone, zoneID string) string {
	return acctest.LoadTestCase("accesspolicywithzoneidupdated.tf", resourceID, zone, zoneID)
}

func TestAccCloudflareAccessPolicy_Group(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGroupConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.group.%", "1"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyGroupConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyGroupConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicygroupconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_MTLS(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyMTLSConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "include.0.certificate.%"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyMTLSConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyMTLSConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicymtlsconfig.tf", resourceID, zone, accountID)
}

func testAccessPolicyCommonNameConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicycommonnameconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_EmailDomain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailDomainConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.domain", "example.com"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyEmailDomainConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyEmailDomainConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyemaildomainconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Emails(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailsConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.email.email", "a@example.com"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyEmailsConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyEmailsConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyemailsconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Everyone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEveryoneConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "include.0.everyone.%"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyEveryoneConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyEveryoneConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyeveryoneconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_IPs(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIPsConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.ip.ip", "10.0.0.1/32"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyIPsConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyIPsConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyipsconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_AuthMethod(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAuthMethodConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.auth_method.auth_method", "hwk"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyAuthMethodConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyAuthMethodConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyauthmethodconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Geo(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGeoConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.geo.country_code", "US"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyGeoConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyGeoConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicygeoconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Okta(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyOktaConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.okta.name", "jacob-group"),
					resource.TestCheckResourceAttr(name, "include.0.okta.identity_provider_id", "225934dc-14e4-4f55-87be-f5d798d23f91"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyOktaConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyOktaConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyoktaconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_PurposeJustification(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyPurposeJustificationConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessPolicyHasPJ(name, cloudflare.AccountIdentifier(accountID)),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "purpose_justification_required", "true"),
					resource.TestCheckResourceAttr(name, "purpose_justification_prompt", "Why should we let you in?"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyPurposeJustificationConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCheckCloudflareAccessPolicyHasPJ(n string, accessIdentifier *cloudflare.ResourceContainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessPolicy ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		var foundAccessPolicy cloudflare.AccessPolicy
		var err error

		foundAccessPolicy, err = client.GetAccessPolicy(context.Background(), accessIdentifier, cloudflare.GetAccessPolicyParams{ApplicationID: rs.Primary.Attributes["application_id"], PolicyID: rs.Primary.ID})
		if err != nil {
			return err
		}

		if foundAccessPolicy.ID != rs.Primary.ID {
			return fmt.Errorf("AccessPolicy not found")
		}

		if !(foundAccessPolicy.PurposeJustificationPrompt != nil && *foundAccessPolicy.PurposeJustificationPrompt == rs.Primary.Attributes["purpose_justification_prompt"]) {
			return fmt.Errorf("AccessPolicy is missing purpose_justification_prompt")
		}

		pjRequired, _ := strconv.ParseBool(rs.Primary.Attributes["purpose_justification_required"])

		if !(foundAccessPolicy.PurposeJustificationRequired != nil && *foundAccessPolicy.PurposeJustificationRequired == pjRequired) {
			return fmt.Errorf("AccessPolicy is missing purpose_justification_required")
		}

		return nil
	}
}

func testAccessPolicyPurposeJustificationConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicypurposejustificationconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ApprovalGroup(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyApprovalGroupConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessPolicyHasApprovalGroups(name, cloudflare.AccountIdentifier(accountID)),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "purpose_justification_required", "true"),
					resource.TestCheckResourceAttr(name, "purpose_justification_prompt", "Why should we let you in?"),
					resource.TestCheckResourceAttr(name, "approval_required", "true"),
					resource.TestCheckResourceAttr(name, "approval_groups.0.email_addresses.0", "test1@example.com"),
					resource.TestCheckResourceAttr(name, "approval_groups.0.email_addresses.1", "test2@example.com"),
					resource.TestCheckResourceAttr(name, "approval_groups.0.email_addresses.2", "test3@example.com"),
					resource.TestCheckResourceAttr(name, "approval_groups.1.email_addresses.0", "test4@example.com"),
					resource.TestCheckResourceAttr(name, "approval_groups.0.approvals_needed", "2"),
					resource.TestCheckResourceAttr(name, "approval_groups.1.approvals_needed", "1"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyApprovalGroupConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCheckCloudflareAccessPolicyHasApprovalGroups(n string, accessIdentifier *cloudflare.ResourceContainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessPolicy ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		var foundAccessPolicy cloudflare.AccessPolicy
		var err error

		foundAccessPolicy, err = client.GetAccessPolicy(context.Background(), accessIdentifier, cloudflare.GetAccessPolicyParams{ApplicationID: rs.Primary.Attributes["application_id"], PolicyID: rs.Primary.ID})
		if err != nil {
			return err
		}

		if foundAccessPolicy.ID != rs.Primary.ID {
			return fmt.Errorf("AccessPolicy not found")
		}

		if !(foundAccessPolicy.ApprovalGroups != nil) {
			return fmt.Errorf("AccessPolicy is missing approval_groups")
		}

		approvalRequired, _ := strconv.ParseBool(rs.Primary.Attributes["approval_required"])

		if !(foundAccessPolicy.ApprovalRequired != nil && *foundAccessPolicy.ApprovalRequired == approvalRequired) {
			return fmt.Errorf("AccessPolicy is missing approval_required")
		}

		return nil
	}
}

func testAccessPolicyApprovalGroupConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyapprovalgroupconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ExternalEvaluation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyExternalEvalautionConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.external_evaluation.evaluate_url", "https://example.com"),
					resource.TestCheckResourceAttr(name, "include.0.external_evaluation.keys_url", "https://example.com/keys"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyExternalEvalautionConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyExternalEvalautionConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyexternalevalautionconfig.tf", resourceID, zone, accountID)
}

/*
Commented out until cloudflare_zero_trust_gateway_settings gets fixed
func TestAccCloudflareAccessPolicy_IsolationRequired(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIsolationRequiredConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "isolation_required", "true"),
				),
			},
		},
	})
}
*/

func testAccessPolicyIsolationRequiredConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyisolationrequiredconfig.tf", resourceID, zone, accountID)
}

func testAccessPolicyReusableConfig(resourceID, accountID string) string {
	return acctest.LoadTestCase("accesspolicyreusableconfig.tf", resourceID, accountID)
}
