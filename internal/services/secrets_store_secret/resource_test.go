package secrets_store_secret_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_secrets_store_secret", &resource.Sweeper{
		Name: "cloudflare_secrets_store_secret",
		F:    testSweepSecretsStoreSecrets,
	})
}

// testSweepSecretsStoreSecrets removes test-created secrets (names starting with
// "cftftest") from all stores in the account. It does not delete the stores
// themselves or any non-test secrets.
func testSweepSecretsStoreSecrets(_ string) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return fmt.Errorf("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	client := acctest.SharedClient()
	ctx := context.Background()

	stores, err := client.SecretsStore.Stores.List(ctx, secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("listing secrets stores: %w", err)
	}

	for _, store := range stores.Result {
		secrets, err := client.SecretsStore.Stores.Secrets.List(ctx, store.ID, secrets_store.StoreSecretListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			return fmt.Errorf("listing secrets in store %s: %w", store.ID, err)
		}
		for _, secret := range secrets.Result {
			// Only delete secrets created by acceptance tests.
			if len(secret.Name) >= 8 && secret.Name[:8] == "cftftest" {
				_, err := client.SecretsStore.Stores.Secrets.Delete(ctx, store.ID, secret.ID, secrets_store.StoreSecretDeleteParams{
					AccountID: cloudflare.F(accountID),
				})
				if err != nil {
					return fmt.Errorf("deleting secret %s: %w", secret.ID, err)
				}
			}
		}
	}

	return nil
}

// getExistingStoreID returns the ID of the first store in the account.
// Secrets Store currently has a one-store-per-account limit.
func getExistingStoreID(t *testing.T) string {
	t.Helper()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("skipping: CLOUDFLARE_ACCOUNT_ID must be set")
	}
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
	return stores.Result[0].ID
}

func TestAccCloudflareSecretsStoreSecret_Workflow(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	storeID := getExistingStoreID(t)
	secretResourceName := "cloudflare_secrets_store_secret." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecretsStoreSecretDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create secret in existing store
				Config: testAccSecretsStoreSecretCreate(rnd, accountID, storeID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("store_id"), knownvalue.StringExact(storeID)),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("initial comment")),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
				},
			},
			{
				// Step 2: Update value via PATCH
				Config: testAccSecretsStoreSecretUpdateValue(rnd, accountID, storeID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(secretResourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("value"), knownvalue.StringExact("updated-secret-value")),
				},
			},
			{
				// Step 3: Update comment via PATCH
				Config: testAccSecretsStoreSecretUpdateComment(rnd, accountID, storeID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(secretResourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("updated comment")),
				},
			},
			{
				// Step 4: Update scopes via PATCH
				// NOTE: scopes must be listed in alphabetical order in the config.
				// The API returns scopes sorted alphabetically regardless of input
				// order. Since scopes is a ListAttribute (ordered), any mismatch
				// between config order and API order causes perpetual drift.
				// A future improvement would be to use SetAttribute for scopes.
				Config: testAccSecretsStoreSecretUpdateScopes(rnd, accountID, storeID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(secretResourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(secretResourceName, tfjsonpath.New("scopes"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("ai_gateway"),
							knownvalue.StringExact("workers"),
						}),
					),
				},
			},
			{
				// Step 5: Import
				ResourceName: secretResourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					secretRS, ok := s.RootModule().Resources[secretResourceName]
					if !ok {
						return "", fmt.Errorf("resource %s not found", secretResourceName)
					}
					return fmt.Sprintf("%s/%s/%s", accountID, storeID, secretRS.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"value", "modified", "status"},
			},
		},
	})
}

func testAccSecretsStoreSecretCreate(rnd, accountID, storeID string) string {
	return acctest.LoadTestCase("secretsstoresecretcreate.tf", rnd, accountID, storeID)
}

func testAccSecretsStoreSecretUpdateValue(rnd, accountID, storeID string) string {
	return acctest.LoadTestCase("secretsstoresecretupdatevalue.tf", rnd, accountID, storeID)
}

func testAccSecretsStoreSecretUpdateComment(rnd, accountID, storeID string) string {
	return acctest.LoadTestCase("secretsstoresecretupdatecomment.tf", rnd, accountID, storeID)
}

func testAccSecretsStoreSecretUpdateScopes(rnd, accountID, storeID string) string {
	return acctest.LoadTestCase("secretsstoresecretupdatescopes.tf", rnd, accountID, storeID)
}

func testAccCheckSecretsStoreSecretDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_secrets_store_secret" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		storeID := rs.Primary.Attributes["store_id"]
		_, err := client.SecretsStore.Stores.Secrets.Get(
			context.Background(),
			storeID,
			rs.Primary.ID,
			secrets_store.StoreSecretGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("secrets store secret %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
