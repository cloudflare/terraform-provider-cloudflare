package dns_record_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	cfold "github.com/cloudflare/cloudflare-go"
	cloudflare "github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	// Clean up test DNS records only
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	records, _, err := client.ListDNSRecords(context.Background(), cfold.ZoneIdentifier(zoneID), cfold.ListDNSRecordsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare DNS records: %s", err))
		return err
	}

	if len(records) == 0 {
		log.Print("[DEBUG] No Cloudflare DNS records to sweep")
		return nil
	}

	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	
	for _, record := range records {
		shouldDelete := false
		
		// Delete test records - those that start with tf-acctest- or contain terraform test patterns
		if strings.HasPrefix(record.Name, "tf-acctest-") || strings.Contains(record.Name, "tf-acctest") {
			shouldDelete = true
		}
		
		// Also clean up apex domain records if they are A/AAAA/CNAME records that could conflict with tests
		// Only delete apex records that are likely from tests (A/AAAA records pointing to test IPs or CNAME records)
		if domain != "" && record.Name == domain {
			if record.Type == "A" && (strings.HasPrefix(record.Content, "192.168.") || strings.HasPrefix(record.Content, "10.0.") || strings.HasPrefix(record.Content, "172.16.")) {
				shouldDelete = true
			} else if record.Type == "AAAA" && strings.HasPrefix(record.Content, "2001:db8:") {
				shouldDelete = true
			} else if record.Type == "CNAME" {
				shouldDelete = true
			}
		}
		
		if shouldDelete {
			tflog.Info(ctx, fmt.Sprintf("Deleting test DNS record ID: %s, Name: %s, Type: %s, Content: %s", record.ID, record.Name, record.Type, record.Content))
			err := client.DeleteDNSRecord(context.Background(), cfold.ZoneIdentifier(zoneID), record.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete DNS record %s: %s", record.ID, err))
			}
		}
	}

	return nil
}

func TestAccCloudflareRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var record cfold.DNSRecord
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
	var afterCreate, afterUpdate cfold.DNSRecord
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
	return acctest.LoadTestCase("dnsrecordcnamecase.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigTagsDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dnsrecordtagsdrift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigSimpleDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dnsrecordsimpledrift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigFQDNNormalize(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dnsrecordfqdnnormalize.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigComputedDrift(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dnsrecordcomputeddrift.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigDriftRepro(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("dnsrecorddriftrepo.tf", rnd, zoneID, domain)
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

func testAccCheckCloudflareRecordRecreated(before, after *cfold.DNSRecord) resource.TestCheckFunc {
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

func testAccManuallyDeleteRecord(record *cfold.DNSRecord, zoneID string) resource.TestCheckFunc {
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

func testAccCheckCloudflareRecordAttributes(record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.10" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordAttributesUpdated(record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.11" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordExists(n string, record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundRecord, err := client.GetDNSRecord(context.Background(), cfold.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundRecord.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*record = foundRecord

		return nil
	}
}

func testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigbasic.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigApex(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigapex.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigLOC(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigloc.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigsrv.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigCAA(resourceName, zoneID, name string, ttl int) string {
	return acctest.LoadTestCase("recordconfigcaa.tf", resourceName, zoneID, name, ttl)
}

func testAccCheckCloudflareRecordConfigProxied(zoneID, domain, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigproxied.tf", zoneID, domain, name, rnd)
}

func testAccCheckCloudflareRecordConfigNewValue(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfignewvalue.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigChangeType(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("recordconfigchangetype.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigChangeHostname(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigchangehostname.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigTtlValidation(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("recordconfigttlvalidation.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, name, zoneName, proxied, ttl string) string {
	return acctest.LoadTestCase("recordconfigexplicitproxied.tf", zoneID, name, zoneName, proxied, ttl)
}

func testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, name, zoneName string) string {
	return acctest.LoadTestCase("recordconfigmxwithpriorityzero.tf", zoneID, name, zoneName)
}

func testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd, zoneName string) string {
	return acctest.LoadTestCase("recordconfighttps.tf", zoneID, rnd, zoneName)
}

func testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigsvcb.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordNullMX(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordnullmx.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigMultipleTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigmultipletags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigNoTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfignotags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordDNSKEY(zoneID, name string) string {
	return acctest.LoadTestCase("recorddnskey.tf", zoneID, name)
}

func suppressTrailingDots(k, old, new string, d *schema.ResourceData) bool {
	newTrimmed := strings.TrimSuffix(new, ".")

	// Ensure to distinguish values consists of dots only.
	if newTrimmed == "" {
		return old == new
	}

	return strings.TrimSuffix(old, ".") == newTrimmed
}
