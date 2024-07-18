package worker_domain_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	scriptContent = `addEventListener('fetch', event => {event.respondWith(new Response('test'))});`
)

func TestAccCloudflareWorkerDomain_Attach(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	var domain cloudflare.WorkersDomain
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_worker_domain." + rnd
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
					resource.TestCheckResourceAttr(name, "service", rnd),
				),
			},
			{
				ImportState:         true,
				ImportStateVerify:   true,
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
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
		if rs.Type != "cloudflare_worker_domain" {
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
