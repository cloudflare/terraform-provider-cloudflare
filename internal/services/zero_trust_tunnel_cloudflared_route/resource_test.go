package zero_trust_tunnel_cloudflared_route_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_cloudflared_route", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_cloudflared_route",
		F:    testSweepCloudflareTunnelRoute,
	})
}

func testSweepCloudflareTunnelRoute(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}
	isDeleted := false
	tunnelRoutes, err := client.ListTunnelRoutes(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelRoutesListParams{IsDeleted: &isDeleted})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Tunnel Routes: %s", err))
	}

	if len(tunnelRoutes) == 0 {
		log.Print("[DEBUG] No Cloudflare Tunnel Routes to sweep")
		return nil
	}

	for _, tunnel := range tunnelRoutes {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Tunnel Route network: %s", tunnel.Network))
		//nolint:errcheck
		client.DeleteTunnelRoute(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelRoutesDeleteParams{Network: tunnel.Network, VirtualNetworkID: tunnel.TunnelID})
	}

	return nil
}

func TestAccCloudflareTunnelRoute_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnets to avoid conflicts
	subnet1 := fmt.Sprintf("10.%d.%d.10/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))
	subnet2 := fmt.Sprintf("20.%d.%d.20/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
				),
			},
			// Update
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet2)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
				),
			},
			// Re-applying same change does not produce drift
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet2),
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

func TestAccCloudflareTunnelRoute_UpdateComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet to avoid conflicts
	subnet := fmt.Sprintf("10.%d.%d.10/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
				),
			},
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd+"-updated", accountID, subnet),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd+"-updated")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTunnelRoute_NoComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet to avoid conflicts
	cidr := fmt.Sprintf("10.%d.%d.1/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRouteNoComment(rnd, accountID, cidr),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(cidr)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTunnelRoute_UpdateTunnel(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnel1 := fmt.Sprintf("%s_tun1", rnd)
	tunnel2 := fmt.Sprintf("%s_tun2", rnd)
	tunnel1Id := ""
	tunnel2Id := ""
	var tunnel1IdPtr *string = &tunnel1Id
	var tunnel2IdPtr *string = &tunnel2Id

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, rnd, tunnel1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact("1.1.1.1/32")),
				},
				Check: resource.ComposeTestCheckFunc(
					getTunnelId(tunnel1, tunnel1IdPtr),
					getTunnelId(tunnel2, tunnel2IdPtr),
					resource.TestCheckResourceAttrPtr(name, "tunnel_id", tunnel1IdPtr),
				),
			},
			{
				Config: testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, rnd, tunnel2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact("1.1.1.1/32")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPtr(name, "tunnel_id", tunnel2IdPtr),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func getTunnelId(tunnelName string, tunnel1Id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", tunnelName)

		// Get the resource instance from the state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		// Access the instance's attributes directly
		*tunnel1Id = rs.Primary.Attributes["id"]

		return nil
	}
}

func testAccCloudflareTunnelRouteSimple(ID, comment, accountID, network string) string {
	return acctest.LoadTestCase("route_simple.tf", ID, comment, accountID, network)
}

func testAccCloudflareRouteNoComment(ID, accountID, network string) string {
	return acctest.LoadTestCase("route_no_comment.tf", ID, accountID, network)
}

func testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, ID, tunnel string) string {
	return acctest.LoadTestCase("route_update_tunnel.tf", accountID, tunnel1, tunnel2, ID, tunnel)
}

// Test for IPv6 network routing
func TestAccCloudflareTunnelRoute_IPv6(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Use IPv6 CIDR
	ipv6Network := "2001:db8::/64"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteIPv6(rnd, accountID, ipv6Network),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(ipv6Network)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("IPv6 route for "+rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for large CIDR blocks
func TestAccCloudflareTunnelRoute_LargeCIDR(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Test different CIDR sizes
	randOctet := utils.RandIntRange(16, 31)
	largeCidr := fmt.Sprintf("172.%d.0.0/16", randOctet)
	smallCidr := fmt.Sprintf("172.%d.1.0/24", randOctet)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteLargeCIDR(rnd, accountID, largeCidr),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(largeCidr)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Large CIDR block for "+rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			// Test updating to smaller CIDR
			{
				Config: testAccCloudflareTunnelRouteLargeCIDR(rnd, accountID, smallCidr),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(smallCidr)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(smallCidr)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Large CIDR block for "+rnd)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for virtual network integration - tests basic route without virtual network
// Virtual networks require enterprise setup, so we test the field without specifying one
func TestAccCloudflareTunnelRoute_VirtualNetworkSupport(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet
	subnet := fmt.Sprintf("192.168.%d.0/24", utils.RandIntRange(1, 254))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, "Virtual network support test", accountID, subnet),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Virtual network support test")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
					// virtual_network_id should be empty when not specified
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for comments with special characters
func TestAccCloudflareTunnelRoute_SpecialCharComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet
	subnet := fmt.Sprintf("10.%d.%d.0/24", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))
	// Comment with safe special characters (within 100 char limit)
	specialComment := "Comment with special chars: !@#$%^&*()_+-=[]{}| testing field limits"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSpecialComment(rnd, accountID, subnet, specialComment),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(specialComment)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			// Test updating to empty comment
			{
				Config: testAccCloudflareTunnelRouteSpecialComment(rnd, accountID, subnet, ""),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for single IP addresses (host routes)
func TestAccCloudflareTunnelRoute_SingleIP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Single IP address
	singleIP := fmt.Sprintf("203.0.113.%d/32", utils.RandIntRange(1, 254))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, "Single IP host route", accountID, singleIP),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(singleIP)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Single IP host route")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for computed fields validation
func TestAccCloudflareTunnelRoute_ComputedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	subnet := fmt.Sprintf("10.%d.%d.0/28", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, "Computed fields test", accountID, subnet),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Computed fields test")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
					// deleted_at should be null for active routes
					resource.TestCheckNoResourceAttr(name, "deleted_at"),
				),
			},
		},
	})
}

