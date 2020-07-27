package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testCertificate = `-----BEGIN CERTIFICATE-----\nMIIDtTCCAp2gAwIBAgIJAMHAwfXZ5/PWMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTYwODI0MTY0MzAxWhcNMTYxMTIyMTY0MzAxWjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAwQHoetcl9+5ikGzV6cMzWtWPJHqXT3wpbEkRU9Yz7lgvddmGdtcGbg/1\nCGZu0jJGkMoppoUo4c3dts3iwqRYmBikUP77wwY2QGmDZw2FvkJCJlKnabIRuGvB\nKwzESIXgKk2016aTP6/dAjEHyo6SeoK8lkIySUvK0fyOVlsiEsCmOpidtnKX/a+5\n0GjB79CJH4ER2lLVZnhePFR/zUOyPxZQQ4naHf7yu/b5jhO0f8fwt+pyFxIXjbEI\ndZliWRkRMtzrHOJIhrmJ2A1J7iOrirbbwillwjjNVUWPf3IJ3M12S9pEewooaeO2\nizNTERcG9HzAacbVRn2Y2SWIyT/18QIDAQABo4GnMIGkMB0GA1UdDgQWBBT/LbE4\n9rWf288N6sJA5BRb6FJIGDB1BgNVHSMEbjBsgBT/LbE49rWf288N6sJA5BRb6FJI\nGKFJpEcwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAfBgNV\nBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZIIJAMHAwfXZ5/PWMAwGA1UdEwQF\nMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAHHFwl0tH0quUYZYO0dZYt4R7SJ0pCm2\n2satiyzHl4OnXcHDpekAo7/a09c6Lz6AU83cKy/+x3/djYHXWba7HpEu0dR3ugQP\nMlr4zrhd9xKZ0KZKiYmtJH+ak4OM4L3FbT0owUZPyjLSlhMtJVcoRp5CJsjAMBUG\nSvD8RX+T01wzox/Qb+lnnNnOlaWpqu8eoOenybxKp1a9ULzIVvN/LAcc+14vioFq\n2swRWtmocBAs8QR9n4uvbpiYvS8eYueDCWMM4fvFfBhaDZ3N9IbtySh3SpFdQDhw\nYbjM2rxXiyLGxB4Bol7QTv4zHif7Zt89FReT/NBy4rzaskDJY5L6xmY=\n-----END CERTIFICATE-----\n`

func TestAccCloudflareAuthenticatedOriginPullsCertificatePerZone(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls_certificate.%s", rnd)
	aopType := "per-zone"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsCertificateConfig(zoneID, rnd, aopType, testCertificate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "zone_id", zoneID),
				),
			},
		},
	})
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificateConfig(zoneID, name, aopType string) string {
	return fmt.Sprintf(`
  resource "cloudflare_authenticated_origin_pulls_certificate" "%[2]s" {
	  zone_id = "%[1]s"
	  certificate = "-----BEGIN CERTIFICATE-----\nMIIEsTCCA5mgAwIBAgISA53fvg2BvlK2QXSkdZewcNo4MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA2MjUyMTAzNDdaFw0y\nMDA5MjMyMTAzNDdaMB4xHDAaBgNVBAMTE3RlcnJhZm9ybS5jZmFwaS5uZXQwdjAQ\nBgcqhkjOPQIBBgUrgQQAIgNiAASBYi00+H4E7uUeogweuutTWvuAz8TC6ClQYemH\nCGA6xKrvSgWwjhvVM9joPhGlbUDbINKhVMdZd7q3DgBinVu9GjjKf1Ajxnr6nEsK\naq37tZmtUFawbqnJHAI+O3uTan+jggJpMIICZTAOBgNVHQ8BAf8EBAMCB4AwHQYD\nVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0O\nBBYEFACS0TnEhBjGvOG127Yn2O1/UCOoMB8GA1UdIwQYMBaAFKhKamMEfd265tE5\nt6ZFZe/zqOyhMG8GCCsGAQUFBwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29j\nc3AuaW50LXgzLmxldHNlbmNyeXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2Nl\ncnQuaW50LXgzLmxldHNlbmNyeXB0Lm9yZy8wHgYDVR0RBBcwFYITdGVycmFmb3Jt\nLmNmYXBpLm5ldDBMBgNVHSAERTBDMAgGBmeBDAECATA3BgsrBgEEAYLfEwEBATAo\nMCYGCCsGAQUFBwIBFhpodHRwOi8vY3BzLmxldHNlbmNyeXB0Lm9yZzCCAQUGCisG\nAQQB1nkCBAIEgfYEgfMA8QB3AF6nc/nfVsDntTZIfdBJ4DJ6kZoMhKESEoQYdZaB\ncUVYAAABcu2CH2EAAAQDAEgwRgIhAK4dA41POH3dCyi/5CN98MbBRAl8a6LyeQls\nJyZ+y1sIAiEAoMtsQKVgf8APT7/DGj/b4OzMO6EBKWcrGkZpTi7nyyQAdgCyHgXM\ni6LNiiBOh2b5K7mKJSBna9r6cOeySVMt74uQXgAAAXLtgh9PAAAEAwBHMEUCIQC1\nnxSRx2fcqG8gw5z0QK5PGktggqIulg2Jrwr20ZfXKwIgGxNlOEucj1t71h4PaLuy\nnBigJo57ztE5t56o0dlUOzEwDQYJKoZIhvcNAQELBQADggEBACy8MS07SVQLMeGK\na3E7jn7mQciQkt063tnIYbvnUTeYQZVe1Rzk6Tm9GyQoL7MIFAvTHbsB9bNzIRrl\nubefCn4s6PHnVyDGiPY/yQgGjymXyxcsfwVnc3XO3i6N8AN1MQuKMx+Kx69sHVpa\nKq9Qlu1HlStlX/eUWMcoDk1WaCJ7xm17npvdWDweDg71Qlgnl6ukggN+cQwKepw5\n4tMnqmhrzMH+xnH2dTIQ10lgB31AlwBSbOUymhg8XN+BIeXW54mBjdxkBd++7+0q\nv7oFDmljpwQSAC2BMU8ah7lwRhQxgTrG0z10Qdje1CJ8ylRHArIeISlx+jBAwKQh\nulkb7Ck=\n-----END CERTIFICATE-----\n"
	  private_key = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDD+Um5v/lCBTCvHEcZlLnSz6XX1fEOk5FxfUdiQvcY5x6WXuu3dDgDf\nvKIS0J6AsxygBwYFK4EEACKhZANiAASBYi00+H4E7uUeogweuutTWvuAz8TC6ClQ\nYemHCGA6xKrvSgWwjhvVM9joPhGlbUDbINKhVMdZd7q3DgBinVu9GjjKf1Ajxnr6\nnEsKaq37tZmtUFawbqnJHAI+O3uTan8=\n-----END EC PRIVATE KEY-----\n"
      type = "%[3]s"
  }`, zoneID, name, aopType)
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cert still exists")
		}
	}

	return nil
}
