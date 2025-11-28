package healthcheck_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/healthchecks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_healthcheck", &resource.Sweeper{
		Name: "cloudflare_healthcheck",
		F:    testSweepCloudflareHealthcheck,
	})
}

func testSweepCloudflareHealthcheck(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping healthchecks sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	healthchecksList, err := client.Healthchecks.List(ctx, healthchecks.HealthcheckListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch healthchecks: %s", err))
		return fmt.Errorf("failed to fetch healthchecks: %w", err)
	}

	if len(healthchecksList.Result) == 0 {
		tflog.Info(ctx, "No healthchecks to sweep")
		return nil
	}

	for _, healthcheck := range healthchecksList.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(healthcheck.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting healthcheck: %s (%s) (zone: %s)", healthcheck.Name, healthcheck.ID, zoneID))
		_, err := client.Healthchecks.Delete(ctx, healthcheck.ID, healthchecks.HealthcheckDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete healthcheck %s: %s", healthcheck.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted healthcheck: %s", healthcheck.ID))
	}

	return nil
}

func TestAccCloudflareHealthcheckTCPExists(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tcp_config").AtMapKey("port"), knownvalue.Int64Exact(80)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tcp_config").AtMapKey("method"), knownvalue.StringExact("connection_established")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func TestAccCloudflareHealthcheckTCPUpdate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareHealthcheckTCP(zoneID, rnd+"-updated", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func TestAccCloudflareHealthcheckHTTPExists(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Healthcheck
	// service does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckHTTP(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("port"), knownvalue.Int64Exact(80)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("method"), knownvalue.StringExact("GET")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func TestAccCloudflareHealthcheckMissingRequired(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckHealthcheckConfigMissingRequired(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("The argument \"name\" is required, but no definition was found.")),
			},
		},
	})
}

func TestAccCloudflareHealthcheck_HTTPS(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckHTTPSComplete(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("address"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("HTTPS")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("HTTPS healthcheck with all options")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("consecutive_fails"), knownvalue.Int64Exact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("consecutive_successes"), knownvalue.Int64Exact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("interval"), knownvalue.Int64Exact(60)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("retries"), knownvalue.Int64Exact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("timeout"), knownvalue.Int64Exact(10)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("suspended"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("allow_insecure"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("expected_body"), knownvalue.StringExact("OK")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("follow_redirects"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("method"), knownvalue.StringExact("GET")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("path"), knownvalue.StringExact("/health")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("port"), knownvalue.Int64Exact(443)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func TestAccCloudflareHealthcheck_TCPCustomPort(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckTCPCustomPort(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("address"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("TCP")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("TCP healthcheck on custom port")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tcp_config").AtMapKey("method"), knownvalue.StringExact("connection_established")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tcp_config").AtMapKey("port"), knownvalue.Int64Exact(8080)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("timeout"), knownvalue.Int64Exact(15)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("retries"), knownvalue.Int64Exact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func TestAccCloudflareHealthcheck_HTTPHead(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_healthcheck.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHealthcheckHTTPHead(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("address"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("HTTP")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("method"), knownvalue.StringExact("HEAD")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("path"), knownvalue.StringExact("/ping")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("http_config").AtMapKey("port"), knownvalue.Int64Exact(8080)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"check_regions", "description", "created_on", "modified_on", "status", "failure_reason", "http_config"},
			},
		},
	})
}

func testAccCheckCloudflareHealthcheckDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_healthcheck" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.Healthchecks.Get(
			context.Background(),
			rs.Primary.ID,
			healthchecks.HealthcheckGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)
		if err == nil {
			return fmt.Errorf("healthcheck still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareHealthcheckTCP(zoneID, name, ID string) string {
	return acctest.LoadTestCase("healthchecktcp.tf", zoneID, name, ID)
}

func testAccCheckCloudflareHealthcheckHTTP(zoneID, ID string) string {
	return acctest.LoadTestCase("healthcheckhttp.tf", zoneID, ID)
}

func testAccCheckHealthcheckConfigMissingRequired(zoneID, ID string) string {
	return acctest.LoadTestCase("acccheckhealthcheckconfigmissingrequired.tf", zoneID, ID)
}

func testAccCheckCloudflareHealthcheckHTTPSComplete(zoneID, name string) string {
	return acctest.LoadTestCase("https_complete.tf", zoneID, name)
}

func testAccCheckCloudflareHealthcheckTCPCustomPort(zoneID, name string) string {
	return acctest.LoadTestCase("tcp_custom_port.tf", zoneID, name)
}

func testAccCheckCloudflareHealthcheckHTTPHead(zoneID, name string) string {
	return acctest.LoadTestCase("http_head_method.tf", zoneID, name)
}
