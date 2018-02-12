package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"time"

	"regexp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudFlareRecord_Basic(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	testStartTime := time.Now().UTC()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigBasic(domain, "tf-acctest-basic"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					testAccCheckCloudFlareRecordAttributes(&record),
					testAccCheckCloudFlareRecordDates(resourceName, &record, testStartTime),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-basic"),
					resource.TestCheckResourceAttr(
						resourceName, "domain", domain),
					resource.TestCheckResourceAttr(
						resourceName, "value", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", fmt.Sprintf("tf-acctest-basic.%s", domain)),
					resource.TestMatchResourceAttr(
						resourceName, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(
						resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(
						resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttr(
						resourceName, "metadata.auto_added", "false"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_Apex(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigApex(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					testAccCheckCloudFlareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "@"),
					resource.TestCheckResourceAttr(
						resourceName, "domain", domain),
					resource.TestCheckResourceAttr(
						resourceName, "value", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_LOC(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigLOC(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "value", "37 46 46.000 N 122 23 35.000 W 0m 100m 0m 0m"),
					resource.TestCheckResourceAttr(
						resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "12"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_SRV(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigSRV(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "hostname", fmt.Sprintf("_xmpp-client._tcp.%s", domain)),
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

func TestAccCloudFlareRecord_Proxied(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigProxied(domain, "tf-acctest-proxied"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
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

func TestAccCloudFlareRecord_Updated(t *testing.T) {
	t.Parallel()
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	recordName := "tf-acctest-update"
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigBasic(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					testAccCheckCloudFlareRecordAttributes(&record),
				),
			},
			{
				Config: testAccCheckCloudFlareRecordConfigNewValue(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &record),
					testAccCheckCloudFlareRecordAttributesUpdated(&record),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_typeForceNewRecord(t *testing.T) {
	t.Parallel()
	var afterCreate, afterUpdate cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	recordName := "tf-acctest-type-force-new"
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigBasic(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudFlareRecordConfigChangeType(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &afterUpdate),
					testAccCheckCloudFlareRecordRecreated(t, &afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_hostnameForceNewRecord(t *testing.T) {
	t.Parallel()
	var afterCreate, afterUpdate cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	recordName := "tf-acctest-hostname-force-new"
	resourceName := fmt.Sprintf("cloudflare_record.foobar")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRecordConfigBasic(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudFlareRecordConfigChangeHostname(domain, recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists(resourceName, &afterUpdate),
					testAccCheckCloudFlareRecordRecreated(t, &afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func testAccCheckCloudFlareRecordRecreated(t *testing.T,
	before, after *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			t.Fatalf("Expected change of Record Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudFlareRecordDestroy(s *terraform.State) error {
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

func testAccCheckCloudFlareRecordAttributes(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if record.Content != "192.168.0.10" {
			return fmt.Errorf("Bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudFlareRecordAttributesUpdated(record *cloudflare.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if record.Content != "192.168.0.11" {
			return fmt.Errorf("Bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudFlareRecordDates(n string, record *cloudflare.DNSRecord, testStartTime time.Time) resource.TestCheckFunc {
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

func testAccCheckCloudFlareRecordExists(n string, record *cloudflare.DNSRecord) resource.TestCheckFunc {
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

func testAccCheckCloudFlareRecordConfigBasic(zone, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "%s"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zone, name)
}

func testAccCheckCloudFlareRecordConfigApex(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%s"
	name = "@"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zone)
}

func testAccCheckCloudFlareRecordConfigLOC(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%[1]s"
	name = "%[1]s"
	data {
	  "lat_degrees" =  "37"
	  "lat_minutes" = "46"
	  "lat_seconds" = "46"
	  "lat_direction" = "N"
	  "long_degrees" = "122"
	  "long_minutes" = "23"
	  "long_seconds" = "35"
	  "long_direction" = "W"
	  "altitude" = 0
	  "size" = 100
	  "precision_horz" = 0
	  "precision_vert" = 0
	}
	type = "LOC"
	ttl = 3600
}`, zone)
}

func testAccCheckCloudFlareRecordConfigSRV(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%[1]s"
	name = "%[1]s"
	data {
	  "priority" = 5
      "weight" = 0
      "port" = 5222
      "target" = "talk.l.google.com"
      "service" = "_xmpp-client"
      "proto" = "_tcp"
      "name" = "%[1]s"
	}
	type = "SRV"
	ttl = 3600
}`, zone)
}

func testAccCheckCloudFlareRecordConfigProxied(zone, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%[1]s"

	name = "%[2]s"
	value = "%[1]s"
	type = "CNAME"
	proxied = true
}`, zone, name)
}

func testAccCheckCloudFlareRecordConfigNewValue(zone, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "%s"
	value = "192.168.0.11"
	type = "A"
	ttl = 3600
}`, zone, name)
}

func testAccCheckCloudFlareRecordConfigChangeType(zone, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%[1]s"

	name = "%[2]s"
	value = "%[1]s"
	type = "CNAME"
	ttl = 3600
}`, zone, name)
}

func testAccCheckCloudFlareRecordConfigChangeHostname(zone, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "%s-changed"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`, zone, name)
}
