package access_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_access_rule", &resource.Sweeper{
		Name: "cloudflare_access_rule",
		F:    testSweepCloudflareAccessRules,
	})
}

func testSweepCloudflareAccessRules(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	// Sweep account-level access rules
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID != "" {
		accountRulesResp, err := client.ListAccountAccessRules(ctx, accountID, cloudflare.AccessRule{}, 1)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch account access rules: %s", err))
			return fmt.Errorf("failed to fetch account access rules: %w", err)
		}

		for _, rule := range accountRulesResp.Result {
			// Use standard filtering helper on the notes field
			if !utils.ShouldSweepResource(rule.Notes) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting account access rule: %s (account: %s)", rule.ID, accountID))
			_, err := client.DeleteAccountAccessRule(ctx, accountID, rule.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete account access rule %s: %s", rule.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted account access rule: %s", rule.ID))
		}
	}

	// Sweep zone-level access rules
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID != "" {
		zoneRulesResp, err := client.ListZoneAccessRules(ctx, zoneID, cloudflare.AccessRule{}, 1)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch zone access rules: %s", err))
			return fmt.Errorf("failed to fetch zone access rules: %w", err)
		}

		for _, rule := range zoneRulesResp.Result {
			// Use standard filtering helper on the notes field
			if !utils.ShouldSweepResource(rule.Notes) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting zone access rule: %s (zone: %s)", rule.ID, zoneID))
			_, err := client.DeleteZoneAccessRule(ctx, zoneID, rule.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete zone access rule %s: %s", rule.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted zone access rule: %s", rule.ID))
		}
	}

	return nil
}

func TestAccCloudflareAccessRule_AccountASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig(accountID, "block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_ZoneASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleZoneConfig(zoneID, "challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleZoneConfig(zoneID, "block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_IPRange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "challenge", "this is notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.value", "104.16.0.0/24"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: testAccessRuleAccountConfig(accountID, "block", "this is updated notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.value", "104.16.0.0/24"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessRule_IPv6(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessRuleAccountConfig(accountID, "block", "this is notes", "ip6", "2001:0db8::", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.value", "2001:0db8::"),
				),
			},
			{
				Config: testAccessRuleAccountConfig(accountID, "block", "this is notes", "ip6", "2001:0db8:0000:0000:0000:0000:0000:0000", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
		},
	})
}

func testAccessRuleAccountConfig(accountID, mode, notes, target, value, rnd string) string {
	return acctest.LoadTestCase("accessruleaccountconfig.tf", accountID, mode, notes, target, value, rnd)
}

func testAccessRuleZoneConfig(zoneID, mode, notes, target, value, rnd string) string {
	return acctest.LoadTestCase("accessrulezoneconfig.tf", zoneID, mode, notes, target, value, rnd)
}

// func TestValidateAccessRuleConfigurationIPRange(t *testing.T) {
// 	ipRangeValid := map[string]bool{
// 		"192.168.0.1/32":           false,
// 		"192.168.0.1/24":           true,
// 		"192.168.0.1/64":           false,
// 		"192.168.0.1/31":           false,
// 		"192.168.0.1/16":           true,
// 		"fd82:0f75:cf0d:d7b3::/64": true,
// 		"fd82:0f75:cf0d:d7b3::/48": true,
// 		"fd82:0f75:cf0d:d7b3::/32": true,
// 		"fd82:0f75:cf0d:d7b3::/63": false,
// 		"fd82:0f75:cf0d:d7b3::/16": false,
// 	}

// 	for ipRange, valid := range ipRangeValid {
// 		warnings, errors := validateAccessRuleConfigurationIPRange(ipRange)
// 		isValid := len(errors) == 0
// 		if len(warnings) != 0 {
// 			t.Fatalf("ipRange is either invalid or valid, no room for warnings")
// 		}
// 		if isValid != valid {
// 			t.Fatalf("%s resulted in %v, expected %v", ipRange, isValid, valid)
// 		}
// 	}
// }
