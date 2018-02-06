package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"time"
)

func TestAccCloudFlareRecord_Basic(t *testing.T) {
	var record cloudflare.DNSRecord
	testStartTime := time.Now().UTC()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &record),
					testAccCheckCloudFlareRecordAttributes(&record),
					testAccCheckCloudFlareRecordDates("cloudflare_record.foobar", &record, testStartTime),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "name", "terraform"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "value", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "proxiable", "false"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "data.%", "0"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "metadata.%", "2"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "metadata.auto_added", "false"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "metadata.managed_by_apps", "false"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_Apex(t *testing.T) {
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigApex, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &record),
					testAccCheckCloudFlareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "name", "@"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "value", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_Proxied(t *testing.T) {
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigProxied, domain, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &record),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "name", "terraform"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "proxied", "true"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "type", "CNAME"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "value", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "proxiable", "true"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "data.%", "0"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_Updated(t *testing.T) {
	var record cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &record),
					testAccCheckCloudFlareRecordAttributes(&record),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "name", "terraform"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "value", "192.168.0.10"),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigNewValue, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &record),
					testAccCheckCloudFlareRecordAttributesUpdated(&record),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "name", "terraform"),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_record.foobar", "value", "192.168.0.11"),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_typeForceNewRecord(t *testing.T) {
	var afterCreate, afterUpdate cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterCreate),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigChangeType, domain, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterUpdate),
					testAccCheckCloudFlareRecordRecreated(t, &afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_hostnameForceNewRecord(t *testing.T) {
	var afterCreate, afterUpdate cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterCreate),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigChangeHostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterUpdate),
					testAccCheckCloudFlareRecordRecreated(t, &afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudFlareRecord_domainForceNewRecord(t *testing.T) {
	var afterCreate, afterUpdate cloudflare.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	otherDomain := os.Getenv("CLOUDFLARE_DOMAIN_2")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterCreate),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckCloudFlareRecordConfigBasic, otherDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRecordExists("cloudflare_record.foobar", &afterUpdate),
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

const testAccCheckCloudFlareRecordConfigBasic = `
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "terraform"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`

const testAccCheckCloudFlareRecordConfigApex = `
resource "cloudflare_record" "foobar" {
	domain = "%s"
	name = "@"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`

const testAccCheckCloudFlareRecordConfigProxied = `
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "terraform"
	value = "%s"
	type = "CNAME"
	proxied = true
}`

const testAccCheckCloudFlareRecordConfigNewValue = `
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "terraform"
	value = "192.168.0.11"
	type = "A"
	ttl = 3600
}`

const testAccCheckCloudFlareRecordConfigChangeType = `
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "terraform"
	value = "%s"
	type = "CNAME"
	ttl = 3600
}`

const testAccCheckCloudFlareRecordConfigChangeHostname = `
resource "cloudflare_record" "foobar" {
	domain = "%s"

	name = "changed"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}`
