package zero_trust_device_managed_networks_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// Config generators for v4 provider (old resource name)

func deviceManagedNetworksConfigV4Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_managed_networks" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "tls"

  config {
    tls_sockaddr = "example.com:443"
    sha256       = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
}`, rnd, accountID)
}


func deviceManagedNetworksConfigV4CustomPort(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_managed_networks" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "tls"

  config {
    tls_sockaddr = "custom.example.com:8443"
    sha256       = "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
  }
}`, rnd, accountID)
}

func deviceManagedNetworksConfigV4IPv6(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_managed_networks" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "tls"

  config {
    tls_sockaddr = "[2001:db8::1]:443"
    sha256       = "fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"
  }
}`, rnd, accountID)
}

// TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_Basic tests basic migration with all fields
func TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_Basic(t *testing.T) {
	// Zero Trust resources require API Key + Email authentication
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_device_managed_networks." + rnd
	tmpDir := t.TempDir()

	v4Config := deviceManagedNetworksConfigV4Basic(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
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
				// Verify resource type changed (renamed from cloudflare_device_managed_networks)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_id"), knownvalue.NotNull()),
				// Verify required fields migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("tls")),
				// Verify config transformed from array to object (TypeList MaxItems:1 → SingleNestedAttribute)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("tls_sockaddr"), knownvalue.StringExact("example.com:443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sha256"), knownvalue.StringExact("b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c")),
			}),
		},
	})
}


// TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_CustomPort tests migration with non-standard port
func TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_CustomPort(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_device_managed_networks." + rnd
	tmpDir := t.TempDir()

	v4Config := deviceManagedNetworksConfigV4CustomPort(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify custom port preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("tls_sockaddr"), knownvalue.StringExact("custom.example.com:8443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sha256"), knownvalue.StringExact("1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")),
			}),
		},
	})
}

// TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_IPv6 tests migration with IPv6 address
func TestMigrateZeroTrustDeviceManagedNetworks_V4ToV5_IPv6(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_device_managed_networks." + rnd
	tmpDir := t.TempDir()

	v4Config := deviceManagedNetworksConfigV4IPv6(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify IPv6 address preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("tls_sockaddr"), knownvalue.StringExact("[2001:db8::1]:443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sha256"), knownvalue.StringExact("fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321")),
			}),
		},
	})
}
