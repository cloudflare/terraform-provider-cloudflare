package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareDevicePostureRuleSerialNumber(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "serial_number"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "windows"),
					resource.TestCheckResourceAttr(name, "input.0.id", "asdf-123"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRuleOsVersion(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.version", "10.0.1"),
					resource.TestCheckResourceAttr(name, "input.0.operator", "=="),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRuleDomainJoined(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "domain_joined"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "windows"),
					resource.TestCheckResourceAttr(name, "input.0.domain", "example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRuleFirewall(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "firewall"),
					resource.TestCheckResourceAttr(name, "description", "firewall description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "windows"),
					resource.TestCheckResourceAttr(name, "input.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRuleDiskEncryption(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDiskEncryption(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "disk_encryption"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.require_all", "true"),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "serial_number"
	description               = "My description"
	schedule                  = "24h"
	match {
		platform = "windows"
	}
	input {
		id = "asdf-123"
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	match {
		platform = "mac"
	}
	input {
		version = "10.0.1"
		operator = "=="
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "domain_joined"
	description               = "My description"
	schedule                  = "24h"
	match {
		platform = "windows"
	}
	input {
		domain = "example.com"
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryption(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "disk_encryption"
	description               = "My description"
	schedule                  = "24h"
	match {
		platform = "mac"
	}
	input {
		require_all = true
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "firewall"
	description               = "firewall description"
	schedule                  = "24h"
	match {
		platform = "windows"
	}
	input {
		enabled = true
	}
}
`, rnd, accountID)
}

func testAccCheckCloudflareDevicePostureRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_device_posture_rule" {
			continue
		}

		_, err := client.DevicePostureRule(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Rule still exists")
		}
	}

	return nil
}
