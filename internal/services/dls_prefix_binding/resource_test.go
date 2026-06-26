package dls_prefix_binding_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/dls"
	"github.com/cloudflare/cloudflare-go/v7/option"
	cloudflareprovider "github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	envAPIToken             = "CLOUDFLARE_API_TOKEN"
	envAPIKey               = "CLOUDFLARE_API_KEY"
	envAPIUserServiceKey    = "CLOUDFLARE_API_USER_SERVICE_KEY"
	envAccountID            = "CLOUDFLARE_ACCOUNT_ID"
	envDLSBYOIPPrefixID     = "CLOUDFLARE_DLS_BYOIP_PREFIX_ID"
	envDLSPrefixBindingCIDR = "CLOUDFLARE_DLS_PREFIX_BINDING_CIDR"
	envDLSRegionKey         = "CLOUDFLARE_DLS_REGION_KEY"
	envDLSAltRegionKey      = "CLOUDFLARE_DLS_ALT_REGION_KEY"

	defaultDLSRegionKey    = "eu"
	defaultDLSAltRegionKey = "isoeu"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6(cloudflareprovider.NewProvider("dev")())(), nil
	},
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_dls_prefix_binding", &resource.Sweeper{
		Name: "cloudflare_dls_prefix_binding",
		F:    testSweepCloudflareDLSPrefixBindings,
	})
}

func testSweepCloudflareDLSPrefixBindings(_ string) error {
	return testAccCleanupCloudflareDLSPrefixBinding(context.Background())
}

