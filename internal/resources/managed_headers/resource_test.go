package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_managed_headers", &resource.Sweeper{
		Name: "cloudflare_managed_headers",
		F:    testSweepCloudflareManagedHeaders,
	})
}

func testSweepCloudflareManagedHeaders(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zone := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zone == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	managedHeaders, err := client.ListZoneManagedHeaders(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListManagedHeadersParams{
		Status: "enabled",
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Zone Managed Headers: %s", err))
	}

	requestHeaders := make([]cloudflare.ManagedHeader, 0, len(managedHeaders.ManagedRequestHeaders))
	for _, h := range managedHeaders.ManagedRequestHeaders {
		tflog.Info(ctx, fmt.Sprintf("Disabling Cloudflare Zone Managed Header ID: %s", h.ID))
		h.Enabled = false
		requestHeaders = append(requestHeaders, h)
	}
	responseHeaders := make([]cloudflare.ManagedHeader, 0, len(managedHeaders.ManagedResponseHeaders))
	for _, h := range managedHeaders.ManagedResponseHeaders {
		tflog.Info(ctx, fmt.Sprintf("Disabling Cloudflare Zone Managed Header ID: %s", h.ID))
		h.Enabled = false
		responseHeaders = append(responseHeaders, h)
	}

	_, err = client.UpdateZoneManagedHeaders(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateManagedHeadersParams{
		ManagedHeaders: cloudflare.ManagedHeaders{
			ManagedRequestHeaders:  requestHeaders,
			ManagedResponseHeaders: responseHeaders,
		},
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to disable Cloudflare Zone Managed Headers: %s", err))
	}

	return nil
}

func TestAccCloudflareManagedHeaders(t *testing.T) {
	t.Parallel()

	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_headers." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedHeaders(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.id", "add_true_client_ip_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.id", "add_visitor_location_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.id", "add_security_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareManagedHeadersRemovedHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.id", "add_true_client_ip_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.id", "add_security_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareManagedHeaders(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_managed_headers" "%[1]s" {
	zone_id  = "%[2]s"
	managed_request_headers {
		id = "add_true_client_ip_headers"
		enabled = true
	}

	managed_request_headers {
		id = "add_visitor_location_headers"
		enabled = true
	}

	managed_response_headers {
		id = "add_security_headers"
		enabled = true
	}
  }`, rnd, zoneID)
}

func testAccCheckCloudflareManagedHeadersRemovedHeader(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_managed_headers" "%[1]s" {
	zone_id  = "%[2]s"
	managed_request_headers {
		id = "add_true_client_ip_headers"
		enabled = true
	}

	managed_response_headers {
		id = "add_security_headers"
		enabled = true
	}
  }`, rnd, zoneID)
}
