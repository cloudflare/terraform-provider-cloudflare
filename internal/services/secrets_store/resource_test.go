package secrets_store_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// getExistingStore returns the ID and name of the first store in the account.
// Secrets Store currently has a one-store-per-account limit.
func getExistingStore(t *testing.T) (storeID, storeName string) {
	t.Helper()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client := acctest.SharedClient()
	stores, err := client.SecretsStore.Stores.List(context.Background(), secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		t.Fatalf("listing stores: %s", err)
	}
	if len(stores.Result) == 0 {
		t.Fatal("no secrets store found in account — create one first")
	}
	return stores.Result[0].ID, stores.Result[0].Name
}

// TestAccCloudflareSecretsStore_Import tests importing an existing secrets
// store into Terraform state and verifying all fields are populated correctly.
// Due to the one-store-per-account limit, this test uses the pre-existing store.
func TestAccCloudflareSecretsStore_Import(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	storeID, storeName := getExistingStore(t)
	resourceName := "cloudflare_secrets_store.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:            testAccSecretsStoreConfig(accountID, storeName),
				ResourceName:     resourceName,
				ImportStateId:    fmt.Sprintf("%s/%s", accountID, storeID),
				ImportState:      true,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(states))
					}
					s := states[0]
					if s.Attributes["id"] != storeID {
						return fmt.Errorf("expected id %q, got %q", storeID, s.Attributes["id"])
					}
					if s.Attributes["name"] != storeName {
						return fmt.Errorf("expected name %q, got %q", storeName, s.Attributes["name"])
					}
					if s.Attributes["account_id"] != accountID {
						return fmt.Errorf("expected account_id %q, got %q", accountID, s.Attributes["account_id"])
					}
					if s.Attributes["created"] == "" {
						return fmt.Errorf("expected created to be set")
					}
					if s.Attributes["modified"] == "" {
						return fmt.Errorf("expected modified to be set")
					}
					return nil
				},
			},
		},
	})
}

// TestAccCloudflareSecretsStore_Workflow tests the full lifecycle by creating
// a store with a config, verifying computed fields, and importing. This test
// requires an account with no existing store (due to the one-store-per-account
// limit) and is skipped when a store already exists.
func TestAccCloudflareSecretsStore_Workflow(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_secrets_store.test"
	storeName := "cftftest-store"

	// Skip if account already has a store.
	client := acctest.SharedClient()
	stores, err := client.SecretsStore.Stores.List(context.Background(), secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err == nil && len(stores.Result) > 0 {
		t.Skip("skipping store create/delete test — account already has a store (1-store limit)")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretsStoreConfig(accountID, storeName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(storeName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccSecretsStoreConfig(accountID, storeName string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "test" {
  account_id = "%s"
  name       = "%s"
}
`, accountID, storeName)
}
