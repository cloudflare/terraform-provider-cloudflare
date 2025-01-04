package zero_trust_device_posture_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestAccCloudflareDevicePostureRule_OsVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.version", "10.0.1"),
					resource.TestCheckResourceAttr(name, "input.operator", "=="),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "mac"),
					resource.TestCheckResourceAttr(name, "input.version", "10.0.1"),
					resource.TestCheckResourceAttr(name, "input.operator", "=="),
					resource.TestCheckResourceAttr(name, "input.os_version_extra", "(a)"),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "os_version"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "linux"),
					resource.TestCheckResourceAttr(name, "input.version", "1.0.0"),
					resource.TestCheckResourceAttr(name, "input.os_distro_name", "ubuntu"),
					resource.TestCheckResourceAttr(name, "input.os_distro_revision", "1.0.0"),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "domain_joined"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "match.0.platform", "windows"),
					resource.TestCheckResourceAttr(name, "input.domain", "example.com"),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
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
					resource.TestCheckResourceAttr(name, "input.enabled", "true"),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
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
					resource.TestCheckResourceAttr(name, "input.require_all", "true"),
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureRuleDestroy,
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
					resource.TestCheckResourceAttr(name, "input.require_all", "false"),
					resource.TestCheckResourceAttr(name, "input.check_disks.0", "C"),
					resource.TestCheckResourceAttr(name, "input.check_disks.1", "D"),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureRuleConfigOsVersion(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigosversion.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigOsVersionExtra(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigosversionextra.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigLinuxDistro(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfiglinuxdistro.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDomainJoined(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdomainjoined.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionRequireAll(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdiskencryptionrequireall.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigDiskEncryptionCheckDisks(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigdiskencryptioncheckdisks.tf", rnd, accountID)
}

func testAccCloudflareDevicePostureRuleConfigFirewall(rnd, accountID string) string {
	return acctest.LoadTestCase("devicepostureruleconfigfirewall.tf", rnd, accountID)
}

func testAccCheckCloudflareDevicePostureRuleDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
