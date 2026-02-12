package v500_test

import (
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_origin_port_range.tf
var v4OriginPortRangeConfig string

//go:embed testdata/v4_edge_ips.tf
var v4EdgeIPsConfig string

//go:embed testdata/v4_origin_direct.tf
var v4OriginDirectConfig string

//go:embed testdata/v4_complex.tf
var v4ComplexConfig string

// TestMigrateSpectrumApplication_Basic tests v4→v5 migration with dns block and origin_direct
func TestMigrateSpectrumApplication_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_spectrum_application." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4BasicConfig, rnd, zoneID, rnd, domain)
	version := acctest.GetLastV4Version()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/22")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

// TestMigrateSpectrumApplication_OriginPortRange tests origin_port_range block → origin_port string conversion
func TestMigrateSpectrumApplication_OriginPortRange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_spectrum_application." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4OriginPortRangeConfig, rnd, zoneID, rnd, domain)
	version := acctest.GetLastV4Version()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_port"), knownvalue.StringExact("3306-3310")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

// TestMigrateSpectrumApplication_EdgeIPs tests edge_ips block handling
func TestMigrateSpectrumApplication_EdgeIPs(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_spectrum_application." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4EdgeIPsConfig, rnd, zoneID, rnd, domain)
	version := acctest.GetLastV4Version()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("type"), knownvalue.StringExact("dynamic")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("edge_ips").AtMapKey("connectivity"), knownvalue.StringExact("ipv4")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.1:23")),
			})...),
	})
}

// TestMigrateSpectrumApplication_OriginDirect tests origin_port as integer (DynamicAttribute number)
func TestMigrateSpectrumApplication_OriginDirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_spectrum_application." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4OriginDirectConfig, rnd, zoneID, rnd, domain)
	version := acctest.GetLastV4Version()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_direct").AtSliceIndex(0), knownvalue.StringExact("tcp://128.66.0.2:3306")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_port"), knownvalue.NumberExact(big.NewFloat(3306))),
			})...),
	})
}

// TestMigrateSpectrumApplication_Complex tests v4→v5 migration with all optional fields
func TestMigrateSpectrumApplication_Complex(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_spectrum_application." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4ComplexConfig, rnd, zoneID, rnd, domain)
	version := acctest.GetLastV4Version()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("protocol"), knownvalue.StringExact("tcp/443")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("type"), knownvalue.StringExact("CNAME")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dns").AtMapKey("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
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