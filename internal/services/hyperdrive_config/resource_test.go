package hyperdrive_config_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
				// don't remove the one for static tests
				if q.ID == "d08bc4a1c3c140aa95e4ceec535f832e" {
					continue
				}

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
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	databaseName := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME")
	databaseHostname := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME")
	databasePort := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT")
	port, _ := strconv.Atoi(databasePort)
	databaseUser := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER")
	databasePassword := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD")

	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOrigin{
		Database: databaseName,
		Host:     databaseHostname,
		Port:     port,
		Scheme:   "postgres",
		User:     databaseUser,
	}

	var configID string
	updatedName := rnd + "-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Hyperdrive(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfig(
					rnd,
					accountID,
					rnd,
					databasePassword,
					origin,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_id"),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_secret"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "caching.max_age"),
					resource.TestCheckNoResourceAttr(resourceName, "caching.stale_while_revalidate"),
				),
			},
			// {
			// 	ResourceName:            resourceName,
			// 	ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// 	ImportStateVerifyIgnore: []string{"origin.password"},
			// },
			{
				Config: testHyperdriveConfigUpdate(
					rnd,
					configID,
					accountID,
					updatedName,
					databasePassword,
					origin,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareHyperdriveConfig_CachingSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	databaseName := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME")
	databaseHostname := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME")
	databasePort := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT")
	port, _ := strconv.Atoi(databasePort)
	databaseUser := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER")
	databasePassword := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOrigin{
		Database: databaseName,
		Host:     databaseHostname,
		Port:     port,
		Scheme:   "postgres",
		User:     databaseUser,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Hyperdrive(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigFullCachingSettings(
					rnd,
					accountID,
					rnd,
					databasePassword,
					origin,
					false,
					60,
					30,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_id"),
					resource.TestCheckNoResourceAttr(resourceName, "origin.access_client_secret"),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "caching.max_age", "60"),
					resource.TestCheckResourceAttr(resourceName, "caching.stale_while_revalidate", "30"),
				),
			},
			// {
			// 	ResourceName:            resourceName,
			// 	ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// 	ImportStateVerifyIgnore: []string{"origin.password"},
			// },
		},
	})
}

func TestAccCloudflareHyperdriveConfig_HyperdriveOverAccess(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	databaseName := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME")
	databaseHostname := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME")
	databaseUser := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER")
	databasePassword := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD")
	accessClientID := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_ID")
	accessClientSecret := os.Getenv("CLOUDFLARE_HYPERDRIVE_ACCESS_CLIENT_SECRET")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOriginWithSecrets{
		HyperdriveConfigOrigin: cfv1.HyperdriveConfigOrigin{
			Database:       databaseName,
			Host:           databaseHostname,
			Scheme:         "postgres",
			User:           databaseUser,
			AccessClientID: accessClientID,
		},
		AccessClientSecret: accessClientSecret,
	}

	var configID string
	updatedName := rnd + "-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_HyperdriveWithAccess(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveOverAccessConfig(
					rnd,
					accountID,
					rnd,
					databasePassword,
					origin,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					// use a test check function to grab the id that was generated by the API
					// func(s *terraform.State) error {
					// 	rs, ok := s.RootModule().Resources[resourceName]
					// 	if !ok {
					// 		return fmt.Errorf("not found: %s", resourceName)
					// 	}

					// 	configID = rs.Primary.Attributes["id"]
					// 	return nil
					// },
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckNoResourceAttr(resourceName, "origin.port"),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_id", accessClientID),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_secret", accessClientSecret),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
				),
			},
			// {
			// 	ResourceName:            resourceName,
			// 	ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// 	ImportStateVerifyIgnore: []string{"origin.password", "origin.access_client_secret"},
			// },
			{
				Config: testHyperdriveOverAccessUpdate(
					rnd,
					configID,
					accountID,
					updatedName,
					databasePassword,
					origin,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckNoResourceAttr(resourceName, "origin.port"),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_id", accessClientID),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_secret", accessClientSecret),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareHyperdriveConfig_Minimum(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	databaseName := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME")
	databaseHostname := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME")
	databasePort := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT")
	port, _ := strconv.Atoi(databasePort)
	databaseUser := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER")
	databasePassword := os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD")
	resourceName := "cloudflare_hyperdrive_config." + rnd

	var origin = cfv1.HyperdriveConfigOrigin{
		Database: databaseName,
		Host:     databaseHostname,
		Port:     port,
		Scheme:   "postgres",
		User:     databaseUser,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Hyperdrive(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Config: testHyperdriveConfigMinimum(
				// 	rnd,
				// 	accountID,
				// 	rnd,
				// 	databasePassword,
				// 	origin,
				// ),
				Config: testHyperdriveConfig(
					rnd,
					accountID,
					rnd,
					databasePassword,
					origin,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckResourceAttr(resourceName, "caching.disabled", "false"),
				),
			},
			// {
			// 	ResourceName:            resourceName,
			// 	ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// 	ImportStateVerifyIgnore: []string{"origin.password"},
			// },
		},
	})
}

func testHyperdriveConfig(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfig.tf",
		rnd, accountId, name, password, origin.Database, origin.Host, fmt.Sprintf("%d", origin.Port), origin.Scheme, origin.User, cacheEnabled)
}

func testHyperdriveConfigMinimum(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin) string {
	return acctest.LoadTestCase("hyperdriveconfigminimum.tf", rnd, accountId, name, password, origin.Database, origin.Host, origin.Port, origin.Scheme, origin.User)
}

func testHyperdriveConfigUpdate(rnd, id, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigupdate.tf",
		rnd,
		id,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		fmt.Sprintf("%d", origin.Port),
		origin.Scheme,
		origin.User,
		cacheEnabled,
	)
}

func testHyperdriveConfigFullCachingSettings(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool, cacheMaxAge int64, cacheStaleWhileRevalidate int64) string {
	return acctest.LoadTestCase("hyperdriveconfigfullcachesettings.tf",
		rnd,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		fmt.Sprintf("%d", origin.Port),
		origin.Scheme,
		origin.User,
		cacheEnabled,
		cacheMaxAge,
		cacheStaleWhileRevalidate,
	)
}

func testHyperdriveOverAccessConfig(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOriginWithSecrets, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigaccess.tf",
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
		cacheEnabled,
	)
}

func testHyperdriveOverAccessUpdate(rnd, id, accountId, name string, password string, origin cfv1.HyperdriveConfigOriginWithSecrets, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigoveraccessupdate.tf",
		rnd,
		id,
		accountId,
		name,
		password,
		origin.Database,
		origin.Host,
		origin.Scheme,
		origin.User,
		origin.AccessClientID,
		origin.AccessClientSecret,
		cacheEnabled,
	)
}
