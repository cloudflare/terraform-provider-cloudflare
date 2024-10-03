package zero_trust_local_domain_fallback_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestAccCloudflareFallbackDomain_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_local_domain_fallback.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		// CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.0.0.1"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
				),
			},
		},
	})
}

func TestAccCloudflareFallbackDomain_DefaultPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_local_domain_fallback.%s", rnd)

	resource.Test(t, resource.TestCase{
		// CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "second example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
				),
			},
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "third example domain", "example.net", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "third example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
					resource.TestCheckResourceAttrSet(name, "policy_id"),
				),
			},
		},
	})
}

func TestAccCloudflareFallbackDomain_WithAttachedPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_local_domain_fallback.%s", rnd)

	resource.Test(t, resource.TestCase{
		// CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "third example domain", "example.net", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "third example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
					resource.TestCheckResourceAttrSet(name, "policy_id"),
				),
			},
		},
	})
}

func testAccCloudflareDefaultFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
	return acctest.LoadTestCase("defaultfallbackdomain.tf", rnd, accountID, description, suffix, dns_server)
}

func testAccCloudflareFallbackDomain(rnd, accountID, description string, suffix string, dns_server string) string {
	return acctest.LoadTestCase("fallbackdomain.tf", rnd, accountID, description, suffix, dns_server)
}

// func testAccCheckCloudflareFallbackDomainDestroy(s *terraform.State) error {
// 	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 	if clientErr != nil {
// 		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 	}

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "cloudflare_zero_trust_local_domain_fallback" {
// 			continue
// 		}

// 		accountID, policyID := parseDevicePolicyID(rs.Primary.ID)

// 		if policyID == "" {
// 			// Deletion of fallback domains should result in a reset to the default fallback domains.
// 			// Default device settings policies, and their fallback domains, cannot be deleted - only reset to default.
// 			result, _ := client.ListFallbackDomains(context.Background(), rs.Primary.ID)
// 			if len(result) == 0 {
// 				return errors.New("deleted Fallback Domain resource has does not include default domains")
// 			}
// 		} else {
// 			// For fallback domains on a non-default device settings policy, only need to check for the deletion of the policy.
// 			_, err := client.GetDeviceSettingsPolicy(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.GetDeviceSettingsPolicyParams{PolicyID: cloudflare.StringPtr(policyID)})
// 			if err == nil {
// 				return fmt.Errorf("device settings policy still exists")
// 			}
// 		}
// 	}

// 	return nil
// }
