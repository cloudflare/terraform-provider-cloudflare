package cloudflare

import (
	"fmt"
	"testing"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudflareLoadBalancerMonitor_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	name := "cloudflare_load_balancer_monitor.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
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
	t.Parallel()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	name := "cloudflare_load_balancer_monitor.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "path", "/custom"),
					resource.TestCheckResourceAttr(name, "header.#", "1"),
					resource.TestCheckResourceAttr(name, "retries", "5"),
					resource.TestCheckResourceAttr(name, "expected_codes", "5xx"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_Update(t *testing.T) {
	t.Parallel()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	var initialId string
	name := "cloudflare_load_balancer_monitor.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
				),
			},
			{
				PreConfig: func() {
					initialId = loadBalancerMonitor.ID
				},
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					func(state *terraform.State) error {
						if initialId != loadBalancerMonitor.ID {
							return fmt.Errorf("wanted update but monitor got recreated (id changed %q -> %q)",
								initialId, loadBalancerMonitor.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerMonitor_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	var initialId string
	name := "cloudflare_load_balancer_monitor.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareLoadBalancerMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					testAccManuallyDeleteLoadBalancerMonitor(name, &loadBalancerMonitor, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerMonitorExists(name, &loadBalancerMonitor),
					func(state *terraform.State) error {
						if initialId == loadBalancerMonitor.ID {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								loadBalancerMonitor.ID)
						}
						return nil
					},
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

		_, err := client.LoadBalancerMonitorDetails(rs.Primary.ID)
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
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundLoadBalancerMonitor, err := client.LoadBalancerMonitorDetails(rs.Primary.ID)
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
				return fmt.Errorf("state value of %s: %s is different than server created value: %s",
					timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("State value of %s: %s should be greater than test start time: %s",
					timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccManuallyDeleteLoadBalancerMonitor(name string, loadBalancerMonitor *cloudflare.LoadBalancerMonitor, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = loadBalancerMonitor.ID
		err := client.DeleteLoadBalancerMonitor(loadBalancerMonitor.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareLoadBalancerMonitorConfigBasic() string {
	return `
resource "cloudflare_load_balancer_monitor" "test" {
  expected_body = "alive"
  expected_codes = "2xx"

}`
}

func testAccCheckCloudflareLoadBalancerMonitorConfigFullySpecified() string {
	return `
resource "cloudflare_load_balancer_monitor" "test" {
  expected_body = "dead" 
  expected_codes = "5xx" 
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 55
  retries = 5
  description = "this is a very weird load balancer"
  header {
    header = "Host"
    values = ["example.com"]
  }
}`
}
