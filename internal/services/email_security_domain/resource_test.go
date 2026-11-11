package email_security_domain_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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
	resource.AddTestSweepers("cloudflare_email_security_domain", &resource.Sweeper{
		Name: "cloudflare_email_security_domain",
		F:    testSweepCloudflareEmailSecurityDomain,
	})
}

func testSweepCloudflareEmailSecurityDomain(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	iter := client.EmailSecurity.Settings.Domains.ListAutoPaging(ctx, email_security.SettingDomainListParams{
		AccountID: cloudflare.F(accountID),
	})

	for iter.Next() {
		domain := iter.Current()
		if !utils.ShouldSweepResource(domain.Domain) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Email Security domain: %s (%s)", domain.ID, domain.Domain))
		_, err := client.EmailSecurity.Settings.Domains.Delete(ctx, domain.ID, email_security.SettingDomainDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Email Security domain %s: %s", domain.ID, err))
		}
	}

	return iter.Err()
}

func TestAccCloudflareEmailSecurityDomain_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_domain.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("domain_basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(fmt.Sprintf("%s.xzy", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_delivery_modes"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("DIRECT")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("drop_dispositions"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_restrictions"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("regions"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("US")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("transport"), knownvalue.StringExact(fmt.Sprintf("transport.%s.xzy", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lookback_hops"), knownvalue.Int64Exact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require_tls_inbound"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require_tls_outbound"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
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

func TestAccCloudflareEmailSecurityDomain_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_domain.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("domain_basic.tf", rnd, accountID),
			},
			{
				Config: acctest.LoadTestCase("domain_updated.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("drop_dispositions"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("MALICIOUS"),
						knownvalue.StringExact("SUSPICIOUS"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_restrictions"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("192.0.2.0/24"),
						knownvalue.StringExact("2001:db8::/32"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("regions"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("DE"),
						knownvalue.StringExact("US"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("transport"), knownvalue.StringExact(fmt.Sprintf("updated.%s.xzy", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lookback_hops"), knownvalue.Int64Exact(5)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require_tls_inbound"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require_tls_outbound"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func testAccCheckCloudflareEmailSecurityDomainDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_email_security_domain" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.EmailSecurity.Settings.Domains.Get(
			context.Background(),
			rs.Primary.ID,
			email_security.SettingDomainGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("email security domain still exists: %s", rs.Primary.ID)
		}

		var apiErr *email_security.Error
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return fmt.Errorf("error checking if email security domain %s was destroyed: %w", rs.Primary.ID, err)
		}
	}

	return nil
}
