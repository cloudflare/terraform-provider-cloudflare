package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAccessRule_ASN(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_rule." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig("challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig("block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
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

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig("challenge", "this is notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "104.16.0.0/24"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig("block", "this is updated notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
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

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig("block", "this is notes", "ip6", "2001:0db8::", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
			{
				Config: testAccessRuleAccountConfig("block", "this is notes", "ip6", "2001:0db8:0000:0000:0000:0000:0000:0000", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
		},
	})
}

func testAccessRuleAccountConfig(mode, notes, target, value, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_rule" "%[5]s" {
  notes = "%[2]s"
  mode = "%[1]s"
  configuration {
    target = "%[3]s"
    value = "%[4]s"
  }
}`, mode, notes, target, value, rnd)
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
