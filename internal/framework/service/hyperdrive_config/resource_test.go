package hyperdrive_config_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_hyperdrive_config", &resource.Sweeper{
		Name: "cloudflare_hyperdrive_config",
		F: func(region string) error {
			client, err := acctest.SharedClient()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			ctx := context.Background()

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", err))
			}

			resp, err := client.ListHyperdriveConfigs(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListHyperdriveConfigParams{})
			if err != nil {
				return err
			}

			for _, q := range resp {
				err := client.DeleteHyperdriveConfig(ctx, cloudflare.AccountIdentifier(accountID), q.ID)
				if err != nil {
					return err
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareHyperdriveConfig_Basic(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires real Postgres instance to be available.")

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cloudflare.HyperdriveConfigOrigin{
		Database: "database",
		Host:     "host.example.com",
		Port:     5432,
		Scheme:   "postgres",
		User:     "user",
	}

	var disabled = true

	var caching = cloudflare.HyperdriveConfigCaching{
		Disabled: &disabled,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigConfig(
					rnd,
					accountID,
					rnd,
					"password",
					origin,
					caching,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", "database"),
					resource.TestCheckResourceAttr(resourceName, "origin.host", "host.example.com"),
					resource.TestCheckResourceAttr(resourceName, "origin.port", "5432"),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", "user"),
					resource.TestCheckResourceAttr(resourceName, "origin.password", "password"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"origin.password"},
			},
		},
	})
}

func TestAccCloudflareHyperdriveConfig_Minimum(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires real Postgres instance to be available.")

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cloudflare.HyperdriveConfigOrigin{
		Database: "database",
		Host:     "host.example.com",
		Port:     5432,
		Scheme:   "postgres",
		User:     "user",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigConfigMinimum(
					rnd,
					accountID,
					rnd,
					"password",
					origin,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", "database"),
					resource.TestCheckResourceAttr(resourceName, "origin.host", "host.example.com"),
					resource.TestCheckResourceAttr(resourceName, "origin.port", "5432"),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", "user"),
					resource.TestCheckResourceAttr(resourceName, "origin.password", "password"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "false"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"origin.password"},
			},
		},
	})
}

func testHyperdriveConfigConfig(
	rnd, accountId, name string, password string, origin cloudflare.HyperdriveConfigOrigin, caching cloudflare.HyperdriveConfigCaching,
) string {
	return fmt.Sprintf(`
		resource "cloudflare_hyperdrive_config" "%[1]s" {
			account_id = "%[2]s"
			name       = "%[3]s"
			origin     = {
				password = "%[4]s"
				database = "%[5]s"
				host     = "%[6]s"
				port     = "%[7]s"
				scheme   = "%[8]s"
				user     = "%[9]s"
			}
			caching = {
				disabled               = %[10]s
			}
		}`,
		rnd, accountId, name, password, origin.Database, origin.Host, fmt.Sprintf("%d", origin.Port), origin.Scheme, origin.User, fmt.Sprintf("%t", *caching.Disabled),
	)
}

func testHyperdriveConfigConfigMinimum(
	rnd, accountId, name string, password string, origin cloudflare.HyperdriveConfigOrigin,
) string {
	return fmt.Sprintf(`
		resource "cloudflare_hyperdrive_config" "%[1]s" {
			account_id = "%[2]s"
			name       = "%[3]s"
			origin     = {
				password   = "%[4]s"
				database   = "%[5]s"
				host       = "%[6]s"
				port       = "%[7]s"
				scheme	   = "%[8]s"
				user       = "%[9]s"
			}
		}`,
		rnd, accountId, name, password, origin.Database, origin.Host, fmt.Sprintf("%d", origin.Port), origin.Scheme, origin.User,
	)
}
