resource "cloudflare_certificate_pack" "%[3]s" {
  zone_id = "%[1]s"
  type = "advanced"
  hosts = [
    "%[2]s",
    "test.%[2]s"
  ]
  validation_method = "http"
  validity_days = 90
  certificate_authority = "lets_encrypt"
  cloudflare_branding = false
}