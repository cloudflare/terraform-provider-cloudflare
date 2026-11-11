package email_security_allow_policy_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_security_allow_policy", &resource.Sweeper{
		Name: "cloudflare_email_security_allow_policy",
		F:    testSweepCloudflareEmailSecurityAllowPolicy,
	})
}

func testSweepCloudflareEmailSecurityAllowPolicy(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	iter := client.EmailSecurity.Settings.AllowPolicies.ListAutoPaging(ctx, email_security.SettingAllowPolicyListParams{
		AccountID: cloudflare.F(accountID),
	})

	for iter.Next() {
		policy := iter.Current()
		if !utils.ShouldSweepResource(policy.Pattern) && !utils.ShouldSweepResource(policy.Comments) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Email Security allow policy: %s (%s)", policy.ID, policy.Pattern))
		_, err := client.EmailSecurity.Settings.AllowPolicies.Delete(ctx, policy.ID, email_security.SettingAllowPolicyDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Email Security allow policy %s: %s", policy.ID, err))
		}
	}

	return iter.Err()
}

func TestAccCloudflareEmailSecurityAllowPolicy_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_allow_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityAllowPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("allowpolicy_basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_trusted_sender"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_acceptable_sender"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_exempt_recipient"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("verify_sender"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_spoof"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recipient"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_sender"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
				},
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

func TestAccCloudflareEmailSecurityAllowPolicy_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_allow_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityAllowPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("allowpolicy_basic.tf", rnd, accountID),
			},
			{
				Config: acctest.LoadTestCase("allowpolicy_updated.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s_updated@example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("EMAIL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_acceptable_sender"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_trusted_sender"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("verify_sender"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(fmt.Sprintf("%s updated", rnd))),
				},
			},
		},
	})
}

// TestAccCloudflareEmailSecurityAllowPolicy_ExemptRecipient uses an EMAIL
// pattern: the API rejects is_exempt_recipient for DOMAIN and IP.
func TestAccCloudflareEmailSecurityAllowPolicy_ExemptRecipient(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_allow_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityAllowPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("allowpolicy_exempt_recipient.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s@example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("EMAIL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_exempt_recipient"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.Null()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareEmailSecurityAllowPolicy_Regex(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_allow_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityAllowPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("allowpolicy_regex.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf(`.*@%s\.example\.com`, rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func testAccCheckCloudflareEmailSecurityAllowPolicyDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_email_security_allow_policy" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.EmailSecurity.Settings.AllowPolicies.Get(
			context.Background(),
			rs.Primary.ID,
			email_security.SettingAllowPolicyGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("email security allow policy still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
