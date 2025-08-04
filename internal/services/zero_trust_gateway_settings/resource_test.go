package zero_trust_gateway_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestAccCloudflareTeamsAccounts_ConfigurationBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountBasic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("footer_text"), knownvalue.StringExact("footer")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("header_text"), knownvalue.StringExact("header")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_subject"), knownvalue.StringExact("hello")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_address"), knownvalue.StringExact("test@cloudflare.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("background_color"), knownvalue.StringExact("#000000")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("logo_path"), knownvalue.StringExact("https://example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("msg")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://hello.com/")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
		},
	})
}

func testAccCloudflareTeamsAccountBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountbasic.tf", rnd, accountID)
}

func TestAccCloudflareTeamsAccounts_ConfigurationMinimal(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					// Verify minimal settings configuration
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
					// Verify computed attributes exist
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
		},
	})
}

func TestAccCloudflareTeamsAccounts_ConfigurationUpdate(t *testing.T) {
	t.Skip("Skipping  due to provider drift bug. The `activity_log` block is incorrectly planned for removal after apply. Check with ZT Gateway team.")
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCloudflareTeamsAccountFull(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("extended_email_matching").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCloudflareTeamsAccountUpdated(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("shallow")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("extended_email_matching").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("updated msg")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://updated.com/")),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
		},
	})
}

func TestAccCloudflareTeamsAccounts_ConfigurationBodyScanningModes(t *testing.T) {
	t.Skip("Skipping  due to provider drift bug. The `activity_log` block is incorrectly planned for removal after apply.")
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountBodyScanningDeep(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),
				},
			},
			{
				Config: testAccCloudflareTeamsAccountBodyScanningShallow(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("shallow")),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
		},
	})
}

func testAccCloudflareTeamsAccountMinimal(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountminimal.tf", rnd, accountID)
}

func testAccCloudflareTeamsAccountFull(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountfull.tf", rnd, accountID)
}

func testAccCloudflareTeamsAccountUpdated(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountupdated.tf", rnd, accountID)
}

func testAccCloudflareTeamsAccountBodyScanningDeep(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    body_scanning = {
      inspection_mode = "deep"
    }
  }
}`, rnd, accountID)
}

func testAccCloudflareTeamsAccountBodyScanningShallow(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    body_scanning = {
      inspection_mode = "shallow"
    }
  }
}`, rnd, accountID)
}
