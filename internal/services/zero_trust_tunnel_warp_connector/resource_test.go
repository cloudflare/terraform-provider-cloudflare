package zero_trust_tunnel_warp_connector_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare6 "github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_warp_connector", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_warp_connector",
		F:    testSweepCloudflareZeroTrustTunnelWARPConnector,
	})
}

func testSweepCloudflareZeroTrustTunnelWARPConnector(region string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping WARP Connector tunnels sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	page, err := client.ZeroTrust.Tunnels.WARPConnector.List(
		ctx,
		zero_trust.TunnelWARPConnectorListParams{
			AccountID: cloudflare6.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list WARP Connector tunnels: %s", err))
		return fmt.Errorf("error listing WARP Connector tunnels for sweep: %w", err)
	}

	tunnelCount := 0
	for page != nil && len(page.Result) > 0 {
		for _, tunnel := range page.Result {
			if tunnel.ID == "" {
				tflog.Debug(ctx, fmt.Sprintf("Skipping WARP Connector tunnel with empty ID: %s", tunnel.Name))
				continue
			}

			if !tunnel.DeletedAt.IsZero() {
				tflog.Debug(ctx, fmt.Sprintf("Skipping already deleted WARP Connector tunnel: %s (%s)", tunnel.Name, tunnel.ID))
				continue
			}

			if !utils.ShouldSweepResource(tunnel.Name) {
				continue
			}

			tunnelCount++
			tflog.Info(ctx, fmt.Sprintf("Deleting WARP Connector tunnel: %s (%s) (account: %s)", tunnel.Name, tunnel.ID, accountID))

			_, err := client.ZeroTrust.Tunnels.WARPConnector.Delete(
				ctx,
				tunnel.ID,
				zero_trust.TunnelWARPConnectorDeleteParams{
					AccountID: cloudflare6.F(accountID),
				},
				option.WithQuery("cascade", "true"),
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete WARP Connector tunnel %s (%s): %s", tunnel.Name, tunnel.ID, err))
			}
		}

		nextPage, err := page.GetNextPage()
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error getting next page: %s", err))
			break
		}
		page = nextPage
	}

	tflog.Info(ctx, fmt.Sprintf("Successfully swept %d WARP Connector tunnel(s)", tunnelCount))
	return nil
}

func TestAccWARPConnectorCreateBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func TestAccWARPConnectorUpdateName(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	name1 := fmt.Sprintf("%s_1", rnd)
	name2 := fmt.Sprintf("%s_2", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func testAccCheckWARPConnectorBasic(accID, name string) string {
	return acctest.LoadTestCase("warp_connector_basic.tf", accID, name)
}

func testAccCheckWARPConnectorUpdateName(accID, resourceName, name string) string {
	return acctest.LoadTestCase("warp_connector_update_name.tf", accID, resourceName, name)
}

func testAccCheckWARPConnectorHA(accID, name string) string {
	return acctest.LoadTestCase("warp_connector_ha.tf", accID, name)
}

func testAccCheckWARPConnectorHAUpdateName(accID, resourceName, name string) string {
	return acctest.LoadTestCase("warp_connector_ha_update_name.tf", accID, resourceName, name)
}

// importStateIDFunc returns an ImportStateIdFunc that formats the import ID as
// "<account_id>/<tunnel_id>", which is what the provider's ImportState expects.
func importStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["account_id"], rs.Primary.ID), nil
	}
}

// TestAccZeroTrustTunnelWARPConnector_Lifecycle exercises the full non-HA CRUD
// lifecycle plus Import in a single test:
//
//  1. Create — verify core attributes and ha defaults to false
//  2. Import — verify all attributes are restored from the API
//  3. Update name — verify the name change is applied in-place (no replacement)
//  4. PlanOnly — verify the updated state is a no-op (no perpetual drift)
func TestAccZeroTrustTunnelWARPConnector_Lifecycle(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)
	updatedName := rnd + "_updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create — basic attributes and ha = false (default).
				Config: testAccCheckWARPConnectorBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "account_tag"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
					resource.TestCheckResourceAttr(resourceName, "ha", "false"),
				),
			},
			{
				// Step 2: Import — all attributes must round-trip through the API.
				// tunnel_secret is write-only (no_refresh) so it cannot be verified.
				// metadata is excluded because the Create response returns {} while
				// the GET response may differ slightly; ha correctness is verified
				// directly by ImportStateVerify on the ha attribute.
				ResourceName:            resourceName,
				ImportStateIdFunc:       importStateIDFunc(resourceName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tunnel_secret", "metadata"},
			},
			{
				// Step 3: Update — rename the tunnel in-place; ha must remain false.
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "ha", "false"),
				),
			},
			{
				// Step 4: PlanOnly — identical config must produce a no-op plan.
				Config:   testAccCheckWARPConnectorUpdateName(accID, rnd, updatedName),
				PlanOnly: true,
			},
		},
	})
}

// TestAccZeroTrustTunnelWARPConnector_LifecycleHA exercises the full HA CRUD
// lifecycle plus Import in a single test:
//
//  1. Create with ha = true — verify ha is stored correctly in state
//  2. PlanOnly — regression check for APIX-1806 / GH-7260: a second plan must
//     not propose a replacement even though ha is not a top-level API response field
//  3. Update name — name is mutable even though ha is immutable; ha must be preserved
//  4. Import — ha must be restored from metadata.ha so that the imported state is usable
func TestAccZeroTrustTunnelWARPConnector_LifecycleHA(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)
	updatedName := rnd + "_updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with ha = true; state must reflect the value.
				Config: testAccCheckWARPConnectorHA(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "account_tag"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
					resource.TestCheckResourceAttr(resourceName, "ha", "true"),
				),
			},
			{
				// Step 2: PlanOnly — must be a no-op.
				// Before the APIX-1806 fix this step showed "+ ha = true # forces replacement".
				Config:   testAccCheckWARPConnectorHA(accID, rnd),
				PlanOnly: true,
			},
			{
				// Step 3: Update name — ha is create-time-only (immutable) but name
				// is mutable.  The rename must happen in-place; ha must stay true.
				Config: testAccCheckWARPConnectorHAUpdateName(accID, rnd, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "ha", "true"),
				),
			},
			{
				// Step 4: Import — ha must be restored from metadata.ha.
				// metadata is excluded: the Create response returns {} while the GET
				// response returns {"ha":true}; this is an API-side inconsistency.
				// tunnel_secret is write-only and cannot be round-tripped.
				ResourceName:            resourceName,
				ImportStateIdFunc:       importStateIDFunc(resourceName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tunnel_secret", "metadata"},
			},
		},
	})
}
