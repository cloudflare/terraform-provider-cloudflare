package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareAccessServiceToken_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName+"-updated", cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName+"-updated", cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})
}

// func TestAccCloudflareAccessServiceTokenUpdateWithExpiration(t *testing.T) {
// 	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
// 	// Service Tokens endpoint does not yet support the API tokens and it
// 	// results in misleading state error messages.
// 	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
// 		defer func(apiToken string) {
// 			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
// 		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
// 		os.Setenv("CLOUDFLARE_API_TOKEN", "")
// 	}

// 	rnd := generateRandomResourceName()
// 	var initialState terraform.ResourceState

// 	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
// 	resourceName := strings.Split(name, ".")[1]
// 	expirationTime := 365

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID), expirationTime),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckCloudflareAccessServiceTokenSaved(name, &initialState),
// 					resource.TestCheckResourceAttr(name, "min_days_for_renewal", strconv.Itoa(expirationTime)),
// 				),
// 				//Expiration of 365 will always force a new resource as long as the tokens expire in 365 days in cloudflare
// 				ExpectNonEmptyPlan: true,
// 			},
// 			{
// 				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID), expirationTime),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(name, "min_days_for_renewal", strconv.Itoa(expirationTime)),
// 					testAccCheckCloudflareAccessServiceTokenRenewed(name, &initialState),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// func testAccCheckCloudflareAccessServiceTokenSaved(n string, resourceState *terraform.ResourceState) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("No Access Token ID is set")
// 		}

// 		*resourceState = *rs

// 		return nil
// 	}
// }

func testAccCheckCloudflareAccessServiceTokenRenewed(n string, oldResourceState *terraform.ResourceState) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Access Token ID is set")
		}

		for _, attribute := range []string{"expires_at", "client_secret"} {
			if rs.Primary.Attributes[attribute] == oldResourceState.Primary.Attributes[attribute] {
				return fmt.Errorf("resource attribute '%s' has not changed. Expected change between old state and new", attribute)
			}
		}

		return nil
	}
}

func TestAccCloudflareAccessServiceToken_Delete(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessServiceToken_WithDuration(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.AccountIdentifier(accountID), "forever"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "forever"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.AccountIdentifier(accountID), "8760h"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID), "forever"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "forever"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID), "8760h"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					resource.TestCheckResourceAttrSet(name, "client_secret_version"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessServiceToken_CreateWithClientSecretFields(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName, resourceName, cloudflare.AccountIdentifier(accountID), 3, "2024-11-15T10:30:00Z"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "client_secret_version", "3"),
					resource.TestCheckResourceAttr(name, "previous_client_secret_expires_at", "2024-11-15T10:30:00Z"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessServiceToken_DefaultClientSecretVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrSet(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
					// Verify client_secret_version defaults to 1 when not specified
					resource.TestCheckResourceAttr(name, "client_secret_version", "1"),
					// Verify previous_client_secret_expires_at is not set when not specified
					resource.TestCheckNoResourceAttr(name, "previous_client_secret_expires_at"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessServiceToken_WithClientSecretVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]
	var clientSecretV1, clientSecretV2 string
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessServiceTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName, resourceName, cloudflare.AccountIdentifier(accountID), 1, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrWith(name, "client_secret", func(value string) error {
						if value == "" {
							return errors.New("client secret is empty")
						}
						clientSecretV1 = value
						return nil
					}),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "client_secret_version", "1"),
				),
			},
			{
				Config: testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName, resourceName+"-updated", cloudflare.AccountIdentifier(accountID), 2, "2024-12-01T05:20:03Z"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					// check that the client_secret is marked for update
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue(name, tfjsonpath.New("client_secret")),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrWith(name, "client_secret", func(value string) error {
						if value == "" {
							return errors.New("client secret is empty")
						}
						if clientSecretV1 == value {
							return errors.New("client secret version 1 and version 2 are the same")
						}
						clientSecretV2 = value
						return nil
					}),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "client_secret_version", "2"),
					resource.TestCheckResourceAttr(name, "previous_client_secret_expires_at", "2024-12-01T05:20:03Z"),
				),
			},
			// Bump previous_client_secret_expires_at
			{
				Config: testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName, resourceName+"-updated", cloudflare.AccountIdentifier(accountID), 2, "2026-12-01T05:20:03Z"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					// check that the client_secret is not marked for update
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(name, tfjsonpath.New("client_secret"), knownvalue.NotNull()),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrWith(name, "client_secret", func(value string) error {
						if value == "" {
							return errors.New("client secret is empty")
						}
						if clientSecretV2 != value {
							return errors.New("client secret changed without a version bump")
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "client_secret_version", "2"),
					resource.TestCheckResourceAttr(name, "previous_client_secret_expires_at", "2026-12-01T05:20:03Z"),
				),
			},
			// Bump previous_client_secret
			{
				Config: testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName, resourceName+"-updated", cloudflare.AccountIdentifier(accountID), 3, "2023-12-01T05:20:03Z"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckResourceAttrWith(name, "client_secret", func(value string) error {
						if value == "" {
							return errors.New("client secret is empty")
						}
						if clientSecretV1 == value {
							return errors.New("client secret version 1 and version 3 are the same")
						}
						if clientSecretV2 == value {
							return errors.New("client secret version 2 and version 3 are the same")
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "client_secret_version", "3"),
					resource.TestCheckResourceAttr(name, "previous_client_secret_expires_at", "2023-12-01T05:20:03Z"),
				),
			},
			{
				RefreshState: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func testCloudflareAccessServiceTokenBasicConfig(resourceName string, tokenName string, identifier *cloudflare.ResourceContainer) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  %[3]s_id = "%[4]s"
  name     = "%[2]s"
  min_days_for_renewal = "0"
}`, resourceName, tokenName, identifier.Type, identifier.Identifier)
}

func testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName string, tokenName string, identifier *cloudflare.ResourceContainer, duration string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  %[3]s_id = "%[4]s"
  name     = "%[2]s"
  min_days_for_renewal = "0"
  duration = "%[5]s"
}`, resourceName, tokenName, identifier.Type, identifier.Identifier, duration)
}

func testCloudflareAccessServiceTokenConfigWithClientSecretVersion(resourceName string, tokenName string, identifier *cloudflare.ResourceContainer, version int, previousExpiresAt string) string {
	config := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
	%[3]s_id = "%[4]s"
	name = "%[2]s"
	client_secret_version = %[5]d`, resourceName, tokenName, identifier.Type, identifier.Identifier, version)

	if previousExpiresAt != "" {
		config += fmt.Sprintf(`
	previous_client_secret_expires_at = "%s"`, previousExpiresAt)
	}

	config += `
}`
	return config
}

func testAccCheckCloudflareAccessServiceTokenDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_service_token" {
			continue
		}

		_, err := client.DeleteAccessServiceToken(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("access service token still exists")
		}

		_, err = client.DeleteAccessServiceToken(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("access service token still exists")
		}
	}

	return nil
}
