package custom_hostname_fallback_origin_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareCustomHostnameFallbackOrigin(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the custom hostname
	// fallback endpoint does not yet support the API tokens for updates and it
	// results in state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomHostnameFallbackOriginDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "origin", fmt.Sprintf("fallback-origin.%s.%s", rnd, domain)),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, subdomain, domain string) string {
	return acctest.LoadTestCase("customhostnamefallbackorigin.tf", zoneID, rnd, subdomain, domain)
}

func TestAccCloudflareCustomHostnameFallbackOriginUpdate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the custom hostname
	// fallback endpoint does not yet support the API tokens for updates and it
	// results in state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	rndUpdate := rnd + "-updated"
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomHostnameFallbackOriginDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "origin", fmt.Sprintf("fallback-origin.%s.%s", rnd, domain)),
				),
			},
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rndUpdate, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "origin", fmt.Sprintf("fallback-origin.%s.%s", rndUpdate, domain)),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameFallbackOriginDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname_fallback_origin" {
			continue
		}

		fallbackOrigin, err := client.CustomHostnameFallbackOrigin(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey])

		// If the fallback origin is in the process of being deleted, that's fine to
		// say it's been deleted as the remote API will take care of it.
		if fallbackOrigin.Status != "pending_deletion" && err == nil {
			return fmt.Errorf("Fallback Origin still exists")
		}
	}

	return nil
}