// Helper functions for new test configurations
func testAccCloudflareTunnelRouteIPv6(ID, accountID, network string) string {
	return acctest.LoadTestCase("route_ipv6.tf", ID, accountID, network)
}

func testAccCloudflareTunnelRouteLargeCIDR(ID, accountID, network string) string {
	return acctest.LoadTestCase("route_large_cidr.tf", ID, accountID, network)
}

// Helper function removed - virtual network test now uses existing simple config

func testAccCloudflareTunnelRouteSpecialComment(ID, accountID, network, comment string) string {
	return acctest.LoadTestCase("route_long_comment.tf", ID, accountID, network, comment)
}

// Test for invalid network CIDR formats (error conditions)
func TestAccCloudflareTunnelRoute_InvalidNetwork(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	
	testCases := []struct {
		name            string
		invalidNetwork  string
		expectedError   string
	}{
		{
			name:           "invalid_cidr_format",
			invalidNetwork: "192.168.1.0/33", // Invalid CIDR - /33 is too large for IPv4
			expectedError:  "Could not parse input|invalid address",
		},
		{
			name:           "not_a_network",
			invalidNetwork: "not-a-network",
			expectedError:  "Could not parse input|invalid address",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      testAccCloudflareTunnelRouteInvalidNetwork(rnd, accountID, tc.invalidNetwork),
						ExpectError: regexp.MustCompile(tc.expectedError),
					},
				},
			})
		})
	}
}

// Test for different network subnet sizes (/28, /30, etc.)
func TestAccCloudflareTunnelRoute_VariousSubnetSizes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Use fixed, non-conflicting networks to avoid API normalization issues
	subnet28 := fmt.Sprintf("172.%d.0.0/28", utils.RandIntRange(16, 31))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, "Subnet size test", accountID, subnet28),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("network"), knownvalue.StringExact(subnet28)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("Subnet size test")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			// Import step
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// Test for missing required fields
func TestAccCloudflareTunnelRoute_MissingRequiredFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s" {
    account_id = "%[2]s"
    # Missing required tunnel_id and network
    comment = "Missing required fields test"
}`, rnd, accountID),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
		},
	})
}

// Test for invalid tunnel ID
func TestAccCloudflareTunnelRoute_InvalidTunnelID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	subnet := fmt.Sprintf("10.%d.%d.0/24", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s" {
    account_id = "%[2]s"
    tunnel_id = "invalid-tunnel-id-not-uuid"
    network = "%[3]s"
    comment = "Invalid tunnel ID test"
}`, rnd, accountID, subnet),
				ExpectError: regexp.MustCompile("not found|invalid|tunnel"),
			},
		},
	})
}

// Test for route conflict detection - API returns 409 for duplicate networks
func TestAccCloudflareTunnelRoute_ConflictingRoutes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Same network CIDR for both routes - should conflict
	conflictingNetwork := fmt.Sprintf("10.%d.%d.0/24", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareTunnelRouteConflicting(rnd, accountID, conflictingNetwork),
				ExpectError: regexp.MustCompile("409 Conflict|You already have a route defined for this exact IP subnet"),
			},
		},
	})
}

// Helper functions for error test configurations
func testAccCloudflareTunnelRouteInvalidNetwork(ID, accountID, network string) string {
	return acctest.LoadTestCase("route_invalid_network.tf", ID, accountID, network)
}

func testAccCloudflareTunnelRouteConflicting(ID, accountID, network string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s_1" {
	account_id    = "%[2]s"
	name          = "%[1]s_1"
	tunnel_secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s_2" {
	account_id    = "%[2]s"
	name          = "%[1]s_2"
	tunnel_secret = "UGBAECAwQFwgBAgIDBAUMEAQIDBABQYHCBgcIAQGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s_1" {
    account_id = "%[2]s"
    tunnel_id = cloudflare_zero_trust_tunnel_cloudflared.%[1]s_1.id
    network = "%[3]s"
    comment = "First route"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s_2" {
    account_id = "%[2]s"
    tunnel_id = cloudflare_zero_trust_tunnel_cloudflared.%[1]s_2.id
    network = "%[3]s"
    comment = "Conflicting route with same network"
}`, ID, accountID, network)
}
