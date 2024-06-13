package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
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
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_domain." + rnd
	hostname := rnd + "." + zoneName

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerDomainDestroy,
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

	client := testAccProvider.Meta().(*cloudflare.API)
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
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s-script" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}

resource "cloudflare_worker_domain" "%[1]s" {
	zone_id = "%[5]s"
	account_id = "%[3]s"
	hostname = "%[4]s"
	service = "%[1]s"
	depends_on = [cloudflare_worker_script.%[1]s-script]
}`, rnd, scriptContent, accountID, hostname, zoneID)
}

func testAccCheckCloudflareWorkerDomainDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_domain" {
			continue
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		r, _ := client.GetWorkersDomain(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)

		if r.ID != "" {
			return fmt.Errorf("worker domain with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
