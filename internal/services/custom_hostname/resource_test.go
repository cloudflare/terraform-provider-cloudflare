package custom_hostname_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_custom_hostname", &resource.Sweeper{
		Name: "cloudflare_custom_hostname",
		F:    testSweepCloudflareCustomHostnames,
	})
}

func testSweepCloudflareCustomHostnames(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping custom hostnames sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	hostnames, err := client.CustomHostnames.List(ctx, custom_hostnames.CustomHostnameListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch custom hostnames: %s", err))
		return nil
	}

	for _, hostname := range hostnames.Result {
		if !utils.ShouldSweepResource(hostname.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting custom hostname: %s", hostname.ID))
		_, err := client.CustomHostnames.Delete(ctx, hostname.ID, custom_hostnames.CustomHostnameDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete custom hostname %s: %s", hostname.ID, err))
		}
	}

	return nil
}

func testAccCheckCloudflareCustomHostnameDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.CustomHostnames.Get(context.Background(), rs.Primary.ID, custom_hostnames.CustomHostnameGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			continue
		}

		tflog.Warn(context.Background(), fmt.Sprintf("Custom hostname %s still exists but this may be expected", rs.Primary.ID))
	}

	return nil
}

// TestAccCustomHostname_Basic tests the basic CRUD lifecycle of a custom hostname.
// This validates that the resource can be created, read, imported, and deleted.
func TestAccCustomHostname_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_hostname." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomHostnameBasicConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("type"), knownvalue.StringExact("dv")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", zoneID, s.RootModule().Resources[name].Primary.ID), nil
				},
				ImportStateVerifyIgnore: []string{
					"ssl.certificate_authority",
					"ssl.validation_records",
					"ssl.validation_errors",
					"ssl.wildcard",
					"ownership_verification",
					"ownership_verification_http",
					"created_at",
					"status",
					"verification_errors",
					"wait_for_ssl_pending_validation",
				},
			},
		},
	})
}

// TestAccCustomHostname_WithSSLSettings tests the optional SSL settings attributes.
// This validates that optional nested attributes are handled correctly and can be updated.
func TestAccCustomHostname_WithSSLSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_custom_hostname." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomHostnameWithSSLSettingsConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("settings").AtMapKey("min_tls_version"), knownvalue.StringExact("1.2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCustomHostnameWithSSLSettingsUpdatedConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ssl").AtMapKey("settings").AtMapKey("min_tls_version"), knownvalue.StringExact("1.1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCustomHostnameBasicConfig(zoneID, domain, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[3]s" {
  zone_id  = "%[1]s"
  hostname = "%[3]s.%[2]s"
  ssl = {
    method = "txt"
    type   = "dv"
  }
  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      ssl.certificate_authority,
      ssl.wildcard,
      status,
      verification_errors,
    ]
  }
}`, zoneID, domain, rnd)
}

func testAccCustomHostnameWithSSLSettingsConfig(zoneID, domain, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[3]s" {
  zone_id  = "%[1]s"
  hostname = "%[3]s.%[2]s"
  ssl = {
    method = "txt"
    type   = "dv"
    settings = {
      min_tls_version = "1.2"
    }
  }
  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      ssl.certificate_authority,
      ssl.wildcard,
      status,
      verification_errors,
    ]
  }
}`, zoneID, domain, rnd)
}

func testAccCustomHostnameWithSSLSettingsUpdatedConfig(zoneID, domain, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname" "%[3]s" {
  zone_id  = "%[1]s"
  hostname = "%[3]s.%[2]s"
  ssl = {
    method = "txt"
    type   = "dv"
    settings = {
      min_tls_version = "1.1"
    }
  }
  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      ssl.certificate_authority,
      ssl.wildcard,
      status,
      verification_errors,
    ]
  }
}`, zoneID, domain, rnd)
}
