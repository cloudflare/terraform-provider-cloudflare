package workers_custom_domain_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
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
	resource.AddTestSweepers("cloudflare_workers_custom_domain", &resource.Sweeper{
		Name: "cloudflare_workers_custom_domain",
		F:    testSweepCloudflareWorkersCustomDomains,
	})
}

func testSweepCloudflareWorkersCustomDomains(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping workers custom domains sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	page, err := client.Workers.Domains.List(ctx, workers.DomainListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch workers custom domains: %s", err))
		return fmt.Errorf("failed to fetch workers custom domains: %w", err)
	}

	for page != nil && len(page.Result) > 0 {
		for _, domain := range page.Result {
			if !utils.ShouldSweepResource(domain.Hostname) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting workers custom domain: %s (account: %s)", domain.Hostname, accountID))
			_, err := client.Workers.Domains.Delete(ctx, domain.ID, workers.DomainDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete workers custom domain %s: %s", domain.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted workers custom domain: %s", domain.ID))
		}

		page, err = page.GetNextPage()
		if err != nil {
			break
		}
	}

	return nil
}

func TestAccCloudflareWorkersCustomDomain_Basic(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_custom_domain." + rnd
	hostname := rnd + "." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainConfig(rnd, accountID, hostname, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("environment"), knownvalue.StringExact("production")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
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

func TestAccCloudflareWorkersCustomDomain_RecreateOnHostnameChange(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_custom_domain." + rnd
	hostname1 := rnd + "." + zoneName
	hostname2 := rnd + "-new." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainConfig(rnd, accountID, hostname1, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname1)),
				},
			},
			{
				Config: testAccCloudflareWorkersCustomDomainConfig(rnd, accountID, hostname2, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname2)),
				},
			},
		},
	})
}

func TestAccCloudflareWorkersCustomDomain_WithZoneName(t *testing.T) {
	t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_custom_domain." + rnd
	hostname := rnd + "." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainWithZoneNameConfig(rnd, accountID, hostname, zoneName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareWorkersCustomDomain_MinimalConfig(t *testing.T) {
	t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_custom_domain." + rnd
	hostname := rnd + "." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainMinimalConfig(rnd, accountID, hostname),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					// zone_id and zone_name should be computed from the hostname
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_name"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCheckCloudflareWorkersCustomDomainDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_custom_domain" {
			continue
		}

		_, err := client.Workers.Domains.Get(context.Background(), rs.Primary.ID, workers.DomainGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("workers custom domain with id %s still exists", rs.Primary.ID)
		}

		var apierr *cloudflare.Error
		if errors.As(err, &apierr) && apierr.StatusCode != http.StatusNotFound {
			return fmt.Errorf("error checking if workers custom domain %s was destroyed: %w", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCloudflareWorkersCustomDomainConfig(rnd, accountID, hostname, zoneID string) string {
	return acctest.LoadTestCase("workerdomainattach.tf", rnd, accountID, hostname, zoneID)
}

func testAccCloudflareWorkersCustomDomainWithZoneNameConfig(rnd, accountID, hostname, zoneName string) string {
	return acctest.LoadTestCase("workerdomainattach_zonename.tf", rnd, accountID, hostname, zoneName)
}

func testAccCloudflareWorkersCustomDomainMinimalConfig(rnd, accountID, hostname string) string {
	return acctest.LoadTestCase("workerdomainattach_minimal.tf", rnd, accountID, hostname)
}

func TestAccUpgradeWorkersCustomDomain_FromPublishedV5(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	hostname := rnd + "." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	config := testAccCloudflareWorkersCustomDomainConfig(rnd, accountID, hostname, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
