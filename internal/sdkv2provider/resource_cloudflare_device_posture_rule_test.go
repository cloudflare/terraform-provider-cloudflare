package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareDevicePostureRule_SerialNumber(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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

func TestAccCloudflareDevicePostureRule_OsVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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

func TestAccCloudflareDevicePostureRule_OsVersionExtra(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.version", "10.0.1"),
					resource.TestCheckResourceAttr(name, "input.0.operator", "=="),
					resource.TestCheckResourceAttr(name, "input.0.os_version_extra", "(a)"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_LinuxOsDistro(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "linux"),
					resource.TestCheckResourceAttr(name, "input.0.version", "1.0.0"),
					resource.TestCheckResourceAttr(name, "input.0.os_distro_name", "ubuntu"),
					resource.TestCheckResourceAttr(name, "input.0.os_distro_revision", "1.0.0"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DomainJoined(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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

func TestAccCloudflareDevicePostureRule_Firewall(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "firewall"),
					resource.TestCheckResourceAttr(name, "description", "firewall description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "expiration", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "windows"),
					resource.TestCheckResourceAttr(name, "input.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DiskEncryption_RequireAll(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDiskEncryptionRequireAll(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "disk_encryption"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "expiration", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.require_all", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_DiskEncryption_CheckDisks(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDiskEncryptionCheckDisks(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "disk_encryption"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "expiration", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.require_all", "false"),
					resource.TestCheckResourceAttr(name, "input.0.check_disks.0", "C"),
					resource.TestCheckResourceAttr(name, "input.0.check_disks.1", "D"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_ClientCertificateV2(t *testing.T) {
	skipForDefaultAccount(t, "Assertion requires an active certificate configured")

	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if v := os.Getenv("CLOUDFLARE_DEVICE_POSTURE_CERTIFICATE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_DEVICE_POSTURE_CERTIFICATE_ID must be set for this acceptance test")
	}

	certificateID := os.Getenv("CLOUDFLARE_DEVICE_POSTURE_CERTIFICATE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigClientCertificateV2(rnd, accountID, certificateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "client_certificate_v2"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "expiration", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "linux"),
					resource.TestCheckResourceAttr(name, "input.0.certificate_id", certificateID),
					resource.TestCheckResourceAttr(name, "input.0.check_private_key", "true"),
					resource.TestCheckResourceAttr(name, "input.0.extended_key_usage.0", "clientAuth"),
					resource.TestCheckResourceAttr(name, "input.0.locations.0.trust_stores.0", "system"),
					resource.TestCheckResourceAttr(name, "input.0.locations.0.paths.0", "/path/to/file"),
				),
			},
		},
	})
}

func TestAccCloudflareDevicePostureRule_Intune(t *testing.T) {
	skipForDefaultAccount(t, "Assertion requires active Intune license.")

	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if v := os.Getenv("CLOUDFLARE_DEVICE_POSTURE_INTUNE_CONNECTION_ID"); v == "" {
		t.Fatal("CLOUDFLARE_DEVICE_POSTURE_INTUNE_CONNECTION_ID must be set for this acceptance test")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)
	connectionID := os.Getenv("CLOUDFLARE_DEVICE_POSTURE_INTUNE_CONNECTION_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleIntune(rnd, accountID, connectionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "intune"),
					resource.TestCheckResourceAttr(name, "description", "intune compliance status"),
					resource.TestCheckResourceAttr(name, "schedule", "24h"),
					resource.TestCheckResourceAttr(name, "expiration", "24h"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.0.compliance_status", "compliant"),
					resource.TestCheckResourceAttr(name, "input.0.connection_id", connectionID),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureRuleConfigSerialNumber(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "serial_number"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
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
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
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

func testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "mac"
	}
	input {
		version = "10.0.1"
		operator = "=="
		os_version_extra = "(a)"
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "linux"
	}
	input {
		version = "1.0.0"
        operator = "<"
		os_distro_name = "ubuntu"
		os_distro_revision = "1.0.0"
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "domain_joined"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "windows"
	}
	input {
		domain = "example.com"
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionRequireAll(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "disk_encryption"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "mac"
	}
	input {
		require_all = true
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionCheckDisks(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "disk_encryption"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "mac"
	}
	input {
		require_all = false
		check_disks = ["C", "D"]
	}
}
`, rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigClientCertificateV2(rnd, accountID, certificateID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "client_certificate_v2"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "linux"
	}
	input {
		certificate_id = "%[3]s"
		check_private_key  = true
		extended_key_usage = ["clientAuth"]
		locations {
			paths = [
				"/path/to/file"
			]
			trust_stores = [
				"system"
			]
		}
	}
}
`, rnd, accountID, certificateID)
}

func testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "firewall"
	description               = "firewall description"
	schedule                  = "24h"
	expiration                = "24h"
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
		if rs.Type != "cloudflare_zero_trust_device_posture_rule" {
			continue
		}

		_, err := client.DevicePostureRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Rule still exists")
		}
	}

	return nil
}

func testAccCloudflareDevicePostureRuleIntune(rnd, accountID, connectionID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "intune"
	description               = "intune compliance status"
	schedule                  = "24h"
	expiration                = "24h"
	match {
		platform = "mac"
	}
	input {
		compliance_status = "compliant"
		connection_id = "%[3]s"
	}
}
`, rnd, accountID, connectionID)
}
