package email_security_trusted_domains_test

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
	resource.AddTestSweepers("cloudflare_email_security_trusted_domains", &resource.Sweeper{
		Name: "cloudflare_email_security_trusted_domains",
		F:    testSweepCloudflareEmailSecurityTrustedDomains,
	})
}

func testSweepCloudflareEmailSecurityTrustedDomains(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	iter := client.EmailSecurity.Settings.TrustedDomains.ListAutoPaging(ctx, email_security.SettingTrustedDomainListParams{
		AccountID: cloudflare.F(accountID),
	})

	for iter.Next() {
		domain := iter.Current()

		if !utils.ShouldSweepResource(domain.Pattern) && !utils.ShouldSweepResource(domain.Comments) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Email Security trusted domain: %s (%s)", domain.ID, domain.Pattern))
		_, err := client.EmailSecurity.Settings.TrustedDomains.Delete(ctx, domain.ID, email_security.SettingTrustedDomainDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Email Security trusted domain %s: %s", domain.ID, err))
		}
	}

	if err := iter.Err(); err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list Email Security trusted domains: %s", err))
		return err
	}

	return nil
}

// TestAccCloudflareEmailSecurityTrustedDomains_Basic covers create, read back and
// import. is_recent is set because the API rejects an entry that enables neither
// is_recent nor is_similarity.
func TestAccCloudflareEmailSecurityTrustedDomains_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_trusted_domains.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityTrustedDomainsDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("trusteddomains_basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recent"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_similarity"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(rnd)),
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

// TestAccCloudflareEmailSecurityTrustedDomains_Update flips both disposition
// flags and changes the pattern, asserting the change is applied in place: only
// account_id forces replacement.
func TestAccCloudflareEmailSecurityTrustedDomains_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_trusted_domains.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityTrustedDomainsDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("trusteddomains_basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recent"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_similarity"), knownvalue.Bool(false)),
				},
			},
			{
				Config: acctest.LoadTestCase("trusteddomains_updated.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s-updated.example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recent"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_similarity"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(fmt.Sprintf("%s updated", rnd))),
				},
			},
		},
	})
}

// TestAccCloudflareEmailSecurityTrustedDomains_SimilarityOnly enables only
// is_similarity. is_recent and is_regex are required, so the fixture sets them
// to false explicitly; the empty post-refresh plan guards against drift on
// those false values.
func TestAccCloudflareEmailSecurityTrustedDomains_SimilarityOnly(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_trusted_domains.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityTrustedDomainsDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("trusteddomains_similarity.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.example.net", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recent"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_similarity"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(false)),
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

// TestAccCloudflareEmailSecurityTrustedDomains_Regex covers is_regex = true,
// where the pattern is validated as a regex rather than as a literal domain.
// is_recent is required, so the fixture sets it to false and relies on
// is_similarity for the API rule that one of the two must be enabled.
func TestAccCloudflareEmailSecurityTrustedDomains_Regex(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_trusted_domains.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityTrustedDomainsDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("trusteddomains_regex.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf(`.*\.%s\.example\.com`, rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_recent"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_similarity"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(true)),
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

func testAccCheckCloudflareEmailSecurityTrustedDomainsDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_email_security_trusted_domains" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.EmailSecurity.Settings.TrustedDomains.Get(
			context.Background(),
			rs.Primary.ID,
			email_security.SettingTrustedDomainGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("email security trusted domain still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
