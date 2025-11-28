package workers_custom_domain_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cfv3 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	scriptContent = `addEventListener('fetch', event => {event.respondWith(new Response('test'))});`
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

	domains, err := client.Workers.Domains.List(ctx, workers.DomainListParams{
		AccountID: cfv3.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch workers custom domains: %s", err))
		return fmt.Errorf("failed to fetch workers custom domains: %w", err)
	}

	if len(domains.Result) == 0 {
		tflog.Info(ctx, "No workers custom domains to sweep")
		return nil
	}

	for _, domain := range domains.Result {
		// Use standard filtering helper on the hostname field
		if !utils.ShouldSweepResource(domain.Hostname) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting workers custom domain: %s (account: %s)", domain.Hostname, accountID))
		err := client.Workers.Domains.Delete(ctx, domain.ID, workers.DomainDeleteParams{
			AccountID: cfv3.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete workers custom domain %s: %s", domain.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted workers custom domain: %s", domain.ID))
	}

	return nil
}

func TestAccCloudflareWorkerDomain_Attach(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	var domain cloudflare.WorkersDomain
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_custom_domain." + rnd
	hostname := rnd + "." + zoneName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerDomainAttach(rnd, accountID, hostname, zoneID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerDomainExists(name, &domain),
					resource.TestCheckResourceAttr(name, "hostname", hostname),
					// resource.TestCheckResourceAttr(name, "service", rnd),
					resource.TestCheckResourceAttr(name, "service", "mute-truth-fdb1"), // while we fix workers_script
				),
			},
			// {
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			// },
		},
	})
}

func getDomainFromApi(accountID, domainID string) (cloudflare.WorkersDomain, error) {
	if accountID == "" {
		return cloudflare.WorkersDomain{}, fmt.Errorf("accountID is required to get a domain")
	}
	if domainID == "" {
		return cloudflare.WorkersDomain{}, fmt.Errorf("domainID is required to get a domain")
	}

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}
	domain, err := client.GetWorkersDomain(context.Background(), cloudflare.AccountIdentifier(accountID), domainID)
	if err != nil {
		fmt.Print(err.Error())
		return cloudflare.WorkersDomain{}, err
	}

	return domain, nil
}

func testAccCheckCloudflareWorkerDomainExists(resourceName string, domain *cloudflare.WorkersDomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worker Domain ID is set")
		}

		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		domainID := rs.Primary.ID
		foundDomain, err := getDomainFromApi(accountID, domainID)
		if err != nil {
			return err
		}

		if foundDomain.ID != domainID {
			return fmt.Errorf("worker domain with id %s not found", domainID)
		}

		*domain = foundDomain
		return nil
	}
}

func testAccCheckCloudflareWorkerDomainAttach(rnd, accountID string, hostname string, zoneID string) string {
	return acctest.LoadTestCase("workerdomainattach.tf", rnd, scriptContent, accountID, hostname, zoneID)
}

func testAccCheckCloudflareWorkerDomainDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_custom_domain" {
			continue
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		r, _ := client.GetWorkersDomain(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)

		if r.ID != "" {
			return fmt.Errorf("worker domain with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
