package connectivity_directory_service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareConnectivityDirectoryServiceDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_connectivity_directory_service." + rnd
	dataSourceName := "data.cloudflare_connectivity_directory_service." + rnd

	// Test configuration values
	serviceName := fmt.Sprintf("test-directory-service-%s", rnd)
	hostname := "ldap.example.com"
	resolverIP1 := "10.0.0.1"
	resolverIP2 := "10.0.0.2"
	httpPort := 389
	httpsPort := 636

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareConnectivityDirectoryServiceDataSource_Basic(
					rnd,
					accountID,
					serviceName,
					hostname,
					resolverIP1,
					resolverIP2,
					httpPort,
					httpsPort,
				),
				Check: resource.ComposeTestCheckFunc(
					// Verify the data source can read the resource
					resource.TestCheckResourceAttrPair(dataSourceName, consts.AccountIDSchemaKey, resourceName, consts.AccountIDSchemaKey),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_id", resourceName, "service_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "type", resourceName, "type"),

					// Verify host configuration matches
					resource.TestCheckResourceAttrPair(dataSourceName, "host.hostname", resourceName, "host.hostname"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host.resolver_network.tunnel_id", resourceName, "host.resolver_network.tunnel_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host.resolver_network.resolver_ips.0", resourceName, "host.resolver_network.resolver_ips.0"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host.resolver_network.resolver_ips.1", resourceName, "host.resolver_network.resolver_ips.1"),

					// Verify ports match
					resource.TestCheckResourceAttrPair(dataSourceName, "http_port", resourceName, "http_port"),
					resource.TestCheckResourceAttrPair(dataSourceName, "https_port", resourceName, "https_port"),

					// Verify computed attributes match
					resource.TestCheckResourceAttrPair(dataSourceName, "created_at", resourceName, "created_at"),
					resource.TestCheckResourceAttrPair(dataSourceName, "updated_at", resourceName, "updated_at"),
				),
			},
		},
	})
}

func testAccCloudflareConnectivityDirectoryServiceDataSource_Basic(rnd, accountID, serviceName, hostname, resolverIP1, resolverIP2 string, httpPort, httpsPort int) string {
	return acctest.LoadTestCase("datasource_basic.tf",
		rnd,
		accountID,
		serviceName,
		hostname,
		resolverIP1,
		resolverIP2,
		httpPort,
		httpsPort,
	)
}
