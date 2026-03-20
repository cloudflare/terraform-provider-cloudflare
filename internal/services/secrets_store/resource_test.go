package secrets_store_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const resourcePrefix = "tfacctest-secrets-store-"

func init() {
	resource.AddTestSweepers("cloudflare_secrets_store", &resource.Sweeper{
		Name: "cloudflare_secrets_store",
		F:    testSweepCloudflareSecretsStore,
	})
}

func testSweepCloudflareSecretsStore(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping secrets store sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.SecretsStore.Stores.List(ctx, secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list secrets stores: %s", err))
		return fmt.Errorf("failed to list secrets stores: %w", err)
	}

	if len(list.Result) == 0 {
		tflog.Info(ctx, "No secrets stores to sweep")
		return nil
	}

	for _, store := range list.Result {
		if !utils.ShouldSweepResource(store.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting secrets store: %s (account: %s)", store.ID, accountID))
		_, err := client.SecretsStore.Stores.Delete(ctx, store.ID, secrets_store.StoreDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete secrets store %s: %s", store.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted secrets store: %s", store.ID))
	}

	return nil
}

func TestAccCloudflareSecretsStore_Basic(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_secrets_store." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreConfigBasic(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "modified_at"},
			},
		},
	})
}

func TestAccCloudflareSecretsStore_NameValidation(t *testing.T) {
	t.Parallel()

	// Names must be lowercase alphanumeric and hyphens only
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid-simple", "mysecrets", false},
		{"valid-hyphen", "my-secrets-store", false},
		{"valid-numbers", "secrets123", false},
		{"invalid-uppercase", "MySecrets", true},
		{"invalid-underscore", "my_secrets", true},
		{"invalid-space", "my secrets", true},
		{"invalid-special", "my@secrets", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				// Test that invalid names are rejected
				config := fmt.Sprintf(`
resource "cloudflare_secrets_store" "test" {
  account_id = "%s"
  name       = "%s"
}
`, os.Getenv("CLOUDFLARE_ACCOUNT_ID"), tc.input)

				resource.Test(t, resource.TestCase{
					PreCheck:                 func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Steps: []resource.TestStep{
						{
							Config:      config,
							ExpectError: regexp.MustCompile("Invalid name"),
						},
					},
				})
			}
		})
	}
}

func testAccCloudflareSecretsStoreConfigBasic(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "%s" {
  account_id = "%s"
  name       = "%s"
}
`, resourceName, accountID, resourceName)
}
