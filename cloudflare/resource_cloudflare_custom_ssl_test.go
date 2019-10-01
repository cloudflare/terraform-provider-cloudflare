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
    certificate = "-----BEGIN CERTIFICATE-----\nMIIFXjCCBEagAwIBAgISAympguRfAsX307ZikP7jl0dwMA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0xOTA4MDUxNDU4NTlaFw0x\nOTExMDMxNDU4NTlaMB4xHDAaBgNVBAMTE3RlcnJhZm9ybS5jZmFwaS5uZXQwggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDMmJ7n1plIwwuA45Q6GeAb13A6\n4/1rOUIHYp+wG/IRyUQijpoSldeYd9pzJOudB200tmNnmMmYIIT7xyewWcEw2LHP\nyEJ2vIuZnlGVDSL9PYrD3T24XRZk4A70wXr9FXxDHAIr+QgGY8YCk9Jv88ySIXIh\nxEf6w/5HYUR46sq+F97QA4w8OhcQCJ5t35ujt9LBxiDlBuh4vDuj4TgkQvtAwB2A\nY+sP+r9Taj5syX2fxGw/OwXqYn+JLoWo52s3790VG3sAUPZ2IoIz+rvs1MAkAer+\nicV6UGSRhBDifg6dZKJmlIJ+I2Jk049wI8imQGLFgM+FMWSoHkRCP9L+obTBAgMB\nAAGjggJoMIICZDAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG\nCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFLx2OSHdcqidsz29u2Ly\numoF9Wb0MB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUF\nBwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZy8wHgYDVR0RBBcwFYITdGVycmFmb3JtLmNmYXBpLm5ldDBMBgNVHSAE\nRTBDMAgGBmeBDAECATA3BgsrBgEEAYLfEwEBATAoMCYGCCsGAQUFBwIBFhpodHRw\nOi8vY3BzLmxldHNlbmNyeXB0Lm9yZzCCAQQGCisGAQQB1nkCBAIEgfUEgfIA8AB1\nAHR+2oMxrTMQkSGcziVPQnDCv/1eQiAIxjc1eeYQe8xWAAABbGKBV4oAAAQDAEYw\nRAIgfMVR6qQguG+a0LmXkDdAVGyVLafenFcJch7qVLPVJiACIEB6Utb1Cuts4Q3P\ndq2c7Srp2OWUwEzGUCjcoYduckg0AHcAKTxRllTIOWW6qlD8WAfUt2+/WHopctyk\nwwz05UVH9HgAAAFsYoFZnAAABAMASDBGAiEA7JnUEyzKzxWvleuyRbQO/e8FYijl\nM7uMRTkI9pbUGUgCIQCUSooINNU5zDYH0a+z/C1ubTGVo6edcj+mmsuqG37RIjAN\nBgkqhkiG9w0BAQsFAAOCAQEAfiBetJTZ51lfQ7GJSXCrejYpzDklz5OpbvR1uHb0\nqiP/G8trGqXyPyjhFGZlFMOag0VQYcsmhEvDCveV67bgziHRthCkPNXMveKHYDRw\njvifAOf1LSRWvzHEKsAyTb5s/qnSOnmH8U2bE2zn45W65ztYjQGJAx8MA908Dcrx\nRoMcVKqmfXVePNY3w8DCZJrc2O17Q8BABjaQvxm2HT48eSD8G+fUoaFIa6nu/FUs\nMustdz8mZg2boAZ1J0yNxa80Y1s44H3kOSYxfgfKYoe1+GGcUKbdM/EF7v4tzYNj\nT11Xy/i+xFc5fa+H2Gpycit3fdkj5TEVPX7ye4yHo2wtPg==\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDMmJ7n1plIwwuA\n45Q6GeAb13A64/1rOUIHYp+wG/IRyUQijpoSldeYd9pzJOudB200tmNnmMmYIIT7\nxyewWcEw2LHPyEJ2vIuZnlGVDSL9PYrD3T24XRZk4A70wXr9FXxDHAIr+QgGY8YC\nk9Jv88ySIXIhxEf6w/5HYUR46sq+F97QA4w8OhcQCJ5t35ujt9LBxiDlBuh4vDuj\n4TgkQvtAwB2AY+sP+r9Taj5syX2fxGw/OwXqYn+JLoWo52s3790VG3sAUPZ2IoIz\n+rvs1MAkAer+icV6UGSRhBDifg6dZKJmlIJ+I2Jk049wI8imQGLFgM+FMWSoHkRC\nP9L+obTBAgMBAAECggEAPs7voW6A2hR+eI/k1j1RTlrB6mJJTtxiB9BgA3lgw9MM\npqsuY1w6tmS83DJOXoOEI/WF6Ky/3oLFMGIALiQvqaYsWAQ7WyYgmQVAOEizIBj/\ne4d0xh9Vm5wpGzw2XHF3F0cG56borsV8aRgmNxYaDBZWakVOb44xhoo2sgQqP1aZ\nrSYUPmLfC2vDyIUiNWZ1f8IJH7IwcvLkaGSzSn8BW/euCz+4B+vz8x3smWKPFM0Q\nkXcL9ciJ1GuDpjhlg006MO0mtcvG76OHX+qPMny1sIFN4villOaHz4SFgYLGRB72\n7Lqyuga+eRrZFarquDkIxSzHdQkesnsK/pKt5GDdUQKBgQD3pXU7X1tDY88mODVg\n3yB/F1V9hgEQZ2N8SGsj8eOhq8AkZqzr9oog/fwYYcgxH1uDn2ESivK/ggGq66ge\nDq9cLzfvOyXaHyLR5DzkdXP3gLoDgH4MTrY9jGcP5Tf6eyUjI+qvDvb+A4qOkVlN\nCH9VA/4zCkM4GiYyWSjw1jpIpQKBgQDTf2eUAMBfRGjQZedfjAxxRb7x6K7VXlME\nW4LBnRZaPnlmm0oPbeDmNlvxmi1BXP8Js2nbY8GmCGe0EtFKZB4qxy5HxbJVzCU5\npt2ZnOp4limGwkEBb/SjdoII0HsNfjddnwvpYx4Vaj0uHH2zF6SqenNlzmg3UwvP\n4FbXr1xk7QKBgQCRKOQ5xCBLtSKEZagsOz3iITxUUosnIWM4Q37B2BS0/GapL6Im\nwiGfSyFM7WwaFyZeVbrh0p6N0NfHZ1DpJXR21Zq02PfMDjory9xBkfNC3aqrSNMZ\nxb2fAECdGaAha7OOEIyMxnnS1SKPhPVSaSuyGqATLO3P4cwH8SlFWl1ZnQKBgGLH\noYfVpgOYvt9+iMbucS1CZwEzLN0I1fs2BmcJSFRT032hz8BPEHhVMTIxUSuzFIbi\nXfGSsPIsAMtw8oEtK43NQ4dQBY/e7g/0KJHDYRt6/uAqwBO8x2TFR8x4GtDdf1xh\nmT2jBnz4BqUPt4G67DSXRmhpM/GK/vxTChxokd2tAoGAYX4s3UGcSQ8U/zxf07y5\n0Z+/5aMqobiUPwX3OAPrJ0MDDelyCtoZzaZ3OrI1/E6DQlUHAgdxvMC5OnP97IU7\nH8akcLR6bN6L9q1CinCJQMx+6fs78qenmaknvVFW5mk35tyEpS6yqauIf1Ddrqxa\n3MA/2EkI8J7LVX1KcuTUGTQ=\n-----END PRIVATE KEY-----\n"
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
