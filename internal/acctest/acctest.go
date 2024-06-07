package acctest

import (
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/stainless-sdks/cloudflare-terraform/internal"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
)

var (
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

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6(internal.NewProvider("dev")())(), nil
	},
}

func TestAccPreCheck_Credentials(t *testing.T) {
	apiKey := os.Getenv(consts.APIKeyEnvVarKey)
	apiToken := os.Getenv(consts.APITokenEnvVarKey)
	userServiceKey := os.Getenv(consts.APIUserServiceKeyEnvVarKey)

	if apiToken == "" && apiKey == "" && userServiceKey == "" {
		t.Fatal("valid credentials are required for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_ZONE_ID` is present.
func TestAccPreCheck_ZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_ACCOUNT_ID` is present.
func TestAccPreCheck_AccountID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ACCOUNT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_DOMAIN` is present.
func TestAccPreCheck_Domain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}

// Test helper method checking `CLOUDFLARE_ALT_DOMAIN` is present.
func TestAccPreCheck_AlternateDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_DOMAIN must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_ALT_ZONE_ID` is present.
func TestAccPreCheck_AlternateZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ALT_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ALT_ZONE_ID must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_EMAIL` is present.
func TestAccPreCheck_Email(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_EMAIL"); v == "" {
		t.Fatal("CLOUDFLARE_EMAIL must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_API_KEY` is present.
func TestAccPreCheck_APIKey(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_KEY"); v == "" {
		t.Fatal("CLOUDFLARE_API_KEY must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_API_TOKEN` is present.
func TestAccPreCheck_APIToken(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_API_TOKEN"); v == "" {
		t.Fatal("CLOUDFLARE_API_TOKEN must be set for acceptance tests")
	}
}

// Test helper method checking `CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN` is present.
func TestAccPreCheck_LogpushToken(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN"); v == "" {
		t.Fatal("CLOUDFLARE_LOGPUSH_OWNERSHIP_TOKEN must be set for this acceptance test.")
	}
}

// Test helper method checking the Workspace One environment variables are present.
func TestAccPreCheck_WorkspaceOne(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_API_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_API_URL must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL"); v == "" {
		t.Fatal("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL must be set for this acceptance test.")
	}
}

// Test helper method checking the required environment variables for Cloudflare Pages
// are present.
func TestAccPreCheck_Pages(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_PAGES_OWNER"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_OWNER must be set for this acceptance test.")
	}

	if v := os.Getenv("CLOUDFLARE_PAGES_REPO"); v == "" {
		t.Fatal("CLOUDFLARE_PAGES_REPO must be set for this acceptance test.")
	}
}

// Test helper method checking `CLOUDFLARE_BYO_IP_PREFIX_ID` is present.
func TestAccPreCheck_BYOIPPrefix(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID"); v == "" {
		t.Skip("Skipping acceptance test as CLOUDFLARE_BYO_IP_PREFIX_ID is not set")
	}
}

// TestAccSkipForDefaultZone is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultZone(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == testAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test for default zone (%s). %s", testAccCloudflareZoneID, reason)
	}
}

// TestAccSkipForDefaultAccount is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultAccount(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == testAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test for default account (%s). %s", testAccCloudflareAccountID, reason)
	}
}

// SharedV1Client returns a common Cloudflare V1 client setup needed for the
// sweeper functions.
func SharedV1Client() (*cfv1.API, error) {
	client, err := cfv1.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL"))

	if err != nil {
		return client, err
	}

	return client, nil
}

// SharedV2Client returns a common Cloudflare V2 client setup needed for the
// sweeper functions.
func SharedV2Client() *cfv2.Client {
	return cfv2.NewClient(
		option.WithAPIKey("CLOUDFLARE_API_KEY"),
		option.WithAPIEmail("CLOUDFLARE_EMAIL"),
	)
}
