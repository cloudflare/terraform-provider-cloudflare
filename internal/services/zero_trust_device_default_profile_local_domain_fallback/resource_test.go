package zero_trust_device_default_profile_local_domain_fallback_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_device_default_profile_local_domain_fallback", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_default_profile_local_domain_fallback",
		F:    testSweepCloudflareZeroTrustDeviceDefaultProfileLocalDomainFallback,
	})
}

func testSweepCloudflareZeroTrustDeviceDefaultProfileLocalDomainFallback(r string) error {
	ctx := context.Background()
	// Device Default Profile Local Domain Fallback is an account-level fallback domain setting for the default device profile.
	// It's a singleton setting per account, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust Device Default Profile Local Domain Fallback doesn't require sweeping (account setting)")
	return nil
}

func TestAccCloudflareFallbackDomain_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_local_domain_fallback.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("example domain")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("dns_server").AtSliceIndex(0), knownvalue.StringExact("1.0.0.1")),
				},
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				PlanOnly: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				PlanOnly: true,
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
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_local_domain_fallback.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("second example domain")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("example.net")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("dns_server").AtSliceIndex(0), knownvalue.StringExact("1.1.1.1")),
				},
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				PlanOnly: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareFallbackDomain_MultipleDomains(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_local_domain_fallback.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain_MultipleDomains(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains"), knownvalue.ListSizeExact(4)),
				},
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain_MultipleDomains(rnd, accountID),
				PlanOnly: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     accountID,
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain_MultipleDomains(rnd, accountID),
				PlanOnly: true,
			},
			{
				Config: testAccCloudflareDefaultFallbackDomain_MultipleDomainsUpdated(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains"), knownvalue.ListSizeExact(3)),
				},
			},
			{
				Config:   testAccCloudflareDefaultFallbackDomain_MultipleDomainsUpdated(rnd, accountID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareDefaultFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
	return acctest.LoadTestCase("defaultfallbackdomain.tf", rnd, accountID, description, suffix, dns_server)
}

func testAccCloudflareDefaultFallbackDomain_MultipleDomains(rnd, accountID string) string {
	return acctest.LoadTestCase("defaultfallbackdomain_multiple-domains.tf", rnd, accountID)
}

func testAccCloudflareDefaultFallbackDomain_MultipleDomainsUpdated(rnd, accountID string) string {
	return acctest.LoadTestCase("defaultfallbackdomain_multiple-domains_updated.tf", rnd, accountID)
}
