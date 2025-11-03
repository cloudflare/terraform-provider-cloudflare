package zero_trust_tunnel_cloudflared_route_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMigrateZeroTrustTunnelCloudflaredRoute_V4ToV5_Basic(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelID := os.Getenv("CLOUDFLARE_TUNNEL_ID")
	if tunnelID == "" {
		t.Skip("Skipping test: CLOUDFLARE_TUNNEL_ID not set")
	}

	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel_route" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = "%[3]s"
  network    = "10.99.88.0/26"
  comment    = "Test tunnel route for migration"
}`, rnd, accountID, tunnelID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Resource should be renamed to cloudflare_zero_trust_tunnel_cloudflared_route
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("tunnel_id"), knownvalue.StringExact(tunnelID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("network"), knownvalue.StringExact("10.99.88.0/26")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("Test tunnel route for migration")),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredRoute_V4ToV5_Minimal(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelID := os.Getenv("CLOUDFLARE_TUNNEL_ID")
	if tunnelID == "" {
		t.Skip("Skipping test: CLOUDFLARE_TUNNEL_ID not set")
	}

	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Minimal config - only required fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel_route" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = "%[3]s"
  network    = "172.31.250.0/28"
}`, rnd, accountID, tunnelID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("tunnel_id"), knownvalue.StringExact(tunnelID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("network"), knownvalue.StringExact("172.31.250.0/28")),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredRoute_V4ToV5_IPv6(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelID := os.Getenv("CLOUDFLARE_TUNNEL_ID")
	if tunnelID == "" {
		t.Skip("Skipping test: CLOUDFLARE_TUNNEL_ID not set")
	}

	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel_route" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = "%[3]s"
  network    = "fd00:cafe:beef::/64"
  comment    = "IPv6 tunnel route"
}`, rnd, accountID, tunnelID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("tunnel_id"), knownvalue.StringExact(tunnelID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("network"), knownvalue.StringExact("fd00:cafe:beef::/64")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_route."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("IPv6 tunnel route")),
			}),
		},
	})
}
