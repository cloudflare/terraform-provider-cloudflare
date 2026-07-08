package hyperdrive_config_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/hyperdrive"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_hyperdrive_config", &resource.Sweeper{
		Name: "cloudflare_hyperdrive_config",
		F: func(region string) error {
			ctx := context.Background()
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", err))
				return fmt.Errorf("error establishing client: %w", err)
			}

			if accountID == "" {
				tflog.Info(ctx, "Skipping Hyperdrive configs sweep: CLOUDFLARE_ACCOUNT_ID not set")
				return nil
			}

			resp, err := client.ListHyperdriveConfigs(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListHyperdriveConfigParams{})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch Hyperdrive configs: %s", err))
				return fmt.Errorf("failed to fetch Hyperdrive configs: %w", err)
			}

			if len(resp) == 0 {
				tflog.Info(ctx, "No Hyperdrive configs to sweep")
				return nil
			}

			for _, q := range resp {
				// don't remove the one for static tests
				if q.ID == "d08bc4a1c3c140aa95e4ceec535f832e" {
					continue
				}

				if !utils.ShouldSweepResource(q.Name) {
					continue
				}

				tflog.Info(ctx, fmt.Sprintf("Deleting Hyperdrive config: %s (%s) (account: %s)", q.Name, q.ID, accountID))
				err := client.DeleteHyperdriveConfig(ctx, cfv1.AccountIdentifier(accountID), q.ID)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete Hyperdrive config %s: %s", q.ID, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted Hyperdrive config: %s", q.ID))
			}

			return nil
		},
	})
}

func testAccCheckCloudflareHyperdriveConfigDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_hyperdrive_config" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.Hyperdrive.Configs.Get(
			context.Background(),
			rs.Primary.ID,
			hyperdrive.ConfigGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("hyperdrive config %s still exists", rs.Primary.ID)
		}
	}

	return nil
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

	updatedName := rnd + "-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Hyperdrive(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
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
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"origin.password"},
			},
			{
				Config: testHyperdriveConfigUpdate(
					rnd,
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
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
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
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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

	updatedName := rnd + "-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_HyperdriveWithAccess(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
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
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"origin.password", "origin.access_client_secret"},
			},
			{
				Config: testHyperdriveOverAccessUpdate(
					rnd,
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
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigMinimum(
					rnd,
					accountID,
					rnd,
					databasePassword,
					origin,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
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

func TestAccCloudflareHyperdriveConfig_ConnectionLimit(t *testing.T) {
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
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testHyperdriveConfigConnectionLimit(
					rnd,
					accountID,
					rnd,
					10,
					databasePassword,
					origin,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "origin_connection_limit", "10"),
					resource.TestCheckResourceAttr(resourceName, "origin.database", databaseName),
					resource.TestCheckResourceAttr(resourceName, "origin.host", databaseHostname),
					resource.TestCheckResourceAttr(resourceName, "origin.port", databasePort),
					resource.TestCheckResourceAttr(resourceName, "origin.scheme", "postgres"),
					resource.TestCheckResourceAttr(resourceName, "origin.user", databaseUser),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
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

// TestAccCloudflareHyperdriveConfig_NoDiffOnConsecutiveApply tests that applying the same
// configuration twice does not result in any changes being detected.
// This is a regression test for https://github.com/cloudflare/terraform-provider-cloudflare/issues/6650
func TestAccCloudflareHyperdriveConfig_NoDiffOnConsecutiveApply(t *testing.T) {
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

	config := testHyperdriveConfig(
		rnd,
		accountID,
		rnd,
		databasePassword,
		origin,
		false,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Hyperdrive(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
		Steps: []resource.TestStep{
			{
				// Initial creation
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
				),
			},
			{
				// Same config - should not produce any changes
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccCloudflareHyperdriveConfig_NoDiffOnConsecutiveApplyWithAccess tests that applying
// the same configuration with access credentials twice does not result in any changes.
// This is a regression test for https://github.com/cloudflare/terraform-provider-cloudflare/issues/6650
func TestAccCloudflareHyperdriveConfig_NoDiffOnConsecutiveApplyWithAccess(t *testing.T) {
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

	config := testHyperdriveOverAccessConfig(
		rnd,
		accountID,
		rnd,
		databasePassword,
		origin,
		true,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_HyperdriveWithAccess(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHyperdriveConfigDestroy,
		Steps: []resource.TestStep{
			{
				// Initial creation
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "origin.password", databasePassword),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_id", accessClientID),
					resource.TestCheckResourceAttr(resourceName, "origin.access_client_secret", accessClientSecret),
				),
			},
			{
				// Same config - should not produce any changes
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// Config helpers

func testHyperdriveConfig(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfig.tf",
		rnd, accountId, name, password, origin.Database, origin.Host, fmt.Sprintf("%d", origin.Port), origin.Scheme, origin.User, cacheEnabled)
}

func testHyperdriveConfigMinimum(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin) string {
	return acctest.LoadTestCase("hyperdriveconfigminimum.tf",
		rnd, accountId, name, password, origin.Database, origin.Host, origin.Port, origin.Scheme, origin.User)
}

func testHyperdriveConfigUpdate(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigupdate.tf",
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

func testHyperdriveConfigConnectionLimit(rnd, accountId, name string, connLimit int, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigconnlimit.tf",
		rnd,
		accountId,
		name,
		connLimit,
		password,
		origin.Database,
		origin.Host,
		origin.Port,
		origin.Scheme,
		origin.User,
		cacheEnabled,
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

func testHyperdriveOverAccessUpdate(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOriginWithSecrets, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigoveraccessupdate.tf",
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

func testHyperdriveConfigDataSource(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfigdatasource.tf",
		rnd, accountId, name, password, origin.Database, origin.Host, origin.Port, origin.Scheme, origin.User, cacheEnabled)
}

func testHyperdriveConfigListDataSource(rnd, accountId, name string, password string, origin cfv1.HyperdriveConfigOrigin, cacheEnabled bool) string {
	return acctest.LoadTestCase("hyperdriveconfiglistdatasource.tf",
		rnd, accountId, name, password, origin.Database, origin.Host, origin.Port, origin.Scheme, origin.User, cacheEnabled)
}
