package custom_hostname_fallback_origin_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_custom_hostname_fallback_origin", &resource.Sweeper{
		Name: "cloudflare_custom_hostname_fallback_origin",
		F:    testSweepCloudflareCustomHostnameFallbackOrigin,
	})
}

func testSweepCloudflareCustomHostnameFallbackOrigin(r string) error {
	ctx := context.Background()
	// Custom Hostname Fallback Origin is a zone-level configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Custom Hostname Fallback Origin doesn't require sweeping (zone setting)")
	return nil
}

func testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, subdomain, domain string) string {
	return acctest.LoadTestCase("customhostnamefallbackorigin.tf", zoneID, rnd, subdomain, domain)
}

func testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, subdomain, domain string) string {
	return acctest.LoadTestCase("customhostnamefallbackorigin_updated.tf", zoneID, rnd, subdomain, domain)
}

func testAccCheckCloudflareCustomHostnameFallbackOriginDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname_fallback_origin" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		fallbackOrigin, err := client.CustomHostnames.FallbackOrigin.Get(
			context.Background(),
			custom_hostnames.FallbackOriginGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)

		// If error, resource is gone - that's expected
		if err != nil {
			continue
		}

		// If pending_deletion, that's acceptable (API will clean up)
		if fallbackOrigin.Status == "pending_deletion" {
			continue
		}

		return fmt.Errorf("Fallback Origin still exists with status: %s", fallbackOrigin.Status)
	}

	return nil
}

func TestAccCloudflareCustomHostnameFallbackOrigin_FullLifecycle(t *testing.T) {
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
			// Step 1: Create
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin.%s.%s", rnd, domain))),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			// Step 2: Drift check - same config, expect empty plan
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Update origin
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin-updated.%s.%s", rnd, domain))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin-updated.%s.%s", rnd, domain))),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			// Step 4: Drift check after update
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 5: Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
			},
		},
	})
}
