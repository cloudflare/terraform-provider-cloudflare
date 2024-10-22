package acctest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// Integration test account ID.
	TestAccCloudflareAccountID string = "f037e56e89293a057740de681ac9abbe"

	// Integration test account zone ID.
	TestAccCloudflareZoneID string = "0da42c8d2132a9ddaf714f9e7c920711"
	// Integration test account zone name.
	TestAccCloudflareZoneName string = "terraform.cfapi.net"

	// Integration test account alternate zone ID.
	TestAccCloudflareAltZoneID string = "b72110c08e3382597095c29ba7e661ea"
	// Integration test account alternate zone name.
	TestAccCloudflareAltZoneName string = "terraform2.cfapi.net"
)

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6(internal.NewProvider("dev")())(), nil
	},
}

func TestAccPreCheck(t *testing.T) {
	TestAccPreCheck_Credentials(t)
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

// Test helper method checking all required Hyperdrive configurations are present.
func TestAccPreCheck_Hyperdrive(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_USER must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD must be set for this acceptance test")
	}
}

func TestAccPreCheck_HyperdriveWithAccess(t *testing.T) {
	TestAccPreCheck_Hyperdrive(t)

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_ID"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_ID must be set for this acceptance test")
	}

	if v := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_SECRET"); v == "" {
		t.Fatal("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_SECRET must be set for this acceptance test")
	}
}

// TestAccSkipForDefaultZone is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultZone(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == TestAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test for default zone (%s). %s", TestAccCloudflareZoneID, reason)
	}
}

// TestAccSkipForDefaultAccount is used for skipping over tests that are not run by
// default on usual acceptance test suite account.
func TestAccSkipForDefaultAccount(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == TestAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test for default account (%s). %s", TestAccCloudflareAccountID, reason)
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

// SharedClient returns a common Cloudflare V2 client setup needed for the
// sweeper functions.
func SharedClient() *cloudflare.Client {
	return cloudflare.NewClient(
		option.WithAPIKey(os.Getenv("CLOUDFLARE_API_KEY")),
		option.WithAPIEmail(os.Getenv("CLOUDFLARE_EMAIL")),
	)
}

// LoadTestCase takes a filename and variadic parameters to build test case output.
//
// Example: If you have a "basic" test case that for `r2_bucket` resource, inside
// of the `testdata` directory, you need to have a `basic.tf` file.
//
//	  $ tree internal/services/r2_bucket/
//	  ├── ...
//	  ├── schema.go
//	  └── testdata
//		  └── basic.tf
//
// The invocation would be `LoadTestCase("basic.tf", rnd, acountID)`.
func LoadTestCase(filename string, parameters ...interface{}) string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	fullPath := filepath.Join(pwd, "testdata", filename)
	f, err := os.ReadFile(fullPath)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(string(f), parameters...)
}
