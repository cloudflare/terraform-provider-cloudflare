SetNestedAttributepackage zero_trust_access_mtls_hostname_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	// Clean up any existing settings before running tests using existing sweeper
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_mtls_hostname_settings", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_mtls_hostname_settings",
		F: func(region string) error {
			ctx := context.Background()

			client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			if clientErr != nil {
				return fmt.Errorf("Failed to create Cloudflare client: %w", clientErr)
			}

			// First clear hostname settings
			deletedSettings := cfv1.UpdateAccessMutualTLSHostnameSettingsParams{
				Settings: []cfv1.AccessMutualTLSHostnameSettings{},
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.AccountIdentifier(accountID), deletedSettings)
			if err != nil {
				return fmt.Errorf("Failed to clear Cloudflare Access Mutual TLS hostname settings: %w", err)
			}

			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			_, err = client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.ZoneIdentifier(zoneID), deletedSettings)
			if err != nil {
				return fmt.Errorf("Failed to clear Cloudflare Access Mutual TLS hostname settings: %w", err)
			}

			// Also clean up ALL certificates to prevent conflicts with certificate tests
			// This ensures certificate tests can create certificates without "already exists" errors
			
			// Clean account certificates - be aggressive to prevent test conflicts
			accountCerts, _, err := client.ListAccessMutualTLSCertificates(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListAccessMutualTLSCertificatesParams{})
			if err == nil {
				for _, cert := range accountCerts {
					// Clear hostnames first, then delete
					client.UpdateAccessMutualTLSCertificate(ctx, cfv1.AccountIdentifier(accountID), cfv1.UpdateAccessMutualTLSCertificateParams{
						ID: cert.ID,
						AssociatedHostnames: []string{},
					})
					client.DeleteAccessMutualTLSCertificate(ctx, cfv1.AccountIdentifier(accountID), cert.ID)
				}
			}

			// Clean zone certificates - be aggressive to prevent test conflicts
			zoneCerts, _, err := client.ListAccessMutualTLSCertificates(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ListAccessMutualTLSCertificatesParams{})
			if err == nil {
				for _, cert := range zoneCerts {
					// Clear hostnames first, then delete
					client.UpdateAccessMutualTLSCertificate(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.UpdateAccessMutualTLSCertificateParams{
						ID: cert.ID,
						AssociatedHostnames: []string{},
					})
					client.DeleteAccessMutualTLSCertificate(ctx, cfv1.ZoneIdentifier(zoneID), cert.ID)
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_Account(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
				},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_Zone(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.ZoneIdentifier(zoneID), domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_MultipleHostnames(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	secondHostname := fmt.Sprintf("test.%s", domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsMultipleConfig(rnd, cfv1.AccountIdentifier(accountID), domain, secondHostname),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(1).AtMapKey("hostname"), knownvalue.StringExact(secondHostname)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(1).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(1).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_Update(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccessMutualTLSHostnameSettingsUpdatedConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_BooleanCombinations(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Test all false values
				Config: testAccessMutualTLSHostnameSettingsBooleanConfig(rnd, cfv1.AccountIdentifier(accountID), domain, false, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
				},
			},
			{
				// Test client_certificate_forwarding = true, china_network = false
				Config: testAccessMutualTLSHostnameSettingsBooleanConfig(rnd, cfv1.AccountIdentifier(accountID), domain, false, true),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
				},
			},
			{
				// Test both false (reset to baseline)
				Config: testAccessMutualTLSHostnameSettingsBooleanConfig(rnd, cfv1.AccountIdentifier(accountID), domain, false, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_Import(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("china_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("settings").AtSliceIndex(0).AtMapKey("client_certificate_forwarding"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:                         name,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        accountID,
				ImportStateVerifyIdentifierAttribute: "account_id",
				ImportStateVerifyIgnore:              []string{"settings"},
			},
		},
	})
}

func testAccessMutualTLSHostnameSettingsConfig(rnd string, identifier *cfv1.ResourceContainer, domain string) string {
	return acctest.LoadTestCase("accessmutualtlshostnamesettingsconfig.tf", rnd, identifier.Type, identifier.Identifier, domain)
}

func testAccessMutualTLSHostnameSettingsMultipleConfig(rnd string, identifier *cfv1.ResourceContainer, domain, secondHostname string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
	%[2]s_id = "%[3]s"
	settings = [
		{
			hostname = "%[4]s"
			client_certificate_forwarding = true
			china_network = false
		},
		{
			hostname = "%[5]s"
			client_certificate_forwarding = false
			china_network = false
		}
	]
}
`, rnd, identifier.Type, identifier.Identifier, domain, secondHostname)
}

func testAccessMutualTLSHostnameSettingsUpdatedConfig(rnd string, identifier *cfv1.ResourceContainer, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
	%[2]s_id = "%[3]s"
	settings = [{
		hostname = "%[4]s"
		client_certificate_forwarding = false
		china_network = false
	}]
}
`, rnd, identifier.Type, identifier.Identifier, domain)
}

func testAccessMutualTLSHostnameSettingsBooleanConfig(rnd string, identifier *cfv1.ResourceContainer, domain string, chinaNetwork, clientCertForwarding bool) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
	%[2]s_id = "%[3]s"
	settings = [{
		hostname = "%[4]s"
		client_certificate_forwarding = %[5]t
		china_network = %[6]t
	}]
}
`, rnd, identifier.Type, identifier.Identifier, domain, clientCertForwarding, chinaNetwork)
}
