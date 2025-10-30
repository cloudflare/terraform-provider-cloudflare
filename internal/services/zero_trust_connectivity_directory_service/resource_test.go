package zero_trust_connectivity_directory_service_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_connectivity_directory_service", &resource.Sweeper{
		Name: "cloudflare_zero_trust_connectivity_directory_service",
		F: func(region string) error {
			client := acctest.SharedClient()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			ctx := context.Background()

			page, err := client.ZeroTrust.Connectivity.Directory.Services.List(
				ctx,
				zero_trust.ConnectivityDirectoryServiceListParams{
					AccountID: cloudflare.F(accountID),
				},
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to list Connectivity Directory Services: %s", err))
				return err
			}

			// Process all pages
			for page != nil && len(page.Result) > 0 {
				for _, service := range page.Result {
					// Skip any static test resources if needed
					// For now, delete all test resources
					err := client.ZeroTrust.Connectivity.Directory.Services.Delete(
						ctx,
						service.ServiceID,
						zero_trust.ConnectivityDirectoryServiceDeleteParams{
							AccountID: cloudflare.F(accountID),
						},
					)
					if err != nil {
						tflog.Error(ctx, fmt.Sprintf("Failed to delete Connectivity Directory Service %s: %s", service.ServiceID, err))
					}
				}

				// Get next page
				page, err = page.GetNextPage()
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to fetch next page: %s", err))
					return err
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareZeroTrustConnectivityDirectoryService_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_zero_trust_connectivity_directory_service." + rnd

	// Test configuration values
	serviceName := fmt.Sprintf("test-directory-service-%s", rnd)
	updatedServiceName := fmt.Sprintf("test-directory-service-%s-updated", rnd)
	hostname := "ldap.example.com"
	resolverIP1 := "10.0.0.1"
	resolverIP2 := "10.0.0.2"
	updatedResolverIP1 := "10.1.0.1"
	updatedResolverIP2 := "10.1.0.2"
	httpPort := 389
	httpsPort := 636
	updatedHttpPort := 3389
	updatedHttpsPort := 3636

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustConnectivityDirectoryService_Basic(
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
					// Check required attributes
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "type", "http"),

					// Check host configuration
					resource.TestCheckResourceAttr(resourceName, "host.hostname", hostname),
					resource.TestCheckResourceAttrSet(resourceName, "host.resolver_network.tunnel_id"),
					resource.TestCheckResourceAttr(resourceName, "host.resolver_network.resolver_ips.0", resolverIP1),
					resource.TestCheckResourceAttr(resourceName, "host.resolver_network.resolver_ips.1", resolverIP2),

					// Check ports
					resource.TestCheckResourceAttr(resourceName, "http_port", fmt.Sprintf("%d", httpPort)),
					resource.TestCheckResourceAttr(resourceName, "https_port", fmt.Sprintf("%d", httpsPort)),

					// Check computed attributes exist
					resource.TestCheckResourceAttrSet(resourceName, "service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Update test step
			{
				Config: testAccCloudflareZeroTrustConnectivityDirectoryService_Update(
					rnd,
					accountID,
					updatedServiceName,
					hostname,
					updatedResolverIP1,
					updatedResolverIP2,
					updatedHttpPort,
					updatedHttpsPort,
				),
				Check: resource.ComposeTestCheckFunc(
					// Check required attributes
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", updatedServiceName),
					resource.TestCheckResourceAttr(resourceName, "type", "http"),

					// Check host configuration (hostname stays the same)
					resource.TestCheckResourceAttr(resourceName, "host.hostname", hostname),
					resource.TestCheckResourceAttrSet(resourceName, "host.resolver_network.tunnel_id"),
					resource.TestCheckResourceAttr(resourceName, "host.resolver_network.resolver_ips.0", updatedResolverIP1),
					resource.TestCheckResourceAttr(resourceName, "host.resolver_network.resolver_ips.1", updatedResolverIP2),

					// Check updated ports
					resource.TestCheckResourceAttr(resourceName, "http_port", fmt.Sprintf("%d", updatedHttpPort)),
					resource.TestCheckResourceAttr(resourceName, "https_port", fmt.Sprintf("%d", updatedHttpsPort)),

					// Check computed attributes exist (service_id should remain the same)
					resource.TestCheckResourceAttrSet(resourceName, "service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Import test
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCloudflareZeroTrustConnectivityDirectoryService_Basic(rnd, accountID, serviceName, hostname, resolverIP1, resolverIP2 string, httpPort, httpsPort int) string {
	return acctest.LoadTestCase("basic.tf",
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

func testAccCloudflareZeroTrustConnectivityDirectoryService_Update(rnd, accountID, serviceName, hostname, resolverIP1, resolverIP2 string, httpPort, httpsPort int) string {
	return acctest.LoadTestCase("update.tf",
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