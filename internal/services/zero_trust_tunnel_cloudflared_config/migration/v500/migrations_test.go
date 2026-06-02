package v500_test

import (
	_ "embed"
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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_origin_request.tf
var v4WithOriginRequestConfig string

//go:embed testdata/v4_deprecated_name.tf
var v4DeprecatedNameConfig string

func TestMigrateZeroTrustTunnelCloudflaredConfig_Basic(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID, tunnelSecret)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		// Verify config is an object (not array) after migration
		statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("http_status:404")),
	}

	// Use MigrationV2TestStepWithStateNormalization because:
	// The v4 PF provider (version=1) stored origin_request as non-null empty ({}) in state when
	// the API returned it. The v5 Read also stores it as non-null when the API returns origin_request={}.
	// This causes a plan showing origin_request: {all-nil} → null, which is not a falsey-to-null change.
	// MigrationV2TestStepWithStateNormalization allows that change to be applied (step 1),
	// which removes origin_request from the API, and then verifies clean plan in step 3.
	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, migrationSteps...),
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

	// Duration fields are tested separately in unit tests.
	v4Config := fmt.Sprintf(v4WithOriginRequestConfig, rnd, accountID, tunnelSecret)

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

	v4Config := fmt.Sprintf(v4DeprecatedNameConfig, rnd, accountID, tunnelSecret)

	stateChecks := []statecheck.StateCheck{
		// After migration, resource type should be cloudflare_zero_trust_tunnel_cloudflared_config
		statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_config."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
	}

	// Use MigrationV2TestStepWithStateNormalization for the same reason as Basic:
	// the API returns origin_request={} even for minimal configs, causing non-falsey-to-null drift.
	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, migrationSteps...),
	})
}
