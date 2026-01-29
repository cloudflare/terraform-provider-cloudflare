package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

const (
	latestV4Version  = "4.52.5"
	currentV5Version = "5.16.0"
)

// Migration Test Configuration
//
// Version constants are defined in internal/version.go:
// - latestV4Version: Last stable v4 release (4.52.5)
// - currentV5Version: Current v5 release (auto-updates with releases)
//
// Based on breaking changes analysis:
// - All breaking changes happened between 4.x and 5.0.0
// - No breaking changes between v5 releases (testing against latest v5)
// - Key changes: cloudflare_record → cloudflare_dns_record, data block → attribute

// Embed migration test configuration files
//
//go:embed testdata/v4_a_record.tf
var v4ARecordConfig string

//go:embed testdata/v5_a_record.tf
var v5ARecordConfig string

//go:embed testdata/v4_caa_record.tf
var v4CAARecordConfig string

//go:embed testdata/v5_caa_record.tf
var v5CAARecordConfig string

//go:embed testdata/v4_mx_record.tf
var v4MXRecordConfig string

//go:embed testdata/v5_mx_record.tf
var v5MXRecordConfig string

//go:embed testdata/v4_srv_record.tf
var v4SRVRecordConfig string

//go:embed testdata/v5_srv_record.tf
var v5SRVRecordConfig string

//go:embed testdata/v4_txt_record.tf
var v4TXTRecordConfig string

//go:embed testdata/v5_txt_record.tf
var v5TXTRecordConfig string

//go:embed testdata/v4_cname_record.tf
var v4CNAMERecordConfig string

//go:embed testdata/v5_cname_record.tf
var v5CNAMERecordConfig string

//go:embed testdata/v4_allow_overwrite.tf
var v4AllowOverwriteConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v4_aaaa_record.tf
var v4AAAARecordConfig string

//go:embed testdata/v5_aaaa_record.tf
var v5AAAARecordConfig string

//go:embed testdata/v4_ns_record.tf
var v4NSRecordConfig string

//go:embed testdata/v5_ns_record.tf
var v5NSRecordConfig string

//go:embed testdata/v4_tags.tf
var v4TagsConfig string

//go:embed testdata/v5_tags.tf
var v5TagsConfig string

//go:embed testdata/v4_ptr_record.tf
var v4PTRRecordConfig string

//go:embed testdata/v5_ptr_record.tf
var v5PTRRecordConfig string

//go:embed testdata/v5_issue6076_basic.tf
var v5Issue6076BasicConfig string

//go:embed testdata/v5_issue6076_updated.tf
var v5Issue6076UpdatedConfig string

// TestMigrateDNSRecordBasicA tests migration of a simple A record from v4 to v5
// Version constant latestV4Version is defined in internal/version.go
func TestMigrateDNSRecordBasicA(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4ARecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5ARecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-a-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Resource should be renamed to cloudflare_dns_record
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("A")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("52.152.96.252")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("ttl"), knownvalue.Float64Exact(1)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("proxied"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("tags"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("tf-applied")})),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordCAARecord tests migration of CAA record with data block conversion
// Using real example from oaistatic_com/dns.tf
func TestMigrateDNSRecordCAARecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4CAARecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5CAARecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-caa-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("CAA")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("ttl"), knownvalue.Float64Exact(1)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("proxied"), knownvalue.Bool(false)),
						// Data should be converted from block to attribute
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("flags"), knownvalue.Float64Exact(0)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("tag"), knownvalue.StringExact("issue")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("value"), knownvalue.StringExact("letsencrypt.org")),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordMXRecord tests migration of MX record with priority
// Using real example from operator_chatgpt_com/mailserver.tf
func TestMigrateDNSRecordMXRecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4MXRecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5MXRecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-mx-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("MX")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("mail.example.com")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("priority"), knownvalue.Float64Exact(10)),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordSRVRecord tests migration of SRV record with complex data
func TestMigrateDNSRecordSRVRecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4SRVRecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5SRVRecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("_sip._tcp.tf-test-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("SRV")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("priority"), knownvalue.Float64Exact(10)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("priority"), knownvalue.Float64Exact(10)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("weight"), knownvalue.Float64Exact(60)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("port"), knownvalue.Float64Exact(5060)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("data").AtMapKey("target"), knownvalue.StringExact("sipserver.example.com")),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordTXTRecord tests migration of TXT record
func TestMigrateDNSRecordTXTRecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4TXTRecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5TXTRecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-txt-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("TXT")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("v=spf1 -all")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("ttl"), knownvalue.Float64Exact(1)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("proxied"), knownvalue.Bool(false)),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordCNAMERecord tests migration of CNAME record
func TestMigrateDNSRecordCNAMERecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4CNAMERecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5CNAMERecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-cname-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("CNAME")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("abc-browser-external.foo.com")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("ttl"), knownvalue.Float64Exact(1)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("proxied"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("tags"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("tf-applied")})),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordWithAllowOverwrite tests migration with v4-only attribute allow_overwrite
func TestMigrateDNSRecordWithAllowOverwrite(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-overwrite-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// V4 config with allow_overwrite (should be removed in v5)
	v4Config := fmt.Sprintf(v4AllowOverwriteConfig, rnd, zoneID, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: latestV4Version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, latestV4Version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("A")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("192.0.2.2")),
				// allow_overwrite should not exist in v5 state
			}),
		},
	})
}

