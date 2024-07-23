package spectrum_application_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_spectrum_applications", &resource.Sweeper{
		Name: "cloudflare_spectrum_applications",
		F:    testSweepCloudflareSpectrumApplications,
	})
}

func testSweepCloudflareSpectrumApplications(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zones, zoneErr := client.ListZones(context.Background(), zone)

	if zoneErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare zones: %s", zoneErr))
	}

	for _, zone := range zones {
		spectrumApps, spectrumErr := client.SpectrumApplications(context.Background(), zone.ID)
		if spectrumErr != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare spectrum applications: %s", zoneErr))
		}

		if len(spectrumApps) == 0 {
			log.Print("[DEBUG] No Cloudflare spectrum applications to sweep")
			return nil
		}

		for _, application := range spectrumApps {
			tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare spectrum application ID: %s", application.ID))
			err := client.DeleteSpectrumApplication(context.Background(), zone.ID, application.ID)

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare spectrum application (%s) in zone ID: %s", application.ID, zone.ID))
			}
		}
	}

	return nil
}

func TestAccCloudflareSpectrumApplication_Basic(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "origin_direct.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://128.66.0.1:23"),
					resource.TestCheckResourceAttr(name, "origin_port", "22"),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_OriginDNS(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigOriginDNS(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "origin_dns.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_dns.0.name", fmt.Sprintf("%s.origin.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "origin_port", "22"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_OriginPortRange(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigOriginPortRange(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22-23"),
					resource.TestCheckResourceAttr(name, "origin_dns.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_dns.0.name", fmt.Sprintf("%s.origin.%s", rnd, domain)),
					resource.TestCheckResourceAttr(name, "origin_port_range.#", "1"),
					resource.TestCheckResourceAttr(name, "origin_port_range.0.start", "2022"),
					resource.TestCheckResourceAttr(name, "origin_port_range.0.end", "2023"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_Update(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	var initialID string
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSpectrumApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://128.66.0.1:23"),
				),
			},
			{
				PreConfig: func() {
					initialID = spectrumApp.ID
				},
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasicUpdated(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					func(state *terraform.State) error {
						if initialID != spectrumApp.ID {
							// want in place update
							return fmt.Errorf("spectrum application id is different after second config applied ( %s -> %s )", initialID, spectrumApp.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "origin_direct.0", "tcp://128.66.0.2:23"),
				),
			},
		},
	})
}

func testAccCheckCloudflareSpectrumApplicationDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_spectrum_application" {
			continue
		}

		_, err := client.SpectrumApplication(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("spectrum application still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func TestAccCloudflareSpectrumApplication_EdgeIPConnectivity(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigEdgeIPConnectivity(zoneID, domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "edge_ips.0.connectivity", "ipv4"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_EdgeIPsMultiple(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigMultipleEdgeIPs(zoneID, domain, rnd, `"172.65.64.13", "172.65.64.49"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "edge_ips.0.ips.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "edge_ips.0.ips.*", "172.65.64.13"),
					resource.TestCheckTypeSetElemAttr(name, "edge_ips.0.ips.*", "172.65.64.49"),
				),
			},
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigMultipleEdgeIPs(zoneID, domain, rnd, `"172.65.64.49", "172.65.64.13"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "edge_ips.0.ips.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "edge_ips.0.ips.*", "172.65.64.13"),
					resource.TestCheckTypeSetElemAttr(name, "edge_ips.0.ips.*", "172.65.64.49"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_BasicSSH(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasicTypes(zoneID, domain, rnd, "ssh", 22),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "ssh"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_BasicRDP(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasicTypes(zoneID, domain, rnd, "rdp", 3389),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "rdp"),
				),
			},
		},
	})
}

func TestAccCloudflareSpectrumApplication_BasicMinecraft(t *testing.T) {
	var spectrumApp cloudflare.SpectrumApplication
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasicTypes(zoneID, domain, rnd, "minecraft", 25565),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &spectrumApp),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
					resource.TestCheckResourceAttr(name, "protocol", "minecraft"),
				),
			},
		},
	})
}

func testAccCheckCloudflareSpectrumApplicationExists(n string, spectrumApp *cloudflare.SpectrumApplication) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundSpectrumApplication, err := client.SpectrumApplication(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		*spectrumApp = foundSpectrumApplication

		return nil
	}
}

func testAccCheckCloudflareSpectrumApplicationIDIsValid(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		if len(rs.Primary.ID) != 32 {
			return fmt.Errorf("invalid id %q, should be a string of length 32", rs.Primary.ID)
		}

		if zoneID, ok := rs.Primary.Attributes[consts.ZoneIDSchemaKey]; !ok || len(zoneID) < 1 {
			return errors.New("zone_id is unset, should always be set with id")
		}
		return nil
	}
}

func testAccManuallyDeleteSpectrumApplication(name string, spectrumApp *cloudflare.SpectrumApplication, initialID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		*initialID = spectrumApp.ID
		err := client.DeleteSpectrumApplication(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareSpectrumApplicationConfigBasic(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigbasic.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigOriginDNS(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigorigindns.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigOriginPortRange(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigoriginportrange.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigBasicUpdated(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigbasicupdated.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigEdgeIPConnectivity(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigedgeipconnectivity.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigEdgeIPsWithoutConnectivity(zoneID, zoneName, ID string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigedgeipswithoutconnectivity.tf", zoneID, zoneName, ID)
}

func testAccCheckCloudflareSpectrumApplicationConfigMultipleEdgeIPs(zoneID, zoneName, ID, IPs string) string {
	return acctest.LoadTestCase("spectrumapplicationconfigmultipleedgeips.tf", zoneID, zoneName, ID, IPs)
}

func testAccCheckCloudflareSpectrumApplicationConfigBasicTypes(zoneID, zoneName, ID, protocol string, port int) string {
	return acctest.LoadTestCase("spectrumapplicationconfigbasictypes.tf", zoneID, zoneName, ID, protocol, port)
}
