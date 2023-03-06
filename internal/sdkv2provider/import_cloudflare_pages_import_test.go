package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testPagesDomains(accountID, projectName, resourceOneId, domainOne, resourceTwoId, domainTwo string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_domain" "%[3]s" {
			account_id = "%[1]s"
			project_name = "%[2]s"
			domain = "%[4]s"
		}

		resource "cloudflare_pages_domain" "%[5]s" {
			account_id = "%[1]s"
			project_name = "%[2]s"
			domain = "%[6]s"
		}
		`, accountID, projectName, resourceOneId, domainOne, resourceTwoId, domainTwo)
}

func TestAccCloudflarePagesDomain_Import(t *testing.T) {
	skipPagesProjectForNonConfiguredDefaultAccount(t)

	resourceOneId := generateRandomResourceName()
	resourceTwoId := generateRandomResourceName()
	resourceOneName := fmt.Sprintf("cloudflare_pages_domain.%s", resourceOneId)
	resourceTwoName := fmt.Sprintf("cloudflare_pages_domain.%s", resourceTwoId)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	projectName := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesDomains(
					accountID,
					projectName,
					resourceOneId,
					resourceOneId+"."+domain,
					resourceTwoId,
					resourceTwoId+"."+domain,
				),
			},
			{
				ResourceName:        resourceOneName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
			{
				ResourceName:        resourceTwoName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
