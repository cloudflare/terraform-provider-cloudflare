package secrets_store_secret_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const secretResourcePrefix = "tfacctest-secrets-store-secret-"

func init() {
	resource.AddTestSweepers("cloudflare_secrets_store_secret", &resource.Sweeper{
		Name: "cloudflare_secrets_store_secret",
		F:    testSweepCloudflareSecretsStoreSecret,
	})
}

func testSweepCloudflareSecretsStoreSecret(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping secrets store secret sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.SecretsStore.Stores.List(ctx, secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list secrets stores: %s", err))
		return fmt.Errorf("failed to list secrets stores: %w", err)
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

func TestAccCloudflareSecretsStoreSecret_Basic(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	storeResourceName := secretResourcePrefix + "store-" + rnd
	secretResourceName := secretResourcePrefix + "secret-" + rnd
	storeName := "cloudflare_secrets_store." + storeResourceName
	secretName := "cloudflare_secrets_store_secret." + secretResourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreSecretConfigBasic(storeResourceName, secretResourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(secretName, "store_id", storeName, "id"),
					resource.TestCheckResourceAttr(secretName, "name", "TEST_SECRET"),
					resource.TestCheckResourceAttr(secretName, "secret_text", "secret-value-123"),
					resource.TestCheckResourceAttrSet(secretName, "id"),
					resource.TestCheckResourceAttrSet(secretName, "status"),
				),
			},
			{
				ResourceName:            secretName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret_text", "created_at", "modified_at"},
			},
		},
	})
}

func TestAccCloudflareSecretsStoreSecret_WithScopes(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	storeResourceName := secretResourcePrefix + "store-" + rnd
	secretResourceName := secretResourcePrefix + "secret-" + rnd
	storeName := "cloudflare_secrets_store." + storeResourceName
	secretName := "cloudflare_secrets_store_secret." + secretResourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreSecretConfigWithScopes(storeResourceName, secretResourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(secretName, "store_id", storeName, "id"),
					resource.TestCheckResourceAttr(secretName, "name", "SCOPED_SECRET"),
					resource.TestCheckResourceAttr(secretName, "secret_text", "scoped-secret-value"),
					resource.TestCheckResourceAttr(secretName, "comment", "Test secret with scopes"),
					resource.TestCheckResourceAttrSet(secretName, "id"),
				),
			},
		},
	})
}

func TestAccCloudflareSecretsStoreSecret_UpdateSecretText(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	storeResourceName := secretResourcePrefix + "store-" + rnd
	secretResourceName := secretResourcePrefix + "secret-" + rnd
	storeName := "cloudflare_secrets_store." + storeResourceName
	secretName := "cloudflare_secrets_store_secret." + secretResourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreSecretConfigBasic(storeResourceName, secretResourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(secretName, "store_id", storeName, "id"),
					resource.TestCheckResourceAttr(secretName, "secret_text", "secret-value-123"),
				),
			},
			{
				Config: testAccCloudflareSecretsStoreSecretConfigUpdated(storeResourceName, secretResourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(secretName, "store_id", storeName, "id"),
					resource.TestCheckResourceAttr(secretName, "secret_text", "updated-secret-value"),
				),
			},
		},
	})
}

func testAccCloudflareSecretsStoreSecretConfigBasic(storeResourceName, secretResourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "%s" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_secrets_store_secret" "%s" {
  account_id  = "%s"
  store_id   = cloudflare_secrets_store.%s.id
  name       = "TEST_SECRET"
  secret_text = "secret-value-123"
}
`, storeResourceName, accountID, storeResourceName,
		secretResourceName, accountID, storeResourceName)
}

func testAccCloudflareSecretsStoreSecretConfigWithScopes(storeResourceName, secretResourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "%s" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_secrets_store_secret" "%s" {
  account_id  = "%s"
  store_id   = cloudflare_secrets_store.%s.id
  name       = "SCOPED_SECRET"
  secret_text = "scoped-secret-value"
  comment    = "Test secret with scopes"
  scopes     = ["workers", "ai_gateway"]
}
`, storeResourceName, accountID, storeResourceName,
		secretResourceName, accountID, storeResourceName)
}

func testAccCloudflareSecretsStoreSecretConfigUpdated(storeResourceName, secretResourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "%s" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_secrets_store_secret" "%s" {
  account_id  = "%s"
  store_id   = cloudflare_secrets_store.%s.id
  name       = "TEST_SECRET"
  secret_text = "updated-secret-value"
}
`, storeResourceName, accountID, storeResourceName,
		secretResourceName, accountID, storeResourceName)
}