// TestMigrateDNSRecordMultipleRecords tests migration of multiple records showing real usage patterns
func TestMigrateDNSRecordMultipleRecords(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(v4MultipleConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: latestV4Version,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state for all records
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, latestV4Version, "v4", "v5", []statecheck.StateCheck{
				// A record checks
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_a", tfjsonpath.New("content"), knownvalue.StringExact("52.152.96.252")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_a", tfjsonpath.New("tags"), knownvalue.ListSizeExact(2)),

				// CNAME record checks (value should be migrated to content)
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_cname", tfjsonpath.New("content"), knownvalue.StringExact(fmt.Sprintf("api-%s.%s", rnd, "terraform.cfapi.net"))),

				// CAA record checks
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_caa", tfjsonpath.New("data").AtMapKey("flags"), knownvalue.Float64Exact(0)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_caa", tfjsonpath.New("data").AtMapKey("tag"), knownvalue.StringExact("issue")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_caa", tfjsonpath.New("data").AtMapKey("value"), knownvalue.StringExact("pki.goog")),

				// TXT record checks
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_txt", tfjsonpath.New("content"), knownvalue.StringExact("v=DMARC1; p=reject; sp=reject;")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd+"_txt", tfjsonpath.New("ttl"), knownvalue.Float64Exact(300)),
			}),
		},
	})
}

// TestMigrateDNSRecordAAAARecord tests migration of AAAA (IPv6) record
func TestMigrateDNSRecordAAAARecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4AAAARecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5AAAARecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-aaaa-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("AAAA")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("2001:db8::1")),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordNSRecord tests migration of NS record
func TestMigrateDNSRecordNSRecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4NSRecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5NSRecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-ns-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("NS")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("ns1.example.com")),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordWithTags tests migration of record with tags
func TestMigrateDNSRecordWithTags(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4TagsConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5TagsConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-tags-%s", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("A")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("192.0.2.3")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("tags"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("env:test"),
							knownvalue.StringExact("managed:terraform"),
						})),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecordPTRRecord tests migration of PTR (reverse DNS) record
func TestMigrateDNSRecordPTRRecord(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  latestV4Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v4PTRRecordConfig, rnd, zoneID, name) },
		},
		{
			name:     "from_v5",
			version:  currentV5Version,
			configFn: func(rnd, zoneID, name string) string { return fmt.Sprintf(v5PTRRecordConfig, rnd, zoneID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("1.2.0.192.in-addr.%s.arpa", rnd)
			tmpDir := t.TempDir()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("PTR")),
						statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("example.com")),
					}),
				},
			})
		})
	}
}

// TestMigrateDNSRecord_Issue6076 tests for GitHub issue #6076
func TestMigrateDNSRecord_Issue6076(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-6076-%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// Config that reproduces the issue
	config := fmt.Sprintf(v5Issue6076BasicConfig, rnd, zoneID, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.8.4 (last version before the rewrite)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.8.4",
					},
				},
				Config:             config,
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 2: Create with v5.9.0, expect apply error
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.9.0",
					},
				},
				Config:      config,
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Error: Provider produced inconsistent result after apply")),
			},
			{
				// Step 3: Apply with current provider
				// This should work without the "inconsistent result" error
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "zone_id", zoneID),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "name", name+"."+domain),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "type", "CNAME"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "content", "kay.ns.cloudflare.com"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "ttl", "1"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "proxied", "true"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "comment", "a comment"),
					resource.TestCheckResourceAttrSet("cloudflare_dns_record."+rnd, "modified_on"),
					resource.TestCheckResourceAttrSet("cloudflare_dns_record."+rnd, "created_on"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				// Step 4: Apply the same config again to verify no drift
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				// Step 5: Update comment and tags
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   fmt.Sprintf(v5Issue6076UpdatedConfig, rnd, zoneID, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "comment", "updated comment for testing"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "tags.#", "2"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "tags.0", "migration"),
					resource.TestCheckResourceAttr("cloudflare_dns_record."+rnd, "tags.1", "test"),
					resource.TestCheckResourceAttrSet("cloudflare_dns_record."+rnd, "comment_modified_on"),
					resource.TestCheckResourceAttrSet("cloudflare_dns_record."+rnd, "tags_modified_on"),
				),
			},
			{
				// Step 6: Apply the updated config again to verify no drift
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   fmt.Sprintf(v5Issue6076UpdatedConfig, rnd, zoneID, name),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // This should pass with our fix
					},
				},
			},
		},
	})
}
