package cloudflare

import (
	"errors"
	"fmt"
	"testing"

	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudflareSpectrumApplication_Basic(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "origin_direct.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://120.120.102.10:23"),
					resource.TestCheckResourceAttr(name, "origin_port", "22"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_OriginDNS(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigOriginDNS(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "origin_dns.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_dns.0.name", fmt.Sprintf("ssh.origin.%s", zone)),
					resource.TestCheckResourceAttr(name, "origin_port", "22"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_Update(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	var initialId string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://120.120.102.10:23"),
				),
			},
			{
				PreConfig: func() {
					initialId = spectrumApp.ID
				},
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasicUpdated(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					func(state *terraform.State) error {
						if initialId != spectrumApp.ID {
							// want in place update
							return fmt.Errorf("spectrum application id is different after second config applied ( %s -> %s )",
								initialId, spectrumApp.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://81.120.102.10:23"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_CreateAfterManualDestroy(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	var initialId string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					testAccManuallyDeleteSpectrumApplication(name, &spectrumApp, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					func(state *terraform.State) error {
						if initialId == spectrumApp.ID {
							return fmt.Errorf("spectrum application id is unchanged even after we thought we deleted it ( %s )",
								spectrumApp.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareSpectrumApplicationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_spectrum_application" {
			continue
		}

		_, err := client.SpectrumApplication(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("spectrum application still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudflareSpectrumApplicationExists(n string, spectrumApp *cloudflare.SpectrumApplication) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundSpectrumApplication, err := client.SpectrumApplication(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		*spectrumApp = foundSpectrumApplication

		return nil
	}
}

func testAccCheckCloudflareSpectrumApplicationIDIsValid(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		if len(rs.Primary.ID) != 32 {
			return fmt.Errorf("invalid id %q, should be a string of length 32", rs.Primary.ID)
		}

		if zoneId, ok := rs.Primary.Attributes["zone_id"]; !ok || len(zoneId) < 1 {
			return errors.New("zone_id is unset, should always be set with id")
		}
		return nil
	}
}

func testAccManuallyDeleteSpectrumApplication(name string, spectrumApp *cloudflare.SpectrumApplication, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = spectrumApp.ID
		err := client.DeleteSpectrumApplication(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareSpectrumApplicationConfigBasic(zoneName, ID string) string {
	return fmt.Sprintf(`
resource "cloudflare_spectrum_application" "%[2]s" {
  zone_id  = "${data.cloudflare_zone.test.id}"
  protocol = "tcp/22"
  
  dns = {
    "type" = "CNAME"
    "name" = "ssh.${data.cloudflare_zone.test.zone}"
  }
  
  origin_direct = ["tcp://120.120.102.10:23"]
  origin_port   = 22
}

data "cloudflare_zone" "test" {
  zone = "%[1]s"
}
`, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigOriginDNS(zoneName, ID string) string {
	return fmt.Sprintf(`
resource "cloudflare_spectrum_application" "%[2]s" {
  zone_id  = "${data.cloudflare_zone.test.id}"
  protocol = "tcp/22"
  
  dns = {
    "type" = "CNAME"
    "name" = "ssh.${data.cloudflare_zone.test.zone}"
  }
  
  origin_dns = {
    name = "ssh.origin.${data.cloudflare_zone.test.zone}"
  }
  origin_port   = 22
}

data "cloudflare_zone" "test" {
  zone = "%[1]s"
}
`, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigBasicUpdated(zoneName, ID string) string {
	return fmt.Sprintf(`
resource "cloudflare_spectrum_application" "%[2]s" {
  zone_id  = "${data.cloudflare_zone.test.id}"
  protocol = "tcp/22"

  dns = {
	"type" = "CNAME"
	"name" = "ssh.${data.cloudflare_zone.test.zone}"
  }

  origin_direct = ["tcp://81.120.102.10:23"]
  origin_port   = 22
}

data "cloudflare_zone" "test" {
	zone = "%[1]s"
}`, zoneName, ID)
}
