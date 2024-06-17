package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessRule_AccountASN(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig(accountID, "block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "AS112"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_ZoneASN(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleZoneConfig(zoneID, "challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleZoneConfig(zoneID, "block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "AS112"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_IPRange(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "challenge", "this is notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "104.16.0.0/24"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig(accountID, "block", "this is updated notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "104.16.0.0/24"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_IPv6(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "block", "this is notes", "ip6", "2001:0db8::", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
			{
				Config: testAccessRuleAccountConfig(accountID, "block", "this is notes", "ip6", "2001:0db8:0000:0000:0000:0000:0000:0000", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
		},
	})
}

func testAccessRuleAccountConfig(accountID, mode, notes, target, value, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_rule" "%[6]s" {
  account_id = "%[1]s"
  notes = "%[3]s"
  mode = "%[2]s"
  configuration {
    target = "%[4]s"
    value = "%[5]s"
  }
}`, accountID, mode, notes, target, value, rnd)
}

func testAccessRuleZoneConfig(zoneID, mode, notes, target, value, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_rule" "%[6]s" {
  zone_id = "%[1]s"
  notes = "%[3]s"
  mode = "%[2]s"
  configuration {
    target = "%[4]s"
    value = "%[5]s"
  }
}`, zoneID, mode, notes, target, value, rnd)
}

func TestValidateAccessRuleConfigurationIPRange(t *testing.T) {
	ipRangeValid := map[string]bool{
		"192.168.0.1/32":           false,
		"192.168.0.1/24":           true,
		"192.168.0.1/64":           false,
		"192.168.0.1/31":           false,
		"192.168.0.1/16":           true,
		"fd82:0f75:cf0d:d7b3::/64": true,
		"fd82:0f75:cf0d:d7b3::/48": true,
		"fd82:0f75:cf0d:d7b3::/32": true,
		"fd82:0f75:cf0d:d7b3::/63": false,
		"fd82:0f75:cf0d:d7b3::/16": false,
	}

	for ipRange, valid := range ipRangeValid {
		warnings, errors := validateAccessRuleConfigurationIPRange(ipRange)
		isValid := len(errors) == 0
		if len(warnings) != 0 {
			t.Fatalf("ipRange is either invalid or valid, no room for warnings")
		}
		if isValid != valid {
			t.Fatalf("%s resulted in %v, expected %v", ipRange, isValid, valid)
		}
	}
}
