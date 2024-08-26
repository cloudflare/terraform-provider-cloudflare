package record_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
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

func init() {
	resource.AddTestSweepers("cloudflare_record", &resource.Sweeper{
		Name: "cloudflare_record",
		F:    testSweepCloudflareRecord,
	})
}

func testSweepCloudflareRecord(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	// Clean up the account level rulesets
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	records, _, err := client.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare DNS records: %s", err))
	}

	if len(records) == 0 {
		log.Print("[DEBUG] No Cloudflare DNS records to sweep")
		return nil
	}

	for _, record := range records {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare DNS record ID: %s", record.ID))
		//nolint:errcheck
		client.DeleteDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), record.ID)
	}

	return nil
}

func TestAccCloudflareRecord_Basic(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	testStartTime := time.Now().UTC()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-basic", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					testAccCheckCloudflareRecordDates(resourceName, &record, testStartTime),
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acctest-basic"),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "value", "192.168.0.10"),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("tf-acctest-basic.%s", zoneName)),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "metadata.auto_added", "false"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "tag2"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CaseInsensitive(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-case-insensitive", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acctest-case-insensitive"),
				),
			},
			{
				Config:   testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-CASE-INSENSITIVE", rnd),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", "tf-acctest-case-insensitive"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Apex(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigApex(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", "@"),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "value", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_LOC(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigLOC(zoneID, "tf-acctest-loc", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "value", "37 46 46.000 N 122 23 35.000 W 0.00 100.00 0.00 0.00"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.0.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.0.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.0.lat_minutes", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.0.lat_seconds", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.0.lat_direction", "N"),
					resource.TestCheckResourceAttr(resourceName, "data.0.long_degrees", "122"),
					resource.TestCheckResourceAttr(resourceName, "data.0.long_minutes", "23"),
					resource.TestCheckResourceAttr(resourceName, "data.0.long_seconds", "35"),
					resource.TestCheckResourceAttr(resourceName, "data.0.long_direction", "W"),
					resource.TestCheckResourceAttr(resourceName, "data.0.altitude", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.0.size", "100"),
					resource.TestCheckResourceAttr(resourceName, "data.0.precision_horz", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.0.precision_vert", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SRV(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("_xmpp-client._tcp.%s", rnd)),
					resource.TestCheckResourceAttr(resourceName, "hostname", fmt.Sprintf("_xmpp-client._tcp.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "value", "0 5222 talk.l.google.com"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.0.priority", "5"),
					resource.TestCheckResourceAttr(resourceName, "data.0.weight", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.0.port", "5222"),
					resource.TestCheckResourceAttr(resourceName, "data.0.target", "talk.l.google.com"),
					resource.TestCheckResourceAttr(resourceName, "data.0.service", "_xmpp-client"),
					resource.TestCheckResourceAttr(resourceName, "data.0.proto", "_tcp"),
					resource.TestCheckResourceAttr(resourceName, "data.0.name", rnd+"."+domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CAA(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.0.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.0.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.0.value", "letsencrypt.org"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 300),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.0.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.0.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.0.value", "letsencrypt.org"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Proxied(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

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
					resource.TestCheckResourceAttr(resourceName, "value", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Updated(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-update"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNewValue(zoneID, recordName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributesUpdated(&record),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_typeForceNewRecord(t *testing.T) {
	t.Parallel()
	var afterCreate, afterUpdate cloudflare.DNSRecord
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-type-force-new"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd),
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

func TestAccCloudflareRecord_hostnameForceNewRecord(t *testing.T) {
	t.Parallel()
	var afterCreate, afterUpdate cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-hostname-force-new"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigChangeHostname(zoneID, recordName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterUpdate),
					testAccCheckCloudflareRecordRecreated(&afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var afterCreate, afterRecreate cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(name, &afterCreate),
					testAccManuallyDeleteRecord(&afterCreate),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(name, &afterRecreate),
					testAccCheckCloudflareRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_TtlValidation(t *testing.T) {
	t.Parallel()
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
				ExpectError: regexp.MustCompile(fmt.Sprintf("error validating record %s: ttl must be set to 1 when `proxied` is true", recordName)),
			},
		},
	})
}

func TestAccCloudflareRecord_ExplicitProxiedFalse(t *testing.T) {
	t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_record." + rnd

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
	t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_record." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "priority", "0"),
					resource.TestCheckResourceAttr(resourceName, "value", "mail.terraform.cfapi.net"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_TtlValidationUpdate(t *testing.T) {
	t.Parallel()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-ttl-validation"
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigProxied(zoneID, domain, recordName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(name, &cloudflare.DNSRecord{}),
				),
			},
			{
				Config:      testAccCheckCloudflareRecordConfigTtlValidation(zoneID, recordName, domain, rnd),
				ExpectError: regexp.MustCompile(fmt.Sprintf("error validating record %s: ttl must be set to 1 when `proxied` is true", recordName)),
			},
		},
	})
}

func TestAccCloudflareRecord_HTTPS(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.0.priority", "1"),
					resource.TestCheckResourceAttr(name, "data.0.target", "."),
					resource.TestCheckResourceAttr(name, "data.0.value", `alpn="h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SVCB(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.0.priority", "2"),
					resource.TestCheckResourceAttr(name, "data.0.target", "foo."),
					resource.TestCheckResourceAttr(name, "data.0.value", `alpn="h3,h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_MXNull(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordNullMX(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "value", "."),
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
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

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
					resource.TestCheckResourceAttr(name, "data.0.flags", "257"),
					resource.TestCheckResourceAttr(name, "data.0.protocol", "13"),
					resource.TestCheckResourceAttr(name, "data.0.algorithm", "2"),
					resource.TestCheckResourceAttr(name, "data.0.public_key", "mdsswUyr3DPW132mOi8V9xESWE8jTo0dxCjjnopKl+GqJxpVXckHAeF+KkxLbxILfDLUT0rAK9iUzy1L53eKGQ=="),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_ClearTags(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMultipleTags(zoneID, rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tags.#", "2"),
					resource.TestCheckResourceAttr(name, "comment", "this is a comment"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNoTags(zoneID, rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tags.#", "0"),
					resource.TestCheckResourceAttr(name, "comment", ""),
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

func testAccCheckCloudflareRecordRecreated(before, after *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of Record Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflareRecordDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_record" {
			continue
		}

		_, err := client.GetDNSRecord(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccManuallyDeleteRecord(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		err := client.DeleteDNSRecord(context.Background(), cloudflare.ZoneIdentifier(record.ZoneID), record.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareRecordAttributes(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.10" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordAttributesUpdated(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.11" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordDates(n string, record *cloudflare.DNSRecord, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": record.CreatedOn, "modified_on": record.ModifiedOn} {
			timeStamp, err := time.Parse(time.RFC3339Nano, rs.Primary.Attributes[timeStampAttr])
			if err != nil {
				return err
			}

			if timeStamp != serverVal {
				return fmt.Errorf("state value of %s: %s is different than server created value: %s", timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("state value of %s: %s should be greater than test start time: %s", timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccCheckCloudflareRecordExists(n string, record *cloudflare.DNSRecord) resource.TestCheckFunc {
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
		foundRecord, err := client.GetDNSRecord(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
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

func testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigbasic.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigApex(zoneID, rnd string) string {
	return acctest.LoadTestCase("recordconfigapex.tf", zoneID, rnd)
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

func testAccCheckCloudflareRecordConfigNewValue(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfignewvalue.tf", zoneID, name, rnd)
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

func testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd string) string {
	return acctest.LoadTestCase("recordconfighttps.tf", zoneID, rnd)
}

func testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd string) string {
	return acctest.LoadTestCase("recordconfigsvcb.tf", zoneID, rnd)
}

func testAccCheckCloudflareRecordNullMX(zoneID, rnd string) string {
	return acctest.LoadTestCase("recordnullmx.tf", rnd, zoneID)
}

func testAccCheckCloudflareRecordConfigMultipleTags(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigmultipletags.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigNoTags(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfignotags.tf", zoneID, name, rnd)
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
