package zero_trust_access_service_token_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID    = os.Getenv("CLOUDFLARE_ZONE_ID")
)

func TestAccCloudflareAccessServiceToken_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// Service Tokens endpoint does not yet support the API tokens and it
	// results in misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName+"-updated", cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckNoResourceAttr(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfig(resourceName, resourceName+"-updated", cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName+"-updated"),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckNoResourceAttr(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
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

// 	rnd := utils.GenerateRandomResourceName()
// 	var initialState terraform.ResourceState

// 	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
// 	resourceName := strings.Split(name, ".")[1]
// 	expirationTime := 365

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { acctest.TestAccPreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessServiceTokenDestroy,
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
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessServiceTokenDestroy,
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_service_token.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessServiceTokenDestroy,
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
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.AccountIdentifier(accountID), "8760h"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckNoResourceAttr(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessServiceTokenDestroy,
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
				),
			},
			{
				Config: testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName, resourceName, cloudflare.ZoneIdentifier(zoneID), "8760h"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttrSet(name, "client_id"),
					resource.TestCheckNoResourceAttr(name, "client_secret"),
					resource.TestCheckResourceAttrSet(name, "expires_at"),
					resource.TestCheckResourceAttr(name, "duration", "8760h"),
				),
			},
		},
	})
}

func testCloudflareAccessServiceTokenBasicConfig(resourceName string, tokenName string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("cloudflareaccessservicetokenbasicconfig.tf", resourceName, tokenName, identifier.Type, identifier.Identifier)
}

func testCloudflareAccessServiceTokenBasicConfigWithDuration(resourceName string, tokenName string, identifier *cloudflare.ResourceContainer, duration string) string {
	return acctest.LoadTestCase("cloudflareaccessservicetokenbasicconfigwithduration.tf", resourceName, tokenName, identifier.Type, identifier.Identifier, duration)
}

func testAccCheckCloudflareAccessServiceTokenDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
