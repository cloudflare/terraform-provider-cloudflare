package zero_trust_network_hostname_route_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZeroTrustNetworkHostnameRoute_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_network_hostname_route." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	secret := generateRandomTunnelSecret(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustNetworkHostnameRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccZeroTrustNetworkHostnameRouteConfig(rnd, accountID, "original", secret),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostname"), knownvalue.StringExact("original.test.example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(fmt.Sprintf("Test hostname route for tf-acctest-%s", rnd))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccZeroTrustNetworkHostnameRouteConfig(rnd, accountID, "update", secret),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostname"), knownvalue.StringExact("update.test.example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(fmt.Sprintf("Test hostname route for tf-acctest-%s", rnd))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
				},
			},
			// Re-applying same change does not produce drift
			{
				Config: testAccZeroTrustNetworkHostnameRouteConfig(rnd, accountID, "update", secret),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccZeroTrustNetworkHostnameRouteConfig(rnd, accountID string, prefix string, secret string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, prefix, secret)
}

func testAccCloudflareZeroTrustNetworkHostnameRouteImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		accountID := rs.Primary.Attributes["account_id"]
		if accountID == "" {
			return "", fmt.Errorf("account_id not found")
		}

		return fmt.Sprintf("%s/%s", accountID, rs.Primary.ID), nil
	}
}

func testAccCheckCloudflareZeroTrustNetworkHostnameRouteDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_network_hostname_route" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		if accountID == "" {
			return fmt.Errorf("account_id not found")
		}
		_, err := client.ZeroTrust.Networks.HostnameRoutes.Get(ctx, rs.Primary.ID, zero_trust.NetworkHostnameRouteGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("zero trust network hostname route still exists")
		}
	}

	return nil
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_network_hostname_route", &resource.Sweeper{
		Name: "cloudflare_zero_trust_network_hostname_route",
		F:    testSweepCloudflareZeroTrustNetworkHostnameRoute,
	})
}

func testSweepCloudflareZeroTrustNetworkHostnameRoute(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// List all hostname routes
	resp, err := client.ZeroTrust.Networks.HostnameRoutes.List(ctx, zero_trust.NetworkHostnameRouteListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list zero trust network hostname routes: %w", err)
	}

	for _, route := range resp.Result {
		// Only delete test resources
		if route.Comment != "" && strings.Contains(route.Comment, "Test hostname route for tf-acctest-") {
			_, err := client.ZeroTrust.Networks.HostnameRoutes.Delete(ctx, route.ID, zero_trust.NetworkHostnameRouteDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				return fmt.Errorf("failed to delete zero trust network hostname route %s: %w", route.ID, err)
			}
		}
	}

	return nil
}
