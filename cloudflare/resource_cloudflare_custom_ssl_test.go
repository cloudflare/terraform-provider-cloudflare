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
    certificate = "-----BEGIN CERTIFICATE-----\nMIIFXDCCBESgAwIBAgISAxQZ/dK8WhbgqBOFP66zZDd1MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0xOTEwMjEwMTI1MjFaFw0y\nMDAxMTkwMTI1MjFaMB4xHDAaBgNVBAMTE3RlcnJhZm9ybS5jZmFwaS5uZXQwggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDcRBwdivjZ+s+kAp0Epy48+/fe\nBLFOmFtlfZYHN4ElK5zEpG8dSIY0GJIOksfm1zZz/VvltuSPgpPzOYFhY3HrDkZg\n1iOfnsw4NvjgXpbcI4v1dgn9e7vQN1EP4SK4hHF5vG7WACIojKmx/T/fBdbAzPgO\nRPhxl93dYOZ3sEKp7wn0Ai8RIOeKWPP7DVsTrzDxokr/nwYhvYOt6fxxe/Z6ysQs\nqXpcEOZY8TZaOgUTgNr9vw2gesE5IcZXy3ipxIqUol6h5OSQx1WHNJvWcfvpY+9P\n7AWSQeYBBPNbKd5YYIw+khoJKTsqAHodVNJi3V6rdD8FUW2Kb+IRKuHYA+pBAgMB\nAAGjggJmMIICYjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG\nCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFKIEjo1TrwyyaQSzDOtm\nkoOpIxASMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUF\nBwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZy8wHgYDVR0RBBcwFYITdGVycmFmb3JtLmNmYXBpLm5ldDBMBgNVHSAE\nRTBDMAgGBmeBDAECATA3BgsrBgEEAYLfEwEBATAoMCYGCCsGAQUFBwIBFhpodHRw\nOi8vY3BzLmxldHNlbmNyeXB0Lm9yZzCCAQIGCisGAQQB1nkCBAIEgfMEgfAA7gB1\nAF6nc/nfVsDntTZIfdBJ4DJ6kZoMhKESEoQYdZaBcUVYAAABbewiHA8AAAQDAEYw\nRAIgPgC5c8nbuh1cp9FcuACBPNp2SsdWmq4q5uiz3cPKotUCIHF/eWflSD1AjItM\n3X9fO7aPyX3IFOpwwyETqQQQu7sZAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDn\nsklTLe+LkF4AAAFt7CIcMQAABAMARjBEAiB13X5HnLppuoGTZGhz40igQNmh03LG\nETHAAsVEgU/N4QIgCfk6fthDKEax0DQGHfIWnFVYOxAhV2jtEWXjiyrQkLAwDQYJ\nKoZIhvcNAQELBQADggEBAF5lKWslzLAOlUEBu4k+M3jgaYooAKsup0B0cHSrZRF6\n85Shif8RX59WnHDblcTvOhmKMcnPu6lGP4bogCjzBFiytvCFkDdL6n09ZiT7eHUp\nn7B5r2qvGLRhkEeN8bkL29Hd5oLOR7vkftMmRFTdGR6vwTfRBAZuqoa2hbO/VVlc\nlcis2tGHQsukM+Z2P+BbLE9cmuUaFAbXPyxeMot6ggJVV269XYXtkcyaOjaC/pLF\nnZ8IW22RnWtMQdUEAG4YNaHQOJ6IfrSP0LcTbMzd69fYWYHkPQu7Js4FtH42OUrj\nREZ7lXEGBO6PbCrwsiXiMEj8sIsBYYHTKHkhKr1rUXU=\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDcRBwdivjZ+s+k\nAp0Epy48+/feBLFOmFtlfZYHN4ElK5zEpG8dSIY0GJIOksfm1zZz/VvltuSPgpPz\nOYFhY3HrDkZg1iOfnsw4NvjgXpbcI4v1dgn9e7vQN1EP4SK4hHF5vG7WACIojKmx\n/T/fBdbAzPgORPhxl93dYOZ3sEKp7wn0Ai8RIOeKWPP7DVsTrzDxokr/nwYhvYOt\n6fxxe/Z6ysQsqXpcEOZY8TZaOgUTgNr9vw2gesE5IcZXy3ipxIqUol6h5OSQx1WH\nNJvWcfvpY+9P7AWSQeYBBPNbKd5YYIw+khoJKTsqAHodVNJi3V6rdD8FUW2Kb+IR\nKuHYA+pBAgMBAAECggEAOuz7qomdpk77UpCiafbBn2328w3FU1XeCYot4zNdmNCQ\niWjGEwruYCx8kmqBEQfdGO2cMpmZjqzroKNvMdsVREFW4ZwX/qFQV++Y6AlWCYF2\n9U1FQetipMLPrFRZ4gwIgX1AF38EpF1xHl02Q5c/yudqqwKXhmgO4iUEsttUEjqS\nEtDjvswSpPasrYhaNqu/Lsm+9BLE3iN3OamAZGD4NVBRhfByd3VnwiqV2sGnF5VL\nFYz4nUn/elaX83ytRrqEsM20vsp44VCup8w3HOZJffVsYsWoHU7qCZG1gx1LLPnv\nG15R6hRyAxuLaywZ+PfFLkNkjH9Cypci3QYhllWOWwKBgQDzbADzZ6Gzl+Wy3f1C\nAQCtAvwo6R6D8j9UrwUB2EH3OV+7ON544Y+PMkGGFdHJ5TJ3dDlcGMUYIf8I6StW\nKBRY0hbTZxL/mLaag5qZWUNq3m9h9teZ8Z1kM7FIDjdXlPpNcVAO+Nh/eWIVIGxb\nfeJALScXPe79+BdWcf2I3Let2wKBgQDnpcy0E9tF8zvpq3Dz/gOQTJTGmul6sGPB\nxmIto8GXnDtBDB6UHXHaOFO+LWrESPwXsxPBg9ESIriN0Mvp1RohDRoET/eI0oRN\nLMeTKnOW1+Fw7GLLnjQojQh0s3Wv4mOuegJEaZN7fI5xJ7KEahWRRmByV8B6xe00\nF02qkVj5EwKBgQCHioANCItVgSL5sfovInfJ3nuiHAxN1DnHYZ0cJdq1WlEf4s6d\n6JsTVRx/GO8zyFeNhD3cNj6o7WUhBRSIaNDLlE/5bs95WwNyjg0rjjSn8St8FQKA\nSbUl8lKomKHgNqgZLxsw+wcyE9i1gtRTLYkpyvqVSnslF1uHWvmdl6j/OQKBgCKi\nuLvIKEYKO1AR8T6aIWBHAwu7B+PvUcscZXube4u2sWllbYEJ8gcF2weZdNhKbV8B\nyJdrpSwIAv45VPPuiAyfD9/LMSDFEUEUy/ZmJ4hLWQrwXUlCq1vQ0o3Yc2VL/UmO\nNp6SBpo1Insqy1dfIUBqfGs8UaxJwdDDFzrEpr//AoGAWRgHlqbKIer0nJy02oWU\nsoNw0pG4ULxtpEriFH6fhs8s8lK5fWOOBpCL3dv5PmImCiME2Ks4CKQr/EPTs4gu\nFWolAEAGRRVF81ohrLTTVk2XLqmY3tAreA2e4ymgJOWtzflTEw6fwI9uONthvmtz\nk6hqKdazK0jNghMBeIawdSE=\n-----END PRIVATE KEY-----\n"
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
