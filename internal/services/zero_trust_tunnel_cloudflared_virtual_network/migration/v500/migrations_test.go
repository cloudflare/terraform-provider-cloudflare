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

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_new_name.tf
var v4NewNameConfig string

//go:embed testdata/v5_new_name.tf
var v5NewNameConfig string

//go:embed testdata/v4_complete.tf
var v4CompleteConfig string

//go:embed testdata/v5_complete.tf
var v5CompleteConfig string

// TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_Basic tests migration of a virtual network
// using the deprecated cloudflare_tunnel_virtual_network resource type with minimal config.
func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_Basic(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-vnet-"+rnd)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_NewName tests migration using the preferred
// v4 name cloudflare_zero_trust_tunnel_virtual_network with minimal config.
func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_NewName(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4NewNameConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5NewNameConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-vnet-new-"+rnd)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_Complete tests migration with all optional
// fields set — verifying comment and is_default_network values are preserved.
func TestMigrateZeroTrustTunnelCloudflaredVirtualNetwork_Complete(t *testing.T) {
	// Zero Trust resources don't support API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4CompleteConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5CompleteConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-vnet-complete-"+rnd)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("comment"), knownvalue.StringExact("Migration test complete")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
					}),
				},
			})
		})
	}
}
