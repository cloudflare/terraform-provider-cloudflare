package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudflareWAFPackage_CreateThenUpdate(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	packageID, err := testAccGetWAFPackage(zoneID)
	if err != nil {
		t.Errorf(err.Error())
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waf_package.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWAFPackageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFPackageConfig(zoneID, packageID, "medium", "simulate", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "package_id", packageID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "sensitivity", "medium"),
					resource.TestCheckResourceAttr(name, "action_mode", "simulate"),
				),
			},
			{
				Config: testAccCheckCloudflareWAFPackageConfig(zoneID, packageID, "low", "block", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "package_id", packageID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "sensitivity", "low"),
					resource.TestCheckResourceAttr(name, "action_mode", "block"),
				),
			},
		},
	})
}

func testAccGetWAFPackage(zoneID string) (string, error) {
	config := Config{}
	if apiToken, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		config.APIToken = apiToken
	} else if apiKey, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		config.APIKey = apiKey
		if email, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
			config.Email = email
		} else {
			return "", fmt.Errorf("Cloudflare email is not set correctly")
		}
	} else {
		return "", fmt.Errorf("Cloudflare credentials are not set correctly")
	}

	client, err := config.Client()
	if err != nil {
		return "", err
	}

	pkgList, err := client.ListWAFPackages(zoneID)
	if err != nil {
		return "", fmt.Errorf("Error while listing WAF packages: %s", err)
	}

	for _, pkg := range pkgList {
		if pkg.DetectionMode == "anomaly" {
			return pkg.ID, nil
		}
	}

	return "", fmt.Errorf("No anomaly package found")
}

func testAccCheckCloudflareWAFPackageDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_waf_package" {
			continue
		}

		pkg, err := client.WAFPackage(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if pkg.Sensitivity != "high" {
			return fmt.Errorf("Expected sensitivity to be reset to high, got: %s", pkg.Sensitivity)
		}
		if pkg.ActionMode != "challenge" {
			return fmt.Errorf("Expected action_mode to be reset to challenge, got: %s", pkg.ActionMode)
		}
	}

	return nil
}

func testAccCheckCloudflareWAFPackageConfig(zoneID, packageID, sensitivity, actionMode, name string) string {
	return fmt.Sprintf(`
				resource "cloudflare_waf_package" "%[5]s" {
					zone_id = "%[1]s"
					package_id = "%[2]s"
					sensitivity = "%[3]s"
					action_mode = "%[4]s"
				}`, zoneID, packageID, sensitivity, actionMode, name)
}
