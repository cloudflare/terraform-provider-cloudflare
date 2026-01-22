package zero_trust_tunnel_cloudflared_config_test

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

func TestMigrateZeroTrustTunnelCloudflaredConfig_Basic(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	// v4 config with basic setup
	// Note: v4 uses block syntax: config { } and ingress_rule { }
	// Note: Use config_src = "cloudflare" since we're managing config via Cloudflare API
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    ingress_rule {
      service = "http_status:404"
    }
  }
}`, rnd, accountID, tunnelSecret)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify config is an object (not array) after migration
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("http_status:404")),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredConfig_WithOriginRequest(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	// v4 config with origin_request settings (without duration fields)
	// Note: v4 uses block syntax: config { }, origin_request { }, ingress_rule { }
	// Note: Duration fields are tested separately in unit tests
	// Note: Use config_src = "cloudflare" since we're managing config via Cloudflare API
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-or-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    origin_request {
      no_happy_eyeballs      = false
      keep_alive_connections = 1024
      http_host_header       = "example.internal"
      http2_origin           = true
    }

    ingress_rule {
      hostname = "test.example.com"
      service  = "http://localhost:8080"
    }

    ingress_rule {
      service = "http_status:404"
    }
  }
}`, rnd, accountID, tunnelSecret)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify origin_request fields migrated correctly
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_happy_eyeballs"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("keep_alive_connections"), knownvalue.Int64Exact(1024)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http_host_header"), knownvalue.StringExact("example.internal")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http2_origin"), knownvalue.Bool(true)),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflaredConfig_DeprecatedName(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	// v4 config using DEPRECATED resource name: cloudflare_tunnel_config
	// Note: v4 uses block syntax: config { } and ingress_rule { }
	// Note: Use config_src = "cloudflare" since we're managing config via Cloudflare API
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-dep-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_tunnel_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    ingress_rule {
      service = "http_status:404"
    }
  }
}`, rnd, accountID, tunnelSecret)

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
			// Step 2: Run migration and verify state - should rename resource type
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// After migration, resource type should be cloudflare_zero_trust_tunnel_cloudflared_config
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			}),
		},
	})
}
