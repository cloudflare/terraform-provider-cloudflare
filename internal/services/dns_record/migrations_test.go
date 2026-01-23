package dns_record_test

import (
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

// TestMigrateDNSRecordBasicA tests migration of a simple A record from v4 to v5
func TestMigrateDNSRecordBasicA(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-a-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = true
  tags    = ["tf-applied"]
  ttl     = 1
  type    = "A"
  content = "52.152.96.252"
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
}

// TestMigrateDNSRecordCAARecord tests migration of CAA record with data block conversion
// Using real example from oaistatic_com/dns.tf
func TestMigrateDNSRecordCAARecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-caa-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// V4 config with data as a block - based on oaistatic_com/dns.tf
	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = false
  ttl     = 1
  type    = "CAA"
  
  data {
    flags = 0
    tag   = "issue"
    value = "letsencrypt.org"
  }
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
}

// TestMigrateDNSRecordMXRecord tests migration of MX record with priority
// Using real example from operator_chatgpt_com/mailserver.tf
func TestMigrateDNSRecordMXRecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-mx-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id  = "%[2]s"
  name     = "%[3]s"
  type     = "MX"
  content  = "mail.example.com"
  priority = 10
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("MX")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("mail.example.com")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("priority"), knownvalue.Float64Exact(10)),
			}),
		},
	})
}

// TestMigrateDNSRecordSRVRecord tests migration of SRV record with complex data
func TestMigrateDNSRecordSRVRecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("_sip._tcp.tf-test-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "SRV"
  ttl     = 3600
  
  data {
    priority = 10
    weight   = 60
    port     = 5060
    target   = "sipserver.example.com"
  }
}`, rnd, zoneID, name)

	// V5 config needs priority at root level
	//	v5Config := fmt.Sprintf(`
	//resource "cloudflare_dns_record" "%[1]s" {
	//  zone_id = "%[2]s"
	//  name    = "%[3]s"
	//  type    = "SRV"
	//  ttl     = 3600
	//  priority = 10
	//
	//  data = {
	//    priority = 10
	//    weight   = 60
	//    port     = 5060
	//    target   = "sipserver.example.com"
	//  }
	//}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
}

// TestMigrateDNSRecordTXTRecord tests migration of TXT record
func TestMigrateDNSRecordTXTRecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-txt-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// V4 config - based on oaiusercontent_com/dns.tf
	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = false
  ttl     = 1
  type    = "TXT"
  value   = "v=spf1 -all"
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("TXT")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("v=spf1 -all")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("ttl"), knownvalue.Float64Exact(1)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("proxied"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateDNSRecordCNAMERecord tests migration of CNAME record
func TestMigrateDNSRecordCNAMERecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-cname-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = true
  tags    = ["tf-applied"]
  ttl     = 1
  type    = "CNAME"
  value   = "abc-browser-external.foo.com"
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
}

// TestMigrateDNSRecordWithAllowOverwrite tests migration with v4-only attribute allow_overwrite
func TestMigrateDNSRecordWithAllowOverwrite(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-overwrite-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// V4 config with allow_overwrite (should be removed in v5)
	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id         = "%[2]s"
  name            = "%[3]s"
  type            = "A"
  value           = "192.0.2.2"
  ttl             = 3600
  allow_overwrite = true
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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

	v4Config := fmt.Sprintf(`
# A record with content field (some configs use content instead of value)
resource "cloudflare_record" "%[1]s_a" {
  zone_id = "%[2]s"
  name    = "api-%[1]s"
  proxied = true
  tags    = ["tf-applied", "production"]
  ttl     = 1
  type    = "A"
  content = "52.152.96.252"
}

# CNAME with value field
resource "cloudflare_record" "%[1]s_cname" {
  zone_id = "%[2]s"
  name    = "www-%[1]s"
  proxied = true
  ttl     = 1
  type    = "CNAME"
  value   = "api-%[1]s.terraform.cfapi.net"
}

# CAA record with data block
resource "cloudflare_record" "%[1]s_caa" {
  zone_id = "%[2]s"
  name    = "caa-%[1]s"
  proxied = false
  ttl     = 1
  type    = "CAA"
  
  data {
    flags = 0
    tag   = "issue"
    value = "pki.goog"
  }
}

# TXT record
resource "cloudflare_record" "%[1]s_txt" {
  zone_id = "%[2]s"
  name    = "_dmarc-%[1]s"
  proxied = false
  ttl     = 300
  type    = "TXT"
  content = "v=DMARC1; p=reject; sp=reject;"
}`, rnd, zoneID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state for all records
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-aaaa-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "AAAA"
  value   = "2001:db8::1"
  ttl     = 3600
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("AAAA")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("2001:db8::1")),
			}),
		},
	})
}

// TestMigrateDNSRecordNSRecord tests migration of NS record
func TestMigrateDNSRecordNSRecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-ns-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "NS"
  value   = "ns1.example.com"
  ttl     = 3600
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("NS")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("ns1.example.com")),
			}),
		},
	})
}

// TestMigrateDNSRecordWithTags tests migration of record with tags
func TestMigrateDNSRecordWithTags(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-tags-%s", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "A"
  value   = "192.0.2.3"
  ttl     = 3600
  tags    = ["env:test", "managed:terraform"]
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
}

// TestMigrateDNSRecordPTRRecord tests migration of PTR (reverse DNS) record
func TestMigrateDNSRecordPTRRecord(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("1.2.0.192.in-addr.%s.arpa", rnd)
	tmpDir := t.TempDir()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	v4Config := fmt.Sprintf(`
resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "PTR"
  value   = "example.com"
  ttl     = 3600
}`, rnd, zoneID, name)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s.%s", name, domain))),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("PTR")),
				statecheck.ExpectKnownValue("cloudflare_dns_record."+rnd, tfjsonpath.New("content"), knownvalue.StringExact("example.com")),
			}),
		},
	})
}

// TestMigrateDNSRecord_Issue6076 tests for GitHub issue #6076
func TestMigrateDNSRecord_Issue6076(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-6076-%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	// Config that reproduces the issue
	config := fmt.Sprintf(`
resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  comment = "a comment"
  name    = "%[3]s"
  type    = "CNAME"
  content = "kay.ns.cloudflare.com"
  ttl     = 1
  proxied = true
}`, rnd, zoneID, name)

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
				Config: fmt.Sprintf(`
resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  comment = "updated comment for testing"
  name    = "%[3]s"
  type    = "CNAME"
  content = "kay.ns.cloudflare.com"
  ttl     = 1
  proxied = true
  tags    = ["test", "migration"]
}`, rnd, zoneID, name),
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
				Config: fmt.Sprintf(`
resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  comment = "updated comment for testing"
  name    = "%[3]s"
  type    = "CNAME"
  content = "kay.ns.cloudflare.com"
  ttl     = 1
  proxied = true
  tags    = ["test", "migration"]
}`, rnd, zoneID, name),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // This should pass with our fix
					},
				},
			},
		},
	})
}
