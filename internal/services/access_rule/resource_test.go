package access_rule_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

func testAccCheckCloudflareAccessRuleDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_rule" {
			continue
		}

		_, err := client.AccountAccessRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudflare_access_rule still exists")
		}

		_, err = client.ZoneAccessRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudflare_access_rule still exists")
		}
	}

	return nil
}

func TestAccCloudflareAccessRule_AccountASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testASN(t, testAccessRuleAccountConfig, accountID, consts.AccountIDSchemaKey, rnd, name)
}

func TestAccCloudflareAccessRule_ZoneASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testASN(t, testAccessRuleZoneConfig, zoneID, consts.ZoneIDSchemaKey, rnd, name)
}

func testASN(t *testing.T, configFn configFunc, configKey, schemaKey, rnd, name string) {
	pathPrefix := "zones"
	if schemaKey == consts.AccountIDSchemaKey {
		pathPrefix = "accounts"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: configFn(configKey, "challenge", "this is notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: configFn(configKey, "block", "this is updated notes", "asn", "AS112", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "asn"),
					resource.TestCheckResourceAttr(name, "configuration.value", "AS112"),
				),
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s/%s", pathPrefix, configKey, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modified_on"},
			},
		},
	})
}

func TestAccCloudflareAccessRule_AccountIPRange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testIPRange(t, testAccessRuleAccountConfig, accountID, consts.AccountIDSchemaKey, rnd, name)
}

func TestAccCloudflareAccessRule_ZoneIPRange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testIPRange(t, testAccessRuleZoneConfig, zoneID, consts.ZoneIDSchemaKey, rnd, name)
}

func testIPRange(t *testing.T, configFn configFunc, configKey, schemaKey, rnd, name string) {
	pathPrefix := "zones"
	if schemaKey == consts.AccountIDSchemaKey {
		pathPrefix = "accounts"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config:      configFn(configKey, "block", "this is notes", "ip_range", "104.16.0.0/25", rnd),
				ExpectError: regexp.MustCompile(`IPv4 .* has prefix length .* are allowed`),
			},
			{
				Config: configFn(configKey, "challenge", "this is notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.value", "104.16.0.0/24"),
				),
			},
			{
				// Note: Only notes + mode can be changed in place.
				Config: configFn(configKey, "block", "this is updated notes", "ip_range", "104.16.0.0/24", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip_range"),
					resource.TestCheckResourceAttr(name, "configuration.value", "104.16.0.0/24"),
				),
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s/%s", pathPrefix, configKey, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modified_on"},
			},
		},
	})
}

func TestAccCloudflareAccessRule_AccountIPv6(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testIPv6(t, testAccessRuleAccountConfig, accountID, consts.AccountIDSchemaKey, rnd, name)
}

func TestAccCloudflareAccessRule_ZoneIPv6(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testIPv6(t, testAccessRuleZoneConfig, zoneID, consts.ZoneIDSchemaKey, rnd, name)
}

func testIPv6(t *testing.T, configFn configFunc, configKey, schemaKey, rnd, name string) {
	pathPrefix := "zones"
	if schemaKey == consts.AccountIDSchemaKey {
		pathPrefix = "accounts"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config:      configFn(configKey, "block", "this is notes", "ip6", "2001:0db8::", rnd),
				ExpectError: regexp.MustCompile(`IPv6 address must be in long form`),
			},
			{
				Config: configFn(configKey, "block", "this is notes", "ip6", "2001:0db8:0000:0000:0000:0000:0000:0000", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
			{
				Config: configFn(configKey, "block", "this is notes", "ip6", "2001:0db8:0000:0000:0000:0000:0000:0000", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip6"),
					resource.TestCheckResourceAttr(name, "configuration.value", "2001:0db8:0000:0000:0000:0000:0000:0000"),
				),
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s/%s", pathPrefix, configKey, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modified_on"},
			},
		},
	})
}

func TestAccCloudflareAccessRule_AccountIPv4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testIPv4(t, testAccessRuleAccountConfig, accountID, consts.AccountIDSchemaKey, rnd, name)
}

func TestAccCloudflareAccessRule_ZoneIPv4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testIPv4(t, testAccessRuleZoneConfig, zoneID, consts.ZoneIDSchemaKey, rnd, name)
}

func testIPv4(t *testing.T, configFn configFunc, configKey, schemaKey, rnd, name string) {
	pathPrefix := "zones"
	if schemaKey == consts.AccountIDSchemaKey {
		pathPrefix = "accounts"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: configFn(configKey, "challenge", "this is notes", "ip", "192.0.2.1", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip"),
					resource.TestCheckResourceAttr(name, "configuration.value", "192.0.2.1"),
				),
			},
			{
				Config: configFn(configKey, "block", "this is updated notes", "ip", "192.0.2.1", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "ip"),
					resource.TestCheckResourceAttr(name, "configuration.value", "192.0.2.1"),
				),
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s/%s", pathPrefix, configKey, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modified_on"},
			},
		},
	})
}

func TestAccCloudflareAccessRule_AccountCountry(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCountry(t, testAccessRuleAccountConfig, accountID, consts.AccountIDSchemaKey, rnd, name)
}

func TestAccCloudflareAccessRule_ZoneCountry(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_access_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testCountry(t, testAccessRuleZoneConfig, zoneID, consts.ZoneIDSchemaKey, rnd, name)
}

func testCountry(t *testing.T, configFn configFunc, configKey, schemaKey, rnd, name string) {
	pathPrefix := "zones"
	if schemaKey == consts.AccountIDSchemaKey {
		pathPrefix = "accounts"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: configFn(configKey, "challenge", "this is notes", "country", "US", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is notes"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
					resource.TestCheckResourceAttr(name, "configuration.target", "country"),
					resource.TestCheckResourceAttr(name, "configuration.value", "US"),
				),
			},
			{
				Config: configFn(configKey, "block", "this is updated notes", "country", "US", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, schemaKey, configKey),
					resource.TestCheckResourceAttr(name, "notes", "this is updated notes"),
					resource.TestCheckResourceAttr(name, "mode", "block"),
					resource.TestCheckResourceAttr(name, "configuration.target", "country"),
					resource.TestCheckResourceAttr(name, "configuration.value", "US"),
				),
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s/%s", pathPrefix, configKey, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modified_on"},
			},
		},
	})
}

type configFunc = func(string, string, string, string, string, string) string

func testAccessRuleAccountConfig(accountID, mode, notes, target, value, rnd string) string {
	return acctest.LoadTestCase("accessruleaccountconfig.tf", accountID, mode, notes, target, value, rnd)
}

func testAccessRuleZoneConfig(zoneID, mode, notes, target, value, rnd string) string {
	return acctest.LoadTestCase("accessrulezoneconfig.tf", zoneID, mode, notes, target, value, rnd)
}