func TestAccCloudflareDLSPrefixBinding(t *testing.T) {
	testAccPreCheckDLSPrefixBinding(t)

	accountID := os.Getenv(envAccountID)
	prefixID := os.Getenv(envDLSBYOIPPrefixID)
	cidr := testAccDLSPrefixBindingCIDR(t)
	regionKey := testAccDLSRegionKeyFromEnv()
	altRegionKey := testAccDLSAltRegionKeyFromEnv()

	resourceName := "cloudflare_dls_prefix_binding.test"
	dataSourceName := "data.cloudflare_dls_prefix_binding.test"
	listDataSourceName := "data.cloudflare_dls_prefix_bindings.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckCredentials(t)
			testAccPreCheckDLSPrefixBinding(t)
			if err := testAccCleanupCloudflareDLSPrefixBinding(context.Background()); err != nil {
				t.Fatalf("failed to clean up pre-existing DLS prefix binding fixture: %s", err)
			}
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDLSPrefixBindingConfig(accountID, prefixID, cidr, regionKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "prefix_id", prefixID),
					resource.TestCheckResourceAttr(resourceName, "cidr", cidr),
					resource.TestCheckResourceAttr(resourceName, "region_key", regionKey),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "prefix_id", prefixID),
					resource.TestCheckResourceAttr(dataSourceName, "cidr", cidr),
					resource.TestCheckResourceAttr(dataSourceName, "region_key", regionKey),
					resource.TestCheckResourceAttrWith(listDataSourceName, "result.#", func(value string) error {
						count, err := strconv.Atoi(value)
						if err != nil {
							return fmt.Errorf("failed to parse list result count %q: %w", value, err)
						}
						if count < 1 {
							return fmt.Errorf("expected at least one DLS prefix binding, got %d", count)
						}
						return nil
					}),
				),
			},
			{
				Config: testAccCloudflareDLSPrefixBindingConfig(accountID, prefixID, cidr, altRegionKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "prefix_id", prefixID),
					resource.TestCheckResourceAttr(resourceName, "cidr", cidr),
					resource.TestCheckResourceAttr(resourceName, "region_key", altRegionKey),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttr(dataSourceName, "prefix_id", prefixID),
					resource.TestCheckResourceAttr(dataSourceName, "cidr", cidr),
					resource.TestCheckResourceAttr(dataSourceName, "region_key", altRegionKey),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccCloudflareDLSPrefixBindingConfig(accountID, prefixID, cidr, regionKey string) string {
	return loadTestCase("dls_prefix_binding.tf", "test", accountID, prefixID, cidr, regionKey)
}

func testAccCleanupCloudflareDLSPrefixBinding(ctx context.Context) error {
	accountID := os.Getenv(envAccountID)
	prefixID := os.Getenv(envDLSBYOIPPrefixID)
	cidr := testAccDLSPrefixBindingCIDRFromEnv()
	if accountID == "" || prefixID == "" || cidr == "" {
		log.Printf("Skipping DLS prefix binding cleanup: %s, %s, or binding CIDR env var is not set", envAccountID, envDLSBYOIPPrefixID)
		return nil
	}

	client := cloudflare.NewClient(testAccClientOptions()...)
	page, err := client.DLS.RegionalServices.PrefixBindings.List(ctx, dls.RegionalServicePrefixBindingListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list DLS prefix bindings: %w", err)
	}

	for page != nil && len(page.Result) > 0 {
		for _, binding := range page.Result {
			if binding.PrefixID != prefixID || binding.CIDR != cidr {
				continue
			}

			log.Printf("Deleting stale DLS prefix binding %s for prefix %s CIDR %s", binding.ID, prefixID, cidr)
			_, err := client.DLS.RegionalServices.PrefixBindings.Delete(ctx, binding.ID, dls.RegionalServicePrefixBindingDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				return fmt.Errorf("failed to delete stale DLS prefix binding %s: %w", binding.ID, err)
			}
		}

		page, err = page.GetNextPage()
		if err != nil {
			return fmt.Errorf("failed to fetch next DLS prefix bindings page: %w", err)
		}
	}

	return nil
}

func testAccPreCheckCredentials(t *testing.T) {
	t.Helper()

	if os.Getenv(envAPIToken) == "" && os.Getenv(envAPIKey) == "" && os.Getenv(envAPIUserServiceKey) == "" {
		t.Fatalf("valid credentials are required for this acceptance test: one of %s, %s, or %s must be set", envAPIToken, envAPIKey, envAPIUserServiceKey)
	}
}

func testAccPreCheckDLSPrefixBinding(t *testing.T) {
	t.Helper()

	for _, envVar := range []string{envAccountID, envDLSBYOIPPrefixID} {
		if os.Getenv(envVar) == "" {
			t.Skipf("%s must be set for this acceptance test", envVar)
		}
	}

	if os.Getenv(envDLSPrefixBindingCIDR) == "" {
		t.Skipf("%s must be set to a reserved CIDR within %s for this acceptance test", envDLSPrefixBindingCIDR, envDLSBYOIPPrefixID)
	}

	if testAccDLSRegionKeyFromEnv() == testAccDLSAltRegionKeyFromEnv() {
		t.Skipf("DLS create and update region keys must be different for the update acceptance test")
	}
}

func testAccDLSPrefixBindingCIDR(t *testing.T) string {
	t.Helper()

	return testAccDLSPrefixBindingCIDRFromEnv()
}

func testAccDLSRegionKeyFromEnv() string {
	if regionKey := os.Getenv(envDLSRegionKey); regionKey != "" {
		return regionKey
	}
	return defaultDLSRegionKey
}

func testAccDLSAltRegionKeyFromEnv() string {
	if regionKey := os.Getenv(envDLSAltRegionKey); regionKey != "" {
		return regionKey
	}
	return defaultDLSAltRegionKey
}

func testAccDLSPrefixBindingCIDRFromEnv() string {
	return os.Getenv(envDLSPrefixBindingCIDR)
}

func testAccClientOptions() []option.RequestOption {
	opts := []option.RequestOption{}
	if value := os.Getenv("CLOUDFLARE_BASE_URL"); value != "" {
		opts = append(opts, option.WithBaseURL(value))
	}
	if value := os.Getenv(envAPIToken); value != "" {
		opts = append(opts, option.WithAPIToken(value))
	}
	if value := os.Getenv(envAPIKey); value != "" {
		opts = append(opts, option.WithAPIKey(value))
	}
	if value := os.Getenv("CLOUDFLARE_EMAIL"); value != "" {
		opts = append(opts, option.WithAPIEmail(value))
	}
	if value := os.Getenv(envAPIUserServiceKey); value != "" {
		opts = append(opts, option.WithUserServiceKey(value))
	}
	return opts
}

func loadTestCase(filename string, parameters ...interface{}) string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	contents, err := os.ReadFile(filepath.Join(pwd, "testdata", filename))
	if err != nil {
		return ""
	}

	return fmt.Sprintf(string(contents), parameters...)
}
