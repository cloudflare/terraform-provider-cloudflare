package connectivity_directory_service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareConnectivityDirectoryService_DataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_connectivity_directory_service." + rnd
	dataSourceName := "data.cloudflare_connectivity_directory_service." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConnectivityDirectoryServiceDataSource(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_id", resourceName, "service_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host.ipv4", resourceName, "host.ipv4"),
				),
			},
		},
	})
}

func TestAccCloudflareConnectivityDirectoryServices_DataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	listDataSourceName := "data.cloudflare_connectivity_directory_services." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConnectivityDirectoryServicesListDataSource(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(listDataSourceName, "services.#"),
				),
			},
		},
	})
}

func testAccConnectivityDirectoryServiceDataSource(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_connectivity_directory_service" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "http"

  host {
    ipv4 = "192.168.1.100"
  }
}

data "cloudflare_connectivity_directory_service" "%[1]s" {
  depends_on = [cloudflare_connectivity_directory_service.%[1]s]
  account_id = "%[2]s"
  service_id = cloudflare_connectivity_directory_service.%[1]s.service_id
}`, rnd, accountID)
}

func testAccConnectivityDirectoryServicesListDataSource(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_connectivity_directory_service" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "http"

  host {
    ipv4 = "192.168.1.100"
  }
}

data "cloudflare_connectivity_directory_services" "%[1]s" {
  depends_on = [cloudflare_connectivity_directory_service.%[1]s]
  account_id = "%[2]s"
  type       = "http"
}`, rnd, accountID)
}
