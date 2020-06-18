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
    certificate = "-----BEGIN CERTIFICATE-----\nMIIFXTCCBEWgAwIBAgISA4vFrUgFW5lDUqt9Qba7adk5MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDAxMjIwMTIxNThaFw0y\nMDA0MjEwMTIxNThaMB4xHDAaBgNVBAMTE3RlcnJhZm9ybS5jZmFwaS5uZXQwggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC/Xq8xjLKyL5WY753wxyebrviv\ncr2z31A9mZNKJjBQfVGmUQVT1x37iqScmO1cpksCxeXGo3+faDDktsm+B01oMw3R\nxVjKlF9l8AI8WT7k8XyhPP4aOr5JWMsi4W9tqk3kJBAiaVMN0NzsQR4MQQhm4iaY\nc3/et/dCZW8MiSXzrvuEIlnORceJuCv7vRVszEjdyrZhkqwxieuJxovpaez5Rfwh\nmF6j0aQO21ZpbPFyvZlDv20LL8X6w5i+7KWuRtR5itBQPEhiDTKUu6zalI5UF1CO\n2p8BG97gqveZYW2fXSolgLthHgsu/rVwBniyzuRTi97vXlCS+6eUJa+fU/P7AgMB\nAAGjggJnMIICYzAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG\nCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFE06N06N95UiXBkr/wxT\nscpgXQgUMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUF\nBwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNy\neXB0Lm9yZy8wHgYDVR0RBBcwFYITdGVycmFmb3JtLmNmYXBpLm5ldDBMBgNVHSAE\nRTBDMAgGBmeBDAECATA3BgsrBgEEAYLfEwEBATAoMCYGCCsGAQUFBwIBFhpodHRw\nOi8vY3BzLmxldHNlbmNyeXB0Lm9yZzCCAQMGCisGAQQB1nkCBAIEgfQEgfEA7wB1\nAPCVpFnyANGCQBAtL5OIjq1L/h1H45nh0DSmsKiqjrJzAAABb8sObycAAAQDAEYw\nRAIgRS7tfvYzWxj0D+aO5tGRtU3nI+GC4m/XHU+onPJ0mKoCIBYD3aG7jPAV4go9\nptjwoh294YBd3xiv33OiX9KjSWspAHYAB7dcG+V9aP/xsMYdIxXHuuZXfFeUt2ru\nvGE6GmnTohwAAAFvyw5vcQAABAMARzBFAiAIqpecJ1UXw3mNGzCnrnPOUAirFwud\nn4pZkk32ykaxVwIhAPEmjAqDrkwtMJkK6dzEvg+rz0MMu+9KwMuTM8cx1ih0MA0G\nCSqGSIb3DQEBCwUAA4IBAQBEHlDZ9Doi/hLbhmw9JLcvPIPCjttjZ/noh167nJzo\nLCpd0//+XhxVml6RSa6sBhT9eiYpyEm4nbCowmcgbuYKDQkGkjyDnOO77XNSP8J2\nQ/RkAbQAR+VRMaiq46Hup/TiVdbqBaOmVRjeplj/7VLylvn8la0cpFe+KgZKlE4d\nd33H2umCPfnQPBDNSs7BAJ88cvN+C0DmNFUn7FMi9IWdtIHuZD36cbucysgKsqni\nHfVK80evEI7ZzSjGpFu0wg+XY1vUjDE9huHjbifMtAPR8jBizwh8RiWe14fnpsiu\nDsf6V17FSFPlmi/7NzaUuGykoCR9MiZdrGuU8mC4aTi9\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/Xq8xjLKyL5WY\n753wxyebrvivcr2z31A9mZNKJjBQfVGmUQVT1x37iqScmO1cpksCxeXGo3+faDDk\ntsm+B01oMw3RxVjKlF9l8AI8WT7k8XyhPP4aOr5JWMsi4W9tqk3kJBAiaVMN0Nzs\nQR4MQQhm4iaYc3/et/dCZW8MiSXzrvuEIlnORceJuCv7vRVszEjdyrZhkqwxieuJ\nxovpaez5RfwhmF6j0aQO21ZpbPFyvZlDv20LL8X6w5i+7KWuRtR5itBQPEhiDTKU\nu6zalI5UF1CO2p8BG97gqveZYW2fXSolgLthHgsu/rVwBniyzuRTi97vXlCS+6eU\nJa+fU/P7AgMBAAECggEAHBDbWcV8OaS/6F+QBs12ch8tqrGFv9kK8BXTY6cJI+zV\nKjKsuNiONaNmM+87tIBQ9PWoFsNIxsyliw2BteRlRlhiePbb6E3tVcpm0Yn3LuV5\ntT34OEmQObqThCiSyn8VEFX3pcxTmW2d0OpV6U8qV8hoB6i8wqGxWiP9LtX4Ym3w\ny0ydYz6QkTOZTV2I80BAGfy1NVDRov/9XqtZWvv63ro9SQMAseQqeeUSthV9Fjhp\nXVj/SUqJtwRQHNEKr4Fx+o/cKapv4id8N2mWjhfcICp5QndF1e9lfVtMngWVFOqf\ncY/89bF5mD8px8/oB7nbN+GOWTFsmihRwI2nA8djmQKBgQDm1BnU7T3Hs2UAbEtN\nP8X1qCt83mrVdb+hibRHl9qlM7+kbAo9JALyjEZ+Q9LTzxaEhGCbsevFqsY1up+W\nEfIC3ev50nFz1NnrXyKR5DC+87HCLE18ip42ZmKoiobTvZd7lRVWbX/rVgMdLemh\nCQuJCncuM6B2SDD0xaCMjkvaKQKBgQDUPQpFoA6Ys0hNllIQmftvcwUp+89z3sId\nm/VyiMqPVWqJHChvD/+VxLNueCNHvkYueEXHC/MZFNHm73Ef1+IXB/z/wuFV2PrN\ndvA3FXz8xloXh2m3NmNbuQ2KN/WXuXtO3Epac+R1if2UDPVBCkAkVOY7tPgYVe1d\nsuAvo9TpgwKBgQDGIuQ1iJtSUyPslAijO42yS3Ng0Q6FQniGscxE9A1jZyMmgPLc\n/o9lIZHVCmTrGUSr5XGD09qdJvTS4+neiHLjkRjgrYpjMh8I0fW7o1NQZaB9G0g8\nEkSyT3p6T8Zh9MkV9KeHM6DtIjy0DFgRudDkBk69IuWnAlq7kPQ6El772QKBgD0/\ntFQtgajdrFL/u4Ug+ufJ2Map7c9xjLGAzY+VHGfK7ajN4HlUs5ykHGgX8Y6Fwbkt\nam2r7Cbj1EOB/DKFWbDt1Dx3IBJnQNHErkQnRl+oWl2J7Z866eeReu/VgGGd3JEA\nj9CUu2yUOwLbzndLnwEdIyg97I8RVSQCOCJndE1DAoGAL3cZky8QTB8nTsQvchZu\nMPQXXDS7WyOgD8RKSKTU2UQtyVm/s7MnUzz5fRDhlLsoKLsw70HjFj4GZGHirGT5\nmTE1qRq0+SDR3OxI+Dn3Xatlpg7xLAAo3nfu+WjNdneBuJfdwko+MNGtbvXkDaAT\nVxHlMS7BpA/Y+FhCI1AH70o=\n-----END PRIVATE KEY-----\n"
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
