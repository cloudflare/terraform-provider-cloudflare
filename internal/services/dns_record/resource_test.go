package dns_record_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"log"

	cloudflare "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_dns_record", &resource.Sweeper{
		Name: "cloudflare_dns_record",
		F:    testSweepCloudflareRecord,
	})
}

func testSweepCloudflareRecord(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up DNS records
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping DNS records sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// List all DNS records using v6 SDK
	records, err := client.DNS.Records.List(ctx, dns.RecordListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare DNS records: %s", err))
		return err
	}

	recordList := records.Result
	if len(recordList) == 0 {
		tflog.Info(ctx, "No Cloudflare DNS records to sweep")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d DNS records to evaluate", len(recordList)))

	deletedCount := 0
	skippedCount := 0

	for _, record := range recordList {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(record.Name) {
			skippedCount++
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting DNS record ID: %s, Name: %s, Type: %s, Content: %s", record.ID, record.Name, record.Type, record.Content))
		_, err := client.DNS.Records.Delete(ctx, record.ID, dns.RecordDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete DNS record %s: %s", record.ID, err))
		} else {
			deletedCount++
		}
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted %d DNS records, skipped %d records", deletedCount, skippedCount))
	return nil
}

func TestAccCloudflareRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-basic", rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("tf-acctest-basic.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "tag2"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Apex(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigApex(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_LOC(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigLOC(zoneID, "tf-acctest-loc."+domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "content", "37 46 46.000 N 122 23 35.000 W 0.00 100.00 0.00 0.00"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_minutes", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_seconds", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_direction", "N"),
					resource.TestCheckResourceAttr(resourceName, "data.long_degrees", "122"),
					resource.TestCheckResourceAttr(resourceName, "data.long_minutes", "23"),
					resource.TestCheckResourceAttr(resourceName, "data.long_seconds", "35"),
					resource.TestCheckResourceAttr(resourceName, "data.long_direction", "W"),
					resource.TestCheckResourceAttr(resourceName, "data.altitude", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.size", "100"),
					resource.TestCheckResourceAttr(resourceName, "data.precision_horz", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.precision_vert", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SRV(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("_xmpp-client._tcp.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "content", "0 5222 talk.l.google.com"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.priority", "5"),
					resource.TestCheckResourceAttr(resourceName, "data.weight", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.port", "5222"),
					resource.TestCheckResourceAttr(resourceName, "data.target", "talk.l.google.com"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CAA(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.value", "letsencrypt.org"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 300),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.value", "letsencrypt.org"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Proxied(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigProxied(zoneID, domain, "tf-acctest-proxied", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "true"),
					resource.TestCheckResourceAttr(resourceName, "proxied", "true"),
					resource.TestCheckResourceAttr(resourceName, "type", "CNAME"),
					resource.TestCheckResourceAttr(resourceName, "content", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Updated(t *testing.T) {
	//t.Parallel()
	var record dns.RecordResponse
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	recordName := "tf-acctest-update"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNewValue(zoneID, recordName, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributesUpdated(&record),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_typeForceNewRecord(t *testing.T) {
	//t.Parallel()
	var afterCreate, afterUpdate dns.RecordResponse
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-type-force-new"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigChangeType(zoneID, recordName, zoneName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterUpdate),
					testAccCheckCloudflareRecordRecreated(&afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_TtlValidation(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-ttl-validation"
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRecordConfigTtlValidation(zoneID, recordName, zoneName, rnd),
				ExpectError: regexp.MustCompile("ttl must be set to 1 when `proxied` is true"),
			},
		},
	})
}

func TestAccCloudflareRecord_ExplicitProxiedFalse(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_dns_record." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "false", "300"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "false"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "true", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "true"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "false", "300"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "false"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_MXWithPriorityZero(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_dns_record." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "priority", "0"),
					resource.TestCheckResourceAttr(resourceName, "content", "mail.terraform.cfapi.net"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_HTTPS(t *testing.T) {
	//t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.priority", "1"),
					resource.TestCheckResourceAttr(name, "data.target", "."),
					resource.TestCheckResourceAttr(name, "data.value", `alpn="h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SVCB(t *testing.T) {
	//t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.priority", "2"),
					resource.TestCheckResourceAttr(name, "data.target", "foo."),
					resource.TestCheckResourceAttr(name, "data.value", `alpn="h3,h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_MXNull(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordNullMX(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "content", "."),
					resource.TestCheckResourceAttr(name, "priority", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_DNSKEY(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending automating setup from https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/.")

	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordDNSKEY(zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", domain),
					resource.TestCheckResourceAttr(name, "type", "DNSKEY"),
					resource.TestCheckResourceAttr(name, "data.flags", "257"),
					resource.TestCheckResourceAttr(name, "data.protocol", "13"),
					resource.TestCheckResourceAttr(name, "data.algorithm", "2"),
					resource.TestCheckResourceAttr(name, "data.public_key", "mdsswUyr3DPW132mOi8V9xESWE8jTo0dxCjjnopKl+GqJxpVXckHAeF+KkxLbxILfDLUT0rAK9iUzy1L53eKGQ=="),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_ClearTags(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMultipleTags(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "tags.#", "2"),
					resource.TestCheckResourceAttr(name, "comment", "this is a comment"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNoTags(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "tags.#", "0"),
					resource.TestCheckNoResourceAttr(name, "comment"),
				),
			},
		},
	})
}

func TestSuppressTrailingDots(t *testing.T) {
	t.Parallel()

	cases := []struct {
		old      string
		new      string
		expected bool
	}{
		{"", "", true},
		{"", "example.com", false},
		{"", "example.com.", false},
		{"", ".", false}, // single dot is used for Null MX record
		{"example.com", "example.com", true},
		{"example.com", "example.com.", true},
		{"example.com", "sub.example.com", false},
		{"sub.example.com", "sub.example.com.", true},
		{".", ".", true},
	}

	for _, c := range cases {
		got := suppressTrailingDots("", c.old, c.new, nil)
		assert.Equal(t, c.expected, got)
	}
}

// TestAccCloudflareRecord_TagsDrift tests for the issue reported in
// https://github.com/cloudflare/terraform-provider-cloudflare/issues/5517
// where DNS records show perpetual drift with tags and other computed fields
func TestAccCloudflareRecord_TagsDrift(t *testing.T) {
	// Don't run in parallel to avoid conflicts
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create resources with various tag configurations
			{
				Config: testAccCheckCloudflareRecordConfigTagsDrift(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Record with empty tags list
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "tags.#", "0"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "type", "A"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "content", "192.168.0.10"),

					// Record with explicit tags (tags are a set, so order may vary)
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_tags", rnd), "tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_tags", rnd), "tags.*", "test:tag1"),
					resource.TestCheckTypeSetElemAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_tags", rnd), "tags.*", "env:test"),

					// Record with settings
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_settings", rnd), "settings.flatten_cname", "false"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_settings", rnd), "tags.#", "0"),

					// Record without tags specified
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_no_tags", rnd), "type", "A"),
				),
			},
			// Step 2: Apply the same configuration again to check for drift
			// This should not show any changes if the provider handles defaults correctly
			{
				Config:             testAccCheckCloudflareRecordConfigTagsDrift(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // We expect NO changes on re-apply
			},
			// Step 3: Import test to verify state consistency
			{
				ResourceName:        fmt.Sprintf("cloudflare_dns_record.%s", rnd),
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				// Ignore computed fields that might differ after import
				ImportStateVerifyIgnore: []string{"comment_modified_on", "created_on", "modified_on", "tags_modified_on"},
			},
			{
				ResourceName:            fmt.Sprintf("cloudflare_dns_record.%s_with_tags", rnd),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"comment_modified_on", "created_on", "modified_on", "tags_modified_on"},
			},
			{
				ResourceName:            fmt.Sprintf("cloudflare_dns_record.%s_with_settings", rnd),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"comment_modified_on", "created_on", "modified_on", "tags_modified_on"},
			},
		},
	})
}

// TestAccCloudflareRecord_ComputedFieldsDrift specifically tests for computed fields
// causing perpetual drift as reported in issue #5517
func TestAccCloudflareRecord_ComputedFieldsDrift(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create minimal DNS record
			{
				Config: testAccCheckCloudflareRecordConfigComputedDrift(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Verify minimal record
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "type", "A"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "content", "192.168.0.20"),
					// Check that computed fields are set
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "created_on"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "modified_on"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "proxiable"),

					// Verify CNAME with settings
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_cname_settings", rnd), "type", "CNAME"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_cname_settings", rnd), "settings.flatten_cname", "false"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_cname_settings", rnd), "tags.#", "0"),

					// Verify record with comment
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_comment", rnd), "comment", "Test comment for drift"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudflare_dns_record.%s_with_comment", rnd), "comment_modified_on"),
				),
			},
			// Step 2: Re-apply to check for drift - this is the critical test
			// If there's a drift issue, this will fail with ExpectNonEmptyPlan: false
			{
				Config:             testAccCheckCloudflareRecordConfigComputedDrift(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Should be no changes on re-apply
			},
			// Step 3: Make a small change to verify updates work correctly
			{
				Config: testAccCheckCloudflareRecordConfigComputedDriftUpdated(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_minimal", rnd), "content", "192.168.0.25"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_with_comment", rnd), "comment", "Updated comment"),
				),
			},
			// Step 4: Re-apply updated config to ensure no drift after update
			{
				Config:             testAccCheckCloudflareRecordConfigComputedDriftUpdated(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Should be no changes on re-apply
			},
		},
	})
}

// TestAccCloudflareRecord_DriftIssue5517 specifically tests for the drift issues
// reported in https://github.com/cloudflare/terraform-provider-cloudflare/issues/5517
// This test attempts to reproduce the exact scenarios users reported
func TestAccCloudflareRecord_DriftIssue5517(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create records that users reported as problematic
			{
				Config: testAccCheckCloudflareRecordConfigDriftRepro(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Check proxied CNAME with settings
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_proxied_cname_with_settings", rnd), "type", "CNAME"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_proxied_cname_with_settings", rnd), "proxied", "true"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_proxied_cname_with_settings", rnd), "settings.flatten_cname", "false"),

					// Check CNAME with mixed case - checking if it preserves case
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_cname_mixed_case", rnd), "type", "CNAME"),
					// For now, just check that it's set, we'll see in step 2 if it causes drift
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudflare_dns_record.%s_cname_mixed_case", rnd), "content"),

					// Check empty tags
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_a_record_empty_tags", rnd), "tags.#", "0"),
				),
			},
			// Step 2: CRITICAL TEST - Re-apply to check for drift
			// This is where users are seeing the issue
			{
				Config:             testAccCheckCloudflareRecordConfigDriftRepro(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // We expect NO changes
			},
		},
	})
}

