package zero_trust_access_policy_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareAccessPolicy_ServiceToken(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyServiceTokenConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("service_token"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAnyServiceTokenConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("any_valid_service_token"), knownvalue.MapSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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

func TestAccCloudflareAccessPolicy_Group(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGroupConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("group"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyMTLSConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("certificate"), knownvalue.MapSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailDomainConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("12h")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailsConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("a@example.com")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEveryoneConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("everyone"), knownvalue.MapSizeExact(0)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIPsConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("10.0.0.1/32")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAuthMethodConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("auth_method").AtMapKey("auth_method"), knownvalue.StringExact("hwk")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGeoConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("geo").AtMapKey("country_code"), knownvalue.StringExact("US")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyOktaConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("okta").AtMapKey("name"), knownvalue.StringExact("jacob-group")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("okta").AtMapKey("identity_provider_id"), knownvalue.StringExact("225934dc-14e4-4f55-87be-f5d798d23f91")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyPurposeJustificationConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_prompt"), knownvalue.StringExact("Why should we let you in?")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyPurposeJustificationConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

// Removed old helper function testAccCheckCloudflareAccessPolicyHasPJ as it used outdated cloudflare-go v1 API
// Modern tests use ConfigStateChecks with statecheck.ExpectKnownValue instead

func testAccessPolicyPurposeJustificationConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicypurposejustificationconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ApprovalGroup(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyApprovalGroupConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_prompt"), knownvalue.StringExact("Why should we let you in?")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("email_addresses").AtSliceIndex(0), knownvalue.StringExact("test1@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("email_addresses").AtSliceIndex(1), knownvalue.StringExact("test2@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("email_addresses").AtSliceIndex(2), knownvalue.StringExact("test3@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(1).AtMapKey("email_addresses").AtSliceIndex(0), knownvalue.StringExact("test4@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("approvals_needed"), knownvalue.Int64Exact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(1).AtMapKey("approvals_needed"), knownvalue.Int64Exact(1)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyApprovalGroupConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

// Removed old helper function testAccCheckCloudflareAccessPolicyHasApprovalGroups as it used outdated cloudflare-go v1 API
// Modern tests use ConfigStateChecks with statecheck.ExpectKnownValue instead

func testAccessPolicyApprovalGroupConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyapprovalgroupconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ExternalEvaluation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyExternalEvalautionConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("external_evaluation").AtMapKey("evaluate_url"), knownvalue.StringExact("https://example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("external_evaluation").AtMapKey("keys_url"), knownvalue.StringExact("https://example.com/keys")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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

func TestAccCloudflareAccessPolicy_ExcludeRules(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyExcludeRulesConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("everyone"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("blocked@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyExcludeRulesConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyExcludeRulesConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyexcluderulesconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_RequireRules(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyRequireRulesConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("geo").AtMapKey("country_code"), knownvalue.StringExact("US")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyRequireRulesConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyRequireRulesConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyrequirerulesconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_DecisionTypes(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	denyResourceName := "cloudflare_zero_trust_access_policy." + rnd + "_deny"
	bypassResourceName := "cloudflare_zero_trust_access_policy." + rnd + "_bypass"
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyDecisionTypesConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check deny policy
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-deny")),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("decision"), knownvalue.StringExact("deny")),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("blocked@example.com")),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(denyResourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(false)),
					// Check bypass policy
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-bypass")),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("decision"), knownvalue.StringExact("bypass")),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("127.0.0.1/32")),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(bypassResourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:        denyResourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				ResourceName:        bypassResourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyDecisionTypesConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyDecisionTypesConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicydecisiontypesconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_IsolationRequired(t *testing.T) {
	t.Skip("this test depends on zero trust gateway settings, which first must be modernized and patched")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIsolationRequiredConfig(rnd, zone, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessPolicyIsolationRequiredConfig(rnd, zone, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccessPolicyIsolationRequiredConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyisolationrequiredconfig.tf", resourceID, zone, accountID)
}

func testAccessPolicyReusableConfig(resourceID, accountID string) string {
	return acctest.LoadTestCase("accesspolicyreusableconfig.tf", resourceID, accountID)
}

func testAccCheckCloudflareZeroTrustAccessPolicyDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_policy" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.ZeroTrust.Access.Policies.Get(
			context.Background(),
			rs.Primary.ID,
			zero_trust.AccessPolicyGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("zero trust access policy still exists")
		}
	}

	return nil
}

func TestAccCloudflareAccessPolicy_DenyOnly(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_policy.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyDenyOnlyConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "decision", "deny"),
					resource.TestCheckResourceAttr(resourceName, "include.0.email.email", "blocked@example.com"),
					resource.TestCheckResourceAttr(resourceName, "approval_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "isolation_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "purpose_justification_required", "false"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccessPolicyDenyOnlyConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicydenyconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_BypassOnly(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_policy.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyBypassOnlyConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "decision", "bypass"),
					resource.TestCheckResourceAttr(resourceName, "include.0.ip.ip", "127.0.0.1/32"),
					resource.TestCheckResourceAttr(resourceName, "approval_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "isolation_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "purpose_justification_required", "false"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccessPolicyBypassOnlyConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicybypassconfig.tf", resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_OptionalBooleans(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_policy.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyOptionalBooleansConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "decision", "allow"),
					resource.TestCheckResourceAttr(resourceName, "include.0.auth_method.auth_method", "hwk"),
					// Verify that omitted boolean fields are not present in state or default to false
					resource.TestCheckNoResourceAttr(resourceName, "approval_required"),
					resource.TestCheckNoResourceAttr(resourceName, "isolation_required"),
					resource.TestCheckNoResourceAttr(resourceName, "purpose_justification_required"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				// Skip verification of boolean fields as they get normalized during import
				ImportStateVerifyIgnore: []string{"approval_required", "isolation_required", "purpose_justification_required"},
			},
		},
	})
}

func testAccessPolicyOptionalBooleansConfig(resourceID, zone, accountID string) string {
	return acctest.LoadTestCase("accesspolicyoptionalboolsconfig.tf", resourceID, zone, accountID)
}
