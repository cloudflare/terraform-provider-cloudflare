package web_analytics_site_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_web_analytics_site", &resource.Sweeper{
		Name: "cloudflare_web_analytics_site",
		F:    testSweepCloudflareWebAnalyticsSites,
	})
}

func testSweepCloudflareWebAnalyticsSites(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping web analytics sites sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	sites, _, err := client.ListWebAnalyticsSites(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWebAnalyticsSitesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch web analytics sites: %s", err))
		return fmt.Errorf("failed to fetch web analytics sites: %w", err)
	}

	if len(sites) == 0 {
		tflog.Info(ctx, "No web analytics sites to sweep")
		return nil
	}

	for _, site := range sites {
		// Use standard filtering helper on the site tag field
		if !utils.ShouldSweepResource(site.SiteTag) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting web analytics site: %s (account: %s)", site.SiteTag, accountID))
		_, err := client.DeleteWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.DeleteWebAnalyticsSiteParams{
			SiteTag: site.SiteTag,
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete web analytics site %s: %s", site.SiteTag, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted web analytics site: %s", site.SiteTag))
	}

	return nil
}

func TestAccCloudflareWebAnalyticsSite_Create_ImportState(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web_analytics_site.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWebAnalyticsSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWebAnalyticsSite(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "site_tag"),
					resource.TestCheckResourceAttr(name, "host", domain),
					resource.TestCheckResourceAttr(name, "auto_install", "false"),
					resource.TestCheckResourceAttrSet(name, "site_token"),
				),
			},
			{
				ResourceName: name,
				PlanOnly:     true,
				ImportState:  true,
				ImportPlanChecks: resource.ImportPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ImportStateIdFunc: testAccCloudflareWebAnalyticsSiteImportStateIdFunc(name),
			},
		},
	})
}

func testAccCheckCloudflareWebAnalyticsSiteDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_web_analytics_site" {
			continue
		}

		_, err := client.GetWebAnalyticsSite(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cloudflare.GetWebAnalyticsSiteParams{
			SiteTag: rs.Primary.Attributes["site_tag"],
		})
		if err == nil {
			return fmt.Errorf("web analytics site still exists")
		}
	}

	return nil
}

func testAccCloudflareWebAnalyticsSite(resourceName, accountID, domain string) string {
	return acctest.LoadTestCase("webanalyticssite.tf", resourceName, accountID, domain)
}

func testAccCloudflareWebAnalyticsSiteImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		accountId := rs.Primary.Attributes["account_id"]
		siteTag := rs.Primary.Attributes["site_tag"]

		return fmt.Sprintf("%s/%s", accountId, siteTag), nil
	}
}