// TestAccCloudflareRecord_ModifiedOnDrift repros issue #6438.
// This ensures that records with data fields (CAA, LOC) and settings don't cause
// "Provider produced inconsistent result after apply" errors.
func TestAccCloudflareRecord_ModifiedOnDrift6438(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	caaResourceName := fmt.Sprintf("cloudflare_dns_record.%s_caa", rnd)
	locResourceName := fmt.Sprintf("cloudflare_dns_record.%s_loc", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create CAA and LOC records with data field
			{
				Config: testAccCheckCloudflareRecordConfigModifiedOnDriftCAA(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Check CAA record
					resource.TestCheckResourceAttr(caaResourceName, "type", "CAA"),
					resource.TestCheckResourceAttr(caaResourceName, "data.flags", "0"),
					resource.TestCheckResourceAttr(caaResourceName, "data.tag", "issue"),
					resource.TestCheckResourceAttr(caaResourceName, "data.value", "letsencrypt.org"),
					// Verify modified_on is set
					resource.TestCheckResourceAttrSet(caaResourceName, "modified_on"),
					resource.TestCheckResourceAttrSet(caaResourceName, "created_on"),
					// Check LOC record
					resource.TestCheckResourceAttr(locResourceName, "type", "LOC"),
					resource.TestCheckResourceAttrSet(locResourceName, "modified_on"),
					resource.TestCheckResourceAttrSet(locResourceName, "created_on"),
				),
				// Use ConfigStateChecks to validate state after apply
				ConfigStateChecks: []statecheck.StateCheck{
					// Ensure modified_on is properly set as a known value
					statecheck.ExpectKnownValue(caaResourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(caaResourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(locResourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(locResourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
				// Use ConfigPlanChecks to validate the plan
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Ensure modified_on starts as unknown for new resources
						plancheck.ExpectUnknownValue(caaResourceName, tfjsonpath.New("modified_on")),
						plancheck.ExpectUnknownValue(locResourceName, tfjsonpath.New("modified_on")),
					},
				},
			},
			// Step 2: Re-apply same config - should not detect changes (no drift)
			{
				Config:             testAccCheckCloudflareRecordConfigModifiedOnDriftCAA(zoneID, rnd, domain),
				ExpectNonEmptyPlan: false, // No changes expected
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Ensure no changes are planned
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Refresh and ensure no drift is detected
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Step 4: Apply multiple times to ensure stability
			{
				Config:             testAccCheckCloudflareRecordConfigModifiedOnDriftCAA(zoneID, rnd, domain),
				ExpectNonEmptyPlan: false, // Still no changes expected
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// After apply, all computed fields should remain stable
					statecheck.ExpectKnownValue(caaResourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(caaResourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(locResourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(locResourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// TestAccCloudflareRecord_SettingsDrift tests that records with settings field
// don't cause drift when the settings are effectively empty or unchanged
func TestAccCloudflareRecord_SettingsDrift(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create record with settings
			{
				Config: testAccCheckCloudflareRecordConfigSettingsDrift(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Check CNAME record with settings
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_a_with_settings", rnd), "type", "CNAME"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_a_with_settings", rnd), "settings.flatten_cname", "false"),
					// Check empty settings record
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_a_empty_settings", rnd), "type", "A"),
				),
			},
			// Step 2: Re-apply - should not detect changes
			{
				Config:             testAccCheckCloudflareRecordConfigSettingsDrift(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Remove settings from first record
			{
				Config: testAccCheckCloudflareRecordConfigSettingsDriftRemoved(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_a_with_settings", rnd), "type", "CNAME"),
				),
			},
			// Step 4: Re-apply - should not detect changes after settings removal
			{
				Config:             testAccCheckCloudflareRecordConfigSettingsDriftRemoved(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccCloudflareRecord_ComprehensiveDriftPrevention is the ultimate test to ensure
// the modified_on drift issue never occurs again. It tests all edge cases and validates
// pre-apply, post-apply, and refresh states.
func TestAccCloudflareRecord_ComprehensiveDriftPrevention(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	// Define resource names for all test cases
	basicA := fmt.Sprintf("cloudflare_dns_record.%s_basic_a", rnd)
	caaRecord := fmt.Sprintf("cloudflare_dns_record.%s_caa", rnd)
	cnameRecord := fmt.Sprintf("cloudflare_dns_record.%s_cname", rnd)
	recordWithTags := fmt.Sprintf("cloudflare_dns_record.%s_with_tags", rnd)
	recordWithSettings := fmt.Sprintf("cloudflare_dns_record.%s_with_settings", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create all types of records
			{
				Config: testAccCheckCloudflareRecordConfigComprehensiveDrift(zoneID, rnd, domain, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// All computed fields should be unknown before first apply
						plancheck.ExpectUnknownValue(basicA, tfjsonpath.New("modified_on")),
						plancheck.ExpectUnknownValue(basicA, tfjsonpath.New("created_on")),
						plancheck.ExpectUnknownValue(caaRecord, tfjsonpath.New("modified_on")),
						plancheck.ExpectUnknownValue(cnameRecord, tfjsonpath.New("modified_on")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// After apply, all computed fields should be known
					statecheck.ExpectKnownValue(basicA, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(basicA, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(caaRecord, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(cnameRecord, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					// Verify tags handling
					statecheck.ExpectKnownValue(recordWithTags, tfjsonpath.New("tags"), knownvalue.SetSizeExact(2)),
					// Verify settings handling
					statecheck.ExpectKnownValue(recordWithSettings, tfjsonpath.New("settings"), knownvalue.NotNull()),
				},
			},
			// Step 2: Immediate re-apply - critical for drift detection
			{
				Config:             testAccCheckCloudflareRecordConfigComprehensiveDrift(zoneID, rnd, domain, false),
				ExpectNonEmptyPlan: false,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Force refresh to ensure Read doesn't cause drift
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Step 4: Make actual changes and verify they're detected correctly
			{
				Config: testAccCheckCloudflareRecordConfigComprehensiveDrift(zoneID, rnd, domain, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Should detect the actual content changes
						plancheck.ExpectResourceAction(basicA, plancheck.ResourceActionUpdate),
						// modified_on should be unknown when there are real changes
						plancheck.ExpectUnknownValue(basicA, tfjsonpath.New("modified_on")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the change was applied
					statecheck.ExpectKnownValue(basicA, tfjsonpath.New("content"), knownvalue.StringExact("192.168.1.2")),
				},
			},
			// Step 5: Re-apply after changes - ensure no drift
			{
				Config:             testAccCheckCloudflareRecordConfigComprehensiveDrift(zoneID, rnd, domain, true),
				ExpectNonEmptyPlan: false,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 6: Multiple refreshes to ensure stability
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// Simple test to isolate the tags drift issue for records without explicit tags
func TestAccCloudflareRecord_SimpleDrift(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSimpleDrift(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "type", "A"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "content", "192.168.0.50"),
				),
			},
			// Second step: Apply same config again - should NOT show drift
			{
				Config:             testAccCheckCloudflareRecordConfigSimpleDrift(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // We expect NO changes
			},
		},
	})
}

// Test CNAME case normalization specifically
func TestAccCloudflareRecord_CNAMECase(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigCNAMECase(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s", rnd), "type", "CNAME"),
				),
			},
			// Second step: Apply same config again - should NOT show drift
			{
				Config:             testAccCheckCloudflareRecordConfigCNAMECase(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // We expect NO changes
			},
		},
	})
}

func TestAccCloudflareRecord_CommentModifiedOn(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigWithoutComment(zoneID, "tf-acctest-basic", rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("tf-acctest-basic.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "type", "A"),
					resource.TestCheckNoResourceAttr(resourceName, "comment_modified_on"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigCommentModified(zoneID, "tf-acctest-basic", rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("tf-acctest-basic.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "comment", "Test comment for drift"),
					resource.TestCheckResourceAttrSet(resourceName, "comment_modified_on"),
				),
			},
		},
	})
}

// Test FQDN normalization for DNS record names
func TestAccCloudflareRecord_FQDNNormalize(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigFQDNNormalize(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					// Check subdomain record
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_subdomain", rnd), "type", "A"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_subdomain", rnd), "content", "192.168.0.100"),

					// Check multi-level subdomain
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_subdomain_multi", rnd), "type", "A"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudflare_dns_record.%s_subdomain_multi", rnd), "content", "192.168.0.102"),
				),
			},
			// Second step: Apply same config again - should NOT show drift for name fields
			{
				Config:             testAccCheckCloudflareRecordConfigFQDNNormalize(zoneID, rnd, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // We expect NO changes
			},
		},
	})
}

func testAccCheckCloudflareRecordConfigCNAMECase(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_cname_case.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigTagsDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_tags_drift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigSimpleDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_simple_drift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigFQDNNormalize(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_fqdn_normalize.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigComputedDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_computed_drift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigDriftRepro(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_drift_repo.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigComputedDriftUpdated(zoneID, rnd, domain string) string {
	// Create an inline config for the update since it's a simple change
	return fmt.Sprintf(`
resource "cloudflare_dns_record" "%[1]s_minimal" {
  zone_id = "%[2]s"
  name    = "tf-acctest-minimal.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.25"  # Changed from .20
  ttl     = 3600
  proxied = false
}

resource "cloudflare_dns_record" "%[1]s_cname_settings" {
  zone_id = "%[2]s"
  name    = "tf-acctest-cname.%[1]s.%[3]s"
  type    = "CNAME"
  content = "target.%[3]s"
  ttl     = 60
  proxied = false

  settings = {
    flatten_cname = false
  }

  tags = []
}

resource "cloudflare_dns_record" "%[1]s_with_comment" {
  zone_id = "%[2]s"
  name    = "tf-acctest-comment.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.21"
  ttl     = 3600
  proxied = false
  comment = "Updated comment"  # Changed comment
  tags    = []
}`, rnd, zoneID, domain)
}

// Test config for CAA record with data field - tests modified_on drift fix
func testAccCheckCloudflareRecordConfigModifiedOnDriftCAA(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_modified_on_drift_caa.tf", rnd, zoneID, domain)
}

// Test config for records with settings field
func testAccCheckCloudflareRecordConfigSettingsDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_settings_drift.tf", rnd, zoneID, domain)
}

// Test config for records with settings removed
func testAccCheckCloudflareRecordConfigSettingsDriftRemoved(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_settings_drift_removed.tf", rnd, zoneID, domain)
}

// Test config for comprehensive drift prevention testing
func testAccCheckCloudflareRecordConfigComprehensiveDrift(zoneID, rnd, domain string, updated bool) string {
	content := "192.168.1.1"
	if updated {
		content = "192.168.1.2"
	}
	return acctest.LoadTestCase("dns_record_comprehensive_drift.tf", rnd, zoneID, domain, content)
}

func testAccCheckCloudflareRecordRecreated(before, after *dns.RecordResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of Record Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflareRecordDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_dns_record" {
			continue
		}

		_, err := client.DNS.Records.Get(context.Background(), rs.Primary.ID, dns.RecordGetParams{
			ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
		})
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			if apierr.StatusCode != 404 {
				return fmt.Errorf("Record still exists")
			}
		}
	}

	return nil
}

func testAccManuallyDeleteRecord(record *dns.RecordResponse, zoneID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()
		_, err := client.DNS.Records.Delete(context.Background(), record.ID, dns.RecordDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareRecordAttributes(record *dns.RecordResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.10" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordAttributesUpdated(record *dns.RecordResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.11" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordExists(n string, record *dns.RecordResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := acctest.SharedClient()
		foundRecord, err := client.DNS.Records.Get(context.Background(), rs.Primary.ID, dns.RecordGetParams{
			ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
		})
		if err != nil {
			return err
		}

		if foundRecord.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*record = *foundRecord

		return nil
	}
}

func testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_basic.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigApex(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_apex.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigLOC(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("record_config_loc.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_srv.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigCAA(resourceName, zoneID, name string, ttl int) string {
	return acctest.LoadTestCase("record_config_caa.tf", resourceName, zoneID, name, ttl)
}

func testAccCheckCloudflareRecordConfigProxied(zoneID, domain, name, rnd string) string {
	return acctest.LoadTestCase("record_config_proxied.tf", zoneID, domain, name, rnd)
}

func testAccCheckCloudflareRecordConfigNewValue(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_new_value.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigChangeType(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("record_config_change_type.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigChangeHostname(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("record_config_change_hostname.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigTtlValidation(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("record_config_ttl_validation.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, name, zoneName, proxied, ttl string) string {
	return acctest.LoadTestCase("record_config_explicit_proxied.tf", zoneID, name, zoneName, proxied, ttl)
}

func testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, name, zoneName string) string {
	return acctest.LoadTestCase("record_config_mx_with_priority_zero.tf", zoneID, name, zoneName)
}

func testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd, zoneName string) string {
	return acctest.LoadTestCase("record_config_https.tf", zoneID, rnd, zoneName)
}

func testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_svcb.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordNullMX(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("record_null_mx.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigMultipleTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_multiple_tags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigNoTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("record_config_no_tags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordDNSKEY(zoneID, name string) string {
	return acctest.LoadTestCase("record_dnskey.tf", zoneID, name)
}

func testAccCheckCloudflareRecordConfigWithoutComment(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_without_comment.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigCommentModified(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("dns_record_comment_modified.tf", zoneID, name, rnd, domain)
}

// TestAccCloudflareRecord_ModifiedOnDrift tests for issues for drift
func TestAccCloudflareRecord_ModifiedOnDrift(t *testing.T) {
	var record dns.RecordResponse
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-acctest-%s", rnd)
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", name)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				// Create a record with settings and tags
				Config: testAccCloudflareRecordModifiedOnDriftInitial(zoneID, name, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "A"),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(resourceName, "proxied", "false"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_on"),
				),
			},
			{
				// Apply the same config again - should show no changes
				Config:             testAccCloudflareRecordModifiedOnDriftInitial(zoneID, name, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // This should pass with our fix, fail without it
			},
			{
				// Test with explicit empty settings
				Config: testAccCloudflareRecordModifiedOnDriftWithEmptySettings(zoneID, name, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.11"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_on"),
				),
			},
			{
				// Apply the same config again - should show no changes
				Config:             testAccCloudflareRecordModifiedOnDriftWithEmptySettings(zoneID, name, domain),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // This should pass with our fix, fail without it
			},
		},
	})
}

func testAccCloudflareRecordModifiedOnDriftInitial(zoneID, name, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "%[2]s"
  type    = "A"
  content = "192.168.0.10"
  ttl     = 1
  proxied = false
  tags    = []
}`, zoneID, name, domain)
}

func testAccCloudflareRecordModifiedOnDriftWithEmptySettings(zoneID, name, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_dns_record" "%[2]s" {
  zone_id  = "%[1]s"
  name     = "%[2]s"
  type     = "A"
  content  = "192.168.0.11"
  ttl      = 1
  proxied  = false
  tags     = []
  settings = {}
}`, zoneID, name, domain)
}

// TestAccCloudflareRecord_ModifiedOnConsistency specifically tests that modified_on
// doesn't change when no actual changes are made to the record.
// This test would catch the issue reported in #6438.
func TestAccCloudflareRecord_ModifiedOnConsistency(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	var initialModifiedOn string
	var initialCreatedOn string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create a DNS record and capture its modified_on timestamp
			{
				Config: testAccCheckCloudflareRecordConfigModifiedOnSimple(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", "A"),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_on"),
					resource.TestCheckResourceAttrSet(resourceName, "created_on"),
					// Capture the initial timestamps
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Resource not found: %s", resourceName)
						}
						initialModifiedOn = rs.Primary.Attributes["modified_on"]
						initialCreatedOn = rs.Primary.Attributes["created_on"]
						log.Printf("Initial modified_on: %s", initialModifiedOn)
						return nil
					},
				),
			},
			// Step 2: Re-apply same config - modified_on should NOT change
			{
				Config: testAccCheckCloudflareRecordConfigModifiedOnSimple(zoneID, rnd, domain),
				PreConfig: func() {
					// Wait to ensure time has passed
					time.Sleep(2 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					// Verify timestamps haven't changed when no actual changes were made
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Resource not found: %s", resourceName)
						}

						currentModifiedOn := rs.Primary.Attributes["modified_on"]
						currentCreatedOn := rs.Primary.Attributes["created_on"]

						if currentModifiedOn != initialModifiedOn {
							return fmt.Errorf("modified_on changed without actual changes: was %s, now %s",
								initialModifiedOn, currentModifiedOn)
						}

						if currentCreatedOn != initialCreatedOn {
							return fmt.Errorf("created_on changed unexpectedly: was %s, now %s",
								initialCreatedOn, currentCreatedOn)
						}

						log.Printf("modified_on remained stable: %s", currentModifiedOn)
						return nil
					},
				),
			},
			// Step 3: Make an actual change and verify modified_on DOES change
			{
				Config: testAccCheckCloudflareRecordConfigModifiedOnUpdated(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.20"),
					// Verify modified_on changed after actual update
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Resource not found: %s", resourceName)
						}

						currentModifiedOn := rs.Primary.Attributes["modified_on"]
						currentCreatedOn := rs.Primary.Attributes["created_on"]

						if currentModifiedOn == initialModifiedOn {
							return fmt.Errorf("modified_on didn't change after actual update: still %s",
								currentModifiedOn)
						}

						if currentCreatedOn != initialCreatedOn {
							return fmt.Errorf("created_on changed unexpectedly: was %s, now %s",
								initialCreatedOn, currentCreatedOn)
						}

						// Update our reference for next check
						initialModifiedOn = currentModifiedOn
						log.Printf("New modified_on after update: %s", currentModifiedOn)
						return nil
					},
				),
			},
			// Step 4: Re-apply without changes and ensure modified_on stays stable
			{
				Config: testAccCheckCloudflareRecordConfigModifiedOnUpdated(zoneID, rnd, domain),
				PreConfig: func() {
					time.Sleep(2 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Resource not found: %s", resourceName)
						}

						currentModifiedOn := rs.Primary.Attributes["modified_on"]

						if currentModifiedOn != initialModifiedOn {
							return fmt.Errorf("modified_on changed without actual changes after update: was %s, now %s",
								initialModifiedOn, currentModifiedOn)
						}

						log.Printf("modified_on remained stable after update: %s", currentModifiedOn)
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareRecordConfigModifiedOnSimple(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "%[2]s.%[3]s"
  type    = "A"
  content = "192.168.0.10"
  ttl     = 3600
}`, zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigModifiedOnUpdated(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "%[2]s.%[3]s"
  type    = "A"
  content = "192.168.0.20"
  ttl     = 3600
}`, zoneID, rnd, domain)
}

func suppressTrailingDots(k, old, new string, d *schema.ResourceData) bool {
	newTrimmed := strings.TrimSuffix(new, ".")

	// Ensure to distinguish values consists of dots only.
	if newTrimmed == "" {
		return old == new
	}

	return strings.TrimSuffix(old, ".") == newTrimmed
}
