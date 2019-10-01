package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareRecord_Basic(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	testStartTime := time.Now().UTC()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-basic"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					testAccCheckCloudflareRecordDates(resourceName, &record, testStartTime),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-basic"),
					resource.TestCheckResourceAttr(
						resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(
						resourceName, "value", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", fmt.Sprintf("tf-acctest-basic.%s", zoneName)),
					resource.TestMatchResourceAttr(
						resourceName, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(
						resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(
						resourceName, "metadata.%", "3"),
					resource.TestCheckResourceAttr(
						resourceName, "metadata.auto_added", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CaseInsensitive(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-case-insensitive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-case-insensitive"),
				),
			},
			{
				Config:   testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-CASE-INSENSITIVE"),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-case-insensitive"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Apex(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigApex(zoneID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "@"),
					resource.TestCheckResourceAttr(
						resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(
						resourceName, "value", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_LOC(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigLOC(zoneID, "tf-acctest-loc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "value", "37 46 46.000 N 122 23 35.000 W 0.00m 100.00m 0.00m 0.00m"),
					resource.TestCheckResourceAttr(
						resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "12"),
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
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSRV(zoneID, "tf-acctest-srv"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", fmt.Sprintf("_xmpp-client._tcp.tf-acctest-srv.%s", domain)),
					resource.TestCheckResourceAttr(
						resourceName, "value", "0	5222	talk.l.google.com"),
					resource.TestCheckResourceAttr(
						resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "7"),
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
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigProxied(zoneID, domain, "tf-acctest-proxied"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "proxiable", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "proxied", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "type", "CNAME"),
					resource.TestCheckResourceAttr(
						resourceName, "value", domain),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
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
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNewValue(zoneID, recordName),
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
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigChangeType(zoneID, recordName, zoneName),
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
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigChangeHostname(zoneID, recordName),
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
	name := "cloudflare_record.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(name, &afterCreate),
					testAccManuallyDeleteRecord(&afterCreate),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name),
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRecordConfigTtlValidation(zoneID, recordName, zoneName),
				ExpectError: regexp.MustCompile(fmt.Sprintf("error validating record %s: ttl must be set to 1 when `proxied` is true", recordName)),
			},
		},
	})
}

func TestAccCloudflareRecord_TtlValidationUpdate(t *testing.T) {
	t.Parallel()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-ttl-validation"
	name := "cloudflare_record.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigProxied(zoneID, domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(name, &cloudflare.DNSRecord{}),
				),
			},
			{
				Config:      testAccCheckCloudflareRecordConfigTtlValidation(zoneID, recordName, domain),
				ExpectError: regexp.MustCompile(fmt.Sprintf("error validating record %s: ttl must be set to 1 when `proxied` is true", recordName)),
			},
		},
	})
}

func testAccCheckCloudflareRecordRecreated(before, after *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("Expected change of Record Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflareRecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_record" {
			continue
		}

		_, err := client.DNSRecord(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccManuallyDeleteRecord(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		err := client.DeleteDNSRecord(record.ZoneID, record.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareRecordAttributes(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if record.Content != "192.168.0.10" {
			return fmt.Errorf("Bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordAttributesUpdated(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if record.Content != "192.168.0.11" {
			return fmt.Errorf("Bad content: %s", record.Content)
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
				return fmt.Errorf("state value of %s: %s is different than server created value: %s",
					timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("State value of %s: %s should be greater than test start time: %s",
					timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccCheckCloudflareRecordExists(n string, record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundRecord, err := client.DNSRecord(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
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

func testAccCheckCloudflareRecordConfigBasic(zoneID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%s"
	name = "%s"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zoneID, name)
}

func testAccCheckCloudflareRecordConfigApex(zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%s"
	name = "@"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zoneID)
}

func testAccCheckCloudflareRecordConfigLOC(zoneID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
  zone_id = "%s"
  name = "%s"
  data = {
    lat_degrees =  "37"
    lat_minutes = "46"
    lat_seconds = 46.000
    lat_direction = "N"
    long_degrees = "122"
    long_minutes = "23"
    long_seconds = 35.000
    long_direction = "W"
    altitude = 0.00
    size = 100.00
    precision_horz = 0.00
    precision_vert = 0.00
  }
  type = "LOC"
  ttl = 3600
}`, zoneID, name)
}

func testAccCheckCloudflareRecordConfigSRV(zoneID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
  zone_id = "%[1]s"
  name = "%[2]s"
  data = {
    priority = 5
    weight = 0
    port = 5222
    target = "talk.l.google.com"
    service = "_xmpp-client"
    proto = "_tcp"
    name = "%[2]s"
  }
  type = "SRV"
  ttl = 3600
}`, zoneID, name)
}

func testAccCheckCloudflareRecordConfigProxied(zoneID, domain, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%[1]s"
	name = "%[3]s"
	value = "%[2]s"
	type = "CNAME"
	proxied = true
}`, zoneID, domain, name)
}

func testAccCheckCloudflareRecordConfigNewValue(zoneID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%s"
	name = "%s"
	value = "192.168.0.11"
	type = "A"
	ttl = 3600
}`, zoneID, name)
}

func testAccCheckCloudflareRecordConfigChangeType(zoneID, name, zoneName string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%s"
	name = "%s"
	value = "%s"
	type = "CNAME"
	ttl = 3600
}`, zoneID, name, zoneName)
}

func testAccCheckCloudflareRecordConfigChangeHostname(zoneID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%s"
	name = "%s-changed"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zoneID, name)
}

func testAccCheckCloudflareRecordConfigTtlValidation(zoneID, name, zoneName string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	zone_id = "%[1]s"
	name = "%[2]s"
	value = "%[3]s"
	type = "CNAME"
	proxied = true
	ttl = 3600
}`, zoneID, name, zoneName)
}
