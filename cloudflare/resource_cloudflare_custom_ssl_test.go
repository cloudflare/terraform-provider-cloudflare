package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareCustomSSL_Basic(t *testing.T) {
	t.Parallel()
	var customSSL cloudflare.ZoneCustomSSL
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_ssl." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomSSLCertBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomSSLExists(resourceName, &customSSL),
					resource.TestCheckResourceAttr(
						resourceName, "zone_id", zoneID),
					resource.TestMatchResourceAttr(
						resourceName, "priority", regexp.MustCompile("^[0-9]\\d*$")),
					resource.TestCheckResourceAttr(
						resourceName, "status", "active"),
					resource.TestMatchResourceAttr(
						resourceName, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(
						resourceName, "custom_ssl_options.bundle_method", "ubiquitous"),
					resource.TestCheckResourceAttr(
						resourceName, "custom_ssl_options.type", "legacy_custom"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomSSLCertBasic(zoneID string, rName string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_ssl" "%[2]s" {
  zone_id = "%[1]s"
  custom_ssl_options = {
    certificate = "-----BEGIN CERTIFICATE-----\nMIIEsTCCA5mgAwIBAgISA53fvg2BvlK2QXSkdZewcNo4MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA2MjUyMTAzNDdaFw0y\nMDA5MjMyMTAzNDdaMB4xHDAaBgNVBAMTE3RlcnJhZm9ybS5jZmFwaS5uZXQwdjAQ\nBgcqhkjOPQIBBgUrgQQAIgNiAASBYi00+H4E7uUeogweuutTWvuAz8TC6ClQYemH\nCGA6xKrvSgWwjhvVM9joPhGlbUDbINKhVMdZd7q3DgBinVu9GjjKf1Ajxnr6nEsK\naq37tZmtUFawbqnJHAI+O3uTan+jggJpMIICZTAOBgNVHQ8BAf8EBAMCB4AwHQYD\nVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0O\nBBYEFACS0TnEhBjGvOG127Yn2O1/UCOoMB8GA1UdIwQYMBaAFKhKamMEfd265tE5\nt6ZFZe/zqOyhMG8GCCsGAQUFBwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29j\nc3AuaW50LXgzLmxldHNlbmNyeXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2Nl\ncnQuaW50LXgzLmxldHNlbmNyeXB0Lm9yZy8wHgYDVR0RBBcwFYITdGVycmFmb3Jt\nLmNmYXBpLm5ldDBMBgNVHSAERTBDMAgGBmeBDAECATA3BgsrBgEEAYLfEwEBATAo\nMCYGCCsGAQUFBwIBFhpodHRwOi8vY3BzLmxldHNlbmNyeXB0Lm9yZzCCAQUGCisG\nAQQB1nkCBAIEgfYEgfMA8QB3AF6nc/nfVsDntTZIfdBJ4DJ6kZoMhKESEoQYdZaB\ncUVYAAABcu2CH2EAAAQDAEgwRgIhAK4dA41POH3dCyi/5CN98MbBRAl8a6LyeQls\nJyZ+y1sIAiEAoMtsQKVgf8APT7/DGj/b4OzMO6EBKWcrGkZpTi7nyyQAdgCyHgXM\ni6LNiiBOh2b5K7mKJSBna9r6cOeySVMt74uQXgAAAXLtgh9PAAAEAwBHMEUCIQC1\nnxSRx2fcqG8gw5z0QK5PGktggqIulg2Jrwr20ZfXKwIgGxNlOEucj1t71h4PaLuy\nnBigJo57ztE5t56o0dlUOzEwDQYJKoZIhvcNAQELBQADggEBACy8MS07SVQLMeGK\na3E7jn7mQciQkt063tnIYbvnUTeYQZVe1Rzk6Tm9GyQoL7MIFAvTHbsB9bNzIRrl\nubefCn4s6PHnVyDGiPY/yQgGjymXyxcsfwVnc3XO3i6N8AN1MQuKMx+Kx69sHVpa\nKq9Qlu1HlStlX/eUWMcoDk1WaCJ7xm17npvdWDweDg71Qlgnl6ukggN+cQwKepw5\n4tMnqmhrzMH+xnH2dTIQ10lgB31AlwBSbOUymhg8XN+BIeXW54mBjdxkBd++7+0q\nv7oFDmljpwQSAC2BMU8ah7lwRhQxgTrG0z10Qdje1CJ8ylRHArIeISlx+jBAwKQh\nulkb7Ck=\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDD+Um5v/lCBTCvHEcZlLnSz6XX1fEOk5FxfUdiQvcY5x6WXuu3dDgDf\nvKIS0J6AsxygBwYFK4EEACKhZANiAASBYi00+H4E7uUeogweuutTWvuAz8TC6ClQ\nYemHCGA6xKrvSgWwjhvVM9joPhGlbUDbINKhVMdZd7q3DgBinVu9GjjKf1Ajxnr6\nnEsKaq37tZmtUFawbqnJHAI+O3uTan8=\n-----END EC PRIVATE KEY-----\n"
    bundle_method = "ubiquitous",
    geo_restrictions = "us"
    type = "legacy_custom"
  }
}`, zoneID, rName)
}

func testAccCheckCloudflareCustomSSLDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_ssl" {
			continue
		}

		err := client.DeleteSSL(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cert still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareCustomSSLExists(n string, customSSL *cloudflare.ZoneCustomSSL) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundCustomSSL, err := client.SSLDetails(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundCustomSSL.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}

		*customSSL = foundCustomSSL

		return nil
	}
}

func TestFlattenCustomSSLOptionsOmitsEmptyGeoRestrictions(t *testing.T) {
	customSSLOptions := cloudflare.ZoneCustomSSLOptions{
		Certificate:     "cert",
		PrivateKey:      "key",
		BundleMethod:    "method",
		GeoRestrictions: nil,
		Type:            "type",
	}

	flattenedSettings := flattenCustomSSLOptions(customSSLOptions)
	if _, ok := flattenedSettings["geo_restrictions"]; ok {
		t.Error("Expected flattenCustomSSLOptions to omit geo_restrictions when nil")
	}
}

func TestFlattenCustomSSLOptionsIncludesGeoRestrictions(t *testing.T) {
	customSSLOptions := cloudflare.ZoneCustomSSLOptions{
		Certificate:  "cert",
		PrivateKey:   "key",
		BundleMethod: "method",
		GeoRestrictions: &cloudflare.ZoneCustomSSLGeoRestrictions{
			Label: "label",
		},
		Type: "type",
	}

	flattenedSettings := flattenCustomSSLOptions(customSSLOptions)
	geoRestrictions, ok := flattenedSettings["geo_restrictions"]
	if !ok {
		t.Error("Expected flattenCustomSSLOptions to include geo_restrictions when not nil")
	}

	if geoRestrictions != "label" {
		t.Error("Expected value of geo_restrictions to match the Label property of provided ZoneCustomSSLGeoRestrictions struct")
	}
}
