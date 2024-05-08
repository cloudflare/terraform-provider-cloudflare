package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccessPolicy_ServiceToken(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyServiceTokenConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.service_token.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessPolicy_AnyServiceToken(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAnyServiceTokenConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessPolicy_WithZoneID(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	updatedName := fmt.Sprintf("%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyWithZoneID(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
				),
			},
			{
				Config: testAccessPolicyWithZoneIDUpdated(rnd, zone, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", updatedName),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
				),
			},
		},
	})
}

func testAccessPolicyServiceTokenConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_service_token" "%[1]s" {
      account_id = "%[3]s"
      name       = "%[1]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        service_token = ["${cloudflare_access_service_token.%[1]s.id}"]
      }
    }

  `, resourceID, zone, accountID)
}

func testAccessPolicyAnyServiceTokenConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        any_valid_service_token = true
      }
    }

  `, resourceID, zone, accountID)
}

func testAccessPolicyWithZoneID(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      zone_id        = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        any_valid_service_token = true
      }
    }

  `, resourceID, zone, zoneID)
}

func testAccessPolicyWithZoneIDUpdated(resourceID, zone, zoneID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s-updated"
      zone_id        = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        any_valid_service_token = true
      }
    }

  `, resourceID, zone, zoneID)
}

func TestAccCloudflareAccessPolicy_Group(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGroupConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.group.#", "1"),
				),
			},
		},
	})
}

func testAccessPolicyGroupConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_group" "%[1]s" {
        account_id     = "%[3]s"
        name           = "%[1]s"

        include {
          ip = ["127.0.0.1/32"]
        }
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        group = [cloudflare_access_group.%[1]s.id]
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_MTLS(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyMTLSConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.certificate", "true"),
				),
			},
		},
	})
}

func testAccessPolicyMTLSConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include {
        certificate = true
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_CommonName(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyCommonNameConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.common_name", "example@example.com"),
				),
			},
		},
	})
}

func testAccessPolicyCommonNameConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        common_name = "example@example.com"
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_EmailDomain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailDomainConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
				),
			},
		},
	})
}

func testAccessPolicyEmailDomainConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"
      session_duration = "12h"

      include {
        email_domain = ["example.com"]
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Emails(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEmailsConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.email.#", "2"),
					resource.TestCheckResourceAttr(name, "include.0.email.0", "a@example.com"),
					resource.TestCheckResourceAttr(name, "include.0.email.1", "b@example.com"),
				),
			},
		},
	})
}

func testAccessPolicyEmailsConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        email = ["a@example.com", "b@example.com"]
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Everyone(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyEveryoneConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.everyone", "true"),
				),
			},
		},
	})
}

func testAccessPolicyEveryoneConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        everyone = true
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_IPs(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIPsConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.ip.#", "2"),
					resource.TestCheckResourceAttr(name, "include.0.ip.0", "10.0.0.1/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip.1", "10.0.0.2/32"),
				),
			},
		},
	})
}

func testAccessPolicyIPsConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        ip = ["10.0.0.1/32", "10.0.0.2/32"]
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_AuthMethod(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyAuthMethodConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.auth_method", "hwk"),
				),
			},
		},
	})
}

func testAccessPolicyAuthMethodConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        auth_method = "hwk"
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Geo(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyGeoConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.geo.#", "2"),
					resource.TestCheckResourceAttr(name, "include.0.geo.0", "US"),
					resource.TestCheckResourceAttr(name, "include.0.geo.1", "AU"),
				),
			},
		},
	})
}

func testAccessPolicyGeoConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        geo = ["US", "AU"]
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_Okta(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyOktaConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.okta.0.name.#", "2"),
					resource.TestCheckResourceAttr(name, "include.0.okta.0.name.0", "jacob-group"),
					resource.TestCheckResourceAttr(name, "include.0.okta.0.name.1", "jacob-group1"),
					resource.TestCheckResourceAttr(name, "include.0.okta.0.identity_provider_id", "225934dc-14e4-4f55-87be-f5d798d23f91"),
				),
			},
		},
	})
}

func testAccessPolicyOktaConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        okta {
          name = ["jacob-group", "jacob-group1"]
          identity_provider_id = "225934dc-14e4-4f55-87be-f5d798d23f91"
        }
      }
    }
  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_PurposeJustification(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyPurposeJustificationConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessPolicyHasPJ(name, cloudflare.AccountIdentifier(accountID)),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "purpose_justification_required", "true"),
					resource.TestCheckResourceAttr(name, "purpose_justification_prompt", "Why should we let you in?"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessPolicyHasPJ(n string, accessIdentifier *cloudflare.ResourceContainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessPolicy ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		var foundAccessPolicy cloudflare.AccessPolicy
		var err error

		foundAccessPolicy, err = client.GetAccessPolicy(context.Background(), accessIdentifier, cloudflare.GetAccessPolicyParams{ApplicationID: rs.Primary.Attributes["application_id"], PolicyID: rs.Primary.ID})
		if err != nil {
			return err
		}

		if foundAccessPolicy.ID != rs.Primary.ID {
			return fmt.Errorf("AccessPolicy not found")
		}

		if !(foundAccessPolicy.PurposeJustificationPrompt != nil && *foundAccessPolicy.PurposeJustificationPrompt == rs.Primary.Attributes["purpose_justification_prompt"]) {
			return fmt.Errorf("AccessPolicy is missing purpose_justification_prompt")
		}

		pjRequired, _ := strconv.ParseBool(rs.Primary.Attributes["purpose_justification_required"])

		if !(foundAccessPolicy.PurposeJustificationRequired != nil && *foundAccessPolicy.PurposeJustificationRequired == pjRequired) {
			return fmt.Errorf("AccessPolicy is missing purpose_justification_required")
		}

		return nil
	}
}

func testAccessPolicyPurposeJustificationConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        email = ["a@example.com", "b@example.com"]
      }

      purpose_justification_required = "true"
      purpose_justification_prompt = "Why should we let you in?"
    }
  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ApprovalGroup(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyApprovalGroupConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessPolicyHasApprovalGroups(name, cloudflare.AccountIdentifier(accountID)),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "purpose_justification_required", "true"),
					resource.TestCheckResourceAttr(name, "purpose_justification_prompt", "Why should we let you in?"),
					resource.TestCheckResourceAttr(name, "approval_required", "true"),
					resource.TestCheckResourceAttr(name, "approval_group.0.email_addresses.0", "test1@example.com"),
					resource.TestCheckResourceAttr(name, "approval_group.0.email_addresses.1", "test2@example.com"),
					resource.TestCheckResourceAttr(name, "approval_group.0.email_addresses.2", "test3@example.com"),
					resource.TestCheckResourceAttr(name, "approval_group.1.email_addresses.0", "test4@example.com"),
					resource.TestCheckResourceAttr(name, "approval_group.0.approvals_needed", "2"),
					resource.TestCheckResourceAttr(name, "approval_group.1.approvals_needed", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessPolicy_Reusable(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyReusableConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", "a@example.com"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessPolicyHasApprovalGroups(n string, accessIdentifier *cloudflare.ResourceContainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessPolicy ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		var foundAccessPolicy cloudflare.AccessPolicy
		var err error

		foundAccessPolicy, err = client.GetAccessPolicy(context.Background(), accessIdentifier, cloudflare.GetAccessPolicyParams{ApplicationID: rs.Primary.Attributes["application_id"], PolicyID: rs.Primary.ID})
		if err != nil {
			return err
		}

		if foundAccessPolicy.ID != rs.Primary.ID {
			return fmt.Errorf("AccessPolicy not found")
		}

		if !(foundAccessPolicy.ApprovalGroups != nil) {
			return fmt.Errorf("AccessPolicy is missing approval_groups")
		}

		approvalRequired, _ := strconv.ParseBool(rs.Primary.Attributes["approval_required"])

		if !(foundAccessPolicy.ApprovalRequired != nil && *foundAccessPolicy.ApprovalRequired == approvalRequired) {
			return fmt.Errorf("AccessPolicy is missing approval_required")
		}

		return nil
	}
}

func testAccessPolicyApprovalGroupConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      purpose_justification_required = "true"
      purpose_justification_prompt = "Why should we let you in?"
      approval_required = "true"

      include {
        email = ["a@example.com", "b@example.com"]
      }

      approval_group {
        email_addresses = ["test1@example.com", "test2@example.com", "test3@example.com"]
        approvals_needed = "2"
      }

      approval_group {
        email_addresses = ["test4@example.com"]
        approvals_needed = "1"
      }
    }
  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_ExternalEvaluation(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyExternalEvalautionConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "include.0.external_evaluation.0.evaluate_url", "https://example.com"),
					resource.TestCheckResourceAttr(name, "include.0.external_evaluation.0.keys_url", "https://example.com/keys"),
				),
			},
		},
	})
}

func testAccessPolicyExternalEvalautionConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = cloudflare_access_application.%[1]s.id
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
		external_evaluation {
			evaluate_url = "https://example.com"
			keys_url = "https://example.com/keys"
		  }
      }
    }

  `, resourceID, zone, accountID)
}

func TestAccCloudflareAccessPolicy_IsolationRequired(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_access_policy." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessPolicyIsolationRequiredConfig(rnd, zone, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "isolation_required", "true"),
				),
			},
		},
	})
}

func testAccessPolicyIsolationRequiredConfig(resourceID, zone, accountID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_teams_account" "%[1]s" {
		account_id = "%[3]s"
		tls_decrypt_enabled = true
		protocol_detection_enabled = true
		activity_log_enabled = true
		url_browser_isolation_enabled = true
		non_identity_browser_isolation_enabled = false
		block_page {
		  name = "%[1]s"
		  enabled = true
		  footer_text = "hello"
		  header_text = "hello"
		  logo_path = "https://example.com"
		  background_color = "#000000"
		  mailto_subject = "hello"
		  mailto_address = "test@cloudflare.com"
		}
		body_scanning {
		  inspection_mode = "deep"
		}
		fips {
		  tls = true
		}
		antivirus {
		  enabled_download_phase = true
		  enabled_upload_phase = false
		  fail_closed = true
		}
		proxy {
		  tcp = true
		  udp = false
		  root_ca = true
		}
		logging {
		  redact_pii = true
		  settings_by_rule_type {
			dns {
			  log_all = false
			  log_blocks = true
			}
			http {
			  log_all = true
			  log_blocks = true
			}
			l4 {
			  log_all = false
			  log_blocks = true
			}
		  }
		}
		ssh_session_log {
		  public_key = "testvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0="
		}
		payload_log {
		  public_key = "EmpOvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0="
		}
	  }

    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
	  depends_on = ["cloudflare_teams_account.%[1]s"]
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = cloudflare_access_application.%[1]s.id
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include {
        email = ["a@example.com", "b@example.com"]
      }

      isolation_required = "true"
    }

  `, resourceID, zone, accountID)
}

func testAccessPolicyReusableConfig(resourceID, accountID string) string {
	return fmt.Sprintf(`
    resource "cloudflare_access_policy" "%[1]s" {
      name           = "%[1]s"
      account_id     = "%[2]s"
      decision       = "allow"
      include {
        email = ["a@example.com"]
      }
    }
  `, resourceID, accountID)
}
