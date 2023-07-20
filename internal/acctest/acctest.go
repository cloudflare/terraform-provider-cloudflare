package acctest

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/provider"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/sdkv2provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
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
		upgradedSdkProvider, err := tf5to6server.UpgradeServer(context.Background(), sdkv2provider.New("dev")().GRPCProvider)
		if err != nil {
			log.Fatal(err)
		}
		providers := []func() tfprotov6.ProviderServer{
			func() tfprotov6.ProviderServer {
				return upgradedSdkProvider
			},
			providerserver.NewProtocol6(provider.New("dev")()),
		}

		return tf6muxserver.NewMuxServer(context.Background(), providers...)
	},
}

func TestAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestAccPreCheck_Zone(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestAccPreCheck_Account(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestAccSkipForDefaultZone(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ZONE_ID") == testAccCloudflareZoneID {
		t.Skipf("Skipping acceptance test for default zone (%s). %s", testAccCloudflareZoneID, reason)
	}
}

func TestAccSkipForDefaultAccount(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == testAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test for default account (%s). %s", testAccCloudflareAccountID, reason)
	}
}

// sharedClient returns a common Cloudflare client setup needed for the
// sweeper functions.
func SharedClient() (*cloudflare.API, error) {
	client, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL"))

	if err != nil {
		return client, err
	}

	return client, nil
}
