package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
    certificate = "-----BEGIN CERTIFICATE-----\nMIIFWzCCBEOgAwIBAgISAw9GlDw8l9n12rmHAKjwRYtFMA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0xOTA3MTUyMzI1MzZaFw0x\nOTEwMTMyMzI1MzZaMBwxGjAYBgNVBAMTEXRmLnRlc3QuY2ZhcGkubmV0MIIBIjAN\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzRS91Cp5T/QmvxB5WoX/BFkf7KPU\noGZM5SpbmU/nKDSG+8H0yNOP2YycxJ1ZlWW6dNsgbBY1Y6XNvBYOUIDklNu56RZT\nGUfUMCQQHlqM04xCb1oZmnwcmqqEy3WxEHQHGNpfARmgD7prDP4XcFsxyLBCBO6U\n7/wbQ1SCP4LOlOi6EgpTYdFIBa8EU6ev0UbaAjXyF411fF4P9KgROonss6SjmSzs\nEro87JV80kMAarTaEenrJMYy+DkS55rEFtAshJ/ov3HvVBzgJl9tiVBxF/+xybM2\nulWR78mIi9byv5WozrdCFKUGxrUbPzrv4C/Zg9ptZMHr2P36thmgl1zAiwIDAQAB\no4ICZzCCAmMwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggr\nBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBSxzSYXcr7fLqOKXxKdX+v4\nWPGuBzAfBgNVHSMEGDAWgBSoSmpjBH3duubRObemRWXv86jsoTBvBggrBgEFBQcB\nAQRjMGEwLgYIKwYBBQUHMAGGImh0dHA6Ly9vY3NwLmludC14My5sZXRzZW5jcnlw\ndC5vcmcwLwYIKwYBBQUHMAKGI2h0dHA6Ly9jZXJ0LmludC14My5sZXRzZW5jcnlw\ndC5vcmcvMBwGA1UdEQQVMBOCEXRmLnRlc3QuY2ZhcGkubmV0MEwGA1UdIARFMEMw\nCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9j\ncHMubGV0c2VuY3J5cHQub3JnMIIBBQYKKwYBBAHWeQIEAgSB9gSB8wDxAHYA4mlL\nribo6UAJ6IYbtjuD1D7n/nSI+6SPKJMBnd3x2/4AAAFr+CucegAABAMARzBFAiEA\n9m+2odMVc4vungmCW10dFce5ot+bVxuRHkk28QoGh/gCIA8BNHm/6yGnUxjzn8J9\n1l3+QlrsZlYe+qS3qApK7EBKAHcAY/Lbzeg7zCzPC3KEJ1drM6SNYXePvXWmOLHH\naFRL2I0AAAFr+CuccgAABAMASDBGAiEAs79AHMEQs95IyfTyMGd2rnk46B/qQwYw\ntKxbFom3aC0CIQCXEXaHnKH21qa6SSFFd996BbY8ZNTAZNna23dIip8UxDANBgkq\nhkiG9w0BAQsFAAOCAQEAW+QluwC2Zxa6WM3bbMAUIAbdUe64nh9H5Ej/GC6niZKZ\nFVQt8nfNcoukmJbjl0YdxZi1QY5eenFg2eYkdy6Loj04Nf8GBEEciWODJaYKbieZ\nzSbVgKL75n7LUfBonavQAjoWIUdW88wL8790QVh0E3vloY3JVCApx0CuQKO0TeaR\nxIcP8QFh3A7EG3T0K5qdfxVTPnf9QLkiYeiXhfbz5EOwCfzA4xYsIDqc86rGFDuh\nFjR5aZV33sGQQFgYAXkHtT58z0sydbXgX4txRnGLQ6n0LlubeiAb0/RhiiHG+VFb\nAK5YZPYP+eod/41MtuKTBEHSrQmODefUZm5VUghJww==\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDNFL3UKnlP9Ca/\nEHlahf8EWR/so9SgZkzlKluZT+coNIb7wfTI04/ZjJzEnVmVZbp02yBsFjVjpc28\nFg5QgOSU27npFlMZR9QwJBAeWozTjEJvWhmafByaqoTLdbEQdAcY2l8BGaAPumsM\n/hdwWzHIsEIE7pTv/BtDVII/gs6U6LoSClNh0UgFrwRTp6/RRtoCNfIXjXV8Xg/0\nqBE6ieyzpKOZLOwSujzslXzSQwBqtNoR6eskxjL4ORLnmsQW0CyEn+i/ce9UHOAm\nX22JUHEX/7HJsza6VZHvyYiL1vK/lajOt0IUpQbGtRs/Ou/gL9mD2m1kwevY/fq2\nGaCXXMCLAgMBAAECggEBAIvvvkQ6o0KiV5oCNLxHOKcP5Y/Ejr7Qb2HkEFLByfqO\nNRku1MgATGTm5MXolIszuhIov6vhT5bqOUNBTY0zFkZY1DevSw6yC6C5yuHbacKk\nL2Tp9xSJ4b7L4gcvDJ4sffdAcpk+khCJZKid7QJ2x7aoRrQ01B4ZScUcsi+CI1JJ\nbkp3LvH+pYfz2/NXGwqsvgxKAIbMTsSv271BWTm0Pr/cQ3AECnKCWnS0/3LNy1SP\nt1xXAuL8vVfkl84RqTYws3cXQBvwo9KobqZFdNvIsGbyIdV4Z8T4KG4QzS2DBFrh\nFzOKBw+uvkBMjJCeBAd1Daha7ylJio+zOvN9j5RGNFkCgYEA7x8FcO4yO6bfTGrd\n/ocbaovSsABHkxmEZb6Ki7EPlQgmrjwmUaNUf/pwoW/RC58d+gHrcBkw8Y9DnUvL\njy9sljKJS+ApS16uILelQ556A0GCYLla5vJ/oDAuJ8kAtxCkPvvdFlHP6GDuODsY\nl6pauLp2VuoOa9IeVs8xlQPiwt0CgYEA246bMyC/gQYIEQj4sIqClUd60mMhjupN\nkbQdMsOjM8sBjf9s3ZIONWXI5UFE/QrdDdMsZBP4xUvX/CgcoXjxsRoIyWFR/vHC\n++o9yKj9XfFP5g41U7eOJj36GiVZ2s/J1qglJ44cFmIBaReGoXopS/ccOTcScDbK\nPOIOArcdFocCgYEAsmhhxeVig1E476oYYaxqTy9tjbVXsa/rMYJdmmYL6zS+r2bf\nbC/Bfw7a9AgaX2JjmkHOaL/S3Zf3aafAg99tVA72ky73gG1u26hJXM8j18QLw6Dn\n6sHpaRophbOZnfyDnx6J0PpPdeDEPB4Tdi07LPKqEqTlB5so2boTE0xn5t0CgYEA\nn9WbSodGoskfSjd7xBmxorccxNiB76bGvZGfx/sAbo4VHaibOlo/mcP1kmAHtycX\nch8Pq/OWIRtrqxgQb8S6PrGzP9dnd+/MgNQwEkpj2OX5woMJc16nT1PDJRGX7mFi\nkLBsC/W6oNjMKhOEYT2rnq/Qjh53f9WDOPtgM73WoTUCgYAnYT8lyRyRVrgm925h\nrzE1uHcK7k6ZgSsIHTy3HLi7GQTzIbAMDmz0ZrTZgyEnrvTSuvHklxreUrBBt1+W\n5HbyCTwlUxkKQcj7QX5IlKe1IpvhSjvplG6bPHvHQWPmnFL4Q0WPRwISWPcoP/Ih\nqHXCFswfipntRh3ghFnACt55xg==\n-----END PRIVATE KEY-----\n"
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
