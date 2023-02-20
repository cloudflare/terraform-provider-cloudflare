package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_load_balancer_monitor", &resource.Sweeper{
		Name: "cloudflare_load_balancer_monitor",
		F:    testSweepCloudflareLoadBalancerMonitors,
	})
}

func testSweepCloudflareLoadBalancerMonitors(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	monitors, err := client.ListLoadBalancerMonitors(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListLoadBalancerMonitorParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Load Balancer Monitors: %s", err))
	}

	if len(monitors) == 0 {
		log.Print("[DEBUG] No Cloudflare Load Balancer Monitors to sweep")
		return nil
	}

	for _, monitor := range monitors {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer Monitor ID: %s", monitor.ID))
		//nolint:errcheck
		client.DeleteLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), monitor.ID)
	}

	return nil
}

func TestAccCloudflareLoadBalancerMonitor_Basic(t *testing.T) {
	testStartTime := time.Now().UTC()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	name := "cloudflare_load_balancer_monitor." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "description", ""),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerMonitorDates(name, &loadBalancerMonitor, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_FullySpecified(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(zoneName, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "path", "/custom"),
					resource.TestCheckResourceAttr(name, "header.#", "1"),
					resource.TestCheckResourceAttr(name, "retries", "5"),
					resource.TestCheckResourceAttr(name, "port", "8080"),
					resource.TestCheckResourceAttr(name, "expected_body", "dead"),
					resource.TestCheckResourceAttr(name, "probe_zone", zoneName),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_EmptyExpectedBody(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigEmptyExpectedBody(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// checking empty string value passes all validations and created
					resource.TestCheckResourceAttr(name, "expected_body", ""),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_TcpFullySpecified(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := "cloudflare_load_balancer_monitor.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigTcpFullySpecified(accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "retries", "5"),
					resource.TestCheckResourceAttr(name, "port", "8080"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_PremiumTypes(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigUDPICMP(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// check we can create one of the correct type
					resource.TestCheckResourceAttr(name, "type", "udp_icmp"),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigICMPPing(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// check we can create one of the correct type
					resource.TestCheckResourceAttr(name, "type", "icmp_ping"),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigSMTP(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "type", "smtp"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_NoRequired(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareLoadBalancerMonitorConfigMissingRequired(accountID),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("expected_codes must be set")),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_Update(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	var initialId string
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
				),
			},
			{
				PreConfig: func() {
					initialId = loadBalancerMonitor.ID
				},
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(zoneName, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					func(state *terraform.State) error {
						if initialId != loadBalancerMonitor.ID {
							return fmt.Errorf("wanted update but monitor got recreated (id changed %q -> %q)", initialId, loadBalancerMonitor.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_CreateAfterManualDestroy(t *testing.T) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	var initialId string
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					testAccManuallyDeleteLoadBalancerMonitor(name, &loadBalancerMonitor, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					func(state *terraform.State) error {
						if initialId == loadBalancerMonitor.ID {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )", loadBalancerMonitor.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_ChangingHeadersCauseReplacement(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigWithHeaders(rnd, domain, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "header.0.header", "Host"),
					resource.TestCheckResourceAttr(name, "header.0.values.0", domain),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigWithHeaders(rnd, fmt.Sprintf("%s.%s", rnd, domain), accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "header.0.header", "Host"),
					resource.TestCheckResourceAttr(name, "header.0.values.0", fmt.Sprintf("%s.%s", rnd, domain)),
				),
			},
		},
	})
}

func testAccCheckCloudflareLoadBalancerMonitorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer_monitor" {
			continue
		}

		_, err := client.GetLoadBalancerMonitor(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer monitor still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareLoadBalancerMonitorExists(n string, load *cloudflare.LoadBalancerMonitor) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundLoadBalancerMonitor, err := client.GetLoadBalancerMonitor(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), rs.Primary.ID)
		if err != nil {
			return err
		}

		*load = foundLoadBalancerMonitor

		return nil
	}
}

func testAccCheckCloudflareLoadBalancerMonitorDates(n string, loadBalancerMonitor *cloudflare.LoadBalancerMonitor, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": *loadBalancerMonitor.CreatedOn, "modified_on": *loadBalancerMonitor.ModifiedOn} {
			timeStamp, err := time.Parse(time.RFC3339Nano, rs.Primary.Attributes[timeStampAttr])
			if err != nil {
				return err
			}

			if timeStamp != serverVal {
				return fmt.Errorf("state value of %s: %s is different than server created value: %s", timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("state value of %s: %s should be greater than test start time: %s", timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccManuallyDeleteLoadBalancerMonitor(name string, loadBalancerMonitor *cloudflare.LoadBalancerMonitor, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = loadBalancerMonitor.ID
		err := client.DeleteLoadBalancerMonitor(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), loadBalancerMonitor.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  expected_body = "alive"
  expected_codes = "2xx"
}`, rnd, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(zoneName, accountID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[3]s" {
  account_id = "%[2]s"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  port = 8080
  description = "this is a very weird load balancer"
  probe_zone = "%[1]s"
  header {
    header = "Host"
    values = ["%[1]s"]
  }
}`, zoneName, accountID, rnd)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigWithHeaders(rnd, hostname, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[3]s"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  port = 8080
  description = "this is a very weird load balancer"
  header {
    header = "Host"
    values = ["%[2]s"]
  }
}`, rnd, hostname, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigEmptyExpectedBody(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  expected_body = ""
  expected_codes = "2xx"
  description = "we don't want to check for a given body"
}`, resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigTcpFullySpecified(accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "test" {
  account_id = "%[1]s"
  type = "tcp"
  method = "connection_established"
  timeout = 9
  interval = 60
  retries = 5
  port = 8080
  description = "this is a very weird tcp load balancer"
}`, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigUDPICMP(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type = "udp_icmp"
  timeout = 2
  interval = 60
  retries = 5
  port = 8080
  description = "test setup udp_icmp"
}`, resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigICMPPing(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type = "icmp_ping"
  timeout = 2
  interval = 60
  retries = 5
  description = "test setup icmp_ping"
}`, resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigSMTP(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type = "smtp"
  timeout = 2
  interval = 60
  retries = 5
  port = 8080
  description = "test setup smtp"
}`, resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigMissingRequired(accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "test" {
	account_id = "%[1]s"
	description = "this is a wrong config"
}`, accountID)
}
