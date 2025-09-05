package spectrum_application_test

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMigrateSpectrumApplication_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_spectrum_application." + rnd

	v4Config := acctest.LoadTestCase("spectrumapplicationmigrationbasic.tf", zoneID, os.Getenv("CLOUDFLARE_DOMAIN"), rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/22")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN")))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

func TestMigrateSpectrumApplication_OriginPortRange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_spectrum_application." + rnd

	v4Config := acctest.LoadTestCase("spectrumapplicationmigrationoriginportrange.tf", zoneID, os.Getenv("CLOUDFLARE_DOMAIN"), rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN")))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_port"), knownvalue.StringExact("3306-3310")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

func TestMigrateSpectrumApplication_EdgeIPs(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_spectrum_application." + rnd

	v4Config := acctest.LoadTestCase("spectrumapplicationmigrationedgeips.tf", zoneID, os.Getenv("CLOUDFLARE_DOMAIN"), rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN")))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("type"), knownvalue.StringExact("dynamic")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("connectivity"), knownvalue.StringExact("ipv4")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

func TestMigrateSpectrumApplication_OriginDirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_spectrum_application." + rnd

	v4Config := acctest.LoadTestCase("spectrumapplicationmigrationorigindirect.tf", zoneID, os.Getenv("CLOUDFLARE_DOMAIN"), rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN")))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.2:3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_port"), knownvalue.NumberExact(big.NewFloat(3306))),
			})...),
	})
}

func TestMigrateSpectrumApplication_Complex(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_spectrum_application." + rnd

	v4Config := acctest.LoadTestCase("spectrumapplicationmigrationcomplex.tf", zoneID, os.Getenv("CLOUDFLARE_DOMAIN"), rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN")))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("type"), knownvalue.StringExact("dynamic")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("connectivity"), knownvalue.StringExact("all")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.3:443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tls"), knownvalue.StringExact("flexible")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("argo_smart_routing"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("proxy_protocol"), knownvalue.StringExact("v1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ip_firewall"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("traffic_type"), knownvalue.StringExact("direct")),
			})...),
	})
}