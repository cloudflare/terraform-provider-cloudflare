package secrets_store_test

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
	resource.AddTestSweepers("cloudflare_secrets_store", &resource.Sweeper{
		Name: "cloudflare_secrets_store",
		F:    testSweepCloudflareSecretsStores,
	})
}

func testSweepCloudflareSecretsStores(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping secrets_store sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.SecretsStore.Stores.List(ctx, secrets_store.StoreListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list secrets stores: %s", err))
		return fmt.Errorf("failed to list secrets stores: %w", err)
	}

	hasStores := false
	for _, store := range list.Result {
		hasStores = true
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

	if !hasStores {
		tflog.Info(ctx, "No secrets stores to sweep")
	}

	return nil
}

func TestAccCloudflareSecretsStore_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_secrets_store." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSecretsStoreConfig(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, resourceName), nil
				},
			},
		},
	})
}

func testAccCloudflareSecretsStoreConfig(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_secrets_store" "%s" {
  account_id = "%s"
  name = "%s"
}
`, rnd, accountID, rnd)
}
