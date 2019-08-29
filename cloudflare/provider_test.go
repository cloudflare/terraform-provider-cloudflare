package cloudflare

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudflare": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

type preCheckFunc = func(*testing.T)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_EMAIL"); v == "" {
		t.Fatal("CLOUDFLARE_EMAIL must be set for acceptance tests")
	}

	if v := os.Getenv("CLOUDFLARE_API_KEY"); v == "" {
		t.Fatal("CLOUDFLARE_API_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("CLOUDFLARE_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}

	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for this acceptance test")
	}
}

func testAccPreCheckAltDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_DOMAIN must be set for this acceptance test")
	}
}

func testAccPreCheckAccount(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ACCOUNT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test")
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

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}
