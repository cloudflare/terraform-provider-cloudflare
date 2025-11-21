package pages_domain_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflarePagesDomain(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_domain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()
	zonePage, err := client.Zones.List(context.Background(), zones.ZoneListParams{
		Account: cloudflare.F(zones.ZoneListParamsAccount{
			ID: cloudflare.F(accountID),
		}),
	})
	if err != nil || zonePage == nil || len(zonePage.Result) == 0 {
		t.Skip("No zones available in account for testing")
	}

	domain := zonePage.Result[0].Name
	fullDomain := rnd + "." + domain

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePagesDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("pagesdomainconfig.tf", rnd, accountID, rnd, fullDomain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("project_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(fullDomain)),
				},
				Check: testAccCheckCloudflarePagesDomainExists(name, accountID, rnd, fullDomain),
				// ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateIdPrefix: fmt.Sprintf("%s/%s/", accountID, rnd),
				ImportStateVerify:   true,
				ImportStateVerifyIgnore: []string{
					"status",
					"validation_data",
					"verification_data",
				},
			},
		},
	})
}

func testAccCheckCloudflarePagesDomainDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "cloudflare_pages_domain" {
			accountID := rs.Primary.Attributes["account_id"]
			projectName := rs.Primary.Attributes["project_name"]
			domainName := rs.Primary.Attributes["name"]

			_, err := client.Pages.Projects.Domains.Get(
				context.Background(),
				projectName,
				domainName,
				pages.ProjectDomainGetParams{
					AccountID: cloudflare.F(accountID),
				},
			)

			if err == nil {
				return fmt.Errorf("pages domain %s for project %s still exists", domainName, projectName)
			}
		}

		if rs.Type == "cloudflare_pages_project" {
			accountID := rs.Primary.Attributes["account_id"]
			projectName := rs.Primary.Attributes["name"]

			_, err := client.Pages.Projects.Get(
				context.Background(),
				projectName,
				pages.ProjectGetParams{
					AccountID: cloudflare.F(accountID),
				},
			)

			if err == nil {
				return fmt.Errorf("pages project %s still exists", projectName)
			}
		}
	}

	return nil
}

func testAccCheckCloudflarePagesDomainExists(resourceName, accountID, projectName, domainName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no pages domain ID is set")
		}

		client := acctest.SharedClient()
		_, err := client.Pages.Projects.Domains.Get(
			context.Background(),
			projectName,
			domainName,
			pages.ProjectDomainGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err != nil {
			return fmt.Errorf("pages domain not found: %s", err)
		}

		return nil
	}
}
