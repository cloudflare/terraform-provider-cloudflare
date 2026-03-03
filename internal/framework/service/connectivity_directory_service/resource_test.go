package connectivity_directory_service_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_connectivity_directory_service", &resource.Sweeper{
		Name: "cloudflare_connectivity_directory_service",
		F: func(region string) error {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			if accountID == "" {
				return fmt.Errorf("CLOUDFLARE_ACCOUNT_ID must be set")
			}

			client := cfv6.NewClient(
				option.WithAPIKey(os.Getenv("CLOUDFLARE_API_KEY")),
				option.WithAPIEmail(os.Getenv("CLOUDFLARE_EMAIL")),
			)

			ctx := context.Background()
			iter := client.Connectivity.Directory.Services.ListAutoPaging(ctx, connectivity.DirectoryServiceListParams{
				AccountID: cfv6.F(accountID),
			})

			for iter.Next() {
				svc := iter.Current()
				_ = client.Connectivity.Directory.Services.Delete(ctx, svc.ServiceID, connectivity.DirectoryServiceDeleteParams{
					AccountID: cfv6.F(accountID),
				})
			}

			return iter.Err()
		},
	})
}

func TestAccCloudflareConnectivityDirectoryService_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_connectivity_directory_service." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConnectivityDirectoryServiceIPv4(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "http"),
					resource.TestCheckResourceAttr(resourceName, "host.ipv4", "192.168.1.100"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			{
				// Import state
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareConnectivityDirectoryService_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_connectivity_directory_service." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConnectivityDirectoryServiceIPv4(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "host.ipv4", "192.168.1.100"),
				),
			},
			{
				Config: testAccConnectivityDirectoryServiceIPv4Updated(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, "host.ipv4", "192.168.1.200"),
					resource.TestCheckResourceAttr(resourceName, "http_port", "8080"),
				),
			},
		},
	})
}

func TestAccCloudflareConnectivityDirectoryService_WithPorts(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_connectivity_directory_service." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConnectivityDirectoryServiceWithPorts(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "host.ipv4", "10.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "http_port", "80"),
					resource.TestCheckResourceAttr(resourceName, "https_port", "443"),
				),
			},
		},
	})
}

func testAccConnectivityDirectoryServiceIPv4(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_connectivity_directory_service" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "http"

  host {
    ipv4 = "192.168.1.100"
  }
}`, rnd, accountID)
}

func testAccConnectivityDirectoryServiceIPv4Updated(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_connectivity_directory_service" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s-updated"
  type       = "http"

  host {
    ipv4 = "192.168.1.200"
  }

  http_port = 8080
}`, rnd, accountID)
}

func testAccConnectivityDirectoryServiceWithPorts(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_connectivity_directory_service" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "http"

  host {
    ipv4 = "10.0.0.1"
  }

  http_port  = 80
  https_port = 443
}`, rnd, accountID)
}
