package zero_trust_tunnel_cloudflared_test

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

func TestMigrateZeroTrustTunnelCloudflared_V4ToV5_Basic(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  secret     = base64encode("test-secret-that-is-at-least-32-bytes-long-for-testing")
  config_src = "local"
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
				// Resource should be renamed to cloudflare_zero_trust_tunnel_cloudflared
				// Field secret should be renamed to tunnel_secret
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", rnd))),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("config_src"), knownvalue.StringExact("local")),
				// tunnel_secret should exist (renamed from secret)
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("tunnel_secret"), knownvalue.NotNull()),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflared_V4ToV5_Minimal(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Minimal config - only required fields
	// Note: config_src is included to prevent v5 PATCH operation (API bug workaround)
	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-minimal-%[1]s"
  secret     = base64encode("minimal-tunnel-secret-that-is-at-least-32-bytes-long-testing")
  config_src = "local"
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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-minimal-%s", rnd))),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("config_src"), knownvalue.StringExact("local")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("tunnel_secret"), knownvalue.NotNull()),
			}),
		},
	})
}

func TestMigrateZeroTrustTunnelCloudflared_V4ToV5_CloudflareConfigSrc(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-cloudflare-config-%[1]s"
  secret     = base64encode("cloudflare-config-tunnel-secret-at-least-32-bytes-long-test")
  config_src = "cloudflare"
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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-cloudflare-config-%s", rnd))),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("config_src"), knownvalue.StringExact("cloudflare")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared."+rnd, tfjsonpath.New("tunnel_secret"), knownvalue.NotNull()),
			}),
		},
	})
}
