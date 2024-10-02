package load_balancer_monitor_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
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
	rnd := utils.GenerateRandomResourceName()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	name := "cloudflare_load_balancer_monitor." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
					resource.TestCheckResourceAttr(name, "consecutive_up", "2"),
					resource.TestCheckResourceAttr(name, "consecutive_down", "2"),
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
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_load_balancer_monitor.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerMonitorDestroy,
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
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
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		*initialId = loadBalancerMonitor.ID
		err := client.DeleteLoadBalancerMonitor(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), loadBalancerMonitor.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigbasic.tf", rnd, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(zoneName, accountID, rnd string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigfullyspecified.tf", zoneName, accountID, rnd)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigWithHeaders(rnd, hostname, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigwithheaders.tf", rnd, hostname, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigEmptyExpectedBody(resourceName, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigemptyexpectedbody.tf", resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigTcpFullySpecified(accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigtcpfullyspecified.tf", accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigUDPICMP(resourceName, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigudpicmp.tf", resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigICMPPing(resourceName, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigicmpping.tf", resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigSMTP(resourceName, accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigsmtp.tf", resourceName, accountID)
}

func testAccCheckCloudflareLoadBalancerMonitorConfigMissingRequired(accountID string) string {
	return acctest.LoadTestCase("loadbalancermonitorconfigmissingrequired.tf", accountID)
}
