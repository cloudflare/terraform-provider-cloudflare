package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	// Provider name for single configuration testing.
	ProviderNameCloudflare = "cloudflare"
)

var (
	// testAccProvider is the "main" provider instance.
	//
	// This Provider can be used in testing code for API calls without requiring
	// the use of saving and referencing specific ProviderFactories instances.
	//
	// testAccPreCheck(t) must be called before using this provider instance.
	testAccProvider *schema.Provider

	// providerFactories are used to instantiate a provider during acceptance
	// testing. The factory function will be invoked for every Terraform CLI
	// command executed to create a provider server to which the CLI can
	// reattach.
	providerFactories map[string]func() (*schema.Provider, error)

	// Integration test account ID.
	testAccCloudflareAccountID string = "f037e56e89293a057740de681ac9abbe"

	// Integration test account zone ID.
	testAccCloudflareZoneID string = "0da42c8d2132a9ddaf714f9e7c920711"
	// Integration test account zone name.
	testAccCloudflareZoneName string = "terraform.cfapi.net"

	// Integration test account alternate zone ID.
	testAccCloudflareAltZoneID string = "b72110c08e3382597095c29ba7e661ea"
	// Integration test account alternate zone name.
	testAccCloudflareAltZoneName string = "terraform2.cfapi.net"
)

func init() {
	testAccProvider = New("dev")()
	providerFactories = map[string]func() (*schema.Provider, error){
		"cloudflare": func() (*schema.Provider, error) {
			return New("dev")(), nil
		},
	}
}
func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = New("dev")()
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

type preCheckFunc = func(*testing.T)

func testAccPreCheck(t *testing.T) {
	testAccPreCheckEmail(t)
	testAccPreCheckApiKey(t)
	testAccPreCheckDomain(t)

	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for this acceptance test")
	}

	err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func testAccPreCheckWithoutZoneID(t *testing.T) {
	testAccPreCheckEmail(t)
	testAccPreCheckApiKey(t)
	testAccPreCheckDomain(t)

	err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func testAccPreCheckAltDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_DOMAIN must be set for this acceptance test")
	}
}

func testAccPreCheckAltZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_ZONE_ID must be set for this acceptance test")
	}
}

func testAccPreCheckAccount(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ACCOUNT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test")
	}
}

func testAccPreCheckEmail(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_EMAIL"); v == "" {
		t.Fatal("CLOUDFLARE_EMAIL must be set for acceptance tests")
	}
}

func testAccPreCheckApiKey(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_KEY"); v == "" {
		t.Fatal("CLOUDFLARE_API_KEY must be set for acceptance tests")
	}
}

func testAccPreCheckApiUserServiceKey(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_USER_SERVICE_KEY"); v == "" {
		t.Fatal("CLOUDFLARE_API_USER_SERVICE_KEY must be set for acceptance tests")
	}

	err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func testAccPreCheckDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}

func testAccPreCheckLogpushToken(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN"); v == "" {
		t.Fatal("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN must be set for this acceptance test")
	}
	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for this acceptance test")
	}
}

func testAccPreCheckWorkspaceOne(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_API_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_API_URL must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL must be set for this acceptance test")
	}
}

func testAccPreCheckPages(t *testing.T) {
	testAccPreCheckAccount(t)

	if v := os.Getenv("CLOUDFLARE_PAGES_OWNER"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_OWNER must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_PAGES_REPO"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_REPO must be set for this acceptance test")
	}
}

func testAccPreCheckBYOIPPrefix(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_BYO_IP_PREFIX_ID is not set")
	}
}

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}

// skipMagicTransitTestForNonConfiguredDefaultZone will force an acceptance test
// to skip instead of running and failing due to not having setup Magic Transit.
// This will allow those who intentionally want to run the test to do so while
// keeping CI sane.
func skipMagicTransitTestForNonConfiguredDefaultZone(t *testing.T) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == testAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test as %s is not configured for Magic Transit", testAccCloudflareZoneID)
	}
}

// skipV1WAFTestForNonConfiguredDefaultZone ignores the V1 WAF test assertions
// as the versions are mutually exclusive and the default zone ID uses V2 WAF.
// This will allow those who intentionally want to run the test to do so while
// keeping CI sane.
func skipV1WAFTestForNonConfiguredDefaultZone(t *testing.T) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == testAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test as %s is using WAF v2 and cannot assert v1 resource configurations", testAccCloudflareZoneID)
	}
}

// skipPagesProjectForNonConfiguredDefaultAccount ignores the pages project tests
// due to not having a dedicated GitHub account setup in Cloudflare for the
// default account. This will allow those who intentionally want to run the test
// to do so while keeping CI sane.
func skipPagesProjectForNonConfiguredDefaultAccount(t *testing.T) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == testAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test as %s is using pages project that isn't setup for CI", testAccCloudflareAccountID)
	}
}

func TestAccProvider_EnsureAtLeastOneCredentialDefined(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if os.Getenv("CLOUDFLARE_API_KEY") != "" {
		t.Setenv("CLOUDFLARE_API_KEY", "")
	}

	if os.Getenv("CLOUDFLARE_API_USER_SERVICE_KEY") != "" {
		t.Setenv("CLOUDFLARE_API_USER_SERVICE_KEY", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRecordConfigBasic(zoneID, rnd, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta(fmt.Sprintf("must provide exactly one of %q, %q or %q.", consts.APIKeySchemaKey, consts.APITokenSchemaKey, consts.APIUserServiceKeySchemaKey))),
			},
		},
	})
}
