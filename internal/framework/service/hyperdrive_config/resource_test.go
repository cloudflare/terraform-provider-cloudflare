package hyperdrive_config_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
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
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			ctx := context.Background()

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", err))
			}

			resp, err := client.ListHyperdriveConfigs(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListHyperdriveConfigParams{})
			if err != nil {
				return err
			}

			for _, q := range resp {
				err := client.DeleteHyperdriveConfig(ctx, cfv1.AccountIdentifier(accountID), q.ID)
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

	var origin = cfv1.HyperdriveConfigOrigin{
		Database: "database",
		Host:     "host.example.com",
		Port:     5432,
		Scheme:   "postgres",
		User:     "user",
	}

	var disabled = true

	var caching = cfv1.HyperdriveConfigCaching{
		Disabled: &disabled,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfig(
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
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_id"),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_secret"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "caching.max_age"),
					resource.TestCheckNoResourceAttr(resourceName, "caching.stale_while_revalidate"),
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

func TestAccCloudflareHyperdriveConfig_CachingSettings(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires real Postgres instance to be available.")

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOrigin{
		Database: "database",
		Host:     "host.example.com",
		Port:     5432,
		Scheme:   "postgres",
		User:     "user",
	}

	var disabled = false

	var caching = cfv1.HyperdriveConfigCaching{
		Disabled:             &disabled,
		MaxAge:               60,
		StaleWhileRevalidate: 30,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigFullCachingSettings(
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
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_id"),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_secret"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "caching.max_age", "60"),
					resource.TestCheckResourceAttr(resourceName, "caching.stale_while_revalidate", "30"),
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

func TestAccCloudflareHyperdriveConfig_HyperdriveOverAccess(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOriginWithSecrets{
		HyperdriveConfigOrigin: cfv1.HyperdriveConfigOrigin{Database: "database",
			Host:           "host.example.com",
			Scheme:         "postgres",
			User:           "user",
			AccessClientID: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.access"},
		AccessClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	var disabled = true

	var caching = cfv1.HyperdriveConfigCaching{
		Disabled: &disabled,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveOverAccessConfig(
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
					resource.TestCheckNoResourceAttr(resourceName, "origin.port"),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", "user"),
					resource.TestCheckResourceAttr(resourceName, "origin.password", "password"),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_id", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.access"),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_secret", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"origin.password", "origin.access_client_secret"},
			},
		},
	})
}

func TestAccCloudflareHyperdriveConfig_Minimum(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires real Postgres instance to be available.")

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOrigin{
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

func testHyperdriveConfig(
	rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, caching cfv1.HyperdriveConfigCaching,
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
		rnd,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		fmt.Sprintf("%d", origin.Port),
		origin.Scheme,
		origin.User,
		fmt.Sprintf("%t", *caching.Disabled),
	)
}

func testHyperdriveConfigFullCachingSettings(
	rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, caching cfv1.HyperdriveConfigCaching,
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
				max_age                = %[11]s
				stale_while_revalidate = %[12]s
			}
		}`,
		rnd,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		fmt.Sprintf("%d", origin.Port),
		origin.Scheme,
		origin.User,
		fmt.Sprintf("%t", *caching.Disabled),
		fmt.Sprintf("%d", caching.MaxAge),
		fmt.Sprintf("%d", caching.StaleWhileRevalidate),
	)
}

func testHyperdriveOverAccessConfig(
	rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOriginWithSecrets, caching cfv1.HyperdriveConfigCaching,
) string {
	return fmt.Sprintf(`
		resource "cloudflare_hyperdrive_config" "%[1]s" {
			account_id = "%[2]s"
			name       = "%[3]s"
			origin     = {
				password             = "%[4]s"
				database             = "%[5]s"
				host                 = "%[6]s"
				scheme               = "%[7]s"
				user                 = "%[8]s"
				access_client_id     = "%[9]s"
				access_client_secret = "%[10]s"
			}
			caching = {
				disabled               = %[11]s
			}
		}`,
		rnd,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		origin.Scheme,
		origin.User,
		origin.AccessClientID,
		origin.AccessClientSecret,
		fmt.Sprintf("%t", *caching.Disabled),
	)
}

func testHyperdriveConfigConfigMinimum(
	rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin,
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
		rnd,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		fmt.Sprintf("%d", origin.Port),
		origin.Scheme,
		origin.User,
	)
}
