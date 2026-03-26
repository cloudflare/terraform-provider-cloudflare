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
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const resourcePrefix = "tfacctest-"

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_secrets_store_secret", &resource.Sweeper{
		Name: "cloudflare_secrets_store_secret",
		F:    testSweepCloudflareSecretsStoreSecrets,
	})
}

func testSweepCloudflareSecretsStoreSecrets(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping secrets_store_secret sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	storeID := os.Getenv("CLOUDFLARE_SECRETS_STORE_ID")
	if storeID == "" {
		tflog.Info(ctx, "Skipping secrets_store_secret sweep: CLOUDFLARE_SECRETS_STORE_ID not set")
		return nil
	}

	list, err := client.SecretsStore.Stores.Secrets.List(ctx, storeID, secrets_store.StoreSecretListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list secrets: %s", err))
		return fmt.Errorf("failed to list secrets: %w", err)
	}

	hasSecrets := false
	for _, secret := range list.Result {
		hasSecrets = true
		if !utils.ShouldSweepResource(secret.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting secret: %s (account: %s)", secret.ID, accountID))
		_, err := client.SecretsStore.Stores.Secrets.Delete(ctx, storeID, secret.ID, secrets_store.StoreSecretDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete secret %s: %s", secret.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted secret: %s", secret.ID))
	}

	if !hasSecrets {
		tflog.Info(ctx, "No secrets to sweep")
	}

	return nil
}

func TestAccCloudflareSecretsStoreSecret_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_secrets_store_secret." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	storeID := os.Getenv("CLOUDFLARE_SECRETS_STORE_ID")
	secretValue := "my-secret-value-" + rnd

	if storeID == "" {
		t.Skip("CLOUDFLARE_SECRETS_STORE_ID not set, skipping test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreSecretConfig(resourceName, accountID, storeID, secretValue),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("store_id"), knownvalue.StringExact(storeID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringExact("active")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s/%s", accountID, storeID, resourceName), nil
				},
			},
		},
	})
}

func testAccCloudflareSecretsStoreSecretConfig(rnd, accountID, storeID, secretValue string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store_secret" "%s" {
  account_id = "%s"
  store_id = "%s"
  name = "%s"
  secret_text = "%s"
}
`, rnd, accountID, storeID, rnd, secretValue)
}
