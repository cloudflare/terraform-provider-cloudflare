package zero_trust_tunnel_cloudflared_virtual_network_test

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

func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_V4ToV5_BasicOldName(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Test old v4 name: cloudflare_tunnel_virtual_network
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel_virtual_network" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-vnet-old-%[1]s"
}`, rnd, accountID)

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
				// Resource should be renamed to cloudflare_zero_trust_tunnel_cloudflared_virtual_network
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-vnet-old-%s", rnd))),
				// Defaults should be added
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_V4ToV5_BasicNewName(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Test new v4 name: cloudflare_zero_trust_tunnel_virtual_network
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_virtual_network" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-vnet-new-%[1]s"
}`, rnd, accountID)

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
				// Resource should be renamed to cloudflare_zero_trust_tunnel_cloudflared_virtual_network
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-vnet-new-%s", rnd))),
				// Defaults should be added
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_V4ToV5_Complete(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Test with all optional fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel_virtual_network" "%[1]s" {
  account_id         = "%[2]s"
  name               = "tf-acc-test-vnet-complete-%[1]s"
  comment            = "Migration test complete"
  is_default_network = false
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-vnet-complete-%s", rnd))),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("Migration test complete")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_V4ToV5_NonDefaultNetwork(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Test with is_default_network = false
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_virtual_network" "%[1]s" {
  account_id         = "%[2]s"
  name               = "tf-acc-test-vnet-nondefault-%[1]s"
  is_default_network = false
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-vnet-nondefault-%s", rnd))),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
				// comment should still get default empty string
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("")),
			}),
		},
	})
}
