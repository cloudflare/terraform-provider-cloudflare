package sdkv2provider

import (
    "context"
    "fmt"
    "os"
    "testing"

    cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareKeylessSSL_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare.keyless_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareKeylessCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareKeylessCertificate(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "port", "24008"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "name", name),
					resource.TestCheckResourceAttr(name, "host", domain),
				),
			},
		},
	})
}

func testAccCheckCloudflareKeylessCertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_keyless_certificate" {
			continue
		}

		keylessCertificates, err := client.ListKeylessSSL(context.Background(), zoneID)

		if err == nil {
			for _, keylessCertificate := range keylessCertificates {
				if keylessCertificate.ID == rs.Primary.Attributes["id"] {
					return fmt.Errorf("Keyless SSL still exists")
				}
			}
		}
	}

	return nil
}

func testAccCloudflareKeylessCertificate(resourceName, zoneId string, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_keyless_certificate" "%[1]s" {
  zone_id       = "%[2]s"
  bundle_method = "ubiquitous"
  name          = "%[1]s"
  host          = "%[3]s"
  port          = 24008
  enabled       = true
  certificate   = "-----BEGIN CERTIFICATE-----\nMIIFQDCCBCigAwIBAgISBITqGEnFOTnEKy0WPSby659TMA0GCSqGSIb3DQEBCwUA\nMDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD\nEwJSMzAeFw0yMTA5MTIyMjE0MjdaFw0yMTEyMTEyMjE0MjZaMB4xHDAaBgNVBAMT\nE3RlcnJhZm9ybS5jZmFwaS5uZXQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQDVHcbuasE2kqFqPagrdUN2OddOkZsujnMe+GVDV65hwK8OFQGRdeiLuXhM\nc4yyAt4eEUNxP+H51HssdKPKPur9lWvBkciHGNvVsoVsWY1QKzhctcZi/TXGi89p\nqnynyMbLSEosr7QXLoVih0i6EgHIZhqT3Iz9MQd5ZymuPnyZBN/DCv32Dhdlueav\n0Q2Dqd7PThmtRBYs5odlF2MNWfwOyxRmJXfI66zTGtdgUTq8Fxk9d/RLt+kIWO7y\nBpMdIUPRVmLwkPO07tFYiG6VtcmTdPMtZsmJwcDABc0qU+U8NpRmigwnLIzsjfwb\nH06wwRhO8N1pQfBPDGtpYp4T3/LDAgMBAAGjggJiMIICXjAOBgNVHQ8BAf8EBAMC\nBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAw\nHQYDVR0OBBYEFK7MWl1qlU2RrL+WlL+QWJjY8poCMB8GA1UdIwQYMBaAFBQusxe3\nWFbLrlAJQOYfr52LFMLGMFUGCCsGAQUFBwEBBEkwRzAhBggrBgEFBQcwAYYVaHR0\ncDovL3IzLm8ubGVuY3Iub3JnMCIGCCsGAQUFBzAChhZodHRwOi8vcjMuaS5sZW5j\nci5vcmcvMDQGA1UdEQQtMCuCE3RlcnJhZm9ybS5jZmFwaS5uZXSCFHRlcnJhZm9y\nbTIuY2ZhcGkubmV0MEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEB\nMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAgYK\nKwYBBAHWeQIEAgSB8wSB8ADuAHUAXNxDkv7mq0VEsV6a1FbmEDf71fpH3KFzlLJe\n5vbHDsoAAAF73EpjEQAABAMARjBEAiB7JXTWsVOKjbRJUhh8nD7BTpo4kYavQ88V\n6AdiTJJTGgIgI9gdMaF0NLpV3SO6J7LvH8ruQ+aTdgmQRoG5o89xVt0AdQD2XJQv\n0XcwIhRUGAgwlFaO400TGTO/3wwvIAvMTvFk4wAAAXvcSmMFAAAEAwBGMEQCIDqN\nolVOMaRyX57A952HltGv7kHvbpP1Cq1Hlx6wtBHvAiBpF6WhzPklj4omAmALxcHR\nmunqNwK1RTZWi0GVAVRQsjANBgkqhkiG9w0BAQsFAAOCAQEAeUhP+bGbtpwREWn6\nbDbGGg5lIBZ1zgzrotM16YcrpzS/BHOpQps7uqmt8aP68aGAyJl3lB2fF2TM8klv\nEoXvG4rGHlRtHZhllCtD1T5f9APKH88F+LoqYyp/m049LZCY/9WCgkXrqNtSbLut\nAr7b1LqvDpyS4m7cW/uG1mk3dsHjmJuwYhk3W/xWyBa6FFxHowbxDSXRGkSJ6SWC\nEXD0YagNvfpm+kNB58pJSIbBpbNL0mJA7gy2BWN58Sb0DMK+gam79QSLrZKdIlq/\nYQWun8yGsH8gHJFWyGcHtnQsGYvMd0Dr7Xf1uIOn/eQujFjF6i9/D5FTxnR5Stbb\nPwneVQ==\n-----END CERTIFICATE-----\n"
}`, resourceName, zoneId, domain)
}